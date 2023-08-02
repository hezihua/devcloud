// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/user/pb/user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	meta "hzh/devcloud/mcenter/common/meta"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户总数
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 用户列表
	// @gotags: json:"items"
	Items []*User `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *UserSet) Reset() {
	*x = UserSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSet) ProtoMessage() {}

func (x *UserSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSet.ProtoReflect.Descriptor instead.
func (*UserSet) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UserSet) GetItems() []*User {
	if x != nil {
		return x.Items
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// @gotags: bson:",inline" json:"spec"
	Spec *CreateUserRequest `protobuf:"bytes,15,opt,name=spec,proto3" json:"spec" bson:",inline"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *User) GetSpec() *CreateUserRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"domain" json: "domain" validate:"required"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty" bson:"domain" validate:"required"`
	// @gotags: bson:"username" json: "username" validate:"required"
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty" bson:"username" validate:"required"`
	// @gotags: bson:"password" json: "password" validate:"required"
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" bson:"password" validate:"required"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_apps_user_pb_user_proto protoreflect.FileDescriptor

var file_apps_user_pb_user_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x64, 0x65, 0x76, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x31, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x75, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x22, 0x63, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x20, 0x5a, 0x1e, 0x68, 0x7a, 0x68, 0x2f,
	0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_apps_user_pb_user_proto_rawDescOnce sync.Once
	file_apps_user_pb_user_proto_rawDescData = file_apps_user_pb_user_proto_rawDesc
)

func file_apps_user_pb_user_proto_rawDescGZIP() []byte {
	file_apps_user_pb_user_proto_rawDescOnce.Do(func() {
		file_apps_user_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_user_pb_user_proto_rawDescData)
	})
	return file_apps_user_pb_user_proto_rawDescData
}

var file_apps_user_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_user_pb_user_proto_goTypes = []interface{}{
	(*UserSet)(nil),           // 0: devcloud.mcenter.user.UserSet
	(*User)(nil),              // 1: devcloud.mcenter.user.User
	(*CreateUserRequest)(nil), // 2: devcloud.mcenter.user.CreateUserRequest
	(*meta.Meta)(nil),         // 3: devcloud.mcenter.meta.Meta
}
var file_apps_user_pb_user_proto_depIdxs = []int32{
	1, // 0: devcloud.mcenter.user.UserSet.items:type_name -> devcloud.mcenter.user.User
	3, // 1: devcloud.mcenter.user.User.meta:type_name -> devcloud.mcenter.meta.Meta
	2, // 2: devcloud.mcenter.user.User.spec:type_name -> devcloud.mcenter.user.CreateUserRequest
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_user_pb_user_proto_init() }
func file_apps_user_pb_user_proto_init() {
	if File_apps_user_pb_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_user_pb_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSet); i {
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
		file_apps_user_pb_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_apps_user_pb_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
			RawDescriptor: file_apps_user_pb_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_user_pb_user_proto_goTypes,
		DependencyIndexes: file_apps_user_pb_user_proto_depIdxs,
		MessageInfos:      file_apps_user_pb_user_proto_msgTypes,
	}.Build()
	File_apps_user_pb_user_proto = out.File
	file_apps_user_pb_user_proto_rawDesc = nil
	file_apps_user_pb_user_proto_goTypes = nil
	file_apps_user_pb_user_proto_depIdxs = nil
}
