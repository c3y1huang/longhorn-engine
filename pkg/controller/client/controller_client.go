package client

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/longhorn/longhorn-engine/pkg/meta"
	"github.com/longhorn/longhorn-engine/pkg/types"
	"github.com/longhorn/longhorn-engine/pkg/util"
	"github.com/longhorn/longhorn-engine/proto/ptypes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type ControllerClient struct {
	grpcAddress string
}

const (
	GRPCServiceTimeout = 1 * time.Minute
)

// NewControllerClient returns new ControllerClient
func NewControllerClient(address string) *ControllerClient {
	return &ControllerClient{
		grpcAddress: util.GetGRPCAddress(address),
	}
}

// GetVolumeInfo returns new types.VolumeInfo object
func GetVolumeInfo(v *ptypes.Volume) *types.VolumeInfo {
	return &types.VolumeInfo{
		Name:                  v.Name,
		Size:                  v.Size,
		ReplicaCount:          int(v.ReplicaCount),
		Endpoint:              v.Endpoint,
		Frontend:              v.Frontend,
		FrontendState:         v.FrontendState,
		IsExpanding:           v.IsExpanding,
		LastExpansionError:    v.LastExpansionError,
		LastExpansionFailedAt: v.LastExpansionFailedAt,
	}
}

// GetControllerReplicaInfo returns new types.ControllerReplicaInfo object
func GetControllerReplicaInfo(cr *ptypes.ControllerReplica) *types.ControllerReplicaInfo {
	return &types.ControllerReplicaInfo{
		Address: cr.Address.Address,
		Mode:    types.Mode(cr.Mode.String()),
	}
}

// GetControllerReplica returns new ptypes.ControllerReplica object
func GetControllerReplica(r *types.ControllerReplicaInfo) *ptypes.ControllerReplica {
	return &ptypes.ControllerReplica{
		Address: &ptypes.ReplicaAddress{
			Address: r.Address,
		},
		Mode: ptypes.ReplicaModeToGRPCReplicaMode(r.Mode),
	}
}

// GetSyncFileInfoList returns a list of types.SyncFileInfo
func GetSyncFileInfoList(list []*ptypes.SyncFileInfo) []types.SyncFileInfo {
	res := []types.SyncFileInfo{}
	for _, info := range list {
		res = append(res, GetSyncFileInfo(info))
	}
	return res
}

// GetSyncFileInfo returns a types.SyncFileInfo object
func GetSyncFileInfo(info *ptypes.SyncFileInfo) types.SyncFileInfo {
	return types.SyncFileInfo{
		FromFileName: info.FromFileName,
		ToFileName:   info.ToFileName,
		ActualSize:   info.ActualSize,
	}
}

// VolumeGet get volume with gRPC client and returns a new VolumeInfo object
func (c *ControllerClient) VolumeGet() (*types.VolumeInfo, error) {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	volume, err := controllerServiceClient.VolumeGet(ctx, &empty.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get volume %v: %v", c.grpcAddress, err)
	}

	return GetVolumeInfo(volume), nil
}

// VolumeStart start volume with gRPC client
func (c *ControllerClient) VolumeStart(replicas ...string) error {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	if _, err := controllerServiceClient.VolumeStart(ctx, &ptypes.VolumeStartRequest{
		ReplicaAddresses: replicas,
	}); err != nil {
		return fmt.Errorf("failed to start volume %v: %v", c.grpcAddress, err)
	}

	return nil
}

