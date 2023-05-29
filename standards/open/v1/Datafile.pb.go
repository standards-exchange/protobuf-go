// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: standards/open/v1/Datafile.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// FieldStatus emum
type DataFile_FieldAnnotation_FieldState int32

const (
	// FIELD_STATE_UNSPECIFIED
	DataFile_FieldAnnotation_FIELD_STATE_UNSPECIFIED DataFile_FieldAnnotation_FieldState = 0
	// COMPLETED
	DataFile_FieldAnnotation_COMPLETED DataFile_FieldAnnotation_FieldState = 1
	// IN_PROGRESS
	DataFile_FieldAnnotation_IN_PROGRESS DataFile_FieldAnnotation_FieldState = 2
	// PENDING_APPROVAL
	DataFile_FieldAnnotation_PENDING_APPROVAL DataFile_FieldAnnotation_FieldState = 3
	// APPROVED
	DataFile_FieldAnnotation_APPROVED DataFile_FieldAnnotation_FieldState = 4
)

// Enum value maps for DataFile_FieldAnnotation_FieldState.
var (
	DataFile_FieldAnnotation_FieldState_name = map[int32]string{
		0: "FIELD_STATE_UNSPECIFIED",
		1: "COMPLETED",
		2: "IN_PROGRESS",
		3: "PENDING_APPROVAL",
		4: "APPROVED",
	}
	DataFile_FieldAnnotation_FieldState_value = map[string]int32{
		"FIELD_STATE_UNSPECIFIED": 0,
		"COMPLETED":               1,
		"IN_PROGRESS":             2,
		"PENDING_APPROVAL":        3,
		"APPROVED":                4,
	}
)

func (x DataFile_FieldAnnotation_FieldState) Enum() *DataFile_FieldAnnotation_FieldState {
	p := new(DataFile_FieldAnnotation_FieldState)
	*p = x
	return p
}

func (x DataFile_FieldAnnotation_FieldState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataFile_FieldAnnotation_FieldState) Descriptor() protoreflect.EnumDescriptor {
	return file_standards_open_v1_Datafile_proto_enumTypes[0].Descriptor()
}

func (DataFile_FieldAnnotation_FieldState) Type() protoreflect.EnumType {
	return &file_standards_open_v1_Datafile_proto_enumTypes[0]
}

func (x DataFile_FieldAnnotation_FieldState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataFile_FieldAnnotation_FieldState.Descriptor instead.
func (DataFile_FieldAnnotation_FieldState) EnumDescriptor() ([]byte, []int) {
	return file_standards_open_v1_Datafile_proto_rawDescGZIP(), []int{0, 3, 0}
}

// Message defining an data file resource.
type DataFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Fully qualified resource name
	// Format: datafiles/{name}
	// Example: datafiles/1234ABC
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The display name of the data room file
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Created time of the reportTemplate.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Update time of the reportTemplate.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Data standard type of the data file's marshalled data
	Type string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	// The data of the file
	Data *anypb.Any `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// Field Annotations
	FieldAnnotations map[string]*DataFile_FieldAnnotation `protobuf:"bytes,7,rep,name=field_annotations,json=fieldAnnotations,proto3" json:"field_annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Attachments
	FileAttachments map[string]*DataFile_FileAttachment `protobuf:"bytes,8,rep,name=file_attachments,json=fileAttachments,proto3" json:"file_attachments,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// DataRoom Reference
	DataroomReference string `protobuf:"bytes,9,opt,name=dataroom_reference,json=dataroomReference,proto3" json:"dataroom_reference,omitempty"`
	// The reference to the user who is owner of the dataRoomFile.
	Owner string `protobuf:"bytes,10,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *DataFile) Reset() {
	*x = DataFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_standards_open_v1_Datafile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataFile) ProtoMessage() {}

func (x *DataFile) ProtoReflect() protoreflect.Message {
	mi := &file_standards_open_v1_Datafile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataFile.ProtoReflect.Descriptor instead.
func (*DataFile) Descriptor() ([]byte, []int) {
	return file_standards_open_v1_Datafile_proto_rawDescGZIP(), []int{0}
}

func (x *DataFile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataFile) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *DataFile) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *DataFile) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *DataFile) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DataFile) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DataFile) GetFieldAnnotations() map[string]*DataFile_FieldAnnotation {
	if x != nil {
		return x.FieldAnnotations
	}
	return nil
}

func (x *DataFile) GetFileAttachments() map[string]*DataFile_FileAttachment {
	if x != nil {
		return x.FileAttachments
	}
	return nil
}

func (x *DataFile) GetDataroomReference() string {
	if x != nil {
		return x.DataroomReference
	}
	return ""
}

func (x *DataFile) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

