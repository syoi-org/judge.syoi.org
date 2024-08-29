package broker

type Delivery struct {
	DeliveryTag int64
	Redelivered bool
	Exchange    string
	RoutingKey  string
	Body        []byte
}
