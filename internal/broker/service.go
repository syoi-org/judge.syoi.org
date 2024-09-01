package broker

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
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
	queueMutex   sync.RWMutex
	queues       map[string]*MessageQueue
	deliveryTags map[int64]string
}

func NewService() Service {
	return &service{
		queues: make(map[string]*MessageQueue),
	}
}

func (s *service) Ack(tag int64) error {
	queue, ok := s.deliveryTags[tag]
	if !ok {
		return fmt.Errorf("delivery tag %d not found", tag)
	}
	messageQueue, ok := s.queues[queue]
	if !ok {
		return fmt.Errorf("queue %s not found", queue)
	}
	messageQueue.Ack(tag)
	return nil
}

func (s *service) Consume(queueName string, consumer string, autoAck bool) (<-chan Delivery, error) {
	s.queueMutex.RLock()
	defer s.queueMutex.RUnlock()
	return nil, nil
}

func (s *service) ExchangeDeclare(name string) error {
	return nil
}

func (s *service) Get(queueName string, autoAck bool) (Delivery, bool, error) {
	s.queueMutex.RLock()
	defer s.queueMutex.RUnlock()
	queue, ok := s.queues[queueName]
	if !ok {
		return Delivery{}, false, fmt.Errorf("queue %s not found", queueName)
	}
	delivery, ok := queue.Get(autoAck)
	if !ok {
		return Delivery{}, false, nil
	}
	return *delivery, true, nil
}

func (s *service) Nack(tag int64, requeue bool) error {
	queue, ok := s.deliveryTags[tag]
	if !ok {
		return fmt.Errorf("delivery tag %d not found", tag)
	}
	messageQueue, ok := s.queues[queue]
	if !ok {
		return fmt.Errorf("queue %s not found", queue)
	}
	messageQueue.Nack(tag, requeue)
	return nil
}

func (s *service) Publish(exchange, key string, msg []byte) error {
	return nil
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

func (s *service) createDeliveryTag(queueName string) error {
	s.queueMutex.Lock()
	defer s.queueMutex.Unlock()
	tagInt, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return fmt.Errorf("failed to generate random int: %w", err)
	}
	tag := tagInt.Int64()
	s.deliveryTags[tag] = queueName
	return nil
}
