// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: http.proto

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

type FailureType int32

const (
	FailureType_UNKNOWN                  FailureType = 0
	FailureType_MALFORMED_REQUEST        FailureType = 1
	FailureType_PERSISTENCE              FailureType = 2
	FailureType_NOT_AUTHORIZED           FailureType = 3
	FailureType_DUPLICATE_MESSAGE        FailureType = 4
	FailureType_UPDATE_NOT_APPLICABLE    FailureType = 5
	FailureType_NOT_FRIENDS              FailureType = 6
	FailureType_DUPLICATE_REQUEST        FailureType = 7
	FailureType_DEPENDENCY_NOT_SATISFIED FailureType = 8
	FailureType_OUT_OF_SYNC              FailureType = 9
	FailureType_NOT_FOUND                FailureType = 10
	FailureType_MESSAGE_ALREADY_EXPIRED  FailureType = 11
)

// Enum value maps for FailureType.
var (
	FailureType_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "MALFORMED_REQUEST",
		2:  "PERSISTENCE",
		3:  "NOT_AUTHORIZED",
		4:  "DUPLICATE_MESSAGE",
		5:  "UPDATE_NOT_APPLICABLE",
		6:  "NOT_FRIENDS",
		7:  "DUPLICATE_REQUEST",
		8:  "DEPENDENCY_NOT_SATISFIED",
		9:  "OUT_OF_SYNC",
		10: "NOT_FOUND",
		11: "MESSAGE_ALREADY_EXPIRED",
	}
	FailureType_value = map[string]int32{
		"UNKNOWN":                  0,
		"MALFORMED_REQUEST":        1,
		"PERSISTENCE":              2,
		"NOT_AUTHORIZED":           3,
		"DUPLICATE_MESSAGE":        4,
		"UPDATE_NOT_APPLICABLE":    5,
		"NOT_FRIENDS":              6,
		"DUPLICATE_REQUEST":        7,
		"DEPENDENCY_NOT_SATISFIED": 8,
		"OUT_OF_SYNC":              9,
		"NOT_FOUND":                10,
		"MESSAGE_ALREADY_EXPIRED":  11,
	}
)

func (x FailureType) Enum() *FailureType {
	p := new(FailureType)
	*p = x
	return p
}

func (x FailureType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FailureType) Descriptor() protoreflect.EnumDescriptor {
	return file_http_proto_enumTypes[0].Descriptor()
}

func (FailureType) Type() protoreflect.EnumType {
	return &file_http_proto_enumTypes[0]
}

func (x FailureType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FailureType.Descriptor instead.
func (FailureType) EnumDescriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{0}
}

type FailureReason struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailureType        FailureType        `protobuf:"varint,1,opt,name=failureType,proto3,enum=http.FailureType" json:"failureType,omitempty"`
	FailureDescription string             `protobuf:"bytes,2,opt,name=failureDescription,proto3" json:"failureDescription,omitempty"`
	ServerInfo         *ServerFailureInfo `protobuf:"bytes,3,opt,name=serverInfo,proto3" json:"serverInfo,omitempty"`
	RetryPolicy        *RetryPolicy       `protobuf:"bytes,4,opt,name=retryPolicy,proto3" json:"retryPolicy,omitempty"`
}

func (x *FailureReason) Reset() {
	*x = FailureReason{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailureReason) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailureReason) ProtoMessage() {}

func (x *FailureReason) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailureReason.ProtoReflect.Descriptor instead.
func (*FailureReason) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{0}
}

func (x *FailureReason) GetFailureType() FailureType {
	if x != nil {
		return x.FailureType
	}
	return FailureType_UNKNOWN
}

func (x *FailureReason) GetFailureDescription() string {
	if x != nil {
		return x.FailureDescription
	}
	return ""
}

func (x *FailureReason) GetServerInfo() *ServerFailureInfo {
	if x != nil {
		return x.ServerInfo
	}
	return nil
}

func (x *FailureReason) GetRetryPolicy() *RetryPolicy {
	if x != nil {
		return x.RetryPolicy
	}
	return nil
}

type RetryPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BackoffTimeMs int64 `protobuf:"varint,1,opt,name=backoffTimeMs,proto3" json:"backoffTimeMs,omitempty"`
	ShouldRetry   bool  `protobuf:"varint,2,opt,name=shouldRetry,proto3" json:"shouldRetry,omitempty"`
}

func (x *RetryPolicy) Reset() {
	*x = RetryPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryPolicy) ProtoMessage() {}

