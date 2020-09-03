package sync

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/longhorn/longhorn-engine/pkg/controller/client"
	"github.com/longhorn/longhorn-engine/pkg/replica"
	replicaClient "github.com/longhorn/longhorn-engine/pkg/replica/client"
	"github.com/longhorn/longhorn-engine/pkg/types"
)

const (
	VolumeHeadName = "volume-head"
)

type Task struct {
	client *client.ControllerClient
}

type TaskError struct {
	ReplicaErrors []ReplicaError
}

type ReplicaError struct {
	Address string
	Message string
}

type SnapshotPurgeStatus struct {
	Error     string `json:"error"`
	IsPurging bool   `json:"isPurging"`
	Progress  int    `json:"progress"`
	State     string `json:"state"`
}

type ReplicaRebuildStatus struct {
	Error              string `json:"error"`
	IsRebuilding       bool   `json:"isRebuilding"`
	Progress           int    `json:"progress"`
	State              string `json:"state"`
	FromReplicaAddress string `json:"fromReplicaAddress"`
}

// NewTaskError returns a new TaskError contains list of the given ReplicaError
func NewTaskError(res ...ReplicaError) *TaskError {
	return &TaskError{
		ReplicaErrors: append([]ReplicaError{}, res...),
	}
}

// Error returns errors in single string
func (e *TaskError) Error() string {
	var errs []string
	for _, re := range e.ReplicaErrors {
		errs = append(errs, re.Error())
	}

	if errs == nil {
		return "Unknown"
	}
	if len(errs) == 1 {
		return errs[0]
	}
	return strings.Join(errs, "; ")
}

// Append the given error to the TaskError.ReplicaErrors
func (e *TaskError) Append(re ReplicaError) {
	e.ReplicaErrors = append(e.ReplicaErrors, re)
}

// HasError returns the number of the ReplicaErrors
func (e *TaskError) HasError() bool {
	return len(e.ReplicaErrors) != 0
}

// NewReplicaError returns ReplicaError object
func NewReplicaError(address string, err error) ReplicaError {
	return ReplicaError{
		Address: address,
		Message: err.Error(),
	}
}

// Error returns message "address: error"
func (e ReplicaError) Error() string {
	return fmt.Sprintf("%v: %v", e.Address, e.Message)
}

// NewTask returns new Task contains the gRPC client for the given address
func NewTask(controller string) *Task {
	return &Task{
		client: client.NewControllerClient(controller),
	}
}

// DeleteSnapshot ensurea no replicas is rebuilding on client and mark all as
// removed
func (t *Task) DeleteSnapshot(snapshot string) error {
	var err error

	replicas, err := t.client.ReplicaList()
	if err != nil {
		return err
	}

	for _, r := range replicas {
		if ok, err := t.isRebuilding(r); err != nil {
			return err
		} else if ok {
			return fmt.Errorf("Can not remove a snapshot because %s is rebuilding", r.Address)
		}
	}

	for _, replica := range replicas {
		if err = t.markSnapshotAsRemoved(replica, snapshot); err != nil {
			return err
		}
	}

	return nil
}

// PurgeSnapshots purges all replica snapshots with gRPC client
func (t *Task) PurgeSnapshots(skip bool) error {
	replicas, err := t.client.ReplicaList()
	if err != nil {
		return err
	}

	for _, r := range replicas {
		if ok, err := t.isRebuilding(r); err != nil {
			return err
		} else if ok {
			return fmt.Errorf("cannot purge snapshots because %s is rebuilding", r.Address)
		}

		if ok, err := t.isPurging(r); err != nil {
			return err
		} else if ok {
			if skip {
				return nil
			}
			return fmt.Errorf("cannot purge snapshots because %s is already purging snapshots", r.Address)
		}
	}

	errorMap := sync.Map{}
	var wg sync.WaitGroup
	wg.Add(len(replicas))

	for _, r := range replicas {
		go func(rep *types.ControllerReplicaInfo) {
			defer wg.Done()

			repClient, err := replicaClient.NewReplicaClient(rep.Address)
			if err != nil {
				errorMap.Store(rep.Address, err)
				return
			}

			if err := repClient.SnapshotPurge(); err != nil {
				errorMap.Store(rep.Address, err)
				return
			}
		}(r)
	}

	wg.Wait()
	for _, r := range replicas {
		if v, ok := errorMap.Load(r.Address); ok {
			err = v.(error)
			if skip && types.IsAlreadyPurgingError(err) {
				continue
			}
			logrus.Errorf("replica %v failed to start snapshot purge: %v", r.Address, err)
		}
	}

	return nil
}

