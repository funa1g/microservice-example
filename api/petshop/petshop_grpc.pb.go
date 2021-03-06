// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package delivery

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PetshopClient is the client API for Petshop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetshopClient interface {
	GetPetList(ctx context.Context, in *GetPetListRequest, opts ...grpc.CallOption) (*PetListResponse, error)
	PostPet(ctx context.Context, in *PostPetRequest, opts ...grpc.CallOption) (*PetResponse, error)
	GetPetById(ctx context.Context, in *GetPetRequest, opts ...grpc.CallOption) (*PetResponse, error)
}

type petshopClient struct {
	cc grpc.ClientConnInterface
}

func NewPetshopClient(cc grpc.ClientConnInterface) PetshopClient {
	return &petshopClient{cc}
}

func (c *petshopClient) GetPetList(ctx context.Context, in *GetPetListRequest, opts ...grpc.CallOption) (*PetListResponse, error) {
	out := new(PetListResponse)
	err := c.cc.Invoke(ctx, "/petshop.Petshop/GetPetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petshopClient) PostPet(ctx context.Context, in *PostPetRequest, opts ...grpc.CallOption) (*PetResponse, error) {
	out := new(PetResponse)
	err := c.cc.Invoke(ctx, "/petshop.Petshop/PostPet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petshopClient) GetPetById(ctx context.Context, in *GetPetRequest, opts ...grpc.CallOption) (*PetResponse, error) {
	out := new(PetResponse)
	err := c.cc.Invoke(ctx, "/petshop.Petshop/GetPetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetshopServer is the server API for Petshop service.
// All implementations must embed UnimplementedPetshopServer
// for forward compatibility
type PetshopServer interface {
	GetPetList(context.Context, *GetPetListRequest) (*PetListResponse, error)
	PostPet(context.Context, *PostPetRequest) (*PetResponse, error)
	GetPetById(context.Context, *GetPetRequest) (*PetResponse, error)
	mustEmbedUnimplementedPetshopServer()
}

// UnimplementedPetshopServer must be embedded to have forward compatible implementations.
type UnimplementedPetshopServer struct {
}

func (UnimplementedPetshopServer) GetPetList(context.Context, *GetPetListRequest) (*PetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPetList not implemented")
}
func (UnimplementedPetshopServer) PostPet(context.Context, *PostPetRequest) (*PetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostPet not implemented")
}
func (UnimplementedPetshopServer) GetPetById(context.Context, *GetPetRequest) (*PetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPetById not implemented")
}
func (UnimplementedPetshopServer) mustEmbedUnimplementedPetshopServer() {}

// UnsafePetshopServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetshopServer will
// result in compilation errors.
type UnsafePetshopServer interface {
	mustEmbedUnimplementedPetshopServer()
}

func RegisterPetshopServer(s grpc.ServiceRegistrar, srv PetshopServer) {
	s.RegisterService(&Petshop_ServiceDesc, srv)
}

func _Petshop_GetPetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetshopServer).GetPetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petshop.Petshop/GetPetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetshopServer).GetPetList(ctx, req.(*GetPetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Petshop_PostPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetshopServer).PostPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petshop.Petshop/PostPet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetshopServer).PostPet(ctx, req.(*PostPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Petshop_GetPetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetshopServer).GetPetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petshop.Petshop/GetPetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetshopServer).GetPetById(ctx, req.(*GetPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Petshop_ServiceDesc is the grpc.ServiceDesc for Petshop service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Petshop_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "petshop.Petshop",
	HandlerType: (*PetshopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPetList",
			Handler:    _Petshop_GetPetList_Handler,
		},
		{
			MethodName: "PostPet",
			Handler:    _Petshop_PostPet_Handler,
		},
		{
			MethodName: "GetPetById",
			Handler:    _Petshop_GetPetById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/petshop/delivery/petshop.proto",
}
