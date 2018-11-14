// Code generated by protoc-gen-go. DO NOT EDIT.
// source: basic/common.proto

package basic

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 基本对象类型
type BasicObjectType int32

const (
	BasicObjectType_OBJTYPE_NULL          BasicObjectType = 0
	BasicObjectType_OBJTYPE_NATURALPERSON BasicObjectType = 1
	BasicObjectType_OBJTYPE_LEGALPERSON   BasicObjectType = 2
	BasicObjectType_OBJTYPE_HOUSEPROPERTY BasicObjectType = 3
	BasicObjectType_OBJTYPE_PROJECT_ATO   BasicObjectType = 4
	BasicObjectType_OBJTYPE_MAX           BasicObjectType = 5
)

var BasicObjectType_name = map[int32]string{
	0: "OBJTYPE_NULL",
	1: "OBJTYPE_NATURALPERSON",
	2: "OBJTYPE_LEGALPERSON",
	3: "OBJTYPE_HOUSEPROPERTY",
	4: "OBJTYPE_PROJECT_ATO",
	5: "OBJTYPE_MAX",
}

var BasicObjectType_value = map[string]int32{
	"OBJTYPE_NULL":          0,
	"OBJTYPE_NATURALPERSON": 1,
	"OBJTYPE_LEGALPERSON":   2,
	"OBJTYPE_HOUSEPROPERTY": 3,
	"OBJTYPE_PROJECT_ATO":   4,
	"OBJTYPE_MAX":           5,
}

func (x BasicObjectType) String() string {
	return proto.EnumName(BasicObjectType_name, int32(x))
}

func (BasicObjectType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_68345cfab19ab0f8, []int{0}
}

// 请求结构
type RequestByCond struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Owner                string   `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestByCond) Reset()         { *m = RequestByCond{} }
func (m *RequestByCond) String() string { return proto.CompactTextString(m) }
func (*RequestByCond) ProtoMessage()    {}
func (*RequestByCond) Descriptor() ([]byte, []int) {
	return fileDescriptor_68345cfab19ab0f8, []int{0}
}

func (m *RequestByCond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestByCond.Unmarshal(m, b)
}
func (m *RequestByCond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestByCond.Marshal(b, m, deterministic)
}
func (m *RequestByCond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestByCond.Merge(m, src)
}
func (m *RequestByCond) XXX_Size() int {
	return xxx_messageInfo_RequestByCond.Size(m)
}
func (m *RequestByCond) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestByCond.DiscardUnknown(m)
}

var xxx_messageInfo_RequestByCond proto.InternalMessageInfo

func (m *RequestByCond) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *RequestByCond) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RequestByCond) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type Response struct {
	Code                 int32          `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Payload              []byte         `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	FabricPrivate        *FabricPrivate `protobuf:"bytes,4,opt,name=fabricPrivate,proto3" json:"fabricPrivate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_68345cfab19ab0f8, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Response) GetFabricPrivate() *FabricPrivate {
	if m != nil {
		return m.FabricPrivate
	}
	return nil
}

// fabric private response
type FabricPrivate struct {
	ChaincodeStatus      int32          `protobuf:"varint,1,opt,name=chaincodeStatus,proto3" json:"chaincodeStatus,omitempty"`
	TransactionID        string         `protobuf:"bytes,2,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
	TxValidationCode     int32          `protobuf:"varint,3,opt,name=txValidationCode,proto3" json:"txValidationCode,omitempty"`
	Endorsement          []*Endorsement `protobuf:"bytes,4,rep,name=endorsement,proto3" json:"endorsement,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FabricPrivate) Reset()         { *m = FabricPrivate{} }
func (m *FabricPrivate) String() string { return proto.CompactTextString(m) }
func (*FabricPrivate) ProtoMessage()    {}
func (*FabricPrivate) Descriptor() ([]byte, []int) {
	return fileDescriptor_68345cfab19ab0f8, []int{2}
}

func (m *FabricPrivate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FabricPrivate.Unmarshal(m, b)
}
func (m *FabricPrivate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FabricPrivate.Marshal(b, m, deterministic)
}
func (m *FabricPrivate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FabricPrivate.Merge(m, src)
}
func (m *FabricPrivate) XXX_Size() int {
	return xxx_messageInfo_FabricPrivate.Size(m)
}
func (m *FabricPrivate) XXX_DiscardUnknown() {
	xxx_messageInfo_FabricPrivate.DiscardUnknown(m)
}

var xxx_messageInfo_FabricPrivate proto.InternalMessageInfo

func (m *FabricPrivate) GetChaincodeStatus() int32 {
	if m != nil {
		return m.ChaincodeStatus
	}
	return 0
}

func (m *FabricPrivate) GetTransactionID() string {
	if m != nil {
		return m.TransactionID
	}
	return ""
}

func (m *FabricPrivate) GetTxValidationCode() int32 {
	if m != nil {
		return m.TxValidationCode
	}
	return 0
}

func (m *FabricPrivate) GetEndorsement() []*Endorsement {
	if m != nil {
		return m.Endorsement
	}
	return nil
}

// fabric endorsement response
type Endorsement struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	EndorseAddr          string   `protobuf:"bytes,2,opt,name=endorseAddr,proto3" json:"endorseAddr,omitempty"`
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Version              int32    `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	Timestamp            int64    `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endorsement) Reset()         { *m = Endorsement{} }
func (m *Endorsement) String() string { return proto.CompactTextString(m) }
func (*Endorsement) ProtoMessage()    {}
func (*Endorsement) Descriptor() ([]byte, []int) {
	return fileDescriptor_68345cfab19ab0f8, []int{3}
}

