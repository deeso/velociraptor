// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hunts.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	proto1 "www.velocidex.com/golang/velociraptor/flows/proto"
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

type HuntOsCondition_OS int32

const (
	HuntOsCondition_ALL     HuntOsCondition_OS = 0
	HuntOsCondition_WINDOWS HuntOsCondition_OS = 1
	HuntOsCondition_LINUX   HuntOsCondition_OS = 2
	HuntOsCondition_OSX     HuntOsCondition_OS = 3
)

var HuntOsCondition_OS_name = map[int32]string{
	0: "ALL",
	1: "WINDOWS",
	2: "LINUX",
	3: "OSX",
}

var HuntOsCondition_OS_value = map[string]int32{
	"ALL":     0,
	"WINDOWS": 1,
	"LINUX":   2,
	"OSX":     3,
}

func (x HuntOsCondition_OS) String() string {
	return proto.EnumName(HuntOsCondition_OS_name, int32(x))
}

func (HuntOsCondition_OS) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0cb9fe450e8d0e17, []int{1, 0}
}

type Hunt_State int32

const (
	Hunt_UNSET    Hunt_State = 0
	Hunt_PAUSED   Hunt_State = 1
	Hunt_RUNNING  Hunt_State = 2
	Hunt_STOPPED  Hunt_State = 3
	Hunt_ARCHIVED Hunt_State = 4
)

var Hunt_State_name = map[int32]string{
	0: "UNSET",
	1: "PAUSED",
	2: "RUNNING",
	3: "STOPPED",
	4: "ARCHIVED",
}

var Hunt_State_value = map[string]int32{
	"UNSET":    0,
	"PAUSED":   1,
	"RUNNING":  2,
	"STOPPED":  3,
	"ARCHIVED": 4,
}

func (x Hunt_State) String() string {
	return proto.EnumName(Hunt_State_name, int32(x))
}

func (Hunt_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0cb9fe450e8d0e17, []int{4, 0}
}

