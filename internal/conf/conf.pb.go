// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v5.26.0
// source: internal/conf/conf.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data   *Data   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http *Server_HTTP `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc *Server_GRPC `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{1}
}

func (x *Server) GetHttp() *Server_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *Data_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Redis    *Data_Redis    `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
	Etcd     *Data_Etcd     `protobuf:"bytes,3,opt,name=etcd,proto3" json:"etcd,omitempty"`
	Oss      *Data_OSS      `protobuf:"bytes,4,opt,name=oss,proto3" json:"oss,omitempty"`
	Leveldb  *Data_Leveldb  `protobuf:"bytes,5,opt,name=leveldb,proto3" json:"leveldb,omitempty"`
	NodeName string         `protobuf:"bytes,6,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"` // 节点号，如node1
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetDatabase() *Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Data) GetRedis() *Data_Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

func (x *Data) GetEtcd() *Data_Etcd {
	if x != nil {
		return x.Etcd
	}
	return nil
}

func (x *Data) GetOss() *Data_OSS {
	if x != nil {
		return x.Oss
	}
	return nil
}

func (x *Data) GetLeveldb() *Data_Leveldb {
	if x != nil {
		return x.Leveldb
	}
	return nil
}

func (x *Data) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

type Server_HTTP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Server_HTTP) Reset() {
	*x = Server_HTTP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_HTTP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_HTTP) ProtoMessage() {}

func (x *Server_HTTP) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_HTTP.ProtoReflect.Descriptor instead.
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Server_HTTP) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_HTTP) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_HTTP) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Server_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Server_GRPC) Reset() {
	*x = Server_GRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_GRPC) ProtoMessage() {}

func (x *Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_GRPC.ProtoReflect.Descriptor instead.
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Data_Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *Data_Database) Reset() {
	*x = Data_Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Database) ProtoMessage() {}

func (x *Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Database.ProtoReflect.Descriptor instead.
func (*Data_Database) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Data_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr         string               `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Password     string               `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Db           uint32               `protobuf:"varint,3,opt,name=db,proto3" json:"db,omitempty"`
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,4,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout *durationpb.Duration `protobuf:"bytes,5,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{2, 1}
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Data_Redis) GetDb() uint32 {
	if x != nil {
		return x.Db
	}
	return 0
}

func (x *Data_Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

type Data_Etcd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterAddrs []string `protobuf:"bytes,1,rep,name=cluster_addrs,json=clusterAddrs,proto3" json:"cluster_addrs,omitempty"`
	UserName     string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Password     string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Data_Etcd) Reset() {
	*x = Data_Etcd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Etcd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Etcd) ProtoMessage() {}

func (x *Data_Etcd) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Etcd.ProtoReflect.Descriptor instead.
func (*Data_Etcd) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{2, 2}
}

func (x *Data_Etcd) GetClusterAddrs() []string {
	if x != nil {
		return x.ClusterAddrs
	}
	return nil
}

func (x *Data_Etcd) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Data_Etcd) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Data_OSS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint  string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"` // s3 addr
	AccessKey string `protobuf:"bytes,2,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey string `protobuf:"bytes,3,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Region    string `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *Data_OSS) Reset() {
	*x = Data_OSS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_OSS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_OSS) ProtoMessage() {}

func (x *Data_OSS) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_OSS.ProtoReflect.Descriptor instead.
func (*Data_OSS) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{2, 3}
}

func (x *Data_OSS) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Data_OSS) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *Data_OSS) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *Data_OSS) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type Data_Leveldb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataDir              string `protobuf:"bytes,1,opt,name=data_dir,json=dataDir,proto3" json:"data_dir,omitempty"`        // data dir
	CacheDir             string `protobuf:"bytes,2,opt,name=cache_dir,json=cacheDir,proto3" json:"cache_dir,omitempty"`     // cache dir
	BlockSize            uint64 `protobuf:"varint,3,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"` // size of leveldb block
	IteratorSamplingRate uint64 `protobuf:"varint,4,opt,name=iterator_sampling_rate,json=iteratorSamplingRate,proto3" json:"iterator_sampling_rate,omitempty"`
	WriteBuffer          uint64 `protobuf:"varint,5,opt,name=write_buffer,json=writeBuffer,proto3" json:"write_buffer,omitempty"`
	CompactionTableTize  uint64 `protobuf:"varint,6,opt,name=compaction_table_tize,json=compactionTableTize,proto3" json:"compaction_table_tize,omitempty"`
}

func (x *Data_Leveldb) Reset() {
	*x = Data_Leveldb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_conf_conf_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Leveldb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Leveldb) ProtoMessage() {}

