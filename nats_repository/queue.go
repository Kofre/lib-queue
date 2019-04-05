package nats_repository

import (
	"devgit.kf.com.br/core/lib-queue/queue_interface"
	"github.com/nats-io/go-nats"
	"github.com/pkg/errors"
)

type Queue struct {
	conn    *nats.Conn
	subject string
}

func NewQueue(conn *nats.Conn, subject string) (*Queue, error) {
	return &Queue{
		conn:    conn,
		subject: subject,
	}, nil
}

func (q *Queue) Publish(message interface{}) error {
	m, ok := message.([]byte)
	if !ok {
		return errors.New("Wrong message type on publish")
	}
	return q.conn.Publish(q.subject, m)
}

func (q *Queue) StartConsume(handler queue_interface.FnConsume) error {
	_, err := q.conn.Subscribe(q.subject, func(m *nats.Msg) {
		handler(q.subject, m.Data)
	})
	if err != nil {
		return errors.Wrap(err, "Error consuming on nats queue")
	}
	return nil
}

func (q *Queue) Close() {
	q.conn.Close()
}
