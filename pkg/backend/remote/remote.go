package remote

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/longhorn/longhorn-engine/pkg/dataconn"
	"github.com/longhorn/longhorn-engine/pkg/replica/client"
	replicaClient "github.com/longhorn/longhorn-engine/pkg/replica/client"
	"github.com/longhorn/longhorn-engine/pkg/types"
	"github.com/longhorn/longhorn-engine/pkg/util"
	"github.com/longhorn/longhorn-engine/proto/ptypes"
)

var (
	pingInveral   = 2 * time.Second
	timeout       = 30 * time.Second
	requestBuffer = 1024
)

// New returns new Factory
func New() types.BackendFactory {
	return &Factory{}
}

// RevisionCounter object
type RevisionCounter struct {
	Counter int64 `json:"counter,string"`
}

// Factory object
type Factory struct {
}

// Remote object
type Remote struct {
	types.ReaderWriterAt
	name              string
	replicaServiceURL string
	closeChan         chan struct{}
	monitorChan       types.MonitorChannel
}

// Close calls replica gRPC client for ReplicaClose method
func (r *Remote) Close() error {
	logrus.Infof("Closing: %s", r.name)
	conn, err := grpc.Dial(r.replicaServiceURL, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ReplicaService %v: %v", r.replicaServiceURL, err)
	}
	defer conn.Close()
	replicaServiceClient := ptypes.NewReplicaServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), client.GRPCServiceCommonTimeout)
	defer cancel()

	if _, err := replicaServiceClient.ReplicaClose(ctx, &empty.Empty{}); err != nil {
		return fmt.Errorf("failed to close replica %v from remote: %v", r.replicaServiceURL, err)
	}

	return nil
}

func (r *Remote) open() error {
	logrus.Infof("Opening: %s", r.name)
	conn, err := grpc.Dial(r.replicaServiceURL, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ReplicaService %v: %v", r.replicaServiceURL, err)
	}
	defer conn.Close()
	replicaServiceClient := ptypes.NewReplicaServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), client.GRPCServiceCommonTimeout)
	defer cancel()

	if _, err := replicaServiceClient.ReplicaOpen(ctx, &empty.Empty{}); err != nil {
		return fmt.Errorf("failed to open replica %v from remote: %v", r.replicaServiceURL, err)
	}

	return nil
}

// Snapshot calls replica gRPC client ReplicaSnapshot
func (r *Remote) Snapshot(name string, userCreated bool, created string, labels map[string]string) error {
	logrus.Infof("Snapshot: %s %s UserCreated %v Created at %v, Labels %v",
		r.name, name, userCreated, created, labels)
	conn, err := grpc.Dial(r.replicaServiceURL, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ReplicaService %v: %v", r.replicaServiceURL, err)
	}
	defer conn.Close()
	replicaServiceClient := ptypes.NewReplicaServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), client.GRPCServiceCommonTimeout)
	defer cancel()

	if _, err := replicaServiceClient.ReplicaSnapshot(ctx, &ptypes.ReplicaSnapshotRequest{
		Name:        name,
		UserCreated: userCreated,
		Created:     created,
		Labels:      labels,
	}); err != nil {
		return fmt.Errorf("failed to snapshot replica %v from remote: %v", r.replicaServiceURL, err)
	}

	return nil
}

// Expand calls replica gRPC client ReplicaExpand
func (r *Remote) Expand(size int64) (err error) {
	logrus.Infof("Expand to size %v", size)
	defer func() {
		err = types.WrapError(err, "failed to expand replica %v from remote", r.replicaServiceURL)
	}()

	conn, err := grpc.Dial(r.replicaServiceURL, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ReplicaService %v: %v", r.replicaServiceURL, err)
	}
	defer conn.Close()
	replicaServiceClient := ptypes.NewReplicaServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), client.GRPCServiceCommonTimeout)
	defer cancel()

	if _, err := replicaServiceClient.ReplicaExpand(ctx, &ptypes.ReplicaExpandRequest{
		Size: size,
	}); err != nil {
		return types.UnmarshalGRPCError(err)
	}

	return nil
}

