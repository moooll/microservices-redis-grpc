// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protocol/price.proto

package protocol

import (
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

type PriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PriceRequest) Reset() {
	*x = PriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_price_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceRequest) ProtoMessage() {}

func (x *PriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_price_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceRequest.ProtoReflect.Descriptor instead.
func (*PriceRequest) Descriptor() ([]byte, []int) {
	return file_protocol_price_proto_rawDescGZIP(), []int{0}
}

type PriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	CompanyName string  `protobuf:"bytes,3,opt,name=companyName,proto3" json:"companyName,omitempty"`
	BuyPrice    float32 `protobuf:"fixed32,4,opt,name=buyPrice,proto3" json:"buyPrice,omitempty"`
	SellPrice   float32 `protobuf:"fixed32,5,opt,name=sellPrice,proto3" json:"sellPrice,omitempty"`
}

func (x *PriceResponse) Reset() {
	*x = PriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_price_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriceResponse) ProtoMessage() {}

func (x *PriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_price_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriceResponse.ProtoReflect.Descriptor instead.
func (*PriceResponse) Descriptor() ([]byte, []int) {
	return file_protocol_price_proto_rawDescGZIP(), []int{1}
}

func (x *PriceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PriceResponse) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *PriceResponse) GetBuyPrice() float32 {
	if x != nil {
		return x.BuyPrice
	}
	return 0
}

func (x *PriceResponse) GetSellPrice() float32 {
	if x != nil {
		return x.SellPrice
	}
	return 0
}

var File_protocol_price_proto protoreflect.FileDescriptor

var file_protocol_price_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x14, 0x0a,
	0x0c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4a, 0x04, 0x08,
	0x01, 0x10, 0x02, 0x22, 0x7b, 0x0a, 0x0d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x75, 0x79, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x62, 0x75, 0x79, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x73, 0x65, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x32, 0x4c, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3c, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x13, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x43,
	0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6f,
	0x6f, 0x6c, 0x6c, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2d, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_price_proto_rawDescOnce sync.Once
	file_protocol_price_proto_rawDescData = file_protocol_price_proto_rawDesc
)

func file_protocol_price_proto_rawDescGZIP() []byte {
	file_protocol_price_proto_rawDescOnce.Do(func() {
		file_protocol_price_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_price_proto_rawDescData)
	})
	return file_protocol_price_proto_rawDescData
}

var file_protocol_price_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protocol_price_proto_goTypes = []interface{}{
	(*PriceRequest)(nil),  // 0: price.PriceRequest
	(*PriceResponse)(nil), // 1: price.PriceResponse
}
var file_protocol_price_proto_depIdxs = []int32{
	0, // 0: price.PriceService.StreamPrice:input_type -> price.PriceRequest
	1, // 1: price.PriceService.StreamPrice:output_type -> price.PriceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protocol_price_proto_init() }
func file_protocol_price_proto_init() {
	if File_protocol_price_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_price_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceRequest); i {
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
		file_protocol_price_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriceResponse); i {
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
			RawDescriptor: file_protocol_price_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_price_proto_goTypes,
		DependencyIndexes: file_protocol_price_proto_depIdxs,
		MessageInfos:      file_protocol_price_proto_msgTypes,
	}.Build()
	File_protocol_price_proto = out.File
	file_protocol_price_proto_rawDesc = nil
	file_protocol_price_proto_goTypes = nil
	file_protocol_price_proto_depIdxs = nil
}