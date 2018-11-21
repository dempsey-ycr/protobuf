// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

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
type ObjectType int32

const (
	ObjectType_OBJTYPE_NULL            ObjectType = 0
	ObjectType_OBJTYPE_NATURALPERSON   ObjectType = 1
	ObjectType_OBJTYPE_LEGALPERSON     ObjectType = 2
	ObjectType_OBJTYPE_HOUSEPROPERTY   ObjectType = 3
	ObjectType_OBJTYPE_PROJECT_ATO     ObjectType = 4
	ObjectType_OBJTYPE_BOUNDARY        ObjectType = 5
	ObjectType_OBJTYPE_ACTION_ATOTRANS ObjectType = 6
	ObjectType_OBJTYPE_ACTION_BITSDAQ  ObjectType = 7
	ObjectType_OBJTYPE_MAX             ObjectType = 8
)

var ObjectType_name = map[int32]string{
	0: "OBJTYPE_NULL",
	1: "OBJTYPE_NATURALPERSON",
	2: "OBJTYPE_LEGALPERSON",
	3: "OBJTYPE_HOUSEPROPERTY",
	4: "OBJTYPE_PROJECT_ATO",
	5: "OBJTYPE_BOUNDARY",
	6: "OBJTYPE_ACTION_ATOTRANS",
	7: "OBJTYPE_ACTION_BITSDAQ",
	8: "OBJTYPE_MAX",
}

var ObjectType_value = map[string]int32{
	"OBJTYPE_NULL":            0,
	"OBJTYPE_NATURALPERSON":   1,
	"OBJTYPE_LEGALPERSON":     2,
	"OBJTYPE_HOUSEPROPERTY":   3,
	"OBJTYPE_PROJECT_ATO":     4,
	"OBJTYPE_BOUNDARY":        5,
	"OBJTYPE_ACTION_ATOTRANS": 6,
	"OBJTYPE_ACTION_BITSDAQ":  7,
	"OBJTYPE_MAX":             8,
}

func (x ObjectType) String() string {
	return proto.EnumName(ObjectType_name, int32(x))
}