type HuntLabelCondition struct {
	Label                []string `protobuf:"bytes,1,rep,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HuntLabelCondition) Reset()         { *m = HuntLabelCondition{} }
func (m *HuntLabelCondition) String() string { return proto.CompactTextString(m) }
func (*HuntLabelCondition) ProtoMessage()    {}
func (*HuntLabelCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cb9fe450e8d0e17, []int{0}
}

func (m *HuntLabelCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HuntLabelCondition.Unmarshal(m, b)
}
func (m *HuntLabelCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HuntLabelCondition.Marshal(b, m, deterministic)
}
func (m *HuntLabelCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HuntLabelCondition.Merge(m, src)
}
func (m *HuntLabelCondition) XXX_Size() int {
	return xxx_messageInfo_HuntLabelCondition.Size(m)
}
func (m *HuntLabelCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_HuntLabelCondition.DiscardUnknown(m)
}

var xxx_messageInfo_HuntLabelCondition proto.InternalMessageInfo

func (m *HuntLabelCondition) GetLabel() []string {
	if m != nil {
		return m.Label
	}
	return nil
}

type HuntOsCondition struct {
	Os                   HuntOsCondition_OS `protobuf:"varint,1,opt,name=os,proto3,enum=proto.HuntOsCondition_OS" json:"os,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HuntOsCondition) Reset()         { *m = HuntOsCondition{} }
func (m *HuntOsCondition) String() string { return proto.CompactTextString(m) }
func (*HuntOsCondition) ProtoMessage()    {}
func (*HuntOsCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cb9fe450e8d0e17, []int{1}
}

func (m *HuntOsCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HuntOsCondition.Unmarshal(m, b)
}
func (m *HuntOsCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HuntOsCondition.Marshal(b, m, deterministic)
}
func (m *HuntOsCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HuntOsCondition.Merge(m, src)
}
func (m *HuntOsCondition) XXX_Size() int {
	return xxx_messageInfo_HuntOsCondition.Size(m)
}
func (m *HuntOsCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_HuntOsCondition.DiscardUnknown(m)
}

var xxx_messageInfo_HuntOsCondition proto.InternalMessageInfo

func (m *HuntOsCondition) GetOs() HuntOsCondition_OS {
	if m != nil {
		return m.Os
	}
	return HuntOsCondition_ALL
}

type HuntCondition struct {
	// Types that are valid to be assigned to UnionField:
	//	*HuntCondition_Labels
	//	*HuntCondition_Os
	UnionField           isHuntCondition_UnionField `protobuf_oneof:"union_field"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *HuntCondition) Reset()         { *m = HuntCondition{} }
func (m *HuntCondition) String() string { return proto.CompactTextString(m) }
func (*HuntCondition) ProtoMessage()    {}
func (*HuntCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cb9fe450e8d0e17, []int{2}
}

func (m *HuntCondition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HuntCondition.Unmarshal(m, b)
}
func (m *HuntCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HuntCondition.Marshal(b, m, deterministic)
}
func (m *HuntCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HuntCondition.Merge(m, src)
}
func (m *HuntCondition) XXX_Size() int {
	return xxx_messageInfo_HuntCondition.Size(m)
}
func (m *HuntCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_HuntCondition.DiscardUnknown(m)
}

var xxx_messageInfo_HuntCondition proto.InternalMessageInfo

type isHuntCondition_UnionField interface {
	isHuntCondition_UnionField()
}

type HuntCondition_Labels struct {
	Labels *HuntLabelCondition `protobuf:"bytes,2,opt,name=labels,proto3,oneof"`
}

type HuntCondition_Os struct {
	Os *HuntOsCondition `protobuf:"bytes,3,opt,name=os,proto3,oneof"`
}

func (*HuntCondition_Labels) isHuntCondition_UnionField() {}

func (*HuntCondition_Os) isHuntCondition_UnionField() {}

func (m *HuntCondition) GetUnionField() isHuntCondition_UnionField {
	if m != nil {
		return m.UnionField
	}
	return nil
}

func (m *HuntCondition) GetLabels() *HuntLabelCondition {
	if x, ok := m.GetUnionField().(*HuntCondition_Labels); ok {
		return x.Labels
	}
	return nil
}

func (m *HuntCondition) GetOs() *HuntOsCondition {
	if x, ok := m.GetUnionField().(*HuntCondition_Os); ok {
		return x.Os
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HuntCondition) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HuntCondition_Labels)(nil),
		(*HuntCondition_Os)(nil),
	}
}

type HuntStats struct {
	TotalClientsScheduled      uint64              `protobuf:"varint,9,opt,name=total_clients_scheduled,json=totalClientsScheduled,proto3" json:"total_clients_scheduled,omitempty"`
	TotalClientsWithResults    uint64              `protobuf:"varint,14,opt,name=total_clients_with_results,json=totalClientsWithResults,proto3" json:"total_clients_with_results,omitempty"`
	TotalClientsWithoutResults uint64              `protobuf:"varint,16,opt,name=total_clients_without_results,json=totalClientsWithoutResults,proto3" json:"total_clients_without_results,omitempty"`
	TotalClientsWithErrors     uint64              `protobuf:"varint,15,opt,name=total_clients_with_errors,json=totalClientsWithErrors,proto3" json:"total_clients_with_errors,omitempty"`
	Stopped                    bool                `protobuf:"varint,1,opt,name=stopped,proto3" json:"stopped,omitempty"`
	AvailableDownloads         *AvailableDownloads `protobuf:"bytes,2,opt,name=available_downloads,json=availableDownloads,proto3" json:"available_downloads,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}            `json:"-"`
	XXX_unrecognized           []byte              `json:"-"`
	XXX_sizecache              int32               `json:"-"`
}

func (m *HuntStats) Reset()         { *m = HuntStats{} }
func (m *HuntStats) String() string { return proto.CompactTextString(m) }
func (*HuntStats) ProtoMessage()    {}
func (*HuntStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cb9fe450e8d0e17, []int{3}
}

func (m *HuntStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HuntStats.Unmarshal(m, b)
}
func (m *HuntStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HuntStats.Marshal(b, m, deterministic)
}
func (m *HuntStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HuntStats.Merge(m, src)
}
func (m *HuntStats) XXX_Size() int {
	return xxx_messageInfo_HuntStats.Size(m)
}
func (m *HuntStats) XXX_DiscardUnknown() {
	xxx_messageInfo_HuntStats.DiscardUnknown(m)
}

var xxx_messageInfo_HuntStats proto.InternalMessageInfo

func (m *HuntStats) GetTotalClientsScheduled() uint64 {
	if m != nil {
		return m.TotalClientsScheduled
	}
	return 0
}

func (m *HuntStats) GetTotalClientsWithResults() uint64 {
	if m != nil {
		return m.TotalClientsWithResults
	}
	return 0
}

func (m *HuntStats) GetTotalClientsWithoutResults() uint64 {
	if m != nil {
		return m.TotalClientsWithoutResults
	}
	return 0
}

func (m *HuntStats) GetTotalClientsWithErrors() uint64 {
	if m != nil {
		return m.TotalClientsWithErrors
	}
	return 0
}

func (m *HuntStats) GetStopped() bool {
	if m != nil {
		return m.Stopped
	}
	return false
}

func (m *HuntStats) GetAvailableDownloads() *AvailableDownloads {
	if m != nil {
		return m.AvailableDownloads
	}
	return nil
}

type Hunt struct {
	HuntId               string                        `protobuf:"bytes,1,opt,name=hunt_id,json=huntId,proto3" json:"hunt_id,omitempty"`
	CreateTime           uint64                        `protobuf:"varint,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Creator              string                        `protobuf:"bytes,12,opt,name=creator,proto3" json:"creator,omitempty"`
	StartTime            uint64                        `protobuf:"varint,21,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Expires              uint64                        `protobuf:"varint,10,opt,name=expires,proto3" json:"expires,omitempty"`
	HuntDescription      string                        `protobuf:"bytes,11,opt,name=hunt_description,json=huntDescription,proto3" json:"hunt_description,omitempty"`
	StartRequest         *proto1.ArtifactCollectorArgs `protobuf:"bytes,16,opt,name=start_request,json=startRequest,proto3" json:"start_request,omitempty"`
	Condition            *HuntCondition                `protobuf:"bytes,4,opt,name=condition,proto3" json:"condition,omitempty"`
	ClientLimit          uint64                        `protobuf:"varint,6,opt,name=client_limit,json=clientLimit,proto3" json:"client_limit,omitempty"`
	Stats                *HuntStats                    `protobuf:"bytes,18,opt,name=stats,proto3" json:"stats,omitempty"`
	Artifacts            []string                      `protobuf:"bytes,17,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	ArtifactSources      []string                      `protobuf:"bytes,19,rep,name=artifact_sources,json=artifactSources,proto3" json:"artifact_sources,omitempty"`
	State                Hunt_State                    `protobuf:"varint,8,opt,name=state,proto3,enum=proto.Hunt_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Hunt) Reset()         { *m = Hunt{} }
func (m *Hunt) String() string { return proto.CompactTextString(m) }
func (*Hunt) ProtoMessage()    {}
func (*Hunt) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cb9fe450e8d0e17, []int{4}
}

func (m *Hunt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hunt.Unmarshal(m, b)
}
func (m *Hunt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hunt.Marshal(b, m, deterministic)
}
func (m *Hunt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hunt.Merge(m, src)
}
func (m *Hunt) XXX_Size() int {
	return xxx_messageInfo_Hunt.Size(m)
}
func (m *Hunt) XXX_DiscardUnknown() {
	xxx_messageInfo_Hunt.DiscardUnknown(m)
}

var xxx_messageInfo_Hunt proto.InternalMessageInfo

func (m *Hunt) GetHuntId() string {
	if m != nil {
		return m.HuntId
	}
	return ""
}

func (m *Hunt) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Hunt) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Hunt) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Hunt) GetExpires() uint64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (m *Hunt) GetHuntDescription() string {
	if m != nil {
		return m.HuntDescription
	}
	return ""
}

