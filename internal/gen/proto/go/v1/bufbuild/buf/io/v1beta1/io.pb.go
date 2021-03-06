// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bufbuild/buf/io/v1beta1/io.proto

package iov1beta1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FileScheme int32

const (
	FileScheme_FILE_SCHEME_UNSPECIFIED FileScheme = 0
	FileScheme_FILE_SCHEME_HTTP        FileScheme = 1
	FileScheme_FILE_SCHEME_HTTPS       FileScheme = 2
	FileScheme_FILE_SCHEME_FILE        FileScheme = 3
	FileScheme_FILE_SCHEME_STDIO       FileScheme = 4
	FileScheme_FILE_SCHEME_NULL        FileScheme = 5
)

var FileScheme_name = map[int32]string{
	0: "FILE_SCHEME_UNSPECIFIED",
	1: "FILE_SCHEME_HTTP",
	2: "FILE_SCHEME_HTTPS",
	3: "FILE_SCHEME_FILE",
	4: "FILE_SCHEME_STDIO",
	5: "FILE_SCHEME_NULL",
}

var FileScheme_value = map[string]int32{
	"FILE_SCHEME_UNSPECIFIED": 0,
	"FILE_SCHEME_HTTP":        1,
	"FILE_SCHEME_HTTPS":       2,
	"FILE_SCHEME_FILE":        3,
	"FILE_SCHEME_STDIO":       4,
	"FILE_SCHEME_NULL":        5,
}

func (x FileScheme) String() string {
	return proto.EnumName(FileScheme_name, int32(x))
}

func (FileScheme) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49d667d72c1020ea, []int{0}
}

type ImageFormat int32

const (
	ImageFormat_IMAGE_FORMAT_UNSPECIFIED ImageFormat = 0
	ImageFormat_IMAGE_FORMAT_BIN         ImageFormat = 1
	ImageFormat_IMAGE_FORMAT_BINGZ       ImageFormat = 2
	ImageFormat_IMAGE_FORMAT_JSON        ImageFormat = 3
	ImageFormat_IMAGE_FORMAT_JSONGZ      ImageFormat = 4
)

var ImageFormat_name = map[int32]string{
	0: "IMAGE_FORMAT_UNSPECIFIED",
	1: "IMAGE_FORMAT_BIN",
	2: "IMAGE_FORMAT_BINGZ",
	3: "IMAGE_FORMAT_JSON",
	4: "IMAGE_FORMAT_JSONGZ",
}

var ImageFormat_value = map[string]int32{
	"IMAGE_FORMAT_UNSPECIFIED": 0,
	"IMAGE_FORMAT_BIN":         1,
	"IMAGE_FORMAT_BINGZ":       2,
	"IMAGE_FORMAT_JSON":        3,
	"IMAGE_FORMAT_JSONGZ":      4,
}

func (x ImageFormat) String() string {
	return proto.EnumName(ImageFormat_name, int32(x))
}

func (ImageFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49d667d72c1020ea, []int{1}
}

type ArchiveFormat int32

const (
	ArchiveFormat_ARCHIVE_FORMAT_UNSPECIFIED ArchiveFormat = 0
	ArchiveFormat_ARCHIVE_FORMAT_TAR         ArchiveFormat = 1
	ArchiveFormat_ARCHIVE_FORMAT_TARGZ       ArchiveFormat = 2
)

var ArchiveFormat_name = map[int32]string{
	0: "ARCHIVE_FORMAT_UNSPECIFIED",
	1: "ARCHIVE_FORMAT_TAR",
	2: "ARCHIVE_FORMAT_TARGZ",
}

var ArchiveFormat_value = map[string]int32{
	"ARCHIVE_FORMAT_UNSPECIFIED": 0,
	"ARCHIVE_FORMAT_TAR":         1,
	"ARCHIVE_FORMAT_TARGZ":       2,
}

func (x ArchiveFormat) String() string {
	return proto.EnumName(ArchiveFormat_name, int32(x))
}

func (ArchiveFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49d667d72c1020ea, []int{2}
}

type GitScheme int32