// VolumeSnapshot start volume with gRPC client
func (c *ControllerClient) VolumeSnapshot(name string, labels map[string]string) (string, error) {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return "", fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	reply, err := controllerServiceClient.VolumeSnapshot(ctx, &ptypes.VolumeSnapshotRequest{
		Name:   name,
		Labels: labels,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create snapshot %v for volume %v: %v", name, c.grpcAddress, err)
	}

	return reply.Name, nil
}

// VolumeRevert reverts volume with gRPC client
func (c *ControllerClient) VolumeRevert(snapshot string) error {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	if _, err := controllerServiceClient.VolumeRevert(ctx, &ptypes.VolumeRevertRequest{
		Name: snapshot,
	}); err != nil {
		return fmt.Errorf("failed to revert to snapshot %v for volume %v: %v", snapshot, c.grpcAddress, err)
	}

	return nil
}

// VolumeExpand expands volume with gRPC client
func (c *ControllerClient) VolumeExpand(size int64) error {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	if _, err := controllerServiceClient.VolumeExpand(ctx, &ptypes.VolumeExpandRequest{
		Size: size,
	}); err != nil {
		return fmt.Errorf("failed to expand to size %v for volume %v: %v", size, c.grpcAddress, err)
	}

	return nil
}

// VolumeFrontendStart starts frontend with gRPC client
func (c *ControllerClient) VolumeFrontendStart(frontend string) error {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	if _, err := controllerServiceClient.VolumeFrontendStart(ctx, &ptypes.VolumeFrontendStartRequest{
		Frontend: frontend,
	}); err != nil {
		return fmt.Errorf("failed to start frontend %v for volume %v: %v", frontend, c.grpcAddress, err)
	}

	return nil
}

// VolumeFrontendShutdown shutdown frontend with gRPC client
func (c *ControllerClient) VolumeFrontendShutdown() error {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	if _, err := controllerServiceClient.VolumeFrontendShutdown(ctx, &empty.Empty{}); err != nil {
		return fmt.Errorf("failed to shutdown frontend for volume %v: %v", c.grpcAddress, err)
	}

	return nil
}

// ReplicaList gets replica list with gRPC client
func (c *ControllerClient) ReplicaList() ([]*types.ControllerReplicaInfo, error) {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	reply, err := controllerServiceClient.ReplicaList(ctx, &empty.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to list replicas for volume %v: %v", c.grpcAddress, err)
	}

	replicas := []*types.ControllerReplicaInfo{}
	for _, cr := range reply.Replicas {
		replicas = append(replicas, GetControllerReplicaInfo(cr))
	}

	return replicas, nil
}

// ReplicaGet gets replica with gRPC client
func (c *ControllerClient) ReplicaGet(address string) (*types.ControllerReplicaInfo, error) {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	cr, err := controllerServiceClient.ReplicaGet(ctx, &ptypes.ReplicaAddress{
		Address: address,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get replica %v for volume %v: %v", address, c.grpcAddress, err)
	}

	return GetControllerReplicaInfo(cr), nil
}

// ReplicaCreate creates replica with gRPC client
func (c *ControllerClient) ReplicaCreate(address string, snapshotRequired bool, mode types.Mode) (*types.ControllerReplicaInfo, error) {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	cr, err := controllerServiceClient.ControllerReplicaCreate(ctx, &ptypes.ControllerReplicaCreateRequest{
		Address:          address,
		SnapshotRequired: snapshotRequired,
		Mode:             ptypes.ReplicaModeToGRPCReplicaMode(mode),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create replica %v for volume %v: %v", address, c.grpcAddress, err)
	}

	return GetControllerReplicaInfo(cr), nil
}

// ReplicaDelete deletes replica with gRPC client
func (c *ControllerClient) ReplicaDelete(address string) error {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	if _, err := controllerServiceClient.ReplicaDelete(ctx, &ptypes.ReplicaAddress{
		Address: address,
	}); err != nil {
		return fmt.Errorf("failed to delete replica %v for volume %v: %v", address, c.grpcAddress, err)
	}

	return nil
}

// ReplicaUpdate updates replica with gRPC client
func (c *ControllerClient) ReplicaUpdate(replica *types.ControllerReplicaInfo) (*types.ControllerReplicaInfo, error) {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	cr, err := controllerServiceClient.ReplicaUpdate(ctx, GetControllerReplica(replica))
	if err != nil {
		return nil, fmt.Errorf("failed to update replica %v for volume %v: %v", replica.Address, c.grpcAddress, err)
	}

	return GetControllerReplicaInfo(cr), nil
}

// ReplicaPrepareRebuild prepare rebuild replica with gRPC client
func (c *ControllerClient) ReplicaPrepareRebuild(address string) ([]types.SyncFileInfo, error) {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	reply, err := controllerServiceClient.ReplicaPrepareRebuild(ctx, &ptypes.ReplicaAddress{
		Address: address,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to prepare rebuilding replica %v for volume %v: %v", address, c.grpcAddress, err)
	}

	return GetSyncFileInfoList(reply.SyncFileInfoList), nil
}

// ReplicaVerifyRebuild verify replica rebuilt with gRPC client
func (c *ControllerClient) ReplicaVerifyRebuild(address string) error {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	if _, err := controllerServiceClient.ReplicaVerifyRebuild(ctx, &ptypes.ReplicaAddress{
		Address: address,
	}); err != nil {
		return fmt.Errorf("failed to verify rebuilded replica %v for volume %v: %v", address, c.grpcAddress, err)
	}

	return nil
}

// JournalList get a list of volume journal with gRPC client
func (c *ControllerClient) JournalList(limit int) error {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	if _, err := controllerServiceClient.JournalList(ctx, &ptypes.JournalListRequest{
		Limit: int64(limit),
	}); err != nil {
		return fmt.Errorf("failed to list journal for volume %v: %v", c.grpcAddress, err)
	}

	return nil
}

// VersionDetailGet get version detail with gRPC client
func (c *ControllerClient) VersionDetailGet() (*meta.VersionOutput, error) {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	reply, err := controllerServiceClient.VersionDetailGet(ctx, &empty.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get version detail: %v", err)
	}

	return &meta.VersionOutput{
		Version:                 reply.Version.Version,
		GitCommit:               reply.Version.GitCommit,
		BuildDate:               reply.Version.BuildDate,
		CLIAPIVersion:           int(reply.Version.CliAPIVersion),
		CLIAPIMinVersion:        int(reply.Version.CliAPIMinVersion),
		ControllerAPIVersion:    int(reply.Version.ControllerAPIVersion),
		ControllerAPIMinVersion: int(reply.Version.ControllerAPIMinVersion),
		DataFormatVersion:       int(reply.Version.DataFormatVersion),
		DataFormatMinVersion:    int(reply.Version.DataFormatMinVersion),
	}, nil

}

// Check the health for gRPC controller server with gRPC client
func (c *ControllerClient) Check() error {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	healthCheckClient := healthpb.NewHealthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	reply, err := healthCheckClient.Check(ctx, &healthpb.HealthCheckRequest{
		Service: "",
	})
	if err != nil {
		return fmt.Errorf("failed to list journal for volume %v: %v", c.grpcAddress, err)
	}

	if reply.Status != healthpb.HealthCheckResponse_SERVING {
		return fmt.Errorf("gRPC controller server is not serving")
	}

	return nil
}

// BackupReplicaMappingCreate creates backup replica mapping with gRPC client
func (c *ControllerClient) BackupReplicaMappingCreate(backupID string, replicaAddress string) error {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	if _, err := controllerServiceClient.BackupReplicaMappingCreate(ctx, &ptypes.BackupReplicaMapping{
		Backup:         backupID,
		ReplicaAddress: replicaAddress,
	}); err != nil {
		return fmt.Errorf("failed to store replica %v for backup %v: %v", replicaAddress, backupID, err)
	}

	return nil
}

// BackupReplicaMappingGet gets the backup replica mapping with gRPC client
func (c *ControllerClient) BackupReplicaMappingGet() (map[string]string, error) {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	br, err := controllerServiceClient.BackupReplicaMappingGet(ctx, &empty.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get backup replica mapping: %v", err)
	}

	return br.BackupReplicaMap, nil
}

// BackupReplicaMappingDelete deletes the backup replica mapping with gRPC client
func (c *ControllerClient) BackupReplicaMappingDelete(backupID string) error {
	conn, err := grpc.Dial(c.grpcAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to ControllerService %v: %v", c.grpcAddress, err)
	}
	defer conn.Close()
	controllerServiceClient := ptypes.NewControllerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), GRPCServiceTimeout)
	defer cancel()

	if _, err := controllerServiceClient.BackupReplicaMappingDelete(ctx, &ptypes.BackupReplicaMappingDeleteRequest{
		Backup: backupID,
	}); err != nil {
		return fmt.Errorf("failed to delete backup %v: %v", backupID, err)
	}

	return nil
}