func (m *Hunt) GetStartRequest() *proto1.ArtifactCollectorArgs {
	if m != nil {
		return m.StartRequest
	}
	return nil
}

func (m *Hunt) GetCondition() *HuntCondition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *Hunt) GetClientLimit() uint64 {
	if m != nil {
		return m.ClientLimit
	}
	return 0
}

func (m *Hunt) GetStats() *HuntStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *Hunt) GetArtifacts() []string {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

func (m *Hunt) GetArtifactSources() []string {
	if m != nil {
		return m.ArtifactSources
	}
	return nil
}

func (m *Hunt) GetState() Hunt_State {
	if m != nil {
		return m.State
	}
	return Hunt_UNSET
}

type ListHuntsRequest struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                uint64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	IncludeArchived      bool     `protobuf:"varint,3,opt,name=include_archived,json=includeArchived,proto3" json:"include_archived,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListHuntsRequest) Reset()         { *m = ListHuntsRequest{} }
func (m *ListHuntsRequest) String() string { return proto.CompactTextString(m) }
func (*ListHuntsRequest) ProtoMessage()    {}
func (*ListHuntsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cb9fe450e8d0e17, []int{5}
}

func (m *ListHuntsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListHuntsRequest.Unmarshal(m, b)
}
func (m *ListHuntsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListHuntsRequest.Marshal(b, m, deterministic)
}
func (m *ListHuntsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListHuntsRequest.Merge(m, src)
}
func (m *ListHuntsRequest) XXX_Size() int {
	return xxx_messageInfo_ListHuntsRequest.Size(m)
}
func (m *ListHuntsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListHuntsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListHuntsRequest proto.InternalMessageInfo

func (m *ListHuntsRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListHuntsRequest) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListHuntsRequest) GetIncludeArchived() bool {
	if m != nil {
		return m.IncludeArchived
	}
	return false
}

type ListHuntsResponse struct {
	Items                []*Hunt  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListHuntsResponse) Reset()         { *m = ListHuntsResponse{} }
func (m *ListHuntsResponse) String() string { return proto.CompactTextString(m) }
func (*ListHuntsResponse) ProtoMessage()    {}
func (*ListHuntsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cb9fe450e8d0e17, []int{6}
}

func (m *ListHuntsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListHuntsResponse.Unmarshal(m, b)
}
func (m *ListHuntsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListHuntsResponse.Marshal(b, m, deterministic)
}
func (m *ListHuntsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListHuntsResponse.Merge(m, src)
}
func (m *ListHuntsResponse) XXX_Size() int {
	return xxx_messageInfo_ListHuntsResponse.Size(m)
}
func (m *ListHuntsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListHuntsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListHuntsResponse proto.InternalMessageInfo

func (m *ListHuntsResponse) GetItems() []*Hunt {
	if m != nil {
		return m.Items
	}
	return nil
}

type GetHuntRequest struct {
	HuntId               string   `protobuf:"bytes,1,opt,name=hunt_id,json=huntId,proto3" json:"hunt_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHuntRequest) Reset()         { *m = GetHuntRequest{} }
func (m *GetHuntRequest) String() string { return proto.CompactTextString(m) }
func (*GetHuntRequest) ProtoMessage()    {}
func (*GetHuntRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cb9fe450e8d0e17, []int{7}
}

func (m *GetHuntRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHuntRequest.Unmarshal(m, b)
}
func (m *GetHuntRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHuntRequest.Marshal(b, m, deterministic)
}
func (m *GetHuntRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHuntRequest.Merge(m, src)
}
func (m *GetHuntRequest) XXX_Size() int {
	return xxx_messageInfo_GetHuntRequest.Size(m)
}
func (m *GetHuntRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHuntRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHuntRequest proto.InternalMessageInfo

func (m *GetHuntRequest) GetHuntId() string {
	if m != nil {
		return m.HuntId
	}
	return ""
}

type GetHuntResultsRequest struct {
	Offset               uint64   `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                uint64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	HuntId               string   `protobuf:"bytes,3,opt,name=hunt_id,json=huntId,proto3" json:"hunt_id,omitempty"`
	Artifact             string   `protobuf:"bytes,4,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHuntResultsRequest) Reset()         { *m = GetHuntResultsRequest{} }
func (m *GetHuntResultsRequest) String() string { return proto.CompactTextString(m) }
func (*GetHuntResultsRequest) ProtoMessage()    {}
func (*GetHuntResultsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cb9fe450e8d0e17, []int{8}
}

func (m *GetHuntResultsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHuntResultsRequest.Unmarshal(m, b)
}
func (m *GetHuntResultsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHuntResultsRequest.Marshal(b, m, deterministic)
}
func (m *GetHuntResultsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHuntResultsRequest.Merge(m, src)
}
func (m *GetHuntResultsRequest) XXX_Size() int {
	return xxx_messageInfo_GetHuntResultsRequest.Size(m)
}
func (m *GetHuntResultsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHuntResultsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHuntResultsRequest proto.InternalMessageInfo

func (m *GetHuntResultsRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetHuntResultsRequest) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetHuntResultsRequest) GetHuntId() string {
	if m != nil {
		return m.HuntId
	}
	return ""
}

func (m *GetHuntResultsRequest) GetArtifact() string {
	if m != nil {
		return m.Artifact
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.HuntOsCondition_OS", HuntOsCondition_OS_name, HuntOsCondition_OS_value)
	proto.RegisterEnum("proto.Hunt_State", Hunt_State_name, Hunt_State_value)
	proto.RegisterType((*HuntLabelCondition)(nil), "proto.HuntLabelCondition")
	proto.RegisterType((*HuntOsCondition)(nil), "proto.HuntOsCondition")
	proto.RegisterType((*HuntCondition)(nil), "proto.HuntCondition")
	proto.RegisterType((*HuntStats)(nil), "proto.HuntStats")
	proto.RegisterType((*Hunt)(nil), "proto.Hunt")
	proto.RegisterType((*ListHuntsRequest)(nil), "proto.ListHuntsRequest")
	proto.RegisterType((*ListHuntsResponse)(nil), "proto.ListHuntsResponse")
	proto.RegisterType((*GetHuntRequest)(nil), "proto.GetHuntRequest")
	proto.RegisterType((*GetHuntResultsRequest)(nil), "proto.GetHuntResultsRequest")
}

func init() {
	proto.RegisterFile("hunts.proto", fileDescriptor_0cb9fe450e8d0e17)
}

var fileDescriptor_0cb9fe450e8d0e17 = []byte{
	// 1477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x72, 0x1b, 0x4b,
	0x15, 0xf6, 0xf8, 0x4f, 0x56, 0x2b, 0xb6, 0xc7, 0x7d, 0xfd, 0x33, 0x11, 0xb8, 0xaa, 0x99, 0x32,
	0x5c, 0x9b, 0x5b, 0x77, 0x0c, 0xb9, 0x55, 0x49, 0x41, 0x11, 0x8c, 0x64, 0x09, 0x4b, 0xc1, 0xb1,
	0xcd, 0x48, 0x46, 0x49, 0x36, 0xa2, 0x35, 0xd3, 0xb2, 0xba, 0x18, 0x4d, 0x2b, 0xd3, 0x3d, 0x51,
	0xcc, 0x92, 0x05, 0x95, 0x15, 0x2c, 0x78, 0x9a, 0xec, 0xd9, 0xf1, 0x06, 0x2c, 0x58, 0x00, 0x2b,
	0x1e, 0x80, 0x15, 0x0b, 0xaa, 0x4f, 0x6b, 0x46, 0xb2, 0xec, 0x64, 0x71, 0x37, 0xd2, 0x74, 0x9f,
	0xef, 0x9c, 0xef, 0x3b, 0xdd, 0x7d, 0x4e, 0x37, 0x2a, 0x0d, 0xd2, 0x58, 0x49, 0x6f, 0x94, 0x08,
	0x25, 0xf0, 0x0a, 0xfc, 0x95, 0xb7, 0xe1, 0xef, 0x58, 0xb2, 0x21, 0x8d, 0x15, 0x0f, 0x8c, 0xb1,
	0x7c, 0xd0, 0x8f, 0xc4, 0x58, 0x1e, 0x1b, 0x1b, 0x4d, 0x14, 0xef, 0xd3, 0x40, 0x75, 0x03, 0x11,
	0x45, 0x2c, 0x50, 0x22, 0x99, 0xa0, 0x4a, 0x80, 0x32, 0x03, 0xf7, 0x87, 0x08, 0x37, 0xd2, 0x58,
	0x9d, 0xd3, 0x1e, 0x8b, 0x4e, 0x45, 0x1c, 0x72, 0xc5, 0x45, 0x8c, 0xb7, 0xd1, 0x4a, 0xa4, 0x67,
	0x1c, 0x8b, 0x2c, 0x1d, 0x16, 0x7d, 0x33, 0x70, 0x23, 0xb4, 0xa9, 0xb1, 0x97, 0x72, 0x0a, 0x3c,
	0x42, 0x8b, 0x42, 0x3a, 0x16, 0xb1, 0x0e, 0x37, 0x9e, 0x3c, 0x36, 0x21, 0xbd, 0x39, 0x8c, 0x77,
	0xd9, 0xf2, 0x17, 0x85, 0x74, 0x3d, 0xb4, 0x78, 0xd9, 0xc2, 0x05, 0xb4, 0x54, 0x39, 0x3f, 0xb7,
	0x17, 0x70, 0x09, 0x15, 0x3a, 0xcd, 0x8b, 0xda, 0x65, 0xa7, 0x65, 0x5b, 0xb8, 0x88, 0x56, 0xce,
	0x9b, 0x17, 0xd7, 0xaf, 0xec, 0x45, 0x0d, 0xb8, 0x6c, 0xbd, 0xb2, 0x97, 0xdc, 0x7f, 0x5b, 0x68,
	0x5d, 0x87, 0x9a, 0x92, 0xfd, 0x0a, 0xad, 0x82, 0x10, 0xe9, 0x2c, 0x12, 0xeb, 0xb0, 0x74, 0x87,
	0xf0, 0x6e, 0x02, 0xd5, 0xdd, 0x7f, 0xfe, 0xef, 0x5f, 0x7f, 0xb5, 0x6c, 0x77, 0xe3, 0x25, 0x55,
	0xc1, 0x80, 0xf4, 0x6e, 0x09, 0x78, 0x36, 0x16, 0xfc, 0x49, 0x08, 0x5c, 0x05, 0xe5, 0x4b, 0x10,
	0x68, 0xf7, 0x61, 0xe5, 0x55, 0x07, 0xa2, 0x60, 0xd7, 0xbe, 0x1c, 0xb1, 0x84, 0x2a, 0x1e, 0xdf,
	0x90, 0xd6, 0xad, 0x54, 0x6c, 0xd8, 0x58, 0xd0, 0x29, 0xfd, 0xf4, 0x9b, 0xbf, 0x6b, 0xfb, 0xd7,
	0xe8, 0xab, 0xf6, 0x80, 0x91, 0x20, 0x73, 0x23, 0x4a, 0x90, 0x21, 0xd0, 0x0e, 0x84, 0x54, 0x92,
	0xf4, 0x45, 0x42, 0xd4, 0x80, 0x4b, 0xa2, 0xf7, 0xd1, 0xab, 0xae, 0xa3, 0x52, 0x1a, 0x73, 0x11,
	0x77, 0xfb, 0x9c, 0x45, 0xa1, 0xfb, 0x61, 0x15, 0x15, 0x35, 0x6f, 0x4b, 0x51, 0x25, 0xf1, 0x9f,
	0x2d, 0xb4, 0xa7, 0x84, 0xa2, 0x51, 0x37, 0x88, 0x38, 0x8b, 0x95, 0xec, 0xca, 0x60, 0xc0, 0xc2,
	0x34, 0x62, 0xa1, 0x53, 0x24, 0xd6, 0xe1, 0x72, 0xb5, 0x03, 0x9a, 0x7e, 0x8d, 0x7f, 0xae, 0x39,
	0x01, 0x4a, 0xe2, 0x74, 0xd8, 0x63, 0x09, 0x11, 0x7d, 0x32, 0x71, 0x22, 0x41, 0x9a, 0x24, 0x2c,
	0x56, 0xd1, 0x2d, 0xc9, 0xdd, 0xe7, 0x64, 0xb8, 0x9b, 0x6d, 0xf0, 0x6d, 0x65, 0x76, 0x7f, 0x07,
	0x82, 0x9d, 0x9a, 0x08, 0xf9, 0x34, 0xfe, 0xa3, 0x85, 0xca, 0x77, 0x15, 0x8d, 0xb9, 0x1a, 0x74,
	0x13, 0x26, 0xd3, 0x48, 0x49, 0x67, 0x03, 0x44, 0x35, 0x41, 0xd4, 0x29, 0xfe, 0x7e, 0xfb, 0x13,
	0x82, 0xb4, 0x0f, 0x99, 0xf8, 0x78, 0x6e, 0xd9, 0xc0, 0x4e, 0x67, 0x8d, 0xbe, 0x31, 0xfa, 0x7b,
	0xb3, 0x32, 0x3a, 0x5c, 0x0d, 0x26, 0x06, 0xfc, 0x17, 0x0b, 0xed, 0xdf, 0x17, 0x22, 0x52, 0x95,
	0x6b, 0xb1, 0x41, 0xcb, 0x15, 0x68, 0x79, 0x81, 0x8f, 0x3e, 0xab, 0x45, 0xa4, 0x6a, 0xaa, 0x67,
	0xff, 0xbe, 0x1e, 0x6d, 0xcf, 0x24, 0x95, 0xe7, 0x25, 0x89, 0x54, 0x65, 0xaa, 0xfe, 0x60, 0xa1,
	0xc7, 0x0f, 0x2c, 0x0f, 0x4b, 0x12, 0x91, 0x48, 0x67, 0x13, 0x14, 0x9d, 0x81, 0xa2, 0x0a, 0x3e,
	0xf8, 0xac, 0x22, 0xe3, 0xe2, 0xb9, 0x8f, 0x1f, 0x58, 0x9c, 0x3a, 0xd8, 0xfc, 0xdd, 0x79, 0x21,
	0x66, 0x1e, 0xdf, 0xa2, 0x82, 0x54, 0x62, 0x34, 0x62, 0x21, 0x94, 0xe2, 0x5a, 0xb5, 0x0b, 0x8c,
	0xaf, 0x71, 0xa7, 0xd9, 0x37, 0x7b, 0xce, 0x25, 0x91, 0x4c, 0x11, 0x35, 0x60, 0xb1, 0xfe, 0x81,
	0x33, 0x00, 0x93, 0xc6, 0xcd, 0x23, 0x6d, 0x8d, 0x82, 0x13, 0xa9, 0xa7, 0x87, 0x34, 0xe6, 0xa3,
	0x34, 0xa2, 0x8a, 0x85, 0xba, 0x76, 0x72, 0x8f, 0x21, 0x8d, 0xe9, 0x0d, 0x4b, 0x3c, 0x3f, 0xe3,
	0xc3, 0x2f, 0xd0, 0x17, 0xf4, 0x1d, 0xe5, 0x11, 0xed, 0x45, 0xac, 0x1b, 0x8a, 0x71, 0x1c, 0x09,
	0x1a, 0xce, 0x17, 0x68, 0x25, 0x43, 0xd4, 0x32, 0x80, 0x8f, 0xe9, 0xbd, 0x39, 0xf7, 0xbf, 0x08,
	0x2d, 0xeb, 0x52, 0xc0, 0x87, 0xa8, 0xa0, 0xe9, 0xba, 0xdc, 0xe4, 0x53, 0xac, 0x6e, 0x42, 0x3e,
	0x45, 0xb7, 0xa0, 0xcd, 0xa4, 0x59, 0xf3, 0x57, 0xb5, 0xbd, 0x19, 0xe2, 0xdf, 0xa2, 0x52, 0x90,
	0x30, 0xaa, 0x58, 0x57, 0xf1, 0x21, 0x03, 0xda, 0xe5, 0xea, 0x09, 0xa0, 0x7f, 0x82, 0x4a, 0x7e,
	0xed, 0x97, 0x35, 0xaa, 0x98, 0x36, 0xe1, 0xef, 0x74, 0x4c, 0xee, 0x93, 0x02, 0x20, 0x63, 0x2a,
	0x89, 0x71, 0x0d, 0x3d, 0x77, 0xfd, 0x54, 0x7f, 0xe9, 0xda, 0x6d, 0xf3, 0x21, 0xf3, 0x91, 0x31,
	0xe8, 0x6f, 0xfc, 0x0c, 0x15, 0x60, 0x24, 0x12, 0xe7, 0x11, 0x68, 0xd9, 0x87, 0xe8, 0x7b, 0x78,
	0xa7, 0x33, 0x10, 0x59, 0x84, 0x7c, 0x85, 0x4e, 0xfc, 0x0c, 0x8d, 0x43, 0x84, 0xa4, 0xa2, 0x89,
	0x32, 0xca, 0x76, 0x40, 0x59, 0x1d, 0x7c, 0x4f, 0xee, 0x2a, 0x3b, 0x78, 0x40, 0x19, 0x0d, 0x54,
	0x4a, 0x23, 0x5d, 0xbf, 0x3a, 0x84, 0x96, 0x88, 0x5a, 0xfa, 0xcb, 0xe8, 0x2b, 0xc2, 0x2c, 0xc8,
	0xeb, 0xa0, 0x02, 0x7b, 0x3f, 0xe2, 0x09, 0x93, 0x0e, 0x02, 0x8a, 0xe7, 0x40, 0xf1, 0xec, 0xa1,
	0xe4, 0x43, 0xc1, 0xe4, 0x0c, 0x8f, 0xf1, 0x3b, 0x71, 0x4b, 0x75, 0xfd, 0x71, 0x6b, 0x42, 0x67,
	0xd1, 0x70, 0x1d, 0xd9, 0xb0, 0x07, 0x21, 0x93, 0x41, 0xc2, 0x47, 0x7a, 0x71, 0x9c, 0x12, 0x2c,
	0x40, 0x19, 0x18, 0xb6, 0x31, 0x5c, 0x1c, 0x5f, 0x4a, 0x32, 0x83, 0xf0, 0x37, 0xb5, 0x4f, 0x6d,
	0x3a, 0x81, 0x3f, 0x58, 0x68, 0xdd, 0x2c, 0x43, 0xc2, 0xde, 0xa6, 0x4c, 0x2a, 0xa8, 0xd2, 0xd2,
	0x93, 0xef, 0x66, 0x47, 0x63, 0x72, 0x4b, 0x9d, 0x66, 0x97, 0x54, 0x25, 0xb9, 0x91, 0xd9, 0x3a,
	0xe1, 0xe7, 0xe7, 0x34, 0x8d, 0x83, 0x81, 0x11, 0x3d, 0xb9, 0xc7, 0xf4, 0x2e, 0x09, 0x73, 0x8a,
	0x4d, 0xfd, 0x10, 0xde, 0x37, 0xa3, 0xbc, 0x01, 0x73, 0x49, 0x54, 0x92, 0x32, 0xff, 0x11, 0x30,
	0xfb, 0x86, 0x18, 0xff, 0xc9, 0x42, 0xc5, 0x1c, 0xe3, 0x2c, 0x83, 0x8c, 0xed, 0x99, 0xce, 0x3f,
	0xed, 0xfb, 0x6f, 0x80, 0xbe, 0x8d, 0xab, 0x73, 0x7d, 0x7d, 0x40, 0x15, 0x19, 0xa6, 0x52, 0x91,
	0x1e, 0x23, 0x92, 0x2a, 0x2e, 0xfb, 0x3c, 0x6f, 0xac, 0x93, 0x0a, 0x51, 0x02, 0x8c, 0x59, 0xfb,
	0xf4, 0xdc, 0x0d, 0x38, 0xb2, 0x79, 0x6c, 0x7f, 0x2a, 0x01, 0xbf, 0x46, 0x8f, 0x4c, 0x06, 0xdd,
	0x88, 0x0f, 0xb9, 0x72, 0x56, 0x61, 0x03, 0x9f, 0x02, 0xf9, 0x8f, 0xb0, 0xf7, 0xa9, 0x6e, 0x31,
	0x73, 0x52, 0x78, 0x14, 0x91, 0x24, 0xd5, 0xeb, 0xe1, 0xf9, 0x25, 0x63, 0x3f, 0xd7, 0xa1, 0xf0,
	0x0f, 0xd0, 0x8a, 0xd4, 0x17, 0x8a, 0x83, 0x21, 0x4d, 0x7b, 0x26, 0x4d, 0xb8, 0x68, 0x7c, 0x63,
	0xc6, 0x2f, 0x51, 0x31, 0x7b, 0x27, 0x48, 0x67, 0x4b, 0x5f, 0xf6, 0xd5, 0x63, 0xe0, 0x3f, 0xc2,
	0x5f, 0x56, 0x48, 0xc4, 0xa5, 0xd2, 0xcc, 0x39, 0x64, 0x86, 0x7b, 0x94, 0x88, 0x30, 0x0d, 0x98,
	0xf4, 0xfc, 0x69, 0x04, 0x4c, 0x91, 0x9d, 0x3f, 0x3b, 0xa4, 0x48, 0x93, 0x80, 0x49, 0xe7, 0x0b,
	0x88, 0x9a, 0x67, 0x75, 0x3f, 0x2a, 0x99, 0x20, 0x1f, 0x0c, 0xbe, 0x99, 0xa1, 0x5a, 0x06, 0x84,
	0xdf, 0x9a, 0xcc, 0x98, 0xb3, 0x06, 0x8f, 0x8e, 0xad, 0x99, 0xcc, 0x3c, 0x9d, 0x1a, 0xab, 0x36,
	0x80, 0xaa, 0x8a, 0x7f, 0xd1, 0xce, 0x3a, 0x9f, 0x9e, 0xd6, 0x8c, 0xd9, 0x2e, 0x3d, 0xd0, 0xee,
	0xd2, 0xb9, 0x76, 0x77, 0x76, 0xdd, 0xf4, 0xcc, 0x22, 0x31, 0xf7, 0x1f, 0x16, 0x5a, 0x81, 0xd0,
	0xfa, 0x9d, 0x72, 0x7d, 0xd1, 0xaa, 0xb7, 0xed, 0x05, 0xdc, 0x40, 0xab, 0x57, 0x95, 0xeb, 0x56,
	0xbd, 0x66, 0x5b, 0xe5, 0x9f, 0xfd, 0xe7, 0xe3, 0xdf, 0x3e, 0x5a, 0x4f, 0x1b, 0xf9, 0x9e, 0xc4,
	0x42, 0xe5, 0x67, 0x80, 0xc4, 0x6c, 0x9c, 0xef, 0x5d, 0x2f, 0x55, 0x24, 0xa0, 0x31, 0x9c, 0x91,
	0x49, 0x65, 0xe3, 0xaf, 0x51, 0xc1, 0xbf, 0xbe, 0xb8, 0x68, 0x5e, 0x9c, 0xd9, 0x8b, 0x65, 0x02,
	0xa1, 0xca, 0x8d, 0x49, 0x7f, 0x4e, 0xd2, 0x38, 0xd6, 0xcf, 0x0e, 0x1a, 0x87, 0x24, 0x61, 0x34,
	0xbc, 0xf5, 0xf0, 0x01, 0x2a, 0xb4, 0xda, 0x97, 0x57, 0x57, 0xf5, 0x9a, 0xbd, 0x54, 0xde, 0x03,
	0xf8, 0x16, 0xc0, 0x07, 0x74, 0xda, 0xcf, 0xf1, 0x57, 0x68, 0xad, 0xe2, 0x9f, 0x36, 0x9a, 0xbf,
	0xa9, 0xd7, 0xec, 0xe5, 0xf2, 0x3e, 0xc0, 0xf6, 0x72, 0x58, 0x8f, 0xb1, 0x98, 0xd0, 0x24, 0x18,
	0xf0, 0x77, 0xfa, 0x8c, 0xfe, 0x0e, 0xd9, 0xe7, 0x5c, 0x2a, 0x6d, 0x96, 0x59, 0xb5, 0xec, 0xa2,
	0x55, 0xd1, 0xef, 0x4b, 0xa6, 0xa0, 0x05, 0x2f, 0xfb, 0x93, 0x91, 0x7e, 0x1a, 0x06, 0x22, 0x8d,
	0x95, 0xe9, 0xb5, 0xbe, 0x19, 0xe0, 0x23, 0x64, 0xf3, 0x38, 0x88, 0xd2, 0x90, 0x75, 0xb3, 0xb0,
	0xf0, 0xb6, 0x5a, 0xf3, 0x37, 0x27, 0xf3, 0x95, 0xc9, 0xb4, 0xfb, 0x14, 0x6d, 0xcd, 0x90, 0xc9,
	0x91, 0x88, 0x25, 0xc3, 0xdf, 0x43, 0x2b, 0x5c, 0xb1, 0xa1, 0x84, 0x07, 0x67, 0xe9, 0x49, 0x69,
	0x66, 0x57, 0x7d, 0x63, 0x71, 0x8f, 0xd0, 0xc6, 0x19, 0x03, 0xb7, 0x4c, 0xe2, 0xde, 0xdc, 0x35,
	0x91, 0xdd, 0x0a, 0xee, 0xef, 0xd1, 0x4e, 0x0e, 0x35, 0x57, 0xf8, 0xb7, 0x4a, 0x6a, 0x26, 0xfe,
	0xd2, 0x6c, 0x7c, 0x5c, 0x46, 0x6b, 0xd9, 0xb1, 0x84, 0x3e, 0x52, 0xf4, 0xf3, 0x71, 0xf5, 0xc7,
	0x6f, 0x8e, 0xc7, 0xe3, 0xb1, 0xf7, 0x8e, 0x45, 0x22, 0xe0, 0x21, 0x7b, 0xef, 0x05, 0x62, 0x78,
	0x7c, 0x23, 0x22, 0x1a, 0xdf, 0x1c, 0x9b, 0xc9, 0x84, 0x8e, 0x94, 0x48, 0x8e, 0xe9, 0x88, 0x9b,
	0x97, 0x7a, 0x6f, 0x15, 0xfe, 0xbe, 0xf9, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x52, 0xa6,
	0x98, 0xe9, 0x0b, 0x00, 0x00,
}
