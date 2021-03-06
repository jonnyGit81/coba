// protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: price_api.proto

package price_api

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
		mi := &file_price_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prices) ProtoMessage() {}

func (x *Prices) ProtoReflect() protoreflect.Message {
	mi := &file_price_api_proto_msgTypes[0]
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
	return file_price_api_proto_rawDescGZIP(), []int{0}
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

type Acknowledge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack string `protobuf:"bytes,1,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (x *Acknowledge) Reset() {
	*x = Acknowledge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Acknowledge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acknowledge) ProtoMessage() {}

func (x *Acknowledge) ProtoReflect() protoreflect.Message {
	mi := &file_price_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acknowledge.ProtoReflect.Descriptor instead.
func (*Acknowledge) Descriptor() ([]byte, []int) {
	return file_price_api_proto_rawDescGZIP(), []int{1}
}

func (x *Acknowledge) GetAck() string {
	if x != nil {
		return x.Ack
	}
	return ""
}

var File_price_api_proto protoreflect.FileDescriptor

var file_price_api_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x58, 0x0a, 0x06, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x73, 0x6b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x73, 0x6b, 0x22, 0x1f, 0x0a, 0x0b, 0x41,
	0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x32, 0x51, 0x0a, 0x03,
	0x41, 0x50, 0x49, 0x12, 0x4a, 0x0a, 0x2d, 0x4d, 0x45, 0x61, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4f, 0x66, 0x50, 0x45, 0x41, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x50, 0x45, 0x12, 0x07, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x0c, 0x2e,
	0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_price_api_proto_rawDescOnce sync.Once
	file_price_api_proto_rawDescData = file_price_api_proto_rawDesc
)

func file_price_api_proto_rawDescGZIP() []byte {
	file_price_api_proto_rawDescOnce.Do(func() {
		file_price_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_price_api_proto_rawDescData)
	})
	return file_price_api_proto_rawDescData
}

var file_price_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_price_api_proto_goTypes = []interface{}{
	(*Prices)(nil),      // 0: Prices
	(*Acknowledge)(nil), // 1: Acknowledge
}
var file_price_api_proto_depIdxs = []int32{
	0, // 0: API.MEasClientOfPEAndStreamingSendPriceUpdateToPE:input_type -> Prices
	1, // 1: API.MEasClientOfPEAndStreamingSendPriceUpdateToPE:output_type -> Acknowledge
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_price_api_proto_init() }
func file_price_api_proto_init() {
	if File_price_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_price_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_price_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Acknowledge); i {
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
			RawDescriptor: file_price_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_price_api_proto_goTypes,
		DependencyIndexes: file_price_api_proto_depIdxs,
		MessageInfos:      file_price_api_proto_msgTypes,
	}.Build()
	File_price_api_proto = out.File
	file_price_api_proto_rawDesc = nil
	file_price_api_proto_goTypes = nil
	file_price_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	MEasClientOfPEAndStreamingSendPriceUpdateToPE(ctx context.Context, opts ...grpc.CallOption) (API_MEasClientOfPEAndStreamingSendPriceUpdateToPEClient, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) MEasClientOfPEAndStreamingSendPriceUpdateToPE(ctx context.Context, opts ...grpc.CallOption) (API_MEasClientOfPEAndStreamingSendPriceUpdateToPEClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[0], "/API/MEasClientOfPEAndStreamingSendPriceUpdateToPE", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIMEasClientOfPEAndStreamingSendPriceUpdateToPEClient{stream}
	return x, nil
}

type API_MEasClientOfPEAndStreamingSendPriceUpdateToPEClient interface {
	Send(*Prices) error
	CloseAndRecv() (*Acknowledge, error)
	grpc.ClientStream
}

type aPIMEasClientOfPEAndStreamingSendPriceUpdateToPEClient struct {
	grpc.ClientStream
}

func (x *aPIMEasClientOfPEAndStreamingSendPriceUpdateToPEClient) Send(m *Prices) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aPIMEasClientOfPEAndStreamingSendPriceUpdateToPEClient) CloseAndRecv() (*Acknowledge, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Acknowledge)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	MEasClientOfPEAndStreamingSendPriceUpdateToPE(API_MEasClientOfPEAndStreamingSendPriceUpdateToPEServer) error
}

// UnimplementedAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (*UnimplementedAPIServer) MEasClientOfPEAndStreamingSendPriceUpdateToPE(API_MEasClientOfPEAndStreamingSendPriceUpdateToPEServer) error {
	return status.Errorf(codes.Unimplemented, "method MEasClientOfPEAndStreamingSendPriceUpdateToPE not implemented")
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_MEasClientOfPEAndStreamingSendPriceUpdateToPE_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(APIServer).MEasClientOfPEAndStreamingSendPriceUpdateToPE(&aPIMEasClientOfPEAndStreamingSendPriceUpdateToPEServer{stream})
}

type API_MEasClientOfPEAndStreamingSendPriceUpdateToPEServer interface {
	SendAndClose(*Acknowledge) error
	Recv() (*Prices, error)
	grpc.ServerStream
}

type aPIMEasClientOfPEAndStreamingSendPriceUpdateToPEServer struct {
	grpc.ServerStream
}

func (x *aPIMEasClientOfPEAndStreamingSendPriceUpdateToPEServer) SendAndClose(m *Acknowledge) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aPIMEasClientOfPEAndStreamingSendPriceUpdateToPEServer) Recv() (*Prices, error) {
	m := new(Prices)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "API",
	HandlerType: (*APIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MEasClientOfPEAndStreamingSendPriceUpdateToPE",
			Handler:       _API_MEasClientOfPEAndStreamingSendPriceUpdateToPE_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "price_api.proto",
}
