// Code generated by protoc-gen-go. DO NOT EDIT.
// source: basic/ato_project.proto

package basic

import (
	fmt "fmt"
	math "math"
	"protobuf/projects/go/protocol/asset"

	proto "github.com/golang/protobuf/proto"
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

// ATO 项目对象
type ProjectATO struct {
	Owner                string     `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Id                   string     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 int32      `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Base                 *BaseATO   `protobuf:"bytes,4,opt,name=base,proto3" json:"base,omitempty"`
	Finance              *Financing `protobuf:"bytes,5,opt,name=finance,proto3" json:"finance,omitempty"`
	ExitPlan             *ExitPlan  `protobuf:"bytes,6,opt,name=exitPlan,proto3" json:"exitPlan,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProjectATO) Reset()         { *m = ProjectATO{} }
func (m *ProjectATO) String() string { return proto.CompactTextString(m) }
func (*ProjectATO) ProtoMessage()    {}
func (*ProjectATO) Descriptor() ([]byte, []int) {
	return fileDescriptor_83981327d982bd68, []int{0}
}

func (m *ProjectATO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectATO.Unmarshal(m, b)
}
func (m *ProjectATO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectATO.Marshal(b, m, deterministic)
}
func (m *ProjectATO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectATO.Merge(m, src)
}
func (m *ProjectATO) XXX_Size() int {
	return xxx_messageInfo_ProjectATO.Size(m)
}
func (m *ProjectATO) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectATO.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectATO proto.InternalMessageInfo

func (m *ProjectATO) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ProjectATO) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProjectATO) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ProjectATO) GetBase() *BaseATO {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *ProjectATO) GetFinance() *Financing {
	if m != nil {
		return m.Finance
	}
	return nil
}

func (m *ProjectATO) GetExitPlan() *ExitPlan {
	if m != nil {
		return m.ExitPlan
	}
	return nil
}

// -----------------------------------------------------基本信息---------------------------------------------------------//
type BaseATO struct {
	OwnerId              string                 `protobuf:"bytes,1,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	OwnerType            int32                  `protobuf:"varint,2,opt,name=ownerType,proto3" json:"ownerType,omitempty"`
	AssetId              string                 `protobuf:"bytes,3,opt,name=assetId,proto3" json:"assetId,omitempty"`
	Name                 string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Picture              []*asset.DocumentStore `protobuf:"bytes,5,rep,name=picture,proto3" json:"picture,omitempty"`
	ShortDesc            string                 `protobuf:"bytes,6,opt,name=shortDesc,proto3" json:"shortDesc,omitempty"`
	EntireDesc           string                 `protobuf:"bytes,7,opt,name=entireDesc,proto3" json:"entireDesc,omitempty"`
	WhitePaper           *asset.DocumentStore   `protobuf:"bytes,8,opt,name=whitePaper,proto3" json:"whitePaper,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *BaseATO) Reset()         { *m = BaseATO{} }
func (m *BaseATO) String() string { return proto.CompactTextString(m) }
func (*BaseATO) ProtoMessage()    {}
func (*BaseATO) Descriptor() ([]byte, []int) {
	return fileDescriptor_83981327d982bd68, []int{1}
}

func (m *BaseATO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseATO.Unmarshal(m, b)
}
func (m *BaseATO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseATO.Marshal(b, m, deterministic)
}
func (m *BaseATO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseATO.Merge(m, src)
}
func (m *BaseATO) XXX_Size() int {
	return xxx_messageInfo_BaseATO.Size(m)
}
func (m *BaseATO) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseATO.DiscardUnknown(m)
}

var xxx_messageInfo_BaseATO proto.InternalMessageInfo

func (m *BaseATO) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *BaseATO) GetOwnerType() int32 {
	if m != nil {
		return m.OwnerType
	}
	return 0
}

func (m *BaseATO) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *BaseATO) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BaseATO) GetPicture() []*asset.DocumentStore {
	if m != nil {
		return m.Picture
	}
	return nil
}

func (m *BaseATO) GetShortDesc() string {
	if m != nil {
		return m.ShortDesc
	}
	return ""
}

func (m *BaseATO) GetEntireDesc() string {
	if m != nil {
		return m.EntireDesc
	}
	return ""
}

func (m *BaseATO) GetWhitePaper() *asset.DocumentStore {
	if m != nil {
		return m.WhitePaper
	}
	return nil
}

// -----------------------------------------------------融资方案---------------------------------------------------------//
type Financing struct {
	TokenType            int32    `protobuf:"varint,1,opt,name=tokenType,proto3" json:"tokenType,omitempty"`
	TokenName            string   `protobuf:"bytes,2,opt,name=tokenName,proto3" json:"tokenName,omitempty"`
	ContractAddr         string   `protobuf:"bytes,3,opt,name=contractAddr,proto3" json:"contractAddr,omitempty"`
	TokenPrice           int32    `protobuf:"varint,4,opt,name=tokenPrice,proto3" json:"tokenPrice,omitempty"`
	TokenUnit            string   `protobuf:"bytes,5,opt,name=tokenUnit,proto3" json:"tokenUnit,omitempty"`
	HardCap              int64    `protobuf:"varint,6,opt,name=hardCap,proto3" json:"hardCap,omitempty"`
	FundGoal             int64    `protobuf:"varint,7,opt,name=fundGoal,proto3" json:"fundGoal,omitempty"`
	FundGoalUnit         string   `protobuf:"bytes,8,opt,name=fundGoalUnit,proto3" json:"fundGoalUnit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Financing) Reset()         { *m = Financing{} }
func (m *Financing) String() string { return proto.CompactTextString(m) }
func (*Financing) ProtoMessage()    {}
func (*Financing) Descriptor() ([]byte, []int) {
	return fileDescriptor_83981327d982bd68, []int{2}
}

func (m *Financing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Financing.Unmarshal(m, b)
}
func (m *Financing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Financing.Marshal(b, m, deterministic)
}
func (m *Financing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Financing.Merge(m, src)
}
func (m *Financing) XXX_Size() int {
	return xxx_messageInfo_Financing.Size(m)
}
func (m *Financing) XXX_DiscardUnknown() {
	xxx_messageInfo_Financing.DiscardUnknown(m)
}

var xxx_messageInfo_Financing proto.InternalMessageInfo

func (m *Financing) GetTokenType() int32 {
	if m != nil {
		return m.TokenType
	}
	return 0
}

func (m *Financing) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *Financing) GetContractAddr() string {
	if m != nil {
		return m.ContractAddr
	}
	return ""
}

func (m *Financing) GetTokenPrice() int32 {
	if m != nil {
		return m.TokenPrice
	}
	return 0
}

func (m *Financing) GetTokenUnit() string {
	if m != nil {
		return m.TokenUnit
	}
	return ""
}

func (m *Financing) GetHardCap() int64 {
	if m != nil {
		return m.HardCap
	}
	return 0
}

func (m *Financing) GetFundGoal() int64 {
	if m != nil {
		return m.FundGoal
	}
	return 0
}

func (m *Financing) GetFundGoalUnit() string {
	if m != nil {
		return m.FundGoalUnit
	}
	return ""
}

// -----------------------------------------------------退出方案(暂不展开)--------------------------------------------------//
type ExitPlan struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExitPlan) Reset()         { *m = ExitPlan{} }
func (m *ExitPlan) String() string { return proto.CompactTextString(m) }
func (*ExitPlan) ProtoMessage()    {}
func (*ExitPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_83981327d982bd68, []int{3}
}

