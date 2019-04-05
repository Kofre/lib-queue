package nats_repository

type QueueRepositoryParams struct {
	natsurl  string
	natsport string
}

func NewQueueRepositoryParams(natsurl string, natsport string) QueueRepositoryParams {
	return QueueRepositoryParams{
		natsurl:  natsurl,
		natsport: natsport,
	}
}
func (q *QueueRepositoryParams) NatsUrl() string {
	return q.natsurl
}

func (q *QueueRepositoryParams) NatsPort() string {
	return q.natsport
}
