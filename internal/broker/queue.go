package broker

import "sync"

type Queue struct {
	Name string
}

type MessageQueue struct {
	Queue
	mu                sync.Mutex
	deliveries        []*Delivery
	unackedDeliveries map[int64]*Delivery
}

func NewQueue(name string) *MessageQueue {
	return &MessageQueue{
		Queue: Queue{
			Name: name,
		},
		deliveries:        make([]*Delivery, 0),
		unackedDeliveries: make(map[int64]*Delivery),
	}
}

func (q *MessageQueue) Ack(deliveryTag int64) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if _, ok := q.unackedDeliveries[deliveryTag]; ok {
		delete(q.unackedDeliveries, deliveryTag)
	}
}

func (q *MessageQueue) Nack(deliveryTag int64, requeue bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if delivery, ok := q.unackedDeliveries[deliveryTag]; ok {
		delete(q.unackedDeliveries, deliveryTag)
		if requeue {
			delivery.Redelivered = true
			q.deliveries = append(q.deliveries, delivery)
		}
	}
}

func (q *MessageQueue) Publish(delivery *Delivery) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.deliveries = append(q.deliveries, delivery)
}

func (q *MessageQueue) Get(autoAck bool) (*Delivery, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.deliveries) == 0 {
		return nil, false
	}
	delivery := q.deliveries[0]
	q.deliveries = q.deliveries[1:]
	if !autoAck {
		q.unackedDeliveries[delivery.DeliveryTag] = delivery
	}
	return delivery, true
}
