package dmdirect

import (
	"fmt"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/src/dmsvr/internal/config"
	deviceauth "github.com/i-Things/things/src/dmsvr/internal/server/deviceauth"
	devicegroup "github.com/i-Things/things/src/dmsvr/internal/server/devicegroup"
	deviceinteract "github.com/i-Things/things/src/dmsvr/internal/server/deviceinteract"
	devicemanage "github.com/i-Things/things/src/dmsvr/internal/server/devicemanage"
	devicemsg "github.com/i-Things/things/src/dmsvr/internal/server/devicemsg"
	firmwaremanage "github.com/i-Things/things/src/dmsvr/internal/server/firmwaremanage"
	otataskmanage "github.com/i-Things/things/src/dmsvr/internal/server/otataskmanage"
	productmanage "github.com/i-Things/things/src/dmsvr/internal/server/productmanage"
	remoteconfig "github.com/i-Things/things/src/dmsvr/internal/server/remoteconfig"
	"github.com/i-Things/things/src/dmsvr/internal/startup"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"sync"
)

type Config = config.Config

var (
	svcCtx     *svc.ServiceContext
	svcOnce    sync.Once
	runSvrOnce sync.Once
	c          config.Config
)

func GetSvcCtx() *svc.ServiceContext {
	svcOnce.Do(func() {
		conf.MustLoad("etc/dm.yaml", &c)
		svcCtx = svc.NewServiceContext(c)
		startup.Init(svcCtx)
		logx.Infof("enabled dmsvr")
	})
	return svcCtx
}

// RunServer 如果是直连模式,同时提供Grpc的能力
func RunServer(svcCtx *svc.ServiceContext) {
	runSvrOnce.Do(func() {
		go Run(svcCtx)
	})

}
func Run(svcCtx *svc.ServiceContext) {
	c := svcCtx.Config
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		dm.RegisterDeviceAuthServer(grpcServer, deviceauth.NewDeviceAuthServer(svcCtx))
		dm.RegisterDeviceManageServer(grpcServer, devicemanage.NewDeviceManageServer(svcCtx))
		dm.RegisterProductManageServer(grpcServer, productmanage.NewProductManageServer(svcCtx))
		dm.RegisterOtaTaskManageServer(grpcServer, otataskmanage.NewOtaTaskManageServer(svcCtx))
		dm.RegisterFirmwareManageServer(grpcServer, firmwaremanage.NewFirmwareManageServer(svcCtx))
		dm.RegisterRemoteConfigServer(grpcServer, remoteconfig.NewRemoteConfigServer(svcCtx))
		dm.RegisterDeviceGroupServer(grpcServer, devicegroup.NewDeviceGroupServer(svcCtx))
		dm.RegisterDeviceInteractServer(grpcServer, deviceinteract.NewDeviceInteractServer(svcCtx))
		dm.RegisterDeviceMsgServer(grpcServer, devicemsg.NewDeviceMsgServer(svcCtx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddUnaryInterceptors(errors.ErrorInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
