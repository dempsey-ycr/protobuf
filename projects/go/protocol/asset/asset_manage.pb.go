// Code generated by protoc-gen-go. DO NOT EDIT.
// source: asset/asset_manage.proto

package asset

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

// 资产类型
type EnumAssetType int32

const (
	EnumAssetType_IHT_ASSETMANAGE_ESTATETYPE_ALL    EnumAssetType = 0
	EnumAssetType_IHT_ASSETMANAGE_ESTATETYPE_VILLA  EnumAssetType = 1
	EnumAssetType_IHT_ASSETMANAGE_ESTATETYPE_NOW    EnumAssetType = 2
	EnumAssetType_IHT_ASSETMANAGE_ESTATETYPE_GROUND EnumAssetType = 3
	EnumAssetType_IHT_ASSETMANAGE_ESTATETYPE_FUTURE EnumAssetType = 4
)

var EnumAssetType_name = map[int32]string{
	0: "IHT_ASSETMANAGE_ESTATETYPE_ALL",
	1: "IHT_ASSETMANAGE_ESTATETYPE_VILLA",
	2: "IHT_ASSETMANAGE_ESTATETYPE_NOW",
	3: "IHT_ASSETMANAGE_ESTATETYPE_GROUND",
	4: "IHT_ASSETMANAGE_ESTATETYPE_FUTURE",
}

var EnumAssetType_value = map[string]int32{
	"IHT_ASSETMANAGE_ESTATETYPE_ALL":    0,
	"IHT_ASSETMANAGE_ESTATETYPE_VILLA":  1,
	"IHT_ASSETMANAGE_ESTATETYPE_NOW":    2,
	"IHT_ASSETMANAGE_ESTATETYPE_GROUND": 3,
	"IHT_ASSETMANAGE_ESTATETYPE_FUTURE": 4,
}

func (x EnumAssetType) String() string {
	return proto.EnumName(EnumAssetType_name, int32(x))
}

func (EnumAssetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_deea2318e0bb06f2, []int{0}
}

// 资产状态
type EnumAssetStatus int32

const (
	EnumAssetStatus_IHT_ASSETMANAGE_ASSETSTATUS_ALL     EnumAssetStatus = 0
	EnumAssetStatus_IHT_ASSETMANAGE_ASSETSTATUS_BEFORE  EnumAssetStatus = 11
	EnumAssetStatus_IHT_ASSETMANAGE_ASSETSTATUS_ING     EnumAssetStatus = 12
	EnumAssetStatus_IHT_ASSETMANAGE_ASSETSTATUS_FAILED  EnumAssetStatus = 13
	EnumAssetStatus_IHT_ASSETMANAGE_ASSETSTATUS_SUCCEED EnumAssetStatus = 14
)

var EnumAssetStatus_name = map[int32]string{
	0:  "IHT_ASSETMANAGE_ASSETSTATUS_ALL",
	11: "IHT_ASSETMANAGE_ASSETSTATUS_BEFORE",
	12: "IHT_ASSETMANAGE_ASSETSTATUS_ING",
	13: "IHT_ASSETMANAGE_ASSETSTATUS_FAILED",
	14: "IHT_ASSETMANAGE_ASSETSTATUS_SUCCEED",
}

var EnumAssetStatus_value = map[string]int32{
	"IHT_ASSETMANAGE_ASSETSTATUS_ALL":     0,
	"IHT_ASSETMANAGE_ASSETSTATUS_BEFORE":  11,
	"IHT_ASSETMANAGE_ASSETSTATUS_ING":     12,
	"IHT_ASSETMANAGE_ASSETSTATUS_FAILED":  13,
	"IHT_ASSETMANAGE_ASSETSTATUS_SUCCEED": 14,
}

func (x EnumAssetStatus) String() string {
	return proto.EnumName(EnumAssetStatus_name, int32(x))
}

func (EnumAssetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_deea2318e0bb06f2, []int{1}
}

