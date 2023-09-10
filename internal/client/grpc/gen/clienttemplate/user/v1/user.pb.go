// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: clienttemplate/user/v1/user.proto

package userv1

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

// Пользователь.
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID пользователя
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Фамилия
	LastName string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// Имя
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// Отчество
	SecondName string `protobuf:"bytes,4,opt,name=second_name,json=secondName,proto3" json:"second_name,omitempty"`
	// Возраст
	Age int32 `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
	// Пароль
	Password string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	// E-mail
	Email string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	// Телефон
	Phone string `protobuf:"bytes,8,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clienttemplate_user_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_clienttemplate_user_v1_user_proto_msgTypes[0]
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
	return file_clienttemplate_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetSecondName() string {
	if x != nil {
		return x.SecondName
	}
	return ""
}

func (x *User) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_clienttemplate_user_v1_user_proto protoreflect.FileDescriptor

var file_clienttemplate_user_v1_user_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0xcd, 0x01, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x80, 0x01, 0x0a, 0x1a,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x55, 0x58, 0xaa, 0x02, 0x16, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clienttemplate_user_v1_user_proto_rawDescOnce sync.Once
	file_clienttemplate_user_v1_user_proto_rawDescData = file_clienttemplate_user_v1_user_proto_rawDesc
)

func file_clienttemplate_user_v1_user_proto_rawDescGZIP() []byte {
	file_clienttemplate_user_v1_user_proto_rawDescOnce.Do(func() {
		file_clienttemplate_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_clienttemplate_user_v1_user_proto_rawDescData)
	})
	return file_clienttemplate_user_v1_user_proto_rawDescData
}

var file_clienttemplate_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_clienttemplate_user_v1_user_proto_goTypes = []interface{}{
	(*User)(nil), // 0: clienttemplate.user.v1.User
}
var file_clienttemplate_user_v1_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_clienttemplate_user_v1_user_proto_init() }
func file_clienttemplate_user_v1_user_proto_init() {
	if File_clienttemplate_user_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clienttemplate_user_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_clienttemplate_user_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_clienttemplate_user_v1_user_proto_goTypes,
		DependencyIndexes: file_clienttemplate_user_v1_user_proto_depIdxs,
		MessageInfos:      file_clienttemplate_user_v1_user_proto_msgTypes,
	}.Build()
	File_clienttemplate_user_v1_user_proto = out.File
	file_clienttemplate_user_v1_user_proto_rawDesc = nil
	file_clienttemplate_user_v1_user_proto_goTypes = nil
	file_clienttemplate_user_v1_user_proto_depIdxs = nil
}
