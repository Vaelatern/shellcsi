package driver

import (
	"context"
	"fmt"
	"io/fs"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Vaelatern/shellcsi/internal/script"
)

func notImplemented(method string) error {
	return status.Errorf(codes.Unimplemented, fmt.Sprintf("method %s not implemented", method))
}

type Config struct {
}

type Driver struct {
	// Safety embeds mandated by the compiler by requiring an unexported function
	csi.UnsafeIdentityServer
	csi.UnsafeNodeServer
	csi.UnsafeControllerServer

	name    string
	version string
	config  Config
	fs      fs.FS
}

func NewDriver(name, version string, c Config) *Driver {
	return &Driver{
		name:    name,
		version: version,
		config:  c,
	}
}

////////////////////////////////////////////
// Identity
////////////////////////////////////////////

func (d *Driver) GetPluginInfo(ctx context.Context, req *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	if !script.Exists("GetPluginInfo") {
		return nil, notImplemented("GetPluginInfo")
	}
	return script.RunViaJson[Config, *csi.GetPluginInfoRequest, *csi.GetPluginInfoResponse](ctx, d.config, req, "GetPluginInfo")
}

func (d *Driver) GetPluginCapabilities(ctx context.Context, req *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	if !script.Exists("GetPluginCapabilities") {
		return nil, notImplemented("GetPluginCapabilities")
	}
	return script.RunViaJson[Config, *csi.GetPluginCapabilitiesRequest, *csi.GetPluginCapabilitiesResponse](ctx, d.config, req, "GetPluginCapabilities")
}

func (d *Driver) Probe(ctx context.Context, req *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	if !script.Exists("Probe") {
		return nil, notImplemented("Probe")
	}
	return script.RunViaJson[Config, *csi.ProbeRequest, *csi.ProbeResponse](ctx, d.config, req, "Probe")
}

////////////////////////////////////////////
// Controller
////////////////////////////////////////////

func (d *Driver) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	if !script.Exists("CreateVolume") {
		return nil, notImplemented("CreateVolume")
	}
	return script.RunViaJson[Config, *csi.CreateVolumeRequest, *csi.CreateVolumeResponse](ctx, d.config, req, "CreateVolume")
}

func (d *Driver) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	if !script.Exists("DeleteVolume") {
		return nil, notImplemented("DeleteVolume")
	}
	return script.RunViaJson[Config, *csi.DeleteVolumeRequest, *csi.DeleteVolumeResponse](ctx, d.config, req, "DeleteVolume")
}

func (d *Driver) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	if !script.Exists("ControllerPublishVolume") {
		return nil, notImplemented("ControllerPublishVolume")
	}
	return script.RunViaJson[Config, *csi.ControllerPublishVolumeRequest, *csi.ControllerPublishVolumeResponse](ctx, d.config, req, "ControllerPublishVolume")
}

func (d *Driver) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	if !script.Exists("ControllerUnpublishVolume") {
		return nil, notImplemented("ControllerUnpublishVolume")
	}
	return script.RunViaJson[Config, *csi.ControllerUnpublishVolumeRequest, *csi.ControllerUnpublishVolumeResponse](ctx, d.config, req, "ControllerUnpublishVolume")
}

func (d *Driver) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	if !script.Exists("ValidateVolumeCapabilities") {
		return nil, notImplemented("ValidateVolumeCapabilities")
	}
	return script.RunViaJson[Config, *csi.ValidateVolumeCapabilitiesRequest, *csi.ValidateVolumeCapabilitiesResponse](ctx, d.config, req, "ValidateVolumeCapabilities")
}

func (d *Driver) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	if !script.Exists("ListVolumes") {
		return nil, notImplemented("ListVolumes")
	}
	return script.RunViaJson[Config, *csi.ListVolumesRequest, *csi.ListVolumesResponse](ctx, d.config, req, "ListVolumes")
}

func (d *Driver) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	if !script.Exists("GetCapacity") {
		return nil, notImplemented("GetCapacity")
	}
	return script.RunViaJson[Config, *csi.GetCapacityRequest, *csi.GetCapacityResponse](ctx, d.config, req, "GetCapacity")
}

func (d *Driver) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	if !script.Exists("ControllerGetCapabilities") {
		return nil, notImplemented("ControllerGetCapabilities")
	}
	return script.RunViaJson[Config, *csi.ControllerGetCapabilitiesRequest, *csi.ControllerGetCapabilitiesResponse](ctx, d.config, req, "ControllerGetCapabilities")
}

func (d *Driver) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	if !script.Exists("CreateSnapshot") {
		return nil, notImplemented("CreateSnapshot")
	}
	return script.RunViaJson[Config, *csi.CreateSnapshotRequest, *csi.CreateSnapshotResponse](ctx, d.config, req, "CreateSnapshot")
}

