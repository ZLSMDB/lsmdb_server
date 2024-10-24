// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: api/etcd/v1/etcd.proto

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

type EtcdPutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EtcdPutRequest) Reset() {
	*x = EtcdPutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_etcd_v1_etcd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EtcdPutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EtcdPutRequest) ProtoMessage() {}

func (x *EtcdPutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_etcd_v1_etcd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EtcdPutRequest.ProtoReflect.Descriptor instead.
func (*EtcdPutRequest) Descriptor() ([]byte, []int) {
	return file_api_etcd_v1_etcd_proto_rawDescGZIP(), []int{0}
}

func (x *EtcdPutRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *EtcdPutRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type EtcdPutReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EtcdPutReply) Reset() {
	*x = EtcdPutReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_etcd_v1_etcd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EtcdPutReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EtcdPutReply) ProtoMessage() {}

func (x *EtcdPutReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_etcd_v1_etcd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EtcdPutReply.ProtoReflect.Descriptor instead.
func (*EtcdPutReply) Descriptor() ([]byte, []int) {
	return file_api_etcd_v1_etcd_proto_rawDescGZIP(), []int{1}
}

func (x *EtcdPutReply) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type EtcdGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *EtcdGetRequest) Reset() {
	*x = EtcdGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_etcd_v1_etcd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EtcdGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EtcdGetRequest) ProtoMessage() {}

func (x *EtcdGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_etcd_v1_etcd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EtcdGetRequest.ProtoReflect.Descriptor instead.
func (*EtcdGetRequest) Descriptor() ([]byte, []int) {
	return file_api_etcd_v1_etcd_proto_rawDescGZIP(), []int{2}
}

func (x *EtcdGetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type EtcdGetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EtcdGetReply) Reset() {
	*x = EtcdGetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_etcd_v1_etcd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EtcdGetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EtcdGetReply) ProtoMessage() {}

func (x *EtcdGetReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_etcd_v1_etcd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EtcdGetReply.ProtoReflect.Descriptor instead.
func (*EtcdGetReply) Descriptor() ([]byte, []int) {
	return file_api_etcd_v1_etcd_proto_rawDescGZIP(), []int{3}
}

func (x *EtcdGetReply) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_api_etcd_v1_etcd_proto protoreflect.FileDescriptor

var file_api_etcd_v1_etcd_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x74, 0x63, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x74,
	0x63, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x74, 0x63, 0x64, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x38, 0x0a, 0x0e, 0x45, 0x74, 0x63, 0x64, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x24, 0x0a, 0x0c, 0x45, 0x74, 0x63,
	0x64, 0x50, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x22, 0x0a, 0x0e, 0x45, 0x74, 0x63, 0x64, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x24, 0x0a, 0x0c, 0x45, 0x74, 0x63, 0x64, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xa0, 0x01, 0x0a, 0x04, 0x45, 0x74,
	0x63, 0x64, 0x12, 0x4b, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x17, 0x2e, 0x65, 0x74, 0x63, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x74, 0x63,
	0x64, 0x50, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x65, 0x74, 0x63, 0x64, 0x2f, 0x70, 0x75, 0x74, 0x12,
	0x4b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x74, 0x63, 0x64, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x65, 0x74, 0x63, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01,
	0x2a, 0x22, 0x09, 0x2f, 0x65, 0x74, 0x63, 0x64, 0x2f, 0x67, 0x65, 0x74, 0x42, 0x2f, 0x5a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x4c, 0x53, 0x4d, 0x44,
	0x42, 0x2f, 0x6c, 0x73, 0x6d, 0x64, 0x62, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x74, 0x63, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_etcd_v1_etcd_proto_rawDescOnce sync.Once
	file_api_etcd_v1_etcd_proto_rawDescData = file_api_etcd_v1_etcd_proto_rawDesc
)

func file_api_etcd_v1_etcd_proto_rawDescGZIP() []byte {
	file_api_etcd_v1_etcd_proto_rawDescOnce.Do(func() {
		file_api_etcd_v1_etcd_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_etcd_v1_etcd_proto_rawDescData)
	})
	return file_api_etcd_v1_etcd_proto_rawDescData
}

var file_api_etcd_v1_etcd_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_etcd_v1_etcd_proto_goTypes = []interface{}{
	(*EtcdPutRequest)(nil), // 0: etcd.v1.EtcdPutRequest
	(*EtcdPutReply)(nil),   // 1: etcd.v1.EtcdPutReply
	(*EtcdGetRequest)(nil), // 2: etcd.v1.EtcdGetRequest
	(*EtcdGetReply)(nil),   // 3: etcd.v1.EtcdGetReply
}
var file_api_etcd_v1_etcd_proto_depIdxs = []int32{
	0, // 0: etcd.v1.Etcd.Put:input_type -> etcd.v1.EtcdPutRequest
	2, // 1: etcd.v1.Etcd.Get:input_type -> etcd.v1.EtcdGetRequest
	1, // 2: etcd.v1.Etcd.Put:output_type -> etcd.v1.EtcdPutReply
	3, // 3: etcd.v1.Etcd.Get:output_type -> etcd.v1.EtcdGetReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_etcd_v1_etcd_proto_init() }
func file_api_etcd_v1_etcd_proto_init() {
	if File_api_etcd_v1_etcd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_etcd_v1_etcd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EtcdPutRequest); i {
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
		file_api_etcd_v1_etcd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EtcdPutReply); i {
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
		file_api_etcd_v1_etcd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EtcdGetRequest); i {
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
		file_api_etcd_v1_etcd_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EtcdGetReply); i {
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
			RawDescriptor: file_api_etcd_v1_etcd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_etcd_v1_etcd_proto_goTypes,
		DependencyIndexes: file_api_etcd_v1_etcd_proto_depIdxs,
		MessageInfos:      file_api_etcd_v1_etcd_proto_msgTypes,
	}.Build()
	File_api_etcd_v1_etcd_proto = out.File
	file_api_etcd_v1_etcd_proto_rawDesc = nil
	file_api_etcd_v1_etcd_proto_goTypes = nil
	file_api_etcd_v1_etcd_proto_depIdxs = nil
}
