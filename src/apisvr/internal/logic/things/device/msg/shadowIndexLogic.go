package msg

import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShadowIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShadowIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShadowIndexLogic {
	return &ShadowIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShadowIndexLogic) ShadowIndex(req *types.DeviceMsgPropertyLatestIndexReq) (resp *types.DeviceMsgShadowIndexResp, err error) {
	dmResp, err := l.svcCtx.DeviceMsg.ShadowIndex(l.ctx, &dm.PropertyLatestIndexReq{
		DeviceName: req.DeviceName,
		ProductID:  req.ProductID,
		DataIDs:    req.DataIDs,
	})
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.ShadowIndex req=%v err=%+v", utils.FuncName(), req, er)
		return nil, er
	}
	info := make([]*types.DeviceMsgShadowIndex, 0, len(dmResp.List))
	for _, v := range dmResp.List {
		info = append(info, &types.DeviceMsgShadowIndex{
			UpdatedDeviceTime: v.UpdatedDeviceTime,
			DataID:            v.DataID,
			Value:             v.Value,
		})
	}
	return &types.DeviceMsgShadowIndexResp{
		List: info,
	}, nil
}