const (
	GitScheme_GIT_SCHEME_UNSPECIFIED GitScheme = 0
	GitScheme_GIT_SCHEME_HTTP        GitScheme = 1
	GitScheme_GIT_SCHEME_HTTPS       GitScheme = 2
	GitScheme_GIT_SCHEME_FILE        GitScheme = 3
	GitScheme_GIT_SCHEME_SSH         GitScheme = 4
)

var GitScheme_name = map[int32]string{
	0: "GIT_SCHEME_UNSPECIFIED",
	1: "GIT_SCHEME_HTTP",
	2: "GIT_SCHEME_HTTPS",
	3: "GIT_SCHEME_FILE",
	4: "GIT_SCHEME_SSH",
}

var GitScheme_value = map[string]int32{
	"GIT_SCHEME_UNSPECIFIED": 0,
	"GIT_SCHEME_HTTP":        1,
	"GIT_SCHEME_HTTPS":       2,
	"GIT_SCHEME_FILE":        3,
	"GIT_SCHEME_SSH":         4,
}

func (x GitScheme) String() string {
	return proto.EnumName(GitScheme_name, int32(x))
}

func (GitScheme) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49d667d72c1020ea, []int{3}
}

type GitSubmoduleBehavior int32

const (
	GitSubmoduleBehavior_GIT_SUBMODULE_BEHAVIOR_UNSPECIFIED GitSubmoduleBehavior = 0
	GitSubmoduleBehavior_GIT_SUBMODULE_BEHAVIOR_NONE        GitSubmoduleBehavior = 1
	GitSubmoduleBehavior_GIT_SUBMODULE_BEHAVIOR_RECURSIVE   GitSubmoduleBehavior = 2
)

var GitSubmoduleBehavior_name = map[int32]string{
	0: "GIT_SUBMODULE_BEHAVIOR_UNSPECIFIED",
	1: "GIT_SUBMODULE_BEHAVIOR_NONE",
	2: "GIT_SUBMODULE_BEHAVIOR_RECURSIVE",
}

var GitSubmoduleBehavior_value = map[string]int32{
	"GIT_SUBMODULE_BEHAVIOR_UNSPECIFIED": 0,
	"GIT_SUBMODULE_BEHAVIOR_NONE":        1,
	"GIT_SUBMODULE_BEHAVIOR_RECURSIVE":   2,
}

func (x GitSubmoduleBehavior) String() string {
	return proto.EnumName(GitSubmoduleBehavior_name, int32(x))
}

func (GitSubmoduleBehavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49d667d72c1020ea, []int{4}
}

type BucketScheme int32

const (
	BucketScheme_BUCKET_SCHEME_UNSPECIFIED BucketScheme = 0
	BucketScheme_BUCKET_SCHEME_DIR         BucketScheme = 1
)

var BucketScheme_name = map[int32]string{
	0: "BUCKET_SCHEME_UNSPECIFIED",
	1: "BUCKET_SCHEME_DIR",
}

var BucketScheme_value = map[string]int32{
	"BUCKET_SCHEME_UNSPECIFIED": 0,
	"BUCKET_SCHEME_DIR":         1,
}

func (x BucketScheme) String() string {
	return proto.EnumName(BucketScheme_name, int32(x))
}

func (BucketScheme) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49d667d72c1020ea, []int{5}
}

type ImageRef struct {
	FileScheme  FileScheme  `protobuf:"varint,1,opt,name=file_scheme,json=fileScheme,proto3,enum=bufbuild.buf.io.v1beta1.FileScheme" json:"file_scheme,omitempty"`
	ImageFormat ImageFormat `protobuf:"varint,2,opt,name=image_format,json=imageFormat,proto3,enum=bufbuild.buf.io.v1beta1.ImageFormat" json:"image_format,omitempty"`
	// path is the path of the image.
	//
	// This path will be local to the system and not normalized.
	// This path will not include the scheme, if any.
	// If the type is stdio or null, this will be empty.
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageRef) Reset()         { *m = ImageRef{} }
func (m *ImageRef) String() string { return proto.CompactTextString(m) }
func (*ImageRef) ProtoMessage()    {}
func (*ImageRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_49d667d72c1020ea, []int{0}
}

