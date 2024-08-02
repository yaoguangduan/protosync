// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: Person.proto

package pbgenv2

import (
	_ "github.com/yaoguangduan/datasync/syncproto"
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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age      *int32                 `protobuf:"varint,1,opt,name=age,proto3,oneof" json:"age,omitempty"`
	VipLevel *VipLevel              `protobuf:"varint,2,opt,name=vipLevel,proto3,enum=VipLevel,oneof" json:"vipLevel,omitempty"`
	Name     *string                `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Actions  map[string]*ActionInfo `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Favor    []string               `protobuf:"bytes,5,rep,name=favor,proto3" json:"favor,omitempty"`
	LoveSeq  []ColorType            `protobuf:"varint,6,rep,packed,name=loveSeq,proto3,enum=ColorType" json:"loveSeq,omitempty"`
	IsGirl   *bool                  `protobuf:"varint,7,opt,name=isGirl,proto3,oneof" json:"isGirl,omitempty"`
	Detail   *IntroDetail           `protobuf:"bytes,8,opt,name=detail,proto3,oneof" json:"detail,omitempty"`
	Data     []byte                 `protobuf:"bytes,9,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_Person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_Person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetAge() int32 {
	if x != nil && x.Age != nil {
		return *x.Age
	}
	return 0
}

func (x *Person) GetVipLevel() VipLevel {
	if x != nil && x.VipLevel != nil {
		return *x.VipLevel
	}
	return VipLevel_Level1
}

func (x *Person) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Person) GetActions() map[string]*ActionInfo {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *Person) GetFavor() []string {
	if x != nil {
		return x.Favor
	}
	return nil
}

func (x *Person) GetLoveSeq() []ColorType {
	if x != nil {
		return x.LoveSeq
	}
	return nil
}

func (x *Person) GetIsGirl() bool {
	if x != nil && x.IsGirl != nil {
		return *x.IsGirl
	}
	return false
}

func (x *Person) GetDetail() *IntroDetail {
	if x != nil {
		return x.Detail
	}
	return nil
}

func (x *Person) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_Person_proto protoreflect.FileDescriptor

var file_Person_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0e, 0x56, 0x69, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x79, 0x6e, 0x63, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x03,
	0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x2a, 0x0a, 0x08, 0x76, 0x69, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x09, 0x2e, 0x56, 0x69, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x48, 0x01, 0x52, 0x08,
	0x76, 0x69, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x07, 0x6c, 0x6f,
	0x76, 0x65, 0x53, 0x65, 0x71, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6c, 0x6f, 0x76, 0x65, 0x53, 0x65, 0x71,
	0x12, 0x1b, 0x0a, 0x06, 0x69, 0x73, 0x47, 0x69, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x03, 0x52, 0x06, 0x69, 0x73, 0x47, 0x69, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x49, 0x6e, 0x74, 0x72, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48, 0x04, 0x52, 0x06, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01,
	0x01, 0x1a, 0x47, 0x0a, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x04, 0xc0, 0xc1, 0x18, 0x01,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x61, 0x67, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x76, 0x69, 0x70,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x69, 0x73, 0x47, 0x69, 0x72, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x67, 0x65, 0x6e, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_Person_proto_rawDescOnce sync.Once
	file_Person_proto_rawDescData = file_Person_proto_rawDesc
)

func file_Person_proto_rawDescGZIP() []byte {
	file_Person_proto_rawDescOnce.Do(func() {
		file_Person_proto_rawDescData = protoimpl.X.CompressGZIP(file_Person_proto_rawDescData)
	})
	return file_Person_proto_rawDescData
}

var file_Person_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Person_proto_goTypes = []any{
	(*Person)(nil),      // 0: Person
	nil,                 // 1: Person.ActionsEntry
	(VipLevel)(0),       // 2: VipLevel
	(ColorType)(0),      // 3: ColorType
	(*IntroDetail)(nil), // 4: IntroDetail
	(*ActionInfo)(nil),  // 5: ActionInfo
}
var file_Person_proto_depIdxs = []int32{
	2, // 0: Person.vipLevel:type_name -> VipLevel
	1, // 1: Person.actions:type_name -> Person.ActionsEntry
	3, // 2: Person.loveSeq:type_name -> ColorType
	4, // 3: Person.detail:type_name -> IntroDetail
	5, // 4: Person.ActionsEntry.value:type_name -> ActionInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_Person_proto_init() }
func file_Person_proto_init() {
	if File_Person_proto != nil {
		return
	}
	file_ColorType_proto_init()
	file_IntroDetail_proto_init()
	file_VipLevel_proto_init()
	file_ActionInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Person_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Person); i {
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
	file_Person_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Person_proto_goTypes,
		DependencyIndexes: file_Person_proto_depIdxs,
		MessageInfos:      file_Person_proto_msgTypes,
	}.Build()
	File_Person_proto = out.File
	file_Person_proto_rawDesc = nil
	file_Person_proto_goTypes = nil
	file_Person_proto_depIdxs = nil
}
