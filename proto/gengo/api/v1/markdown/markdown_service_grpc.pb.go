// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: api/v1/markdown/markdown_service.proto

package markdown

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
	MarkdownService_ParseMarkdown_FullMethodName          = "/memos.api.v1.markdown.MarkdownService/ParseMarkdown"
	MarkdownService_RestoreMarkdownNodes_FullMethodName   = "/memos.api.v1.markdown.MarkdownService/RestoreMarkdownNodes"
	MarkdownService_StringifyMarkdownNodes_FullMethodName = "/memos.api.v1.markdown.MarkdownService/StringifyMarkdownNodes"
	MarkdownService_GetLinkMetadata_FullMethodName        = "/memos.api.v1.markdown.MarkdownService/GetLinkMetadata"
)

// MarkdownServiceClient is the client API for MarkdownService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarkdownServiceClient interface {
	// ParseMarkdown parses the given markdown content and returns a list of nodes.
	ParseMarkdown(ctx context.Context, in *ParseMarkdownRequest, opts ...grpc.CallOption) (*ParseMarkdownResponse, error)
	// RestoreMarkdownNodes restores the given nodes to markdown content.
	RestoreMarkdownNodes(ctx context.Context, in *RestoreMarkdownNodesRequest, opts ...grpc.CallOption) (*RestoreMarkdownNodesResponse, error)
	// StringifyMarkdownNodes stringify the given nodes to plain text content.
	StringifyMarkdownNodes(ctx context.Context, in *StringifyMarkdownNodesRequest, opts ...grpc.CallOption) (*StringifyMarkdownNodesResponse, error)
	// GetLinkMetadata returns metadata for a given link.
	GetLinkMetadata(ctx context.Context, in *GetLinkMetadataRequest, opts ...grpc.CallOption) (*LinkMetadata, error)
}

type markdownServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarkdownServiceClient(cc grpc.ClientConnInterface) MarkdownServiceClient {
	return &markdownServiceClient{cc}
}

func (c *markdownServiceClient) ParseMarkdown(ctx context.Context, in *ParseMarkdownRequest, opts ...grpc.CallOption) (*ParseMarkdownResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ParseMarkdownResponse)
	err := c.cc.Invoke(ctx, MarkdownService_ParseMarkdown_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *markdownServiceClient) RestoreMarkdownNodes(ctx context.Context, in *RestoreMarkdownNodesRequest, opts ...grpc.CallOption) (*RestoreMarkdownNodesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RestoreMarkdownNodesResponse)
	err := c.cc.Invoke(ctx, MarkdownService_RestoreMarkdownNodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *markdownServiceClient) StringifyMarkdownNodes(ctx context.Context, in *StringifyMarkdownNodesRequest, opts ...grpc.CallOption) (*StringifyMarkdownNodesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StringifyMarkdownNodesResponse)
	err := c.cc.Invoke(ctx, MarkdownService_StringifyMarkdownNodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *markdownServiceClient) GetLinkMetadata(ctx context.Context, in *GetLinkMetadataRequest, opts ...grpc.CallOption) (*LinkMetadata, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LinkMetadata)
	err := c.cc.Invoke(ctx, MarkdownService_GetLinkMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarkdownServiceServer is the server API for MarkdownService service.
// All implementations must embed UnimplementedMarkdownServiceServer
// for forward compatibility.
type MarkdownServiceServer interface {
	// ParseMarkdown parses the given markdown content and returns a list of nodes.
	ParseMarkdown(context.Context, *ParseMarkdownRequest) (*ParseMarkdownResponse, error)
	// RestoreMarkdownNodes restores the given nodes to markdown content.
	RestoreMarkdownNodes(context.Context, *RestoreMarkdownNodesRequest) (*RestoreMarkdownNodesResponse, error)
	// StringifyMarkdownNodes stringify the given nodes to plain text content.
	StringifyMarkdownNodes(context.Context, *StringifyMarkdownNodesRequest) (*StringifyMarkdownNodesResponse, error)
	// GetLinkMetadata returns metadata for a given link.
	GetLinkMetadata(context.Context, *GetLinkMetadataRequest) (*LinkMetadata, error)
	mustEmbedUnimplementedMarkdownServiceServer()
}

// UnimplementedMarkdownServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMarkdownServiceServer struct{}

func (UnimplementedMarkdownServiceServer) ParseMarkdown(context.Context, *ParseMarkdownRequest) (*ParseMarkdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseMarkdown not implemented")
}
func (UnimplementedMarkdownServiceServer) RestoreMarkdownNodes(context.Context, *RestoreMarkdownNodesRequest) (*RestoreMarkdownNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreMarkdownNodes not implemented")
}
func (UnimplementedMarkdownServiceServer) StringifyMarkdownNodes(context.Context, *StringifyMarkdownNodesRequest) (*StringifyMarkdownNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StringifyMarkdownNodes not implemented")
}
func (UnimplementedMarkdownServiceServer) GetLinkMetadata(context.Context, *GetLinkMetadataRequest) (*LinkMetadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinkMetadata not implemented")
}
func (UnimplementedMarkdownServiceServer) mustEmbedUnimplementedMarkdownServiceServer() {}
func (UnimplementedMarkdownServiceServer) testEmbeddedByValue()                         {}

// UnsafeMarkdownServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarkdownServiceServer will
// result in compilation errors.
type UnsafeMarkdownServiceServer interface {
	mustEmbedUnimplementedMarkdownServiceServer()
}

func RegisterMarkdownServiceServer(s grpc.ServiceRegistrar, srv MarkdownServiceServer) {
	// If the following call pancis, it indicates UnimplementedMarkdownServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MarkdownService_ServiceDesc, srv)
}

func _MarkdownService_ParseMarkdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseMarkdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarkdownServiceServer).ParseMarkdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarkdownService_ParseMarkdown_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarkdownServiceServer).ParseMarkdown(ctx, req.(*ParseMarkdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarkdownService_RestoreMarkdownNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreMarkdownNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarkdownServiceServer).RestoreMarkdownNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarkdownService_RestoreMarkdownNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarkdownServiceServer).RestoreMarkdownNodes(ctx, req.(*RestoreMarkdownNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarkdownService_StringifyMarkdownNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringifyMarkdownNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarkdownServiceServer).StringifyMarkdownNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarkdownService_StringifyMarkdownNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarkdownServiceServer).StringifyMarkdownNodes(ctx, req.(*StringifyMarkdownNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarkdownService_GetLinkMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLinkMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarkdownServiceServer).GetLinkMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarkdownService_GetLinkMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarkdownServiceServer).GetLinkMetadata(ctx, req.(*GetLinkMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MarkdownService_ServiceDesc is the grpc.ServiceDesc for MarkdownService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarkdownService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "memos.api.v1.markdown.MarkdownService",
	HandlerType: (*MarkdownServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParseMarkdown",
			Handler:    _MarkdownService_ParseMarkdown_Handler,
		},
		{
			MethodName: "RestoreMarkdownNodes",
			Handler:    _MarkdownService_RestoreMarkdownNodes_Handler,
		},
		{
			MethodName: "StringifyMarkdownNodes",
			Handler:    _MarkdownService_StringifyMarkdownNodes_Handler,
		},
		{
			MethodName: "GetLinkMetadata",
			Handler:    _MarkdownService_GetLinkMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/markdown/markdown_service.proto",
}
