// Code generated by protoc-gen-go. DO NOT EDIT.
// source: artifact.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "www.velocidex.com/golang/velociraptor/proto"
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

type Tool struct {
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url           string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	ServeLocally  bool   `protobuf:"varint,3,opt,name=serve_locally,json=serveLocally,proto3" json:"serve_locally,omitempty"`
	FilestorePath string `protobuf:"bytes,4,opt,name=filestore_path,json=filestorePath,proto3" json:"filestore_path,omitempty"`
	// The name of the client on the endpoint.
	Filename string `protobuf:"bytes,5,opt,name=filename,proto3" json:"filename,omitempty"`
	// Hex encoded sha256 hash.
	Hash                 string   `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tool) Reset()         { *m = Tool{} }
func (m *Tool) String() string { return proto.CompactTextString(m) }
func (*Tool) ProtoMessage()    {}
func (*Tool) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{0}
}

func (m *Tool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tool.Unmarshal(m, b)
}
func (m *Tool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tool.Marshal(b, m, deterministic)
}
func (m *Tool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tool.Merge(m, src)
}
func (m *Tool) XXX_Size() int {
	return xxx_messageInfo_Tool.Size(m)
}
func (m *Tool) XXX_DiscardUnknown() {
	xxx_messageInfo_Tool.DiscardUnknown(m)
}

var xxx_messageInfo_Tool proto.InternalMessageInfo

func (m *Tool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tool) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Tool) GetServeLocally() bool {
	if m != nil {
		return m.ServeLocally
	}
	return false
}

func (m *Tool) GetFilestorePath() string {
	if m != nil {
		return m.FilestorePath
	}
	return ""
}

func (m *Tool) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Tool) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type ThirdParty struct {
	Tools                []*Tool  `protobuf:"bytes,1,rep,name=tools,proto3" json:"tools,omitempty"`
	Version              uint64   `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThirdParty) Reset()         { *m = ThirdParty{} }
func (m *ThirdParty) String() string { return proto.CompactTextString(m) }
func (*ThirdParty) ProtoMessage()    {}
func (*ThirdParty) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{1}
}

func (m *ThirdParty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThirdParty.Unmarshal(m, b)
}
func (m *ThirdParty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThirdParty.Marshal(b, m, deterministic)
}
func (m *ThirdParty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThirdParty.Merge(m, src)
}
func (m *ThirdParty) XXX_Size() int {
	return xxx_messageInfo_ThirdParty.Size(m)
}
func (m *ThirdParty) XXX_DiscardUnknown() {
	xxx_messageInfo_ThirdParty.DiscardUnknown(m)
}

var xxx_messageInfo_ThirdParty proto.InternalMessageInfo

func (m *ThirdParty) GetTools() []*Tool {
	if m != nil {
		return m.Tools
	}
	return nil
}

