// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/product.proto

package proto

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

// ProductInfoClient is the client API for ProductInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductInfoClient interface {
	SearchProducts(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (ProductInfo_SearchProductsClient, error)
}

type productInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewProductInfoClient(cc grpc.ClientConnInterface) ProductInfoClient {
	return &productInfoClient{cc}
}

func (c *productInfoClient) SearchProducts(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (ProductInfo_SearchProductsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductInfo_ServiceDesc.Streams[0], "/product.ProductInfo/searchProducts", opts...)
	if err != nil {
		return nil, err
	}
	x := &productInfoSearchProductsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductInfo_SearchProductsClient interface {
	Recv() (*Product, error)
	grpc.ClientStream
}

type productInfoSearchProductsClient struct {
	grpc.ClientStream
}

func (x *productInfoSearchProductsClient) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductInfoServer is the server API for ProductInfo service.
// All implementations must embed UnimplementedProductInfoServer
// for forward compatibility
type ProductInfoServer interface {
	SearchProducts(*SearchQuery, ProductInfo_SearchProductsServer) error
	mustEmbedUnimplementedProductInfoServer()
}

// UnimplementedProductInfoServer must be embedded to have forward compatible implementations.
type UnimplementedProductInfoServer struct {
}

func (UnimplementedProductInfoServer) SearchProducts(*SearchQuery, ProductInfo_SearchProductsServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
}
func (UnimplementedProductInfoServer) mustEmbedUnimplementedProductInfoServer() {}

// UnsafeProductInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductInfoServer will
// result in compilation errors.
type UnsafeProductInfoServer interface {
	mustEmbedUnimplementedProductInfoServer()
}

func RegisterProductInfoServer(s grpc.ServiceRegistrar, srv ProductInfoServer) {
	s.RegisterService(&ProductInfo_ServiceDesc, srv)
}

func _ProductInfo_SearchProducts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductInfoServer).SearchProducts(m, &productInfoSearchProductsServer{stream})
}

type ProductInfo_SearchProductsServer interface {
	Send(*Product) error
	grpc.ServerStream
}

type productInfoSearchProductsServer struct {
	grpc.ServerStream
}

func (x *productInfoSearchProductsServer) Send(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

// ProductInfo_ServiceDesc is the grpc.ServiceDesc for ProductInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductInfo",
	HandlerType: (*ProductInfoServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "searchProducts",
			Handler:       _ProductInfo_SearchProducts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/product.proto",
}
