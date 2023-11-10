package relationDB

import (
	"database/sql"
	"github.com/i-Things/things/shared/stores"
	"gorm.io/gorm/clause"
)

func Migrate() error {
	db := stores.GetCommonConn(nil)
	var needInitColumn bool
	if !db.Migrator().HasTable(&SysUserInfo{}) {
		//需要初始化表
		needInitColumn = true
	}
	err := db.AutoMigrate(
		&SysUserInfo{},
		&SysRoleInfo{},
		&SysRoleMenu{},
		&SysMenuInfo{},
		&SysLoginLog{},
		&SysOperLog{},
		&SysApiInfo{},
		&SysApiAuth{},
	)
	if err != nil {
		return err
	}
	if needInitColumn {
		return migrateTableColumn()
	}
	return err
}
func migrateTableColumn() error {
	db := stores.GetCommonConn(nil).Clauses(clause.OnConflict{DoNothing: true})
	if err := db.CreateInBatches(&MigrateUserInfo, 100).Error; err != nil {
		return err
	}
	if err := db.CreateInBatches(&MigrateRoleInfo, 100).Error; err != nil {
		return err
	}
	if err := db.CreateInBatches(&MigrateMenuInfo, 100).Error; err != nil {
		return err
	}
	if err := db.CreateInBatches(&MigrateRoleMenu, 100).Error; err != nil {
		return err
	}
	if err := db.CreateInBatches(&MigrateApiInfo, 100).Error; err != nil {
		return err
	}
	if err := db.CreateInBatches(&MigrateApiAuth, 100).Error; err != nil {
		return err
	}
	return nil
}

func init() {
	for i := int64(1); i <= 100; i++ {
		MigrateRoleMenu = append(MigrateRoleMenu, SysRoleMenu{
			RoleID: 1,
			MenuID: i,
		})
	}
}

