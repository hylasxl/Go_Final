// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: api/proto/airport.proto

package airportpb

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
	AirportService_CreateAirport_FullMethodName                               = "/auth.AirportService/CreateAirport"
	AirportService_GetListAirports_FullMethodName                             = "/auth.AirportService/GetListAirports"
	AirportService_GetAirport_FullMethodName                                  = "/auth.AirportService/GetAirport"
	AirportService_UpdateAirport_FullMethodName                               = "/auth.AirportService/UpdateAirport"
	AirportService_DeleteAirport_FullMethodName                               = "/auth.AirportService/DeleteAirport"
	AirportService_AirportGetDepartureFlights_FullMethodName                  = "/auth.AirportService/AirportGetDepartureFlights"
	AirportService_AirportGetArrivalFlights_FullMethodName                    = "/auth.AirportService/AirportGetArrivalFlights"
	AirportService_AirportGetDepartureFlightsAndArrivalFlights_FullMethodName = "/auth.AirportService/AirportGetDepartureFlightsAndArrivalFlights"
)

// AirportServiceClient is the client API for AirportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AirportServiceClient interface {
	CreateAirport(ctx context.Context, in *CreateAirportRequest, opts ...grpc.CallOption) (*AirportResponse, error)
	GetListAirports(ctx context.Context, in *GetListAirportRequest, opts ...grpc.CallOption) (*GetListAirportResponse, error)
	GetAirport(ctx context.Context, in *GetAirportRequest, opts ...grpc.CallOption) (*GetAirportResponse, error)
	UpdateAirport(ctx context.Context, in *UpdateAirportRequest, opts ...grpc.CallOption) (*AirportResponse, error)
	DeleteAirport(ctx context.Context, in *DeleteAirportRequest, opts ...grpc.CallOption) (*AirportResponse, error)
	AirportGetDepartureFlights(ctx context.Context, in *AirportGetDepartureFlightRequest, opts ...grpc.CallOption) (*AirportGetDepartureFlightResponse, error)
	AirportGetArrivalFlights(ctx context.Context, in *AirportGetArrivalFlightRequest, opts ...grpc.CallOption) (*AirportGetArrivalFlightResponse, error)
	AirportGetDepartureFlightsAndArrivalFlights(ctx context.Context, in *AirportGetDepartureFlightsAndArrivalFlightRequest, opts ...grpc.CallOption) (*AirportGetDepartureFlightsAndArrivalFlightResponse, error)
}

type airportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAirportServiceClient(cc grpc.ClientConnInterface) AirportServiceClient {
	return &airportServiceClient{cc}
}