func (m *ImageRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageRef.Unmarshal(m, b)
}
func (m *ImageRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageRef.Marshal(b, m, deterministic)
}
func (m *ImageRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageRef.Merge(m, src)
}
func (m *ImageRef) XXX_Size() int {
	return xxx_messageInfo_ImageRef.Size(m)
}
func (m *ImageRef) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageRef.DiscardUnknown(m)
}

var xxx_messageInfo_ImageRef proto.InternalMessageInfo

func (m *ImageRef) GetFileScheme() FileScheme {
	if m != nil {
		return m.FileScheme
	}
	return FileScheme_FILE_SCHEME_UNSPECIFIED
}

func (m *ImageRef) GetImageFormat() ImageFormat {
	if m != nil {
		return m.ImageFormat
	}
	return ImageFormat_IMAGE_FORMAT_UNSPECIFIED
}

func (m *ImageRef) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ArchiveRef struct {
	FileScheme    FileScheme    `protobuf:"varint,1,opt,name=file_scheme,json=fileScheme,proto3,enum=bufbuild.buf.io.v1beta1.FileScheme" json:"file_scheme,omitempty"`
	ArchiveFormat ArchiveFormat `protobuf:"varint,2,opt,name=archive_format,json=archiveFormat,proto3,enum=bufbuild.buf.io.v1beta1.ArchiveFormat" json:"archive_format,omitempty"`
	// path is the path of the archive.
	//
	// This path will be local to the system and not normalized.
	// This path will not include the scheme, if any.
	// If the type is stdio or devnull, this will be empty.
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	StripComponents      uint32   `protobuf:"varint,4,opt,name=strip_components,json=stripComponents,proto3" json:"strip_components,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArchiveRef) Reset()         { *m = ArchiveRef{} }
func (m *ArchiveRef) String() string { return proto.CompactTextString(m) }
func (*ArchiveRef) ProtoMessage()    {}
func (*ArchiveRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_49d667d72c1020ea, []int{1}
}

func (m *ArchiveRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArchiveRef.Unmarshal(m, b)
}
func (m *ArchiveRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArchiveRef.Marshal(b, m, deterministic)
}
func (m *ArchiveRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArchiveRef.Merge(m, src)
}
func (m *ArchiveRef) XXX_Size() int {
	return xxx_messageInfo_ArchiveRef.Size(m)
}
func (m *ArchiveRef) XXX_DiscardUnknown() {
	xxx_messageInfo_ArchiveRef.DiscardUnknown(m)
}

var xxx_messageInfo_ArchiveRef proto.InternalMessageInfo

func (m *ArchiveRef) GetFileScheme() FileScheme {
	if m != nil {
		return m.FileScheme
	}
	return FileScheme_FILE_SCHEME_UNSPECIFIED
}

func (m *ArchiveRef) GetArchiveFormat() ArchiveFormat {
	if m != nil {
		return m.ArchiveFormat
	}
	return ArchiveFormat_ARCHIVE_FORMAT_UNSPECIFIED
}

func (m *ArchiveRef) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ArchiveRef) GetStripComponents() uint32 {
	if m != nil {
		return m.StripComponents
	}
	return 0
}

type GitRef struct {
	GitScheme GitScheme `protobuf:"varint,1,opt,name=git_scheme,json=gitScheme,proto3,enum=bufbuild.buf.io.v1beta1.GitScheme" json:"git_scheme,omitempty"`
	// path is the path of the archive.
	//
	// This path will be local to the system and not normalized.
	// This path will not include the scheme, if any.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Types that are valid to be assigned to Reference:
	//	*GitRef_Branch
	//	*GitRef_Tag
	Reference            isGitRef_Reference   `protobuf_oneof:"reference"`
	GitSubmoduleBehavior GitSubmoduleBehavior `protobuf:"varint,5,opt,name=git_submodule_behavior,json=gitSubmoduleBehavior,proto3,enum=bufbuild.buf.io.v1beta1.GitSubmoduleBehavior" json:"git_submodule_behavior,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GitRef) Reset()         { *m = GitRef{} }
func (m *GitRef) String() string { return proto.CompactTextString(m) }
func (*GitRef) ProtoMessage()    {}
func (*GitRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_49d667d72c1020ea, []int{2}
}

func (m *GitRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitRef.Unmarshal(m, b)
}
func (m *GitRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitRef.Marshal(b, m, deterministic)
}
func (m *GitRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitRef.Merge(m, src)
}
func (m *GitRef) XXX_Size() int {
	return xxx_messageInfo_GitRef.Size(m)
}
func (m *GitRef) XXX_DiscardUnknown() {
	xxx_messageInfo_GitRef.DiscardUnknown(m)
}

var xxx_messageInfo_GitRef proto.InternalMessageInfo

func (m *GitRef) GetGitScheme() GitScheme {
	if m != nil {
		return m.GitScheme
	}
	return GitScheme_GIT_SCHEME_UNSPECIFIED
}

func (m *GitRef) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type isGitRef_Reference interface {
	isGitRef_Reference()
}

type GitRef_Branch struct {
	Branch string `protobuf:"bytes,3,opt,name=branch,proto3,oneof"`
}

type GitRef_Tag struct {
	Tag string `protobuf:"bytes,4,opt,name=tag,proto3,oneof"`
}

func (*GitRef_Branch) isGitRef_Reference() {}

func (*GitRef_Tag) isGitRef_Reference() {}

func (m *GitRef) GetReference() isGitRef_Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *GitRef) GetBranch() string {
	if x, ok := m.GetReference().(*GitRef_Branch); ok {
		return x.Branch
	}
	return ""
}