func (m *ThirdParty) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type FieldDescriptor struct {
	FriendlyName string   `protobuf:"bytes,1,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Repeated     bool     `protobuf:"varint,3,opt,name=repeated,proto3" json:"repeated,omitempty"`
	Type         string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Doc          string   `protobuf:"bytes,5,opt,name=doc,proto3" json:"doc,omitempty"`
	Labels       []string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty"`
	Default      string   `protobuf:"bytes,7,opt,name=default,proto3" json:"default,omitempty"`
	// Same as doc field.
	Description          string   `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldDescriptor) Reset()         { *m = FieldDescriptor{} }
func (m *FieldDescriptor) String() string { return proto.CompactTextString(m) }
func (*FieldDescriptor) ProtoMessage()    {}
func (*FieldDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{2}
}

func (m *FieldDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldDescriptor.Unmarshal(m, b)
}
func (m *FieldDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldDescriptor.Marshal(b, m, deterministic)
}
func (m *FieldDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldDescriptor.Merge(m, src)
}
func (m *FieldDescriptor) XXX_Size() int {
	return xxx_messageInfo_FieldDescriptor.Size(m)
}
func (m *FieldDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_FieldDescriptor proto.InternalMessageInfo

func (m *FieldDescriptor) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

func (m *FieldDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FieldDescriptor) GetRepeated() bool {
	if m != nil {
		return m.Repeated
	}
	return false
}

func (m *FieldDescriptor) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *FieldDescriptor) GetDoc() string {
	if m != nil {
		return m.Doc
	}
	return ""
}

func (m *FieldDescriptor) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *FieldDescriptor) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

func (m *FieldDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type EnumValue struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumValue) Reset()         { *m = EnumValue{} }
func (m *EnumValue) String() string { return proto.CompactTextString(m) }
func (*EnumValue) ProtoMessage()    {}
func (*EnumValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{3}
}

func (m *EnumValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumValue.Unmarshal(m, b)
}
func (m *EnumValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumValue.Marshal(b, m, deterministic)
}
func (m *EnumValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumValue.Merge(m, src)
}
func (m *EnumValue) XXX_Size() int {
	return xxx_messageInfo_EnumValue.Size(m)
}
func (m *EnumValue) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumValue.DiscardUnknown(m)
}

var xxx_messageInfo_EnumValue proto.InternalMessageInfo

func (m *EnumValue) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *EnumValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TypeDescriptor struct {
	Doc string `protobuf:"bytes,1,opt,name=doc,proto3" json:"doc,omitempty"`
	// Same as doc field.
	Description  string             `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	Fields       []*FieldDescriptor `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	Name         string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	FriendlyName string             `protobuf:"bytes,7,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	Kind         string             `protobuf:"bytes,4,opt,name=kind,proto3" json:"kind,omitempty"`
	// The fields are all part of a single one of structure. NOTE:
	// Currently we only support an all or nothing approach to oneof -
	// there can be at most a single oneof group within the proto and
	// it implies that all the fields belong to it.
	Oneof                bool         `protobuf:"varint,5,opt,name=oneof,proto3" json:"oneof,omitempty"`
	Default              string       `protobuf:"bytes,6,opt,name=default,proto3" json:"default,omitempty"`
	AllowedValues        []*EnumValue `protobuf:"bytes,8,rep,name=allowed_values,json=allowedValues,proto3" json:"allowed_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TypeDescriptor) Reset()         { *m = TypeDescriptor{} }
func (m *TypeDescriptor) String() string { return proto.CompactTextString(m) }
func (*TypeDescriptor) ProtoMessage()    {}
func (*TypeDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{4}
}

func (m *TypeDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeDescriptor.Unmarshal(m, b)
}
func (m *TypeDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeDescriptor.Marshal(b, m, deterministic)
}
func (m *TypeDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeDescriptor.Merge(m, src)
}
func (m *TypeDescriptor) XXX_Size() int {
	return xxx_messageInfo_TypeDescriptor.Size(m)
}
func (m *TypeDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_TypeDescriptor proto.InternalMessageInfo

func (m *TypeDescriptor) GetDoc() string {
	if m != nil {
		return m.Doc
	}
	return ""
}

func (m *TypeDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TypeDescriptor) GetFields() []*FieldDescriptor {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *TypeDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TypeDescriptor) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

func (m *TypeDescriptor) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *TypeDescriptor) GetOneof() bool {
	if m != nil {
		return m.Oneof
	}
	return false
}

func (m *TypeDescriptor) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

func (m *TypeDescriptor) GetAllowedValues() []*EnumValue {
	if m != nil {
		return m.AllowedValues
	}
	return nil
}

type Types struct {
	Items                []*TypeDescriptor `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Types) Reset()         { *m = Types{} }
func (m *Types) String() string { return proto.CompactTextString(m) }
func (*Types) ProtoMessage()    {}
func (*Types) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{5}
}

func (m *Types) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Types.Unmarshal(m, b)
}
func (m *Types) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Types.Marshal(b, m, deterministic)
}
func (m *Types) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Types.Merge(m, src)
}
func (m *Types) XXX_Size() int {
	return xxx_messageInfo_Types.Size(m)
}
func (m *Types) XXX_DiscardUnknown() {
	xxx_messageInfo_Types.DiscardUnknown(m)
}

var xxx_messageInfo_Types proto.InternalMessageInfo

func (m *Types) GetItems() []*TypeDescriptor {
	if m != nil {
		return m.Items
	}
	return nil
}

