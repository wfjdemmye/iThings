package remoteconfiglogic

import (
	"context"
	"encoding/json"
	"github.com/i-Things/things/shared/devices"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/internal/domain/deviceMsg"
	"github.com/i-Things/things/src/dmsvr/internal/domain/deviceMsg/msgRemoteConfig"
	"github.com/i-Things/things/src/dmsvr/internal/repo/relationDB"
	devicemanage "github.com/i-Things/things/src/dmsvr/internal/server/devicemanage"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoteConfigPushAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PrcDB *relationDB.ProductRemoteConfigRepo
}

func NewRemoteConfigPushAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoteConfigPushAllLogic {
	return &RemoteConfigPushAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		PrcDB:  relationDB.NewProductRemoteConfigRepo(ctx),
	}
}

func (l *RemoteConfigPushAllLogic) RemoteConfigPushAll(in *dm.RemoteConfigPushAllReq) (*dm.Response, error) {
	//1. 根据产品id获取配置json
	respConfig, err := l.PrcDB.FindOneByFilter(l.ctx, relationDB.RemoteConfigFilter{
		ProductID: in.ProductID,
	})
	if err != nil {
		l.Errorf("%s.RemoteConfigLastRead failure err:%v", utils.FuncName(), err)
		return nil, err
	}

	//2. 根据产品id获取产品下的所有设备信息
	respDevices, err := devicemanage.NewDeviceManageServer(l.svcCtx).DeviceInfoIndex(l.ctx, &dm.DeviceInfoIndexReq{
		ProductID: in.ProductID,
	})
	if err != nil {
		l.Errorf("%s.RemoteConfigLastRead failure err:%v", utils.FuncName(), err)
		return nil, err
	}

	//3. for循环所有设备发送消息给设备
	for _, v := range respDevices.List {
		resp := &msgRemoteConfig.RemoteConfigMsg{
			Method:  "push",
			Code:    0,
			Payload: respConfig.Content,
		}
		respBytes, _ := json.Marshal(resp)
		msg := deviceMsg.PublishMsg{
			Handle:     devices.Config,
			Type:       msgRemoteConfig.TypePush,
			Payload:    respBytes,
			Timestamp:  time.Now().UnixMilli(),
			ProductID:  v.ProductID,
			DeviceName: v.DeviceName,
		}
		er := l.svcCtx.PubDev.PublishToDev(l.ctx, &msg)
		if er != nil {
			l.Errorf("%s.PublishToDev failure err:%v", utils.FuncName(), er)
			return nil, errors.System.AddDetail(er)
		}
	}
	if err != nil {
		l.Errorf("RemoteConfigPushAll.DeviceRemoteConfigUpdate err=%+v", err)
		return nil, errors.System.AddDetail(err)
	}
	return &dm.Response{}, nil
}
