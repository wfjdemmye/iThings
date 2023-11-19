package devicemsglogic

import (
	"context"
	"github.com/i-Things/things/src/dmsvr/internal/domain/deviceMsg/msgSdkLog"
	"github.com/i-Things/things/src/dmsvr/pb/dm"

	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/errors"

	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SdkLogIndexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSdkLogIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SdkLogIndexLogic {
	return &SdkLogIndexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取设备sdk调试日志
func (l *SdkLogIndexLogic) SdkLogIndex(in *dm.SdkLogIndexReq) (*dm.SdkLogIndexResp, error) {
	filter := msgSdkLog.SdkLogFilter{
		ProductID:  in.ProductID,
		DeviceName: in.DeviceName,
		LogLevel:   int(in.LogLevel),
	}
	logs, err := l.svcCtx.SDKLogRepo.GetDeviceSDKLog(l.ctx, filter, def.PageInfo2{
		TimeStart: in.TimeStart,
		TimeEnd:   in.TimeEnd,
		Page:      in.Page.GetPage(),
		Size:      in.Page.GetSize(),
	})
	if err != nil {
		return nil, errors.Database.AddDetail(err)
	}
	var data []*dm.SdkLogIndex
	for _, v := range logs {
		data = append(data, ToDataSdkLogIndex(v))
	}
	total, err := l.svcCtx.SDKLogRepo.GetCountLog(l.ctx, filter, def.PageInfo2{
		TimeStart: in.TimeStart,
		TimeEnd:   in.TimeEnd,
		Page:      in.Page.GetPage(),
		Size:      in.Page.GetSize(),
	})
	if err != nil {
		return nil, errors.Database.AddDetail(err)
	}
	//todo 总数未统计
	return &dm.SdkLogIndexResp{List: data, Total: total}, nil
}
