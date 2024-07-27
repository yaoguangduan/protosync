// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: VipLevel_sync.proto

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

type VipLevel int32

const (
	VipLevel_Level1 VipLevel = 0 //
	VipLevel_Level2 VipLevel = 1 //
	VipLevel_Level3 VipLevel = 2 //
)

// Enum value maps for VipLevel.
var (
	VipLevel_name = map[int32]string{
		0: "Level1",
		1: "Level2",
		2: "Level3",
	}
	VipLevel_value = map[string]int32{
		"Level1": 0,
		"Level2": 1,
		"Level3": 2,
	}
)

func (x VipLevel) Enum() *VipLevel {
	p := new(VipLevel)
	*p = x
	return p
}

func (x VipLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VipLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_VipLevel_sync_proto_enumTypes[0].Descriptor()
}

func (VipLevel) Type() protoreflect.EnumType {
	return &file_VipLevel_sync_proto_enumTypes[0]
}

func (x VipLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VipLevel.Descriptor instead.
func (VipLevel) EnumDescriptor() ([]byte, []int) {
	return file_VipLevel_sync_proto_rawDescGZIP(), []int{0}
}

var File_VipLevel_sync_proto protoreflect.FileDescriptor

var file_VipLevel_sync_proto_rawDesc = []byte{
	0x0a, 0x13, 0x56, 0x69, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x2e, 0x0a, 0x08, 0x56, 0x69, 0x70, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x31, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x32, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x33, 0x10, 0x02, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_VipLevel_sync_proto_rawDescOnce sync.Once
	file_VipLevel_sync_proto_rawDescData = file_VipLevel_sync_proto_rawDesc
)

func file_VipLevel_sync_proto_rawDescGZIP() []byte {
	file_VipLevel_sync_proto_rawDescOnce.Do(func() {
		file_VipLevel_sync_proto_rawDescData = protoimpl.X.CompressGZIP(file_VipLevel_sync_proto_rawDescData)
	})
	return file_VipLevel_sync_proto_rawDescData
}

var file_VipLevel_sync_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_VipLevel_sync_proto_goTypes = []any{
	(VipLevel)(0), // 0: VipLevel
}
var file_VipLevel_sync_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_VipLevel_sync_proto_init() }
func file_VipLevel_sync_proto_init() {
	if File_VipLevel_sync_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_VipLevel_sync_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_VipLevel_sync_proto_goTypes,
		DependencyIndexes: file_VipLevel_sync_proto_depIdxs,
		EnumInfos:         file_VipLevel_sync_proto_enumTypes,
	}.Build()
	File_VipLevel_sync_proto = out.File
	file_VipLevel_sync_proto_rawDesc = nil
	file_VipLevel_sync_proto_goTypes = nil
	file_VipLevel_sync_proto_depIdxs = nil
}
