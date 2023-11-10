package deviceinteractlogic

import (
	"context"
	"github.com/i-Things/things/shared/devices"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/disvr/internal/domain/deviceMsg"
	"strings"
	"time"

	"github.com/i-Things/things/src/disvr/internal/svc"
	"github.com/i-Things/things/src/disvr/pb/di"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送消息给设备
func (l *SendMsgLogic) SendMsg(in *di.SendMsgReq) (*di.SendMsgResp, error) {
	l.Infof("%s topic:%v payload:%v", utils.FuncName(), in.GetTopic(), string(in.GetPayload()))
	topicInfo, err := devices.GetTopicInfo(in.Topic)
	if err != nil {
		return nil, errors.Parameter.AddMsg("topic 不正确").AddDetail(err)
	}
	if topicInfo.Direction == devices.Up {
		return nil, errors.Parameter.AddMsg("只能发给设备")
	}
	if err = CheckIsOnline(l.ctx, l.svcCtx, devices.Core{
		ProductID:  topicInfo.ProductID,
		DeviceName: topicInfo.DeviceName,
	}); err != nil {
		return nil, err
	}

	er := l.svcCtx.PubDev.PublishToDev(l.ctx, &deviceMsg.PublishMsg{
		Timestamp:  time.Now().UnixMilli(),
		Payload:    in.Payload,
		Handle:     strings.TrimPrefix(topicInfo.TopicHead, "$"),
		Type:       topicInfo.Types[0],
		ProductID:  topicInfo.ProductID,
		DeviceName: topicInfo.DeviceName,
	})
	if er != nil {
		l.Errorf("%s.PublishToDev failure err:%v", utils.FuncName(), er)
		return nil, er
	}
	return &di.SendMsgResp{}, nil
}
