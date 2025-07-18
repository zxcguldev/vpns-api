// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.6
// source: integrator/v1/integrator_api.proto

package integrator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProcessPaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderID       string                 `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessPaymentRequest) Reset() {
	*x = ProcessPaymentRequest{}
	mi := &file_integrator_v1_integrator_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessPaymentRequest) ProtoMessage() {}

func (x *ProcessPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrator_v1_integrator_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessPaymentRequest.ProtoReflect.Descriptor instead.
func (*ProcessPaymentRequest) Descriptor() ([]byte, []int) {
	return file_integrator_v1_integrator_api_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessPaymentRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

type ProcessPaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessPaymentResponse) Reset() {
	*x = ProcessPaymentResponse{}
	mi := &file_integrator_v1_integrator_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessPaymentResponse) ProtoMessage() {}

func (x *ProcessPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrator_v1_integrator_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessPaymentResponse.ProtoReflect.Descriptor instead.
func (*ProcessPaymentResponse) Descriptor() ([]byte, []int) {
	return file_integrator_v1_integrator_api_proto_rawDescGZIP(), []int{1}
}

type CheckBanRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientID      string                 `protobuf:"bytes,1,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckBanRequest) Reset() {
	*x = CheckBanRequest{}
	mi := &file_integrator_v1_integrator_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckBanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckBanRequest) ProtoMessage() {}

func (x *CheckBanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integrator_v1_integrator_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckBanRequest.ProtoReflect.Descriptor instead.
func (*CheckBanRequest) Descriptor() ([]byte, []int) {
	return file_integrator_v1_integrator_api_proto_rawDescGZIP(), []int{2}
}

func (x *CheckBanRequest) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

type CheckBanResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsBanned      bool                   `protobuf:"varint,1,opt,name=IsBanned,proto3" json:"IsBanned,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckBanResponse) Reset() {
	*x = CheckBanResponse{}
	mi := &file_integrator_v1_integrator_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckBanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckBanResponse) ProtoMessage() {}

func (x *CheckBanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integrator_v1_integrator_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckBanResponse.ProtoReflect.Descriptor instead.
func (*CheckBanResponse) Descriptor() ([]byte, []int) {
	return file_integrator_v1_integrator_api_proto_rawDescGZIP(), []int{3}
}

func (x *CheckBanResponse) GetIsBanned() bool {
	if x != nil {
		return x.IsBanned
	}
	return false
}

var File_integrator_v1_integrator_api_proto protoreflect.FileDescriptor

const file_integrator_v1_integrator_api_proto_rawDesc = "" +
	"\n" +
	"\"integrator/v1/integrator_api.proto\x12\x0evpns.Client.V1\"1\n" +
	"\x15ProcessPaymentRequest\x12\x18\n" +
	"\aOrderID\x18\x01 \x01(\tR\aOrderID\"\x18\n" +
	"\x16ProcessPaymentResponse\"-\n" +
	"\x0fCheckBanRequest\x12\x1a\n" +
	"\bClientID\x18\x01 \x01(\tR\bClientID\".\n" +
	"\x10CheckBanResponse\x12\x1a\n" +
	"\bIsBanned\x18\x01 \x01(\bR\bIsBanned2\xc7\x01\n" +
	"\x11IntegratorService\x12a\n" +
	"\x0eProcessPayment\x12%.vpns.Client.V1.ProcessPaymentRequest\x1a&.vpns.Client.V1.ProcessPaymentResponse\"\x00\x12O\n" +
	"\bCheckBan\x12\x1f.vpns.Client.V1.CheckBanRequest\x1a .vpns.Client.V1.CheckBanResponse\"\x00B5Z3github.com/zxcguldev/vpns-api/integrator;integratorb\x06proto3"

var (
	file_integrator_v1_integrator_api_proto_rawDescOnce sync.Once
	file_integrator_v1_integrator_api_proto_rawDescData []byte
)

func file_integrator_v1_integrator_api_proto_rawDescGZIP() []byte {
	file_integrator_v1_integrator_api_proto_rawDescOnce.Do(func() {
		file_integrator_v1_integrator_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_integrator_v1_integrator_api_proto_rawDesc), len(file_integrator_v1_integrator_api_proto_rawDesc)))
	})
	return file_integrator_v1_integrator_api_proto_rawDescData
}

var file_integrator_v1_integrator_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_integrator_v1_integrator_api_proto_goTypes = []any{
	(*ProcessPaymentRequest)(nil),  // 0: vpns.Client.V1.ProcessPaymentRequest
	(*ProcessPaymentResponse)(nil), // 1: vpns.Client.V1.ProcessPaymentResponse
	(*CheckBanRequest)(nil),        // 2: vpns.Client.V1.CheckBanRequest
	(*CheckBanResponse)(nil),       // 3: vpns.Client.V1.CheckBanResponse
}
var file_integrator_v1_integrator_api_proto_depIdxs = []int32{
	0, // 0: vpns.Client.V1.IntegratorService.ProcessPayment:input_type -> vpns.Client.V1.ProcessPaymentRequest
	2, // 1: vpns.Client.V1.IntegratorService.CheckBan:input_type -> vpns.Client.V1.CheckBanRequest
	1, // 2: vpns.Client.V1.IntegratorService.ProcessPayment:output_type -> vpns.Client.V1.ProcessPaymentResponse
	3, // 3: vpns.Client.V1.IntegratorService.CheckBan:output_type -> vpns.Client.V1.CheckBanResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_integrator_v1_integrator_api_proto_init() }
func file_integrator_v1_integrator_api_proto_init() {
	if File_integrator_v1_integrator_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_integrator_v1_integrator_api_proto_rawDesc), len(file_integrator_v1_integrator_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_integrator_v1_integrator_api_proto_goTypes,
		DependencyIndexes: file_integrator_v1_integrator_api_proto_depIdxs,
		MessageInfos:      file_integrator_v1_integrator_api_proto_msgTypes,
	}.Build()
	File_integrator_v1_integrator_api_proto = out.File
	file_integrator_v1_integrator_api_proto_goTypes = nil
	file_integrator_v1_integrator_api_proto_depIdxs = nil
}