// DataBlockBase 基础结构
type DataBlockBase struct {
	DocTag               string   `protobuf:"bytes,1,opt,name=docTag,proto3" json:"docTag,omitempty"`
	DbType               int32    `protobuf:"varint,2,opt,name=dbType,proto3" json:"dbType,omitempty"`
	DbProfile            string   `protobuf:"bytes,3,opt,name=dbProfile,proto3" json:"dbProfile,omitempty"`
	CreateTime           string   `protobuf:"bytes,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataBlockBase) Reset()         { *m = DataBlockBase{} }
func (m *DataBlockBase) String() string { return proto.CompactTextString(m) }
func (*DataBlockBase) ProtoMessage()    {}
func (*DataBlockBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_deea2318e0bb06f2, []int{0}
}

func (m *DataBlockBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBlockBase.Unmarshal(m, b)
}
func (m *DataBlockBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBlockBase.Marshal(b, m, deterministic)
}
func (m *DataBlockBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBlockBase.Merge(m, src)
}
func (m *DataBlockBase) XXX_Size() int {
	return xxx_messageInfo_DataBlockBase.Size(m)
}
func (m *DataBlockBase) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBlockBase.DiscardUnknown(m)
}

var xxx_messageInfo_DataBlockBase proto.InternalMessageInfo

func (m *DataBlockBase) GetDocTag() string {
	if m != nil {
		return m.DocTag
	}
	return ""
}

func (m *DataBlockBase) GetDbType() int32 {
	if m != nil {
		return m.DbType
	}
	return 0
}

func (m *DataBlockBase) GetDbProfile() string {
	if m != nil {
		return m.DbProfile
	}
	return ""
}

func (m *DataBlockBase) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *DataBlockBase) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

// 文件结构信息
type DocumentStore struct {
	Dtype                int32    `protobuf:"varint,1,opt,name=dtype,proto3" json:"dtype,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Hash                 string   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocumentStore) Reset()         { *m = DocumentStore{} }
func (m *DocumentStore) String() string { return proto.CompactTextString(m) }
func (*DocumentStore) ProtoMessage()    {}
func (*DocumentStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_deea2318e0bb06f2, []int{1}
}

func (m *DocumentStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentStore.Unmarshal(m, b)
}
func (m *DocumentStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentStore.Marshal(b, m, deterministic)
}
func (m *DocumentStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentStore.Merge(m, src)
}
func (m *DocumentStore) XXX_Size() int {
	return xxx_messageInfo_DocumentStore.Size(m)
}
func (m *DocumentStore) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentStore.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentStore proto.InternalMessageInfo

func (m *DocumentStore) GetDtype() int32 {
	if m != nil {
		return m.Dtype
	}
	return 0
}

func (m *DocumentStore) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *DocumentStore) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterEnum("asset.EnumAssetType", EnumAssetType_name, EnumAssetType_value)
	proto.RegisterEnum("asset.EnumAssetStatus", EnumAssetStatus_name, EnumAssetStatus_value)
	proto.RegisterType((*DataBlockBase)(nil), "asset.DataBlockBase")
	proto.RegisterType((*DocumentStore)(nil), "asset.DocumentStore")
}

func init() { proto.RegisterFile("asset/asset_manage.proto", fileDescriptor_deea2318e0bb06f2) }