func (m *GitRef) GetTag() string {
	if x, ok := m.GetReference().(*GitRef_Tag); ok {
		return x.Tag
	}
	return ""
}

func (m *GitRef) GetGitSubmoduleBehavior() GitSubmoduleBehavior {
	if m != nil {
		return m.GitSubmoduleBehavior
	}
	return GitSubmoduleBehavior_GIT_SUBMODULE_BEHAVIOR_UNSPECIFIED
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GitRef) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GitRef_Branch)(nil),
		(*GitRef_Tag)(nil),
	}
}

type BucketRef struct {
	BucketScheme BucketScheme `protobuf:"varint,1,opt,name=bucket_scheme,json=bucketScheme,proto3,enum=bufbuild.buf.io.v1beta1.BucketScheme" json:"bucket_scheme,omitempty"`
	// path is the path of the bucket.
	//
	// This path will be local to the system and not normalized.
	// This path will not include the scheme, if any.
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BucketRef) Reset()         { *m = BucketRef{} }
func (m *BucketRef) String() string { return proto.CompactTextString(m) }
func (*BucketRef) ProtoMessage()    {}
func (*BucketRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_49d667d72c1020ea, []int{3}
}

func (m *BucketRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketRef.Unmarshal(m, b)
}
func (m *BucketRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketRef.Marshal(b, m, deterministic)
}
func (m *BucketRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketRef.Merge(m, src)
}
func (m *BucketRef) XXX_Size() int {
	return xxx_messageInfo_BucketRef.Size(m)
}
func (m *BucketRef) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketRef.DiscardUnknown(m)
}

var xxx_messageInfo_BucketRef proto.InternalMessageInfo

func (m *BucketRef) GetBucketScheme() BucketScheme {
	if m != nil {
		return m.BucketScheme
	}
	return BucketScheme_BUCKET_SCHEME_UNSPECIFIED
}

func (m *BucketRef) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type SourceRef struct {
	// Types that are valid to be assigned to Value:
	//	*SourceRef_ArchiveRef
	//	*SourceRef_GitRef
	//	*SourceRef_BucketRef
	Value                isSourceRef_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SourceRef) Reset()         { *m = SourceRef{} }
func (m *SourceRef) String() string { return proto.CompactTextString(m) }
func (*SourceRef) ProtoMessage()    {}
func (*SourceRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_49d667d72c1020ea, []int{4}
}

func (m *SourceRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceRef.Unmarshal(m, b)
}
func (m *SourceRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceRef.Marshal(b, m, deterministic)
}
func (m *SourceRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceRef.Merge(m, src)
}
func (m *SourceRef) XXX_Size() int {
	return xxx_messageInfo_SourceRef.Size(m)
}
func (m *SourceRef) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceRef.DiscardUnknown(m)
}

