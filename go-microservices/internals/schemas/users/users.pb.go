// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: protofiles/users.proto

package users

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginDTO) Reset() {
	*x = LoginDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginDTO) ProtoMessage() {}

func (x *LoginDTO) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginDTO.ProtoReflect.Descriptor instead.
func (*LoginDTO) Descriptor() ([]byte, []int) {
	return file_protofiles_users_proto_rawDescGZIP(), []int{0}
}

func (x *LoginDTO) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginDTO) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterDTO) Reset() {
	*x = RegisterDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDTO) ProtoMessage() {}

func (x *RegisterDTO) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterDTO.ProtoReflect.Descriptor instead.
func (*RegisterDTO) Descriptor() ([]byte, []int) {
	return file_protofiles_users_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterDTO) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *RegisterDTO) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *RegisterDTO) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterDTO) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ApiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatCode    int32      `protobuf:"varint,1,opt,name=stat_code,json=statCode,proto3" json:"stat_code,omitempty"`
	StatMessage string     `protobuf:"bytes,2,opt,name=stat_message,json=statMessage,proto3" json:"stat_message,omitempty"`
	Data        *anypb.Any `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *ApiResponse) Reset() {
	*x = ApiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protofiles_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResponse) ProtoMessage() {}

func (x *ApiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protofiles_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResponse.ProtoReflect.Descriptor instead.
func (*ApiResponse) Descriptor() ([]byte, []int) {
	return file_protofiles_users_proto_rawDescGZIP(), []int{2}
}

func (x *ApiResponse) GetStatCode() int32 {
	if x != nil {
		return x.StatCode
	}
	return 0
}

func (x *ApiResponse) GetStatMessage() string {
	if x != nil {
		return x.StatMessage
	}
	return ""
}

func (x *ApiResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_protofiles_users_proto protoreflect.FileDescriptor

var file_protofiles_users_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x08, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x44, 0x54, 0x4f, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x79, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x32, 0x75, 0x0a, 0x05, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x32, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44,
	0x54, 0x4f, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x1a, 0x12, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protofiles_users_proto_rawDescOnce sync.Once
	file_protofiles_users_proto_rawDescData = file_protofiles_users_proto_rawDesc
)

func file_protofiles_users_proto_rawDescGZIP() []byte {
	file_protofiles_users_proto_rawDescOnce.Do(func() {
		file_protofiles_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_protofiles_users_proto_rawDescData)
	})
	return file_protofiles_users_proto_rawDescData
}

var file_protofiles_users_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protofiles_users_proto_goTypes = []interface{}{
	(*LoginDTO)(nil),    // 0: users.LoginDTO
	(*RegisterDTO)(nil), // 1: users.RegisterDTO
	(*ApiResponse)(nil), // 2: users.ApiResponse
	(*anypb.Any)(nil),   // 3: google.protobuf.Any
}
var file_protofiles_users_proto_depIdxs = []int32{
	3, // 0: users.ApiResponse.data:type_name -> google.protobuf.Any
	0, // 1: users.Users.LoginAuth:input_type -> users.LoginDTO
	1, // 2: users.Users.RegisterAuth:input_type -> users.RegisterDTO
	2, // 3: users.Users.LoginAuth:output_type -> users.ApiResponse
	2, // 4: users.Users.RegisterAuth:output_type -> users.ApiResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protofiles_users_proto_init() }
func file_protofiles_users_proto_init() {
	if File_protofiles_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protofiles_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginDTO); i {
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
		file_protofiles_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterDTO); i {
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
		file_protofiles_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiResponse); i {
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
	file_protofiles_users_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protofiles_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protofiles_users_proto_goTypes,
		DependencyIndexes: file_protofiles_users_proto_depIdxs,
		MessageInfos:      file_protofiles_users_proto_msgTypes,
	}.Build()
	File_protofiles_users_proto = out.File
	file_protofiles_users_proto_rawDesc = nil
	file_protofiles_users_proto_goTypes = nil
	file_protofiles_users_proto_depIdxs = nil
}
