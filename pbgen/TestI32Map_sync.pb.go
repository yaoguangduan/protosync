// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: TestI32Map_sync.proto

package pbgen

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

type TestI32Map struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *int32  `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Addition *string `protobuf:"bytes,2,opt,name=addition,proto3,oneof" json:"addition,omitempty"`
}

func (x *TestI32Map) Reset() {
	*x = TestI32Map{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TestI32Map_sync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestI32Map) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestI32Map) ProtoMessage() {}

func (x *TestI32Map) ProtoReflect() protoreflect.Message {
	mi := &file_TestI32Map_sync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestI32Map.ProtoReflect.Descriptor instead.
func (*TestI32Map) Descriptor() ([]byte, []int) {
	return file_TestI32Map_sync_proto_rawDescGZIP(), []int{0}
}

func (x *TestI32Map) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *TestI32Map) GetAddition() string {
	if x != nil && x.Addition != nil {
		return *x.Addition
	}
	return ""
}

var File_TestI32Map_sync_proto protoreflect.FileDescriptor

var file_TestI32Map_sync_proto_rawDesc = []byte{
	0x0a, 0x15, 0x54, 0x65, 0x73, 0x74, 0x49, 0x33, 0x32, 0x4d, 0x61, 0x70, 0x5f, 0x73, 0x79, 0x6e,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x49,
	0x33, 0x32, 0x4d, 0x61, 0x70, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x61, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f,
	0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_TestI32Map_sync_proto_rawDescOnce sync.Once
	file_TestI32Map_sync_proto_rawDescData = file_TestI32Map_sync_proto_rawDesc
)

func file_TestI32Map_sync_proto_rawDescGZIP() []byte {
	file_TestI32Map_sync_proto_rawDescOnce.Do(func() {
		file_TestI32Map_sync_proto_rawDescData = protoimpl.X.CompressGZIP(file_TestI32Map_sync_proto_rawDescData)
	})
	return file_TestI32Map_sync_proto_rawDescData
}

var file_TestI32Map_sync_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TestI32Map_sync_proto_goTypes = []any{
	(*TestI32Map)(nil), // 0: TestI32Map
}
var file_TestI32Map_sync_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TestI32Map_sync_proto_init() }
func file_TestI32Map_sync_proto_init() {
	if File_TestI32Map_sync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TestI32Map_sync_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TestI32Map); i {
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
	file_TestI32Map_sync_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_TestI32Map_sync_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TestI32Map_sync_proto_goTypes,
		DependencyIndexes: file_TestI32Map_sync_proto_depIdxs,
		MessageInfos:      file_TestI32Map_sync_proto_msgTypes,
	}.Build()
	File_TestI32Map_sync_proto = out.File
	file_TestI32Map_sync_proto_rawDesc = nil
	file_TestI32Map_sync_proto_goTypes = nil
	file_TestI32Map_sync_proto_depIdxs = nil
}
