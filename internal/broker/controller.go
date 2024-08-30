package broker

import (
	"context"

	"github.com/syoi-org/judy/internal/broker/pb"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type controller struct {
	pb.UnimplementedBrokerServer
	Service Service
	Logger  *zap.SugaredLogger
}

type ControllerParams struct {
	fx.In
	Service Service
	Logger  *zap.SugaredLogger
}

func NewController(p ControllerParams) pb.BrokerServer {
	return &controller{
		Service: p.Service,
		Logger:  p.Logger,
	}
}

func (c *controller) Consume(consumeRequest *pb.ConsumeRequest, consumeResponse grpc.ServerStreamingServer[pb.Delivery]) error {
	deliveries, err := c.Service.Consume(consumeRequest.Queue, consumeRequest.Consumer, consumeRequest.AutoAck)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to consume: %v", err)
		c.Logger.Error(err)
		return err
	}
	for delivery := range deliveries {
		deliveryPb := &pb.Delivery{
			DeliveryTag: delivery.DeliveryTag,
			Redelivered: delivery.Redelivered,
			Exchange:    delivery.Exchange,
			RoutingKey:  delivery.RoutingKey,
			Body:        delivery.Body,
		}
		if err := consumeResponse.SendMsg(deliveryPb); err != nil {
			err = status.Errorf(codes.Internal, "failed to send delivery: %v", err)
			c.Logger.Error(err)
			return err
		}
	}
	return nil
}

func (c *controller) ExchangeDeclare(ctx context.Context, exchangeDeclareRequest *pb.ExchangeDeclareRequest) (*pb.ExchangeDeclareResponse, error) {
	if err := c.Service.ExchangeDeclare(exchangeDeclareRequest.Name); err != nil {
		err = status.Errorf(codes.Internal, "failed to declare exchange: %v", err)
		c.Logger.Error(err)
		return nil, err
	}
	return nil, nil
}

func (c *controller) Get(ctx context.Context, getRequest *pb.GetRequest) (*pb.GetResponse, error) {
	delivery, ok, err := c.Service.Get(getRequest.Queue, getRequest.AutoAck)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to get delivery: %v", err)
		c.Logger.Error(err)
		return nil, err
	}
	return &pb.GetResponse{
		Delivery: &pb.Delivery{
			DeliveryTag: delivery.DeliveryTag,
			Redelivered: delivery.Redelivered,
			Exchange:    delivery.Exchange,
			RoutingKey:  delivery.RoutingKey,
			Body:        delivery.Body,
		},
		Ok: ok,
	}, nil
}

func (c *controller) Publish(ctx context.Context, publishRequest *pb.PublishRequest) (*pb.PublishResponse, error) {
	if err := c.Service.Publish(publishRequest.Exchange, publishRequest.Key, publishRequest.Body); err != nil {
		err = status.Errorf(codes.Internal, "failed to publish: %v", err)
		c.Logger.Error(err)
		return nil, err
	}
	return &pb.PublishResponse{}, nil
}

func (c *controller) QueueBind(ctx context.Context, queueBindRequest *pb.QueueBindRequest) (*pb.QueueBindResponse, error) {
	if err := c.Service.QueueBind(queueBindRequest.Name, queueBindRequest.Key, queueBindRequest.Exchange); err != nil {
		err = status.Errorf(codes.Internal, "failed to bind queue: %v", err)
		c.Logger.Error(err)
		return nil, err
	}
	return &pb.QueueBindResponse{}, nil
}

func (c *controller) QueueDeclare(ctx context.Context, queueDeclareRequest *pb.QueueDeclareRequest) (*pb.QueueDeclareResponse, error) {
	queue, err := c.Service.QueueDeclare(queueDeclareRequest.Name)
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to declare queue: %v", err)
		c.Logger.Error(err)
		return nil, err
	}
	return &pb.QueueDeclareResponse{
		Name: queue.Name,
	}, nil
}
