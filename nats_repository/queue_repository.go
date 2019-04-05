package nats_repository

import (
	"github.com/nats-io/go-nats"
	"github.com/pkg/errors"
)

type QueueRepository struct {
	natsConn *nats.Conn
}

func NewQueueRepository(params QueueRepositoryParams) (*QueueRepository, error) {
	conn, err := nats.Connect(params.NatsUrl() + ":" + params.NatsPort())
	if err != nil {
		return nil, errors.Wrap(err, "cannot get connection with nats")
	}
	return &QueueRepository{
		natsConn: conn,
	}, nil
}

func (q *QueueRepository) GetConnection() *nats.Conn {
	return q.natsConn
}
