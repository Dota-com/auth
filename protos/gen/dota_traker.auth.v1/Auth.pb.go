// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: proto/Auth/Auth.proto

package auth

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

type AuthRegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login     string `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Password2 string `protobuf:"bytes,3,opt,name=Password2,proto3" json:"Password2,omitempty"`
	SteamId   int64  `protobuf:"varint,4,opt,name=SteamId,proto3" json:"SteamId,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *AuthRegistrationRequest) Reset() {
	*x = AuthRegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Auth_Auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRegistrationRequest) ProtoMessage() {}

func (x *AuthRegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Auth_Auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRegistrationRequest.ProtoReflect.Descriptor instead.
func (*AuthRegistrationRequest) Descriptor() ([]byte, []int) {
	return file_proto_Auth_Auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthRegistrationRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *AuthRegistrationRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AuthRegistrationRequest) GetPassword2() string {
	if x != nil {
		return x.Password2
	}
	return ""
}

func (x *AuthRegistrationRequest) GetSteamId() int64 {
	if x != nil {
		return x.SteamId
	}
	return 0
}

func (x *AuthRegistrationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AuthLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *AuthLoginRequest) Reset() {
	*x = AuthLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Auth_Auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLoginRequest) ProtoMessage() {}

func (x *AuthLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Auth_Auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLoginRequest.ProtoReflect.Descriptor instead.
func (*AuthLoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_Auth_Auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthLoginRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *AuthLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthRegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *AuthRegistrationResponse) Reset() {
	*x = AuthRegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Auth_Auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRegistrationResponse) ProtoMessage() {}

func (x *AuthRegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Auth_Auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRegistrationResponse.ProtoReflect.Descriptor instead.
func (*AuthRegistrationResponse) Descriptor() ([]byte, []int) {
	return file_proto_Auth_Auth_proto_rawDescGZIP(), []int{2}
}

func (x *AuthRegistrationResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthRolesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RolesFlag bool `protobuf:"varint,1,opt,name=RolesFlag,proto3" json:"RolesFlag,omitempty"`
}

func (x *AuthRolesResponse) Reset() {
	*x = AuthRolesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Auth_Auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRolesResponse) ProtoMessage() {}

func (x *AuthRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Auth_Auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRolesResponse.ProtoReflect.Descriptor instead.
func (*AuthRolesResponse) Descriptor() ([]byte, []int) {
	return file_proto_Auth_Auth_proto_rawDescGZIP(), []int{3}
}

func (x *AuthRolesResponse) GetRolesFlag() bool {
	if x != nil {
		return x.RolesFlag
	}
	return false
}

type AuthRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *AuthRolesRequest) Reset() {
	*x = AuthRolesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Auth_Auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRolesRequest) ProtoMessage() {}

func (x *AuthRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Auth_Auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRolesRequest.ProtoReflect.Descriptor instead.
func (*AuthRolesRequest) Descriptor() ([]byte, []int) {
	return file_proto_Auth_Auth_proto_rawDescGZIP(), []int{4}
}

func (x *AuthRolesRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_proto_Auth_Auth_proto protoreflect.FileDescriptor

var file_proto_Auth_Auth_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x75, 0x74, 0x68, 0x2f, 0x41, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x99, 0x01,
	0x0a, 0x17, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x53, 0x74, 0x65, 0x61,
	0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x44, 0x0a, 0x10, 0x41, 0x75, 0x74,
	0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x30, 0x0a, 0x18, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x31, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x46,
	0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x73,
	0x46, 0x6c, 0x61, 0x67, 0x22, 0x2a, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x32, 0xe2, 0x01, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x43, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x64, 0x6f, 0x74, 0x61, 0x5f, 0x74, 0x72,
	0x61, 0x6b, 0x65, 0x72, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_Auth_Auth_proto_rawDescOnce sync.Once
	file_proto_Auth_Auth_proto_rawDescData = file_proto_Auth_Auth_proto_rawDesc
)

func file_proto_Auth_Auth_proto_rawDescGZIP() []byte {
	file_proto_Auth_Auth_proto_rawDescOnce.Do(func() {
		file_proto_Auth_Auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_Auth_Auth_proto_rawDescData)
	})
	return file_proto_Auth_Auth_proto_rawDescData
}

var file_proto_Auth_Auth_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_Auth_Auth_proto_goTypes = []interface{}{
	(*AuthRegistrationRequest)(nil),  // 0: auth.AuthRegistrationRequest
	(*AuthLoginRequest)(nil),         // 1: auth.AuthLoginRequest
	(*AuthRegistrationResponse)(nil), // 2: auth.AuthRegistrationResponse
	(*AuthRolesResponse)(nil),        // 3: auth.AuthRolesResponse
	(*AuthRolesRequest)(nil),         // 4: auth.AuthRolesRequest
}
var file_proto_Auth_Auth_proto_depIdxs = []int32{
	1, // 0: auth.AuthServer.AuthLogin:input_type -> auth.AuthLoginRequest
	0, // 1: auth.AuthServer.AuthRegistration:input_type -> auth.AuthRegistrationRequest
	4, // 2: auth.AuthServer.AuthRoles:input_type -> auth.AuthRolesRequest
	2, // 3: auth.AuthServer.AuthLogin:output_type -> auth.AuthRegistrationResponse
	2, // 4: auth.AuthServer.AuthRegistration:output_type -> auth.AuthRegistrationResponse
	3, // 5: auth.AuthServer.AuthRoles:output_type -> auth.AuthRolesResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_Auth_Auth_proto_init() }
func file_proto_Auth_Auth_proto_init() {
	if File_proto_Auth_Auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_Auth_Auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRegistrationRequest); i {
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
		file_proto_Auth_Auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthLoginRequest); i {
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
		file_proto_Auth_Auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRegistrationResponse); i {
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
		file_proto_Auth_Auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRolesResponse); i {
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
		file_proto_Auth_Auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRolesRequest); i {
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
			RawDescriptor: file_proto_Auth_Auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_Auth_Auth_proto_goTypes,
		DependencyIndexes: file_proto_Auth_Auth_proto_depIdxs,
		MessageInfos:      file_proto_Auth_Auth_proto_msgTypes,
	}.Build()
	File_proto_Auth_Auth_proto = out.File
	file_proto_Auth_Auth_proto_rawDesc = nil
	file_proto_Auth_Auth_proto_goTypes = nil
	file_proto_Auth_Auth_proto_depIdxs = nil
}
