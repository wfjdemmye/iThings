package svc

import (
	"context"
	"github.com/i-Things/things/src/timed/timedjobsvr/client/timedmanage"
	"github.com/i-Things/things/src/timed/timedjobsvr/timedjobdirect"
	"os"

	"github.com/i-Things/things/shared/conf"
	"github.com/i-Things/things/shared/domain/schema"
	"github.com/i-Things/things/shared/oss"
	"github.com/i-Things/things/shared/stores"
	"github.com/i-Things/things/src/disvr/internal/config"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg/msgHubLog"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg/msgSdkLog"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg/msgThing"
	"github.com/i-Things/things/src/disvr/internal/repo/event/publish/pubApp"
	"github.com/i-Things/things/src/disvr/internal/repo/event/publish/pubDev"
	"github.com/i-Things/things/src/disvr/internal/repo/relationDB"
	"github.com/i-Things/things/src/disvr/internal/repo/tdengine/hubLogRepo"
	"github.com/i-Things/things/src/disvr/internal/repo/tdengine/schemaDataRepo"
	"github.com/i-Things/things/src/disvr/internal/repo/tdengine/sdkLogRepo"
	devicemanage "github.com/i-Things/things/src/dmsvr/client/devicemanage"
	"github.com/i-Things/things/src/dmsvr/client/firmwaremanage"
	"github.com/i-Things/things/src/dmsvr/client/otataskmanage"
	productmanage "github.com/i-Things/things/src/dmsvr/client/productmanage"
	remoteconfig "github.com/i-Things/things/src/dmsvr/client/remoteconfig"
	"github.com/i-Things/things/src/dmsvr/dmdirect"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	PubDev        pubDev.PubDev
	PubApp        pubApp.PubApp
	SchemaMsgRepo msgThing.SchemaDataRepo
	HubLogRepo    msgHubLog.HubLogRepo
	SchemaRepo    schema.ReadRepo
	SDKLogRepo    msgSdkLog.SDKLogRepo
	DeviceM       devicemanage.DeviceManage
	ProductM      productmanage.ProductManage
	RemoteConfig  remoteconfig.RemoteConfig
	Cache         kv.Store
	OssClient     *oss.Client
	FirmwareM     firmwaremanage.FirmwareManage
	OtaTaskM      otataskmanage.OtaTaskManage
	TimedM        timedmanage.TimedManage
}

func NewServiceContext(c config.Config) *ServiceContext {
	var (
		deviceM      devicemanage.DeviceManage
		productM     productmanage.ProductManage
		remoteConfig remoteconfig.RemoteConfig
		otataskM     otataskmanage.OtaTaskManage
		firmwareM    firmwaremanage.FirmwareManage
		timedM       timedmanage.TimedManage
	)

	hubLog := hubLogRepo.NewHubLogRepo(c.TDengine.DataSource)
	sdkLog := sdkLogRepo.NewSDKLogRepo(c.TDengine.DataSource)
	stores.InitConn(c.Database)
	err := relationDB.Migrate(c.Database)
	if err != nil {
		logx.Error("disvr 数据库初始化失败 err", err)
		os.Exit(-1)
	}
	pd, err := pubDev.NewPubDev(c.Event)
	if err != nil {
		logx.Error("NewPubDev err", err)
		os.Exit(-1)
	}
	pa, err := pubApp.NewPubApp(c.Event)
	if err != nil {
		logx.Error("NewPubApp err", err)
		os.Exit(-1)
	}
	if c.DmRpc.Mode == conf.ClientModeGrpc {
		deviceM = devicemanage.NewDeviceManage(zrpc.MustNewClient(c.DmRpc.Conf))
		productM = productmanage.NewProductManage(zrpc.MustNewClient(c.DmRpc.Conf))
		remoteConfig = remoteconfig.NewRemoteConfig(zrpc.MustNewClient(c.DmRpc.Conf))
		otataskM = otataskmanage.NewOtaTaskManage(zrpc.MustNewClient(c.DmRpc.Conf))
		firmwareM = firmwaremanage.NewFirmwareManage(zrpc.MustNewClient(c.DmRpc.Conf))
	} else {
		deviceM = dmdirect.NewDeviceManage(c.DmRpc.RunProxy)
		productM = dmdirect.NewProductManage(c.DmRpc.RunProxy)
		remoteConfig = dmdirect.NewRemoteConfig(c.DmRpc.RunProxy)
		otataskM = dmdirect.NewOtaTaskManage(c.DmRpc.RunProxy)
		firmwareM = dmdirect.NewFirmwareManage(c.DmRpc.RunProxy)
	}
	if c.TimedJobRpc.Mode == conf.ClientModeGrpc {
		timedM = timedmanage.NewTimedManage(zrpc.MustNewClient(c.TimedJobRpc.Conf))
	} else {
		timedM = timedjobdirect.NewTimedJob(c.TimedJobRpc.RunProxy)
	}
	tr := schema.NewReadRepo(func(ctx context.Context, productID string) (*schema.Model, error) {
		info, err := productM.ProductSchemaTslRead(ctx, &dm.ProductSchemaTslReadReq{ProductID: productID})
		if err != nil {
			return nil, err
		}
		return schema.ValidateWithFmt([]byte(info.Tsl))
	})
	cache := kv.NewStore(c.CacheRedis)
	ossClient := oss.NewOssClient(c.OssConf)
	if ossClient == nil {
		logx.Error("NewOss err")
		os.Exit(-1)
	}
	deviceData := schemaDataRepo.NewSchemaDataRepo(c.TDengine.DataSource, tr.GetSchemaModel, cache)
	return &ServiceContext{
		PubApp:        pa,
		Config:        c,
		SchemaRepo:    tr,
		PubDev:        pd,
		Cache:         cache,
		SchemaMsgRepo: deviceData,
		HubLogRepo:    hubLog,
		SDKLogRepo:    sdkLog,
		ProductM:      productM,
		DeviceM:       deviceM,
		RemoteConfig:  remoteConfig,
		OssClient:     ossClient,
		OtaTaskM:      otataskM,
		FirmwareM:     firmwareM,
		TimedM:        timedM,
	}
}
