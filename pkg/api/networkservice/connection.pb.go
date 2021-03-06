// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connection.proto

package networkservice

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type State int32

const (
	State_UP   State = 0
	State_DOWN State = 1
)

var State_name = map[int32]string{
	0: "UP",
	1: "DOWN",
}

var State_value = map[string]int32{
	"UP":   0,
	"DOWN": 1,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{0}
}

type ConnectionEventType int32

const (
	ConnectionEventType_INITIAL_STATE_TRANSFER ConnectionEventType = 0
	ConnectionEventType_UPDATE                 ConnectionEventType = 1
	ConnectionEventType_DELETE                 ConnectionEventType = 2
)

var ConnectionEventType_name = map[int32]string{
	0: "INITIAL_STATE_TRANSFER",
	1: "UPDATE",
	2: "DELETE",
}

var ConnectionEventType_value = map[string]int32{
	"INITIAL_STATE_TRANSFER": 0,
	"UPDATE":                 1,
	"DELETE":                 2,
}

func (x ConnectionEventType) String() string {
	return proto.EnumName(ConnectionEventType_name, int32(x))
}

func (ConnectionEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{1}
}

type Mechanism struct {
	Cls                  string            `protobuf:"bytes,1,opt,name=cls,proto3" json:"cls,omitempty"`
	Type                 string            `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Parameters           map[string]string `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Mechanism) Reset()         { *m = Mechanism{} }
func (m *Mechanism) String() string { return proto.CompactTextString(m) }
func (*Mechanism) ProtoMessage()    {}
func (*Mechanism) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{0}
}

func (m *Mechanism) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mechanism.Unmarshal(m, b)
}
func (m *Mechanism) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mechanism.Marshal(b, m, deterministic)
}
func (m *Mechanism) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mechanism.Merge(m, src)
}
func (m *Mechanism) XXX_Size() int {
	return xxx_messageInfo_Mechanism.Size(m)
}
func (m *Mechanism) XXX_DiscardUnknown() {
	xxx_messageInfo_Mechanism.DiscardUnknown(m)
}

var xxx_messageInfo_Mechanism proto.InternalMessageInfo

func (m *Mechanism) GetCls() string {
	if m != nil {
		return m.Cls
	}
	return ""
}

func (m *Mechanism) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Mechanism) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type PathSegment struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Token                string               `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Expires              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=expires,proto3" json:"expires,omitempty"`
	Metrics              map[string]string    `protobuf:"bytes,5,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PathSegment) Reset()         { *m = PathSegment{} }
func (m *PathSegment) String() string { return proto.CompactTextString(m) }
func (*PathSegment) ProtoMessage()    {}
func (*PathSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{1}
}

func (m *PathSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PathSegment.Unmarshal(m, b)
}
func (m *PathSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PathSegment.Marshal(b, m, deterministic)
}
func (m *PathSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathSegment.Merge(m, src)
}
func (m *PathSegment) XXX_Size() int {
	return xxx_messageInfo_PathSegment.Size(m)
}
func (m *PathSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_PathSegment.DiscardUnknown(m)
}

var xxx_messageInfo_PathSegment proto.InternalMessageInfo

func (m *PathSegment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PathSegment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PathSegment) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PathSegment) GetExpires() *timestamp.Timestamp {
	if m != nil {
		return m.Expires
	}
	return nil
}

func (m *PathSegment) GetMetrics() map[string]string {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type Path struct {
	Index                uint32         `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	PathSegments         []*PathSegment `protobuf:"bytes,2,rep,name=path_segments,json=pathSegments,proto3" json:"path_segments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Path) Reset()         { *m = Path{} }
func (m *Path) String() string { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()    {}
func (*Path) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{2}
}

func (m *Path) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Path.Unmarshal(m, b)
}
func (m *Path) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Path.Marshal(b, m, deterministic)
}
func (m *Path) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Path.Merge(m, src)
}
func (m *Path) XXX_Size() int {
	return xxx_messageInfo_Path.Size(m)
}
func (m *Path) XXX_DiscardUnknown() {
	xxx_messageInfo_Path.DiscardUnknown(m)
}

var xxx_messageInfo_Path proto.InternalMessageInfo

func (m *Path) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Path) GetPathSegments() []*PathSegment {
	if m != nil {
		return m.PathSegments
	}
	return nil
}

type Connection struct {
	Id                         string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NetworkService             string             `protobuf:"bytes,2,opt,name=network_service,json=networkService,proto3" json:"network_service,omitempty"`
	Mechanism                  *Mechanism         `protobuf:"bytes,3,opt,name=mechanism,proto3" json:"mechanism,omitempty"`
	Context                    *ConnectionContext `protobuf:"bytes,4,opt,name=context,proto3" json:"context,omitempty"`
	Labels                     map[string]string  `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Path                       *Path              `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	NetworkServiceEndpointName string             `protobuf:"bytes,7,opt,name=network_service_endpoint_name,json=networkServiceEndpointName,proto3" json:"network_service_endpoint_name,omitempty"`
	State                      State              `protobuf:"varint,9,opt,name=state,proto3,enum=connection.State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}           `json:"-"`
	XXX_unrecognized           []byte             `json:"-"`
	XXX_sizecache              int32              `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{3}
}

func (m *Connection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connection.Unmarshal(m, b)
}
func (m *Connection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connection.Marshal(b, m, deterministic)
}
func (m *Connection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connection.Merge(m, src)
}
func (m *Connection) XXX_Size() int {
	return xxx_messageInfo_Connection.Size(m)
}
func (m *Connection) XXX_DiscardUnknown() {
	xxx_messageInfo_Connection.DiscardUnknown(m)
}

var xxx_messageInfo_Connection proto.InternalMessageInfo

func (m *Connection) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Connection) GetNetworkService() string {
	if m != nil {
		return m.NetworkService
	}
	return ""
}

func (m *Connection) GetMechanism() *Mechanism {
	if m != nil {
		return m.Mechanism
	}
	return nil
}

func (m *Connection) GetContext() *ConnectionContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Connection) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Connection) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *Connection) GetNetworkServiceEndpointName() string {
	if m != nil {
		return m.NetworkServiceEndpointName
	}
	return ""
}

func (m *Connection) GetState() State {
	if m != nil {
		return m.State
	}
	return State_UP
}

type ConnectionEvent struct {
	Type                 ConnectionEventType    `protobuf:"varint,1,opt,name=type,proto3,enum=connection.ConnectionEventType" json:"type,omitempty"`
	Connections          map[string]*Connection `protobuf:"bytes,2,rep,name=connections,proto3" json:"connections,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ConnectionEvent) Reset()         { *m = ConnectionEvent{} }
