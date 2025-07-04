// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.6
// source: client/v1/client_api.proto

package client

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type PaymentType int32

const (
	PaymentType_PaymentTypeUnspecified PaymentType = 0
	PaymentType_PaymentTypeCrypto      PaymentType = 1
	PaymentType_PaymentTypeAifory      PaymentType = 2
)

// Enum value maps for PaymentType.
var (
	PaymentType_name = map[int32]string{
		0: "PaymentTypeUnspecified",
		1: "PaymentTypeCrypto",
		2: "PaymentTypeAifory",
	}
	PaymentType_value = map[string]int32{
		"PaymentTypeUnspecified": 0,
		"PaymentTypeCrypto":      1,
		"PaymentTypeAifory":      2,
	}
)

func (x PaymentType) Enum() *PaymentType {
	p := new(PaymentType)
	*p = x
	return p
}

func (x PaymentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentType) Descriptor() protoreflect.EnumDescriptor {
	return file_client_v1_client_api_proto_enumTypes[0].Descriptor()
}

func (PaymentType) Type() protoreflect.EnumType {
	return &file_client_v1_client_api_proto_enumTypes[0]
}

func (x PaymentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentType.Descriptor instead.
func (PaymentType) EnumDescriptor() ([]byte, []int) {
	return file_client_v1_client_api_proto_rawDescGZIP(), []int{0}
}

type OrderType int32

const (
	OrderType_OrderTypeUnspecified OrderType = 0
	OrderType_OrderTypeCreate      OrderType = 1
	OrderType_OrderTypeRefresh     OrderType = 2
)

// Enum value maps for OrderType.
var (
	OrderType_name = map[int32]string{
		0: "OrderTypeUnspecified",
		1: "OrderTypeCreate",
		2: "OrderTypeRefresh",
	}
	OrderType_value = map[string]int32{
		"OrderTypeUnspecified": 0,
		"OrderTypeCreate":      1,
		"OrderTypeRefresh":     2,
	}
)

func (x OrderType) Enum() *OrderType {
	p := new(OrderType)
	*p = x
	return p
}

func (x OrderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderType) Descriptor() protoreflect.EnumDescriptor {
	return file_client_v1_client_api_proto_enumTypes[1].Descriptor()
}

func (OrderType) Type() protoreflect.EnumType {
	return &file_client_v1_client_api_proto_enumTypes[1]
}

func (x OrderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderType.Descriptor instead.
func (OrderType) EnumDescriptor() ([]byte, []int) {
	return file_client_v1_client_api_proto_rawDescGZIP(), []int{1}
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PaymentType   PaymentType            `protobuf:"varint,1,opt,name=PaymentType,proto3,enum=Vpns.Client.V1.PaymentType" json:"PaymentType,omitempty"`
	OrderType     OrderType              `protobuf:"varint,2,opt,name=OrderType,proto3,enum=Vpns.Client.V1.OrderType" json:"OrderType,omitempty"`
	TariffID      int64                  `protobuf:"varint,3,opt,name=TariffID,proto3" json:"TariffID,omitempty"`
	ClientID      string                 `protobuf:"bytes,4,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_client_v1_client_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_v1_client_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_client_v1_client_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetPaymentType() PaymentType {
	if x != nil {
		return x.PaymentType
	}
	return PaymentType_PaymentTypeUnspecified
}

func (x *CreateOrderRequest) GetOrderType() OrderType {
	if x != nil {
		return x.OrderType
	}
	return OrderType_OrderTypeUnspecified
}

func (x *CreateOrderRequest) GetTariffID() int64 {
	if x != nil {
		return x.TariffID
	}
	return 0
}

func (x *CreateOrderRequest) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderID       string                 `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_client_v1_client_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_v1_client_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_client_v1_client_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderResponse) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	OrderType     OrderType              `protobuf:"varint,2,opt,name=OrderType,proto3,enum=Vpns.Client.V1.OrderType" json:"OrderType,omitempty"`
	OrderStatus   string                 `protobuf:"bytes,3,opt,name=OrderStatus,proto3" json:"OrderStatus,omitempty"`
	PaymentType   PaymentType            `protobuf:"varint,4,opt,name=PaymentType,proto3,enum=Vpns.Client.V1.PaymentType" json:"PaymentType,omitempty"`
	ClientID      string                 `protobuf:"bytes,5,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	Tariff        *Tariff                `protobuf:"bytes,6,opt,name=Tariff,proto3" json:"Tariff,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_client_v1_client_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_client_v1_client_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_client_v1_client_api_proto_rawDescGZIP(), []int{2}
}

func (x *Order) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Order) GetOrderType() OrderType {
	if x != nil {
		return x.OrderType
	}
	return OrderType_OrderTypeUnspecified
}

