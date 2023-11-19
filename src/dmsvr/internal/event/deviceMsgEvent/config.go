package deviceMsgEvent

import (
	"context"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/internal/domain/deviceMsg"
	remoteconfiglogic "github.com/i-Things/things/src/dmsvr/internal/logic/remoteconfig"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type ConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigLogic {
	return &ConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigLogic) Handle(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	l.Infof("%s req=%+v", utils.FuncName(), msg)

	//获取最新配置
	resp1, err := remoteconfiglogic.NewRemoteConfigLastReadLogic(l.ctx, l.svcCtx).RemoteConfigLastRead(&dm.RemoteConfigLastReadReq{
		ProductID: msg.ProductID,
	})

	resp := &deviceMsg.CommonMsg{
		Method:    deviceMsg.RemoteConfigReply,
		Timestamp: time.Now().UnixMilli(),
		Data:      resp1.Info.Content,
	}

	return &deviceMsg.PublishMsg{
		Handle:     msg.Handle,
		Type:       msg.Type,
		Payload:    resp.Bytes(),
		ProductID:  msg.ProductID,
		DeviceName: msg.DeviceName,
	}, nil
}