func (m *ConnectionEvent) String() string { return proto.CompactTextString(m) }
func (*ConnectionEvent) ProtoMessage()    {}
func (*ConnectionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{4}
}

func (m *ConnectionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionEvent.Unmarshal(m, b)
}
func (m *ConnectionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionEvent.Marshal(b, m, deterministic)
}
func (m *ConnectionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionEvent.Merge(m, src)
}
func (m *ConnectionEvent) XXX_Size() int {
	return xxx_messageInfo_ConnectionEvent.Size(m)
}
func (m *ConnectionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionEvent proto.InternalMessageInfo

func (m *ConnectionEvent) GetType() ConnectionEventType {
	if m != nil {
		return m.Type
	}
	return ConnectionEventType_INITIAL_STATE_TRANSFER
}

func (m *ConnectionEvent) GetConnections() map[string]*Connection {
	if m != nil {
		return m.Connections
	}
	return nil
}

type MonitorScopeSelector struct {
	PathSegments         []*PathSegment `protobuf:"bytes,1,rep,name=path_segments,json=pathSegments,proto3" json:"path_segments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MonitorScopeSelector) Reset()         { *m = MonitorScopeSelector{} }
func (m *MonitorScopeSelector) String() string { return proto.CompactTextString(m) }
func (*MonitorScopeSelector) ProtoMessage()    {}
func (*MonitorScopeSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{5}
}

func (m *MonitorScopeSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorScopeSelector.Unmarshal(m, b)
}
func (m *MonitorScopeSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorScopeSelector.Marshal(b, m, deterministic)
}
func (m *MonitorScopeSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorScopeSelector.Merge(m, src)
}
func (m *MonitorScopeSelector) XXX_Size() int {
	return xxx_messageInfo_MonitorScopeSelector.Size(m)
}
func (m *MonitorScopeSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorScopeSelector.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorScopeSelector proto.InternalMessageInfo

func (m *MonitorScopeSelector) GetPathSegments() []*PathSegment {
	if m != nil {
		return m.PathSegments
	}
	return nil
}

func init() {
	proto.RegisterEnum("connection.State", State_name, State_value)
	proto.RegisterEnum("connection.ConnectionEventType", ConnectionEventType_name, ConnectionEventType_value)
	proto.RegisterType((*Mechanism)(nil), "connection.Mechanism")
	proto.RegisterMapType((map[string]string)(nil), "connection.Mechanism.ParametersEntry")
	proto.RegisterType((*PathSegment)(nil), "connection.PathSegment")
	proto.RegisterMapType((map[string]string)(nil), "connection.PathSegment.MetricsEntry")
	proto.RegisterType((*Path)(nil), "connection.Path")
	proto.RegisterType((*Connection)(nil), "connection.Connection")
	proto.RegisterMapType((map[string]string)(nil), "connection.Connection.LabelsEntry")
	proto.RegisterType((*ConnectionEvent)(nil), "connection.ConnectionEvent")
	proto.RegisterMapType((map[string]*Connection)(nil), "connection.ConnectionEvent.ConnectionsEntry")
	proto.RegisterType((*MonitorScopeSelector)(nil), "connection.MonitorScopeSelector")
}

func init() { proto.RegisterFile("connection.proto", fileDescriptor_51baa40a1cc6b48b) }

var fileDescriptor_51baa40a1cc6b48b = []byte{
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x4e, 0xdb, 0x4a,
	0x10, 0xc6, 0xf9, 0xe3, 0x64, 0xc2, 0x8f, 0xd9, 0xc3, 0x01, 0x1f, 0x57, 0x55, 0xa3, 0x88, 0x8a,
	0x08, 0x21, 0xa7, 0x0a, 0xbd, 0xa0, 0x51, 0x5b, 0x29, 0x05, 0x57, 0x8a, 0x04, 0x69, 0xe4, 0x98,
	0x56, 0xe2, 0x26, 0x72, 0x9c, 0x69, 0xe2, 0x26, 0xf6, 0x5a, 0xf6, 0x42, 0xc9, 0x93, 0xf5, 0x01,
	0xfa, 0x3e, 0xbd, 0xef, 0x5d, 0xe5, 0xb5, 0x1d, 0x2f, 0x21, 0x45, 0xe2, 0xca, 0xbb, 0x33, 0xdf,
	0xec, 0xcc, 0x7c, 0xf3, 0x8d, 0x41, 0xb6, 0xa9, 0xe7, 0xa1, 0xcd, 0x1c, 0xea, 0x69, 0x7e, 0x40,
	0x19, 0x25, 0x90, 0x59, 0xd4, 0xaa, 0xcf, 0xe6, 0x3e, 0x86, 0x0d, 0xe6, 0xb8, 0x18, 0x32, 0xcb,
	0xf5, 0xb3, 0x53, 0x8c, 0x56, 0xf7, 0x33, 0xb4, 0x4d, 0x3d, 0x86, 0x77, 0x2c, 0x76, 0xd4, 0x7e,
	0x48, 0x50, 0xbe, 0x44, 0x7b, 0x62, 0x79, 0x4e, 0xe8, 0x12, 0x19, 0xf2, 0xf6, 0x2c, 0x54, 0xa4,
	0xaa, 0x54, 0x2f, 0x1b, 0xd1, 0x91, 0x10, 0x28, 0x44, 0x6f, 0x2b, 0x39, 0x6e, 0xe2, 0x67, 0xa2,
	0x03, 0xf8, 0x56, 0x60, 0xb9, 0xc8, 0x30, 0x08, 0x95, 0x7c, 0x35, 0x5f, 0xaf, 0x34, 0x5f, 0x6a,
	0x42, 0x85, 0x8b, 0x07, 0xb5, 0xde, 0x02, 0xa7, 0x7b, 0x2c, 0x98, 0x1b, 0x42, 0xa0, 0xfa, 0x0e,
	0xb6, 0x97, 0xdc, 0x51, 0xfe, 0x29, 0xce, 0xd3, 0xfc, 0x53, 0x9c, 0x93, 0x5d, 0x28, 0xde, 0x5a,
	0xb3, 0x9b, 0xb4, 0x80, 0xf8, 0xd2, 0xca, 0x9d, 0x4a, 0xb5, 0xdf, 0x12, 0x54, 0x7a, 0x16, 0x9b,
	0xf4, 0x71, 0xec, 0xa2, 0xc7, 0xa2, 0x4a, 0x3d, 0xcb, 0xc5, 0x24, 0x98, 0x9f, 0xc9, 0x16, 0xe4,
	0x9c, 0x51, 0x12, 0x9a, 0x73, 0x46, 0xd1, 0x6b, 0x8c, 0x4e, 0xd1, 0x53, 0xf2, 0xf1, 0x6b, 0xfc,
	0x42, 0x5e, 0xc3, 0x3a, 0xde, 0xf9, 0x4e, 0x80, 0xa1, 0x52, 0xa8, 0x4a, 0xf5, 0x4a, 0x53, 0xd5,
	0xc6, 0x94, 0x8e, 0x67, 0x18, 0x73, 0x34, 0xbc, 0xf9, 0xaa, 0x99, 0x29, 0x9f, 0x46, 0x0a, 0x25,
	0xef, 0x61, 0xdd, 0x45, 0x16, 0x38, 0x76, 0xa8, 0x14, 0x39, 0x05, 0x07, 0x22, 0x05, 0x42, 0x65,
	0xda, 0x65, 0x0c, 0x8b, 0x19, 0x48, 0x83, 0xd4, 0x16, 0x6c, 0x88, 0x8e, 0x27, 0xf5, 0x7e, 0x0d,
	0x85, 0x28, 0x41, 0x84, 0x70, 0xbc, 0x11, 0xde, 0xf1, 0xa8, 0x4d, 0x23, 0xbe, 0x90, 0xb7, 0xb0,
	0xe9, 0x5b, 0x6c, 0x32, 0x08, 0xe3, 0xfc, 0xa1, 0x92, 0xe3, 0xf5, 0xed, 0xff, 0xa5, 0x3e, 0x63,
	0xc3, 0xcf, 0x2e, 0x61, 0xed, 0x67, 0x1e, 0xe0, 0x6c, 0x01, 0x4c, 0x28, 0x94, 0x16, 0x14, 0x1e,
	0xc2, 0xb6, 0x87, 0xec, 0x3b, 0x0d, 0xa6, 0x83, 0x10, 0x83, 0x5b, 0xc7, 0x4e, 0xcb, 0xdb, 0x4a,
	0xcc, 0xfd, 0xd8, 0x4a, 0x4e, 0xa0, 0xec, 0xa6, 0x3a, 0xe0, 0x7c, 0x57, 0x9a, 0xff, 0xad, 0x14,
	0x89, 0x91, 0xe1, 0x22, 0x52, 0x13, 0x7d, 0x26, 0xa3, 0x10, 0x49, 0x4d, 0x95, 0x9b, 0x55, 0x77,
	0x16, 0x5b, 0x8c, 0x34, 0x88, 0xb4, 0xa0, 0x34, 0xb3, 0x86, 0x38, 0x4b, 0x67, 0x52, 0x13, 0x33,
	0x66, 0x71, 0xda, 0x05, 0x07, 0xc5, 0x13, 0x49, 0x22, 0xc8, 0x01, 0x14, 0x22, 0x22, 0x94, 0x12,
	0x4f, 0x2c, 0x2f, 0xb3, 0x65, 0x70, 0x2f, 0x69, 0xc3, 0xf3, 0xa5, 0xfe, 0x07, 0xe8, 0x8d, 0x7c,
	0xea, 0x78, 0x6c, 0xc0, 0xf5, 0xb7, 0xce, 0xd9, 0x50, 0xef, 0xb3, 0xa1, 0x27, 0x90, 0x6e, 0xa4,
	0xca, 0x43, 0x28, 0x86, 0xcc, 0x62, 0xa8, 0x94, 0xab, 0x52, 0x7d, 0xab, 0xb9, 0x23, 0x66, 0xea,
	0x47, 0x0e, 0x23, 0xf6, 0xab, 0x6f, 0xa0, 0x22, 0x14, 0xfa, 0x24, 0x85, 0xfc, 0x92, 0x60, 0x3b,
	0xeb, 0x57, 0xbf, 0x8d, 0x36, 0xe4, 0x24, 0xd9, 0x65, 0x89, 0xa7, 0x7d, 0xb1, 0x9a, 0x1a, 0x0e,
	0x35, 0xe7, 0x3e, 0x26, 0xcb, 0xde, 0x85, 0x4a, 0x86, 0x4b, 0xa5, 0x74, 0xfc, 0x48, 0xac, 0x70,
	0x4f, 0x08, 0x16, 0x1f, 0x50, 0x3f, 0x83, 0xbc, 0x0c, 0x58, 0xd1, 0xd8, 0xb1, 0xd8, 0x58, 0xa5,
	0xb9, 0xb7, 0x3a, 0x9f, 0xd8, 0xb0, 0x09, 0xbb, 0x97, 0xd4, 0x73, 0x18, 0x0d, 0xfa, 0x36, 0xf5,
	0xb1, 0x8f, 0x33, 0xb4, 0x19, 0x0d, 0x1e, 0x2e, 0x83, 0xf4, 0x84, 0x65, 0x38, 0xfa, 0x1f, 0x8a,
	0x7c, 0x22, 0xa4, 0x04, 0xb9, 0xab, 0x9e, 0xbc, 0x46, 0xfe, 0x81, 0xc2, 0xf9, 0xa7, 0x2f, 0x5d,
	0x59, 0x3a, 0xea, 0xc0, 0xbf, 0x2b, 0x58, 0x23, 0x2a, 0xec, 0x75, 0xba, 0x1d, 0xb3, 0xd3, 0xbe,
	0x18, 0xf4, 0xcd, 0xb6, 0xa9, 0x0f, 0x4c, 0xa3, 0xdd, 0xed, 0x7f, 0xd4, 0x0d, 0x79, 0x8d, 0x00,
	0x94, 0xae, 0x7a, 0xe7, 0x6d, 0x53, 0x97, 0xa5, 0xe8, 0x7c, 0xae, 0x5f, 0xe8, 0xa6, 0x2e, 0xe7,
	0x9a, 0xdf, 0x60, 0x27, 0xa9, 0x5d, 0x58, 0xbc, 0x2b, 0x20, 0x0f, 0x8c, 0x21, 0xa9, 0xde, 0x5b,
	0xa1, 0x15, 0x0d, 0xab, 0xcf, 0x1e, 0x99, 0xcd, 0x2b, 0xe9, 0x43, 0xeb, 0xfa, 0x74, 0xec, 0xb0,
	0xc9, 0xcd, 0x50, 0xb3, 0xa9, 0xdb, 0x48, 0x54, 0x9a, 0x28, 0xd9, 0xc5, 0x70, 0xd2, 0xb0, 0x7c,
	0xa7, 0xe1, 0x4f, 0xc7, 0xfc, 0x7b, 0xdf, 0x3d, 0x2c, 0xf1, 0xff, 0xe1, 0xc9, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xac, 0xa5, 0x75, 0x12, 0x8e, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MonitorConnectionClient is the client API for MonitorConnection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitorConnectionClient interface {
	MonitorConnections(ctx context.Context, in *MonitorScopeSelector, opts ...grpc.CallOption) (MonitorConnection_MonitorConnectionsClient, error)
}

type monitorConnectionClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitorConnectionClient(cc grpc.ClientConnInterface) MonitorConnectionClient {
	return &monitorConnectionClient{cc}
}

func (c *monitorConnectionClient) MonitorConnections(ctx context.Context, in *MonitorScopeSelector, opts ...grpc.CallOption) (MonitorConnection_MonitorConnectionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MonitorConnection_serviceDesc.Streams[0], "/connection.MonitorConnection/MonitorConnections", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitorConnectionMonitorConnectionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MonitorConnection_MonitorConnectionsClient interface {
	Recv() (*ConnectionEvent, error)
	grpc.ClientStream
}

type monitorConnectionMonitorConnectionsClient struct {
	grpc.ClientStream
}

func (x *monitorConnectionMonitorConnectionsClient) Recv() (*ConnectionEvent, error) {
	m := new(ConnectionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MonitorConnectionServer is the server API for MonitorConnection service.
type MonitorConnectionServer interface {
	MonitorConnections(*MonitorScopeSelector, MonitorConnection_MonitorConnectionsServer) error
}

// UnimplementedMonitorConnectionServer can be embedded to have forward compatible implementations.
type UnimplementedMonitorConnectionServer struct {
}

func (*UnimplementedMonitorConnectionServer) MonitorConnections(req *MonitorScopeSelector, srv MonitorConnection_MonitorConnectionsServer) error {
	return status.Errorf(codes.Unimplemented, "method MonitorConnections not implemented")
}

func RegisterMonitorConnectionServer(s *grpc.Server, srv MonitorConnectionServer) {
	s.RegisterService(&_MonitorConnection_serviceDesc, srv)
}

func _MonitorConnection_MonitorConnections_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MonitorScopeSelector)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MonitorConnectionServer).MonitorConnections(m, &monitorConnectionMonitorConnectionsServer{stream})
}

type MonitorConnection_MonitorConnectionsServer interface {
	Send(*ConnectionEvent) error
	grpc.ServerStream
}

type monitorConnectionMonitorConnectionsServer struct {
	grpc.ServerStream
}

func (x *monitorConnectionMonitorConnectionsServer) Send(m *ConnectionEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _MonitorConnection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connection.MonitorConnection",
	HandlerType: (*MonitorConnectionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorConnections",
			Handler:       _MonitorConnection_MonitorConnections_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "connection.proto",
}
