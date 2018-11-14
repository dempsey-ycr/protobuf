// Code generated by protoc-gen-go. DO NOT EDIT.
// source: asset/asset.proto

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

// 资产信息
type AssetInfo struct {
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pid                  string           `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Name                 string           `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string           `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Ntype                string           `protobuf:"bytes,5,opt,name=ntype,proto3" json:"ntype,omitempty"`
	Currency             string           `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	FinanceSum           int64            `protobuf:"varint,7,opt,name=financeSum,proto3" json:"financeSum,omitempty"`
	SplitNum             int32            `protobuf:"varint,8,opt,name=splitNum,proto3" json:"splitNum,omitempty"`
	BeginTime            string           `protobuf:"bytes,9,opt,name=beginTime,proto3" json:"beginTime,omitempty"`
	EndTime              string           `protobuf:"bytes,10,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Base                 *DataBlockBase   `protobuf:"bytes,11,opt,name=base,proto3" json:"base,omitempty"`
	QualifiedFile        *DocumentStore   `protobuf:"bytes,12,opt,name=qualifiedFile,proto3" json:"qualifiedFile,omitempty"`
	WhitePaper           *DocumentStore   `protobuf:"bytes,13,opt,name=whitePaper,proto3" json:"whitePaper,omitempty"`
	Pictures             []*DocumentStore `protobuf:"bytes,14,rep,name=pictures,proto3" json:"pictures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AssetInfo) Reset()         { *m = AssetInfo{} }
func (m *AssetInfo) String() string { return proto.CompactTextString(m) }
func (*AssetInfo) ProtoMessage()    {}
func (*AssetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_06b463ba992d3361, []int{0}
}

func (m *AssetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetInfo.Unmarshal(m, b)
}
func (m *AssetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetInfo.Marshal(b, m, deterministic)
}
func (m *AssetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetInfo.Merge(m, src)
}
func (m *AssetInfo) XXX_Size() int {
	return xxx_messageInfo_AssetInfo.Size(m)
}
func (m *AssetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AssetInfo proto.InternalMessageInfo

func (m *AssetInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AssetInfo) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *AssetInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AssetInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AssetInfo) GetNtype() string {
	if m != nil {
		return m.Ntype
	}
	return ""
}

func (m *AssetInfo) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *AssetInfo) GetFinanceSum() int64 {
	if m != nil {
		return m.FinanceSum
	}
	return 0
}

func (m *AssetInfo) GetSplitNum() int32 {
	if m != nil {
		return m.SplitNum
	}
	return 0
}

func (m *AssetInfo) GetBeginTime() string {
	if m != nil {
		return m.BeginTime
	}
	return ""
}

func (m *AssetInfo) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *AssetInfo) GetBase() *DataBlockBase {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *AssetInfo) GetQualifiedFile() *DocumentStore {
	if m != nil {
		return m.QualifiedFile
	}
	return nil
}

func (m *AssetInfo) GetWhitePaper() *DocumentStore {
	if m != nil {
		return m.WhitePaper
	}
	return nil
}

func (m *AssetInfo) GetPictures() []*DocumentStore {
	if m != nil {
		return m.Pictures
	}
	return nil
}