func (d *Driver) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	if !script.Exists("DeleteSnapshot") {
		return nil, notImplemented("DeleteSnapshot")
	}
	return script.RunViaJson[Config, *csi.DeleteSnapshotRequest, *csi.DeleteSnapshotResponse](ctx, d.config, req, "DeleteSnapshot")
}

func (d *Driver) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	if !script.Exists("ListSnapshots") {
		return nil, notImplemented("ListSnapshots")
	}
	return script.RunViaJson[Config, *csi.ListSnapshotsRequest, *csi.ListSnapshotsResponse](ctx, d.config, req, "ListSnapshots")
}

func (d *Driver) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	if !script.Exists("ControllerExpandVolume") {
		return nil, notImplemented("ControllerExpandVolume")
	}
	return script.RunViaJson[Config, *csi.ControllerExpandVolumeRequest, *csi.ControllerExpandVolumeResponse](ctx, d.config, req, "ControllerExpandVolume")
}

func (d *Driver) ControllerGetVolume(ctx context.Context, req *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	if !script.Exists("ControllerGetVolume") {
		return nil, notImplemented("ControllerGetVolume")
	}
	return script.RunViaJson[Config, *csi.ControllerGetVolumeRequest, *csi.ControllerGetVolumeResponse](ctx, d.config, req, "ControllerGetVolume")
}

func (d *Driver) ControllerModifyVolume(ctx context.Context, req *csi.ControllerModifyVolumeRequest) (*csi.ControllerModifyVolumeResponse, error) {
	if !script.Exists("ControllerModifyVolume") {
		return nil, notImplemented("ControllerModifyVolume")
	}
	return script.RunViaJson[Config, *csi.ControllerModifyVolumeRequest, *csi.ControllerModifyVolumeResponse](ctx, d.config, req, "ControllerModifyVolume")
}

////////////////////////////////////////////
// Node
////////////////////////////////////////////

func (d *Driver) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	if !script.Exists("NodeStageVolume") {
		return nil, notImplemented("NodeStageVolume")
	}
	return script.RunViaJson[Config, *csi.NodeStageVolumeRequest, *csi.NodeStageVolumeResponse](ctx, d.config, req, "NodeStageVolume")
}

func (d *Driver) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	if !script.Exists("NodeUnstageVolume") {
		return nil, notImplemented("NodeUnstageVolume")
	}
	return script.RunViaJson[Config, *csi.NodeUnstageVolumeRequest, *csi.NodeUnstageVolumeResponse](ctx, d.config, req, "NodeUnstageVolume")
}

func (d *Driver) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	if !script.Exists("NodePublishVolume") {
		return nil, notImplemented("NodePublishVolume")
	}
	return script.RunViaJson[Config, *csi.NodePublishVolumeRequest, *csi.NodePublishVolumeResponse](ctx, d.config, req, "NodePublishVolume")
}

func (d *Driver) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	if !script.Exists("NodeUnpublishVolume") {
		return nil, notImplemented("NodeUnpublishVolume")
	}
	return script.RunViaJson[Config, *csi.NodeUnpublishVolumeRequest, *csi.NodeUnpublishVolumeResponse](ctx, d.config, req, "NodeUnpublishVolume")
}

func (d *Driver) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	if !script.Exists("NodeGetCapabilities") {
		return nil, notImplemented("NodeGetCapabilities")
	}
	return script.RunViaJson[Config, *csi.NodeGetCapabilitiesRequest, *csi.NodeGetCapabilitiesResponse](ctx, d.config, req, "NodeGetCapabilities")
}

func (d *Driver) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	if !script.Exists("NodeGetInfo") {
		return nil, notImplemented("NodeGetInfo")
	}
	return script.RunViaJson[Config, *csi.NodeGetInfoRequest, *csi.NodeGetInfoResponse](ctx, d.config, req, "NodeGetInfo")
}

func (d *Driver) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	if !script.Exists("NodeGetVolumeStats") {
		return nil, notImplemented("NodeGetVolumeStats")
	}
	return script.RunViaJson[Config, *csi.NodeGetVolumeStatsRequest, *csi.NodeGetVolumeStatsResponse](ctx, d.config, req, "NodeGetVolumeStats")
}

func (d *Driver) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	if !script.Exists("NodeExpandVolume") {
		return nil, notImplemented("NodeExpandVolume")
	}
	return script.RunViaJson[Config, *csi.NodeExpandVolumeRequest, *csi.NodeExpandVolumeResponse](ctx, d.config, req, "NodeExpandVolume")
}
