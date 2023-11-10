package ctxs

import (
	"context"
	"github.com/i-Things/things/shared/utils/cast"
)

type UserCtx struct {
	IsOpen    bool   //是否开放认证用户
	UserID    int64  //用户id（开放认证用户值为0）
	Role      int64  //用户角色（开放认证用户值为0）
	IsAllData bool   //是否所有数据权限（开放认证用户值为true）
	IP        string //用户的ip地址
	Os        string //操作系统
	InnerCtx
}

type InnerCtx struct {
	AllArea bool //内部使用,不限制区域
}

func SetUserCtx(ctx context.Context, userCtx *UserCtx) context.Context {
	return context.WithValue(ctx, UserInfoKey, userCtx)
}
func SetInnerCtx(ctx context.Context, inner InnerCtx) context.Context {
	uc := GetUserCtx(ctx)
	if uc == nil {
		return ctx
	}
	uc.InnerCtx = inner
	return SetUserCtx(ctx, uc)
}

func GetInnerCtx(ctx context.Context) InnerCtx {
	uc := GetUserCtx(ctx)
	if uc == nil {
		return InnerCtx{}
	}
	return uc.InnerCtx
}

// 使用该函数前必须传了UserCtx
func GetUserCtx(ctx context.Context) *UserCtx {
	val, ok := ctx.Value(UserInfoKey).(*UserCtx)
	if !ok { //这里线上不能获取不到
		return nil
	}
	return val
}

// 使用该函数前必须传了UserCtx
func GetUserCtxOrNil(ctx context.Context) *UserCtx {
	val, ok := ctx.Value(UserInfoKey).(*UserCtx)
	if !ok { //这里线上不能获取不到
		return nil
	}
	return val
}

type MetadataCtx = map[string][]string

func SetMetaCtx(ctx context.Context, maps MetadataCtx) context.Context {
	return context.WithValue(ctx, MetadataKey, maps)
}
func GetMetaCtx(ctx context.Context) MetadataCtx {
	val, ok := ctx.Value(MetadataKey).(MetadataCtx)
	if !ok {
		return nil
	}
	return val
}

func GetMetaVal(ctx context.Context, field string) []string {
	mdCtx := GetMetaCtx(ctx)
	if val, ok := mdCtx[field]; !ok {
		return nil
	} else {
		return val
	}
}

// 获取meta里的项目ID（企业版功能）
func GetMetaProjectID(ctx context.Context) int64 {
	items := GetMetaVal(ctx, string(MetaFieldProjectID))
	if len(items) == 0 {
		return 0
	} else {
		return cast.ToInt64(items[0])
	}
}

// 指定项目id（企业版功能）
func SetMetaProjectID(ctx context.Context, projectID int64) {
	mc := GetMetaCtx(ctx)
	projectIDStr := cast.ToString(projectID)
	mc[string(MetaFieldProjectID)] = []string{projectIDStr}
}

// 获取meta里的项目ID（企业版功能）
func ClearMetaProjectID(ctx context.Context) {
	mc := GetMetaCtx(ctx)
	delete(mc, string(MetaFieldProjectID))
}
