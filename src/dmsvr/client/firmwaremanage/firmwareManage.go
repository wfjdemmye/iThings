// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package firmwaremanage

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AccessAuthReq               = dm.AccessAuthReq
	DeviceCore                  = dm.DeviceCore
	DeviceGatewayBindDevice     = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq       = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp      = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiDeleteReq = dm.DeviceGatewayMultiDeleteReq
	DeviceGatewaySign           = dm.DeviceGatewaySign
	DeviceInfo                  = dm.DeviceInfo
	DeviceInfoCountReq          = dm.DeviceInfoCountReq
	DeviceInfoCountResp         = dm.DeviceInfoCountResp
	DeviceInfoDeleteReq         = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq          = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp         = dm.DeviceInfoIndexResp
	DeviceInfoReadReq           = dm.DeviceInfoReadReq
	DeviceRegisterReq           = dm.DeviceRegisterReq
	DeviceRegisterResp          = dm.DeviceRegisterResp
	DeviceTypeCountReq          = dm.DeviceTypeCountReq
	DeviceTypeCountResp         = dm.DeviceTypeCountResp
	Firmware                    = dm.Firmware
	FirmwareInfo                = dm.FirmwareInfo
	FirmwareInfoDeleteReq       = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp      = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq        = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp       = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq         = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp        = dm.FirmwareInfoReadResp
	FirmwareResp                = dm.FirmwareResp
	GroupDeviceIndexReq         = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp        = dm.GroupDeviceIndexResp
	GroupDeviceMultiCreateReq   = dm.GroupDeviceMultiCreateReq
	GroupDeviceMultiDeleteReq   = dm.GroupDeviceMultiDeleteReq
	GroupInfo                   = dm.GroupInfo
	GroupInfoCreateReq          = dm.GroupInfoCreateReq
	GroupInfoDeleteReq          = dm.GroupInfoDeleteReq
	GroupInfoIndexReq           = dm.GroupInfoIndexReq
	GroupInfoIndexResp          = dm.GroupInfoIndexResp
	GroupInfoReadReq            = dm.GroupInfoReadReq
	GroupInfoUpdateReq          = dm.GroupInfoUpdateReq
	LoginAuthReq                = dm.LoginAuthReq
	OtaCommonResp               = dm.OtaCommonResp
	OtaFirmwareDeviceInfoReq    = dm.OtaFirmwareDeviceInfoReq
	OtaFirmwareDeviceInfoResp   = dm.OtaFirmwareDeviceInfoResp
	OtaFirmwareFile             = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq     = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp    = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo         = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq          = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp         = dm.OtaFirmwareFileResp
	OtaPageInfo                 = dm.OtaPageInfo
	OtaTaskBatchReq             = dm.OtaTaskBatchReq
	OtaTaskBatchResp            = dm.OtaTaskBatchResp
	OtaTaskCancleReq            = dm.OtaTaskCancleReq
	OtaTaskCreatResp            = dm.OtaTaskCreatResp
	OtaTaskCreateReq            = dm.OtaTaskCreateReq
	OtaTaskDeviceCancleReq      = dm.OtaTaskDeviceCancleReq
	OtaTaskDeviceIndexReq       = dm.OtaTaskDeviceIndexReq
	OtaTaskDeviceIndexResp      = dm.OtaTaskDeviceIndexResp
	OtaTaskDeviceInfo           = dm.OtaTaskDeviceInfo
	OtaTaskDeviceProcessReq     = dm.OtaTaskDeviceProcessReq
	OtaTaskDeviceReadReq        = dm.OtaTaskDeviceReadReq
	OtaTaskIndexReq             = dm.OtaTaskIndexReq
	OtaTaskIndexResp            = dm.OtaTaskIndexResp
	OtaTaskInfo                 = dm.OtaTaskInfo
	OtaTaskReadReq              = dm.OtaTaskReadReq
	OtaTaskReadResp             = dm.OtaTaskReadResp
	PageInfo                    = dm.PageInfo
	PageInfo_OrderBy            = dm.PageInfo_OrderBy
	Point                       = dm.Point
	ProductCustom               = dm.ProductCustom
	ProductCustomReadReq        = dm.ProductCustomReadReq
	ProductInfo                 = dm.ProductInfo
	ProductInfoDeleteReq        = dm.ProductInfoDeleteReq
	ProductInfoIndexReq         = dm.ProductInfoIndexReq
	ProductInfoIndexResp        = dm.ProductInfoIndexResp
	ProductInfoReadReq          = dm.ProductInfoReadReq
	ProductRemoteConfig         = dm.ProductRemoteConfig
	ProductSchemaCreateReq      = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq      = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq       = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp      = dm.ProductSchemaIndexResp
	ProductSchemaInfo           = dm.ProductSchemaInfo
	ProductSchemaTslImportReq   = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq     = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp    = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq      = dm.ProductSchemaUpdateReq
	RemoteConfigCreateReq       = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq        = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp       = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq     = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp    = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq      = dm.RemoteConfigPushAllReq
	Response                    = dm.Response
	RootCheckReq                = dm.RootCheckReq

	FirmwareManage interface {
		// 新增固件升级包
		FirmwareInfoCreate(ctx context.Context, in *Firmware, opts ...grpc.CallOption) (*FirmwareResp, error)
		FirmwareInfoUpdate(ctx context.Context, in *FirmwareInfo, opts ...grpc.CallOption) (*OtaCommonResp, error)
		FirmwareInfoDelete(ctx context.Context, in *FirmwareInfoDeleteReq, opts ...grpc.CallOption) (*FirmwareInfoDeleteResp, error)
		FirmwareInfoIndex(ctx context.Context, in *FirmwareInfoIndexReq, opts ...grpc.CallOption) (*FirmwareInfoIndexResp, error)
		FirmwareInfoRead(ctx context.Context, in *FirmwareInfoReadReq, opts ...grpc.CallOption) (*FirmwareInfoReadResp, error)
		// 附件信息更新
		OtaFirmwareFileUpdate(ctx context.Context, in *OtaFirmwareFileReq, opts ...grpc.CallOption) (*OtaFirmwareFileResp, error)
		// 附件列表搜索
		OtaFirmwareFileIndex(ctx context.Context, in *OtaFirmwareFileIndexReq, opts ...grpc.CallOption) (*OtaFirmwareFileIndexResp, error)
		// 获取固件包对应设备版本列表
		OtaFirmwareDeviceInfo(ctx context.Context, in *OtaFirmwareDeviceInfoReq, opts ...grpc.CallOption) (*OtaFirmwareDeviceInfoResp, error)
	}

	defaultFirmwareManage struct {
		cli zrpc.Client
	}

	directFirmwareManage struct {
		svcCtx *svc.ServiceContext
		svr    dm.FirmwareManageServer
	}
)