func (x *Order) GetOrderStatus() string {
	if x != nil {
		return x.OrderStatus
	}
	return ""
}

func (x *Order) GetPaymentType() PaymentType {
	if x != nil {
		return x.PaymentType
	}
	return PaymentType_PaymentTypeUnspecified
}

func (x *Order) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *Order) GetTariff() *Tariff {
	if x != nil {
		return x.Tariff
	}
	return nil
}

func (x *Order) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Order) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Tariff struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TariffID      int64                  `protobuf:"varint,1,opt,name=TariffID,proto3" json:"TariffID,omitempty"` // Add other Tariff fields as needed
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tariff) Reset() {
	*x = Tariff{}
	mi := &file_client_v1_client_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tariff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tariff) ProtoMessage() {}

func (x *Tariff) ProtoReflect() protoreflect.Message {
	mi := &file_client_v1_client_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tariff.ProtoReflect.Descriptor instead.
func (*Tariff) Descriptor() ([]byte, []int) {
	return file_client_v1_client_api_proto_rawDescGZIP(), []int{3}
}

func (x *Tariff) GetTariffID() int64 {
	if x != nil {
		return x.TariffID
	}
	return 0
}

var File_client_v1_client_api_proto protoreflect.FileDescriptor

const file_client_v1_client_api_proto_rawDesc = "" +
	"\n" +
	"\x1aclient/v1/client_api.proto\x12\x0eVpns.Client.V1\x1a\x1fgoogle/protobuf/timestamp.proto\"\xc4\x01\n" +
	"\x12CreateOrderRequest\x12=\n" +
	"\vPaymentType\x18\x01 \x01(\x0e2\x1b.Vpns.Client.V1.PaymentTypeR\vPaymentType\x127\n" +
	"\tOrderType\x18\x02 \x01(\x0e2\x19.Vpns.Client.V1.OrderTypeR\tOrderType\x12\x1a\n" +
	"\bTariffID\x18\x03 \x01(\x03R\bTariffID\x12\x1a\n" +
	"\bClientID\x18\x04 \x01(\tR\bClientID\"/\n" +
	"\x13CreateOrderResponse\x12\x18\n" +
	"\aOrderID\x18\x01 \x01(\tR\aOrderID\"\xf1\x02\n" +
	"\x05Order\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\tR\x02ID\x127\n" +
	"\tOrderType\x18\x02 \x01(\x0e2\x19.Vpns.Client.V1.OrderTypeR\tOrderType\x12 \n" +
	"\vOrderStatus\x18\x03 \x01(\tR\vOrderStatus\x12=\n" +
	"\vPaymentType\x18\x04 \x01(\x0e2\x1b.Vpns.Client.V1.PaymentTypeR\vPaymentType\x12\x1a\n" +
	"\bClientID\x18\x05 \x01(\tR\bClientID\x12.\n" +
	"\x06Tariff\x18\x06 \x01(\v2\x16.Vpns.Client.V1.TariffR\x06Tariff\x128\n" +
	"\tCreatedAt\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tCreatedAt\x128\n" +
	"\tUpdatedAt\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tUpdatedAt\"$\n" +
	"\x06Tariff\x12\x1a\n" +
	"\bTariffID\x18\x01 \x01(\x03R\bTariffID*W\n" +
	"\vPaymentType\x12\x1a\n" +
	"\x16PaymentTypeUnspecified\x10\x00\x12\x15\n" +
	"\x11PaymentTypeCrypto\x10\x01\x12\x15\n" +
	"\x11PaymentTypeAifory\x10\x02*P\n" +
	"\tOrderType\x12\x18\n" +
	"\x14OrderTypeUnspecified\x10\x00\x12\x13\n" +
	"\x0fOrderTypeCreate\x10\x01\x12\x14\n" +
	"\x10OrderTypeRefresh\x10\x022i\n" +
	"\rClientService\x12X\n" +
	"\vCreateOrder\x12\".Vpns.Client.V1.CreateOrderRequest\x1a#.Vpns.Client.V1.CreateOrderResponse\"\x00B-Z+github.com/zxcguldev/vpns-api/client;clientb\x06proto3"

var (
	file_client_v1_client_api_proto_rawDescOnce sync.Once
	file_client_v1_client_api_proto_rawDescData []byte
)

func file_client_v1_client_api_proto_rawDescGZIP() []byte {
	file_client_v1_client_api_proto_rawDescOnce.Do(func() {
		file_client_v1_client_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_client_v1_client_api_proto_rawDesc), len(file_client_v1_client_api_proto_rawDesc)))
	})
	return file_client_v1_client_api_proto_rawDescData
}