func (c *airportServiceClient) CreateAirport(ctx context.Context, in *CreateAirportRequest, opts ...grpc.CallOption) (*AirportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AirportResponse)
	err := c.cc.Invoke(ctx, AirportService_CreateAirport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airportServiceClient) GetListAirports(ctx context.Context, in *GetListAirportRequest, opts ...grpc.CallOption) (*GetListAirportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetListAirportResponse)
	err := c.cc.Invoke(ctx, AirportService_GetListAirports_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airportServiceClient) GetAirport(ctx context.Context, in *GetAirportRequest, opts ...grpc.CallOption) (*GetAirportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAirportResponse)
	err := c.cc.Invoke(ctx, AirportService_GetAirport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airportServiceClient) UpdateAirport(ctx context.Context, in *UpdateAirportRequest, opts ...grpc.CallOption) (*AirportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AirportResponse)
	err := c.cc.Invoke(ctx, AirportService_UpdateAirport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airportServiceClient) DeleteAirport(ctx context.Context, in *DeleteAirportRequest, opts ...grpc.CallOption) (*AirportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AirportResponse)
	err := c.cc.Invoke(ctx, AirportService_DeleteAirport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airportServiceClient) AirportGetDepartureFlights(ctx context.Context, in *AirportGetDepartureFlightRequest, opts ...grpc.CallOption) (*AirportGetDepartureFlightResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AirportGetDepartureFlightResponse)
	err := c.cc.Invoke(ctx, AirportService_AirportGetDepartureFlights_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airportServiceClient) AirportGetArrivalFlights(ctx context.Context, in *AirportGetArrivalFlightRequest, opts ...grpc.CallOption) (*AirportGetArrivalFlightResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AirportGetArrivalFlightResponse)
	err := c.cc.Invoke(ctx, AirportService_AirportGetArrivalFlights_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airportServiceClient) AirportGetDepartureFlightsAndArrivalFlights(ctx context.Context, in *AirportGetDepartureFlightsAndArrivalFlightRequest, opts ...grpc.CallOption) (*AirportGetDepartureFlightsAndArrivalFlightResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AirportGetDepartureFlightsAndArrivalFlightResponse)
	err := c.cc.Invoke(ctx, AirportService_AirportGetDepartureFlightsAndArrivalFlights_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AirportServiceServer is the server API for AirportService service.
// All implementations must embed UnimplementedAirportServiceServer
// for forward compatibility.
type AirportServiceServer interface {
	CreateAirport(context.Context, *CreateAirportRequest) (*AirportResponse, error)
	GetListAirports(context.Context, *GetListAirportRequest) (*GetListAirportResponse, error)
	GetAirport(context.Context, *GetAirportRequest) (*GetAirportResponse, error)
	UpdateAirport(context.Context, *UpdateAirportRequest) (*AirportResponse, error)
	DeleteAirport(context.Context, *DeleteAirportRequest) (*AirportResponse, error)
	AirportGetDepartureFlights(context.Context, *AirportGetDepartureFlightRequest) (*AirportGetDepartureFlightResponse, error)
	AirportGetArrivalFlights(context.Context, *AirportGetArrivalFlightRequest) (*AirportGetArrivalFlightResponse, error)
	AirportGetDepartureFlightsAndArrivalFlights(context.Context, *AirportGetDepartureFlightsAndArrivalFlightRequest) (*AirportGetDepartureFlightsAndArrivalFlightResponse, error)
	mustEmbedUnimplementedAirportServiceServer()
}

// UnimplementedAirportServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAirportServiceServer struct{}

func (UnimplementedAirportServiceServer) CreateAirport(context.Context, *CreateAirportRequest) (*AirportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAirport not implemented")
}
func (UnimplementedAirportServiceServer) GetListAirports(context.Context, *GetListAirportRequest) (*GetListAirportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListAirports not implemented")
}
func (UnimplementedAirportServiceServer) GetAirport(context.Context, *GetAirportRequest) (*GetAirportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAirport not implemented")
}
func (UnimplementedAirportServiceServer) UpdateAirport(context.Context, *UpdateAirportRequest) (*AirportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAirport not implemented")
}
func (UnimplementedAirportServiceServer) DeleteAirport(context.Context, *DeleteAirportRequest) (*AirportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAirport not implemented")
}
func (UnimplementedAirportServiceServer) AirportGetDepartureFlights(context.Context, *AirportGetDepartureFlightRequest) (*AirportGetDepartureFlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AirportGetDepartureFlights not implemented")
}
func (UnimplementedAirportServiceServer) AirportGetArrivalFlights(context.Context, *AirportGetArrivalFlightRequest) (*AirportGetArrivalFlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AirportGetArrivalFlights not implemented")
}
func (UnimplementedAirportServiceServer) AirportGetDepartureFlightsAndArrivalFlights(context.Context, *AirportGetDepartureFlightsAndArrivalFlightRequest) (*AirportGetDepartureFlightsAndArrivalFlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AirportGetDepartureFlightsAndArrivalFlights not implemented")
}
func (UnimplementedAirportServiceServer) mustEmbedUnimplementedAirportServiceServer() {}
func (UnimplementedAirportServiceServer) testEmbeddedByValue()                        {}

// UnsafeAirportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AirportServiceServer will
// result in compilation errors.
type UnsafeAirportServiceServer interface {
	mustEmbedUnimplementedAirportServiceServer()
}

func RegisterAirportServiceServer(s grpc.ServiceRegistrar, srv AirportServiceServer) {
	// If the following call pancis, it indicates UnimplementedAirportServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AirportService_ServiceDesc, srv)
}

func _AirportService_CreateAirport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAirportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirportServiceServer).CreateAirport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AirportService_CreateAirport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirportServiceServer).CreateAirport(ctx, req.(*CreateAirportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AirportService_GetListAirports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListAirportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirportServiceServer).GetListAirports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AirportService_GetListAirports_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirportServiceServer).GetListAirports(ctx, req.(*GetListAirportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AirportService_GetAirport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAirportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirportServiceServer).GetAirport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AirportService_GetAirport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirportServiceServer).GetAirport(ctx, req.(*GetAirportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AirportService_UpdateAirport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAirportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirportServiceServer).UpdateAirport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AirportService_UpdateAirport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirportServiceServer).UpdateAirport(ctx, req.(*UpdateAirportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AirportService_DeleteAirport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAirportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirportServiceServer).DeleteAirport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AirportService_DeleteAirport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirportServiceServer).DeleteAirport(ctx, req.(*DeleteAirportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AirportService_AirportGetDepartureFlights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AirportGetDepartureFlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirportServiceServer).AirportGetDepartureFlights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AirportService_AirportGetDepartureFlights_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirportServiceServer).AirportGetDepartureFlights(ctx, req.(*AirportGetDepartureFlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AirportService_AirportGetArrivalFlights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AirportGetArrivalFlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirportServiceServer).AirportGetArrivalFlights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AirportService_AirportGetArrivalFlights_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirportServiceServer).AirportGetArrivalFlights(ctx, req.(*AirportGetArrivalFlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AirportService_AirportGetDepartureFlightsAndArrivalFlights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AirportGetDepartureFlightsAndArrivalFlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirportServiceServer).AirportGetDepartureFlightsAndArrivalFlights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AirportService_AirportGetDepartureFlightsAndArrivalFlights_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirportServiceServer).AirportGetDepartureFlightsAndArrivalFlights(ctx, req.(*AirportGetDepartureFlightsAndArrivalFlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AirportService_ServiceDesc is the grpc.ServiceDesc for AirportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AirportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AirportService",
	HandlerType: (*AirportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAirport",
			Handler:    _AirportService_CreateAirport_Handler,
		},
		{
			MethodName: "GetListAirports",
			Handler:    _AirportService_GetListAirports_Handler,
		},
		{
			MethodName: "GetAirport",
			Handler:    _AirportService_GetAirport_Handler,
		},
		{
			MethodName: "UpdateAirport",
			Handler:    _AirportService_UpdateAirport_Handler,
		},
		{
			MethodName: "DeleteAirport",
			Handler:    _AirportService_DeleteAirport_Handler,
		},
		{
			MethodName: "AirportGetDepartureFlights",
			Handler:    _AirportService_AirportGetDepartureFlights_Handler,
		},
		{
			MethodName: "AirportGetArrivalFlights",
			Handler:    _AirportService_AirportGetArrivalFlights_Handler,
		},
		{
			MethodName: "AirportGetDepartureFlightsAndArrivalFlights",
			Handler:    _AirportService_AirportGetDepartureFlightsAndArrivalFlights_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/airport.proto",
}