func NewFirmwareManage(cli zrpc.Client) FirmwareManage {
	return &defaultFirmwareManage{
		cli: cli,
	}
}

func NewDirectFirmwareManage(svcCtx *svc.ServiceContext, svr dm.FirmwareManageServer) FirmwareManage {
	return &directFirmwareManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 新增固件升级包
func (m *defaultFirmwareManage) FirmwareInfoCreate(ctx context.Context, in *Firmware, opts ...grpc.CallOption) (*FirmwareResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.FirmwareInfoCreate(ctx, in, opts...)
}

// 新增固件升级包
func (d *directFirmwareManage) FirmwareInfoCreate(ctx context.Context, in *Firmware, opts ...grpc.CallOption) (*FirmwareResp, error) {
	return d.svr.FirmwareInfoCreate(ctx, in)
}

func (m *defaultFirmwareManage) FirmwareInfoUpdate(ctx context.Context, in *FirmwareInfo, opts ...grpc.CallOption) (*OtaCommonResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.FirmwareInfoUpdate(ctx, in, opts...)
}

func (d *directFirmwareManage) FirmwareInfoUpdate(ctx context.Context, in *FirmwareInfo, opts ...grpc.CallOption) (*OtaCommonResp, error) {
	return d.svr.FirmwareInfoUpdate(ctx, in)
}

func (m *defaultFirmwareManage) FirmwareInfoDelete(ctx context.Context, in *FirmwareInfoDeleteReq, opts ...grpc.CallOption) (*FirmwareInfoDeleteResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.FirmwareInfoDelete(ctx, in, opts...)
}

func (d *directFirmwareManage) FirmwareInfoDelete(ctx context.Context, in *FirmwareInfoDeleteReq, opts ...grpc.CallOption) (*FirmwareInfoDeleteResp, error) {
	return d.svr.FirmwareInfoDelete(ctx, in)
}

func (m *defaultFirmwareManage) FirmwareInfoIndex(ctx context.Context, in *FirmwareInfoIndexReq, opts ...grpc.CallOption) (*FirmwareInfoIndexResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.FirmwareInfoIndex(ctx, in, opts...)
}

func (d *directFirmwareManage) FirmwareInfoIndex(ctx context.Context, in *FirmwareInfoIndexReq, opts ...grpc.CallOption) (*FirmwareInfoIndexResp, error) {
	return d.svr.FirmwareInfoIndex(ctx, in)
}

func (m *defaultFirmwareManage) FirmwareInfoRead(ctx context.Context, in *FirmwareInfoReadReq, opts ...grpc.CallOption) (*FirmwareInfoReadResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.FirmwareInfoRead(ctx, in, opts...)
}

func (d *directFirmwareManage) FirmwareInfoRead(ctx context.Context, in *FirmwareInfoReadReq, opts ...grpc.CallOption) (*FirmwareInfoReadResp, error) {
	return d.svr.FirmwareInfoRead(ctx, in)
}

// 附件信息更新
func (m *defaultFirmwareManage) OtaFirmwareFileUpdate(ctx context.Context, in *OtaFirmwareFileReq, opts ...grpc.CallOption) (*OtaFirmwareFileResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.OtaFirmwareFileUpdate(ctx, in, opts...)
}

// 附件信息更新
func (d *directFirmwareManage) OtaFirmwareFileUpdate(ctx context.Context, in *OtaFirmwareFileReq, opts ...grpc.CallOption) (*OtaFirmwareFileResp, error) {
	return d.svr.OtaFirmwareFileUpdate(ctx, in)
}

// 附件列表搜索
func (m *defaultFirmwareManage) OtaFirmwareFileIndex(ctx context.Context, in *OtaFirmwareFileIndexReq, opts ...grpc.CallOption) (*OtaFirmwareFileIndexResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.OtaFirmwareFileIndex(ctx, in, opts...)
}

// 附件列表搜索
func (d *directFirmwareManage) OtaFirmwareFileIndex(ctx context.Context, in *OtaFirmwareFileIndexReq, opts ...grpc.CallOption) (*OtaFirmwareFileIndexResp, error) {
	return d.svr.OtaFirmwareFileIndex(ctx, in)
}

// 获取固件包对应设备版本列表
func (m *defaultFirmwareManage) OtaFirmwareDeviceInfo(ctx context.Context, in *OtaFirmwareDeviceInfoReq, opts ...grpc.CallOption) (*OtaFirmwareDeviceInfoResp, error) {
	client := dm.NewFirmwareManageClient(m.cli.Conn())
	return client.OtaFirmwareDeviceInfo(ctx, in, opts...)
}

// 获取固件包对应设备版本列表
func (d *directFirmwareManage) OtaFirmwareDeviceInfo(ctx context.Context, in *OtaFirmwareDeviceInfoReq, opts ...grpc.CallOption) (*OtaFirmwareDeviceInfoResp, error) {
	return d.svr.OtaFirmwareDeviceInfo(ctx, in)
}
