// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.3.0
// source: diffservice.proto

package diffstore

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diffservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_diffservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_diffservice_proto_rawDescGZIP(), []int{0}
}

type GetDiffsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainDigest   string   `protobuf:"bytes,2,opt,name=mainDigest,proto3" json:"mainDigest,omitempty"`
	RightDigests []string `protobuf:"bytes,3,rep,name=rightDigests,proto3" json:"rightDigests,omitempty"`
}

func (x *GetDiffsRequest) Reset() {
	*x = GetDiffsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diffservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDiffsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDiffsRequest) ProtoMessage() {}

func (x *GetDiffsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_diffservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDiffsRequest.ProtoReflect.Descriptor instead.
func (*GetDiffsRequest) Descriptor() ([]byte, []int) {
	return file_diffservice_proto_rawDescGZIP(), []int{1}
}

func (x *GetDiffsRequest) GetMainDigest() string {
	if x != nil {
		return x.MainDigest
	}
	return ""
}

func (x *GetDiffsRequest) GetRightDigests() []string {
	if x != nil {
		return x.RightDigests
	}
	return nil
}

type GetDiffsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Diffs []byte `protobuf:"bytes,1,opt,name=diffs,proto3" json:"diffs,omitempty"`
}

func (x *GetDiffsResponse) Reset() {
	*x = GetDiffsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diffservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDiffsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDiffsResponse) ProtoMessage() {}

func (x *GetDiffsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_diffservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDiffsResponse.ProtoReflect.Descriptor instead.
func (*GetDiffsResponse) Descriptor() ([]byte, []int) {
	return file_diffservice_proto_rawDescGZIP(), []int{2}
}

func (x *GetDiffsResponse) GetDiffs() []byte {
	if x != nil {
		return x.Diffs
	}
	return nil
}

type PurgeDigestsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digests  []string `protobuf:"bytes,1,rep,name=digests,proto3" json:"digests,omitempty"`
	PurgeGCS bool     `protobuf:"varint,2,opt,name=purgeGCS,proto3" json:"purgeGCS,omitempty"`
}

func (x *PurgeDigestsRequest) Reset() {
	*x = PurgeDigestsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diffservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurgeDigestsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurgeDigestsRequest) ProtoMessage() {}

