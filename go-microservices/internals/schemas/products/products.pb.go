// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: protofiles/products.proto

package products

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServiceName int32

const (
	ServiceName_PingProducts ServiceName = 0
)

// Enum value maps for ServiceName.
var (
	ServiceName_name = map[int32]string{
		0: "PingProducts",
	}
	ServiceName_value = map[string]int32{
		"PingProducts": 0,
	}
)

func (x ServiceName) Enum() *ServiceName {
	p := new(ServiceName)
	*p = x
	return p
}

func (x ServiceName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceName) Descriptor() protoreflect.EnumDescriptor {
	return file_protofiles_products_proto_enumTypes[0].Descriptor()
}

func (ServiceName) Type() protoreflect.EnumType {
	return &file_protofiles_products_proto_enumTypes[0]
}

func (x ServiceName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceName.Descriptor instead.
func (ServiceName) EnumDescriptor() ([]byte, []int) {
	return file_protofiles_products_proto_rawDescGZIP(), []int{0}
}

type GrpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatCode    int32      `protobuf:"varint,1,opt,name=stat_code,json=statCode,proto3" json:"stat_code,omitempty"`
	StatMessage string     `protobuf:"bytes,2,opt,name=stat_message,json=statMessage,proto3" json:"stat_message,omitempty"`
	Data        *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *GrpcResponse) Reset() {
	*x = GrpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_products_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcResponse) ProtoMessage() {}

func (x *GrpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_products_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcResponse.ProtoReflect.Descriptor instead.
func (*GrpcResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_products_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcResponse) GetStatCode() int32 {
	if x != nil {
		return x.StatCode
	}
	return 0
}

func (x *GrpcResponse) GetStatMessage() string {
	if x != nil {
		return x.StatMessage
	}
	return ""
}

func (x *GrpcResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_protofiles_products_proto protoreflect.FileDescriptor

var file_protofiles_products_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01,
	0x0a, 0x0c, 0x47, 0x72, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x74, 0x61, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x1f, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x10, 0x00, 0x32, 0x53, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x50, 0x69,
	0x6e, 0x67, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x47, 0x72,
	0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protofiles_products_proto_rawDescOnce sync.Once
	file_protofiles_products_proto_rawDescData = file_protofiles_products_proto_rawDesc
)

func file_protofiles_products_proto_rawDescGZIP() []byte {
	file_protofiles_products_proto_rawDescOnce.Do(func() {
		file_protofiles_products_proto_rawDescData = protoimpl.X.CompressGZIP(file_protofiles_products_proto_rawDescData)
	})
	return file_protofiles_products_proto_rawDescData
}

var file_protofiles_products_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protofiles_products_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protofiles_products_proto_goTypes = []interface{}{
	(ServiceName)(0),      // 0: products.ServiceName
	(*GrpcResponse)(nil),  // 1: products.GrpcResponse
	(*anypb.Any)(nil),     // 2: google.protobuf.Any
	(*emptypb.Empty)(nil), // 3: google.protobuf.Empty
}
var file_protofiles_products_proto_depIdxs = []int32{
	2, // 0: products.GrpcResponse.data:type_name -> google.protobuf.Any
	3, // 1: products.ProductsService.PingProducts:input_type -> google.protobuf.Empty
	1, // 2: products.ProductsService.PingProducts:output_type -> products.GrpcResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protofiles_products_proto_init() }
func file_protofiles_products_proto_init() {
	if File_protofiles_products_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protofiles_products_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcResponse); i {
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
	file_protofiles_products_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protofiles_products_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protofiles_products_proto_goTypes,
		DependencyIndexes: file_protofiles_products_proto_depIdxs,
		EnumInfos:         file_protofiles_products_proto_enumTypes,
		MessageInfos:      file_protofiles_products_proto_msgTypes,
	}.Build()
	File_protofiles_products_proto = out.File
	file_protofiles_products_proto_rawDesc = nil
	file_protofiles_products_proto_goTypes = nil
	file_protofiles_products_proto_depIdxs = nil
}
