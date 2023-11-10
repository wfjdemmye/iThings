package firmwaremanagelogic

import (
	"context"

	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/src/dmsvr/internal/repo/relationDB"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type FirmwareInfoReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFirmwareInfoReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FirmwareInfoReadLogic {
	return &FirmwareInfoReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FirmwareInfoReadLogic) FirmwareInfoRead(in *dm.FirmwareInfoReadReq) (*dm.FirmwareInfoReadResp, error) {
	di, err := relationDB.NewOtaFirmwareRepo(l.ctx).FindOneByFilter(l.ctx, relationDB.OtaFirmwareFilter{
		FirmwareID: in.FirmwareID,
	})
	if err != nil {
		return nil, err
	}

	df, err := relationDB.NewOtaFirmwareFileRepo(l.ctx).FindByFilter(l.ctx, relationDB.OtaFirmwareFileFilter{
		FirmwareID: in.FirmwareID,
	}, &def.PageInfo{Size: 20, Page: 0})
	if err != nil {
		return nil, err
	}
	return ToFirmwareRespInfo(di, df...), nil
}
