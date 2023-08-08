// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: getGroups.proto

package protos

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

type GetGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SelfUserId *UUID `protobuf:"bytes,1,opt,name=selfUserId,proto3" json:"selfUserId,omitempty"`
}

func (x *GetGroupsRequest) Reset() {
	*x = GetGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_getGroups_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupsRequest) ProtoMessage() {}

func (x *GetGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_getGroups_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupsRequest.ProtoReflect.Descriptor instead.
func (*GetGroupsRequest) Descriptor() ([]byte, []int) {
	return file_getGroups_proto_rawDescGZIP(), []int{0}
}

func (x *GetGroupsRequest) GetSelfUserId() *UUID {
	if x != nil {
		return x.SelfUserId
	}
	return nil
}

type GetGroupsResponseWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []*GetGroupsResponse `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *GetGroupsResponseWrapper) Reset() {
	*x = GetGroupsResponseWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_getGroups_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupsResponseWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupsResponseWrapper) ProtoMessage() {}

func (x *GetGroupsResponseWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_getGroups_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupsResponseWrapper.ProtoReflect.Descriptor instead.
func (*GetGroupsResponseWrapper) Descriptor() ([]byte, []int) {
	return file_getGroups_proto_rawDescGZIP(), []int{1}
}

func (x *GetGroupsResponseWrapper) GetResponse() []*GetGroupsResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type GetGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group               *Conversation `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	LastUpdateEventTime int64         `protobuf:"varint,2,opt,name=lastUpdateEventTime,proto3" json:"lastUpdateEventTime,omitempty"`
}

func (x *GetGroupsResponse) Reset() {
	*x = GetGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_getGroups_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupsResponse) ProtoMessage() {}

func (x *GetGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_getGroups_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupsResponse.ProtoReflect.Descriptor instead.
func (*GetGroupsResponse) Descriptor() ([]byte, []int) {
	return file_getGroups_proto_rawDescGZIP(), []int{2}
}

func (x *GetGroupsResponse) GetGroup() *Conversation {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *GetGroupsResponse) GetLastUpdateEventTime() int64 {
	if x != nil {
		return x.LastUpdateEventTime
	}
	return 0
}

var File_getGroups_proto protoreflect.FileDescriptor

var file_getGroups_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x67, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a, 0x0a, 0x75, 0x75,
	0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x0a, 0x73, 0x65, 0x6c, 0x66, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x55, 0x55, 0x49, 0x44,
	0x52, 0x0a, 0x73, 0x65, 0x6c, 0x66, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x77, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x30, 0x0a, 0x13, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_getGroups_proto_rawDescOnce sync.Once
	file_getGroups_proto_rawDescData = file_getGroups_proto_rawDesc
)

func file_getGroups_proto_rawDescGZIP() []byte {
	file_getGroups_proto_rawDescOnce.Do(func() {
		file_getGroups_proto_rawDescData = protoimpl.X.CompressGZIP(file_getGroups_proto_rawDescData)
	})
	return file_getGroups_proto_rawDescData
}

var file_getGroups_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_getGroups_proto_goTypes = []interface{}{
	(*GetGroupsRequest)(nil),         // 0: getGroups.GetGroupsRequest
	(*GetGroupsResponseWrapper)(nil), // 1: getGroups.GetGroupsResponseWrapper
	(*GetGroupsResponse)(nil),        // 2: getGroups.GetGroupsResponse
	(*UUID)(nil),                     // 3: uuid.UUID
	(*Conversation)(nil),             // 4: conversation.Conversation
}
var file_getGroups_proto_depIdxs = []int32{
	3, // 0: getGroups.GetGroupsRequest.selfUserId:type_name -> uuid.UUID
	2, // 1: getGroups.GetGroupsResponseWrapper.response:type_name -> getGroups.GetGroupsResponse
	4, // 2: getGroups.GetGroupsResponse.group:type_name -> conversation.Conversation
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_getGroups_proto_init() }
func file_getGroups_proto_init() {
	if File_getGroups_proto != nil {
		return
	}
	file_uuid_proto_init()
	file_conversation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_getGroups_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupsRequest); i {
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
		file_getGroups_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupsResponseWrapper); i {
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
		file_getGroups_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupsResponse); i {
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
			RawDescriptor: file_getGroups_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_getGroups_proto_goTypes,
		DependencyIndexes: file_getGroups_proto_depIdxs,
		MessageInfos:      file_getGroups_proto_msgTypes,
	}.Build()
	File_getGroups_proto = out.File
	file_getGroups_proto_rawDesc = nil
	file_getGroups_proto_goTypes = nil
	file_getGroups_proto_depIdxs = nil
}