var xxx_messageInfo_SourceRef proto.InternalMessageInfo

type isSourceRef_Value interface {
	isSourceRef_Value()
}

type SourceRef_ArchiveRef struct {
	ArchiveRef *ArchiveRef `protobuf:"bytes,1,opt,name=archive_ref,json=archiveRef,proto3,oneof"`
}

type SourceRef_GitRef struct {
	GitRef *GitRef `protobuf:"bytes,2,opt,name=git_ref,json=gitRef,proto3,oneof"`
}

type SourceRef_BucketRef struct {
	BucketRef *BucketRef `protobuf:"bytes,3,opt,name=bucket_ref,json=bucketRef,proto3,oneof"`
}

func (*SourceRef_ArchiveRef) isSourceRef_Value() {}

func (*SourceRef_GitRef) isSourceRef_Value() {}

func (*SourceRef_BucketRef) isSourceRef_Value() {}

func (m *SourceRef) GetValue() isSourceRef_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SourceRef) GetArchiveRef() *ArchiveRef {
	if x, ok := m.GetValue().(*SourceRef_ArchiveRef); ok {
		return x.ArchiveRef
	}
	return nil
}

func (m *SourceRef) GetGitRef() *GitRef {
	if x, ok := m.GetValue().(*SourceRef_GitRef); ok {
		return x.GitRef
	}
	return nil
}

func (m *SourceRef) GetBucketRef() *BucketRef {
	if x, ok := m.GetValue().(*SourceRef_BucketRef); ok {
		return x.BucketRef
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SourceRef) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SourceRef_ArchiveRef)(nil),
		(*SourceRef_GitRef)(nil),
		(*SourceRef_BucketRef)(nil),
	}
}

type InputRef struct {
	// Types that are valid to be assigned to Value:
	//	*InputRef_ImageRef
	//	*InputRef_SourceRef
	Value                isInputRef_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *InputRef) Reset()         { *m = InputRef{} }
func (m *InputRef) String() string { return proto.CompactTextString(m) }
func (*InputRef) ProtoMessage()    {}
func (*InputRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_49d667d72c1020ea, []int{5}
}

func (m *InputRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputRef.Unmarshal(m, b)
}
func (m *InputRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputRef.Marshal(b, m, deterministic)
}
func (m *InputRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputRef.Merge(m, src)
}
func (m *InputRef) XXX_Size() int {
	return xxx_messageInfo_InputRef.Size(m)
}
func (m *InputRef) XXX_DiscardUnknown() {
	xxx_messageInfo_InputRef.DiscardUnknown(m)
}

var xxx_messageInfo_InputRef proto.InternalMessageInfo

type isInputRef_Value interface {
	isInputRef_Value()
}

type InputRef_ImageRef struct {
	ImageRef *ImageRef `protobuf:"bytes,1,opt,name=image_ref,json=imageRef,proto3,oneof"`
}

type InputRef_SourceRef struct {
	SourceRef *SourceRef `protobuf:"bytes,2,opt,name=source_ref,json=sourceRef,proto3,oneof"`
}

func (*InputRef_ImageRef) isInputRef_Value() {}

func (*InputRef_SourceRef) isInputRef_Value() {}

func (m *InputRef) GetValue() isInputRef_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *InputRef) GetImageRef() *ImageRef {
	if x, ok := m.GetValue().(*InputRef_ImageRef); ok {
		return x.ImageRef
	}
	return nil
}

func (m *InputRef) GetSourceRef() *SourceRef {
	if x, ok := m.GetValue().(*InputRef_SourceRef); ok {
		return x.SourceRef
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*InputRef) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*InputRef_ImageRef)(nil),
		(*InputRef_SourceRef)(nil),
	}
}

