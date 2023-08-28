// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: standards/open/v1/fieldOptionsExtentions.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FDXOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question    string `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Number      string `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	Conditional string `protobuf:"bytes,4,opt,name=conditional,proto3" json:"conditional,omitempty"`
	H1          string `protobuf:"bytes,5,opt,name=h1,proto3" json:"h1,omitempty"`
	H2          string `protobuf:"bytes,6,opt,name=h2,proto3" json:"h2,omitempty"`
	H3          string `protobuf:"bytes,7,opt,name=h3,proto3" json:"h3,omitempty"`
	H4          string `protobuf:"bytes,10,opt,name=h4,proto3" json:"h4,omitempty"`
	Validation  string `protobuf:"bytes,8,opt,name=validation,proto3" json:"validation,omitempty"`
	Default     string `protobuf:"bytes,9,opt,name=default,proto3" json:"default,omitempty"`
	Required    bool   `protobuf:"varint,11,opt,name=required,proto3" json:"required,omitempty"`
	JsonName    string `protobuf:"bytes,12,opt,name=jsonName,proto3" json:"jsonName,omitempty"`
}

func (x *FDXOptions) Reset() {
	*x = FDXOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_standards_open_v1_fieldOptionsExtentions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FDXOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FDXOptions) ProtoMessage() {}

func (x *FDXOptions) ProtoReflect() protoreflect.Message {
	mi := &file_standards_open_v1_fieldOptionsExtentions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FDXOptions.ProtoReflect.Descriptor instead.
func (*FDXOptions) Descriptor() ([]byte, []int) {
	return file_standards_open_v1_fieldOptionsExtentions_proto_rawDescGZIP(), []int{0}
}

func (x *FDXOptions) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *FDXOptions) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FDXOptions) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *FDXOptions) GetConditional() string {
	if x != nil {
		return x.Conditional
	}
	return ""
}

func (x *FDXOptions) GetH1() string {
	if x != nil {
		return x.H1
	}
	return ""
}

func (x *FDXOptions) GetH2() string {
	if x != nil {
		return x.H2
	}
	return ""
}

func (x *FDXOptions) GetH3() string {
	if x != nil {
		return x.H3
	}
	return ""
}

func (x *FDXOptions) GetH4() string {
	if x != nil {
		return x.H4
	}
	return ""
}

func (x *FDXOptions) GetValidation() string {
	if x != nil {
		return x.Validation
	}
	return ""
}

func (x *FDXOptions) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *FDXOptions) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *FDXOptions) GetJsonName() string {
	if x != nil {
		return x.JsonName
	}
	return ""
}

var file_standards_open_v1_fieldOptionsExtentions_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FDXOptions)(nil),
		Field:         70001,
		Name:          "standards.open.v1.fdx_options",
		Tag:           "bytes,70001,opt,name=fdx_options",
		Filename:      "standards/open/v1/fieldOptionsExtentions.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional standards.open.v1.FDXOptions fdx_options = 70001;
	E_FdxOptions = &file_standards_open_v1_fieldOptionsExtentions_proto_extTypes[0]
)

var File_standards_open_v1_fieldOptionsExtentions_proto protoreflect.FileDescriptor

var file_standards_open_v1_fieldOptionsExtentions_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x02, 0x0a, 0x0a, 0x46, 0x44, 0x58, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02,
	0x68, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x68, 0x31, 0x12, 0x0e, 0x0a, 0x02,
	0x68, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x68, 0x32, 0x12, 0x0e, 0x0a, 0x02,
	0x68, 0x33, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x68, 0x33, 0x12, 0x0e, 0x0a, 0x02,
	0x68, 0x34, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x68, 0x34, 0x12, 0x1e, 0x0a, 0x0a,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x5f,
	0x0a, 0x0b, 0x66, 0x64, 0x78, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf1, 0xa2, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x44, 0x58, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x66, 0x64, 0x78, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2d, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_standards_open_v1_fieldOptionsExtentions_proto_rawDescOnce sync.Once
	file_standards_open_v1_fieldOptionsExtentions_proto_rawDescData = file_standards_open_v1_fieldOptionsExtentions_proto_rawDesc
)

func file_standards_open_v1_fieldOptionsExtentions_proto_rawDescGZIP() []byte {
	file_standards_open_v1_fieldOptionsExtentions_proto_rawDescOnce.Do(func() {
		file_standards_open_v1_fieldOptionsExtentions_proto_rawDescData = protoimpl.X.CompressGZIP(file_standards_open_v1_fieldOptionsExtentions_proto_rawDescData)
	})
	return file_standards_open_v1_fieldOptionsExtentions_proto_rawDescData
}

var file_standards_open_v1_fieldOptionsExtentions_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_standards_open_v1_fieldOptionsExtentions_proto_goTypes = []interface{}{
	(*FDXOptions)(nil),                // 0: standards.open.v1.FDXOptions
	(*descriptorpb.FieldOptions)(nil), // 1: google.protobuf.FieldOptions
}
var file_standards_open_v1_fieldOptionsExtentions_proto_depIdxs = []int32{
	1, // 0: standards.open.v1.fdx_options:extendee -> google.protobuf.FieldOptions
	0, // 1: standards.open.v1.fdx_options:type_name -> standards.open.v1.FDXOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_standards_open_v1_fieldOptionsExtentions_proto_init() }
func file_standards_open_v1_fieldOptionsExtentions_proto_init() {
	if File_standards_open_v1_fieldOptionsExtentions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_standards_open_v1_fieldOptionsExtentions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FDXOptions); i {
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
			RawDescriptor: file_standards_open_v1_fieldOptionsExtentions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_standards_open_v1_fieldOptionsExtentions_proto_goTypes,
		DependencyIndexes: file_standards_open_v1_fieldOptionsExtentions_proto_depIdxs,
		MessageInfos:      file_standards_open_v1_fieldOptionsExtentions_proto_msgTypes,
		ExtensionInfos:    file_standards_open_v1_fieldOptionsExtentions_proto_extTypes,
	}.Build()
	File_standards_open_v1_fieldOptionsExtentions_proto = out.File
	file_standards_open_v1_fieldOptionsExtentions_proto_rawDesc = nil
	file_standards_open_v1_fieldOptionsExtentions_proto_goTypes = nil
	file_standards_open_v1_fieldOptionsExtentions_proto_depIdxs = nil
}
