// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: unlockables.proto

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

type SnapSource int32

const (
	SnapSource_SNAP_SOURCE_UNKNOWN SnapSource = 0
	SnapSource_SNAP_SOURCE_CAMERA  SnapSource = 1
	SnapSource_SNAP_SOURCE_MEMORY  SnapSource = 2
)

// Enum value maps for SnapSource.
var (
	SnapSource_name = map[int32]string{
		0: "SNAP_SOURCE_UNKNOWN",
		1: "SNAP_SOURCE_CAMERA",
		2: "SNAP_SOURCE_MEMORY",
	}
	SnapSource_value = map[string]int32{
		"SNAP_SOURCE_UNKNOWN": 0,
		"SNAP_SOURCE_CAMERA":  1,
		"SNAP_SOURCE_MEMORY":  2,
	}
)

func (x SnapSource) Enum() *SnapSource {
	p := new(SnapSource)
	*p = x
	return p
}

func (x SnapSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SnapSource) Descriptor() protoreflect.EnumDescriptor {
	return file_unlockables_proto_enumTypes[0].Descriptor()
}

func (SnapSource) Type() protoreflect.EnumType {
	return &file_unlockables_proto_enumTypes[0]
}

func (x SnapSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SnapSource.Descriptor instead.
func (SnapSource) EnumDescriptor() ([]byte, []int) {
	return file_unlockables_proto_rawDescGZIP(), []int{0}
}

type LensSource int32

const (
	LensSource_LENS_SOURCE_UNKNOWN      LensSource = 0
	LensSource_LENS_SOURCE_CAMERA       LensSource = 1
	LensSource_LENS_SOURCE_VIDEOCHAT    LensSource = 2
	LensSource_LENS_SOURCE_SNAPCODE     LensSource = 3
	LensSource_LENS_SOURCE_LOGIN_CAMERA LensSource = 4
	LensSource_LENS_SOURCE_PREVIEW      LensSource = 5
	LensSource_LENS_SOURCE_MEMORIES     LensSource = 6
	LensSource_LENS_SOURCE_ON_DEMAND    LensSource = 7
	LensSource_LENS_SOURCE_DISCOVER     LensSource = 8
	LensSource_LENS_SOURCE_STORY        LensSource = 9
)

// Enum value maps for LensSource.
var (
	LensSource_name = map[int32]string{
		0: "LENS_SOURCE_UNKNOWN",
		1: "LENS_SOURCE_CAMERA",
		2: "LENS_SOURCE_VIDEOCHAT",
		3: "LENS_SOURCE_SNAPCODE",
		4: "LENS_SOURCE_LOGIN_CAMERA",
		5: "LENS_SOURCE_PREVIEW",
		6: "LENS_SOURCE_MEMORIES",
		7: "LENS_SOURCE_ON_DEMAND",
		8: "LENS_SOURCE_DISCOVER",
		9: "LENS_SOURCE_STORY",
	}
	LensSource_value = map[string]int32{
		"LENS_SOURCE_UNKNOWN":      0,
		"LENS_SOURCE_CAMERA":       1,
		"LENS_SOURCE_VIDEOCHAT":    2,
		"LENS_SOURCE_SNAPCODE":     3,
		"LENS_SOURCE_LOGIN_CAMERA": 4,
		"LENS_SOURCE_PREVIEW":      5,
		"LENS_SOURCE_MEMORIES":     6,
		"LENS_SOURCE_ON_DEMAND":    7,
		"LENS_SOURCE_DISCOVER":     8,
		"LENS_SOURCE_STORY":        9,
	}
)

func (x LensSource) Enum() *LensSource {
	p := new(LensSource)
	*p = x
	return p
}

func (x LensSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LensSource) Descriptor() protoreflect.EnumDescriptor {
	return file_unlockables_proto_enumTypes[1].Descriptor()
}

func (LensSource) Type() protoreflect.EnumType {
	return &file_unlockables_proto_enumTypes[1]
}

func (x LensSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LensSource.Descriptor instead.
func (LensSource) EnumDescriptor() ([]byte, []int) {
	return file_unlockables_proto_rawDescGZIP(), []int{1}
}

type SnapDoc_Unlockables struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnlockablesSnapInfo *UnlockablesSnapInfo `protobuf:"bytes,1,opt,name=unlockablesSnapInfo,proto3" json:"unlockablesSnapInfo,omitempty"`
	EncryptedGeoData    []byte               `protobuf:"bytes,2,opt,name=encryptedGeoData,proto3" json:"encryptedGeoData,omitempty"`
}

func (x *SnapDoc_Unlockables) Reset() {
	*x = SnapDoc_Unlockables{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unlockables_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapDoc_Unlockables) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapDoc_Unlockables) ProtoMessage() {}

func (x *SnapDoc_Unlockables) ProtoReflect() protoreflect.Message {
	mi := &file_unlockables_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapDoc_Unlockables.ProtoReflect.Descriptor instead.
func (*SnapDoc_Unlockables) Descriptor() ([]byte, []int) {
	return file_unlockables_proto_rawDescGZIP(), []int{0}
}

func (x *SnapDoc_Unlockables) GetUnlockablesSnapInfo() *UnlockablesSnapInfo {
	if x != nil {
		return x.UnlockablesSnapInfo
	}
	return nil
}

func (x *SnapDoc_Unlockables) GetEncryptedGeoData() []byte {
	if x != nil {
		return x.EncryptedGeoData
	}
	return nil
}

type UnlockablesSnapInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SnapSessionId string         `protobuf:"bytes,1,opt,name=snapSessionId,proto3" json:"snapSessionId,omitempty"`
	Filters       []*FilterInfo  `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
	Lenses        []*LensInfo    `protobuf:"bytes,3,rep,name=lenses,proto3" json:"lenses,omitempty"`
	Stickers      []*StickerInfo `protobuf:"bytes,4,rep,name=stickers,proto3" json:"stickers,omitempty"`
	SnapSource    SnapSource     `protobuf:"varint,5,opt,name=snapSource,proto3,enum=unlockables.SnapSource" json:"snapSource,omitempty"`
}

func (x *UnlockablesSnapInfo) Reset() {
	*x = UnlockablesSnapInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unlockables_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockablesSnapInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockablesSnapInfo) ProtoMessage() {}

func (x *UnlockablesSnapInfo) ProtoReflect() protoreflect.Message {
	mi := &file_unlockables_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockablesSnapInfo.ProtoReflect.Descriptor instead.
func (*UnlockablesSnapInfo) Descriptor() ([]byte, []int) {
	return file_unlockables_proto_rawDescGZIP(), []int{1}
}

func (x *UnlockablesSnapInfo) GetSnapSessionId() string {
	if x != nil {
		return x.SnapSessionId
	}
	return ""
}

func (x *UnlockablesSnapInfo) GetFilters() []*FilterInfo {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *UnlockablesSnapInfo) GetLenses() []*LensInfo {
	if x != nil {
		return x.Lenses
	}
	return nil
}

func (x *UnlockablesSnapInfo) GetStickers() []*StickerInfo {
	if x != nil {
		return x.Stickers
	}
	return nil
}

func (x *UnlockablesSnapInfo) GetSnapSource() SnapSource {
	if x != nil {
		return x.SnapSource
	}
	return SnapSource_SNAP_SOURCE_UNKNOWN
}

type FilterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnlockableId int64 `protobuf:"varint,1,opt,name=unlockableId,proto3" json:"unlockableId,omitempty"`
	IsUco        bool  `protobuf:"varint,2,opt,name=isUco,proto3" json:"isUco,omitempty"`
}

func (x *FilterInfo) Reset() {
	*x = FilterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unlockables_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterInfo) ProtoMessage() {}

func (x *FilterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_unlockables_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterInfo.ProtoReflect.Descriptor instead.
func (*FilterInfo) Descriptor() ([]byte, []int) {
	return file_unlockables_proto_rawDescGZIP(), []int{2}
}

func (x *FilterInfo) GetUnlockableId() int64 {
	if x != nil {
		return x.UnlockableId
	}
	return 0
}

func (x *FilterInfo) GetIsUco() bool {
	if x != nil {
		return x.IsUco
	}
	return false
}

type LensInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnlockableId int64      `protobuf:"varint,1,opt,name=unlockableId,proto3" json:"unlockableId,omitempty"`
	OptionId     int64      `protobuf:"varint,2,opt,name=optionId,proto3" json:"optionId,omitempty"`
	Source       LensSource `protobuf:"varint,3,opt,name=source,proto3,enum=unlockables.LensSource" json:"source,omitempty"`
}

func (x *LensInfo) Reset() {
	*x = LensInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unlockables_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LensInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LensInfo) ProtoMessage() {}

func (x *LensInfo) ProtoReflect() protoreflect.Message {
	mi := &file_unlockables_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LensInfo.ProtoReflect.Descriptor instead.
func (*LensInfo) Descriptor() ([]byte, []int) {
	return file_unlockables_proto_rawDescGZIP(), []int{3}
}

func (x *LensInfo) GetUnlockableId() int64 {
	if x != nil {
		return x.UnlockableId
	}
	return 0
}

func (x *LensInfo) GetOptionId() int64 {
	if x != nil {
		return x.OptionId
	}
	return 0
}

func (x *LensInfo) GetSource() LensSource {
	if x != nil {
		return x.Source
	}
	return LensSource_LENS_SOURCE_UNKNOWN
}

type StickerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnlockableId int64 `protobuf:"varint,1,opt,name=unlockableId,proto3" json:"unlockableId,omitempty"`
	StickerId    int64 `protobuf:"varint,2,opt,name=stickerId,proto3" json:"stickerId,omitempty"`
}

func (x *StickerInfo) Reset() {
	*x = StickerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unlockables_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StickerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StickerInfo) ProtoMessage() {}

func (x *StickerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_unlockables_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StickerInfo.ProtoReflect.Descriptor instead.
func (*StickerInfo) Descriptor() ([]byte, []int) {
	return file_unlockables_proto_rawDescGZIP(), []int{4}
}

func (x *StickerInfo) GetUnlockableId() int64 {
	if x != nil {
		return x.UnlockableId
	}
	return 0
}

func (x *StickerInfo) GetStickerId() int64 {
	if x != nil {
		return x.StickerId
	}
	return 0
}

var File_unlockables_proto protoreflect.FileDescriptor

var file_unlockables_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x22, 0x95, 0x01, 0x0a, 0x13, 0x53, 0x6e, 0x61, 0x70, 0x44, 0x6f, 0x63, 0x5f, 0x55, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x52, 0x0a, 0x13, 0x75, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x53, 0x6e, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x53,
	0x6e, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x13, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x53, 0x6e, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x10,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x47, 0x65, 0x6f, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65,
	0x64, 0x47, 0x65, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x22, 0x8c, 0x02, 0x0a, 0x13, 0x55, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x53, 0x6e, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x24, 0x0a, 0x0d, 0x73, 0x6e, 0x61, 0x70, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6e, 0x61, 0x70, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x6c, 0x65, 0x6e,
	0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x06, 0x6c, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x08, 0x73, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x37,
	0x0a, 0x0a, 0x73, 0x6e, 0x61, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x2e, 0x53, 0x6e, 0x61, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x73, 0x6e, 0x61,
	0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x46, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x61,
	0x62, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x75, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x55,
	0x63, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x55, 0x63, 0x6f, 0x22,
	0x7b, 0x0a, 0x08, 0x4c, 0x65, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x75,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x75, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x4c, 0x65, 0x6e, 0x73, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x4f, 0x0a, 0x0b,
	0x53, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x75,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x55, 0x0a,
	0x0a, 0x53, 0x6e, 0x61, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x53,
	0x4e, 0x41, 0x50, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x4e, 0x41, 0x50, 0x5f, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x43, 0x41, 0x4d, 0x45, 0x52, 0x41, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12,
	0x53, 0x4e, 0x41, 0x50, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4d, 0x45, 0x4d, 0x4f,
	0x52, 0x59, 0x10, 0x02, 0x2a, 0x8f, 0x02, 0x0a, 0x0a, 0x4c, 0x65, 0x6e, 0x73, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x45, 0x4e, 0x53, 0x5f, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12,
	0x4c, 0x45, 0x4e, 0x53, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x43, 0x41, 0x4d, 0x45,
	0x52, 0x41, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x45, 0x4e, 0x53, 0x5f, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x43, 0x48, 0x41, 0x54, 0x10, 0x02, 0x12,
	0x18, 0x0a, 0x14, 0x4c, 0x45, 0x4e, 0x53, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53,
	0x4e, 0x41, 0x50, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x4c, 0x45, 0x4e,
	0x53, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x43,
	0x41, 0x4d, 0x45, 0x52, 0x41, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x45, 0x4e, 0x53, 0x5f,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x10, 0x05,
	0x12, 0x18, 0x0a, 0x14, 0x4c, 0x45, 0x4e, 0x53, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f,
	0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x49, 0x45, 0x53, 0x10, 0x06, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x45,
	0x4e, 0x53, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4d,
	0x41, 0x4e, 0x44, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x45, 0x4e, 0x53, 0x5f, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x10, 0x08, 0x12,
	0x15, 0x0a, 0x11, 0x4c, 0x45, 0x4e, 0x53, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53,
	0x54, 0x4f, 0x52, 0x59, 0x10, 0x09, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_unlockables_proto_rawDescOnce sync.Once
	file_unlockables_proto_rawDescData = file_unlockables_proto_rawDesc
)

func file_unlockables_proto_rawDescGZIP() []byte {
	file_unlockables_proto_rawDescOnce.Do(func() {
		file_unlockables_proto_rawDescData = protoimpl.X.CompressGZIP(file_unlockables_proto_rawDescData)
	})
	return file_unlockables_proto_rawDescData
}

var file_unlockables_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_unlockables_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_unlockables_proto_goTypes = []interface{}{
	(SnapSource)(0),             // 0: unlockables.SnapSource
	(LensSource)(0),             // 1: unlockables.LensSource
	(*SnapDoc_Unlockables)(nil), // 2: unlockables.SnapDoc_Unlockables
	(*UnlockablesSnapInfo)(nil), // 3: unlockables.UnlockablesSnapInfo
	(*FilterInfo)(nil),          // 4: unlockables.FilterInfo
	(*LensInfo)(nil),            // 5: unlockables.LensInfo
	(*StickerInfo)(nil),         // 6: unlockables.StickerInfo
}
var file_unlockables_proto_depIdxs = []int32{
	3, // 0: unlockables.SnapDoc_Unlockables.unlockablesSnapInfo:type_name -> unlockables.UnlockablesSnapInfo
	4, // 1: unlockables.UnlockablesSnapInfo.filters:type_name -> unlockables.FilterInfo
	5, // 2: unlockables.UnlockablesSnapInfo.lenses:type_name -> unlockables.LensInfo
	6, // 3: unlockables.UnlockablesSnapInfo.stickers:type_name -> unlockables.StickerInfo
	0, // 4: unlockables.UnlockablesSnapInfo.snapSource:type_name -> unlockables.SnapSource
	1, // 5: unlockables.LensInfo.source:type_name -> unlockables.LensSource
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_unlockables_proto_init() }
func file_unlockables_proto_init() {
	if File_unlockables_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_unlockables_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapDoc_Unlockables); i {
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
		file_unlockables_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockablesSnapInfo); i {
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
		file_unlockables_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterInfo); i {
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
		file_unlockables_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LensInfo); i {
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
		file_unlockables_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StickerInfo); i {
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
			RawDescriptor: file_unlockables_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_unlockables_proto_goTypes,
		DependencyIndexes: file_unlockables_proto_depIdxs,
		EnumInfos:         file_unlockables_proto_enumTypes,
		MessageInfos:      file_unlockables_proto_msgTypes,
	}.Build()
	File_unlockables_proto = out.File
	file_unlockables_proto_rawDesc = nil
	file_unlockables_proto_goTypes = nil
	file_unlockables_proto_depIdxs = nil
}