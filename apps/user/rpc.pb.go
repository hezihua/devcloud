// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/user/pb/rpc.proto

package user

import (
	request "github.com/infraboard/mcube/http/request"
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

type DESCRIBE_BY int32

const (
	DESCRIBE_BY_USER_ID  DESCRIBE_BY = 0
	DESCRIBE_BY_USERNAME DESCRIBE_BY = 1
)

// Enum value maps for DESCRIBE_BY.
var (
	DESCRIBE_BY_name = map[int32]string{
		0: "USER_ID",
		1: "USERNAME",
	}
	DESCRIBE_BY_value = map[string]int32{
		"USER_ID":  0,
		"USERNAME": 1,
	}
)

func (x DESCRIBE_BY) Enum() *DESCRIBE_BY {
	p := new(DESCRIBE_BY)
	*p = x
	return p
}

func (x DESCRIBE_BY) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DESCRIBE_BY) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_user_pb_rpc_proto_enumTypes[0].Descriptor()
}

func (DESCRIBE_BY) Type() protoreflect.EnumType {
	return &file_apps_user_pb_rpc_proto_enumTypes[0]
}

func (x DESCRIBE_BY) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DESCRIBE_BY.Descriptor instead.
func (DESCRIBE_BY) EnumDescriptor() ([]byte, []int) {
	return file_apps_user_pb_rpc_proto_rawDescGZIP(), []int{0}
}

type QueryUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Keywords string               `protobuf:"bytes,2,opt,name=keywords,proto3" json:"keywords,omitempty"`
}

func (x *QueryUserRequest) Reset() {
	*x = QueryUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserRequest) ProtoMessage() {}

func (x *QueryUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserRequest.ProtoReflect.Descriptor instead.
func (*QueryUserRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryUserRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryUserRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

type DescribeUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DescribeBy    DESCRIBE_BY `protobuf:"varint,1,opt,name=describe_by,json=describeBy,proto3,enum=devcloud.mcenter.user.DESCRIBE_BY" json:"describe_by,omitempty"`
	DescribeValue string      `protobuf:"bytes,2,opt,name=describe_value,json=describeValue,proto3" json:"describe_value,omitempty"`
}

func (x *DescribeUserRequest) Reset() {
	*x = DescribeUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeUserRequest) ProtoMessage() {}

func (x *DescribeUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeUserRequest.ProtoReflect.Descriptor instead.
func (*DescribeUserRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeUserRequest) GetDescribeBy() DESCRIBE_BY {
	if x != nil {
		return x.DescribeBy
	}
	return DESCRIBE_BY_USER_ID
}

func (x *DescribeUserRequest) GetDescribeValue() string {
	if x != nil {
		return x.DescribeValue
	}
	return ""
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_rpc_proto_rawDescGZIP(), []int{2}
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_rpc_proto_rawDescGZIP(), []int{3}
}

var File_apps_user_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_user_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a,
	0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x22, 0x81, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e,
	0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x42,
	0x59, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x12, 0x25, 0x0a,
	0x0e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2a, 0x28,
	0x0a, 0x0b, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x42, 0x59, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x53,
	0x45, 0x52, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x32, 0xb4, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43,
	0x12, 0x54, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x27, 0x2e,
	0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x12, 0x57, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2a, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x20, 0x5a, 0x1e, 0x68, 0x7a, 0x68, 0x2f, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_user_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_user_pb_rpc_proto_rawDescData = file_apps_user_pb_rpc_proto_rawDesc
)

func file_apps_user_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_user_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_user_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_user_pb_rpc_proto_rawDescData)
	})
	return file_apps_user_pb_rpc_proto_rawDescData
}

var file_apps_user_pb_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_user_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_user_pb_rpc_proto_goTypes = []interface{}{
	(DESCRIBE_BY)(0),            // 0: devcloud.mcenter.user.DESCRIBE_BY
	(*QueryUserRequest)(nil),    // 1: devcloud.mcenter.user.QueryUserRequest
	(*DescribeUserRequest)(nil), // 2: devcloud.mcenter.user.DescribeUserRequest
	(*UpdateUserRequest)(nil),   // 3: devcloud.mcenter.user.UpdateUserRequest
	(*DeleteUserRequest)(nil),   // 4: devcloud.mcenter.user.DeleteUserRequest
	(*request.PageRequest)(nil), // 5: infraboard.mcube.page.PageRequest
	(*UserSet)(nil),             // 6: devcloud.mcenter.user.UserSet
	(*User)(nil),                // 7: devcloud.mcenter.user.User
}
var file_apps_user_pb_rpc_proto_depIdxs = []int32{
	5, // 0: devcloud.mcenter.user.QueryUserRequest.page:type_name -> infraboard.mcube.page.PageRequest
	0, // 1: devcloud.mcenter.user.DescribeUserRequest.describe_by:type_name -> devcloud.mcenter.user.DESCRIBE_BY
	1, // 2: devcloud.mcenter.user.RPC.QueryUser:input_type -> devcloud.mcenter.user.QueryUserRequest
	2, // 3: devcloud.mcenter.user.RPC.DescribeUser:input_type -> devcloud.mcenter.user.DescribeUserRequest
	6, // 4: devcloud.mcenter.user.RPC.QueryUser:output_type -> devcloud.mcenter.user.UserSet
	7, // 5: devcloud.mcenter.user.RPC.DescribeUser:output_type -> devcloud.mcenter.user.User
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apps_user_pb_rpc_proto_init() }
func file_apps_user_pb_rpc_proto_init() {
	if File_apps_user_pb_rpc_proto != nil {
		return
	}
	file_apps_user_pb_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_user_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserRequest); i {
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
		file_apps_user_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeUserRequest); i {
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
		file_apps_user_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
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
		file_apps_user_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
			RawDescriptor: file_apps_user_pb_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_user_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_user_pb_rpc_proto_depIdxs,
		EnumInfos:         file_apps_user_pb_rpc_proto_enumTypes,
		MessageInfos:      file_apps_user_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_user_pb_rpc_proto = out.File
	file_apps_user_pb_rpc_proto_rawDesc = nil
	file_apps_user_pb_rpc_proto_goTypes = nil
	file_apps_user_pb_rpc_proto_depIdxs = nil
}