// SetRevisionCounter calls replica gRPC client RevisionCounterSet
func (r *Remote) SetRevisionCounter(counter int64) error {
	logrus.Infof("Set revision counter of %s to : %v", r.name, counter)

	conn, err := grpc.Dial(r.replicaServiceURL, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ReplicaService %v: %v", r.replicaServiceURL, err)
	}
	defer conn.Close()
	replicaServiceClient := ptypes.NewReplicaServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), client.GRPCServiceCommonTimeout)
	defer cancel()

	if _, err := replicaServiceClient.RevisionCounterSet(ctx, &ptypes.RevisionCounterSetRequest{
		Counter: counter,
	}); err != nil {
		return fmt.Errorf("failed to set revision counter to %d for replica %v from remote: %v", counter, r.replicaServiceURL, err)
	}

	return nil

}

// Size returns replica size from replica gRPC server
func (r *Remote) Size() (int64, error) {
	replicaInfo, err := r.info()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(replicaInfo.Size, 10, 0)
}

// SectorSize returns replica sector size from replica
// gRPC server
func (r *Remote) SectorSize() (int64, error) {
	replicaInfo, err := r.info()
	if err != nil {
		return 0, err
	}
	return replicaInfo.SectorSize, nil
}

// RemainSnapshots returns remain snapshot from replica
// gRPC server
func (r *Remote) RemainSnapshots() (int, error) {
	replicaInfo, err := r.info()
	if err != nil {
		return 0, err
	}
	if replicaInfo.State != "open" && replicaInfo.State != "dirty" && replicaInfo.State != "rebuilding" {
		return 0, fmt.Errorf("Invalid state %v for counting snapshots", replicaInfo.State)
	}
	return replicaInfo.RemainSnapshots, nil
}

// GetRevisionCounter returns revision counter
// from gRPC server
func (r *Remote) GetRevisionCounter() (int64, error) {
	replicaInfo, err := r.info()
	if err != nil {
		return 0, err
	}
	if replicaInfo.State != "open" && replicaInfo.State != "dirty" {
		return 0, fmt.Errorf("Invalid state %v for getting revision counter", replicaInfo.State)
	}
	return replicaInfo.RevisionCounter, nil
}

// info calls replica gRPC client ReplicaGet and returns
// a ReplicaInfo object
func (r *Remote) info() (*types.ReplicaInfo, error) {
	conn, err := grpc.Dial(r.replicaServiceURL, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("cannot connect to ReplicaService %v: %v", r.replicaServiceURL, err)
	}
	defer conn.Close()
	replicaServiceClient := ptypes.NewReplicaServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), client.GRPCServiceCommonTimeout)
	defer cancel()

	resp, err := replicaServiceClient.ReplicaGet(ctx, &empty.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get replica %v info from remote: %v", r.replicaServiceURL, err)
	}

	return replicaClient.GetReplicaInfo(resp.Replica), nil
}

// Create checks the replica.State from replica rGPC server and
// monitors the given address connection with a inifinit ping
// loop 
func (rf *Factory) Create(address string) (types.Backend, error) {
	logrus.Infof("Connecting to remote: %s", address)

	controlAddress, dataAddress, _, _, err := util.ParseAddresses(address)
	if err != nil {
		return nil, err
	}

	r := &Remote{
		name:              address,
		replicaServiceURL: controlAddress,
		// We don't want sender to wait for receiver, because receiver may
		// has been already notified
		closeChan:   make(chan struct{}, 5),
		monitorChan: make(types.MonitorChannel, 5),
	}

	replica, err := r.info()
	if err != nil {
		return nil, err
	}

	if replica.State != "closed" {
		return nil, fmt.Errorf("Replica must be closed, Can not add in state: %s", replica.State)
	}

	conn, err := net.Dial("tcp", dataAddress)
	if err != nil {
		return nil, err
	}

	dataConnClient := dataconn.NewClient(conn)
	r.ReaderWriterAt = dataConnClient

	if err := r.open(); err != nil {
		return nil, err
	}

	go r.monitorPing(dataConnClient)

	return r, nil
}

// monitorPing starts an infinit loop that pings the given
// client
func (r *Remote) monitorPing(client *dataconn.Client) {
	ticker := time.NewTicker(pingInveral)
	defer ticker.Stop()

	for {
		select {
		case <-r.closeChan:
			r.monitorChan <- nil
			return
		case <-ticker.C:
			if err := client.Ping(); err != nil {
				client.SetError(err)
				r.monitorChan <- err
				return
			}
		}
	}
}

// GetMonitorChannel returns Remote.monitorChan
func (r *Remote) GetMonitorChannel() types.MonitorChannel {
	return r.monitorChan
}

// StopMonitoring sents close to Remote channel
func (r *Remote) StopMonitoring() {
	r.closeChan <- struct{}{}
}