// PurgeSnapshotStatus get all replica snapshot purge status with gRPC client,
// returns replicaStatusMap with SnapshotPurgeStatus mapped to the replica
// address
func (t *Task) PurgeSnapshotStatus() (map[string]*SnapshotPurgeStatus, error) {
	replicaStatusMap := make(map[string]*SnapshotPurgeStatus)

	replicas, err := t.client.ReplicaList()
	if err != nil {
		return nil, err
	}

	for _, r := range replicas {
		if r.Mode == types.ERR {
			continue
		}
		repClient, err := replicaClient.NewReplicaClient(r.Address)
		if err != nil {
			return nil, err
		}

		status, err := repClient.SnapshotPurgeStatus()
		if err != nil {
			replicaStatusMap[r.Address] = &SnapshotPurgeStatus{
				Error: fmt.Sprintf("failed to get snapshot purge status of %v: %v", r.Address, err),
			}
			continue
		}
		replicaStatusMap[r.Address] = &SnapshotPurgeStatus{
			Error:     status.Error,
			IsPurging: status.IsPurging,
			Progress:  int(status.Progress),
			State:     status.State,
		}
	}

	return replicaStatusMap, nil
}

// getNameAndIndex returns the given snapshot name if it exist in chain
func getNameAndIndex(chain []string, snapshot string) (string, int) {
	index := find(chain, snapshot)
	if index < 0 {
		snapshot = fmt.Sprintf("volume-snap-%s.img", snapshot)
		index = find(chain, snapshot)
	}

	if index < 0 {
		return "", index
	}

	return snapshot, index
}

// isRebuilding returns true if replica is rebuilding
func (t *Task) isRebuilding(replicaInController *types.ControllerReplicaInfo) (bool, error) {
	repClient, err := replicaClient.NewReplicaClient(replicaInController.Address)
	if err != nil {
		return false, err
	}

	replica, err := repClient.GetReplica()
	if err != nil {
		return false, err
	}

	return replica.Rebuilding, nil
}

// isPurgin returns true if replica status is purging
func (t *Task) isPurging(replicaInController *types.ControllerReplicaInfo) (bool, error) {
	repClient, err := replicaClient.NewReplicaClient(replicaInController.Address)
	if err != nil {
		return false, err
	}

	status, err := repClient.SnapshotPurgeStatus()
	if err != nil {
		return false, err
	}

	return status.IsPurging, nil
}

// isDirty returns true if replica is dirty
func (t *Task) isDirty(replicaInController *types.ControllerReplicaInfo) (bool, error) {
	repClient, err := replicaClient.NewReplicaClient(replicaInController.Address)
	if err != nil {
		return false, err
	}

	replica, err := repClient.GetReplica()
	if err != nil {
		return false, err
	}

	return replica.Dirty, nil
}

// markSnapshotAsRemoved mark replica disk as removed with gRPC client
func (t *Task) markSnapshotAsRemoved(replicaInController *types.ControllerReplicaInfo, snapshot string) error {
	if replicaInController.Mode != types.RW {
		return fmt.Errorf("Can only mark snapshot as removed from replica in mode RW, got %s", replicaInController.Mode)
	}

	repClient, err := replicaClient.NewReplicaClient(replicaInController.Address)
	if err != nil {
		return err
	}

	if err := repClient.MarkDiskAsRemoved(snapshot); err != nil {
		return err
	}

	return nil
}