// File Attachment message definition
type DataFile_FileAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the supporting file
	// Format: fileAttachments/{fileAttachment}
	//
	// The supportingFile identifier has a one-to-one
	// mapping to a file in GCS. On creation, a check
	// is done to see whether the file exists and
	// a FAILED_PRECONDITION error is returned in the
	// case that it does not exist.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// The name of the file
	Filename string `protobuf:"bytes,6,opt,name=filename,proto3" json:"filename,omitempty"`
	// Download URI for the file.
	//
	// The URI is a [self-signed URI](https://cloud.google.com/storage/docs/access-control/signed-urls)
	// which expires after creation.
	//
	// Suggested usage is to immediately access the file on receiving the link,
	// as the link will not remain valid after the expiry time.
	DownloadUri string `protobuf:"bytes,7,opt,name=download_uri,json=downloadUri,proto3" json:"download_uri,omitempty"`
	// The email address of the entity that uploaded the file.
	Uploader string `protobuf:"bytes,8,opt,name=uploader,proto3" json:"uploader,omitempty"`
	// The time the file was uploaded.
	UploadTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=upload_time,json=uploadTime,proto3" json:"upload_time,omitempty"` // Can extend with metadata such as upload_time, uploaded_by, etc..
}

func (x *DataFile_FileAttachment) Reset() {
	*x = DataFile_FileAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_standards_open_v1_Datafile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataFile_FileAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataFile_FileAttachment) ProtoMessage() {}

func (x *DataFile_FileAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_standards_open_v1_Datafile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataFile_FileAttachment.ProtoReflect.Descriptor instead.
func (*DataFile_FileAttachment) Descriptor() ([]byte, []int) {
	return file_standards_open_v1_Datafile_proto_rawDescGZIP(), []int{0, 2}
}

func (x *DataFile_FileAttachment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataFile_FileAttachment) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *DataFile_FileAttachment) GetDownloadUri() string {
	if x != nil {
		return x.DownloadUri
	}
	return ""
}

func (x *DataFile_FileAttachment) GetUploader() string {
	if x != nil {
		return x.Uploader
	}
	return ""
}

func (x *DataFile_FileAttachment) GetUploadTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UploadTime
	}
	return nil
}

// Field Annotation message definition
type DataFile_FieldAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reference field key in the flat object structure where the field name is preceded by all parent level object
	// names separated by a '.'.
	ReferenceKey string `protobuf:"bytes,1,opt,name=reference_key,json=referenceKey,proto3" json:"reference_key,omitempty"`
	// Person responsible for completing the field
	ResponsibleIndividuals []string `protobuf:"bytes,2,rep,name=responsible_individuals,json=responsibleIndividuals,proto3" json:"responsible_individuals,omitempty"`
	// Status of the field
	FieldStatus DataFile_FieldAnnotation_FieldState `protobuf:"varint,3,opt,name=field_status,json=fieldStatus,proto3,enum=standards.open.v1.DataFile_FieldAnnotation_FieldState" json:"field_status,omitempty"`
	// Comments on field
	Comments []string `protobuf:"bytes,4,rep,name=comments,proto3" json:"comments,omitempty"`
	// Field description
	FieldDescription string `protobuf:"bytes,5,opt,name=field_description,json=fieldDescription,proto3" json:"field_description,omitempty"`
	// Update time of field
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Field Tags
	Tags []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	// File Attachment Reference
	FileAttachments []string `protobuf:"bytes,8,rep,name=file_attachments,json=fileAttachments,proto3" json:"file_attachments,omitempty"`
}

func (x *DataFile_FieldAnnotation) Reset() {
	*x = DataFile_FieldAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_standards_open_v1_Datafile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataFile_FieldAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataFile_FieldAnnotation) ProtoMessage() {}

func (x *DataFile_FieldAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_standards_open_v1_Datafile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataFile_FieldAnnotation.ProtoReflect.Descriptor instead.
func (*DataFile_FieldAnnotation) Descriptor() ([]byte, []int) {
	return file_standards_open_v1_Datafile_proto_rawDescGZIP(), []int{0, 3}
}

func (x *DataFile_FieldAnnotation) GetReferenceKey() string {
	if x != nil {
		return x.ReferenceKey
	}
	return ""
}

func (x *DataFile_FieldAnnotation) GetResponsibleIndividuals() []string {
	if x != nil {
		return x.ResponsibleIndividuals
	}
	return nil
}

func (x *DataFile_FieldAnnotation) GetFieldStatus() DataFile_FieldAnnotation_FieldState {
	if x != nil {
		return x.FieldStatus
	}
	return DataFile_FieldAnnotation_FIELD_STATE_UNSPECIFIED
}

func (x *DataFile_FieldAnnotation) GetComments() []string {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *DataFile_FieldAnnotation) GetFieldDescription() string {
	if x != nil {
		return x.FieldDescription
	}
	return ""
}

func (x *DataFile_FieldAnnotation) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *DataFile_FieldAnnotation) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *DataFile_FieldAnnotation) GetFileAttachments() []string {
	if x != nil {
		return x.FileAttachments
	}
	return nil
}

var File_standards_open_v1_Datafile_proto protoreflect.FileDescriptor