var (
	MigrateUserInfo = []SysUserInfo{
		{UserID: 1740358057038188544, UserName: sql.NullString{String: "administrator", Valid: true}, Password: "4f0fded4a38abe7a3ea32f898bb82298", Role: 1, NickName: "iThings管理员"},
	}
	MigrateRoleInfo = []SysRoleInfo{{ID: 1, Name: "admin"}}
	MigrateRoleMenu []SysRoleMenu
	MigrateMenuInfo = []SysMenuInfo{
		{ID: 2, ParentID: 1, Type: 0, Order: 2, Name: "设备管理", Path: "/deviceMangers", Component: "./deviceMangers/index.tsx", Icon: "icon_data_01", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 3, ParentID: 1, Type: 0, Order: 9, Name: "系统管理", Path: "/systemManagers", Component: "./systemManagers/index.tsx", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 4, ParentID: 1, Type: 0, Order: 4, Name: "运维监控", Path: "/operationsMonitorings", Component: "./operationsMonitorings/index.tsx", Icon: "icon_hvac", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 6, ParentID: 2, Type: 2, Order: 1, Name: "产品", Path: "/deviceMangers/product/index", Component: "./deviceMangers/product/index", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 7, ParentID: 2, Type: 0, Order: 1, Name: "产品详情", Path: "/deviceMangers/product/detail/:id", Component: "./deviceMangers/product/detail/index", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 1},
		{ID: 8, ParentID: 2, Type: 0, Order: 2, Name: "设备", Path: "/deviceMangers/device/index", Component: "./deviceMangers/device/index", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 9, ParentID: 2, Type: 0, Order: 2, Name: "设备详情", Path: "/deviceMangers/device/detail/:id/:name", Component: "./deviceMangers/device/detail/index", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 1},
		{ID: 10, ParentID: 3, Type: 0, Order: 1, Name: "用户管理", Path: "/systemMangers/user/index", Component: "./systemMangers/user/index", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 11, ParentID: 3, Type: 2, Order: 2, Name: "角色管理", Path: "/systemMangers/role/index", Component: "./systemMangers/role/index", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 12, ParentID: 3, Type: 2, Order: 3, Name: "菜单列表", Path: "/systemMangers/menu/index", Component: "./systemMangers/menu/index", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 13, ParentID: 4, Type: 0, Order: 1, Name: "固件升级", Path: "/operationsMonitorings/firmwareUpgrades/index", Component: "./operationsMonitorings/firmwareUpgrades/index.tsx", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 15, ParentID: 4, Type: 0, Order: 3, Name: "资源管理", Path: "/operationsMonitorings/resourceManagements/index", Component: "./operationsMonitorings/resourceManagements/index.tsx", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 16, ParentID: 4, Type: 0, Order: 4, Name: "远程配置", Path: "/operationsMonitorings/remoteConfiguration/index", Component: "./operationsMonitorings/remoteConfiguration/index.tsx", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 18, ParentID: 4, Type: 2, Order: 6, Name: "在线调试", Path: "/operationsMonitorings/onlineDebugs/index", Component: "./operationsMonitorings/onlineDebugs/index.tsx", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 23, ParentID: 2, Type: 0, Order: 3, Name: "分组", Path: "/deviceMangers/group/index", Component: "./deviceMangers/group/index.tsx", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 24, ParentID: 2, Type: 0, Order: 3, Name: "分组详情", Path: "/deviceMangers/group/detail/:id", Component: "./deviceMangers/group/detail/index.tsx", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 1},
		{ID: 25, ParentID: 4, Type: 0, Order: 7, Name: "日志服务", Path: "/operationsMonitorings/logService/index", Component: "./operationsMonitorings/logService/index.tsx", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 35, ParentID: 1, Type: 1, Order: 1, Name: "首页", Path: "/home", Component: "./home/index.tsx", Icon: "icon_dosing", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 38, ParentID: 3, Type: 1, Order: 5, Name: "日志管理", Path: "/systemManagers/log/index", Component: "./systemManagers/log/index.tsx", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 39, ParentID: 38, Type: 1, Order: 1, Name: "操作日志", Path: "/systemMangers/log/operationLog/index", Component: "./systemMangers/log/operationLog/index.tsx", Icon: "icon_dosing", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 41, ParentID: 38, Type: 1, Order: 2, Name: "登录日志", Path: "/systemMangers/log/loginLog/index", Component: "./systemMangers/log/loginLog/index", Icon: "icon_heat", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 42, ParentID: 3, Type: 1, Order: 4, Name: "接口管理", Path: "/systemMangers/api/index", Component: "./systemMangers/api/index", Icon: "icon_system", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 43, ParentID: 1, Type: 1, Order: 5, Name: "告警管理", Path: "/alarmMangers", Component: "./alarmMangers/index", Icon: "icon_ap", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 44, ParentID: 43, Type: 1, Order: 1, Name: "告警配置", Path: "/alarmMangers/alarmConfiguration/index", Component: "./alarmMangers/alarmConfiguration/index", Icon: "icon_ap", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 53, ParentID: 43, Type: 1, Order: 5, Name: "新增告警配置", Path: "/alarmMangers/alarmConfiguration/save", Component: "./alarmMangers/alarmConfiguration/addAlarmConfig/index", Icon: "icon_ap", Redirect: "", BackgroundUrl: "", HideInMenu: 1},
		{ID: 54, ParentID: 43, Type: 1, Order: 5, Name: "告警日志", Path: "/alarmMangers/alarmConfiguration/log/detail/:id/:level", Component: "./alarmMangers/alarmLog/index", Icon: "icon_ap", Redirect: "", BackgroundUrl: "", HideInMenu: 1},
		{ID: 62, ParentID: 43, Type: 1, Order: 5, Name: "告警记录", Path: "/alarmMangers/alarmConfiguration/log", Component: "./alarmMangers/alarmRecord/index", Icon: "icon_ap", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 50, ParentID: 1, Type: 1, Order: 5, Name: "规则引擎", Path: "/ruleEngine", Component: "./ruleEngine/index.tsx", Icon: "icon_dosing", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 51, ParentID: 50, Type: 1, Order: 1, Name: "场景联动", Path: "/ruleEngine/scene/index", Component: "./ruleEngine/scene/index.tsx", Icon: "icon_device", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		//视频服务菜单项
		{ID: 63, ParentID: 1, Type: 1, Order: 2, Name: "视频服务", Path: "/videoMangers", Component: "./videoMangers", Icon: "icon_heat", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 64, ParentID: 63, Type: 1, Order: 1, Name: "服务管理", Path: "/videoMangers/vidsrvmgr/index", Component: "./videoMangers/vidsrvmgr/index.tsx", Icon: "icon_heat", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 65, ParentID: 63, Type: 1, Order: 3, Name: "视频广场", Path: "/videoMangers/plaza/index", Component: "./videoMangers/plaza/index.tsx", Icon: "icon_heat", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 66, ParentID: 63, Type: 1, Order: 2, Name: "接入管理", Path: "/videoMangers/vdevice/index", Component: "./videoMangers/vdevice/index.tsx", Icon: "icon_heat", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 67, ParentID: 63, Type: 1, Order: 4, Name: "视频回放", Path: "/videoMangers/playback/index", Component: "./videoMangers/playback/index.tsx", Icon: "icon_heat", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 68, ParentID: 63, Type: 1, Order: 2, Name: "录像计划", Path: "/videoMangers/recordplan/index", Component: "./videoMangers/recordplan/index.tsx", Icon: "icon_heat", Redirect: "", BackgroundUrl: "", HideInMenu: 2},
		{ID: 70, ParentID: 63, Type: 1, Order: 1, Name: "服务详细", Path: "/videoMangers/vidsrvmgr/detail/:id", Component: "./videoMangers/vidsrvmgr/detail/index", Icon: "icon_heat", Redirect: "", BackgroundUrl: "", HideInMenu: 1},
	}
	MigrateApiInfo = []SysApiInfo{
		{Route: "/api/v1/things/product/info/update", Method: 2, Name: "更新产品", BusinessType: 2, Desc: "", Group: "产品管理"},
		{Route: "/api/v1/things/product/info/create", Method: 2, Name: "新增产品", BusinessType: 1, Desc: "", Group: "产品管理"},
		{Route: "/api/v1/things/product/info/read", Method: 2, Name: "获取产品详情", BusinessType: 4, Desc: "", Group: "产品管理"},
		{Route: "/api/v1/things/product/info/delete", Method: 2, Name: "删除产品", BusinessType: 3, Desc: "", Group: "产品管理"},
		{Route: "/api/v1/things/product/info/index", Method: 2, Name: "获取产品列表", BusinessType: 4, Desc: "", Group: "产品管理"},
		{Route: "/api/v1/things/product/custom/read", Method: 2, Name: "获取产品自定义信息", BusinessType: 4, Desc: "", Group: "产品自定义信息"},
		{Route: "/api/v1/things/product/custom/update", Method: 2, Name: "更新产品自定义信息", BusinessType: 2, Desc: "", Group: "产品自定义信息"},
		{Route: "/api/v1/things/product/schema/index", Method: 2, Name: "获取产品物模型列表", BusinessType: 4, Desc: "", Group: "物模型"},
		{Route: "/api/v1/things/product/schema/tsl-import", Method: 2, Name: "导入物模型tsl", BusinessType: 1, Desc: "", Group: "物模型"},
		{Route: "/api/v1/things/product/schema/tsl-read", Method: 2, Name: "获取产品物模型tsl", BusinessType: 4, Desc: "", Group: "物模型"},
		{Route: "/api/v1/things/product/schema/create", Method: 2, Name: "新增物模型功能", BusinessType: 1, Desc: "", Group: "物模型"},
		{Route: "/api/v1/things/product/schema/update", Method: 2, Name: "更新物模型功能", BusinessType: 2, Desc: "", Group: "物模型"},
		{Route: "/api/v1/things/product/schema/delete", Method: 2, Name: "删除物模型功能", BusinessType: 3, Desc: "", Group: "物模型"},
		{Route: "/api/v1/things/product/remote-config/create", Method: 2, Name: "创建配置", BusinessType: 1, Desc: "", Group: "产品远程配置"},
		{Route: "/api/v1/things/product/remote-config/index", Method: 2, Name: "获取配置列表", BusinessType: 4, Desc: "", Group: "产品远程配置"},
		{Route: "/api/v1/things/product/remote-config/push-all", Method: 2, Name: "推送配置", BusinessType: 5, Desc: "", Group: "产品远程配置"},
		{Route: "/api/v1/things/product/remote-config/lastest-read", Method: 2, Name: "获取最新配置", BusinessType: 4, Desc: "", Group: "产品远程配置"},
		{Route: "/api/v1/things/group/info/create", Method: 2, Name: "创建分组", BusinessType: 1, Desc: "", Group: "设备分组"},
		{Route: "/api/v1/things/group/info/index", Method: 2, Name: "获取分组列表", BusinessType: 4, Desc: "", Group: "设备分组"},
		{Route: "/api/v1/things/group/info/read", Method: 2, Name: "获取分组详情信息", BusinessType: 4, Desc: "", Group: "设备分组"},
		{Route: "/api/v1/things/group/info/update", Method: 2, Name: "更新分组信息", BusinessType: 2, Desc: "", Group: "设备分组"},
		{Route: "/api/v1/things/group/info/delete", Method: 2, Name: "删除分组", BusinessType: 3, Desc: "", Group: "设备分组"},
		{Route: "/api/v1/things/group/device/index", Method: 2, Name: "获取分组设备列表", BusinessType: 4, Desc: "", Group: "设备分组"},
		{Route: "/api/v1/things/group/device/multi-create", Method: 2, Name: "添加分组设备(支持批量)", BusinessType: 1, Desc: "", Group: "设备分组"},
		{Route: "/api/v1/things/group/device/multi-delete", Method: 2, Name: "删除分组设备(支持批量)", BusinessType: 3, Desc: "", Group: "设备分组"},
		{Route: "/api/v1/things/device/info/index", Method: 2, Name: "获取设备列表", BusinessType: 4, Desc: "", Group: "设备管理"},
		{Route: "/api/v1/things/device/info/read", Method: 2, Name: "获取设备详情", BusinessType: 4, Desc: "", Group: "设备管理"},
		{Route: "/api/v1/things/device/info/create", Method: 2, Name: "新增设备", BusinessType: 1, Desc: "", Group: "设备管理"},
		{Route: "/api/v1/things/device/info/delete", Method: 2, Name: "删除设备", BusinessType: 3, Desc: "", Group: "设备管理"},
		{Route: "/api/v1/things/device/info/update", Method: 2, Name: "更新设备", BusinessType: 2, Desc: "", Group: "设备管理"},
		{Route: "/api/v1/things/device/info/count", Method: 2, Name: "设备统计详情", BusinessType: 4, Desc: "", Group: "设备管理"},
		{Route: "/api/v1/things/device/info/multi-import", Method: 2, Name: "批量导入设备", BusinessType: 1, Desc: "", Group: "设备管理"},
		{Route: "/api/v1/things/device/auth/login", Method: 2, Name: "设备登录认证", BusinessType: 5, Desc: "", Group: "设备鉴权"},
		{Route: "/api/v1/things/device/auth/root-check", Method: 2, Name: "鉴定mqtt账号root权限", BusinessType: 5, Desc: "", Group: "设备鉴权"},
		{Route: "/api/v1/things/device/auth/access", Method: 2, Name: "设备操作认证", BusinessType: 5, Desc: "", Group: "设备鉴权"},
		{Route: "/api/v1/things/device/msg/property-log/index", Method: 2, Name: "获取单个id属性历史记录", BusinessType: 4, Desc: "", Group: "设备消息"},
		{Route: "/api/v1/things/device/msg/sdk-log/index", Method: 2, Name: "获取设备本地日志", BusinessType: 4, Desc: "", Group: "设备消息"},
		{Route: "/api/v1/things/device/msg/hub-log/index", Method: 2, Name: "获取云端诊断日志", BusinessType: 4, Desc: "", Group: "设备消息"},
		{Route: "/api/v1/things/device/msg/property-latest/index", Method: 2, Name: "获取最新属性", BusinessType: 4, Desc: "", Group: "设备消息"},
		{Route: "/api/v1/things/device/msg/event-log/index", Method: 2, Name: "获取物模型事件历史记录", BusinessType: 4, Desc: "", Group: "设备消息"},
		{Route: "/api/v1/things/device/interact/send-action", Method: 2, Name: "同步调用设备行为", BusinessType: 5, Desc: "", Group: "设备交互"},
		{Route: "/api/v1/things/device/interact/send-property", Method: 2, Name: "同步调用设备属性", BusinessType: 5, Desc: "", Group: "设备交互"},
		{Route: "/api/v1/things/device/interact/multi-send-property", Method: 2, Name: "批量调用设备属性", BusinessType: 5, Desc: "", Group: "设备交互"},
		{Route: "/api/v1/things/device/interact/get-property-reply", Method: 2, Name: "请求设备获取设备最新属性", BusinessType: 4, Desc: "", Group: "设备交互"},
		{Route: "/api/v1/things/device/interact/send-msg", Method: 2, Name: "发送消息给设备", BusinessType: 5, Desc: "", Group: "设备交互"},
		{Route: "/api/v1/things/device/gateway/multi-create", Method: 2, Name: "批量添加网关子设备", BusinessType: 1, Desc: "", Group: "网关子设备管理"},
		{Route: "/api/v1/things/device/gateway/multi-delete", Method: 2, Name: "批量解绑网关子设备", BusinessType: 3, Desc: "", Group: "网关子设备管理"},
		{Route: "/api/v1/things/device/gateway/index", Method: 2, Name: "获取子设备列表", BusinessType: 4, Desc: "", Group: "网关子设备管理"},
		{Route: "/api/v1/system/log/login/index", Method: 2, Name: "获取登录日志列表", BusinessType: 4, Desc: "", Group: "日志管理"},
		{Route: "/api/v1/system/log/oper/index", Method: 2, Name: "获取操作日志列表", BusinessType: 4, Desc: "", Group: "日志管理"},
		{Route: "/api/v1/system/role/create", Method: 2, Name: "添加角色", BusinessType: 1, Desc: "", Group: "角色管理"},
		{Route: "/api/v1/system/role/index", Method: 2, Name: "获取角色列表", BusinessType: 4, Desc: "", Group: "角色管理"},
		{Route: "/api/v1/system/role/update", Method: 2, Name: "更新角色", BusinessType: 2, Desc: "", Group: "角色管理"},
		{Route: "/api/v1/system/role/delete", Method: 2, Name: "删除角色", BusinessType: 3, Desc: "", Group: "角色管理"},
		{Route: "/api/v1/system/role/role-menu/update", Method: 2, Name: "更新角色对应菜单列表", BusinessType: 2, Desc: "", Group: "角色管理"},
		{Route: "/api/v1/system/menu/create", Method: 2, Name: "添加菜单", BusinessType: 1, Desc: "", Group: "菜单管理"},
		{Route: "/api/v1/system/menu/index", Method: 2, Name: "获取菜单列表", BusinessType: 4, Desc: "", Group: "菜单管理"},
		{Route: "/api/v1/system/menu/update", Method: 2, Name: "更新菜单", BusinessType: 2, Desc: "", Group: "菜单管理"},
		{Route: "/api/v1/system/menu/delete", Method: 2, Name: "删除菜单", BusinessType: 3, Desc: "", Group: "菜单管理"},
		{Route: "/api/v1/system/user/create", Method: 2, Name: "创建用户信息", BusinessType: 1, Desc: "", Group: "用户管理"},
		{Route: "/api/v1/system/user/captcha", Method: 2, Name: "获取验证码", BusinessType: 5, Desc: "", Group: "用户管理"},
		{Route: "/api/v1/system/user/login", Method: 2, Name: "登录", BusinessType: 5, Desc: "", Group: "用户管理"},
		{Route: "/api/v1/system/user/delete", Method: 2, Name: "删除用户", BusinessType: 3, Desc: "", Group: "用户管理"},
		{Route: "/api/v1/system/user/read", Method: 2, Name: "获取用户信息", BusinessType: 4, Desc: "", Group: "用户管理"},
		{Route: "/api/v1/system/user/update", Method: 2, Name: "更新用户基本数据", BusinessType: 2, Desc: "", Group: "用户管理"},
		{Route: "/api/v1/system/user/index", Method: 2, Name: "获取用户信息列表", BusinessType: 4, Desc: "", Group: "用户管理"},
		{Route: "/api/v1/system/user/resource-read", Method: 2, Name: "获取用户资源", BusinessType: 4, Desc: "", Group: "用户管理"},
		{Route: "/api/v1/system/common/config", Method: 2, Name: "获取系统配置", BusinessType: 4, Desc: "", Group: "系统配置"},
		{Route: "/api/v1/system/api/create", Method: 2, Name: "添加接口", BusinessType: 1, Desc: "", Group: "接口管理"},
		{Route: "/api/v1/system/api/index", Method: 2, Name: "获取接口列表", BusinessType: 4, Desc: "", Group: "接口管理"},
		{Route: "/api/v1/system/api/update", Method: 2, Name: "更新接口", BusinessType: 2, Desc: "", Group: "接口管理"},
		{Route: "/api/v1/system/api/delete", Method: 2, Name: "删除接口", BusinessType: 3, Desc: "", Group: "接口管理"},
		{Route: "/api/v1/system/auth/api/index", Method: 2, Name: "获取API权限列表", BusinessType: 4, Desc: "", Group: "权限管理"},
		{Route: "/api/v1/system/auth/api/multiUpdate", Method: 2, Name: "更新API权限", BusinessType: 2, Desc: "", Group: "权限管理"},
		{Route: "/api/v1/things/rule/scene/info/read", Method: 2, Name: "获取场景信息", BusinessType: 4, Desc: "", Group: "场景联动"},
		{Route: "/api/v1/things/rule/scene/info/index", Method: 2, Name: "获取场景列表", BusinessType: 4, Desc: "", Group: "场景联动"},
		{Route: "/api/v1/things/rule/scene/info/create", Method: 2, Name: "创建场景信息", BusinessType: 1, Desc: "", Group: "场景联动"},
		{Route: "/api/v1/things/rule/scene/info/update", Method: 2, Name: "更新场景信息", BusinessType: 2, Desc: "", Group: "场景联动"},
		{Route: "/api/v1/things/rule/scene/info/delete", Method: 2, Name: "删除场景信息", BusinessType: 3, Desc: "", Group: "场景联动"},
		{Route: "/api/v1/things/rule/flow/info/index", Method: 2, Name: "获取流列表", BusinessType: 4, Desc: "", Group: "流"},
		{Route: "/api/v1/things/rule/flow/info/create", Method: 2, Name: "创建流", BusinessType: 1, Desc: "", Group: "流"},
		{Route: "/api/v1/things/rule/flow/info/update", Method: 2, Name: "修改流", BusinessType: 2, Desc: "", Group: "流"},
		{Route: "/api/v1/things/rule/flow/info/delete", Method: 2, Name: "删除流", BusinessType: 3, Desc: "", Group: "流"},
		{Route: "/api/v1/things/rule/alarm/info/create", Method: 2, Name: "新增告警", BusinessType: 1, Desc: "", Group: "告警管理"},
		{Route: "/api/v1/things/rule/alarm/info/update", Method: 2, Name: "更新告警", BusinessType: 2, Desc: "", Group: "告警管理"},
		{Route: "/api/v1/things/rule/alarm/info/delete", Method: 2, Name: "删除告警", BusinessType: 3, Desc: "", Group: "告警管理"},
		{Route: "/api/v1/things/rule/alarm/info/index", Method: 2, Name: "获取告警信息列表", BusinessType: 4, Desc: "", Group: "告警管理"},
		{Route: "/api/v1/things/rule/alarm/info/read", Method: 2, Name: "获取告警详情", BusinessType: 4, Desc: "", Group: "告警管理"},
		{Route: "/api/v1/things/rule/alarm/scene/delete", Method: 2, Name: "删除告警和场景的关联", BusinessType: 3, Desc: "", Group: "场景联动"},
		{Route: "/api/v1/things/rule/alarm/log/index", Method: 2, Name: "获取告警流水日志记录列表", BusinessType: 4, Desc: "", Group: "告警日志"},
		{Route: "/api/v1/things/rule/alarm/record/index", Method: 2, Name: "获取告警记录列表", BusinessType: 4, Desc: "", Group: "告警记录"},
		{Route: "/api/v1/things/rule/alarm/deal-record/create", Method: 2, Name: "新增告警处理记录", BusinessType: 1, Desc: "", Group: "处理记录"},
		{Route: "/api/v1/things/rule/alarm/deal-record/index", Method: 2, Name: "获取告警处理记录列表", BusinessType: 4, Desc: "", Group: "处理记录"},
		{Route: "/api/v1/things/rule/alarm/scene/multi-update", Method: 2, Name: "更新告警和场景的关联", BusinessType: 2, Desc: "", Group: "场景联动"},
		//视频服务API接口
		{Route: "/api/v1/things/vidmgr/info/count", Method: 2, Name: "流服务器统计", BusinessType: 1, Desc: "", Group: "视频服务"},
		{Route: "/api/v1/things/vidmgr/info/create", Method: 2, Name: "新增流服务器", BusinessType: 1, Desc: "", Group: "视频服务"},
		{Route: "/api/v1/things/vidmgr/info/delete", Method: 2, Name: "删除流服务器", BusinessType: 1, Desc: "", Group: "视频服务"},
		{Route: "/api/v1/things/vidmgr/info/index", Method: 2, Name: "获取流服务器列表", BusinessType: 1, Desc: "", Group: "视频服务"},
		{Route: "/api/v1/things/vidmgr/info/read", Method: 2, Name: "获取流服详细", BusinessType: 1, Desc: "", Group: "视频服务"},
		{Route: "/api/v1/things/vidmgr/info/update", Method: 2, Name: "更新流服务器", BusinessType: 1, Desc: "", Group: "视频服务"},
	}
	MigrateApiAuth = []SysApiAuth{
		{PType: "p", V0: "1", V1: "/api/v1/things/product/info/update", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/info/create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/info/read", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/info/delete", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/info/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/schema/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/schema/tsl-import", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/schema/tsl-read", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/schema/create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/schema/update", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/schema/delete", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/remote-config/create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/remote-config/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/remote-config/push-all", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/remote-config/lastest-read", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/custom/read", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/product/custom/update", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/group/info/create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/group/info/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/group/info/read", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/group/info/update", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/group/info/delete", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/group/device/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/group/device/multi-create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/group/device/multi-delete", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/info/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/info/read", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/info/create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/info/delete", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/info/update", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/info/count", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/auth/login", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/auth/root-check", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/auth/access", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/msg/property-log/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/msg/sdk-log/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/msg/hub-log/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/msg/property-latest/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/msg/event-log/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/interact/send-action", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/interact/send-property", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/interact/multi-send-property", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/interact/send-msg", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/gateway/multi-create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/gateway/multi-delete", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/device/gateway/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/log/login/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/log/oper/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/role/create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/role/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/role/update", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/role/delete", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/role/role-menu/update", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/menu/create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/menu/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/menu/update", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/menu/delete", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/user/create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/user/captcha", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/user/login", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/user/delete", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/user/read", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/user/update", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/user/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/user/resource-read", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/common/config", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/api/create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/api/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/api/update", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/api/delete", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/auth/api/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/system/auth/api/multiUpdate", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/scene/info/read", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/scene/info/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/scene/info/create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/scene/info/update", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/scene/info/delete", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/flow/info/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/flow/info/create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/flow/info/update", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/flow/info/delete", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/alarm/info/create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/alarm/info/update", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/alarm/info/delete", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/alarm/info/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/alarm/info/read", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/alarm/scene/delete", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/alarm/log/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/alarm/record/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/alarm/deal-record/create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/alarm/deal-record/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/rule/alarm/scene/multi-update", V2: "2", V3: "", V4: "", V5: ""},

		//视频服务API接口
		{PType: "p", V0: "1", V1: "/api/v1/things/vidmgr/info/count", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/vidmgr/info/create", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/vidmgr/info/delete", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/vidmgr/info/index", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/vidmgr/info/read", V2: "2", V3: "", V4: "", V5: ""},
		{PType: "p", V0: "1", V1: "/api/v1/things/vidmgr/info/update", V2: "2", V3: "", V4: "", V5: ""},
	}
)
