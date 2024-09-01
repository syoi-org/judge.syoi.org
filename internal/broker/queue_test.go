package broker

import "testing"

func TestAck(t *testing.T) {
	queue := NewQueue("test")
	delivery := Delivery{
		DeliveryTag: 1,
		Exchange:    "test",
		RoutingKey:  "test",
		Body:        []byte("test"),
	}
	queue.unackedDeliveries[1] = &delivery
	queue.Ack(1)
	if len(queue.unackedDeliveries) != 0 {
		t.Errorf("expected unacked deliveries to be 0, got %d", len(queue.unackedDeliveries))
	}
}

func TestNack(t *testing.T) {
	queue := NewQueue("test")
	delivery := Delivery{
		DeliveryTag: 1,
		Exchange:    "test",
		RoutingKey:  "test",
		Body:        []byte("test"),
	}
	queue.unackedDeliveries[1] = &delivery
	delivery2 := Delivery{
		DeliveryTag: 2,
		Exchange:    "test",
		RoutingKey:  "test",
		Body:        []byte("test"),
	}
	queue.unackedDeliveries[2] = &delivery2
	queue.Nack(1, true)
	queue.Nack(2, false)
	if len(queue.unackedDeliveries) != 0 {
		t.Errorf("expected unacked deliveries to be 0, got %d", len(queue.unackedDeliveries))
	}
	if len(queue.deliveries) != 1 {
		t.Errorf("expected deliveries to have length 1, got %d", len(queue.deliveries))
	}
	if queue.deliveries[0] != &delivery {
		t.Errorf("expected delivery to be %v, got %v", &delivery, queue.deliveries[0])
	}
}

func TestQueuePublish(t *testing.T) {
	queue := NewQueue("test")
	delivery := Delivery{
		DeliveryTag: 1,
		Exchange:    "test",
		RoutingKey:  "test",
		Body:        []byte("test"),
	}
	queue.Publish(&delivery)
	if len(queue.deliveries) != 1 {
		t.Errorf("expected deliveries to be 1, got %d", len(queue.deliveries))
	}
	if queue.deliveries[0] != &delivery {
		t.Errorf("expected delivery to be %v, got %v", &delivery, queue.deliveries[0])
	}
}

func TestQueueGet(t *testing.T) {
	queue := NewQueue("test")
	delivery := Delivery{
		DeliveryTag: 1,
		Exchange:    "test",
		RoutingKey:  "test",
		Body:        []byte("test"),
	}
	delivery2 := Delivery{
		DeliveryTag: 2,
		Exchange:    "test",
		RoutingKey:  "test",
		Body:        []byte("test"),
	}
	queue.deliveries = append(queue.deliveries, &delivery, &delivery2)
	gotDelivery, ok := queue.Get(true)
	if !ok {
		t.Errorf("expected delivery to be available")
	}
	if gotDelivery != &delivery {
		t.Errorf("expected delivery to be %v, got %v", &delivery, gotDelivery)
	}
	if len(queue.deliveries) != 1 {
		t.Errorf("expected deliveries to have length 1, got %d", len(queue.deliveries))
	}
	if len(queue.unackedDeliveries) != 0 {
		t.Errorf("expected unacked deliveries to have length 0, got %d", len(queue.unackedDeliveries))
	}
	gotDelivery, ok = queue.Get(false)
	if !ok {
		t.Errorf("expected delivery to be available")
	}
	if gotDelivery != &delivery2 {
		t.Errorf("expected delivery to be %v, got %v", &delivery2, gotDelivery)
	}
	if len(queue.deliveries) != 0 {
		t.Errorf("expected deliveries to have length 0, got %d", len(queue.deliveries))
	}
	if len(queue.unackedDeliveries) != 1 {
		t.Errorf("expected unacked deliveries to have length 1, got %d", len(queue.unackedDeliveries))
	}
}
