package server

import (
	"context"
	"encoding/json"
	"github.com/i-Things/things/shared/clients"
	"github.com/i-Things/things/shared/conf"
	"github.com/i-Things/things/shared/events/topics"
	"github.com/i-Things/things/src/dmsvr/internal/domain/deviceMsg"
	"github.com/nats-io/nats.go"
)

type (
	NatsClient struct {
		client *clients.NatsClient
	}
)

var (
	natsJsConsumerName = "dmsvr"
)

func newNatsClient(conf conf.EventConf) (*NatsClient, error) {
	nc, err := clients.NewNatsClient2(conf.Mode, natsJsConsumerName, conf.Nats)
	if err != nil {
		return nil, err
	}
	return &NatsClient{client: nc}, nil
}

func (n *NatsClient) Subscribe(handle Handle) error {
	err := n.client.QueueSubscribe(topics.DmActionCheckDelay, natsJsConsumerName,
		func(ctx context.Context, msg []byte, natsMsg *nats.Msg) error {
			tempInfo := deviceMsg.PublishMsg{}
			err := json.Unmarshal(msg, &tempInfo)
			if err != nil {
				return err
			}
			return handle(ctx).ActionCheck(&tempInfo)
		})
	if err != nil {
		return err
	}
	return nil
}
