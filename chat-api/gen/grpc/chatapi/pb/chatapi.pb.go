// Code generated with goa v3.3.1, DO NOT EDIT.
//
// chatapi protocol buffer definition
//
// Command:
// $ goa gen chat-api/design

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.12.1
// source: chatapi.proto

package chatapipb

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

type GetchatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// room id
	Id int32 `protobuf:"zigzag32,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetchatRequest) Reset() {
	*x = GetchatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetchatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetchatRequest) ProtoMessage() {}

func (x *GetchatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chatapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetchatRequest.ProtoReflect.Descriptor instead.
func (*GetchatRequest) Descriptor() ([]byte, []int) {
	return file_chatapi_proto_rawDescGZIP(), []int{0}
}

func (x *GetchatRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GoaChatCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field []*GoaChat `protobuf:"bytes,1,rep,name=field,proto3" json:"field,omitempty"`
}

func (x *GoaChatCollection) Reset() {
	*x = GoaChatCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoaChatCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoaChatCollection) ProtoMessage() {}

func (x *GoaChatCollection) ProtoReflect() protoreflect.Message {
	mi := &file_chatapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoaChatCollection.ProtoReflect.Descriptor instead.
func (*GoaChatCollection) Descriptor() ([]byte, []int) {
	return file_chatapi_proto_rawDescGZIP(), []int{1}
}

func (x *GoaChatCollection) GetField() []*GoaChat {
	if x != nil {
		return x.Field
	}
	return nil
}

// All chat
type GoaChat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// room id
	Id int32 `protobuf:"zigzag32,1,opt,name=id,proto3" json:"id,omitempty"`
	// user id
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// room name
	RoomName string `protobuf:"bytes,3,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`
	// member
	Member string `protobuf:"bytes,4,opt,name=member,proto3" json:"member,omitempty"`
	// chat
	Chat   string `protobuf:"bytes,5,opt,name=chat,proto3" json:"chat,omitempty"`
	PostDt string `protobuf:"bytes,6,opt,name=post_dt,json=postDt,proto3" json:"post_dt,omitempty"`
}

func (x *GoaChat) Reset() {
	*x = GoaChat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoaChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoaChat) ProtoMessage() {}

func (x *GoaChat) ProtoReflect() protoreflect.Message {
	mi := &file_chatapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoaChat.ProtoReflect.Descriptor instead.
func (*GoaChat) Descriptor() ([]byte, []int) {
	return file_chatapi_proto_rawDescGZIP(), []int{2}
}

func (x *GoaChat) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GoaChat) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GoaChat) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *GoaChat) GetMember() string {
	if x != nil {
		return x.Member
	}
	return ""
}

func (x *GoaChat) GetChat() string {
	if x != nil {
		return x.Chat
	}
	return ""
}

func (x *GoaChat) GetPostDt() string {
	if x != nil {
		return x.PostDt
	}
	return ""
}

var File_chatapi_proto protoreflect.FileDescriptor

var file_chatapi_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x68, 0x61, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x68, 0x61, 0x74, 0x61, 0x70, 0x69, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x63,
	0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x47, 0x6f,
	0x61, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x26, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x6f, 0x61, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x07, 0x47, 0x6f, 0x61, 0x43,
	0x68, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x68, 0x61, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x64, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x44, 0x74, 0x32, 0x49,
	0x0a, 0x07, 0x43, 0x68, 0x61, 0x74, 0x61, 0x70, 0x69, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x63, 0x68, 0x61, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x63, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x6f, 0x61, 0x43, 0x68, 0x61, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x61, 0x70, 0x69, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chatapi_proto_rawDescOnce sync.Once
	file_chatapi_proto_rawDescData = file_chatapi_proto_rawDesc
)

func file_chatapi_proto_rawDescGZIP() []byte {
	file_chatapi_proto_rawDescOnce.Do(func() {
		file_chatapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatapi_proto_rawDescData)
	})
	return file_chatapi_proto_rawDescData
}

var file_chatapi_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chatapi_proto_goTypes = []interface{}{
	(*GetchatRequest)(nil),    // 0: chatapi.GetchatRequest
	(*GoaChatCollection)(nil), // 1: chatapi.GoaChatCollection
	(*GoaChat)(nil),           // 2: chatapi.GoaChat
}
var file_chatapi_proto_depIdxs = []int32{
	2, // 0: chatapi.GoaChatCollection.field:type_name -> chatapi.GoaChat
	0, // 1: chatapi.Chatapi.Getchat:input_type -> chatapi.GetchatRequest
	1, // 2: chatapi.Chatapi.Getchat:output_type -> chatapi.GoaChatCollection
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chatapi_proto_init() }
func file_chatapi_proto_init() {
	if File_chatapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetchatRequest); i {
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
		file_chatapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoaChatCollection); i {
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
		file_chatapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoaChat); i {
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
			RawDescriptor: file_chatapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chatapi_proto_goTypes,
		DependencyIndexes: file_chatapi_proto_depIdxs,
		MessageInfos:      file_chatapi_proto_msgTypes,
	}.Build()
	File_chatapi_proto = out.File
	file_chatapi_proto_rawDesc = nil
	file_chatapi_proto_goTypes = nil
	file_chatapi_proto_depIdxs = nil
}
