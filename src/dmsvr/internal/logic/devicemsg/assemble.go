package devicemsglogic

import (
	"github.com/i-Things/things/src/dmsvr/internal/domain/deviceMsg/msgHubLog"
	"github.com/i-Things/things/src/dmsvr/internal/domain/deviceMsg/msgSdkLog"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
)

func ToDataHubLogIndex(log *msgHubLog.HubLog) *dm.HubLogIndex {
	return &dm.HubLogIndex{
		Timestamp:  log.Timestamp.UnixMilli(),
		Action:     log.Action,
		RequestID:  log.RequestID,
		TranceID:   log.TranceID,
		Topic:      log.Topic,
		Content:    log.Content,
		ResultType: log.ResultType,
	}
}

// SDK调试日志
func ToDataSdkLogIndex(log *msgSdkLog.SDKLog) *dm.SdkLogIndex {
	return &dm.SdkLogIndex{
		Timestamp: log.Timestamp.UnixMilli(),
		Loglevel:  log.LogLevel,
		Content:   log.Content,
	}
}
