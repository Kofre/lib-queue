package rabbitmq_repository

import (
	"devgit.kf.com.br/core/lib-queue/queue_interface"
	"github.com/streadway/amqp"
	"strconv"
)

type QueueRepository struct {
	params     QueueRepositoryParams
	connection *amqp.Connection
}

func NewQueueRepository(params QueueRepositoryParams) (*QueueRepository, error) {
	queueRp := QueueRepository{params: params}
	auth := amqp.PlainAuth{Username: queueRp.params.Login(), Password: queueRp.params.Password()}
	var arrAuth []amqp.Authentication
	arrAuth = append(arrAuth, &auth)
	config := amqp.Config{
		SASL: arrAuth,
	}
	conn, err := amqp.DialConfig("amqp://"+queueRp.params.Host()+":"+strconv.Itoa(queueRp.params.Port())+queueRp.params.VHost(), config)
	if err != nil {
		return nil, err
	}
	queueRp.connection = conn
	return &queueRp, nil
}

func (q *QueueRepository) QueueBind(params QueueBindParams) error {
	activeChannel, err := q.connection.Channel()
	if err != nil {
		return err
	}
	if err := activeChannel.QueueBind(
		params.Name(),
		params.Key(),
		params.Exchange(),
		params.NoWait(),
		params.Args()); err != nil {
		return err
	}
	return nil
}

func (q *QueueRepository) QueueDeclare(params QueueParams, withErrorQueue bool) (queue_interface.IQueue, error) {
	if withErrorQueue {
		err, errorQueueName := q.errorQueueDeclare(params)
		if err != nil {
			return nil, err
		}
		args := params.Args()
		args["x-dead-letter-exchange"] = "Error"
		args["x-dead-letter-routing-key"] = errorQueueName
		params.SetArgs(args)
	}
	return q.queueDeclare(params)
}

func (q *QueueRepository) ExchangeDeclare(params QueueParams) error {
	activeChannel, err := q.connection.Channel()
	if err != nil {
		return err
	}
	if err := activeChannel.ExchangeDeclare(
		params.Name(),
		params.Kind(),
		params.Durable(),
		params.AutoDelete(),
		params.Internal(),
		params.NoWait(),
		params.Args()); err != nil {
		return err
	}
	return nil
}

func (q *QueueRepository) errorQueueDeclare(params QueueParams) (error, string) {

	errorQueueName := params.Name() + "-error"

	exchangeParams := NewQueueParams("Error")
	q.ExchangeDeclare(exchangeParams)
	qParam := NewQueueParams(errorQueueName)
	_, err := q.queueDeclare(qParam)

	if err != nil {
		return err, errorQueueName
	}

	err = q.QueueBind(NewQueueBindParams(errorQueueName, errorQueueName, "Error"))
	if err != nil {
		return err, errorQueueName
	}
	return nil, errorQueueName
}

func (q *QueueRepository) queueDeclare(params QueueParams) (queue_interface.IQueue, error) {
	queue, err := NewQueue(params, q.connection)
	if err != nil {
		return nil, err
	}
	return queue, nil
}

func (q *QueueRepository) GetConnection() *amqp.Connection {
	return q.connection
}
