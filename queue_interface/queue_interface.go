package queue_interface

type FnConsume func(queueName string, msg []byte) bool

type IQueue interface {
	Publish(message interface{}) error
	StartConsume(handler FnConsume) error
	Close()
}

type IMultiThreadQueue interface {
	IQueue
	GetThreadCount() int
	SetThreadLimit(limit int)
	WaitQueue()
}
