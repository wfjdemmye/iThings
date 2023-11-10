package info

import (
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/apisvr/internal/logic"
	"github.com/i-Things/things/src/apisvr/internal/types"
	"github.com/i-Things/things/src/vidsvr/pb/vid"
)

func vidmgrInfoToApi(v *vid.VidmgrInfo) *types.VidmgrInfo {
	return &types.VidmgrInfo{
		CreatedTime:  v.CreatedTime,              //创建时间 只读
		VidmgrID:     v.VidmgrID,                 //服务id 只读
		VidmgrName:   v.VidmgrName,               //服务名称
		VidmgrIpV4:   v.VidmgrIpV4,               //服务IP
		VidmgrPort:   v.VidmgrPort,               //服务端口
		VidmgrType:   v.VidmgrType,               //服务类型:1:zlmediakit,2:srs,3:monibuca
		VidmgrSecret: v.VidmgrSecret,             //服务秘钥
		VidmgrStatus: v.VidmgrStatus,             //服务状态: 1：离线 2：在线  3：未激活
		Desc:         utils.ToNullString(v.Desc), //描述
		Tags:         logic.ToTagsType(v.Tags),
	}
}