func (m *ExitPlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitPlan.Unmarshal(m, b)
}
func (m *ExitPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitPlan.Marshal(b, m, deterministic)
}
func (m *ExitPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitPlan.Merge(m, src)
}
func (m *ExitPlan) XXX_Size() int {
	return xxx_messageInfo_ExitPlan.Size(m)
}
func (m *ExitPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitPlan.DiscardUnknown(m)
}

var xxx_messageInfo_ExitPlan proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProjectATO)(nil), "basic.ProjectATO")
	proto.RegisterType((*BaseATO)(nil), "basic.BaseATO")
	proto.RegisterType((*Financing)(nil), "basic.Financing")
	proto.RegisterType((*ExitPlan)(nil), "basic.ExitPlan")
}

func init() { proto.RegisterFile("basic/ato_project.proto", fileDescriptor_83981327d982bd68) }

var fileDescriptor_83981327d982bd68 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x55, 0xd2, 0x66, 0x49, 0xee, 0xd0, 0x40, 0xd6, 0x24, 0xac, 0x09, 0xa1, 0x2a, 0x4f, 0x15,
	0x48, 0x99, 0x34, 0xf8, 0x81, 0xc2, 0x00, 0xed, 0x85, 0x55, 0xa6, 0x3c, 0x4f, 0xae, 0x73, 0xb7,
	0x1a, 0x56, 0x3b, 0x72, 0x5c, 0x15, 0x9e, 0xf9, 0x04, 0x3e, 0x89, 0x1f, 0x43, 0xbe, 0x89, 0xd3,
	0x22, 0xf1, 0x12, 0xdd, 0x7b, 0xee, 0x71, 0xee, 0x39, 0xc7, 0x86, 0xe7, 0x6b, 0xd9, 0x69, 0x75,
	0x29, 0xbd, 0xbd, 0x6b, 0x9d, 0xfd, 0x86, 0xca, 0xd7, 0xad, 0xb3, 0xde, 0xb2, 0x8c, 0x06, 0x17,
	0x5c, 0x76, 0x1d, 0xfa, 0x4b, 0xfa, 0xde, 0x6d, 0xa5, 0x91, 0x0f, 0xd8, 0x13, 0xaa, 0x3f, 0x09,
	0xc0, 0xb2, 0x3f, 0xb2, 0x58, 0xdd, 0xb2, 0x73, 0xc8, 0xec, 0xde, 0xa0, 0xe3, 0xc9, 0x2c, 0x99,
	0x97, 0xa2, 0x6f, 0xd8, 0x19, 0xa4, 0xba, 0xe1, 0x29, 0x41, 0xa9, 0x6e, 0x18, 0x83, 0xa9, 0xff,
	0xd9, 0x22, 0x9f, 0xcc, 0x92, 0x79, 0x26, 0xa8, 0x66, 0x15, 0x4c, 0xd7, 0xb2, 0x43, 0x3e, 0x9d,
	0x25, 0xf3, 0xd3, 0xab, 0xb3, 0x9a, 0x16, 0xd7, 0xef, 0x64, 0x87, 0x8b, 0xd5, 0xad, 0xa0, 0x19,
	0x7b, 0x05, 0xf9, 0xbd, 0x36, 0xd2, 0x28, 0xe4, 0x19, 0xd1, 0x9e, 0x0d, 0xb4, 0x8f, 0x84, 0x6a,
	0xf3, 0x20, 0x22, 0x81, 0xbd, 0x86, 0x02, 0x7f, 0x68, 0xbf, 0x7c, 0x94, 0x86, 0x9f, 0x10, 0xf9,
	0xe9, 0x40, 0xfe, 0x30, 0xc0, 0x62, 0x24, 0x54, 0xbf, 0x53, 0xc8, 0x87, 0x55, 0x8c, 0x43, 0x4e,
	0xaa, 0x6f, 0x9a, 0xc1, 0x44, 0x6c, 0xd9, 0x0b, 0x28, 0xa9, 0x5c, 0x05, 0xed, 0x29, 0x69, 0x3f,
	0x00, 0xe1, 0x1c, 0xe5, 0x73, 0xd3, 0x90, 0xaf, 0x52, 0xc4, 0x36, 0xd8, 0x35, 0x72, 0xdb, 0x5b,
	0x2b, 0x05, 0xd5, 0xac, 0x86, 0xbc, 0xd5, 0xca, 0xef, 0x5c, 0xb0, 0x32, 0x99, 0x9f, 0x5e, 0x9d,
	0xd7, 0x44, 0xaf, 0xaf, 0xad, 0xda, 0x6d, 0xd1, 0xf8, 0x2f, 0xde, 0x3a, 0x14, 0x91, 0x14, 0x76,
	0x77, 0x1b, 0xeb, 0xfc, 0x35, 0x76, 0x8a, 0xfc, 0x94, 0xe2, 0x00, 0xb0, 0x97, 0x00, 0x68, 0xbc,
	0x76, 0x48, 0xe3, 0x9c, 0xc6, 0x47, 0x08, 0x7b, 0x0b, 0xb0, 0xdf, 0x68, 0x8f, 0x4b, 0xd9, 0xa2,
	0xe3, 0x05, 0xc5, 0xf1, 0xff, 0x85, 0x47, 0xbc, 0xea, 0x57, 0x0a, 0xe5, 0x98, 0x6c, 0x50, 0xe0,
	0xed, 0x77, 0x34, 0xe4, 0x3e, 0xe9, 0xdd, 0x8f, 0xc0, 0x38, 0xfd, 0x1c, 0x8c, 0xf6, 0x37, 0x7d,
	0x00, 0x58, 0x05, 0x4f, 0x94, 0x35, 0xde, 0x49, 0xe5, 0x17, 0x4d, 0xe3, 0x86, 0x80, 0xfe, 0xc1,
	0x82, 0x07, 0x3a, 0xb0, 0x74, 0x5a, 0xf5, 0x59, 0x65, 0xe2, 0x08, 0x19, 0x37, 0x7c, 0x35, 0xda,
	0xd3, 0xf5, 0xc7, 0x0d, 0x01, 0x08, 0xe9, 0x6f, 0xa4, 0x6b, 0xde, 0xcb, 0x96, 0xd2, 0x99, 0x88,
	0xd8, 0xb2, 0x0b, 0x28, 0xee, 0x77, 0xa6, 0xf9, 0x64, 0xe5, 0x23, 0x25, 0x33, 0x11, 0x63, 0x1f,
	0x74, 0xc5, 0x9a, 0x7e, 0x5b, 0xf4, 0xba, 0x8e, 0xb1, 0x0a, 0xa0, 0x88, 0x2f, 0x66, 0x7d, 0x42,
	0x8f, 0xfe, 0xcd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x3b, 0x70, 0x51, 0x30, 0x03, 0x00,
	0x00,
}