var file_client_v1_client_api_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_client_v1_client_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_client_v1_client_api_proto_goTypes = []any{
	(PaymentType)(0),              // 0: Vpns.Client.V1.PaymentType
	(OrderType)(0),                // 1: Vpns.Client.V1.OrderType
	(*CreateOrderRequest)(nil),    // 2: Vpns.Client.V1.CreateOrderRequest
	(*CreateOrderResponse)(nil),   // 3: Vpns.Client.V1.CreateOrderResponse
	(*Order)(nil),                 // 4: Vpns.Client.V1.Order
	(*Tariff)(nil),                // 5: Vpns.Client.V1.Tariff
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_client_v1_client_api_proto_depIdxs = []int32{
	0, // 0: Vpns.Client.V1.CreateOrderRequest.PaymentType:type_name -> Vpns.Client.V1.PaymentType
	1, // 1: Vpns.Client.V1.CreateOrderRequest.OrderType:type_name -> Vpns.Client.V1.OrderType
	1, // 2: Vpns.Client.V1.Order.OrderType:type_name -> Vpns.Client.V1.OrderType
	0, // 3: Vpns.Client.V1.Order.PaymentType:type_name -> Vpns.Client.V1.PaymentType
	5, // 4: Vpns.Client.V1.Order.Tariff:type_name -> Vpns.Client.V1.Tariff
	6, // 5: Vpns.Client.V1.Order.CreatedAt:type_name -> google.protobuf.Timestamp
	6, // 6: Vpns.Client.V1.Order.UpdatedAt:type_name -> google.protobuf.Timestamp
	2, // 7: Vpns.Client.V1.ClientService.CreateOrder:input_type -> Vpns.Client.V1.CreateOrderRequest
	3, // 8: Vpns.Client.V1.ClientService.CreateOrder:output_type -> Vpns.Client.V1.CreateOrderResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_client_v1_client_api_proto_init() }
func file_client_v1_client_api_proto_init() {
	if File_client_v1_client_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_client_v1_client_api_proto_rawDesc), len(file_client_v1_client_api_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_v1_client_api_proto_goTypes,
		DependencyIndexes: file_client_v1_client_api_proto_depIdxs,
		EnumInfos:         file_client_v1_client_api_proto_enumTypes,
		MessageInfos:      file_client_v1_client_api_proto_msgTypes,
	}.Build()
	File_client_v1_client_api_proto = out.File
	file_client_v1_client_api_proto_goTypes = nil
	file_client_v1_client_api_proto_depIdxs = nil
}
