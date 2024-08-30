package broker

type Message []byte

type Delivery struct {
	DeliveryTag int64
	Redelivered bool
	Exchange    string
	RoutingKey  string
	Body        Message
}
