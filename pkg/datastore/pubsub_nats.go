package datastore

import (
	"context"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/otyang/icd-10/pkg/logger"
)

var _ IPubSub = (*PubSub)(nil)

type IPubSub interface {
	Publish(ctx context.Context, topic string, data any) error
	Subscribe(ctx context.Context, topic string, subcriptionHandler SubHandler) error
	Close()
}

type SubHandler = nats.MsgHandler

type PubSub struct {
	logger   logger.Interface
	natsConn *nats.Conn
}

func NewNatsFromCredential(url, pathToCredsFile string, log logger.Interface) *PubSub {
	nc, err := nats.Connect(
		url,
		nats.UserCredentials(pathToCredsFile),
	)
	if err != nil {
		log.Fatal("error establishing nats connection -" + err.Error())
	}
	return &PubSub{logger: log, natsConn: nc}
}

func NewNatsCluster(servers []string, log logger.Interface) *PubSub {
	nc, err := nats.Connect(strings.Join(servers, ","))
	if err != nil {
		log.Fatal("error establishing nats connection -" + err.Error())
	}
	return &PubSub{logger: log, natsConn: nc}
}

func NewNatsDefault(log logger.Interface) *PubSub {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("error establishing nats connection -" + err.Error())
	}
	return &PubSub{logger: nil, natsConn: nc}
}

func (ps *PubSub) Publish(ctx context.Context, topic string, data any) error {
	ec, err := nats.NewEncodedConn(ps.natsConn, nats.JSON_ENCODER)
	if err != nil {
		ps.logger.Error("Nats Publisher Encoder Error -" + err.Error())
		return err
	}

	if err := ec.Publish(topic, data); err != nil {
		ps.logger.Error("Nats Publisher Error -" + err.Error())
		return err
	}
	return nil
}

func (ps *PubSub) Subscribe(ctx context.Context, topic string, subcriptionHandler SubHandler) error {
	_, err := ps.natsConn.Subscribe(topic, subcriptionHandler)
	if err != nil {
		ps.logger.Error("Nats Subscriber Error - " + err.Error())
		return err
	}
	return nil
}

func (ps *PubSub) Close() {
	if ps.natsConn != nil {
		if err := ps.natsConn.Drain(); err != nil {
			ps.logger.Error("error draining nats: " + err.Error())
		}
		ps.natsConn.Close()
	}
}