func init() {
	proto.RegisterEnum("bufbuild.buf.io.v1beta1.FileScheme", FileScheme_name, FileScheme_value)
	proto.RegisterEnum("bufbuild.buf.io.v1beta1.ImageFormat", ImageFormat_name, ImageFormat_value)
	proto.RegisterEnum("bufbuild.buf.io.v1beta1.ArchiveFormat", ArchiveFormat_name, ArchiveFormat_value)
	proto.RegisterEnum("bufbuild.buf.io.v1beta1.GitScheme", GitScheme_name, GitScheme_value)
	proto.RegisterEnum("bufbuild.buf.io.v1beta1.GitSubmoduleBehavior", GitSubmoduleBehavior_name, GitSubmoduleBehavior_value)
	proto.RegisterEnum("bufbuild.buf.io.v1beta1.BucketScheme", BucketScheme_name, BucketScheme_value)
	proto.RegisterType((*ImageRef)(nil), "bufbuild.buf.io.v1beta1.ImageRef")
	proto.RegisterType((*ArchiveRef)(nil), "bufbuild.buf.io.v1beta1.ArchiveRef")
	proto.RegisterType((*GitRef)(nil), "bufbuild.buf.io.v1beta1.GitRef")
	proto.RegisterType((*BucketRef)(nil), "bufbuild.buf.io.v1beta1.BucketRef")
	proto.RegisterType((*SourceRef)(nil), "bufbuild.buf.io.v1beta1.SourceRef")
	proto.RegisterType((*InputRef)(nil), "bufbuild.buf.io.v1beta1.InputRef")
}

func init() {
	proto.RegisterFile("bufbuild/buf/io/v1beta1/io.proto", fileDescriptor_49d667d72c1020ea)
}

