// protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: match_api.proto

package match_api

import (
	context "context"
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

type MatchEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MatchEmpty) Reset() {
	*x = MatchEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchEmpty) ProtoMessage() {}

func (x *MatchEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_match_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchEmpty.ProtoReflect.Descriptor instead.
func (*MatchEmpty) Descriptor() ([]byte, []int) {
	return file_match_api_proto_rawDescGZIP(), []int{0}
}

type PriceStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prices *Prices `protobuf:"bytes,1,opt,name=Prices,proto3" json:"Prices,omitempty"`
}

func (x *PriceStreamResponse) Reset() {
	*x = PriceStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceStreamResponse) ProtoMessage() {}

func (x *PriceStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_match_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceStreamResponse.ProtoReflect.Descriptor instead.
func (*PriceStreamResponse) Descriptor() ([]byte, []int) {
	return file_match_api_proto_rawDescGZIP(), []int{1}
}

func (x *PriceStreamResponse) GetPrices() *Prices {
	if x != nil {
		return x.Prices
	}
	return nil
}

type Prices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Time   string `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Bid    string `protobuf:"bytes,3,opt,name=bid,proto3" json:"bid,omitempty"`
	Ask    string `protobuf:"bytes,4,opt,name=ask,proto3" json:"ask,omitempty"`
}

func (x *Prices) Reset() {
	*x = Prices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prices) ProtoMessage() {}

func (x *Prices) ProtoReflect() protoreflect.Message {
	mi := &file_match_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prices.ProtoReflect.Descriptor instead.
func (*Prices) Descriptor() ([]byte, []int) {
	return file_match_api_proto_rawDescGZIP(), []int{2}
}

func (x *Prices) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Prices) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Prices) GetBid() string {
	if x != nil {
		return x.Bid
	}
	return ""
}

func (x *Prices) GetAsk() string {
	if x != nil {
		return x.Ask
	}
	return ""
}

var File_match_api_proto protoreflect.FileDescriptor

var file_match_api_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x0c, 0x0a, 0x0a,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x40, 0x0a, 0x13, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x06, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x06, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x22, 0x58, 0x0a, 0x06,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x62, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x61, 0x73, 0x6b, 0x32, 0x5f, 0x0a, 0x08, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x41,
	0x50, 0x49, 0x12, 0x53, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_match_api_proto_rawDescOnce sync.Once
	file_match_api_proto_rawDescData = file_match_api_proto_rawDesc
)

func file_match_api_proto_rawDescGZIP() []byte {
	file_match_api_proto_rawDescOnce.Do(func() {
		file_match_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_match_api_proto_rawDescData)
	})
	return file_match_api_proto_rawDescData
}

var file_match_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_match_api_proto_goTypes = []interface{}{
	(*MatchEmpty)(nil),          // 0: match_api.MatchEmpty
	(*PriceStreamResponse)(nil), // 1: match_api.PriceStreamResponse
	(*Prices)(nil),              // 2: match_api.Prices
}
var file_match_api_proto_depIdxs = []int32{
	2, // 0: match_api.PriceStreamResponse.Prices:type_name -> match_api.Prices
	0, // 1: match_api.MatchAPI.RequestPricesStreaming:input_type -> match_api.MatchEmpty
	1, // 2: match_api.MatchAPI.RequestPricesStreaming:output_type -> match_api.PriceStreamResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_match_api_proto_init() }
func file_match_api_proto_init() {
	if File_match_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_match_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchEmpty); i {
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
		file_match_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceStreamResponse); i {
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
		file_match_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prices); i {
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
			RawDescriptor: file_match_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_match_api_proto_goTypes,
		DependencyIndexes: file_match_api_proto_depIdxs,
		MessageInfos:      file_match_api_proto_msgTypes,
	}.Build()
	File_match_api_proto = out.File
	file_match_api_proto_rawDesc = nil
	file_match_api_proto_goTypes = nil
	file_match_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MatchAPIClient is the client API for MatchAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MatchAPIClient interface {
	RequestPricesStreaming(ctx context.Context, in *MatchEmpty, opts ...grpc.CallOption) (MatchAPI_RequestPricesStreamingClient, error)
}

type matchAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchAPIClient(cc grpc.ClientConnInterface) MatchAPIClient {
	return &matchAPIClient{cc}
}

func (c *matchAPIClient) RequestPricesStreaming(ctx context.Context, in *MatchEmpty, opts ...grpc.CallOption) (MatchAPI_RequestPricesStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MatchAPI_serviceDesc.Streams[0], "/match_api.MatchAPI/RequestPricesStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &matchAPIRequestPricesStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MatchAPI_RequestPricesStreamingClient interface {
	Recv() (*PriceStreamResponse, error)
	grpc.ClientStream
}

type matchAPIRequestPricesStreamingClient struct {
	grpc.ClientStream
}

func (x *matchAPIRequestPricesStreamingClient) Recv() (*PriceStreamResponse, error) {
	m := new(PriceStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MatchAPIServer is the server API for MatchAPI service.
type MatchAPIServer interface {
	RequestPricesStreaming(*MatchEmpty, MatchAPI_RequestPricesStreamingServer) error
}

// UnimplementedMatchAPIServer can be embedded to have forward compatible implementations.
type UnimplementedMatchAPIServer struct {
}

func (*UnimplementedMatchAPIServer) RequestPricesStreaming(*MatchEmpty, MatchAPI_RequestPricesStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method RequestPricesStreaming not implemented")
}

func RegisterMatchAPIServer(s *grpc.Server, srv MatchAPIServer) {
	s.RegisterService(&_MatchAPI_serviceDesc, srv)
}

func _MatchAPI_RequestPricesStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MatchEmpty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MatchAPIServer).RequestPricesStreaming(m, &matchAPIRequestPricesStreamingServer{stream})
}

type MatchAPI_RequestPricesStreamingServer interface {
	Send(*PriceStreamResponse) error
	grpc.ServerStream
}

type matchAPIRequestPricesStreamingServer struct {
	grpc.ServerStream
}

func (x *matchAPIRequestPricesStreamingServer) Send(m *PriceStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MatchAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "match_api.MatchAPI",
	HandlerType: (*MatchAPIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestPricesStreaming",
			Handler:       _MatchAPI_RequestPricesStreaming_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "match_api.proto",
}