// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: api/proto/service.proto

package api

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

// PlaylistClient is the client API for Playlist service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlaylistClient interface {
	Play(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	Pause(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	Next(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	Prev(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	AddSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Response, error)
	GetCurrentSong(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	DeleteSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Response, error)
	UpdateNextSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Response, error)
}

type playlistClient struct {
	cc grpc.ClientConnInterface
}

func NewPlaylistClient(cc grpc.ClientConnInterface) PlaylistClient {
	return &playlistClient{cc}
}

func (c *playlistClient) Play(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Playlist/Play", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) Pause(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Playlist/Pause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) Next(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Playlist/Next", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) Prev(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Playlist/Prev", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) AddSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Playlist/AddSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) GetCurrentSong(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Playlist/GetCurrentSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) DeleteSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Playlist/DeleteSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playlistClient) UpdateNextSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Playlist/UpdateNextSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaylistServer is the server API for Playlist service.
// All implementations must embed UnimplementedPlaylistServer
// for forward compatibility
type PlaylistServer interface {
	Play(context.Context, *Empty) (*Response, error)
	Pause(context.Context, *Empty) (*Response, error)
	Next(context.Context, *Empty) (*Response, error)
	Prev(context.Context, *Empty) (*Response, error)
	AddSong(context.Context, *Song) (*Response, error)
	GetCurrentSong(context.Context, *Empty) (*Response, error)
	DeleteSong(context.Context, *Song) (*Response, error)
	UpdateNextSong(context.Context, *Song) (*Response, error)
	mustEmbedUnimplementedPlaylistServer()
}

// UnimplementedPlaylistServer must be embedded to have forward compatible implementations.
type UnimplementedPlaylistServer struct {
}

func (UnimplementedPlaylistServer) Play(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Play not implemented")
}
func (UnimplementedPlaylistServer) Pause(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedPlaylistServer) Next(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Next not implemented")
}
func (UnimplementedPlaylistServer) Prev(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prev not implemented")
}
func (UnimplementedPlaylistServer) AddSong(context.Context, *Song) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSong not implemented")
}
func (UnimplementedPlaylistServer) GetCurrentSong(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentSong not implemented")
}
func (UnimplementedPlaylistServer) DeleteSong(context.Context, *Song) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSong not implemented")
}
func (UnimplementedPlaylistServer) UpdateNextSong(context.Context, *Song) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNextSong not implemented")
}
func (UnimplementedPlaylistServer) mustEmbedUnimplementedPlaylistServer() {}

// UnsafePlaylistServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlaylistServer will
// result in compilation errors.
type UnsafePlaylistServer interface {
	mustEmbedUnimplementedPlaylistServer()
}

func RegisterPlaylistServer(s grpc.ServiceRegistrar, srv PlaylistServer) {
	s.RegisterService(&Playlist_ServiceDesc, srv)
}

func _Playlist_Play_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).Play(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Playlist/Play",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).Play(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Playlist/Pause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).Pause(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_Next_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).Next(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Playlist/Next",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).Next(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_Prev_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).Prev(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Playlist/Prev",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).Prev(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_AddSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Song)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).AddSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Playlist/AddSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).AddSong(ctx, req.(*Song))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_GetCurrentSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).GetCurrentSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Playlist/GetCurrentSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).GetCurrentSong(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_DeleteSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Song)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).DeleteSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Playlist/DeleteSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).DeleteSong(ctx, req.(*Song))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playlist_UpdateNextSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Song)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaylistServer).UpdateNextSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Playlist/UpdateNextSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaylistServer).UpdateNextSong(ctx, req.(*Song))
	}
	return interceptor(ctx, in, info, handler)
}

// Playlist_ServiceDesc is the grpc.ServiceDesc for Playlist service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Playlist_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Playlist",
	HandlerType: (*PlaylistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Play",
			Handler:    _Playlist_Play_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _Playlist_Pause_Handler,
		},
		{
			MethodName: "Next",
			Handler:    _Playlist_Next_Handler,
		},
		{
			MethodName: "Prev",
			Handler:    _Playlist_Prev_Handler,
		},
		{
			MethodName: "AddSong",
			Handler:    _Playlist_AddSong_Handler,
		},
		{
			MethodName: "GetCurrentSong",
			Handler:    _Playlist_GetCurrentSong_Handler,
		},
		{
			MethodName: "DeleteSong",
			Handler:    _Playlist_DeleteSong_Handler,
		},
		{
			MethodName: "UpdateNextSong",
			Handler:    _Playlist_UpdateNextSong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/service.proto",
}