type ArtifactParameter struct {
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Default     string `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// For type = choice
	Choices              []string `protobuf:"bytes,6,rep,name=choices,proto3" json:"choices,omitempty"`
	FriendlyName         string   `protobuf:"bytes,5,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactParameter) Reset()         { *m = ArtifactParameter{} }
func (m *ArtifactParameter) String() string { return proto.CompactTextString(m) }
func (*ArtifactParameter) ProtoMessage()    {}
func (*ArtifactParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{6}
}

func (m *ArtifactParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactParameter.Unmarshal(m, b)
}
func (m *ArtifactParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactParameter.Marshal(b, m, deterministic)
}
func (m *ArtifactParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactParameter.Merge(m, src)
}
func (m *ArtifactParameter) XXX_Size() int {
	return xxx_messageInfo_ArtifactParameter.Size(m)
}
func (m *ArtifactParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactParameter.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactParameter proto.InternalMessageInfo

func (m *ArtifactParameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArtifactParameter) GetDefault() string {
	if m != nil {
		return m.Default
	}
	return ""
}

func (m *ArtifactParameter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ArtifactParameter) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ArtifactParameter) GetChoices() []string {
	if m != nil {
		return m.Choices
	}
	return nil
}

func (m *ArtifactParameter) GetFriendlyName() string {
	if m != nil {
		return m.FriendlyName
	}
	return ""
}

type ArtifactSource struct {
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Precondition         string   `protobuf:"bytes,1,opt,name=precondition,proto3" json:"precondition,omitempty"`
	Query                string   `protobuf:"bytes,6,opt,name=query,proto3" json:"query,omitempty"`
	Queries              []string `protobuf:"bytes,2,rep,name=queries,proto3" json:"queries,omitempty"`
	PostProcess          []string `protobuf:"bytes,5,rep,name=post_process,json=postProcess,proto3" json:"post_process,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtifactSource) Reset()         { *m = ArtifactSource{} }
func (m *ArtifactSource) String() string { return proto.CompactTextString(m) }
func (*ArtifactSource) ProtoMessage()    {}
func (*ArtifactSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{7}
}

func (m *ArtifactSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactSource.Unmarshal(m, b)
}
func (m *ArtifactSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactSource.Marshal(b, m, deterministic)
}
func (m *ArtifactSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactSource.Merge(m, src)
}
func (m *ArtifactSource) XXX_Size() int {
	return xxx_messageInfo_ArtifactSource.Size(m)
}
func (m *ArtifactSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactSource.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactSource proto.InternalMessageInfo

func (m *ArtifactSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ArtifactSource) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ArtifactSource) GetPrecondition() string {
	if m != nil {
		return m.Precondition
	}
	return ""
}

func (m *ArtifactSource) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *ArtifactSource) GetQueries() []string {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *ArtifactSource) GetPostProcess() []string {
	if m != nil {
		return m.PostProcess
	}
	return nil
}

type Report struct {
	// Each report type will be handled differently. Read about the
	// different types in reporting.go
	Type                 string               `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Template             string               `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	Parameters           []*ArtifactParameter `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{8}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Report) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *Report) GetParameters() []*ArtifactParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type Artifact struct {
	Name                string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description         string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Author              string   `protobuf:"bytes,12,opt,name=author,proto3" json:"author,omitempty"`
	Reference           []string `protobuf:"bytes,5,rep,name=reference,proto3" json:"reference,omitempty"`
	RequiredPermissions []string `protobuf:"bytes,13,rep,name=required_permissions,json=requiredPermissions,proto3" json:"required_permissions,omitempty"`
	// An optional list of tool descriptions. These are only used to
	// initialize Velociraptor if there is no previous tool
	// definition. It will not override existing tools. The user may
	// override the tool name in order to control where it will be
	// downloaded from.
	Tools []*Tool `protobuf:"bytes,15,rep,name=tools,proto3" json:"tools,omitempty"`
	// If set here this applies to all sources.
	Precondition string               `protobuf:"bytes,8,opt,name=precondition,proto3" json:"precondition,omitempty"`
	Parameters   []*ArtifactParameter `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty"`
	// Unfortunately we can not use enum due to limitations in
	// yaml->json->protobuf conversion.
	Type     string            `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	Sources  []*ArtifactSource `protobuf:"bytes,4,rep,name=sources,proto3" json:"sources,omitempty"`
	Includes []string          `protobuf:"bytes,9,rep,name=includes,proto3" json:"includes,omitempty"`
	Reports  []*Report         `protobuf:"bytes,11,rep,name=reports,proto3" json:"reports,omitempty"`
	// Internal use only
	Raw                  string   `protobuf:"bytes,7,opt,name=raw,proto3" json:"raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{9}
}

func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (m *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(m, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Artifact) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Artifact) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Artifact) GetReference() []string {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *Artifact) GetRequiredPermissions() []string {
	if m != nil {
		return m.RequiredPermissions
	}
	return nil
}

func (m *Artifact) GetTools() []*Tool {
	if m != nil {
		return m.Tools
	}
	return nil
}

func (m *Artifact) GetPrecondition() string {
	if m != nil {
		return m.Precondition
	}
	return ""
}

func (m *Artifact) GetParameters() []*ArtifactParameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *Artifact) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Artifact) GetSources() []*ArtifactSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *Artifact) GetIncludes() []string {
	if m != nil {
		return m.Includes
	}
	return nil
}

func (m *Artifact) GetReports() []*Report {
	if m != nil {
		return m.Reports
	}
	return nil
}

func (m *Artifact) GetRaw() string {
	if m != nil {
		return m.Raw
	}
	return ""
}

