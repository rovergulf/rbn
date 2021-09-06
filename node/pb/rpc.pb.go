// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: node/pb/rpc.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Command int32

const (
	Command_NoOp   Command = 0
	Command_Sync   Command = 1
	Command_Add    Command = 2
	Command_Update Command = 3
	Command_Get    Command = 4
	Command_List   Command = 5
	Command_Verify Command = 6
	Command_Drop   Command = 7
)

// Enum value maps for Command.
var (
	Command_name = map[int32]string{
		0: "NoOp",
		1: "Sync",
		2: "Add",
		3: "Update",
		4: "Get",
		5: "List",
		6: "Verify",
		7: "Drop",
	}
	Command_value = map[string]int32{
		"NoOp":   0,
		"Sync":   1,
		"Add":    2,
		"Update": 3,
		"Get":    4,
		"List":   5,
		"Verify": 6,
		"Drop":   7,
	}
)

func (x Command) Enum() *Command {
	p := new(Command)
	*p = x
	return p
}

func (x Command) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Command) Descriptor() protoreflect.EnumDescriptor {
	return file_node_pb_rpc_proto_enumTypes[0].Descriptor()
}

func (Command) Type() protoreflect.EnumType {
	return &file_node_pb_rpc_proto_enumTypes[0]
}

func (x Command) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Command.Descriptor instead.
func (Command) EnumDescriptor() ([]byte, []int) {
	return file_node_pb_rpc_proto_rawDescGZIP(), []int{0}
}

type Entity int32

const (
	Entity_Unknown     Entity = 0
	Entity_Genesis     Entity = 1
	Entity_Balance     Entity = 2
	Entity_Transaction Entity = 3
	Entity_Block       Entity = 4
	Entity_BlockHeader Entity = 5
	Entity_State       Entity = 6
	Entity_Peer        Entity = 7
)

// Enum value maps for Entity.
var (
	Entity_name = map[int32]string{
		0: "Unknown",
		1: "Genesis",
		2: "Balance",
		3: "Transaction",
		4: "Block",
		5: "BlockHeader",
		6: "State",
		7: "Peer",
	}
	Entity_value = map[string]int32{
		"Unknown":     0,
		"Genesis":     1,
		"Balance":     2,
		"Transaction": 3,
		"Block":       4,
		"BlockHeader": 5,
		"State":       6,
		"Peer":        7,
	}
)

func (x Entity) Enum() *Entity {
	p := new(Entity)
	*p = x
	return p
}

func (x Entity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Entity) Descriptor() protoreflect.EnumDescriptor {
	return file_node_pb_rpc_proto_enumTypes[1].Descriptor()
}

func (Entity) Type() protoreflect.EnumType {
	return &file_node_pb_rpc_proto_enumTypes[1]
}

func (x Entity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Entity.Descriptor instead.
func (Entity) EnumDescriptor() ([]byte, []int) {
	return file_node_pb_rpc_proto_rawDescGZIP(), []int{1}
}

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd       Command `protobuf:"varint,1,opt,name=cmd,proto3,enum=rovergulf.proto.rbn.v1.Command" json:"cmd,omitempty"`
	Entity    Entity  `protobuf:"varint,2,opt,name=entity,proto3,enum=rovergulf.proto.rbn.v1.Entity" json:"entity,omitempty"`
	Data      []byte  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp int64   `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_node_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *CallRequest) GetCmd() Command {
	if x != nil {
		return x.Cmd
	}
	return Command_NoOp
}

func (x *CallRequest) GetEntity() Entity {
	if x != nil {
		return x.Entity
	}
	return Entity_Unknown
}

func (x *CallRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CallRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_node_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *CallResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CallResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Ip      string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port    int32  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	NodeId  string `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_node_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_node_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *Node) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Node) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Node) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Node) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

var File_node_pb_rpc_proto protoreflect.FileDescriptor

var file_node_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x67, 0x75, 0x6c, 0x66, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x62, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xaa, 0x01, 0x0a, 0x0b,
	0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x03, 0x63,
	0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x72, 0x6f, 0x76, 0x65, 0x72,
	0x67, 0x75, 0x6c, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x62, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x36,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e,
	0x2e, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x67, 0x75, 0x6c, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x72, 0x62, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x3a, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x5d, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x2a, 0x5b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x6f, 0x4f, 0x70, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x10, 0x04,
	0x12, 0x08, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x72, 0x6f, 0x70, 0x10, 0x07,
	0x2a, 0x71, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x65, 0x6e, 0x65, 0x73,
	0x69, 0x73, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x10,
	0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x10, 0x04, 0x12, 0x0f, 0x0a,
	0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x10, 0x05, 0x12, 0x09,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x65, 0x65,
	0x72, 0x10, 0x07, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x6f, 0x76, 0x65, 0x72, 0x67, 0x75, 0x6c, 0x66, 0x2f, 0x72, 0x62, 0x6e, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_pb_rpc_proto_rawDescOnce sync.Once
	file_node_pb_rpc_proto_rawDescData = file_node_pb_rpc_proto_rawDesc
)

func file_node_pb_rpc_proto_rawDescGZIP() []byte {
	file_node_pb_rpc_proto_rawDescOnce.Do(func() {
		file_node_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_pb_rpc_proto_rawDescData)
	})
	return file_node_pb_rpc_proto_rawDescData
}

var file_node_pb_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_node_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_node_pb_rpc_proto_goTypes = []interface{}{
	(Command)(0),         // 0: rovergulf.proto.rbn.v1.Command
	(Entity)(0),          // 1: rovergulf.proto.rbn.v1.Entity
	(*CallRequest)(nil),  // 2: rovergulf.proto.rbn.v1.CallRequest
	(*CallResponse)(nil), // 3: rovergulf.proto.rbn.v1.CallResponse
	(*Node)(nil),         // 4: rovergulf.proto.rbn.v1.Node
}
var file_node_pb_rpc_proto_depIdxs = []int32{
	0, // 0: rovergulf.proto.rbn.v1.CallRequest.cmd:type_name -> rovergulf.proto.rbn.v1.Command
	1, // 1: rovergulf.proto.rbn.v1.CallRequest.entity:type_name -> rovergulf.proto.rbn.v1.Entity
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_node_pb_rpc_proto_init() }
func file_node_pb_rpc_proto_init() {
	if File_node_pb_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_node_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
		file_node_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
			RawDescriptor: file_node_pb_rpc_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_node_pb_rpc_proto_goTypes,
		DependencyIndexes: file_node_pb_rpc_proto_depIdxs,
		EnumInfos:         file_node_pb_rpc_proto_enumTypes,
		MessageInfos:      file_node_pb_rpc_proto_msgTypes,
	}.Build()
	File_node_pb_rpc_proto = out.File
	file_node_pb_rpc_proto_rawDesc = nil
	file_node_pb_rpc_proto_goTypes = nil
	file_node_pb_rpc_proto_depIdxs = nil
}