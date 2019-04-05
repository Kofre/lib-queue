package nats_repository

import (
	"github.com/onsi/gomega"
	"testing"
)

func TestQueueRepositoryNewQueue(t *testing.T) {
	gomega.RegisterTestingT(t)
	qrParams := NewQueueRepositoryParams("srvqueue.gatewaygps.interno.ntopus.com.br", "4222")
	qR, err := NewQueueRepository(qrParams)
	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
	q, err := NewQueue(qR.natsConn, "testSubject")
	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
	gomega.Expect("testSubject").Should(gomega.BeEquivalentTo(q.subject))
	gomega.Expect(qR.natsConn).ShouldNot(gomega.BeNil())
}

//
//func TestQueueRepositoryPublish(t *testing.T) {
//	gomega.RegisterTestingT(t)
//	mu := sync.Mutex{}
//	mu.Lock()
//	receivedMsg := ""
//	mu.Unlock()
//	conn := getMockNatsConn()
//	_, err := conn.Subscribe("testSubject", func(msg *nats.Msg) {
//		mu.Lock()
//		receivedMsg = string(msg.Data)
//		mu.Unlock()
//	})
//	time.Sleep(100 * time.Millisecond)
//	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
//	qR, err := NewQueueRepository("srvqueue.gatewaygps.interno.ntopus.com.br", "4222", "testSubject")
//	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
//	err = qR.Publish([]byte("messagePublished"))
//	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
//	gomega.Eventually(func() string {
//		mu.Lock()
//		defer mu.Unlock()
//		return receivedMsg
//	}).Should(gomega.BeEquivalentTo("messagePublished"))
//}
//
//func TestQueueRepositoryPublishWrongMessage(t *testing.T) {
//	gomega.RegisterTestingT(t)
//	mu := sync.Mutex{}
//	mu.Lock()
//	receivedMsg := ""
//	mu.Unlock()
//	conn := getMockNatsConn()
//	_, err := conn.Subscribe("testSubject", func(msg *nats.Msg) {
//		mu.Lock()
//		receivedMsg = string(msg.Data)
//		mu.Unlock()
//	})
//	time.Sleep(100 * time.Millisecond)
//	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
//	qR, err := NewQueueRepository("srvqueue.gatewaygps.interno.ntopus.com.br", "4222", "testSubject")
//	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
//	err = qR.Publish("messagePublished")
//	gomega.Expect(err).Should(gomega.HaveOccurred())
//	gomega.Eventually(func() string {
//		mu.Lock()
//		defer mu.Unlock()
//		return receivedMsg
//	}).Should(gomega.BeEquivalentTo(""))
//}
//
//func TestQueueRepositoryConsuming(t *testing.T) {
//	gomega.RegisterTestingT(t)
//	mu := sync.Mutex{}
//	mu.Lock()
//	receivedMsg := ""
//	mu.Unlock()
//	conn := getMockNatsConn()
//	qR, err := NewQueueRepository("srvqueue.gatewaygps.interno.ntopus.com.br", "4222", "testSubject")
//	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
//	err = qR.StartConsume(func(msg []byte) bool {
//		receivedMsg = string(msg)
//		return true
//	})
//	time.Sleep(100 * time.Millisecond)
//	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
//	err = conn.Publish("testSubject", []byte("message"))
//	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
//	gomega.Eventually(func() string {
//		mu.Lock()
//		defer mu.Unlock()
//		return receivedMsg
//	}).Should(gomega.BeEquivalentTo("message"))
//}
//
//func getMockNatsConn() *nats.Conn {
//	conn, err := nats.Connect("srvqueue.gatewaygps.interno.ntopus.com.br:4222")
//	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
//	return conn
//}