// find the index of the given item in list
func find(list []string, item string) int {
	for i, val := range list {
		if val == item {
			return i
		}
	}
	return -1
}

// AddRestoreReplica creates replica with gRPC cliet
func (t *Task) AddRestoreReplica(replica string) error {
	volume, err := t.client.VolumeGet()
	if err != nil {
		return err
	}

	if volume.ReplicaCount == 0 {
		return t.client.VolumeStart(replica)
	}

	if err := t.checkRestoreReplicaSize(replica, volume.Size); err != nil {
		return err
	}

	logrus.Infof("Adding restore replica %s in WO mode", replica)

	// The replica mode will become RW after the first restoration complete.
	// And the rebuilding flag in the replica server won't be set since this is not normal rebuilding.
	if _, err = t.client.ReplicaCreate(replica, false, types.WO); err != nil {
		return err
	}

	return nil
}

// checkRestoreReplicaSize get replica size with gRPC client, returns error if
// replicaSize not equal to the given volumeSize
func (t *Task) checkRestoreReplicaSize(address string, volumeSize int64) error {
	replicaCli, err := replicaClient.NewReplicaClient(address)
	if err != nil {
		return err
	}

	replicaInfo, err := replicaCli.GetReplica()
	if err != nil {
		return err
	}
	replicaSize, err := strconv.ParseInt(replicaInfo.Size, 10, 64)
	if err != nil {
		return err
	}
	if replicaSize != volumeSize {
		return fmt.Errorf("rebuilding replica size %v is not the same as volume size %v", replicaSize, volumeSize)
	}

	return nil
}

// VerifyRebuildReplica verified rebuild status with gRPC client for the given
// address
func (t *Task) VerifyRebuildReplica(address string) error {
	if err := t.client.ReplicaVerifyRebuild(address); err != nil {
		return err
	}
	return nil
}

// AddReplica sync file to the new replica
func (t *Task) AddReplica(replica string) error {
	volume, err := t.client.VolumeGet()
	if err != nil {
		return err
	}

	if volume.ReplicaCount == 0 {
		return t.client.VolumeStart(replica)
	}

	if err := t.checkAndExpandReplica(replica, volume.Size); err != nil {
		return err
	}

	if err := t.checkAndResetFailedRebuild(replica); err != nil {
		return err
	}

	logrus.Infof("Adding replica %s in WO mode", replica)
	_, err = t.client.ReplicaCreate(replica, true, types.WO)
	if err != nil {
		return err
	}

	_, toClient, fromAddress, _, err := t.getTransferClients(replica)
	if err != nil {
		return err
	}

	if err := toClient.SetRebuilding(true); err != nil {
		return err
	}

	resp, err := t.client.ReplicaPrepareRebuild(replica)
	if err != nil {
		return err
	}

	if checkIfVolumeHeadExists(resp) {
		return fmt.Errorf("sync file list shouldn't contain volume head")
	}

	if err = toClient.SyncFiles(fromAddress, resp); err != nil {
		return err
	}

	if err := t.reloadAndVerify(replica, toClient); err != nil {
		return err
	}

	return nil
}

// checkAndResetFailedRebuild open replica for the given address and set
// rebuilding
func (t *Task) checkAndResetFailedRebuild(address string) error {
	client, err := replicaClient.NewReplicaClient(address)
	if err != nil {
		return err
	}

	replica, err := client.GetReplica()
	if err != nil {
		return err
	}

	if replica.State == "closed" && replica.Rebuilding {
		if err := client.OpenReplica(); err != nil {
			return err
		}

		if err := client.SetRebuilding(false); err != nil {
			return err
		}

		return client.Close()
	}

	return nil
}

