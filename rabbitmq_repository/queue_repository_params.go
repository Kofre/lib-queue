package rabbitmq_repository

type QueueRepositoryParams struct {
	login    string
	password string
	host     string
	vhost    string
	port     int
}

func NewQueueRepositoryParams(login string, password string, host string, port int) QueueRepositoryParams {
	return QueueRepositoryParams{
		login:    login,
		password: password,
		host:     host,
		vhost:    "/",
		port:     port,
	}
}

func (q *QueueRepositoryParams) Port() int {
	return q.port
}

func (q *QueueRepositoryParams) Host() string {
	return q.host
}

func (q *QueueRepositoryParams) Password() string {
	return q.password
}

func (q *QueueRepositoryParams) Login() string {
	return q.login
}

func (q *QueueRepositoryParams) VHost() string {
	return q.vhost
}

func (q *QueueRepositoryParams) SetVHost(vhost string) {
	q.vhost = vhost
}