var file_standards_open_v1_Datafile_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc4, 0x0b, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5e, 0x0a,
	0x11, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x46, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5b, 0x0a,
	0x10, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61,
	0x72, 0x64, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x46, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x61,
	0x74, 0x61, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x61, 0x74, 0x61, 0x72, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x1a,
	0x70, 0x0a, 0x15, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x41, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x6e, 0x0a, 0x14, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x40, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0xd4, 0x01, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03,
	0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x1a, 0xfe, 0x03, 0x0a, 0x0f, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x37, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x5f, 0x69, 0x6e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x16, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x49,
	0x6e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x73, 0x12, 0x59, 0x0a, 0x0c, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x36, 0x2e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x2b, 0x0a, 0x11, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12,
	0x29, 0x0a, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x6d, 0x0a, 0x0a, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x49, 0x45, 0x4c,
	0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52,
	0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x41,
	0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x04, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x73, 0x2d, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x73,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_standards_open_v1_Datafile_proto_rawDescOnce sync.Once
	file_standards_open_v1_Datafile_proto_rawDescData = file_standards_open_v1_Datafile_proto_rawDesc
)

func file_standards_open_v1_Datafile_proto_rawDescGZIP() []byte {
	file_standards_open_v1_Datafile_proto_rawDescOnce.Do(func() {
		file_standards_open_v1_Datafile_proto_rawDescData = protoimpl.X.CompressGZIP(file_standards_open_v1_Datafile_proto_rawDescData)
	})
	return file_standards_open_v1_Datafile_proto_rawDescData
}

var file_standards_open_v1_Datafile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_standards_open_v1_Datafile_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_standards_open_v1_Datafile_proto_goTypes = []interface{}{
	(DataFile_FieldAnnotation_FieldState)(0), // 0: standards.open.v1.DataFile.FieldAnnotation.FieldState
	(*DataFile)(nil),                         // 1: standards.open.v1.DataFile
	nil,                                      // 2: standards.open.v1.DataFile.FieldAnnotationsEntry
	nil,                                      // 3: standards.open.v1.DataFile.FileAttachmentsEntry
	(*DataFile_FileAttachment)(nil),          // 4: standards.open.v1.DataFile.FileAttachment
	(*DataFile_FieldAnnotation)(nil),         // 5: standards.open.v1.DataFile.FieldAnnotation
	(*timestamppb.Timestamp)(nil),            // 6: google.protobuf.Timestamp
	(*anypb.Any)(nil),                        // 7: google.protobuf.Any
}
var file_standards_open_v1_Datafile_proto_depIdxs = []int32{
	6,  // 0: standards.open.v1.DataFile.create_time:type_name -> google.protobuf.Timestamp
	6,  // 1: standards.open.v1.DataFile.update_time:type_name -> google.protobuf.Timestamp
	7,  // 2: standards.open.v1.DataFile.data:type_name -> google.protobuf.Any
	2,  // 3: standards.open.v1.DataFile.field_annotations:type_name -> standards.open.v1.DataFile.FieldAnnotationsEntry
	3,  // 4: standards.open.v1.DataFile.file_attachments:type_name -> standards.open.v1.DataFile.FileAttachmentsEntry
	5,  // 5: standards.open.v1.DataFile.FieldAnnotationsEntry.value:type_name -> standards.open.v1.DataFile.FieldAnnotation
	4,  // 6: standards.open.v1.DataFile.FileAttachmentsEntry.value:type_name -> standards.open.v1.DataFile.FileAttachment
	6,  // 7: standards.open.v1.DataFile.FileAttachment.upload_time:type_name -> google.protobuf.Timestamp
	0,  // 8: standards.open.v1.DataFile.FieldAnnotation.field_status:type_name -> standards.open.v1.DataFile.FieldAnnotation.FieldState
	6,  // 9: standards.open.v1.DataFile.FieldAnnotation.update_time:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_standards_open_v1_Datafile_proto_init() }
func file_standards_open_v1_Datafile_proto_init() {
	if File_standards_open_v1_Datafile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_standards_open_v1_Datafile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataFile); i {
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
		file_standards_open_v1_Datafile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataFile_FileAttachment); i {
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
		file_standards_open_v1_Datafile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataFile_FieldAnnotation); i {
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
			RawDescriptor: file_standards_open_v1_Datafile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_standards_open_v1_Datafile_proto_goTypes,
		DependencyIndexes: file_standards_open_v1_Datafile_proto_depIdxs,
		EnumInfos:         file_standards_open_v1_Datafile_proto_enumTypes,
		MessageInfos:      file_standards_open_v1_Datafile_proto_msgTypes,
	}.Build()
	File_standards_open_v1_Datafile_proto = out.File
	file_standards_open_v1_Datafile_proto_rawDesc = nil
	file_standards_open_v1_Datafile_proto_goTypes = nil
	file_standards_open_v1_Datafile_proto_depIdxs = nil
}
