// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: protosync.options.proto

package syncproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_protosync_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50200,
		Name:          "sync_gen",
		Tag:           "varint,50200,opt,name=sync_gen",
		Filename:      "protosync.options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50201,
		Name:          "sync_key",
		Tag:           "varint,50201,opt,name=sync_key",
		Filename:      "protosync.options.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional bool sync_gen = 50200;
	E_SyncGen = &file_protosync_options_proto_extTypes[0]
	// optional int32 sync_key = 50201;
	E_SyncKey = &file_protosync_options_proto_extTypes[1]
)

var File_protosync_options_proto protoreflect.FileDescriptor

var file_protosync_options_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3f, 0x0a, 0x08, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x67, 0x65, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x98, 0x88, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x47, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x3a, 0x3f, 0x0a, 0x08,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x99, 0x88, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6f, 0x67,
	0x75, 0x61, 0x6e, 0x67, 0x64, 0x75, 0x61, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x79, 0x6e,
	0x63, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_protosync_options_proto_goTypes = []any{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
}
var file_protosync_options_proto_depIdxs = []int32{
	0, // 0: sync_gen:extendee -> google.protobuf.MessageOptions
	0, // 1: sync_key:extendee -> google.protobuf.MessageOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protosync_options_proto_init() }
func file_protosync_options_proto_init() {
	if File_protosync_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protosync_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_protosync_options_proto_goTypes,
		DependencyIndexes: file_protosync_options_proto_depIdxs,
		ExtensionInfos:    file_protosync_options_proto_extTypes,
	}.Build()
	File_protosync_options_proto = out.File
	file_protosync_options_proto_rawDesc = nil
	file_protosync_options_proto_goTypes = nil
	file_protosync_options_proto_depIdxs = nil
}
