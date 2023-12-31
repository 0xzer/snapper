// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: uuid.proto

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

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncodedId []byte `protobuf:"bytes,1,opt,name=encodedId,proto3" json:"encodedId,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uuid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_uuid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_uuid_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetEncodedId() []byte {
	if x != nil {
		return x.EncodedId
	}
	return nil
}

type UUIDMobile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HighBits int64 `protobuf:"fixed64,1,opt,name=highBits,proto3" json:"highBits,omitempty"`
	LowBits  int64 `protobuf:"fixed64,2,opt,name=lowBits,proto3" json:"lowBits,omitempty"`
}

func (x *UUIDMobile) Reset() {
	*x = UUIDMobile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uuid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUIDMobile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUIDMobile) ProtoMessage() {}

func (x *UUIDMobile) ProtoReflect() protoreflect.Message {
	mi := &file_uuid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUIDMobile.ProtoReflect.Descriptor instead.
func (*UUIDMobile) Descriptor() ([]byte, []int) {
	return file_uuid_proto_rawDescGZIP(), []int{1}
}

func (x *UUIDMobile) GetHighBits() int64 {
	if x != nil {
		return x.HighBits
	}
	return 0
}

func (x *UUIDMobile) GetLowBits() int64 {
	if x != nil {
		return x.LowBits
	}
	return 0
}

var File_uuid_proto protoreflect.FileDescriptor

var file_uuid_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x22, 0x24, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0a, 0x55, 0x55, 0x49, 0x44,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x69, 0x67, 0x68, 0x42, 0x69,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x10, 0x52, 0x08, 0x68, 0x69, 0x67, 0x68, 0x42, 0x69,
	0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x77, 0x42, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x10, 0x52, 0x07, 0x6c, 0x6f, 0x77, 0x42, 0x69, 0x74, 0x73, 0x42, 0x0b, 0x5a, 0x09,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_uuid_proto_rawDescOnce sync.Once
	file_uuid_proto_rawDescData = file_uuid_proto_rawDesc
)

func file_uuid_proto_rawDescGZIP() []byte {
	file_uuid_proto_rawDescOnce.Do(func() {
		file_uuid_proto_rawDescData = protoimpl.X.CompressGZIP(file_uuid_proto_rawDescData)
	})
	return file_uuid_proto_rawDescData
}

var file_uuid_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_uuid_proto_goTypes = []interface{}{
	(*UUID)(nil),       // 0: uuid.UUID
	(*UUIDMobile)(nil), // 1: uuid.UUIDMobile
}
var file_uuid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_uuid_proto_init() }
func file_uuid_proto_init() {
	if File_uuid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_uuid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_uuid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUIDMobile); i {
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
			RawDescriptor: file_uuid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_uuid_proto_goTypes,
		DependencyIndexes: file_uuid_proto_depIdxs,
		MessageInfos:      file_uuid_proto_msgTypes,
	}.Build()
	File_uuid_proto = out.File
	file_uuid_proto_rawDesc = nil
	file_uuid_proto_goTypes = nil
	file_uuid_proto_depIdxs = nil
}
