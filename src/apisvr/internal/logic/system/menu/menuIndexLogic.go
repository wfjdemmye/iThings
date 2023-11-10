package menu

import (
	"context"
	"github.com/i-Things/things/src/apisvr/internal/logic/system"
	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuIndexLogic {
	return &MenuIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuIndexLogic) MenuIndex(req *types.MenuIndexReq) (resp *types.MenuIndexResp, err error) {

	info, err := l.svcCtx.MenuRpc.MenuIndex(l.ctx, &sys.MenuIndexReq{
		Name: req.Name,
		Path: req.Path,
	})
	if err != nil {
		return nil, err
	}

	var menuInfo []*types.MenuData

	menuInfo = make([]*types.MenuData, 0, len(menuInfo))
	for _, i := range info.List {
		menuInfo = append(menuInfo, system.ToMenuInfoApi(i))
	}

	return &types.MenuIndexResp{menuInfo}, nil
}