// 数据入链请求结构
type RequestAssetPut struct {
	Key                  string     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Fun                  string     `protobuf:"bytes,2,opt,name=fun,proto3" json:"fun,omitempty"`
	Value                *AssetInfo `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RequestAssetPut) Reset()         { *m = RequestAssetPut{} }
func (m *RequestAssetPut) String() string { return proto.CompactTextString(m) }
func (*RequestAssetPut) ProtoMessage()    {}
func (*RequestAssetPut) Descriptor() ([]byte, []int) {
	return fileDescriptor_06b463ba992d3361, []int{1}
}

func (m *RequestAssetPut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestAssetPut.Unmarshal(m, b)
}
func (m *RequestAssetPut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestAssetPut.Marshal(b, m, deterministic)
}
func (m *RequestAssetPut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestAssetPut.Merge(m, src)
}
func (m *RequestAssetPut) XXX_Size() int {
	return xxx_messageInfo_RequestAssetPut.Size(m)
}
func (m *RequestAssetPut) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestAssetPut.DiscardUnknown(m)
}

var xxx_messageInfo_RequestAssetPut proto.InternalMessageInfo

func (m *RequestAssetPut) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RequestAssetPut) GetFun() string {
	if m != nil {
		return m.Fun
	}
	return ""
}

func (m *RequestAssetPut) GetValue() *AssetInfo {
	if m != nil {
		return m.Value
	}
	return nil
}

// 资产查询返回结构
type ResponseAssetGet struct {
	Code                 int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *AssetInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ResponseAssetGet) Reset()         { *m = ResponseAssetGet{} }
func (m *ResponseAssetGet) String() string { return proto.CompactTextString(m) }
func (*ResponseAssetGet) ProtoMessage()    {}
func (*ResponseAssetGet) Descriptor() ([]byte, []int) {
	return fileDescriptor_06b463ba992d3361, []int{2}
}

func (m *ResponseAssetGet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseAssetGet.Unmarshal(m, b)
}
func (m *ResponseAssetGet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseAssetGet.Marshal(b, m, deterministic)
}
func (m *ResponseAssetGet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseAssetGet.Merge(m, src)
}
func (m *ResponseAssetGet) XXX_Size() int {
	return xxx_messageInfo_ResponseAssetGet.Size(m)
}
func (m *ResponseAssetGet) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseAssetGet.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseAssetGet proto.InternalMessageInfo

func (m *ResponseAssetGet) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseAssetGet) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ResponseAssetGet) GetData() *AssetInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

// 资产富查询请求结构
type RequestAssetRichQuery struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Rtype                int32    `protobuf:"varint,2,opt,name=rtype,proto3" json:"rtype,omitempty"`
	Pid                  string   `protobuf:"bytes,3,opt,name=pid,proto3" json:"pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestAssetRichQuery) Reset()         { *m = RequestAssetRichQuery{} }
func (m *RequestAssetRichQuery) String() string { return proto.CompactTextString(m) }
func (*RequestAssetRichQuery) ProtoMessage()    {}
func (*RequestAssetRichQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_06b463ba992d3361, []int{3}
}

func (m *RequestAssetRichQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestAssetRichQuery.Unmarshal(m, b)
}
func (m *RequestAssetRichQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestAssetRichQuery.Marshal(b, m, deterministic)
}
func (m *RequestAssetRichQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestAssetRichQuery.Merge(m, src)
}
func (m *RequestAssetRichQuery) XXX_Size() int {
	return xxx_messageInfo_RequestAssetRichQuery.Size(m)
}
func (m *RequestAssetRichQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestAssetRichQuery.DiscardUnknown(m)
}

var xxx_messageInfo_RequestAssetRichQuery proto.InternalMessageInfo

func (m *RequestAssetRichQuery) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *RequestAssetRichQuery) GetRtype() int32 {
	if m != nil {
		return m.Rtype
	}
	return 0
}

func (m *RequestAssetRichQuery) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func init() {
	proto.RegisterType((*AssetInfo)(nil), "asset.AssetInfo")
	proto.RegisterType((*RequestAssetPut)(nil), "asset.RequestAssetPut")
	proto.RegisterType((*ResponseAssetGet)(nil), "asset.ResponseAssetGet")
	proto.RegisterType((*RequestAssetRichQuery)(nil), "asset.RequestAssetRichQuery")
}

func init() { proto.RegisterFile("asset/asset.proto", fileDescriptor_06b463ba992d3361) }