var fileDescriptor_deea2318e0bb06f2 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0xae, 0x93, 0x40,
	0x14, 0x76, 0xee, 0x2d, 0x37, 0xe9, 0x51, 0x74, 0x32, 0x31, 0x86, 0x85, 0xa9, 0x95, 0xfa, 0xd3,
	0x74, 0xa1, 0x0b, 0x9f, 0x60, 0x5a, 0xa6, 0x95, 0x04, 0x69, 0x85, 0x41, 0xe3, 0x8a, 0x4c, 0x61,
	0x6c, 0x1b, 0x4b, 0x69, 0x60, 0x58, 0xf8, 0x30, 0x3e, 0x91, 0x5b, 0x1f, 0xc8, 0x30, 0x50, 0x6b,
	0xee, 0x02, 0x36, 0xe4, 0x7c, 0x3f, 0xe7, 0xe4, 0xfb, 0xc8, 0x80, 0x25, 0xca, 0x52, 0xaa, 0xf7,
	0xfa, 0x1b, 0x67, 0xe2, 0x24, 0x76, 0xf2, 0xdd, 0xb9, 0xc8, 0x55, 0x4e, 0x0c, 0xcd, 0xd9, 0xbf,
	0x10, 0x98, 0x8e, 0x50, 0x62, 0x7e, 0xcc, 0x93, 0x1f, 0x73, 0x51, 0x4a, 0xf2, 0x0c, 0xee, 0xd2,
	0x3c, 0xe1, 0x62, 0x67, 0xa1, 0x31, 0x9a, 0x0e, 0x83, 0x16, 0x69, 0x7e, 0xcb, 0x7f, 0x9e, 0xa5,
	0x75, 0x33, 0x46, 0x53, 0x23, 0x68, 0x11, 0x79, 0x0e, 0xc3, 0x74, 0xbb, 0x29, 0xf2, 0xef, 0x87,
	0xa3, 0xb4, 0x6e, 0xf5, 0xca, 0x95, 0x20, 0x23, 0x80, 0xa4, 0x90, 0x42, 0x49, 0x7e, 0xc8, 0xa4,
	0x35, 0xd0, 0xf2, 0x7f, 0x4c, 0xad, 0x57, 0xe7, 0xf4, 0xa2, 0x1b, 0x8d, 0x7e, 0x65, 0xec, 0xcf,
	0x60, 0x3a, 0x79, 0x52, 0x65, 0xf2, 0xa4, 0x42, 0x95, 0x17, 0x92, 0x3c, 0x05, 0x23, 0x55, 0x75,
	0x0a, 0xa4, 0x53, 0x34, 0xa0, 0x0e, 0x57, 0x2a, 0xa1, 0xaa, 0xf2, 0x12, 0xae, 0x41, 0x84, 0xc0,
	0x60, 0x2f, 0xca, 0x7d, 0x9b, 0x4b, 0xcf, 0xb3, 0xdf, 0x08, 0x4c, 0x76, 0xaa, 0x32, 0x5a, 0xff,
	0x00, 0x5d, 0xc1, 0x86, 0x91, 0xfb, 0x91, 0xc7, 0x34, 0x0c, 0x19, 0xff, 0x44, 0x7d, 0xba, 0x62,
	0x31, 0x0b, 0x39, 0xe5, 0x8c, 0x7f, 0xdb, 0xb0, 0x98, 0x7a, 0x1e, 0x7e, 0x40, 0x5e, 0xc1, 0xb8,
	0xc3, 0xf3, 0xc5, 0xf5, 0x3c, 0x8a, 0x51, 0xcf, 0x25, 0x7f, 0xfd, 0x15, 0xdf, 0x90, 0xd7, 0xf0,
	0xb2, 0xc3, 0xb3, 0x0a, 0xd6, 0x91, 0xef, 0xe0, 0xdb, 0x1e, 0xdb, 0x32, 0xe2, 0x51, 0xc0, 0xf0,
	0x60, 0xf6, 0x07, 0xc1, 0x93, 0x7f, 0x6d, 0xc2, 0xa6, 0xf5, 0x04, 0x5e, 0xdc, 0x5f, 0xd5, 0x73,
	0xbd, 0x1e, 0x85, 0x6d, 0xa1, 0x37, 0x60, 0x77, 0x99, 0xe6, 0x6c, 0xb9, 0x0e, 0x18, 0x7e, 0xd8,
	0x77, 0xcc, 0xf5, 0x57, 0xf8, 0x51, 0xdf, 0xb1, 0x25, 0x75, 0x3d, 0xe6, 0x60, 0x93, 0xbc, 0x85,
	0x49, 0x97, 0x2f, 0x8c, 0x16, 0x0b, 0xc6, 0x1c, 0xfc, 0x78, 0x7b, 0xa7, 0x5f, 0xe9, 0x87, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x0f, 0x0b, 0x0d, 0xc1, 0x02, 0x00, 0x00,
}
