package role

import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDeleteLogic) RoleDelete(req *types.RoleDeleteReq) error {
	resp, err := l.svcCtx.RoleRpc.RoleDelete(l.ctx, &sys.RoleDeleteReq{
		Id: req.ID,
	})
	if err != nil {
		err := errors.Fmt(err)
		l.Errorf("%s.rpc.RoleDelete req=%v err=%+v", utils.FuncName(), req, err)
		return err
	}
	if resp == nil {
		l.Errorf("%s.rpc.RoleDelete return nil req=%+v", utils.FuncName(), req)
		return errors.System.AddDetail("RoleDelete rpc return nil")
	}
	return nil
}
