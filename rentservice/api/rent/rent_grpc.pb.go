// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: rent.proto

package proto

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
	RentService_GetRentsByLandlord_FullMethodName = "/rentservice.RentService/GetRentsByLandlord"
	RentService_GetRentsByRenter_FullMethodName   = "/rentservice.RentService/GetRentsByRenter"
	RentService_GetRentedDates_FullMethodName     = "/rentservice.RentService/GetRentedDates"
	RentService_CreateRent_FullMethodName         = "/rentservice.RentService/CreateRent"
	RentService_CloseRent_FullMethodName          = "/rentservice.RentService/CloseRent"
)

// RentServiceClient is the client API for RentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RentServiceClient interface {
	GetRentsByLandlord(ctx context.Context, in *GetRentByLandlordRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetRentsByRenter(ctx context.Context, in *GetRentByRenterRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetRentedDates(ctx context.Context, in *GetRentedDatesRequest, opts ...grpc.CallOption) (*GetRentedDatesResponse, error)
	CreateRent(ctx context.Context, in *CreateRentRequest, opts ...grpc.CallOption) (*CreateRentResponse, error)
	CloseRent(ctx context.Context, in *CloseRentRequest, opts ...grpc.CallOption) (*CloseRentResponse, error)
}

type rentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRentServiceClient(cc grpc.ClientConnInterface) RentServiceClient {
	return &rentServiceClient{cc}
}

func (c *rentServiceClient) GetRentsByLandlord(ctx context.Context, in *GetRentByLandlordRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, RentService_GetRentsByLandlord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentServiceClient) GetRentsByRenter(ctx context.Context, in *GetRentByRenterRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, RentService_GetRentsByRenter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentServiceClient) GetRentedDates(ctx context.Context, in *GetRentedDatesRequest, opts ...grpc.CallOption) (*GetRentedDatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRentedDatesResponse)
	err := c.cc.Invoke(ctx, RentService_GetRentedDates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentServiceClient) CreateRent(ctx context.Context, in *CreateRentRequest, opts ...grpc.CallOption) (*CreateRentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRentResponse)
	err := c.cc.Invoke(ctx, RentService_CreateRent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentServiceClient) CloseRent(ctx context.Context, in *CloseRentRequest, opts ...grpc.CallOption) (*CloseRentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CloseRentResponse)
	err := c.cc.Invoke(ctx, RentService_CloseRent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RentServiceServer is the server API for RentService service.
// All implementations must embed UnimplementedRentServiceServer
// for forward compatibility.
type RentServiceServer interface {
	GetRentsByLandlord(context.Context, *GetRentByLandlordRequest) (*GetResponse, error)
	GetRentsByRenter(context.Context, *GetRentByRenterRequest) (*GetResponse, error)
	GetRentedDates(context.Context, *GetRentedDatesRequest) (*GetRentedDatesResponse, error)
	CreateRent(context.Context, *CreateRentRequest) (*CreateRentResponse, error)
	CloseRent(context.Context, *CloseRentRequest) (*CloseRentResponse, error)
	mustEmbedUnimplementedRentServiceServer()
}

// UnimplementedRentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRentServiceServer struct{}

func (UnimplementedRentServiceServer) GetRentsByLandlord(context.Context, *GetRentByLandlordRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRentsByLandlord not implemented")
}
func (UnimplementedRentServiceServer) GetRentsByRenter(context.Context, *GetRentByRenterRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRentsByRenter not implemented")
}
func (UnimplementedRentServiceServer) GetRentedDates(context.Context, *GetRentedDatesRequest) (*GetRentedDatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRentedDates not implemented")
}
func (UnimplementedRentServiceServer) CreateRent(context.Context, *CreateRentRequest) (*CreateRentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRent not implemented")
}
func (UnimplementedRentServiceServer) CloseRent(context.Context, *CloseRentRequest) (*CloseRentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseRent not implemented")
}
func (UnimplementedRentServiceServer) mustEmbedUnimplementedRentServiceServer() {}
func (UnimplementedRentServiceServer) testEmbeddedByValue()                     {}

// UnsafeRentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RentServiceServer will
// result in compilation errors.
type UnsafeRentServiceServer interface {
	mustEmbedUnimplementedRentServiceServer()
}

func RegisterRentServiceServer(s grpc.ServiceRegistrar, srv RentServiceServer) {
	// If the following call pancis, it indicates UnimplementedRentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RentService_ServiceDesc, srv)
}

func _RentService_GetRentsByLandlord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRentByLandlordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).GetRentsByLandlord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RentService_GetRentsByLandlord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).GetRentsByLandlord(ctx, req.(*GetRentByLandlordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentService_GetRentsByRenter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRentByRenterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).GetRentsByRenter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RentService_GetRentsByRenter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).GetRentsByRenter(ctx, req.(*GetRentByRenterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentService_GetRentedDates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRentedDatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).GetRentedDates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RentService_GetRentedDates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).GetRentedDates(ctx, req.(*GetRentedDatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentService_CreateRent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).CreateRent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RentService_CreateRent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).CreateRent(ctx, req.(*CreateRentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentService_CloseRent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentServiceServer).CloseRent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RentService_CloseRent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentServiceServer).CloseRent(ctx, req.(*CloseRentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RentService_ServiceDesc is the grpc.ServiceDesc for RentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rentservice.RentService",
	HandlerType: (*RentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRentsByLandlord",
			Handler:    _RentService_GetRentsByLandlord_Handler,
		},
		{
			MethodName: "GetRentsByRenter",
			Handler:    _RentService_GetRentsByRenter_Handler,
		},
		{
			MethodName: "GetRentedDates",
			Handler:    _RentService_GetRentedDates_Handler,
		},
		{
			MethodName: "CreateRent",
			Handler:    _RentService_CreateRent_Handler,
		},
		{
			MethodName: "CloseRent",
			Handler:    _RentService_CloseRent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rent.proto",
}
