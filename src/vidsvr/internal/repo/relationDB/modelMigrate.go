package relationDB

import (
	"github.com/i-Things/things/shared/stores"
)

func Migrate() error {
	db := stores.GetCommonConn(nil)
	return db.AutoMigrate(
		&VidmgrInfo{},
		&VidmgrConfig{},
		&VidmgrStream{},
	)
}