var fileDescriptor_06b463ba992d3361 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0xe3, 0xb8, 0x4d, 0x26, 0xb4, 0x84, 0x55, 0x91, 0x56, 0x15, 0x42, 0x56, 0x84, 0x90,
	0x4f, 0x05, 0x15, 0x4e, 0xdc, 0xa8, 0x10, 0x88, 0x0b, 0x2a, 0x5b, 0x38, 0x22, 0xb4, 0xb1, 0xc7,
	0xe9, 0xaa, 0xf6, 0xae, 0xbb, 0x1f, 0xa0, 0xfc, 0x52, 0xfe, 0x0e, 0xda, 0x71, 0x62, 0xdc, 0x43,
	0xb8, 0x44, 0xf3, 0xde, 0x9b, 0xc9, 0x93, 0xf7, 0x3d, 0x78, 0x22, 0x9d, 0x43, 0xff, 0x8a, 0x7e,
	0x2f, 0x3a, 0x6b, 0xbc, 0x61, 0x19, 0x81, 0x73, 0x3e, 0x52, 0x7e, 0xb6, 0x52, 0xcb, 0x0d, 0xf6,
	0x0b, 0xab, 0x3f, 0x29, 0xcc, 0xdf, 0x47, 0xfa, 0xb3, 0xae, 0x0d, 0x3b, 0x85, 0x89, 0xaa, 0x78,
	0x92, 0x27, 0xc5, 0x5c, 0x4c, 0x54, 0xc5, 0x96, 0x90, 0x76, 0xaa, 0xe2, 0x13, 0x22, 0xe2, 0xc8,
	0x18, 0x4c, 0xb5, 0x6c, 0x91, 0xa7, 0x44, 0xd1, 0xcc, 0x72, 0x58, 0x54, 0xe8, 0x4a, 0xab, 0x3a,
	0xaf, 0x8c, 0xe6, 0x53, 0x92, 0xc6, 0x14, 0x3b, 0x83, 0x4c, 0xfb, 0x6d, 0x87, 0x3c, 0x23, 0xad,
	0x07, 0xec, 0x1c, 0x66, 0x65, 0xb0, 0x16, 0x75, 0xb9, 0xe5, 0x47, 0x24, 0x0c, 0x98, 0x3d, 0x07,
	0xa8, 0x95, 0x96, 0xba, 0xc4, 0x9b, 0xd0, 0xf2, 0xe3, 0x3c, 0x29, 0x52, 0x31, 0x62, 0xe2, 0xad,
	0xeb, 0x1a, 0xe5, 0xbf, 0x84, 0x96, 0xcf, 0xf2, 0xa4, 0xc8, 0xc4, 0x80, 0xd9, 0x33, 0x98, 0xaf,
	0x71, 0xa3, 0xf4, 0x37, 0xd5, 0x22, 0x9f, 0xd3, 0x1f, 0xff, 0x23, 0x18, 0x87, 0x63, 0xd4, 0x15,
	0x69, 0x40, 0xda, 0x1e, 0xb2, 0x02, 0xa6, 0x6b, 0xe9, 0x90, 0x2f, 0xf2, 0xa4, 0x58, 0x5c, 0x9e,
	0x5d, 0xf4, 0x0f, 0xf9, 0x41, 0x7a, 0x79, 0xd5, 0x98, 0xf2, 0xee, 0x4a, 0x3a, 0x14, 0xb4, 0xc1,
	0xde, 0xc1, 0xc9, 0x7d, 0x90, 0x8d, 0xaa, 0x15, 0x56, 0x1f, 0x55, 0x83, 0xfc, 0xd1, 0xc3, 0x13,
	0x53, 0x86, 0x16, 0xb5, 0xbf, 0xf1, 0xc6, 0xa2, 0x78, 0xb8, 0xca, 0xde, 0x02, 0xfc, 0xbe, 0x55,
	0x1e, 0xaf, 0x65, 0x87, 0x96, 0x9f, 0xfc, 0xe7, 0x70, 0xb4, 0xc7, 0x5e, 0xc3, 0xac, 0x53, 0xa5,
	0x0f, 0x16, 0x1d, 0x3f, 0xcd, 0xd3, 0x83, 0x37, 0xc3, 0xd6, 0xea, 0x07, 0x3c, 0x16, 0x78, 0x1f,
	0xd0, 0x79, 0xca, 0xf7, 0x3a, 0xf8, 0x18, 0xe7, 0x1d, 0x6e, 0x77, 0xf9, 0xc6, 0x31, 0x32, 0x75,
	0xd0, 0xfb, 0x80, 0xeb, 0xa0, 0xd9, 0x4b, 0xc8, 0x7e, 0xc9, 0x26, 0xf4, 0x09, 0x2f, 0x2e, 0x97,
	0x3b, 0x97, 0xa1, 0x23, 0xa2, 0x97, 0x57, 0x35, 0x2c, 0x05, 0xba, 0xce, 0x68, 0x87, 0xa4, 0x7d,
	0x42, 0x1f, 0xcb, 0x51, 0x9a, 0x0a, 0xc9, 0x20, 0x13, 0x34, 0xc7, 0xe7, 0x6e, 0xd1, 0x39, 0xb9,
	0xc1, 0x9d, 0xcb, 0x1e, 0xb2, 0x17, 0x30, 0xad, 0xa4, 0x97, 0x07, 0x8d, 0x48, 0x5d, 0x7d, 0x87,
	0xa7, 0xe3, 0xcf, 0x10, 0xaa, 0xbc, 0xfd, 0x1a, 0xd0, 0x6e, 0x63, 0xa7, 0x1a, 0xd5, 0x2a, 0xbf,
	0x73, 0xeb, 0x41, 0x64, 0x2d, 0x35, 0x6d, 0xd2, 0xb3, 0x04, 0xf6, 0x3d, 0x4e, 0x87, 0x1e, 0xaf,
	0x8f, 0xa8, 0xfe, 0x6f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x95, 0xd6, 0x27, 0x01, 0x34, 0x03,
	0x00, 0x00,
}
