package msg

import (
	"context"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/apisvr/internal/logic"
	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SdkLogIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSdkLogIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SdkLogIndexLogic {
	return &SdkLogIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SdkLogIndexLogic) SdkLogIndex(req *types.DeviceMsgSdkLogIndexReq) (resp *types.DeviceMsgSdkIndexResp, err error) {
	dmReq := &dm.SdkLogIndexReq{
		DeviceName: req.DeviceName,
		ProductID:  req.ProductID, //产品id
		LogLevel:   int64(req.LogLevel),
		TimeStart:  req.TimeStart,
		TimeEnd:    req.TimeEnd,
		Page:       logic.ToDiPageRpc(req.Page),
	}

	dmResp, err := l.svcCtx.DeviceMsg.SdkLogIndex(l.ctx, dmReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.GetDeviceSDKLog req=%v err=%+v", utils.FuncName(), req, er)
		return nil, er
	}
	info := make([]*types.DeviceMsgSdkIndex, 0, len(dmResp.List))
	for _, v := range dmResp.List {
		info = append(info, &types.DeviceMsgSdkIndex{
			Timestamp: v.Timestamp,
			Loglevel:  v.Loglevel,
			Content:   v.Content,
		})
	}
	return &types.DeviceMsgSdkIndexResp{List: info, Total: dmResp.Total}, err

}
