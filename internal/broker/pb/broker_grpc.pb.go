// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.4
// source: broker.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Broker_Consume_FullMethodName         = "/Broker/Consume"
	Broker_ExchangeDeclare_FullMethodName = "/Broker/ExchangeDeclare"
	Broker_Get_FullMethodName             = "/Broker/Get"
	Broker_Publish_FullMethodName         = "/Broker/Publish"
	Broker_QueueBind_FullMethodName       = "/Broker/QueueBind"
	Broker_QueueDeclare_FullMethodName    = "/Broker/QueueDeclare"
)

// BrokerClient is the client API for Broker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrokerClient interface {
	Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Delivery], error)
	ExchangeDeclare(ctx context.Context, in *ExchangeDeclareRequest, opts ...grpc.CallOption) (*ExchangeDeclareResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	QueueBind(ctx context.Context, in *QueueBindRequest, opts ...grpc.CallOption) (*QueueBindResponse, error)
	QueueDeclare(ctx context.Context, in *QueueDeclareRequest, opts ...grpc.CallOption) (*QueueDeclareResponse, error)
}

type brokerClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerClient(cc grpc.ClientConnInterface) BrokerClient {
	return &brokerClient{cc}
}

func (c *brokerClient) Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Delivery], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Broker_ServiceDesc.Streams[0], Broker_Consume_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ConsumeRequest, Delivery]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Broker_ConsumeClient = grpc.ServerStreamingClient[Delivery]

func (c *brokerClient) ExchangeDeclare(ctx context.Context, in *ExchangeDeclareRequest, opts ...grpc.CallOption) (*ExchangeDeclareResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExchangeDeclareResponse)
	err := c.cc.Invoke(ctx, Broker_ExchangeDeclare_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, Broker_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, Broker_Publish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) QueueBind(ctx context.Context, in *QueueBindRequest, opts ...grpc.CallOption) (*QueueBindResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueueBindResponse)
	err := c.cc.Invoke(ctx, Broker_QueueBind_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerClient) QueueDeclare(ctx context.Context, in *QueueDeclareRequest, opts ...grpc.CallOption) (*QueueDeclareResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueueDeclareResponse)
	err := c.cc.Invoke(ctx, Broker_QueueDeclare_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerServer is the server API for Broker service.
// All implementations must embed UnimplementedBrokerServer
// for forward compatibility.
type BrokerServer interface {
	Consume(*ConsumeRequest, grpc.ServerStreamingServer[Delivery]) error
	ExchangeDeclare(context.Context, *ExchangeDeclareRequest) (*ExchangeDeclareResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	QueueBind(context.Context, *QueueBindRequest) (*QueueBindResponse, error)
	QueueDeclare(context.Context, *QueueDeclareRequest) (*QueueDeclareResponse, error)
	mustEmbedUnimplementedBrokerServer()
}

// UnimplementedBrokerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBrokerServer struct{}

func (UnimplementedBrokerServer) Consume(*ConsumeRequest, grpc.ServerStreamingServer[Delivery]) error {
	return status.Errorf(codes.Unimplemented, "method Consume not implemented")
}
func (UnimplementedBrokerServer) ExchangeDeclare(context.Context, *ExchangeDeclareRequest) (*ExchangeDeclareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExchangeDeclare not implemented")
}
func (UnimplementedBrokerServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBrokerServer) Publish(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedBrokerServer) QueueBind(context.Context, *QueueBindRequest) (*QueueBindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueueBind not implemented")
}
func (UnimplementedBrokerServer) QueueDeclare(context.Context, *QueueDeclareRequest) (*QueueDeclareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueueDeclare not implemented")
}
func (UnimplementedBrokerServer) mustEmbedUnimplementedBrokerServer() {}
func (UnimplementedBrokerServer) testEmbeddedByValue()                {}

// UnsafeBrokerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerServer will
// result in compilation errors.
type UnsafeBrokerServer interface {
	mustEmbedUnimplementedBrokerServer()
}

func RegisterBrokerServer(s grpc.ServiceRegistrar, srv BrokerServer) {
	// If the following call pancis, it indicates UnimplementedBrokerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Broker_ServiceDesc, srv)
}

func _Broker_Consume_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BrokerServer).Consume(m, &grpc.GenericServerStream[ConsumeRequest, Delivery]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Broker_ConsumeServer = grpc.ServerStreamingServer[Delivery]

func _Broker_ExchangeDeclare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeDeclareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).ExchangeDeclare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Broker_ExchangeDeclare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).ExchangeDeclare(ctx, req.(*ExchangeDeclareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Broker_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Broker_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_QueueBind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueBindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).QueueBind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Broker_QueueBind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).QueueBind(ctx, req.(*QueueBindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broker_QueueDeclare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueDeclareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).QueueDeclare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Broker_QueueDeclare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).QueueDeclare(ctx, req.(*QueueDeclareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Broker_ServiceDesc is the grpc.ServiceDesc for Broker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Broker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Broker",
	HandlerType: (*BrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExchangeDeclare",
			Handler:    _Broker_ExchangeDeclare_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Broker_Get_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _Broker_Publish_Handler,
		},
		{
			MethodName: "QueueBind",
			Handler:    _Broker_QueueBind_Handler,
		},
		{
			MethodName: "QueueDeclare",
			Handler:    _Broker_QueueDeclare_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Consume",
			Handler:       _Broker_Consume_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "broker.proto",
}