func (ObjectType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
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
	return fileDescriptor_8f954d82c0b891f6, []int{0}
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
	return fileDescriptor_8f954d82c0b891f6, []int{1}
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
	return fileDescriptor_8f954d82c0b891f6, []int{2}
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
	return fileDescriptor_8f954d82c0b891f6, []int{3}
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
	proto.RegisterEnum("common.ObjectType", ObjectType_name, ObjectType_value)
	proto.RegisterType((*RequestByCond)(nil), "common.RequestByCond")
	proto.RegisterType((*Response)(nil), "common.Response")
	proto.RegisterType((*FabricPrivate)(nil), "common.FabricPrivate")
	proto.RegisterType((*Endorsement)(nil), "common.Endorsement")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x71, 0x6e, 0x6d, 0x4f, 0x1a, 0x3a, 0x3a, 0xbd, 0x99, 0xcb, 0xc2, 0x8a, 0x58, 0x58,
	0x5d, 0x14, 0xa9, 0x88, 0x15, 0x2b, 0x27, 0x31, 0x90, 0x2a, 0xc4, 0x66, 0xe2, 0x20, 0xb2, 0xaa,
	0x26, 0xf6, 0x50, 0x8c, 0xea, 0x99, 0xe0, 0x99, 0x14, 0xf2, 0x14, 0x3c, 0x01, 0x2f, 0xc3, 0xc3,
	0xf0, 0x1c, 0xc8, 0x8e, 0x9d, 0x38, 0x65, 0xe5, 0xf9, 0xbf, 0xff, 0x9c, 0x5f, 0xe7, 0x8c, 0x6d,
	0x38, 0x0e, 0x65, 0x92, 0x48, 0xf1, 0x72, 0xfd, 0xb8, 0x5c, 0xa4, 0x52, 0x4b, 0x6c, 0xad, 0x55,
	0x77, 0x08, 0x1d, 0xca, 0xbf, 0x2f, 0xb9, 0xd2, 0xbd, 0x55, 0x5f, 0x8a, 0x08, 0x11, 0x1a, 0x7a,
	0xb5, 0xe0, 0xa6, 0x61, 0x19, 0x76, 0x93, 0xe6, 0x67, 0x7c, 0x0c, 0xb5, 0x38, 0x32, 0x6b, 0x96,
	0x61, 0x1f, 0xd0, 0x5a, 0x1c, 0xe1, 0x09, 0x34, 0xe5, 0x0f, 0xc1, 0x53, 0xb3, 0x9e, 0xa3, 0xb5,
	0xe8, 0xfe, 0x32, 0x60, 0x9f, 0x72, 0xb5, 0x90, 0x42, 0xf1, 0x2c, 0x26, 0x94, 0xd1, 0x26, 0x26,
	0x3b, 0xa3, 0x09, 0x7b, 0x09, 0x57, 0x8a, 0xdd, 0xf2, 0x22, 0xab, 0x94, 0x99, 0xb3, 0x60, 0xab,
	0x3b, 0xc9, 0xa2, 0x3c, 0xf2, 0x90, 0x96, 0x12, 0xdf, 0x40, 0xe7, 0x0b, 0x9b, 0xa7, 0x71, 0xe8,
	0xa7, 0xf1, 0x3d, 0xd3, 0xdc, 0x6c, 0x58, 0x86, 0xdd, 0xbe, 0x3a, 0xbd, 0x2c, 0xb6, 0x79, 0x5b,
	0x35, 0xe9, 0x6e, 0x6d, 0xf7, 0x8f, 0x01, 0x9d, 0x9d, 0x02, 0xb4, 0xe1, 0x28, 0xfc, 0xca, 0x62,
	0x91, 0xcd, 0x33, 0xd1, 0x4c, 0x2f, 0x55, 0x31, 0xe1, 0x43, 0x8c, 0x2f, 0xa0, 0xa3, 0x53, 0x26,
	0x14, 0x0b, 0x75, 0x2c, 0xc5, 0x70, 0x50, 0x8c, 0xbc, 0x0b, 0xf1, 0x02, 0x88, 0xfe, 0xf9, 0x89,
	0xdd, 0xc5, 0x11, 0xcb, 0x48, 0x3f, 0x5b, 0xb9, 0x9e, 0x07, 0xfe, 0xc7, 0xf1, 0x35, 0xb4, 0xb9,
	0x88, 0x64, 0xaa, 0x78, 0xc2, 0x85, 0x36, 0x1b, 0x56, 0xdd, 0x6e, 0x5f, 0x1d, 0x97, 0x8b, 0xb8,
	0x5b, 0x8b, 0x56, 0xeb, 0xba, 0xbf, 0x0d, 0x68, 0x57, 0x4c, 0x3c, 0x83, 0x96, 0xaa, 0x4e, 0x5e,
	0x28, 0xb4, 0x36, 0xf1, 0x4e, 0x14, 0xa5, 0xc5, 0xb8, 0x55, 0x84, 0xcf, 0xe1, 0x40, 0xc5, 0xb7,
	0x82, 0xe9, 0x65, 0xca, 0x8b, 0x7b, 0xde, 0x82, 0xec, 0x1d, 0xdc, 0xf3, 0x54, 0xc5, 0x52, 0xe4,
	0x77, 0xdc, 0xa4, 0xa5, 0xcc, 0xfa, 0x74, 0x9c, 0x70, 0xa5, 0x59, 0xb2, 0x30, 0x9b, 0x96, 0x61,
	0xd7, 0xe9, 0x16, 0x5c, 0xfc, 0x35, 0x00, 0xbc, 0xf9, 0x37, 0x1e, 0xea, 0x20, 0xfb, 0x56, 0x08,
	0x1c, 0x7a, 0xbd, 0xeb, 0x60, 0xe6, 0xbb, 0x37, 0xe3, 0xe9, 0x68, 0x44, 0x1e, 0xe1, 0x13, 0x38,
	0xdd, 0x10, 0x27, 0x98, 0x52, 0x67, 0xe4, 0xbb, 0x74, 0xe2, 0x8d, 0x89, 0x81, 0xe7, 0x70, 0x5c,
	0x5a, 0x23, 0xf7, 0xdd, 0xc6, 0xa8, 0x55, 0x7b, 0xde, 0x7b, 0xd3, 0x89, 0xeb, 0x53, 0xcf, 0x77,
	0x69, 0x30, 0x23, 0xf5, 0x6a, 0x8f, 0x4f, 0xbd, 0x6b, 0xb7, 0x1f, 0xdc, 0x38, 0x81, 0x47, 0x1a,
	0x78, 0x02, 0xa4, 0x34, 0x7a, 0xde, 0x74, 0x3c, 0x70, 0xe8, 0x8c, 0x34, 0xf1, 0x19, 0x9c, 0x97,
	0xd4, 0xe9, 0x07, 0x43, 0x6f, 0x9c, 0x55, 0x07, 0xd4, 0x19, 0x4f, 0x48, 0x0b, 0x9f, 0xc2, 0xd9,
	0x03, 0xb3, 0x37, 0x0c, 0x26, 0x03, 0xe7, 0x23, 0xd9, 0xc3, 0x23, 0x68, 0x97, 0xde, 0x07, 0xe7,
	0x33, 0xd9, 0x9f, 0xb7, 0xf2, 0x3f, 0xe7, 0xd5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x70, 0x0d,
	0xc0, 0x10, 0x50, 0x03, 0x00, 0x00,
}