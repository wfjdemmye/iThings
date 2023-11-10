package user

import (
	"context"
	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"
	"github.com/i-Things/things/src/syssvr/pb/sys"
	"github.com/zeromicro/go-zero/core/logx"
)

type ReadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadLogic {
	return &ReadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReadLogic) Read(req *types.UserReadReq) (resp *types.UserInfo, err error) {
	info, err := l.svcCtx.UserRpc.UserRead(l.ctx, &sys.UserReadReq{UserID: req.UserID})
	if err != nil {
		return nil, err
	}

	return UserInfoToApi(info), nil
}