// checkAndExpandReplica checks the given volume size against the client
// replica size and expand if its greater
func (t *Task) checkAndExpandReplica(address string, size int64) error {
	client, err := replicaClient.NewReplicaClient(address)
	if err != nil {
		return err
	}

	replica, err := client.GetReplica()
	if err != nil {
		return err
	}

	replicaSize, err := strconv.ParseInt(replica.Size, 10, 64)
	if err != nil {
		return err
	}

	
	if replicaSize > size {
		return fmt.Errorf("cannot add a larger replica to the engine")
	} else if replicaSize < size {
		logrus.Infof("Prepare to expand new replica to size %v", size)
		needClose := false
		if replica.State == "closed" {
			if err := client.OpenReplica(); err != nil {
				return err
			}
			needClose = true
		}
		if _, err := client.ExpandReplica(size); err != nil {
			return err
		}
		if needClose {
			return client.Close()
		}
	}

	return nil
}

// reloadAndVerify checks replica numbers, enables rebuilded replica and updates
// the rebuilding to false
func (t *Task) reloadAndVerify(address string, repClient *replicaClient.ReplicaClient) error {
	_, err := repClient.ReloadReplica()
	if err != nil {
		return err
	}

	if err := t.client.ReplicaVerifyRebuild(address); err != nil {
		return err
	}

	if err := repClient.SetRebuilding(false); err != nil {
		return err
	}
	return nil
}

// checkIfVolumeHeadExists ensures the "volume-head" not exist in the given
// infoList
func checkIfVolumeHeadExists(infoList []types.SyncFileInfo) bool {
	// volume head has been synced by PrepareRebuild()
	for _, info := range infoList {
		if strings.Contains(info.FromFileName, VolumeHeadName) {
			return true
		}
	}
	return false
}

// getTransferClients returns a healthy replica as source and a WO replica as
// the target
func (t *Task) getTransferClients(address string) (
	*replicaClient.ReplicaClient, *replicaClient.ReplicaClient, string, string, error) {
	from, err := t.getFromReplica()
	if err != nil {
		return nil, nil, "", "", err
	}
	logrus.Infof("Using replica %s as the source for rebuild ", from.Address)

	fromClient, err := replicaClient.NewReplicaClient(from.Address)
	if err != nil {
		return nil, nil, "", "", err
	}

	to, err := t.getToReplica(address)
	if err != nil {
		return nil, nil, "", "", err
	}
	logrus.Infof("Using replica %s as the target for rebuild ", to.Address)

	toClient, err := replicaClient.NewReplicaClient(to.Address)
	if err != nil {
		return nil, nil, "", "", err
	}

	return fromClient, toClient, from.Address, to.Address, nil
}

// getFromReplica return the first found good replica (RW)
func (t *Task) getFromReplica() (*types.ControllerReplicaInfo, error) {
	replicas, err := t.client.ReplicaList()
	if err != nil {
		return &types.ControllerReplicaInfo{}, err
	}

	for _, r := range replicas {
		if r.Mode == types.RW {
			return r, nil
		}
	}

	return &types.ControllerReplicaInfo{}, fmt.Errorf("Failed to find good replica to copy from")
}

// getToReplica returns replica in WO for the given address
func (t *Task) getToReplica(address string) (*types.ControllerReplicaInfo, error) {
	replicas, err := t.client.ReplicaList()
	if err != nil {
		return &types.ControllerReplicaInfo{}, err
	}

	for _, r := range replicas {
		if r.Address == address {
			if r.Mode != types.WO {
				return &types.ControllerReplicaInfo{}, fmt.Errorf("Replica %s is not in mode WO got: %s", address, r.Mode)
			}
			return r, nil
		}
	}

	return &types.ControllerReplicaInfo{}, fmt.Errorf("Failed to find target replica to copy to")
}

// getNonBackingDisks returns object contains all replica maps to its name if 
// its disk name is not the BackingFile
func getNonBackingDisks(address string) (map[string]types.DiskInfo, error) {
	repClient, err := replicaClient.NewReplicaClient(address)
	if err != nil {
		return nil, err
	}

	r, err := repClient.GetReplica()
	if err != nil {
		return nil, err
	}

	disks := make(map[string]types.DiskInfo)
	for name, disk := range r.Disks {
		if name == r.BackingFile {
			continue
		}
		disks[name] = disk
	}

	return disks, err
}

