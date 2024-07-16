// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: user/rpc/user.proto

package rpc

import (
	entity "github.com/rayfiyo/kotatuneko-rest/gen/user/entity"
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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *entity.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_rpc_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetUser() *entity.User {
	if x != nil {
		return x.User
	}
	return nil
}

type SelectByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // user_name
}

func (x *SelectByNameRequest) Reset() {
	*x = SelectByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_rpc_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectByNameRequest) ProtoMessage() {}

func (x *SelectByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectByNameRequest.ProtoReflect.Descriptor instead.
func (*SelectByNameRequest) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_proto_rawDescGZIP(), []int{1}
}

func (x *SelectByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SelectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *entity.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SelectResponse) Reset() {
	*x = SelectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_rpc_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectResponse) ProtoMessage() {}

func (x *SelectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectResponse.ProtoReflect.Descriptor instead.
func (*SelectResponse) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_proto_rawDescGZIP(), []int{2}
}

func (x *SelectResponse) GetUser() *entity.User {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *entity.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_rpc_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetUser() *entity.User {
	if x != nil {
		return x.User
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // user_id (UUID)
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_rpc_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_rpc_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_user_rpc_user_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_user_rpc_user_proto protoreflect.FileDescriptor

var file_user_rpc_user_proto_rawDesc = []byte{
	0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x70, 0x63, 0x1a,
	0x16, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x29, 0x0a, 0x13, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x0e, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x1f, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x8b, 0x01, 0x0a,
	0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x70, 0x63, 0x42, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x79, 0x66, 0x69, 0x79, 0x6f, 0x2f, 0x6b,
	0x6f, 0x74, 0x61, 0x74, 0x75, 0x6e, 0x65, 0x6b, 0x6f, 0x2d, 0x72, 0x65, 0x73, 0x74, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0xa2, 0x02, 0x03, 0x55, 0x52,
	0x58, 0xaa, 0x02, 0x08, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x70, 0x63, 0xca, 0x02, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x5c, 0x52, 0x70, 0x63, 0xe2, 0x02, 0x14, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x52,
	0x70, 0x63, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x09, 0x55, 0x73, 0x65, 0x72, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_user_rpc_user_proto_rawDescOnce sync.Once
	file_user_rpc_user_proto_rawDescData = file_user_rpc_user_proto_rawDesc
)

func file_user_rpc_user_proto_rawDescGZIP() []byte {
	file_user_rpc_user_proto_rawDescOnce.Do(func() {
		file_user_rpc_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_rpc_user_proto_rawDescData)
	})
	return file_user_rpc_user_proto_rawDescData
}

var file_user_rpc_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_rpc_user_proto_goTypes = []any{
	(*CreateRequest)(nil),       // 0: user.rpc.CreateRequest
	(*SelectByNameRequest)(nil), // 1: user.rpc.SelectByNameRequest
	(*SelectResponse)(nil),      // 2: user.rpc.SelectResponse
	(*UpdateRequest)(nil),       // 3: user.rpc.UpdateRequest
	(*DeleteRequest)(nil),       // 4: user.rpc.DeleteRequest
	(*entity.User)(nil),         // 5: user.entity.User
}
var file_user_rpc_user_proto_depIdxs = []int32{
	5, // 0: user.rpc.CreateRequest.user:type_name -> user.entity.User
	5, // 1: user.rpc.SelectResponse.user:type_name -> user.entity.User
	5, // 2: user.rpc.UpdateRequest.user:type_name -> user.entity.User
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_user_rpc_user_proto_init() }
func file_user_rpc_user_proto_init() {
	if File_user_rpc_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_rpc_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRequest); i {
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
		file_user_rpc_user_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SelectByNameRequest); i {
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
		file_user_rpc_user_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SelectResponse); i {
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
		file_user_rpc_user_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRequest); i {
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
		file_user_rpc_user_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRequest); i {
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
			RawDescriptor: file_user_rpc_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_rpc_user_proto_goTypes,
		DependencyIndexes: file_user_rpc_user_proto_depIdxs,
		MessageInfos:      file_user_rpc_user_proto_msgTypes,
	}.Build()
	File_user_rpc_user_proto = out.File
	file_user_rpc_user_proto_rawDesc = nil
	file_user_rpc_user_proto_goTypes = nil
	file_user_rpc_user_proto_depIdxs = nil
}