func (m *Endorsement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endorsement.Unmarshal(m, b)
}
func (m *Endorsement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endorsement.Marshal(b, m, deterministic)
}
func (m *Endorsement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endorsement.Merge(m, src)
}
func (m *Endorsement) XXX_Size() int {
	return xxx_messageInfo_Endorsement.Size(m)
}
func (m *Endorsement) XXX_DiscardUnknown() {
	xxx_messageInfo_Endorsement.DiscardUnknown(m)
}

var xxx_messageInfo_Endorsement proto.InternalMessageInfo

func (m *Endorsement) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Endorsement) GetEndorseAddr() string {
	if m != nil {
		return m.EndorseAddr
	}
	return ""
}

func (m *Endorsement) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Endorsement) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Endorsement) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("basic.BasicObjectType", BasicObjectType_name, BasicObjectType_value)
	proto.RegisterType((*RequestByCond)(nil), "basic.RequestByCond")
	proto.RegisterType((*Response)(nil), "basic.Response")
	proto.RegisterType((*FabricPrivate)(nil), "basic.FabricPrivate")
	proto.RegisterType((*Endorsement)(nil), "basic.Endorsement")
}

func init() { proto.RegisterFile("basic/common.proto", fileDescriptor_68345cfab19ab0f8) }

var fileDescriptor_68345cfab19ab0f8 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x71, 0xdb, 0x0c, 0x76, 0xba, 0xd2, 0xe8, 0x30, 0x20, 0x48, 0x5c, 0x44, 0x15, 0x17,
	0xd1, 0x2e, 0x8a, 0x34, 0xb8, 0xe2, 0xae, 0x2d, 0x01, 0x36, 0x95, 0x25, 0x72, 0x53, 0xc4, 0xae,
	0x26, 0x37, 0x31, 0xc3, 0x68, 0xb1, 0x43, 0xec, 0x0e, 0xfa, 0x12, 0x3c, 0x01, 0x3c, 0x0c, 0x6f,
	0x86, 0x1c, 0x92, 0x36, 0x85, 0x3b, 0x9f, 0xef, 0x3f, 0xe7, 0xd7, 0x7f, 0x6c, 0x03, 0xae, 0x98,
	0x16, 0xe9, 0xf3, 0x54, 0xe5, 0xb9, 0x92, 0xe3, 0xa2, 0x54, 0x46, 0xa1, 0x53, 0xb1, 0xd1, 0x19,
	0x0c, 0x28, 0xff, 0xba, 0xe6, 0xda, 0x4c, 0x37, 0x33, 0x25, 0x33, 0x44, 0xe8, 0x99, 0x4d, 0xc1,
	0x3d, 0xe2, 0x93, 0xc0, 0xa1, 0xd5, 0x19, 0xef, 0x43, 0x47, 0x64, 0x5e, 0xc7, 0x27, 0xc1, 0x21,
	0xed, 0x88, 0x0c, 0x8f, 0xc1, 0x51, 0xdf, 0x24, 0x2f, 0xbd, 0x6e, 0x85, 0xfe, 0x16, 0xa3, 0x1f,
	0x04, 0xee, 0x51, 0xae, 0x0b, 0x25, 0x35, 0xb7, 0x36, 0xa9, 0xca, 0xb6, 0x36, 0xf6, 0x8c, 0x1e,
	0xdc, 0xcd, 0xb9, 0xd6, 0xec, 0x9a, 0xd7, 0x5e, 0x4d, 0x69, 0x95, 0x82, 0x6d, 0x6e, 0x14, 0xcb,
	0x2a, 0xcb, 0x23, 0xda, 0x94, 0xf8, 0x0a, 0x06, 0x9f, 0xd8, 0xaa, 0x14, 0x69, 0x5c, 0x8a, 0x5b,
	0x66, 0xb8, 0xd7, 0xf3, 0x49, 0xd0, 0x3f, 0x3d, 0x1e, 0x57, 0xf1, 0xc7, 0x6f, 0xda, 0x1a, 0xdd,
	0x6f, 0x1d, 0xfd, 0x26, 0x30, 0xd8, 0x6b, 0xc0, 0x00, 0x86, 0xe9, 0x67, 0x26, 0xa4, 0x8d, 0xb3,
	0x30, 0xcc, 0xac, 0x75, 0x1d, 0xf0, 0x5f, 0x8c, 0xcf, 0x60, 0x60, 0x4a, 0x26, 0x35, 0x4b, 0x8d,
	0x50, 0xf2, 0xec, 0x75, 0x9d, 0x78, 0x1f, 0xe2, 0x09, 0xb8, 0xe6, 0xfb, 0x07, 0x76, 0x23, 0x32,
	0x66, 0xc9, 0xcc, 0x6e, 0xdc, 0xad, 0x0c, 0xff, 0xe3, 0xf8, 0x12, 0xfa, 0x5c, 0x66, 0xaa, 0xd4,
	0x3c, 0xe7, 0xd2, 0x78, 0x3d, 0xbf, 0x1b, 0xf4, 0x4f, 0xb1, 0xde, 0x23, 0xdc, 0x29, 0xb4, 0xdd,
	0x36, 0xfa, 0x45, 0xa0, 0xdf, 0x12, 0xf1, 0x11, 0x1c, 0xe8, 0x76, 0xf0, 0xba, 0x42, 0x7f, 0xeb,
	0x3e, 0xc9, 0xb2, 0xb2, 0x4e, 0xdb, 0x46, 0xf8, 0x14, 0x0e, 0xb5, 0xb8, 0x96, 0xcc, 0xac, 0x4b,
	0x5e, 0xdf, 0xf2, 0x0e, 0xd8, 0x17, 0xb8, 0xe5, 0xa5, 0x16, 0x4a, 0x56, 0x37, 0xec, 0xd0, 0xa6,
	0xb4, 0x73, 0x46, 0xe4, 0x5c, 0x1b, 0x96, 0x17, 0x9e, 0xe3, 0x93, 0xa0, 0x4b, 0x77, 0xe0, 0xe4,
	0x27, 0x81, 0xe1, 0xd4, 0xae, 0x10, 0xad, 0xbe, 0xf0, 0xd4, 0x24, 0xf6, 0xbb, 0xb8, 0x70, 0x14,
	0x4d, 0xcf, 0x93, 0xcb, 0x38, 0xbc, 0xba, 0x58, 0xce, 0xe7, 0xee, 0x1d, 0x7c, 0x02, 0x0f, 0xb7,
	0x64, 0x92, 0x2c, 0xe9, 0x64, 0x1e, 0x87, 0x74, 0x11, 0x5d, 0xb8, 0x04, 0x1f, 0xc3, 0x83, 0x46,
	0x9a, 0x87, 0x6f, 0xb7, 0x42, 0xa7, 0x3d, 0xf3, 0x2e, 0x5a, 0x2e, 0xc2, 0x98, 0x46, 0x71, 0x48,
	0x93, 0x4b, 0xb7, 0xdb, 0x9e, 0x89, 0x69, 0x74, 0x1e, 0xce, 0x92, 0xab, 0x49, 0x12, 0xb9, 0x3d,
	0x1c, 0x42, 0xbf, 0x11, 0xde, 0x4f, 0x3e, 0xba, 0xce, 0xea, 0xa0, 0xfa, 0xec, 0x2f, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x9e, 0x99, 0x9d, 0xf6, 0x02, 0x03, 0x00, 0x00,
}
