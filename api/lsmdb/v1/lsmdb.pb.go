// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: api/lsmdb/v1/lsmdb.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// 命名必须使用下划线，否则出现传递不了参数的情况
type OpenDBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}

func (x *OpenDBRequest) Reset() {
	*x = OpenDBRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenDBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenDBRequest) ProtoMessage() {}

func (x *OpenDBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenDBRequest.ProtoReflect.Descriptor instead.
func (*OpenDBRequest) Descriptor() ([]byte, []int) {
	return file_api_lsmdb_v1_lsmdb_proto_rawDescGZIP(), []int{0}
}

func (x *OpenDBRequest) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

// The response message containing the value about register success or not .
type OpenDBReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *OpenDBReply) Reset() {
	*x = OpenDBReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenDBReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenDBReply) ProtoMessage() {}

func (x *OpenDBReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenDBReply.ProtoReflect.Descriptor instead.
func (*OpenDBReply) Descriptor() ([]byte, []int) {
	return file_api_lsmdb_v1_lsmdb_proto_rawDescGZIP(), []int{1}
}

func (x *OpenDBReply) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type PutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PutRequest) Reset() {
	*x = PutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRequest) ProtoMessage() {}

func (x *PutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRequest.ProtoReflect.Descriptor instead.
func (*PutRequest) Descriptor() ([]byte, []int) {
	return file_api_lsmdb_v1_lsmdb_proto_rawDescGZIP(), []int{2}
}

func (x *PutRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type PutReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data bool `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PutReply) Reset() {
	*x = PutReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutReply) ProtoMessage() {}

func (x *PutReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutReply.ProtoReflect.Descriptor instead.
func (*PutReply) Descriptor() ([]byte, []int) {
	return file_api_lsmdb_v1_lsmdb_proto_rawDescGZIP(), []int{3}
}

func (x *PutReply) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

type PutStrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PutStrRequest) Reset() {
	*x = PutStrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutStrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutStrRequest) ProtoMessage() {}

func (x *PutStrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutStrRequest.ProtoReflect.Descriptor instead.
func (*PutStrRequest) Descriptor() ([]byte, []int) {
	return file_api_lsmdb_v1_lsmdb_proto_rawDescGZIP(), []int{4}
}

func (x *PutStrRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutStrRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PutStrReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data bool `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PutStrReply) Reset() {
	*x = PutStrReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutStrReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutStrReply) ProtoMessage() {}

func (x *PutStrReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutStrReply.ProtoReflect.Descriptor instead.
func (*PutStrReply) Descriptor() ([]byte, []int) {
	return file_api_lsmdb_v1_lsmdb_proto_rawDescGZIP(), []int{5}
}

func (x *PutStrReply) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

// The request message containing mac addr and ip.
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_api_lsmdb_v1_lsmdb_proto_rawDescGZIP(), []int{6}
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// The response message containing the value about register success or not .
type GetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetReply) Reset() {
	*x = GetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReply) ProtoMessage() {}

func (x *GetReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_lsmdb_v1_lsmdb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReply.ProtoReflect.Descriptor instead.
func (*GetReply) Descriptor() ([]byte, []int) {
	return file_api_lsmdb_v1_lsmdb_proto_rawDescGZIP(), []int{7}
}

func (x *GetReply) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_api_lsmdb_v1_lsmdb_proto protoreflect.FileDescriptor

var file_api_lsmdb_v1_lsmdb_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x73, 0x6d, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6c,
	0x73, 0x6d, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x73, 0x6d, 0x64,
	0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x0b,
	0x4f, 0x70, 0x65, 0x6e, 0x44, 0x42, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x34, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x08, 0x50, 0x75, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x37, 0x0a, 0x0d, 0x50, 0x75, 0x74, 0x53, 0x74,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x21, 0x0a, 0x0b, 0x50, 0x75, 0x74, 0x53, 0x74, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x1e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x20, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xc2, 0x02, 0x0a, 0x05, 0x4c, 0x73, 0x6d, 0x64, 0x62, 0x12,
	0x52, 0x0a, 0x06, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x42, 0x12, 0x17, 0x2e, 0x6c, 0x73, 0x6d, 0x64,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6c, 0x73, 0x6d, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70,
	0x65, 0x6e, 0x44, 0x42, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x6c, 0x73, 0x6d, 0x64, 0x62, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x64, 0x62, 0x12, 0x46, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x14, 0x2e, 0x6c, 0x73, 0x6d,
	0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x6c, 0x73, 0x6d, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22,
	0x0a, 0x2f, 0x6c, 0x73, 0x6d, 0x64, 0x62, 0x2f, 0x70, 0x75, 0x74, 0x12, 0x52, 0x0a, 0x06, 0x50,
	0x75, 0x74, 0x53, 0x74, 0x72, 0x12, 0x17, 0x2e, 0x6c, 0x73, 0x6d, 0x64, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x75, 0x74, 0x53, 0x74, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x6c, 0x73, 0x6d, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x74, 0x53, 0x74, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a,
	0x22, 0x0d, 0x2f, 0x6c, 0x73, 0x6d, 0x64, 0x62, 0x2f, 0x70, 0x75, 0x74, 0x73, 0x74, 0x72, 0x12,
	0x49, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x6c, 0x73, 0x6d, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6c,
	0x73, 0x6d, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x6c, 0x73, 0x6d, 0x64, 0x62,
	0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x7d, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x4c, 0x53, 0x4d, 0x44, 0x42, 0x2f,
	0x6c, 0x73, 0x6d, 0x64, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6c, 0x73, 0x6d, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_lsmdb_v1_lsmdb_proto_rawDescOnce sync.Once
	file_api_lsmdb_v1_lsmdb_proto_rawDescData = file_api_lsmdb_v1_lsmdb_proto_rawDesc
)

func file_api_lsmdb_v1_lsmdb_proto_rawDescGZIP() []byte {
	file_api_lsmdb_v1_lsmdb_proto_rawDescOnce.Do(func() {
		file_api_lsmdb_v1_lsmdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_lsmdb_v1_lsmdb_proto_rawDescData)
	})
	return file_api_lsmdb_v1_lsmdb_proto_rawDescData
}

