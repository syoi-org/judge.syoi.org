package broker

import "testing"

func TestQueueDeclare(t *testing.T) {
	brokerService := NewService()
	queue, err := brokerService.QueueDeclare("test")
	if err != nil {
		t.Fatal(err)
	}
	if queue.Name != "test" {
		t.Errorf("expected queue name to be 'test', got %s", queue.Name)
	}
	if _, ok := brokerService.(*service).queues["test"]; !ok {
		t.Errorf("expected queue to be in service.queues")
	}
}

func TestQueueDeclareWithRandomName(t *testing.T) {
	brokerService := NewService()
	queue, err := brokerService.QueueDeclare("")
	if err != nil {
		t.Fatal(err)
	}
	if queue.Name == "" {
		t.Errorf("expected queue name to be non-empty, got %s", queue.Name)
	}
	if _, ok := brokerService.(*service).queues[queue.Name]; !ok {
		t.Errorf("expected queue to be in service.queues")
	}
}
