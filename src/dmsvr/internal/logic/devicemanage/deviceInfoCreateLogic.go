package devicemanagelogic

import (
	"context"
	"github.com/i-Things/things/shared/ctxs"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/stores"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/dmsvr/internal/logic"
	"github.com/i-Things/things/src/dmsvr/internal/repo/relationDB"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceInfoCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PiDB *relationDB.ProductInfoRepo
	DiDB *relationDB.DeviceInfoRepo
}

func NewDeviceInfoCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceInfoCreateLogic {
	return &DeviceInfoCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		PiDB:   relationDB.NewProductInfoRepo(ctx),
		DiDB:   relationDB.NewDeviceInfoRepo(ctx),
	}
}

/*
发现返回true 没有返回false
*/
func (l *DeviceInfoCreateLogic) CheckDevice(in *dm.DeviceInfo) (bool, error) {
	_, err := l.DiDB.FindOneByFilter(l.ctx, relationDB.DeviceFilter{ProductID: in.ProductID, DeviceNames: []string{in.DeviceName}})
	if err == nil {
		return true, nil
	}
	if errors.Cmp(err, errors.NotFind) {
		return false, nil
	}
	return false, err
}

/*
发现返回true 没有返回false
*/
func (l *DeviceInfoCreateLogic) CheckProduct(in *dm.DeviceInfo) (bool, error) {
	_, err := l.PiDB.FindOneByFilter(l.ctx, relationDB.ProductFilter{ProductIDs: []string{in.ProductID}})
	if err == nil {
		return true, nil
	}
	if errors.Cmp(err, errors.NotFind) {
		return false, nil
	}
	return false, err
}

// 新增设备
func (l *DeviceInfoCreateLogic) DeviceInfoCreate(in *dm.DeviceInfo) (resp *dm.Response, err error) {

	if in.ProductID == "" && in.ProductName != "" { //通过唯一的产品名 查找唯一的产品ID
		if pid, err := l.PiDB.FindOneByFilter(l.ctx, relationDB.ProductFilter{ProductNames: []string{in.ProductName}}); err != nil {
			return nil, err
		} else {
			in.ProductID = pid.ProductID
		}
	}

	find, err := l.CheckDevice(in)
	if err != nil {
		l.Errorf("%s.CheckDevice in=%v\n", utils.FuncName(), in)
		return nil, errors.Database.AddDetail(err)
	} else if find == true {
		return nil, errors.Duplicate.WithMsgf("设备名称重复:%s", in.DeviceName).AddDetail("DeviceName:" + in.DeviceName)
	}

	find, err = l.CheckProduct(in)
	if err != nil {
		l.Errorf("%s.CheckProduct in=%v", utils.FuncName(), in)
		return nil, errors.Database.AddDetail(err)
	} else if find == false {
		return nil, errors.Parameter.AddDetail("not find product id:" + cast.ToString(in.ProductID))
	}

	err = l.InitDevice(in)
	if err != nil {
		return nil, err
	}

	projectID := stores.ProjectID(ctxs.GetMetaProjectID(l.ctx))
	if projectID == 0 { //如果没有传项目,则分配到未分类项目中
		projectID = def.NotClassified
	}
	di := relationDB.DmDeviceInfo{
		ProjectID:  projectID,
		ProductID:  in.ProductID,  // 产品id
		DeviceName: in.DeviceName, // 设备名称
		Position:   logic.ToStorePoint(in.Position),
		AreaID:     def.NotClassified, //设备默认都是未分类
	}

	if in.Secret != "" {
		di.Secret = in.Secret
	} else {
		di.Secret = utils.GetRandomBase64(20)
	}
	di.Tags = in.Tags
	if di.Tags == nil {
		di.Tags = map[string]string{}
	}

	if in.LogLevel != def.Unknown {
		di.LogLevel = def.LogClose
	}
	if in.Address != nil {
		di.Address = in.Address.Value
	}

	if in.DeviceAlias != nil {
		di.DeviceAlias = in.DeviceAlias.Value
	}
	if in.MobileOperator != 0 {
		di.MobileOperator = in.MobileOperator
	}

	if in.Iccid != nil {
		di.Iccid = utils.AnyToNullString(in.Iccid)
	}
	if in.Phone != nil {
		di.Phone = utils.AnyToNullString(in.Phone)
	}

	err = l.DiDB.Insert(l.ctx, &di)
	if err != nil {
		l.Errorf("AddDevice.DeviceInfo.Insert err=%+v", err)
		return nil, errors.Database.AddDetail(err)
	}

	return &dm.Response{}, nil
}

func (l *DeviceInfoCreateLogic) InitDevice(in *dm.DeviceInfo) error {
	pt, err := l.svcCtx.SchemaRepo.GetSchemaModel(l.ctx, in.ProductID)
	if err != nil {
		return errors.System.AddDetail(err)
	}
	err = l.svcCtx.SchemaManaRepo.InitDevice(l.ctx, pt, in.ProductID, in.DeviceName)
	if err != nil {
		return errors.Database.AddDetail(err)
	}
	err = l.svcCtx.SDKLogRepo.InitDevice(l.ctx, in.ProductID, in.DeviceName)
	if err != nil {
		return errors.Database.AddDetail(err)
	}
	return nil
}