func (x *Data_Leveldb) ProtoReflect() protoreflect.Message {
	mi := &file_internal_conf_conf_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Leveldb.ProtoReflect.Descriptor instead.
func (*Data_Leveldb) Descriptor() ([]byte, []int) {
	return file_internal_conf_conf_proto_rawDescGZIP(), []int{2, 4}
}

func (x *Data_Leveldb) GetDataDir() string {
	if x != nil {
		return x.DataDir
	}
	return ""
}

func (x *Data_Leveldb) GetCacheDir() string {
	if x != nil {
		return x.CacheDir
	}
	return ""
}

func (x *Data_Leveldb) GetBlockSize() uint64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

func (x *Data_Leveldb) GetIteratorSamplingRate() uint64 {
	if x != nil {
		return x.IteratorSamplingRate
	}
	return 0
}

func (x *Data_Leveldb) GetWriteBuffer() uint64 {
	if x != nil {
		return x.WriteBuffer
	}
	return 0
}

func (x *Data_Leveldb) GetCompactionTableTize() uint64 {
	if x != nil {
		return x.CompactionTableTize
	}
	return 0
}

var File_internal_conf_conf_proto protoreflect.FileDescriptor

var file_internal_conf_conf_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb8, 0x02, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x2b, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x2b, 0x0a,
	0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x69, 0x0a, 0x04, 0x48, 0x54,
	0x54, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x69, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43, 0x12, 0x18, 0x0a,
	0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a, 0x07, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x22, 0xe2, 0x07, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x12, 0x29,
	0x0a, 0x04, 0x65, 0x74, 0x63, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x45,
	0x74, 0x63, 0x64, 0x52, 0x04, 0x65, 0x74, 0x63, 0x64, 0x12, 0x26, 0x0a, 0x03, 0x6f, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4f, 0x53, 0x53, 0x52, 0x03, 0x6f, 0x73,
	0x73, 0x12, 0x32, 0x0a, 0x07, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x64, 0x62, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x64, 0x62, 0x52, 0x07, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x64, 0x62, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x1a, 0x3a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0xc5,
	0x01, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x64, 0x62, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x1a, 0x64, 0x0a, 0x04, 0x45, 0x74, 0x63, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x77, 0x0a, 0x03,
	0x4f, 0x53, 0x53, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x1a, 0xed, 0x01, 0x0a, 0x07, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x64,
	0x62, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x44, 0x69, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x63, 0x68, 0x65, 0x44, 0x69, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x69, 0x74, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x77, 0x72, 0x69, 0x74, 0x65, 0x42, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x54, 0x69, 0x7a, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x6c, 0x73, 0x6d, 0x64, 0x62, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_conf_conf_proto_rawDescOnce sync.Once
	file_internal_conf_conf_proto_rawDescData = file_internal_conf_conf_proto_rawDesc
)

func file_internal_conf_conf_proto_rawDescGZIP() []byte {
	file_internal_conf_conf_proto_rawDescOnce.Do(func() {
		file_internal_conf_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_conf_conf_proto_rawDescData)
	})
	return file_internal_conf_conf_proto_rawDescData
}

var file_internal_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_internal_conf_conf_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),           // 0: kratos.api.Bootstrap
	(*Server)(nil),              // 1: kratos.api.Server
	(*Data)(nil),                // 2: kratos.api.Data
	(*Server_HTTP)(nil),         // 3: kratos.api.Server.HTTP
	(*Server_GRPC)(nil),         // 4: kratos.api.Server.GRPC
	(*Data_Database)(nil),       // 5: kratos.api.Data.Database
	(*Data_Redis)(nil),          // 6: kratos.api.Data.Redis
	(*Data_Etcd)(nil),           // 7: kratos.api.Data.Etcd
	(*Data_OSS)(nil),            // 8: kratos.api.Data.OSS
	(*Data_Leveldb)(nil),        // 9: kratos.api.Data.Leveldb
	(*durationpb.Duration)(nil), // 10: google.protobuf.Duration
}
var file_internal_conf_conf_proto_depIdxs = []int32{
	1,  // 0: kratos.api.Bootstrap.server:type_name -> kratos.api.Server
	2,  // 1: kratos.api.Bootstrap.data:type_name -> kratos.api.Data
	3,  // 2: kratos.api.Server.http:type_name -> kratos.api.Server.HTTP
	4,  // 3: kratos.api.Server.grpc:type_name -> kratos.api.Server.GRPC
	5,  // 4: kratos.api.Data.database:type_name -> kratos.api.Data.Database
	6,  // 5: kratos.api.Data.redis:type_name -> kratos.api.Data.Redis
	7,  // 6: kratos.api.Data.etcd:type_name -> kratos.api.Data.Etcd
	8,  // 7: kratos.api.Data.oss:type_name -> kratos.api.Data.OSS
	9,  // 8: kratos.api.Data.leveldb:type_name -> kratos.api.Data.Leveldb
	10, // 9: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
	10, // 10: kratos.api.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	10, // 11: kratos.api.Data.Redis.read_timeout:type_name -> google.protobuf.Duration
	10, // 12: kratos.api.Data.Redis.write_timeout:type_name -> google.protobuf.Duration
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_internal_conf_conf_proto_init() }
func file_internal_conf_conf_proto_init() {
	if File_internal_conf_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_conf_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
		file_internal_conf_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_internal_conf_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_internal_conf_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_HTTP); i {
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
		file_internal_conf_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_GRPC); i {
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
		file_internal_conf_conf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Database); i {
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
		file_internal_conf_conf_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Redis); i {
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
		file_internal_conf_conf_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Etcd); i {
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
		file_internal_conf_conf_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_OSS); i {
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
		file_internal_conf_conf_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Leveldb); i {
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
			RawDescriptor: file_internal_conf_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_conf_conf_proto_goTypes,
		DependencyIndexes: file_internal_conf_conf_proto_depIdxs,
		MessageInfos:      file_internal_conf_conf_proto_msgTypes,
	}.Build()
	File_internal_conf_conf_proto = out.File
	file_internal_conf_conf_proto_rawDesc = nil
	file_internal_conf_conf_proto_goTypes = nil
	file_internal_conf_conf_proto_depIdxs = nil
}
