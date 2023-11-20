package schema

import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/apisvr/internal/logic"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index(req *types.ProductSchemaIndexReq) (resp *types.ProductSchemaIndexResp, err error) {
	dmReq := &dm.ProductSchemaIndexReq{
		Page:        logic.ToDmPageRpc(req.Page),
		ProductID:   req.ProductID,
		Type:        req.Type,
		Tag:         req.Tag,
		Identifiers: req.Identifiers,
	}
	dmResp, err := l.svcCtx.ProductM.ProductSchemaIndex(l.ctx, dmReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.ProductSchemaIndex req=%v err=%+v", utils.FuncName(), req, er)
		return nil, er
	}
	pis := make([]*types.ProductSchemaInfo, 0, len(dmResp.List))
	for _, v := range dmResp.List {
		pi := ToSchemaInfoTypes(v)
		pis = append(pis, pi)
	}
	return &types.ProductSchemaIndexResp{
		Total: dmResp.Total,
		List:  pis,
	}, nil
}
