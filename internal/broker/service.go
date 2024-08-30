package broker

import (
	"crypto/rand"
	"fmt"
	"sync"

	"github.com/oklog/ulid/v2"
)

type Service interface {
	Consume(queue string, consumer string, autoAck bool) (<-chan Delivery, error)
	ExchangeDeclare(name string) error
	Publish(exchange, key string, msg []byte) error
	Get(queue string, autoAck bool) (msg Delivery, ok bool, err error)
	QueueBind(name, key, exchange string) error
	QueueDeclare(name string) (Queue, error)
}

type service struct {
	queueMutex sync.RWMutex
	queues     map[string]*MessageQueue
}

func NewService() Service {
	return &service{}
}

func (s *service) Consume(queueName string, consumer string, autoAck bool) (<-chan Delivery, error) {
	s.queueMutex.RLock()
	defer s.queueMutex.RUnlock()
	return nil, nil
}

func (s *service) ExchangeDeclare(name string) error {
	return nil
}

func (s *service) Publish(exchange, key string, msg []byte) error {
	return nil
}

func (s *service) Get(queueName string, autoAck bool) (Delivery, bool, error) {
	s.queueMutex.RLock()
	defer s.queueMutex.RUnlock()
	queue := s.queues[queueName]
	delivery, ok := queue.Get(autoAck)
	if !ok {
		return Delivery{}, false, nil
	}
	return *delivery, true, nil
}

func (s *service) QueueBind(name, key, exchange string) error {
	return nil
}

func (s *service) QueueDeclare(name string) (Queue, error) {
	s.queueMutex.Lock()
	defer s.queueMutex.Unlock()
	if queue, ok := s.queues[name]; ok {
		return queue.Queue, nil
	}
	if name == "" {
		randomUlid, err := ulid.New(ulid.Now(), rand.Reader)
		if err != nil {
			return Queue{}, fmt.Errorf("failed to generate random ulid: %w", err)
		}
		name = randomUlid.String()
	}
	queue := NewQueue(name)
	s.queues[name] = queue
	return queue.Queue, nil
}
