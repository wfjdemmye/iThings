package relationDB

import (
	"github.com/i-Things/things/shared/stores"
)

func Migrate() error {
	db := stores.GetCommonConn(nil)
	return db.AutoMigrate(
		//&DmProductInfo{},
		//&DmDeviceInfo{},
		//&DmProductCustom{},
		//&DmProductSchema{},
		//&DmGroupInfo{},
		//&DmGroupDevice{},
		//&DmGatewayDevice{},
		//&DmProductRemoteConfig{},
		&VidmgrInfo{},
		&VidmgrConfig{},
	)
}
