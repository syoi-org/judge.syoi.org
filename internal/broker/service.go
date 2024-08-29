package broker

type Service interface {
	Consume(queue string, consumer string, autoAck bool) (<-chan Delivery, error)
	ExchangeDeclare(name string) error
	Publish(exchange, key string, msg []byte) error
	Get(queue string, autoAck bool) (msg Delivery, ok bool, err error)
	QueueBind(name, key, exchange string) error
	QueueDeclare(name string) (Queue, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Consume(queue string, consumer string, autoAck bool) (<-chan Delivery, error) {
	return nil, nil
}

func (s *service) ExchangeDeclare(name string) error {
	return nil
}

func (s *service) Publish(exchange, key string, msg []byte) error {
	return nil
}

func (s *service) Get(queue string, autoAck bool) (Delivery, bool, error) {
	return Delivery{}, false, nil
}

func (s *service) QueueBind(name, key, exchange string) error {
	return nil
}

func (s *service) QueueDeclare(name string) (Queue, error) {
	return Queue{}, nil
}
