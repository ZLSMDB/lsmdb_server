// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: api/lsmdb/v1/lsmdb.proto

package v1

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

const (
	Lsmdb_OpenDB_FullMethodName  = "/lsmdb.v1.Lsmdb/OpenDB"
	Lsmdb_Put_FullMethodName     = "/lsmdb.v1.Lsmdb/Put"
	Lsmdb_PutStr_FullMethodName  = "/lsmdb.v1.Lsmdb/PutStr"
	Lsmdb_Get_FullMethodName     = "/lsmdb.v1.Lsmdb/Get"
	Lsmdb_CloseDB_FullMethodName = "/lsmdb.v1.Lsmdb/CloseDB"
)

// LsmdbClient is the client API for Lsmdb service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LsmdbClient interface {
	// Sends a register
	OpenDB(ctx context.Context, in *OpenDBRequest, opts ...grpc.CallOption) (*OpenDBReply, error)
	// rpc Put (stream PutRequest) returns (PutReply) {
	// option (google.api.http) = {
	// post: "/lsmdb/put"
	// body: "*"
	// };
	// }
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutReply, error)
	PutStr(ctx context.Context, in *PutStrRequest, opts ...grpc.CallOption) (*PutStrReply, error)
	// rpc Get (GetRequest) returns (stream GetReply) {
	// option (google.api.http) = {
	// get: "/lsmdb/get/{key}"
	// };
	// }
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	CloseDB(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseReply, error)
}

type lsmdbClient struct {
	cc grpc.ClientConnInterface
}

func NewLsmdbClient(cc grpc.ClientConnInterface) LsmdbClient {
	return &lsmdbClient{cc}
}

func (c *lsmdbClient) OpenDB(ctx context.Context, in *OpenDBRequest, opts ...grpc.CallOption) (*OpenDBReply, error) {
	out := new(OpenDBReply)
	err := c.cc.Invoke(ctx, Lsmdb_OpenDB_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lsmdbClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutReply, error) {
	out := new(PutReply)
	err := c.cc.Invoke(ctx, Lsmdb_Put_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lsmdbClient) PutStr(ctx context.Context, in *PutStrRequest, opts ...grpc.CallOption) (*PutStrReply, error) {
	out := new(PutStrReply)
	err := c.cc.Invoke(ctx, Lsmdb_PutStr_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lsmdbClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, Lsmdb_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lsmdbClient) CloseDB(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseReply, error) {
	out := new(CloseReply)
	err := c.cc.Invoke(ctx, Lsmdb_CloseDB_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LsmdbServer is the server API for Lsmdb service.
// All implementations must embed UnimplementedLsmdbServer
// for forward compatibility
type LsmdbServer interface {
	// Sends a register
	OpenDB(context.Context, *OpenDBRequest) (*OpenDBReply, error)
	// rpc Put (stream PutRequest) returns (PutReply) {
	// option (google.api.http) = {
	// post: "/lsmdb/put"
	// body: "*"
	// };
	// }
	Put(context.Context, *PutRequest) (*PutReply, error)
	PutStr(context.Context, *PutStrRequest) (*PutStrReply, error)
	// rpc Get (GetRequest) returns (stream GetReply) {
	// option (google.api.http) = {
	// get: "/lsmdb/get/{key}"
	// };
	// }
	Get(context.Context, *GetRequest) (*GetReply, error)
	CloseDB(context.Context, *CloseRequest) (*CloseReply, error)
	mustEmbedUnimplementedLsmdbServer()
}

// UnimplementedLsmdbServer must be embedded to have forward compatible implementations.
type UnimplementedLsmdbServer struct {
}

func (UnimplementedLsmdbServer) OpenDB(context.Context, *OpenDBRequest) (*OpenDBReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenDB not implemented")
}
func (UnimplementedLsmdbServer) Put(context.Context, *PutRequest) (*PutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedLsmdbServer) PutStr(context.Context, *PutStrRequest) (*PutStrReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutStr not implemented")
}
func (UnimplementedLsmdbServer) Get(context.Context, *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLsmdbServer) CloseDB(context.Context, *CloseRequest) (*CloseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseDB not implemented")
}
func (UnimplementedLsmdbServer) mustEmbedUnimplementedLsmdbServer() {}

// UnsafeLsmdbServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LsmdbServer will
// result in compilation errors.
type UnsafeLsmdbServer interface {
	mustEmbedUnimplementedLsmdbServer()
}

func RegisterLsmdbServer(s grpc.ServiceRegistrar, srv LsmdbServer) {
	s.RegisterService(&Lsmdb_ServiceDesc, srv)
}

func _Lsmdb_OpenDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LsmdbServer).OpenDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lsmdb_OpenDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LsmdbServer).OpenDB(ctx, req.(*OpenDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lsmdb_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LsmdbServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lsmdb_Put_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LsmdbServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lsmdb_PutStr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutStrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LsmdbServer).PutStr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lsmdb_PutStr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LsmdbServer).PutStr(ctx, req.(*PutStrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lsmdb_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LsmdbServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lsmdb_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LsmdbServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lsmdb_CloseDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LsmdbServer).CloseDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lsmdb_CloseDB_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LsmdbServer).CloseDB(ctx, req.(*CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Lsmdb_ServiceDesc is the grpc.ServiceDesc for Lsmdb service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lsmdb_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lsmdb.v1.Lsmdb",
	HandlerType: (*LsmdbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenDB",
			Handler:    _Lsmdb_OpenDB_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Lsmdb_Put_Handler,
		},
		{
			MethodName: "PutStr",
			Handler:    _Lsmdb_PutStr_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Lsmdb_Get_Handler,
		},
		{
			MethodName: "CloseDB",
			Handler:    _Lsmdb_CloseDB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/lsmdb/v1/lsmdb.proto",
}