func (x *PurgeDigestsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_diffservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurgeDigestsRequest.ProtoReflect.Descriptor instead.
func (*PurgeDigestsRequest) Descriptor() ([]byte, []int) {
	return file_diffservice_proto_rawDescGZIP(), []int{3}
}

func (x *PurgeDigestsRequest) GetDigests() []string {
	if x != nil {
		return x.Digests
	}
	return nil
}

func (x *PurgeDigestsRequest) GetPurgeGCS() bool {
	if x != nil {
		return x.PurgeGCS
	}
	return false
}

type UnavailableDigestsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DigestFailures map[string]*DigestFailureResponse `protobuf:"bytes,1,rep,name=digestFailures,proto3" json:"digestFailures,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UnavailableDigestsResponse) Reset() {
	*x = UnavailableDigestsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diffservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnavailableDigestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnavailableDigestsResponse) ProtoMessage() {}

func (x *UnavailableDigestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_diffservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnavailableDigestsResponse.ProtoReflect.Descriptor instead.
func (*UnavailableDigestsResponse) Descriptor() ([]byte, []int) {
	return file_diffservice_proto_rawDescGZIP(), []int{4}
}

func (x *UnavailableDigestsResponse) GetDigestFailures() map[string]*DigestFailureResponse {
	if x != nil {
		return x.DigestFailures
	}
	return nil
}

type DigestFailureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Digest string `protobuf:"bytes,1,opt,name=Digest,proto3" json:"Digest,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=Reason,proto3" json:"Reason,omitempty"`
	TS     int64  `protobuf:"varint,3,opt,name=TS,proto3" json:"TS,omitempty"`
}

func (x *DigestFailureResponse) Reset() {
	*x = DigestFailureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_diffservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DigestFailureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DigestFailureResponse) ProtoMessage() {}

func (x *DigestFailureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_diffservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DigestFailureResponse.ProtoReflect.Descriptor instead.
func (*DigestFailureResponse) Descriptor() ([]byte, []int) {
	return file_diffservice_proto_rawDescGZIP(), []int{5}
}

func (x *DigestFailureResponse) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *DigestFailureResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *DigestFailureResponse) GetTS() int64 {
	if x != nil {
		return x.TS
	}
	return 0
}

var File_diffservice_proto protoreflect.FileDescriptor

var file_diffservice_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x69, 0x66, 0x66, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x69, 0x66, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x55, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x69,
	0x66, 0x66, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61,
	0x69, 0x6e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x61, 0x69, 0x6e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x69,
	0x67, 0x68, 0x74, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x69, 0x67, 0x68, 0x74, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x22, 0x28,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x44, 0x69, 0x66, 0x66, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x69, 0x66, 0x66, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x64, 0x69, 0x66, 0x66, 0x73, 0x22, 0x4b, 0x0a, 0x13, 0x50, 0x75, 0x72, 0x67,
	0x65, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x75, 0x72,
	0x67, 0x65, 0x47, 0x43, 0x53, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x70, 0x75, 0x72,
	0x67, 0x65, 0x47, 0x43, 0x53, 0x22, 0xe4, 0x01, 0x0a, 0x1a, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x0e, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x64,
	0x69, 0x66, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x1a, 0x63, 0x0a, 0x13, 0x44, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x36, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x57, 0x0a, 0x15,
	0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x53, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x54, 0x53, 0x32, 0x97, 0x02, 0x0a, 0x0b, 0x44, 0x69, 0x66, 0x66, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x44, 0x69, 0x66, 0x66,
	0x73, 0x12, 0x1a, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x69, 0x66, 0x66, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x64, 0x69, 0x66, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x66,
	0x66, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x12,
	0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x69, 0x67, 0x65, 0x73,
	0x74, 0x73, 0x12, 0x10, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x25, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a,
	0x0c, 0x50, 0x75, 0x72, 0x67, 0x65, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1e, 0x2e,
	0x64, 0x69, 0x66, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x75, 0x72, 0x67, 0x65, 0x44,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x64, 0x69, 0x66, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x2c, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x2e, 0x64, 0x69, 0x66, 0x66,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x64, 0x69,
	0x66, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x64, 0x69, 0x66, 0x66, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_diffservice_proto_rawDescOnce sync.Once
	file_diffservice_proto_rawDescData = file_diffservice_proto_rawDesc
)

func file_diffservice_proto_rawDescGZIP() []byte {
	file_diffservice_proto_rawDescOnce.Do(func() {
		file_diffservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_diffservice_proto_rawDescData)
	})
	return file_diffservice_proto_rawDescData
}

var file_diffservice_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_diffservice_proto_goTypes = []interface{}{
	(*Empty)(nil),                      // 0: diffstore.Empty
	(*GetDiffsRequest)(nil),            // 1: diffstore.GetDiffsRequest
	(*GetDiffsResponse)(nil),           // 2: diffstore.GetDiffsResponse
	(*PurgeDigestsRequest)(nil),        // 3: diffstore.PurgeDigestsRequest
	(*UnavailableDigestsResponse)(nil), // 4: diffstore.UnavailableDigestsResponse
	(*DigestFailureResponse)(nil),      // 5: diffstore.DigestFailureResponse
	nil,                                // 6: diffstore.UnavailableDigestsResponse.DigestFailuresEntry
}
var file_diffservice_proto_depIdxs = []int32{
	6, // 0: diffstore.UnavailableDigestsResponse.digestFailures:type_name -> diffstore.UnavailableDigestsResponse.DigestFailuresEntry
	5, // 1: diffstore.UnavailableDigestsResponse.DigestFailuresEntry.value:type_name -> diffstore.DigestFailureResponse
	1, // 2: diffstore.DiffService.GetDiffs:input_type -> diffstore.GetDiffsRequest
	0, // 3: diffstore.DiffService.UnavailableDigests:input_type -> diffstore.Empty
	3, // 4: diffstore.DiffService.PurgeDigests:input_type -> diffstore.PurgeDigestsRequest
	0, // 5: diffstore.DiffService.Ping:input_type -> diffstore.Empty
	2, // 6: diffstore.DiffService.GetDiffs:output_type -> diffstore.GetDiffsResponse
	4, // 7: diffstore.DiffService.UnavailableDigests:output_type -> diffstore.UnavailableDigestsResponse
	0, // 8: diffstore.DiffService.PurgeDigests:output_type -> diffstore.Empty
	0, // 9: diffstore.DiffService.Ping:output_type -> diffstore.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_diffservice_proto_init() }
func file_diffservice_proto_init() {
	if File_diffservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_diffservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_diffservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDiffsRequest); i {
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
		file_diffservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDiffsResponse); i {
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
		file_diffservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurgeDigestsRequest); i {
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
		file_diffservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnavailableDigestsResponse); i {
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
		file_diffservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DigestFailureResponse); i {
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
			RawDescriptor: file_diffservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_diffservice_proto_goTypes,
		DependencyIndexes: file_diffservice_proto_depIdxs,
		MessageInfos:      file_diffservice_proto_msgTypes,
	}.Build()
	File_diffservice_proto = out.File
	file_diffservice_proto_rawDesc = nil
	file_diffservice_proto_goTypes = nil
	file_diffservice_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DiffServiceClient is the client API for DiffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiffServiceClient interface {
	// Same functionality as Get in the diff.DiffStore interface.
	GetDiffs(ctx context.Context, in *GetDiffsRequest, opts ...grpc.CallOption) (*GetDiffsResponse, error)
	// Same functionality asSee UnavailableDigests in the diff.DiffStore interface.
	UnavailableDigests(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UnavailableDigestsResponse, error)
	//Same functionality asSee PurgeDigestset in the diff.DiffStore interface.
	PurgeDigests(ctx context.Context, in *PurgeDigestsRequest, opts ...grpc.CallOption) (*Empty, error)
	// Ping is used to test connection.
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type diffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiffServiceClient(cc grpc.ClientConnInterface) DiffServiceClient {
	return &diffServiceClient{cc}
}

func (c *diffServiceClient) GetDiffs(ctx context.Context, in *GetDiffsRequest, opts ...grpc.CallOption) (*GetDiffsResponse, error) {
	out := new(GetDiffsResponse)
	err := c.cc.Invoke(ctx, "/diffstore.DiffService/GetDiffs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diffServiceClient) UnavailableDigests(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UnavailableDigestsResponse, error) {
	out := new(UnavailableDigestsResponse)
	err := c.cc.Invoke(ctx, "/diffstore.DiffService/UnavailableDigests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diffServiceClient) PurgeDigests(ctx context.Context, in *PurgeDigestsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/diffstore.DiffService/PurgeDigests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diffServiceClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/diffstore.DiffService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiffServiceServer is the server API for DiffService service.
type DiffServiceServer interface {
	// Same functionality as Get in the diff.DiffStore interface.
	GetDiffs(context.Context, *GetDiffsRequest) (*GetDiffsResponse, error)
	// Same functionality asSee UnavailableDigests in the diff.DiffStore interface.
	UnavailableDigests(context.Context, *Empty) (*UnavailableDigestsResponse, error)
	//Same functionality asSee PurgeDigestset in the diff.DiffStore interface.
	PurgeDigests(context.Context, *PurgeDigestsRequest) (*Empty, error)
	// Ping is used to test connection.
	Ping(context.Context, *Empty) (*Empty, error)
}

// UnimplementedDiffServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDiffServiceServer struct {
}

func (*UnimplementedDiffServiceServer) GetDiffs(context.Context, *GetDiffsRequest) (*GetDiffsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiffs not implemented")
}
func (*UnimplementedDiffServiceServer) UnavailableDigests(context.Context, *Empty) (*UnavailableDigestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnavailableDigests not implemented")
}
func (*UnimplementedDiffServiceServer) PurgeDigests(context.Context, *PurgeDigestsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurgeDigests not implemented")
}
func (*UnimplementedDiffServiceServer) Ping(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterDiffServiceServer(s *grpc.Server, srv DiffServiceServer) {
	s.RegisterService(&_DiffService_serviceDesc, srv)
}

func _DiffService_GetDiffs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiffsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServiceServer).GetDiffs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diffstore.DiffService/GetDiffs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServiceServer).GetDiffs(ctx, req.(*GetDiffsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiffService_UnavailableDigests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServiceServer).UnavailableDigests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diffstore.DiffService/UnavailableDigests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServiceServer).UnavailableDigests(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiffService_PurgeDigests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeDigestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServiceServer).PurgeDigests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diffstore.DiffService/PurgeDigests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServiceServer).PurgeDigests(ctx, req.(*PurgeDigestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiffService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiffServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/diffstore.DiffService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiffServiceServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiffService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "diffstore.DiffService",
	HandlerType: (*DiffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDiffs",
			Handler:    _DiffService_GetDiffs_Handler,
		},
		{
			MethodName: "UnavailableDigests",
			Handler:    _DiffService_UnavailableDigests_Handler,
		},
		{
			MethodName: "PurgeDigests",
			Handler:    _DiffService_PurgeDigests_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _DiffService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "diffservice.proto",
}
