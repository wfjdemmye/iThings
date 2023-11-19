package devicemanagelogic

import (
	"context"
	"database/sql"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/devices"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/stores"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/internal/domain/deviceMsg"
	"github.com/i-Things/things/src/dmsvr/internal/domain/deviceMsg/msgSdkLog"
	"github.com/i-Things/things/src/dmsvr/internal/logic"
	"github.com/i-Things/things/src/dmsvr/internal/repo/relationDB"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceInfoUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PiDB *relationDB.ProductInfoRepo
	DiDB *relationDB.DeviceInfoRepo
}

func NewDeviceInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceInfoUpdateLogic {
	return &DeviceInfoUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		PiDB:   relationDB.NewProductInfoRepo(ctx),
		DiDB:   relationDB.NewDeviceInfoRepo(ctx),
	}
}

func (l *DeviceInfoUpdateLogic) SetDevicePoByDto(old *relationDB.DmDeviceInfo, data *dm.DeviceInfo) {
	if data.AreaID != 0 {
		old.AreaID = stores.AreaID(data.AreaID)
	}

	if data.Tags != nil {
		old.Tags = data.Tags
	}
	if data.LogLevel != def.Unknown {
		old.LogLevel = data.LogLevel
	}

	if data.Imei != "" {
		old.Imei = data.Imei
	}
	if data.Mac != "" {
		old.Mac = data.Mac
	}
	if data.Version != nil {
		old.Version = data.Version.GetValue()
	}
	if data.HardInfo != "" {
		old.HardInfo = data.HardInfo
	}
	if data.SoftInfo != "" {
		old.SoftInfo = data.SoftInfo
	}

	if data.IsOnline != def.Unknown {
		old.IsOnline = data.IsOnline
		if data.IsOnline == def.True { //需要处理第一次上线的情况,一般在网关代理登录时需要处理
			now := sql.NullTime{
				Valid: true,
				Time:  time.Now(),
			}
			if old.FirstLogin.Valid == false {
				old.FirstLogin = now
			}
			old.LastLogin = now
		}
	}

	if data.Address != nil {
		old.Address = data.Address.Value
	}
	if data.Position != nil {
		old.Position = logic.ToStorePoint(data.Position)
	}

	if data.DeviceAlias != nil {
		old.DeviceAlias = data.DeviceAlias.Value
	}
	if data.MobileOperator != 0 {
		old.MobileOperator = data.MobileOperator
	}
	if data.Iccid != nil {
		old.Iccid = utils.AnyToNullString(data.Iccid)
	}
	if data.Phone != nil {
		old.Phone = utils.AnyToNullString(data.Phone)
	}
}

// 更新设备
func (l *DeviceInfoUpdateLogic) DeviceInfoUpdate(in *dm.DeviceInfo) (*dm.Response, error) {
	if in.ProductID == "" && in.ProductName != "" { //通过唯一的产品名 查找唯一的产品ID
		if pid, err := l.PiDB.FindOneByFilter(l.ctx, relationDB.ProductFilter{ProductNames: []string{in.ProductName}}); err != nil {
			return nil, err
		} else {
			in.ProductID = pid.ProductID
		}
	}
	dmDiPo, err := l.DiDB.FindOneByFilter(l.ctx, relationDB.DeviceFilter{ProductID: in.ProductID, DeviceNames: []string{in.DeviceName}})
	if err != nil {
		if errors.Cmp(err, errors.NotFind) {
			return nil, errors.NotFind.AddDetailf("not find device productID=%s deviceName=%s",
				in.ProductID, in.DeviceName)
		}
		return nil, errors.Database.AddDetail(err)
	}

	l.SetDevicePoByDto(dmDiPo, in)

	err = l.DiDB.Update(l.ctx, dmDiPo)
	if err != nil {
		l.Errorf("DeviceInfoUpdate.DeviceInfo.Update err=%+v", err)
		return nil, err
	}

	if in.LogLevel != def.Unknown {
		di, err := l.DiDB.FindOneByFilter(l.ctx, relationDB.DeviceFilter{ProductID: in.ProductID, DeviceNames: []string{in.DeviceName}, WithProduct: true})
		if err != nil {
			return nil, err
		}
		resp := deviceMsg.NewRespCommonMsg(l.ctx, deviceMsg.GetStatus, "")
		resp.Data = map[string]any{"logLevel": di.LogLevel}

		msg := deviceMsg.PublishMsg{
			Handle:     devices.Log,
			Type:       msgSdkLog.TypeUpdate,
			Payload:    resp.AddStatus(errors.OK).Bytes(),
			Timestamp:  time.Now().UnixMilli(),
			ProductID:  di.ProductID,
			DeviceName: di.DeviceName,
		}
		er := l.svcCtx.PubDev.PublishToDev(l.ctx, &msg)
		if er != nil {
			l.Errorf("%s.PublishToDev failure err:%v", utils.FuncName(), er)
			return nil, err
		}
	}
	return &dm.Response{}, nil
}