func (x *RetryPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryPolicy.ProtoReflect.Descriptor instead.
func (*RetryPolicy) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{1}
}

func (x *RetryPolicy) GetBackoffTimeMs() int64 {
	if x != nil {
		return x.BackoffTimeMs
	}
	return 0
}

func (x *RetryPolicy) GetShouldRetry() bool {
	if x != nil {
		return x.ShouldRetry
	}
	return false
}

type ServerFailureInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RateLimited        bool  `protobuf:"varint,1,opt,name=rateLimited,proto3" json:"rateLimited,omitempty"`
	ErrorRegion        int32 `protobuf:"varint,2,opt,name=errorRegion,proto3" json:"errorRegion,omitempty"`
	IsIntentionalError bool  `protobuf:"varint,3,opt,name=isIntentionalError,proto3" json:"isIntentionalError,omitempty"`
}

func (x *ServerFailureInfo) Reset() {
	*x = ServerFailureInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerFailureInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerFailureInfo) ProtoMessage() {}

func (x *ServerFailureInfo) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerFailureInfo.ProtoReflect.Descriptor instead.
func (*ServerFailureInfo) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{2}
}

func (x *ServerFailureInfo) GetRateLimited() bool {
	if x != nil {
		return x.RateLimited
	}
	return false
}

func (x *ServerFailureInfo) GetErrorRegion() int32 {
	if x != nil {
		return x.ErrorRegion
	}
	return 0
}

func (x *ServerFailureInfo) GetIsIntentionalError() bool {
	if x != nil {
		return x.IsIntentionalError
	}
	return false
}

var File_http_proto protoreflect.FileDescriptor

var file_http_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x74,
	0x74, 0x70, 0x22, 0xe2, 0x01, 0x0a, 0x0d, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x46, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x33, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x52,
	0x65, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x72,
	0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x55, 0x0a, 0x0b, 0x52, 0x65, 0x74, 0x72, 0x79,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66,
	0x66, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x62,
	0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x52, 0x65, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x52, 0x65, 0x74, 0x72, 0x79, 0x22, 0x87,
	0x01, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x73, 0x49, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x8b, 0x02, 0x0a, 0x0b, 0x46, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x41, 0x4c, 0x46, 0x4f, 0x52, 0x4d,
	0x45, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b,
	0x50, 0x45, 0x52, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a,
	0x0e, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10,
	0x03, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x4d,
	0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e,
	0x44, 0x53, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54,
	0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x07, 0x12, 0x1c, 0x0a, 0x18, 0x44,
	0x45, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x41,
	0x54, 0x49, 0x53, 0x46, 0x49, 0x45, 0x44, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x55, 0x54,
	0x5f, 0x4f, 0x46, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x0a, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x50,
	0x49, 0x52, 0x45, 0x44, 0x10, 0x0b, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_http_proto_rawDescOnce sync.Once
	file_http_proto_rawDescData = file_http_proto_rawDesc
)

func file_http_proto_rawDescGZIP() []byte {
	file_http_proto_rawDescOnce.Do(func() {
		file_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_proto_rawDescData)
	})
	return file_http_proto_rawDescData
}

var file_http_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_http_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_http_proto_goTypes = []interface{}{
	(FailureType)(0),          // 0: http.FailureType
	(*FailureReason)(nil),     // 1: http.FailureReason
	(*RetryPolicy)(nil),       // 2: http.RetryPolicy
	(*ServerFailureInfo)(nil), // 3: http.ServerFailureInfo
}
var file_http_proto_depIdxs = []int32{
	0, // 0: http.FailureReason.failureType:type_name -> http.FailureType
	3, // 1: http.FailureReason.serverInfo:type_name -> http.ServerFailureInfo
	2, // 2: http.FailureReason.retryPolicy:type_name -> http.RetryPolicy
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_http_proto_init() }
func file_http_proto_init() {
	if File_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailureReason); i {
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
		file_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryPolicy); i {
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
		file_http_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerFailureInfo); i {
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
			RawDescriptor: file_http_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_http_proto_goTypes,
		DependencyIndexes: file_http_proto_depIdxs,
		EnumInfos:         file_http_proto_enumTypes,
		MessageInfos:      file_http_proto_msgTypes,
	}.Build()
	File_http_proto = out.File
	file_http_proto_rawDesc = nil
	file_http_proto_goTypes = nil
	file_http_proto_depIdxs = nil
}