type ArtifactDescriptors struct {
	Items                []*Artifact `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ArtifactDescriptors) Reset()         { *m = ArtifactDescriptors{} }
func (m *ArtifactDescriptors) String() string { return proto.CompactTextString(m) }
func (*ArtifactDescriptors) ProtoMessage()    {}
func (*ArtifactDescriptors) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1932e98ed811590, []int{10}
}

func (m *ArtifactDescriptors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactDescriptors.Unmarshal(m, b)
}
func (m *ArtifactDescriptors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactDescriptors.Marshal(b, m, deterministic)
}
func (m *ArtifactDescriptors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactDescriptors.Merge(m, src)
}
func (m *ArtifactDescriptors) XXX_Size() int {
	return xxx_messageInfo_ArtifactDescriptors.Size(m)
}
func (m *ArtifactDescriptors) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactDescriptors.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactDescriptors proto.InternalMessageInfo

func (m *ArtifactDescriptors) GetItems() []*Artifact {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Tool)(nil), "proto.Tool")
	proto.RegisterType((*ThirdParty)(nil), "proto.third_party")
	proto.RegisterType((*FieldDescriptor)(nil), "proto.FieldDescriptor")
	proto.RegisterType((*EnumValue)(nil), "proto.EnumValue")
	proto.RegisterType((*TypeDescriptor)(nil), "proto.TypeDescriptor")
	proto.RegisterType((*Types)(nil), "proto.Types")
	proto.RegisterType((*ArtifactParameter)(nil), "proto.ArtifactParameter")
	proto.RegisterType((*ArtifactSource)(nil), "proto.ArtifactSource")
	proto.RegisterType((*Report)(nil), "proto.Report")
	proto.RegisterType((*Artifact)(nil), "proto.Artifact")
	proto.RegisterType((*ArtifactDescriptors)(nil), "proto.ArtifactDescriptors")
}

func init() {
	proto.RegisterFile("artifact.proto", fileDescriptor_a1932e98ed811590)
}

var fileDescriptor_a1932e98ed811590 = []byte{
	// 1815 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0x41, 0x6f, 0x1c, 0x49,
	0xf5, 0x57, 0xdb, 0x19, 0x7b, 0x5c, 0x8e, 0x9d, 0xfd, 0xd7, 0xe6, 0x1f, 0x9a, 0x1c, 0x96, 0xb7,
	0x13, 0x2d, 0x38, 0xe0, 0xb4, 0xc5, 0x66, 0xa3, 0x04, 0xb3, 0x42, 0xcc, 0x24, 0x66, 0xd7, 0x91,
	0xe3, 0x71, 0xda, 0x56, 0xa2, 0xcd, 0xc5, 0x2a, 0x77, 0xbf, 0x99, 0x2e, 0x52, 0x53, 0xd5, 0xa9,
	0xaa, 0xf6, 0x78, 0xb8, 0x70, 0x40, 0x42, 0x5c, 0x56, 0x42, 0x20, 0x0e, 0x20, 0x3e, 0x00, 0x37,
	0xe0, 0x03, 0x70, 0xe3, 0x33, 0x20, 0xc1, 0x69, 0x81, 0xaf, 0xc1, 0x01, 0x55, 0x75, 0xf7, 0x74,
	0x8f, 0x1d, 0x40, 0x1c, 0x38, 0x4d, 0x57, 0xf5, 0xab, 0xf7, 0x7e, 0xef, 0xbd, 0xdf, 0xfb, 0x75,
	0x0d, 0xd9, 0x64, 0xda, 0xf2, 0x11, 0x4b, 0x6c, 0x94, 0x6b, 0x65, 0x15, 0xed, 0xf8, 0x9f, 0xdb,
	0x37, 0xfd, 0xcf, 0x8e, 0xc1, 0x09, 0x93, 0x96, 0x27, 0xe5, 0xcb, 0xde, 0x6f, 0x02, 0x72, 0xed,
	0x44, 0x29, 0x41, 0x29, 0xb9, 0x26, 0xd9, 0x04, 0xc3, 0x00, 0x82, 0xad, 0xb5, 0xd8, 0x3f, 0xd3,
	0x77, 0xc8, 0x72, 0xa1, 0x45, 0xb8, 0xe4, 0xb7, 0xdc, 0x23, 0xbd, 0x43, 0x36, 0x0c, 0xea, 0x73,
	0x3c, 0x15, 0x2a, 0x61, 0x42, 0xcc, 0xc2, 0x65, 0x08, 0xb6, 0xba, 0xf1, 0x75, 0xbf, 0x79, 0x50,
	0xee, 0xd1, 0x0f, 0xc8, 0xe6, 0x88, 0x0b, 0x34, 0x56, 0x69, 0x3c, 0xcd, 0x99, 0xcd, 0xc2, 0x6b,
	0xde, 0xc3, 0xc6, 0x7c, 0xf7, 0x88, 0xd9, 0x8c, 0xde, 0x26, 0x5d, 0xb7, 0xe1, 0xa3, 0x76, 0xbc,
	0xc1, 0x7c, 0xed, 0xd0, 0x64, 0xcc, 0x64, 0xe1, 0x4a, 0x89, 0xc6, 0x3d, 0xf7, 0x9e, 0x92, 0x75,
	0x9b, 0x71, 0x9d, 0x9e, 0xe6, 0x4c, 0xdb, 0x19, 0x7d, 0x9f, 0x74, 0xac, 0x52, 0xc2, 0x84, 0x01,
	0x2c, 0x6f, 0xad, 0x7f, 0xb8, 0x5e, 0x26, 0x14, 0xb9, 0x64, 0xe2, 0xf2, 0x0d, 0x0d, 0xc9, 0xea,
	0x39, 0x6a, 0xc3, 0x95, 0xf4, 0x39, 0x5c, 0x8b, 0xeb, 0x65, 0xef, 0x8b, 0x80, 0xdc, 0xf8, 0x1e,
	0x47, 0x91, 0x3e, 0x41, 0x93, 0x68, 0x9e, 0x5b, 0xa5, 0x5d, 0x6e, 0x23, 0xcd, 0x51, 0xa6, 0x62,
	0x76, 0xda, 0x2a, 0xc5, 0xf5, 0x7a, 0xf3, 0xb0, 0x02, 0xe6, 0xdf, 0x2d, 0xb5, 0xca, 0x74, 0x9b,
	0x74, 0x35, 0xe6, 0xc8, 0x2c, 0xa6, 0x55, 0x3d, 0xe6, 0x6b, 0x67, 0x6f, 0x67, 0x39, 0x56, 0x15,
	0xf0, 0xcf, 0xae, 0xac, 0xa9, 0x4a, 0xaa, 0x9c, 0xdd, 0x23, 0xbd, 0x45, 0x56, 0x04, 0x3b, 0x43,
	0x61, 0xc2, 0x15, 0x58, 0xde, 0x5a, 0x8b, 0xab, 0x95, 0x4b, 0x20, 0xc5, 0x11, 0x2b, 0x84, 0x0d,
	0x57, 0xbd, 0x75, 0xbd, 0xa4, 0x40, 0xd6, 0xd3, 0x0a, 0xba, 0x4b, 0xaf, 0xeb, 0xdf, 0xb6, 0xb7,
	0x7a, 0x0f, 0xc8, 0xda, 0x9e, 0x2c, 0x26, 0x2f, 0x98, 0x28, 0x90, 0xde, 0x24, 0x9d, 0x73, 0xf7,
	0xe0, 0x73, 0xea, 0xc4, 0xe5, 0xe2, 0x6d, 0xc9, 0xf4, 0xfe, 0xba, 0x44, 0x36, 0x4f, 0x66, 0x39,
	0xb6, 0x0a, 0x53, 0xe1, 0x0d, 0x1a, 0xbc, 0x97, 0xa2, 0xaf, 0x5d, 0x89, 0x4e, 0x23, 0xb2, 0x32,
	0x72, 0xf5, 0x35, 0xe1, 0x92, 0x6f, 0xcf, 0xad, 0xaa, 0x3d, 0x97, 0x8a, 0x1e, 0x57, 0x56, 0x73,
	0x28, 0xcb, 0xad, 0xba, 0x5e, 0x69, 0xc8, 0xea, 0xdb, 0x1b, 0xf2, 0x9a, 0xcb, 0xb4, 0x2e, 0xb0,
	0x7b, 0x76, 0xd9, 0x2a, 0x89, 0x6a, 0xe4, 0x4b, 0xdc, 0x8d, 0xcb, 0x05, 0x1d, 0x36, 0xc5, 0xf4,
	0xb4, 0x1a, 0x3c, 0xf8, 0xdb, 0x3f, 0xfe, 0xfe, 0xc7, 0x60, 0x87, 0xde, 0x3b, 0xc9, 0x10, 0xbe,
	0x6f, 0x94, 0x04, 0x94, 0x89, 0x4a, 0x31, 0x85, 0xca, 0x0e, 0x7c, 0x99, 0x60, 0xa4, 0x34, 0xd8,
	0x8c, 0x1b, 0x70, 0xdd, 0x8b, 0x9a, 0x1e, 0x3c, 0x24, 0x9b, 0x4c, 0x08, 0x35, 0xc5, 0xf4, 0xd4,
	0x1b, 0x9a, 0xb0, 0xeb, 0x73, 0x7d, 0xa7, 0xca, 0x75, 0x5e, 0xfe, 0x78, 0xa3, 0xb2, 0xf3, 0x2b,
	0xd3, 0xfb, 0x88, 0x74, 0x5c, 0x89, 0x0d, 0xfd, 0x06, 0xe9, 0x70, 0x8b, 0x93, 0x9a, 0xc3, 0xff,
	0x5f, 0x73, 0x78, 0xa1, 0xfe, 0x71, 0x69, 0xd3, 0xfb, 0xe5, 0x32, 0xf9, 0xbf, 0x7e, 0x35, 0xda,
	0x47, 0x4c, 0xb3, 0x09, 0x5a, 0xd4, 0x6f, 0x9d, 0xdb, 0x16, 0x6d, 0x96, 0xfe, 0x2d, 0x6d, 0x96,
	0xaf, 0x36, 0xee, 0x6d, 0x84, 0x0d, 0xc9, 0x6a, 0x92, 0x29, 0x9e, 0x60, 0xcd, 0xcf, 0x7a, 0x79,
	0xb5, 0x45, 0x9d, 0xab, 0x2d, 0xda, 0xfd, 0x22, 0xf8, 0x8b, 0x2b, 0xf4, 0x9f, 0x03, 0xf2, 0xa7,
	0xa0, 0x4e, 0xc0, 0xc0, 0x84, 0xcd, 0x80, 0x25, 0x09, 0xe6, 0x16, 0xf2, 0x3a, 0x1b, 0x03, 0xd3,
	0x8c, 0x27, 0x19, 0x30, 0x8d, 0xc0, 0x52, 0xd7, 0x09, 0xab, 0xc0, 0x66, 0x08, 0x26, 0x51, 0x39,
	0x42, 0xae, 0xb9, 0x6b, 0x85, 0x02, 0xbc, 0xc0, 0xa4, 0x70, 0x70, 0x23, 0x38, 0x1c, 0x9e, 0xec,
	0xed, 0x02, 0x13, 0xa2, 0xed, 0xc5, 0x9d, 0x37, 0x56, 0x73, 0x39, 0x36, 0x70, 0x0f, 0xf8, 0x08,
	0x66, 0xaa, 0x00, 0x89, 0x98, 0x82, 0x51, 0x13, 0xb4, 0x19, 0x97, 0x63, 0x40, 0x61, 0xd0, 0xfb,
	0x7e, 0x53, 0xa0, 0x9e, 0x41, 0xc2, 0x24, 0x14, 0x32, 0x67, 0xc9, 0x6b, 0xc0, 0x68, 0x1c, 0xc1,
	0x48, 0xab, 0x09, 0x3c, 0x3d, 0x1e, 0x1e, 0x42, 0x61, 0x9c, 0xb9, 0xb3, 0x74, 0xcb, 0x23, 0xa6,
	0x0d, 0x6e, 0xdd, 0x85, 0x17, 0xcf, 0x0f, 0x60, 0x54, 0xc8, 0xc4, 0xa3, 0xe8, 0xfd, 0x64, 0x85,
	0x6c, 0xd6, 0xa9, 0x1d, 0xab, 0x42, 0x27, 0x48, 0x7f, 0x1b, 0xb4, 0x29, 0x3d, 0xf8, 0x55, 0xe0,
	0xd9, 0xf6, 0xf3, 0x80, 0xfe, 0x34, 0x70, 0x7c, 0x73, 0xaf, 0x40, 0x8d, 0x4a, 0x62, 0xd5, 0x8a,
	0x0d, 0xc6, 0x9f, 0x8d, 0x60, 0x7f, 0x04, 0x52, 0x59, 0x30, 0x68, 0x61, 0x8a, 0x50, 0x54, 0x48,
	0xcd, 0xfc, 0x0c, 0x36, 0x47, 0xb8, 0x35, 0x28, 0x46, 0x11, 0x9c, 0xb4, 0x37, 0x13, 0x35, 0xc9,
	0xb9, 0x40, 0x0d, 0x53, 0x2e, 0x04, 0x8c, 0x51, 0xa2, 0x66, 0x16, 0x81, 0x55, 0xc9, 0x4e, 0xb9,
	0xcd, 0xca, 0xc8, 0x0e, 0x46, 0x54, 0xd1, 0xe6, 0xf3, 0x60, 0x91, 0x1d, 0x9e, 0x02, 0x83, 0xd7,
	0x1e, 0x37, 0xd2, 0xa4, 0x0f, 0xad, 0x97, 0x55, 0x79, 0x9b, 0xd9, 0xa8, 0x91, 0x1f, 0x2a, 0x8b,
	0xc0, 0xad, 0xaf, 0xe7, 0x19, 0x02, 0x97, 0x16, 0x75, 0xae, 0x84, 0x13, 0xc3, 0x32, 0xac, 0xb2,
	0x19, 0xea, 0x06, 0x69, 0xcb, 0xa7, 0x89, 0x16, 0xa9, 0x98, 0x91, 0xeb, 0xb9, 0xc6, 0x44, 0xc9,
	0x94, 0x7b, 0x3c, 0x9e, 0xe2, 0x83, 0x27, 0x1e, 0xcf, 0x77, 0xe8, 0xc7, 0x7d, 0xdf, 0x03, 0xbc,
	0xc8, 0x35, 0x1a, 0x27, 0xea, 0x8e, 0x18, 0x67, 0x08, 0xe8, 0x66, 0xd1, 0x87, 0x9b, 0xd3, 0xa5,
	0xee, 0x60, 0x03, 0x33, 0x5e, 0xf0, 0x4c, 0x3f, 0x22, 0x1d, 0x5f, 0x98, 0x4a, 0x18, 0xde, 0xf3,
	0x21, 0x42, 0x7a, 0xab, 0x0f, 0x93, 0x42, 0x58, 0x7e, 0x4f, 0x70, 0x89, 0x3e, 0x9a, 0xb7, 0x8a,
	0x4b, 0x63, 0x6a, 0xc9, 0xaa, 0x7b, 0xe0, 0x58, 0x8a, 0xdc, 0xda, 0xe0, 0x95, 0x3f, 0x77, 0x42,
	0xe3, 0xe7, 0xe5, 0x36, 0xd8, 0x8c, 0xd9, 0xb2, 0x05, 0xba, 0x90, 0xc0, 0x25, 0x28, 0x9d, 0xa2,
	0x8e, 0x60, 0x28, 0xc5, 0x0c, 0x54, 0x61, 0xf3, 0xc2, 0x96, 0x44, 0x73, 0x1d, 0x15, 0xcc, 0xd8,
	0x79, 0x7f, 0x84, 0x70, 0xb9, 0x24, 0x4a, 0x08, 0x4c, 0x2c, 0xa6, 0x51, 0x5c, 0x87, 0xa2, 0x86,
	0x5c, 0xcf, 0x95, 0xb1, 0xa7, 0xb9, 0x56, 0x09, 0x1a, 0x13, 0x76, 0x7c, 0xe8, 0x23, 0x1f, 0xfa,
	0x29, 0xfd, 0xb4, 0x0f, 0x82, 0x1b, 0xeb, 0x58, 0xf2, 0xe6, 0x0a, 0x88, 0x33, 0xf4, 0x38, 0xd4,
	0x39, 0x6a, 0x1f, 0x51, 0xa3, 0x29, 0x84, 0x35, 0xbe, 0x89, 0xce, 0x27, 0x54, 0x3e, 0xb9, 0x1c,
	0x47, 0xf1, 0xba, 0xdb, 0x39, 0x2a, 0x37, 0x76, 0xef, 0xfa, 0x09, 0xbe, 0x43, 0xde, 0x7f, 0x99,
	0xa1, 0xc6, 0x45, 0x02, 0x8e, 0xd1, 0x1a, 0xc7, 0x42, 0x48, 0x99, 0x65, 0x51, 0xef, 0x0f, 0x4b,
	0x64, 0x25, 0xc6, 0x5c, 0x69, 0x4b, 0x9f, 0x55, 0x5a, 0x52, 0x36, 0xee, 0x5b, 0x1e, 0xe2, 0x7d,
	0xfa, 0x4d, 0x27, 0x6f, 0x0e, 0xa0, 0xf6, 0x56, 0xbb, 0xf0, 0x6c, 0x78, 0xb8, 0x7f, 0x32, 0x8c,
	0xf7, 0x0f, 0x3f, 0x39, 0x7d, 0xd2, 0xdf, 0x3f, 0xf8, 0x6c, 0x1b, 0x8e, 0x86, 0xc7, 0x27, 0xa7,
	0x47, 0xf1, 0xf0, 0xf1, 0xde, 0xf1, 0xf1, 0xfe, 0xe1, 0x27, 0x95, 0x0c, 0xdd, 0x26, 0x5d, 0x8b,
	0x93, 0xdc, 0x71, 0xa9, 0xd2, 0xb5, 0xf9, 0x9a, 0x3e, 0x22, 0xa4, 0x99, 0xff, 0x70, 0xd9, 0xcb,
	0x69, 0x58, 0xc9, 0xe9, 0x15, 0xd1, 0x8c, 0x5b, 0xb6, 0xbb, 0x9f, 0x97, 0xea, 0xf4, 0xe3, 0x80,
	0xfc, 0x28, 0xe8, 0x57, 0x98, 0x80, 0x9b, 0xf9, 0xd0, 0xa4, 0x4d, 0x87, 0xaa, 0x8e, 0x5d, 0x9e,
	0xc0, 0xa6, 0x4b, 0x6e, 0x08, 0x35, 0xd6, 0xdc, 0xf7, 0xdc, 0xc9, 0x05, 0xd6, 0x4e, 0x5d, 0x12,
	0x6d, 0xbf, 0x29, 0xe6, 0x28, 0x53, 0x47, 0x4d, 0x25, 0x21, 0x51, 0xd2, 0xe2, 0x85, 0x8d, 0x7a,
	0xbf, 0x23, 0xa4, 0x5b, 0x23, 0xa6, 0xbf, 0x0f, 0xda, 0xf2, 0x3e, 0xf8, 0x75, 0x29, 0x22, 0xbf,
	0x08, 0xe8, 0xcf, 0x2e, 0x89, 0x48, 0x03, 0x27, 0x82, 0xe3, 0x4c, 0x15, 0x22, 0x75, 0x08, 0x0a,
	0xc9, 0xdf, 0x14, 0x08, 0x4c, 0xa6, 0x5e, 0x6e, 0x5d, 0x0c, 0xc6, 0x25, 0xa4, 0xca, 0x9a, 0x08,
	0xfa, 0x4e, 0x57, 0x46, 0x85, 0x00, 0x93, 0x64, 0x38, 0x41, 0x97, 0xb3, 0x9b, 0x20, 0x8d, 0xec,
	0x35, 0x24, 0xcc, 0xe2, 0x58, 0x79, 0xfa, 0xf8, 0xa1, 0x4d, 0x95, 0x2d, 0x65, 0xf1, 0x80, 0xcb,
	0xe2, 0x22, 0x1a, 0x68, 0x35, 0x35, 0xa8, 0x4d, 0xf4, 0x38, 0xd3, 0x6a, 0x82, 0x9f, 0x72, 0x77,
	0xab, 0x9b, 0x55, 0x32, 0xf2, 0x7c, 0x51, 0x45, 0x7c, 0xa7, 0x06, 0x3b, 0x1e, 0xf8, 0x5d, 0xfa,
	0xb5, 0x97, 0x8e, 0x8d, 0x8b, 0x02, 0x66, 0xc0, 0xea, 0x99, 0x9f, 0x50, 0x55, 0xd7, 0xf2, 0x92,
	0x12, 0x3c, 0x24, 0x2b, 0xac, 0xb0, 0x99, 0xd2, 0xe1, 0x75, 0xef, 0xed, 0x2b, 0xde, 0xdb, 0x97,
	0xe9, 0x97, 0xfa, 0x7e, 0xf7, 0x4a, 0x09, 0xe2, 0xca, 0x9c, 0x3e, 0x21, 0x6b, 0x1a, 0x47, 0xa8,
	0x51, 0x26, 0x58, 0x4d, 0xca, 0x57, 0xfd, 0x59, 0xa0, 0xef, 0xb9, 0x66, 0x57, 0xaf, 0x1a, 0x19,
	0x6b, 0x5c, 0x34, 0x07, 0xa9, 0x21, 0x37, 0x35, 0xbe, 0x29, 0xb8, 0xc6, 0xf4, 0x34, 0x47, 0x3d,
	0xe1, 0x5e, 0x6f, 0x4c, 0xb8, 0xe1, 0x1d, 0x7e, 0xd7, 0x3b, 0xdc, 0xa5, 0x8f, 0x9a, 0xd1, 0xab,
	0xad, 0xa1, 0x65, 0xdd, 0xca, 0xee, 0x72, 0xa8, 0x77, 0x6b, 0xfb, 0xa3, 0xc6, 0xbc, 0xb9, 0xdf,
	0xde, 0xf8, 0x97, 0xf7, 0xdb, 0xcb, 0x02, 0xd9, 0xfd, 0x9f, 0x09, 0xa4, 0xf8, 0x6f, 0xc6, 0x6b,
	0xf0, 0xa1, 0x47, 0xb0, 0x4d, 0xbf, 0x7e, 0xd4, 0x7c, 0x92, 0xcb, 0xd8, 0xb9, 0x56, 0xe7, 0xbc,
	0xf5, 0x5d, 0x6f, 0x6a, 0xd0, 0xf2, 0x4f, 0x87, 0x95, 0x6e, 0x10, 0x9f, 0xcf, 0xb7, 0xbd, 0xb7,
	0x07, 0xf4, 0xbe, 0x63, 0xbc, 0xad, 0xb4, 0x63, 0xfe, 0xe1, 0x8e, 0xe0, 0x71, 0x39, 0x6d, 0x87,
	0xc3, 0xf8, 0x59, 0xff, 0x60, 0x1b, 0xf6, 0x5e, 0xec, 0x1d, 0x9e, 0x6c, 0xc3, 0xf1, 0x5e, 0xfc,
	0x62, 0x2f, 0xae, 0x94, 0xe3, 0x15, 0x59, 0x2d, 0xf3, 0x32, 0xe1, 0xb5, 0x85, 0x9b, 0xd6, 0xe2,
	0x37, 0x7b, 0x70, 0xd7, 0x87, 0xba, 0x43, 0xff, 0xb3, 0xcc, 0xc5, 0xb5, 0x43, 0xfa, 0x92, 0x74,
	0xb9, 0x4c, 0x44, 0x91, 0xa2, 0x09, 0xd7, 0x3c, 0x21, 0xe6, 0x80, 0x1b, 0x42, 0x2c, 0x7e, 0xf4,
	0x7c, 0x35, 0x26, 0xa8, 0xc7, 0x4e, 0x52, 0x74, 0x55, 0x72, 0x3f, 0x68, 0x51, 0x3c, 0x77, 0x46,
	0x27, 0x64, 0xb5, 0x14, 0x0e, 0x13, 0xae, 0x7b, 0xd0, 0x1b, 0x15, 0xe8, 0x52, 0x5d, 0xdb, 0x7d,
	0x6e, 0x78, 0xe7, 0xed, 0x9d, 0xfb, 0x5c, 0x59, 0x94, 0x96, 0xbb, 0x3f, 0x64, 0x0b, 0xda, 0x7e,
	0x99, 0x7b, 0x75, 0x0c, 0xfa, 0x88, 0x2c, 0x6b, 0x36, 0x2d, 0xef, 0xd8, 0xcd, 0x90, 0xb8, 0x9a,
	0x6b, 0x36, 0x85, 0xcf, 0xfa, 0xcf, 0x0e, 0xae, 0x5c, 0x57, 0xa2, 0xd8, 0x1d, 0xd9, 0xfd, 0xa1,
	0x17, 0xd0, 0x19, 0x99, 0xf6, 0x65, 0x53, 0xae, 0xa9, 0x66, 0xb9, 0x01, 0xd6, 0x7c, 0x35, 0xdd,
	0xe7, 0x4f, 0x63, 0x61, 0xd8, 0x99, 0xc0, 0x6d, 0x48, 0x55, 0x52, 0x4c, 0x50, 0xfa, 0x5b, 0x01,
	0x9b, 0x45, 0xcd, 0xb5, 0xd0, 0x5f, 0xfd, 0x84, 0x00, 0x76, 0xa6, 0x8a, 0xb9, 0xb4, 0x56, 0xec,
	0x74, 0xd7, 0x39, 0x77, 0x2d, 0x62, 0x92, 0x89, 0xd9, 0x0f, 0xaa, 0x5b, 0xd9, 0x24, 0xea, 0x7d,
	0x4c, 0xde, 0xad, 0x1d, 0x34, 0xd7, 0x66, 0x43, 0x3f, 0x58, 0xbc, 0x5d, 0xdf, 0xb8, 0xd4, 0xf3,
	0xea, 0x5e, 0x3d, 0x78, 0xf8, 0xea, 0xc1, 0x74, 0x3a, 0x8d, 0xce, 0x51, 0xa8, 0x84, 0xa7, 0x78,
	0x11, 0x25, 0x6a, 0xb2, 0x33, 0x56, 0x82, 0xc9, 0xf1, 0x4e, 0xb9, 0xa9, 0x99, 0x73, 0xb8, 0x33,
	0x6f, 0xdf, 0x8e, 0x77, 0x74, 0xb6, 0xe2, 0x7f, 0xee, 0xff, 0x33, 0x00, 0x00, 0xff, 0xff, 0xfd,
	0x5d, 0xe4, 0xc9, 0x71, 0x0f, 0x00, 0x00,
}
