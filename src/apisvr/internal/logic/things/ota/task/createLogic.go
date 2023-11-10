package task

import (
	"context"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.OtaTaskCreateReq) error {
	otaTaskCreateReq := dm.OtaTaskCreateReq{
		FirmwareID:  req.FirmwareID,
		Type:        req.Type,
		UpgradeType: req.UpgradeType,
	}
	if req.DeviceList != "" {
		otaTaskCreateReq.DeviceList = &wrappers.StringValue{
			Value: req.DeviceList,
		}
	}
	if req.VersionList != "" {
		otaTaskCreateReq.VersionList = &wrappers.StringValue{
			Value: req.VersionList,
		}
	}
	_, err := l.svcCtx.OtaTaskM.OtaTaskCreate(l.ctx, &otaTaskCreateReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s|rpc.OtaTaskCreate|req=%v|err=%+v", utils.FuncName(), req, er)
		return er
	}
	return nil
}
