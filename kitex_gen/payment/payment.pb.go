// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.0
// source: idl/payment.proto

package rpcPayment

import (
	context "context"
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

type CreatePaymentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oid    string  `protobuf:"bytes,1,opt,name=oid,proto3" json:"oid,omitempty"`
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	UserId string  `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreatePaymentReq) Reset() {
	*x = CreatePaymentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaymentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentReq) ProtoMessage() {}

func (x *CreatePaymentReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentReq.ProtoReflect.Descriptor instead.
func (*CreatePaymentReq) Descriptor() ([]byte, []int) {
	return file_idl_payment_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePaymentReq) GetOid() string {
	if x != nil {
		return x.Oid
	}
	return ""
}

func (x *CreatePaymentReq) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreatePaymentReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreatePaymentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	PaymentUrl string `protobuf:"bytes,2,opt,name=payment_url,json=paymentUrl,proto3" json:"payment_url,omitempty"`
}

func (x *CreatePaymentResp) Reset() {
	*x = CreatePaymentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaymentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentResp) ProtoMessage() {}

func (x *CreatePaymentResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentResp.ProtoReflect.Descriptor instead.
func (*CreatePaymentResp) Descriptor() ([]byte, []int) {
	return file_idl_payment_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePaymentResp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CreatePaymentResp) GetPaymentUrl() string {
	if x != nil {
		return x.PaymentUrl
	}
	return ""
}

type PayCallbackReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawData []byte `protobuf:"bytes,1,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
}

func (x *PayCallbackReq) Reset() {
	*x = PayCallbackReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayCallbackReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayCallbackReq) ProtoMessage() {}

func (x *PayCallbackReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayCallbackReq.ProtoReflect.Descriptor instead.
func (*PayCallbackReq) Descriptor() ([]byte, []int) {
	return file_idl_payment_proto_rawDescGZIP(), []int{2}
}

func (x *PayCallbackReq) GetRawData() []byte {
	if x != nil {
		return x.RawData
	}
	return nil
}

type PayCallbackResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
}

func (x *PayCallbackResp) Reset() {
	*x = PayCallbackResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayCallbackResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayCallbackResp) ProtoMessage() {}

func (x *PayCallbackResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayCallbackResp.ProtoReflect.Descriptor instead.
func (*PayCallbackResp) Descriptor() ([]byte, []int) {
	return file_idl_payment_proto_rawDescGZIP(), []int{3}
}

func (x *PayCallbackResp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

type PayNotifyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawData []byte `protobuf:"bytes,1,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
}

func (x *PayNotifyReq) Reset() {
	*x = PayNotifyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayNotifyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayNotifyReq) ProtoMessage() {}

func (x *PayNotifyReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayNotifyReq.ProtoReflect.Descriptor instead.
func (*PayNotifyReq) Descriptor() ([]byte, []int) {
	return file_idl_payment_proto_rawDescGZIP(), []int{4}
}

func (x *PayNotifyReq) GetRawData() []byte {
	if x != nil {
		return x.RawData
	}
	return nil
}

type PayNotifyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64 `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
}

func (x *PayNotifyResp) Reset() {
	*x = PayNotifyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_payment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayNotifyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayNotifyResp) ProtoMessage() {}

func (x *PayNotifyResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_payment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayNotifyResp.ProtoReflect.Descriptor instead.
func (*PayNotifyResp) Descriptor() ([]byte, []int) {
	return file_idl_payment_proto_rawDescGZIP(), []int{5}
}

func (x *PayNotifyResp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

var File_idl_payment_proto protoreflect.FileDescriptor

var file_idl_payment_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x55, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x2b, 0x0a, 0x0e, 0x50, 0x61,
	0x79, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08,
	0x72, 0x61, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x72, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x22, 0x32, 0x0a, 0x0f, 0x50, 0x61, 0x79, 0x43, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x29, 0x0a, 0x0c, 0x50,
	0x61, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x72,
	0x61, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72,
	0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x22, 0x30, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xdc, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50,
	0x61, 0x79, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x43, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x50, 0x61, 0x79,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x15, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x50, 0x61, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x6f, 0x6d, 0x61, 0x6c,
	0x6c, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x3b, 0x72, 0x70, 0x63, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_payment_proto_rawDescOnce sync.Once
	file_idl_payment_proto_rawDescData = file_idl_payment_proto_rawDesc
)

func file_idl_payment_proto_rawDescGZIP() []byte {
	file_idl_payment_proto_rawDescOnce.Do(func() {
		file_idl_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_payment_proto_rawDescData)
	})
	return file_idl_payment_proto_rawDescData
}

var file_idl_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_idl_payment_proto_goTypes = []interface{}{
	(*CreatePaymentReq)(nil),  // 0: payment.CreatePaymentReq
	(*CreatePaymentResp)(nil), // 1: payment.CreatePaymentResp
	(*PayCallbackReq)(nil),    // 2: payment.PayCallbackReq
	(*PayCallbackResp)(nil),   // 3: payment.PayCallbackResp
	(*PayNotifyReq)(nil),      // 4: payment.PayNotifyReq
	(*PayNotifyResp)(nil),     // 5: payment.PayNotifyResp
}
var file_idl_payment_proto_depIdxs = []int32{
	0, // 0: payment.PaymentService.CreatePayment:input_type -> payment.CreatePaymentReq
	2, // 1: payment.PaymentService.PayCallback:input_type -> payment.PayCallbackReq
	4, // 2: payment.PaymentService.PayNotify:input_type -> payment.PayNotifyReq
	1, // 3: payment.PaymentService.CreatePayment:output_type -> payment.CreatePaymentResp
	3, // 4: payment.PaymentService.PayCallback:output_type -> payment.PayCallbackResp
	5, // 5: payment.PaymentService.PayNotify:output_type -> payment.PayNotifyResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_idl_payment_proto_init() }
func file_idl_payment_proto_init() {
	if File_idl_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePaymentReq); i {
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
		file_idl_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePaymentResp); i {
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
		file_idl_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayCallbackReq); i {
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
		file_idl_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayCallbackResp); i {
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
		file_idl_payment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayNotifyReq); i {
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
		file_idl_payment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayNotifyResp); i {
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
			RawDescriptor: file_idl_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_payment_proto_goTypes,
		DependencyIndexes: file_idl_payment_proto_depIdxs,
		MessageInfos:      file_idl_payment_proto_msgTypes,
	}.Build()
	File_idl_payment_proto = out.File
	file_idl_payment_proto_rawDesc = nil
	file_idl_payment_proto_goTypes = nil
	file_idl_payment_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.11.3. DO NOT EDIT.

type PaymentService interface {
	CreatePayment(ctx context.Context, req *CreatePaymentReq) (res *CreatePaymentResp, err error)
	PayCallback(ctx context.Context, req *PayCallbackReq) (res *PayCallbackResp, err error)
	PayNotify(ctx context.Context, req *PayNotifyReq) (res *PayNotifyResp, err error)
}
