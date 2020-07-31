// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.1
// source: user.proto

package services

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UserScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserInfo `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserScoreRequest) Reset() {
	*x = UserScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserScoreRequest) ProtoMessage() {}

func (x *UserScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserScoreRequest.ProtoReflect.Descriptor instead.
func (*UserScoreRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserScoreRequest) GetUsers() []*UserInfo {
	if x != nil {
		return x.Users
	}
	return nil
}

type UserScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserInfo `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserScoreResponse) Reset() {
	*x = UserScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserScoreResponse) ProtoMessage() {}

func (x *UserScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserScoreResponse.ProtoReflect.Descriptor instead.
func (*UserScoreResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserScoreResponse) GetUsers() []*UserInfo {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x22, 0x3d, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x32, 0xda, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x59, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x42,
	0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1a, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x52, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x79, 0x54, 0x57, 0x53, 0x12,
	0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_proto_goTypes = []interface{}{
	(*UserScoreRequest)(nil),  // 0: services.UserScoreRequest
	(*UserScoreResponse)(nil), // 1: services.UserScoreResponse
	(*UserInfo)(nil),          // 2: services.UserInfo
}
var file_user_proto_depIdxs = []int32{
	2, // 0: services.UserScoreRequest.users:type_name -> services.UserInfo
	2, // 1: services.UserScoreResponse.users:type_name -> services.UserInfo
	0, // 2: services.UserService.GetUserScore:input_type -> services.UserScoreRequest
	0, // 3: services.UserService.GetUserScoreStream:input_type -> services.UserScoreRequest
	0, // 4: services.UserService.GetUserScoreByClientStream:input_type -> services.UserScoreRequest
	0, // 5: services.UserService.GetUserScoreByTWS:input_type -> services.UserScoreRequest
	1, // 6: services.UserService.GetUserScore:output_type -> services.UserScoreResponse
	1, // 7: services.UserService.GetUserScoreStream:output_type -> services.UserScoreResponse
	1, // 8: services.UserService.GetUserScoreByClientStream:output_type -> services.UserScoreResponse
	1, // 9: services.UserService.GetUserScoreByTWS:output_type -> services.UserScoreResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserScoreRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserScoreResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	//传统方法
	GetUserScore(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error)
	//服务端流模式返回
	GetUserScoreStream(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (UserService_GetUserScoreStreamClient, error)
	//接收客户端流模式请求
	GetUserScoreByClientStream(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserScoreByClientStreamClient, error)
	//双向流
	GetUserScoreByTWS(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserScoreByTWSClient, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserScore(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error) {
	out := new(UserScoreResponse)
	err := c.cc.Invoke(ctx, "/services.UserService/GetUserScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserScoreStream(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (UserService_GetUserScoreStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[0], "/services.UserService/GetUserScoreStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUserScoreStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_GetUserScoreStreamClient interface {
	Recv() (*UserScoreResponse, error)
	grpc.ClientStream
}

type userServiceGetUserScoreStreamClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUserScoreStreamClient) Recv() (*UserScoreResponse, error) {
	m := new(UserScoreResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) GetUserScoreByClientStream(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserScoreByClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[1], "/services.UserService/GetUserScoreByClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUserScoreByClientStreamClient{stream}
	return x, nil
}

type UserService_GetUserScoreByClientStreamClient interface {
	Send(*UserScoreRequest) error
	CloseAndRecv() (*UserScoreResponse, error)
	grpc.ClientStream
}

type userServiceGetUserScoreByClientStreamClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUserScoreByClientStreamClient) Send(m *UserScoreRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceGetUserScoreByClientStreamClient) CloseAndRecv() (*UserScoreResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UserScoreResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) GetUserScoreByTWS(ctx context.Context, opts ...grpc.CallOption) (UserService_GetUserScoreByTWSClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[2], "/services.UserService/GetUserScoreByTWS", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceGetUserScoreByTWSClient{stream}
	return x, nil
}

type UserService_GetUserScoreByTWSClient interface {
	Send(*UserScoreRequest) error
	Recv() (*UserScoreResponse, error)
	grpc.ClientStream
}

type userServiceGetUserScoreByTWSClient struct {
	grpc.ClientStream
}

func (x *userServiceGetUserScoreByTWSClient) Send(m *UserScoreRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceGetUserScoreByTWSClient) Recv() (*UserScoreResponse, error) {
	m := new(UserScoreResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	//传统方法
	GetUserScore(context.Context, *UserScoreRequest) (*UserScoreResponse, error)
	//服务端流模式返回
	GetUserScoreStream(*UserScoreRequest, UserService_GetUserScoreStreamServer) error
	//接收客户端流模式请求
	GetUserScoreByClientStream(UserService_GetUserScoreByClientStreamServer) error
	//双向流
	GetUserScoreByTWS(UserService_GetUserScoreByTWSServer) error
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUserScore(context.Context, *UserScoreRequest) (*UserScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserScore not implemented")
}
func (*UnimplementedUserServiceServer) GetUserScoreStream(*UserScoreRequest, UserService_GetUserScoreStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserScoreStream not implemented")
}
func (*UnimplementedUserServiceServer) GetUserScoreByClientStream(UserService_GetUserScoreByClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserScoreByClientStream not implemented")
}
func (*UnimplementedUserServiceServer) GetUserScoreByTWS(UserService_GetUserScoreByTWSServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserScoreByTWS not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUserScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.UserService/GetUserScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserScore(ctx, req.(*UserScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserScoreStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserScoreRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetUserScoreStream(m, &userServiceGetUserScoreStreamServer{stream})
}

type UserService_GetUserScoreStreamServer interface {
	Send(*UserScoreResponse) error
	grpc.ServerStream
}

type userServiceGetUserScoreStreamServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUserScoreStreamServer) Send(m *UserScoreResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_GetUserScoreByClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).GetUserScoreByClientStream(&userServiceGetUserScoreByClientStreamServer{stream})
}

type UserService_GetUserScoreByClientStreamServer interface {
	SendAndClose(*UserScoreResponse) error
	Recv() (*UserScoreRequest, error)
	grpc.ServerStream
}

type userServiceGetUserScoreByClientStreamServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUserScoreByClientStreamServer) SendAndClose(m *UserScoreResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceGetUserScoreByClientStreamServer) Recv() (*UserScoreRequest, error) {
	m := new(UserScoreRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UserService_GetUserScoreByTWS_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).GetUserScoreByTWS(&userServiceGetUserScoreByTWSServer{stream})
}

type UserService_GetUserScoreByTWSServer interface {
	Send(*UserScoreResponse) error
	Recv() (*UserScoreRequest, error)
	grpc.ServerStream
}

type userServiceGetUserScoreByTWSServer struct {
	grpc.ServerStream
}

func (x *userServiceGetUserScoreByTWSServer) Send(m *UserScoreResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceGetUserScoreByTWSServer) Recv() (*UserScoreRequest, error) {
	m := new(UserScoreRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserScore",
			Handler:    _UserService_GetUserScore_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUserScoreStream",
			Handler:       _UserService_GetUserScoreStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetUserScoreByClientStream",
			Handler:       _UserService_GetUserScoreByClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetUserScoreByTWS",
			Handler:       _UserService_GetUserScoreByTWS_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "user.proto",
}
