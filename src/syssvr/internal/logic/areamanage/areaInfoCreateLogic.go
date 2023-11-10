package areamanagelogic

import (
	"context"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/stores"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/syssvr/internal/logic"
	"github.com/i-Things/things/src/syssvr/internal/repo/relationDB"
	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"
	"github.com/zeromicro/go-zero/core/logx"
)

type AreaInfoCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	AiDB *relationDB.AreaInfoRepo
}

func NewAreaInfoCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AreaInfoCreateLogic {
	return &AreaInfoCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		AiDB:   relationDB.NewAreaInfoRepo(ctx),
	}
}

// 新增区域
func (l *AreaInfoCreateLogic) AreaInfoCreate(in *sys.AreaInfo) (*sys.Response, error) {
	if in.ProjectID == 0 || in.AreaName == "" || in.ParentAreaID == 0 || ////root节点不为0
		in.ParentAreaID == def.NotClassified { //未分类不能有下属的区域
		return nil, errors.Parameter
	}

	projPo, err := checkProject(l.ctx, in.ProjectID)
	if err != nil {
		return nil, errors.Database.AddDetail(err).WithMsg("检查项目出错")
	} else if projPo == nil {
		return nil, errors.Parameter.AddDetail(in.ProjectID).WithMsg("检查项目不存在")
	}

	if in.ParentAreaID != def.RootNode { //有选了父级项目区域
		if _, err := checkParentArea(l.ctx, in.ParentAreaID, true); err != nil {
			return nil, err
		}
	}

	areaPo := &relationDB.SysAreaInfo{
		AreaID:       stores.AreaID(l.svcCtx.AreaID.GetSnowflakeId()),
		ParentAreaID: in.ParentAreaID,                //创建时必填
		ProjectID:    stores.ProjectID(in.ProjectID), //创建时必填
		AreaName:     in.AreaName,
		Position:     logic.ToStorePoint(in.Position),
		Desc:         utils.ToEmptyString(in.Desc),
	}

	err = l.AiDB.Insert(l.ctx, areaPo)
	if err != nil {
		l.Errorf("%s.Insert err=%+v", utils.FuncName(), err)
		return nil, errors.System.AddDetail(err)
	}

	return &sys.Response{}, nil
}