var file_api_lsmdb_v1_lsmdb_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_lsmdb_v1_lsmdb_proto_goTypes = []interface{}{
	(*OpenDBRequest)(nil), // 0: lsmdb.v1.OpenDBRequest
	(*OpenDBReply)(nil),   // 1: lsmdb.v1.OpenDBReply
	(*PutRequest)(nil),    // 2: lsmdb.v1.PutRequest
	(*PutReply)(nil),      // 3: lsmdb.v1.PutReply
	(*PutStrRequest)(nil), // 4: lsmdb.v1.PutStrRequest
	(*PutStrReply)(nil),   // 5: lsmdb.v1.PutStrReply
	(*GetRequest)(nil),    // 6: lsmdb.v1.GetRequest
	(*GetReply)(nil),      // 7: lsmdb.v1.GetReply
}
var file_api_lsmdb_v1_lsmdb_proto_depIdxs = []int32{
	0, // 0: lsmdb.v1.Lsmdb.OpenDB:input_type -> lsmdb.v1.OpenDBRequest
	2, // 1: lsmdb.v1.Lsmdb.Put:input_type -> lsmdb.v1.PutRequest
	4, // 2: lsmdb.v1.Lsmdb.PutStr:input_type -> lsmdb.v1.PutStrRequest
	6, // 3: lsmdb.v1.Lsmdb.Get:input_type -> lsmdb.v1.GetRequest
	1, // 4: lsmdb.v1.Lsmdb.OpenDB:output_type -> lsmdb.v1.OpenDBReply
	3, // 5: lsmdb.v1.Lsmdb.Put:output_type -> lsmdb.v1.PutReply
	5, // 6: lsmdb.v1.Lsmdb.PutStr:output_type -> lsmdb.v1.PutStrReply
	7, // 7: lsmdb.v1.Lsmdb.Get:output_type -> lsmdb.v1.GetReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_lsmdb_v1_lsmdb_proto_init() }
func file_api_lsmdb_v1_lsmdb_proto_init() {
	if File_api_lsmdb_v1_lsmdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_lsmdb_v1_lsmdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenDBRequest); i {
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
		file_api_lsmdb_v1_lsmdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenDBReply); i {
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
		file_api_lsmdb_v1_lsmdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRequest); i {
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
		file_api_lsmdb_v1_lsmdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutReply); i {
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
		file_api_lsmdb_v1_lsmdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutStrRequest); i {
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
		file_api_lsmdb_v1_lsmdb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutStrReply); i {
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
		file_api_lsmdb_v1_lsmdb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_api_lsmdb_v1_lsmdb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReply); i {
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
			RawDescriptor: file_api_lsmdb_v1_lsmdb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_lsmdb_v1_lsmdb_proto_goTypes,
		DependencyIndexes: file_api_lsmdb_v1_lsmdb_proto_depIdxs,
		MessageInfos:      file_api_lsmdb_v1_lsmdb_proto_msgTypes,
	}.Build()
	File_api_lsmdb_v1_lsmdb_proto = out.File
	file_api_lsmdb_v1_lsmdb_proto_rawDesc = nil
	file_api_lsmdb_v1_lsmdb_proto_goTypes = nil
	file_api_lsmdb_v1_lsmdb_proto_depIdxs = nil
}
