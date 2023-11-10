package dataUpdate

import (
	"context"
	"github.com/i-Things/things/shared/conf"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/events"
)

type (
	DataUpdate interface {
		ProductCustomUpdate(ctx context.Context, info *events.DeviceUpdateInfo) error
		ProductSchemaUpdate(ctx context.Context, info *events.DeviceUpdateInfo) error
		DeviceLogLevelUpdate(ctx context.Context, info *events.DeviceUpdateInfo) error
		DeviceGatewayUpdate(ctx context.Context, info *events.GatewayUpdateInfo) error
		DeviceRemoteConfigUpdate(ctx context.Context, info *events.DeviceUpdateInfo) error
	}
)

func NewDataUpdate(c conf.EventConf) (DataUpdate, error) {
	switch c.Mode {
	case conf.EventModeNats:
		return newNatsClient(c.Nats)
	case conf.EventModeNatsJs:
		return newNatsJsClient(c.Nats)
	}
	return nil, errors.Parameter.AddMsgf("mode:%v not support", c.Mode)
}