var fileDescriptor_49d667d72c1020ea = []byte{
	// 794 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xcd, 0x92, 0xda, 0x46,
	0x10, 0xc7, 0x77, 0x80, 0x5d, 0x5b, 0xcd, 0xee, 0x7a, 0x3c, 0xc6, 0xac, 0x62, 0x27, 0x31, 0x51,
	0x1c, 0xd7, 0x86, 0xaa, 0x40, 0x39, 0xb9, 0xe5, 0x14, 0x24, 0x04, 0xc8, 0xe1, 0xc3, 0x25, 0xc1,
	0x1e, 0xf6, 0xa2, 0x92, 0xe4, 0x11, 0x4c, 0x2d, 0x20, 0x4a, 0x48, 0x5c, 0x72, 0xce, 0x21, 0x8f,
	0x90, 0xbc, 0x43, 0x1e, 0x29, 0x55, 0x79, 0x93, 0xa4, 0x34, 0x08, 0x90, 0xb4, 0x88, 0x5c, 0x72,
	0x9b, 0xee, 0xf9, 0x77, 0xeb, 0xd7, 0x3d, 0x33, 0x2d, 0xa8, 0xd9, 0xa1, 0x6b, 0x87, 0x6c, 0xfe,
	0xa9, 0x69, 0x87, 0x6e, 0x93, 0x79, 0xcd, 0xcd, 0x7b, 0x9b, 0x06, 0xd6, 0xfb, 0x26, 0xf3, 0x1a,
	0x2b, 0xdf, 0x0b, 0x3c, 0x72, 0xb3, 0x53, 0x34, 0xec, 0xd0, 0x6d, 0x30, 0xaf, 0x11, 0x2b, 0xa4,
	0x3f, 0x11, 0x3c, 0xd5, 0x16, 0xd6, 0x94, 0xea, 0xd4, 0x25, 0x6d, 0x28, 0xbb, 0x6c, 0x4e, 0xcd,
	0xb5, 0x33, 0xa3, 0x0b, 0x2a, 0xa2, 0x1a, 0xba, 0xbd, 0xfe, 0xfe, 0xeb, 0x46, 0x4e, 0x6c, 0xa3,
	0xc3, 0xe6, 0xd4, 0xe0, 0x52, 0x1d, 0xdc, 0xfd, 0x9a, 0x74, 0xe1, 0x92, 0x45, 0x19, 0x4d, 0xd7,
	0xf3, 0x17, 0x56, 0x20, 0x16, 0x78, 0x9a, 0xb7, 0xb9, 0x69, 0xf8, 0xe7, 0x3b, 0x5c, 0xab, 0x97,
	0xd9, 0xc1, 0x20, 0x04, 0x4a, 0x2b, 0x2b, 0x98, 0x89, 0xc5, 0x1a, 0xba, 0x15, 0x74, 0xbe, 0x96,
	0xfe, 0x46, 0x00, 0x2d, 0xdf, 0x99, 0xb1, 0xcd, 0xff, 0x48, 0x3c, 0x80, 0x6b, 0x6b, 0x9b, 0x33,
	0xcd, 0xfc, 0x2e, 0x37, 0x51, 0x8c, 0x10, 0x53, 0x5f, 0x59, 0x49, 0xf3, 0x18, 0x37, 0xf9, 0x16,
	0xf0, 0x3a, 0xf0, 0xd9, 0xca, 0x74, 0xbc, 0xc5, 0xca, 0x5b, 0xd2, 0x65, 0xb0, 0x16, 0x4b, 0x35,
	0x74, 0x7b, 0xa5, 0x3f, 0xe3, 0x7e, 0x65, 0xef, 0x96, 0xfe, 0x41, 0x70, 0xd1, 0x65, 0x41, 0x54,
	0x5e, 0x0b, 0x60, 0xca, 0x82, 0x74, 0x75, 0x52, 0x2e, 0x54, 0x97, 0x05, 0x71, 0x71, 0xc2, 0x74,
	0xb7, 0xdc, 0xc3, 0x14, 0x12, 0x30, 0x22, 0x5c, 0xd8, 0xbe, 0xb5, 0x74, 0x62, 0xc4, 0xde, 0x99,
	0x1e, 0xdb, 0x84, 0x40, 0x31, 0xb0, 0xa6, 0x9c, 0x2c, 0x72, 0x47, 0x06, 0x71, 0xa0, 0xca, 0x21,
	0x42, 0x7b, 0xe1, 0x7d, 0x0a, 0xe7, 0xd4, 0xb4, 0xe9, 0xcc, 0xda, 0x30, 0xcf, 0x17, 0xcf, 0x39,
	0xd0, 0x77, 0x27, 0x81, 0x76, 0x51, 0x72, 0x1c, 0xa4, 0x57, 0xa6, 0x47, 0xbc, 0x72, 0x19, 0x04,
	0x9f, 0xba, 0xd4, 0xa7, 0x4b, 0x87, 0x4a, 0x0f, 0x20, 0xc8, 0xa1, 0xf3, 0x40, 0x79, 0x0f, 0x3e,
	0xc0, 0x95, 0xcd, 0x8d, 0x74, 0x1b, 0xbe, 0xc9, 0xfd, 0xea, 0x36, 0x34, 0xee, 0xc4, 0xa5, 0x9d,
	0xb0, 0x8e, 0x35, 0x43, 0xfa, 0x0b, 0x81, 0x60, 0x78, 0xa1, 0xef, 0xf0, 0x0b, 0xd5, 0x81, 0xf2,
	0xee, 0x2a, 0xf8, 0xd4, 0xe5, 0xdf, 0x2a, 0x9f, 0xb8, 0x50, 0x87, 0xab, 0xd8, 0x3b, 0xd3, 0xc1,
	0x3a, 0x5c, 0xcc, 0x1f, 0xe1, 0x49, 0xd4, 0xb4, 0x28, 0x47, 0x81, 0xe7, 0x78, 0x73, 0xaa, 0x4b,
	0xdb, 0xf8, 0x8b, 0xe9, 0xf6, 0xd4, 0x15, 0x80, 0xb8, 0xe2, 0x28, 0xbc, 0xc8, 0xc3, 0xa5, 0xff,
	0x28, 0x77, 0x9b, 0x41, 0xb0, 0x77, 0x86, 0xfc, 0x04, 0xce, 0x37, 0xd6, 0x3c, 0xa4, 0xd2, 0x1f,
	0xd1, 0x0b, 0x5f, 0xae, 0x42, 0x9e, 0xfa, 0x27, 0x10, 0xb6, 0x6f, 0xf3, 0x50, 0xdc, 0x57, 0xa7,
	0x1f, 0xe6, 0x36, 0xf1, 0x53, 0xb6, 0x9b, 0x11, 0x0a, 0xc0, 0x9a, 0x77, 0x2b, 0x51, 0x5b, 0x3e,
	0xdc, 0xbe, 0xb1, 0x11, 0xdc, 0x7a, 0x67, 0xec, 0xe1, 0xea, 0xbf, 0x23, 0x80, 0xc3, 0xa3, 0x24,
	0xaf, 0xe1, 0xa6, 0xa3, 0xf5, 0x55, 0xd3, 0x50, 0x7a, 0xea, 0x40, 0x35, 0x27, 0x43, 0xe3, 0xa3,
	0xaa, 0x68, 0x1d, 0x4d, 0x6d, 0xe3, 0x33, 0x52, 0x01, 0x9c, 0xdc, 0xec, 0x8d, 0xc7, 0x1f, 0x31,
	0x22, 0x2f, 0xe1, 0x79, 0xd6, 0x6b, 0xe0, 0x42, 0x56, 0x1c, 0xad, 0x71, 0x31, 0x2b, 0x36, 0xc6,
	0x6d, 0x6d, 0x84, 0x4b, 0x59, 0xf1, 0x70, 0xd2, 0xef, 0xe3, 0xf3, 0xfa, 0x6f, 0x08, 0xca, 0x89,
	0xd9, 0x44, 0x3e, 0x07, 0x51, 0x1b, 0xb4, 0xba, 0xaa, 0xd9, 0x19, 0xe9, 0x83, 0xd6, 0xf8, 0x31,
	0x5d, 0x6a, 0x57, 0xd6, 0x86, 0x18, 0x91, 0x2a, 0x90, 0xac, 0xb7, 0x7b, 0x8f, 0x0b, 0x11, 0x48,
	0xca, 0xff, 0xc1, 0x18, 0x0d, 0x71, 0x91, 0xdc, 0xc0, 0x8b, 0x47, 0xee, 0xee, 0x3d, 0x2e, 0xd5,
	0x2d, 0xb8, 0x4a, 0x8d, 0x1c, 0xf2, 0x25, 0xbc, 0x6a, 0xe9, 0x4a, 0x4f, 0xbb, 0xcb, 0xc1, 0xa9,
	0x02, 0xc9, 0xec, 0x8f, 0x5b, 0x3a, 0x46, 0x44, 0x84, 0xca, 0x63, 0x7f, 0x84, 0x54, 0xff, 0x05,
	0x84, 0xfd, 0x00, 0x21, 0xaf, 0xa0, 0xda, 0xd5, 0xc6, 0xc7, 0xcf, 0xe1, 0x05, 0x3c, 0x4b, 0xec,
	0xc5, 0xc7, 0x50, 0x01, 0x9c, 0x71, 0x46, 0xa7, 0x90, 0x96, 0xc6, 0x87, 0x40, 0xe0, 0x3a, 0xe1,
	0x34, 0x8c, 0x1e, 0x2e, 0xd5, 0x7f, 0x45, 0x50, 0x39, 0x36, 0x2d, 0xc8, 0x3b, 0x90, 0xb8, 0x78,
	0x22, 0x0f, 0x46, 0xed, 0x49, 0x5f, 0x35, 0x65, 0xb5, 0xd7, 0xba, 0xd3, 0x46, 0x7a, 0x06, 0xea,
	0x0d, 0xbc, 0xce, 0xd1, 0x0d, 0x47, 0x43, 0x15, 0x23, 0xf2, 0x16, 0x6a, 0x39, 0x02, 0x5d, 0x55,
	0x26, 0xba, 0xa1, 0xdd, 0xa9, 0xb8, 0x50, 0x6f, 0xc3, 0x65, 0x72, 0x7c, 0x90, 0x2f, 0xe0, 0x33,
	0x79, 0xa2, 0xfc, 0xac, 0xe6, 0xb4, 0xe2, 0x25, 0x3c, 0x4f, 0x6f, 0xb7, 0x35, 0x1d, 0x23, 0xb9,
	0x7c, 0x2f, 0x30, 0x2f, 0x7e, 0x02, 0xf6, 0x05, 0xff, 0x03, 0xff, 0xf0, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xa5, 0xd5, 0xa8, 0xf4, 0xa5, 0x07, 0x00, 0x00,
}