// GetSnapshotsInfo populates the DiskInfo mapping to the snapshot name, all
// heads will be marked as volume-head
func GetSnapshotsInfo(replicas []*types.ControllerReplicaInfo) (outputDisks map[string]types.DiskInfo, err error) {
	defer func() {
		err = errors.Wrapf(err, "BUG: cannot get snapshot info")
	}()
	for _, r := range replicas {
		if r.Mode != types.RW {
			continue
		}

		disks, err := getNonBackingDisks(r.Address)
		if err != nil {
			return nil, err
		}

		newDisks := make(map[string]types.DiskInfo)
		for name, disk := range disks {
			snapshot := ""

			if !replica.IsHeadDisk(name) {
				snapshot, err = replica.GetSnapshotNameFromDiskName(name)
				if err != nil {
					return nil, err
				}
			} else {
				snapshot = VolumeHeadName
			}
			children := map[string]bool{}
			for childDisk := range disk.Children {
				child := ""
				if !replica.IsHeadDisk(childDisk) {
					child, err = replica.GetSnapshotNameFromDiskName(childDisk)
					if err != nil {
						return nil, err
					}
				} else {
					child = VolumeHeadName
				}
				children[child] = true
			}
			parent := ""
			if disk.Parent != "" {
				parent, err = replica.GetSnapshotNameFromDiskName(disk.Parent)
				if err != nil {
					return nil, err
				}
			}
			info := types.DiskInfo{
				Name:        snapshot,
				Parent:      parent,
				Removed:     disk.Removed,
				UserCreated: disk.UserCreated,
				Children:    children,
				Created:     disk.Created,
				Size:        disk.Size,
				Labels:      disk.Labels,
			}
			newDisks[snapshot] = info
		}
		// we treat the healthy replica with the most snapshots as the
		// source of the truth, since that means something are still in
		// progress and haven't completed yet.
		if len(newDisks) > len(outputDisks) {
			outputDisks = newDisks
		}
	}
	return outputDisks, nil
}

// StartWithReplicas Backend and Frontend Volumes streams
func (t *Task) StartWithReplicas(replicas []string) error {
	volume, err := t.client.VolumeGet()
	if err != nil {
		return err
	}

	if volume.ReplicaCount != 0 {
		return fmt.Errorf("cannot add multiple replicas if volume is already up")
	}

	return t.client.VolumeStart(replicas...)
}

// RebuildStatus returns replicaStatusMap contains all restoring replica for
// its rebuilding status mapping to the replica address
func (t *Task) RebuildStatus() (map[string]*ReplicaRebuildStatus, error) {
	replicaStatusMap := make(map[string]*ReplicaRebuildStatus)

	replicas, err := t.client.ReplicaList()
	if err != nil {
		return nil, err
	}

	for _, r := range replicas {
		if r.Mode != types.WO {
			continue
		}
		repClient, err := replicaClient.NewReplicaClient(r.Address)
		if err != nil {
			return nil, err
		}

		restoreStatus, err := repClient.RestoreStatus()
		if err != nil {
			return nil, errors.Wrapf(err, "failed to check the restore status before fetching the rebuild status")
		}
		if restoreStatus.DestFileName != "" {
			logrus.Debugf("Skip checking rebuild status since the volume is a restore/DR volume")
			return replicaStatusMap, nil
		}

		status, err := repClient.ReplicaRebuildStatus()
		if err != nil {
			replicaStatusMap[r.Address] = &ReplicaRebuildStatus{
				Error: fmt.Sprintf("failed to get replica rebuild status of %v: %v", r.Address, err),
			}
			continue
		}
		replicaStatusMap[r.Address] = &ReplicaRebuildStatus{
			Error:              status.Error,
			IsRebuilding:       status.IsRebuilding,
			Progress:           int(status.Progress),
			State:              status.State,
			FromReplicaAddress: status.FromReplicaAddress,
		}
	}

	return replicaStatusMap, nil
}
