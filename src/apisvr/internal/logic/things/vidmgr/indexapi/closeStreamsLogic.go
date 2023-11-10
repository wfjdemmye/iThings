package indexapi

import (
	"context"
	"encoding/json"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CloseStreamsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCloseStreamsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CloseStreamsLogic {
	return &CloseStreamsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CloseStreamsLogic) CloseStreams(req *types.IndexApiReq) (resp *types.IndexApiCloseStreamsResp, err error) {
	// todo: add your logic here and delete this line
	data, err := proxyMediaServer(l.ctx, l.svcCtx, CLOSESTREAMS, req.VidmgrID)
	dataRecv := new(types.IndexApiCloseStreamsResp)
	json.Unmarshal(data, dataRecv)
	return dataRecv, err
}
