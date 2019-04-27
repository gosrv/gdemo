// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logic.proto

package netproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 错误码
type E_Code int32

const (
	E_Code_E_ERROR                 E_Code = 0
	E_Code_E_OK                    E_Code = 1
	E_Code_E_RELOGIN               E_Code = 2
	E_Code_E_NONE_EXIST            E_Code = 3
	E_Code_E_UNKNOWN               E_Code = 4
	E_Code_E_SERVER_INTERNAL_ERROR E_Code = 5
	E_Code_E_INVALID_PARAM         E_Code = 6
	E_Code_E_INVALID_OPT           E_Code = 7
	E_Code_E_LIMIT_GOLD            E_Code = 20
	E_Code_E_LIMIT_DIAMOND         E_Code = 21
	E_Code_E_LIMIT_TIMES           E_Code = 22
	E_Code_E_LIMIT_CHAPTER_LEVEL   E_Code = 23
	E_Code_E_LIMIT_PLAYER_LEVEL    E_Code = 24
	E_Code_E_LIMIT_CAPACITY        E_Code = 25
)

var E_Code_name = map[int32]string{
	0:  "E_ERROR",
	1:  "E_OK",
	2:  "E_RELOGIN",
	3:  "E_NONE_EXIST",
	4:  "E_UNKNOWN",
	5:  "E_SERVER_INTERNAL_ERROR",
	6:  "E_INVALID_PARAM",
	7:  "E_INVALID_OPT",
	20: "E_LIMIT_GOLD",
	21: "E_LIMIT_DIAMOND",
	22: "E_LIMIT_TIMES",
	23: "E_LIMIT_CHAPTER_LEVEL",
	24: "E_LIMIT_PLAYER_LEVEL",
	25: "E_LIMIT_CAPACITY",
}
var E_Code_value = map[string]int32{
	"E_ERROR":                 0,
	"E_OK":                    1,
	"E_RELOGIN":               2,
	"E_NONE_EXIST":            3,
	"E_UNKNOWN":               4,
	"E_SERVER_INTERNAL_ERROR": 5,
	"E_INVALID_PARAM":         6,
	"E_INVALID_OPT":           7,
	"E_LIMIT_GOLD":            20,
	"E_LIMIT_DIAMOND":         21,
	"E_LIMIT_TIMES":           22,
	"E_LIMIT_CHAPTER_LEVEL":   23,
	"E_LIMIT_PLAYER_LEVEL":    24,
	"E_LIMIT_CAPACITY":        25,
}

func (x E_Code) String() string {
	return proto.EnumName(E_Code_name, int32(x))
}
func (E_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{0}
}

// 前100号的消息留给系统用
type EMsgIds int32

const (
	EMsgIds_ECS_None EMsgIds = 0
	EMsgIds_ECS_P2P  EMsgIds = 101
	EMsgIds_ESC_P2P  EMsgIds = 102
	// 心跳
	EMsgIds_ECS_Tick             EMsgIds = 201
	EMsgIds_ESC_Tick             EMsgIds = 202
	EMsgIds_ECS_Login            EMsgIds = 203
	EMsgIds_ESC_Login            EMsgIds = 204
	EMsgIds_ECS_ShopAddDiamond   EMsgIds = 205
	EMsgIds_ESC_ShopAddDiamond   EMsgIds = 206
	EMsgIds_ECS_ShopBuyGold      EMsgIds = 207
	EMsgIds_ESC_ShopBuyGold      EMsgIds = 208
	EMsgIds_ECS_HeroDraw         EMsgIds = 209
	EMsgIds_ESC_HeroDraw         EMsgIds = 210
	EMsgIds_ECS_HeroUplevel      EMsgIds = 211
	EMsgIds_ESC_HeroUpLevel      EMsgIds = 212
	EMsgIds_ECS_HeroEquip        EMsgIds = 213
	EMsgIds_ESC_HeroEquip        EMsgIds = 214
	EMsgIds_ECS_ChapterGetPrize  EMsgIds = 215
	EMsgIds_ESC_ChapterGetPrize  EMsgIds = 216
	EMsgIds_ECS_ChapterChallenge EMsgIds = 217
	EMsgIds_ESC_ChapterChallenge EMsgIds = 218
)

var EMsgIds_name = map[int32]string{
	0:   "ECS_None",
	101: "ECS_P2P",
	102: "ESC_P2P",
	201: "ECS_Tick",
	202: "ESC_Tick",
	203: "ECS_Login",
	204: "ESC_Login",
	205: "ECS_ShopAddDiamond",
	206: "ESC_ShopAddDiamond",
	207: "ECS_ShopBuyGold",
	208: "ESC_ShopBuyGold",
	209: "ECS_HeroDraw",
	210: "ESC_HeroDraw",
	211: "ECS_HeroUplevel",
	212: "ESC_HeroUpLevel",
	213: "ECS_HeroEquip",
	214: "ESC_HeroEquip",
	215: "ECS_ChapterGetPrize",
	216: "ESC_ChapterGetPrize",
	217: "ECS_ChapterChallenge",
	218: "ESC_ChapterChallenge",
}
var EMsgIds_value = map[string]int32{
	"ECS_None":             0,
	"ECS_P2P":              101,
	"ESC_P2P":              102,
	"ECS_Tick":             201,
	"ESC_Tick":             202,
	"ECS_Login":            203,
	"ESC_Login":            204,
	"ECS_ShopAddDiamond":   205,
	"ESC_ShopAddDiamond":   206,
	"ECS_ShopBuyGold":      207,
	"ESC_ShopBuyGold":      208,
	"ECS_HeroDraw":         209,
	"ESC_HeroDraw":         210,
	"ECS_HeroUplevel":      211,
	"ESC_HeroUpLevel":      212,
	"ECS_HeroEquip":        213,
	"ESC_HeroEquip":        214,
	"ECS_ChapterGetPrize":  215,
	"ESC_ChapterGetPrize":  216,
	"ECS_ChapterChallenge": 217,
	"ESC_ChapterChallenge": 218,
}

func (x EMsgIds) String() string {
	return proto.EnumName(EMsgIds_name, int32(x))
}
func (EMsgIds) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{1}
}

// 心跳
type CS_Tick struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_Tick) Reset()         { *m = CS_Tick{} }
func (m *CS_Tick) String() string { return proto.CompactTextString(m) }
func (*CS_Tick) ProtoMessage()    {}
func (*CS_Tick) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{0}
}
func (m *CS_Tick) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_Tick.Unmarshal(m, b)
}
func (m *CS_Tick) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_Tick.Marshal(b, m, deterministic)
}
func (dst *CS_Tick) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_Tick.Merge(dst, src)
}
func (m *CS_Tick) XXX_Size() int {
	return xxx_messageInfo_CS_Tick.Size(m)
}
func (m *CS_Tick) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_Tick.DiscardUnknown(m)
}

var xxx_messageInfo_CS_Tick proto.InternalMessageInfo

type SC_Tick struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_Tick) Reset()         { *m = SC_Tick{} }
func (m *SC_Tick) String() string { return proto.CompactTextString(m) }
func (*SC_Tick) ProtoMessage()    {}
func (*SC_Tick) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{1}
}
func (m *SC_Tick) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_Tick.Unmarshal(m, b)
}
func (m *SC_Tick) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_Tick.Marshal(b, m, deterministic)
}
func (dst *SC_Tick) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_Tick.Merge(dst, src)
}
func (m *SC_Tick) XXX_Size() int {
	return xxx_messageInfo_SC_Tick.Size(m)
}
func (m *SC_Tick) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_Tick.DiscardUnknown(m)
}

var xxx_messageInfo_SC_Tick proto.InternalMessageInfo

// 登录
type CS_Login struct {
	Token                string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_Login) Reset()         { *m = CS_Login{} }
func (m *CS_Login) String() string { return proto.CompactTextString(m) }
func (*CS_Login) ProtoMessage()    {}
func (*CS_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{2}
}
func (m *CS_Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_Login.Unmarshal(m, b)
}
func (m *CS_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_Login.Marshal(b, m, deterministic)
}
func (dst *CS_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_Login.Merge(dst, src)
}
func (m *CS_Login) XXX_Size() int {
	return xxx_messageInfo_CS_Login.Size(m)
}
func (m *CS_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_Login.DiscardUnknown(m)
}

var xxx_messageInfo_CS_Login proto.InternalMessageInfo

func (m *CS_Login) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SC_Login struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_Login) Reset()         { *m = SC_Login{} }
func (m *SC_Login) String() string { return proto.CompactTextString(m) }
func (*SC_Login) ProtoMessage()    {}
func (*SC_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{3}
}
func (m *SC_Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_Login.Unmarshal(m, b)
}
func (m *SC_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_Login.Marshal(b, m, deterministic)
}
func (dst *SC_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_Login.Merge(dst, src)
}
func (m *SC_Login) XXX_Size() int {
	return xxx_messageInfo_SC_Login.Size(m)
}
func (m *SC_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_Login.DiscardUnknown(m)
}

var xxx_messageInfo_SC_Login proto.InternalMessageInfo

func (m *SC_Login) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

// 冲钻石
type CS_ShopAddDiamond struct {
	Idx                  int32    `protobuf:"varint,1,opt,name=idx" json:"idx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_ShopAddDiamond) Reset()         { *m = CS_ShopAddDiamond{} }
func (m *CS_ShopAddDiamond) String() string { return proto.CompactTextString(m) }
func (*CS_ShopAddDiamond) ProtoMessage()    {}
func (*CS_ShopAddDiamond) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{4}
}
func (m *CS_ShopAddDiamond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_ShopAddDiamond.Unmarshal(m, b)
}
func (m *CS_ShopAddDiamond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_ShopAddDiamond.Marshal(b, m, deterministic)
}
func (dst *CS_ShopAddDiamond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_ShopAddDiamond.Merge(dst, src)
}
func (m *CS_ShopAddDiamond) XXX_Size() int {
	return xxx_messageInfo_CS_ShopAddDiamond.Size(m)
}
func (m *CS_ShopAddDiamond) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_ShopAddDiamond.DiscardUnknown(m)
}

var xxx_messageInfo_CS_ShopAddDiamond proto.InternalMessageInfo

func (m *CS_ShopAddDiamond) GetIdx() int32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

type SC_ShopAddDiamond struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_ShopAddDiamond) Reset()         { *m = SC_ShopAddDiamond{} }
func (m *SC_ShopAddDiamond) String() string { return proto.CompactTextString(m) }
func (*SC_ShopAddDiamond) ProtoMessage()    {}
func (*SC_ShopAddDiamond) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{5}
}
func (m *SC_ShopAddDiamond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_ShopAddDiamond.Unmarshal(m, b)
}
func (m *SC_ShopAddDiamond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_ShopAddDiamond.Marshal(b, m, deterministic)
}
func (dst *SC_ShopAddDiamond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_ShopAddDiamond.Merge(dst, src)
}
func (m *SC_ShopAddDiamond) XXX_Size() int {
	return xxx_messageInfo_SC_ShopAddDiamond.Size(m)
}
func (m *SC_ShopAddDiamond) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_ShopAddDiamond.DiscardUnknown(m)
}

var xxx_messageInfo_SC_ShopAddDiamond proto.InternalMessageInfo

func (m *SC_ShopAddDiamond) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

// 购买金币
type CS_ShopBuyGold struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_ShopBuyGold) Reset()         { *m = CS_ShopBuyGold{} }
func (m *CS_ShopBuyGold) String() string { return proto.CompactTextString(m) }
func (*CS_ShopBuyGold) ProtoMessage()    {}
func (*CS_ShopBuyGold) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{6}
}
func (m *CS_ShopBuyGold) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_ShopBuyGold.Unmarshal(m, b)
}
func (m *CS_ShopBuyGold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_ShopBuyGold.Marshal(b, m, deterministic)
}
func (dst *CS_ShopBuyGold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_ShopBuyGold.Merge(dst, src)
}
func (m *CS_ShopBuyGold) XXX_Size() int {
	return xxx_messageInfo_CS_ShopBuyGold.Size(m)
}
func (m *CS_ShopBuyGold) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_ShopBuyGold.DiscardUnknown(m)
}

var xxx_messageInfo_CS_ShopBuyGold proto.InternalMessageInfo

type SC_ShopBuyGold struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_ShopBuyGold) Reset()         { *m = SC_ShopBuyGold{} }
func (m *SC_ShopBuyGold) String() string { return proto.CompactTextString(m) }
func (*SC_ShopBuyGold) ProtoMessage()    {}
func (*SC_ShopBuyGold) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{7}
}
func (m *SC_ShopBuyGold) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_ShopBuyGold.Unmarshal(m, b)
}
func (m *SC_ShopBuyGold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_ShopBuyGold.Marshal(b, m, deterministic)
}
func (dst *SC_ShopBuyGold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_ShopBuyGold.Merge(dst, src)
}
func (m *SC_ShopBuyGold) XXX_Size() int {
	return xxx_messageInfo_SC_ShopBuyGold.Size(m)
}
func (m *SC_ShopBuyGold) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_ShopBuyGold.DiscardUnknown(m)
}

var xxx_messageInfo_SC_ShopBuyGold proto.InternalMessageInfo

func (m *SC_ShopBuyGold) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

// 抽英雄
type CS_HeroDraw struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_HeroDraw) Reset()         { *m = CS_HeroDraw{} }
func (m *CS_HeroDraw) String() string { return proto.CompactTextString(m) }
func (*CS_HeroDraw) ProtoMessage()    {}
func (*CS_HeroDraw) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{8}
}
func (m *CS_HeroDraw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_HeroDraw.Unmarshal(m, b)
}
func (m *CS_HeroDraw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_HeroDraw.Marshal(b, m, deterministic)
}
func (dst *CS_HeroDraw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_HeroDraw.Merge(dst, src)
}
func (m *CS_HeroDraw) XXX_Size() int {
	return xxx_messageInfo_CS_HeroDraw.Size(m)
}
func (m *CS_HeroDraw) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_HeroDraw.DiscardUnknown(m)
}

var xxx_messageInfo_CS_HeroDraw proto.InternalMessageInfo

func (m *CS_HeroDraw) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type SC_HeroDraw struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_HeroDraw) Reset()         { *m = SC_HeroDraw{} }
func (m *SC_HeroDraw) String() string { return proto.CompactTextString(m) }
func (*SC_HeroDraw) ProtoMessage()    {}
func (*SC_HeroDraw) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{9}
}
func (m *SC_HeroDraw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_HeroDraw.Unmarshal(m, b)
}
func (m *SC_HeroDraw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_HeroDraw.Marshal(b, m, deterministic)
}
func (dst *SC_HeroDraw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_HeroDraw.Merge(dst, src)
}
func (m *SC_HeroDraw) XXX_Size() int {
	return xxx_messageInfo_SC_HeroDraw.Size(m)
}
func (m *SC_HeroDraw) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_HeroDraw.DiscardUnknown(m)
}

var xxx_messageInfo_SC_HeroDraw proto.InternalMessageInfo

func (m *SC_HeroDraw) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

// 升级英雄等级
type CS_HeroUplevel struct {
	HeroId               int32    `protobuf:"varint,1,opt,name=heroId" json:"heroId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_HeroUplevel) Reset()         { *m = CS_HeroUplevel{} }
func (m *CS_HeroUplevel) String() string { return proto.CompactTextString(m) }
func (*CS_HeroUplevel) ProtoMessage()    {}
func (*CS_HeroUplevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{10}
}
func (m *CS_HeroUplevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_HeroUplevel.Unmarshal(m, b)
}
func (m *CS_HeroUplevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_HeroUplevel.Marshal(b, m, deterministic)
}
func (dst *CS_HeroUplevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_HeroUplevel.Merge(dst, src)
}
func (m *CS_HeroUplevel) XXX_Size() int {
	return xxx_messageInfo_CS_HeroUplevel.Size(m)
}
func (m *CS_HeroUplevel) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_HeroUplevel.DiscardUnknown(m)
}

var xxx_messageInfo_CS_HeroUplevel proto.InternalMessageInfo

func (m *CS_HeroUplevel) GetHeroId() int32 {
	if m != nil {
		return m.HeroId
	}
	return 0
}

type SC_HeroUpLevel struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_HeroUpLevel) Reset()         { *m = SC_HeroUpLevel{} }
func (m *SC_HeroUpLevel) String() string { return proto.CompactTextString(m) }
func (*SC_HeroUpLevel) ProtoMessage()    {}
func (*SC_HeroUpLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{11}
}
func (m *SC_HeroUpLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_HeroUpLevel.Unmarshal(m, b)
}
func (m *SC_HeroUpLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_HeroUpLevel.Marshal(b, m, deterministic)
}
func (dst *SC_HeroUpLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_HeroUpLevel.Merge(dst, src)
}
func (m *SC_HeroUpLevel) XXX_Size() int {
	return xxx_messageInfo_SC_HeroUpLevel.Size(m)
}
func (m *SC_HeroUpLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_HeroUpLevel.DiscardUnknown(m)
}

var xxx_messageInfo_SC_HeroUpLevel proto.InternalMessageInfo

func (m *SC_HeroUpLevel) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

// 装备英雄
type CS_HeroEquip struct {
	HeroId               int32    `protobuf:"varint,1,opt,name=heroId" json:"heroId,omitempty"`
	EquipId              int32    `protobuf:"varint,2,opt,name=equipId" json:"equipId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_HeroEquip) Reset()         { *m = CS_HeroEquip{} }
func (m *CS_HeroEquip) String() string { return proto.CompactTextString(m) }
func (*CS_HeroEquip) ProtoMessage()    {}
func (*CS_HeroEquip) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{12}
}
func (m *CS_HeroEquip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_HeroEquip.Unmarshal(m, b)
}
func (m *CS_HeroEquip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_HeroEquip.Marshal(b, m, deterministic)
}
func (dst *CS_HeroEquip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_HeroEquip.Merge(dst, src)
}
func (m *CS_HeroEquip) XXX_Size() int {
	return xxx_messageInfo_CS_HeroEquip.Size(m)
}
func (m *CS_HeroEquip) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_HeroEquip.DiscardUnknown(m)
}

var xxx_messageInfo_CS_HeroEquip proto.InternalMessageInfo

func (m *CS_HeroEquip) GetHeroId() int32 {
	if m != nil {
		return m.HeroId
	}
	return 0
}

func (m *CS_HeroEquip) GetEquipId() int32 {
	if m != nil {
		return m.EquipId
	}
	return 0
}

type SC_HeroEquip struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_HeroEquip) Reset()         { *m = SC_HeroEquip{} }
func (m *SC_HeroEquip) String() string { return proto.CompactTextString(m) }
func (*SC_HeroEquip) ProtoMessage()    {}
func (*SC_HeroEquip) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{13}
}
func (m *SC_HeroEquip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_HeroEquip.Unmarshal(m, b)
}
func (m *SC_HeroEquip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_HeroEquip.Marshal(b, m, deterministic)
}
func (dst *SC_HeroEquip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_HeroEquip.Merge(dst, src)
}
func (m *SC_HeroEquip) XXX_Size() int {
	return xxx_messageInfo_SC_HeroEquip.Size(m)
}
func (m *SC_HeroEquip) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_HeroEquip.DiscardUnknown(m)
}

var xxx_messageInfo_SC_HeroEquip proto.InternalMessageInfo

func (m *SC_HeroEquip) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

// 领取挂机奖励
type CS_ChapterGetPrize struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_ChapterGetPrize) Reset()         { *m = CS_ChapterGetPrize{} }
func (m *CS_ChapterGetPrize) String() string { return proto.CompactTextString(m) }
func (*CS_ChapterGetPrize) ProtoMessage()    {}
func (*CS_ChapterGetPrize) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{14}
}
func (m *CS_ChapterGetPrize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_ChapterGetPrize.Unmarshal(m, b)
}
func (m *CS_ChapterGetPrize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_ChapterGetPrize.Marshal(b, m, deterministic)
}
func (dst *CS_ChapterGetPrize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_ChapterGetPrize.Merge(dst, src)
}
func (m *CS_ChapterGetPrize) XXX_Size() int {
	return xxx_messageInfo_CS_ChapterGetPrize.Size(m)
}
func (m *CS_ChapterGetPrize) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_ChapterGetPrize.DiscardUnknown(m)
}

var xxx_messageInfo_CS_ChapterGetPrize proto.InternalMessageInfo

type SC_ChapterGetPrize struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_ChapterGetPrize) Reset()         { *m = SC_ChapterGetPrize{} }
func (m *SC_ChapterGetPrize) String() string { return proto.CompactTextString(m) }
func (*SC_ChapterGetPrize) ProtoMessage()    {}
func (*SC_ChapterGetPrize) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{15}
}
func (m *SC_ChapterGetPrize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_ChapterGetPrize.Unmarshal(m, b)
}
func (m *SC_ChapterGetPrize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_ChapterGetPrize.Marshal(b, m, deterministic)
}
func (dst *SC_ChapterGetPrize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_ChapterGetPrize.Merge(dst, src)
}
func (m *SC_ChapterGetPrize) XXX_Size() int {
	return xxx_messageInfo_SC_ChapterGetPrize.Size(m)
}
func (m *SC_ChapterGetPrize) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_ChapterGetPrize.DiscardUnknown(m)
}

var xxx_messageInfo_SC_ChapterGetPrize proto.InternalMessageInfo

func (m *SC_ChapterGetPrize) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

// 关卡挑战
type CS_ChapterChallenge struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_ChapterChallenge) Reset()         { *m = CS_ChapterChallenge{} }
func (m *CS_ChapterChallenge) String() string { return proto.CompactTextString(m) }
func (*CS_ChapterChallenge) ProtoMessage()    {}
func (*CS_ChapterChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{16}
}
func (m *CS_ChapterChallenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_ChapterChallenge.Unmarshal(m, b)
}
func (m *CS_ChapterChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_ChapterChallenge.Marshal(b, m, deterministic)
}
func (dst *CS_ChapterChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_ChapterChallenge.Merge(dst, src)
}
func (m *CS_ChapterChallenge) XXX_Size() int {
	return xxx_messageInfo_CS_ChapterChallenge.Size(m)
}
func (m *CS_ChapterChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_ChapterChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_CS_ChapterChallenge proto.InternalMessageInfo

type SC_ChapterChallenge struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_ChapterChallenge) Reset()         { *m = SC_ChapterChallenge{} }
func (m *SC_ChapterChallenge) String() string { return proto.CompactTextString(m) }
func (*SC_ChapterChallenge) ProtoMessage()    {}
func (*SC_ChapterChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{17}
}
func (m *SC_ChapterChallenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_ChapterChallenge.Unmarshal(m, b)
}
func (m *SC_ChapterChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_ChapterChallenge.Marshal(b, m, deterministic)
}
func (dst *SC_ChapterChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_ChapterChallenge.Merge(dst, src)
}
func (m *SC_ChapterChallenge) XXX_Size() int {
	return xxx_messageInfo_SC_ChapterChallenge.Size(m)
}
func (m *SC_ChapterChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_ChapterChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_SC_ChapterChallenge proto.InternalMessageInfo

func (m *SC_ChapterChallenge) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

// p2p消息
type CS_P2P struct {
	From                 int64    `protobuf:"varint,1,opt,name=from" json:"from,omitempty"`
	To                   int64    `protobuf:"varint,2,opt,name=to" json:"to,omitempty"`
	Msg                  []byte   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_P2P) Reset()         { *m = CS_P2P{} }
func (m *CS_P2P) String() string { return proto.CompactTextString(m) }
func (*CS_P2P) ProtoMessage()    {}
func (*CS_P2P) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{18}
}
func (m *CS_P2P) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_P2P.Unmarshal(m, b)
}
func (m *CS_P2P) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_P2P.Marshal(b, m, deterministic)
}
func (dst *CS_P2P) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_P2P.Merge(dst, src)
}
func (m *CS_P2P) XXX_Size() int {
	return xxx_messageInfo_CS_P2P.Size(m)
}
func (m *CS_P2P) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_P2P.DiscardUnknown(m)
}

var xxx_messageInfo_CS_P2P proto.InternalMessageInfo

func (m *CS_P2P) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *CS_P2P) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *CS_P2P) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

// p2p消息
type SC_P2P struct {
	From                 int64    `protobuf:"varint,1,opt,name=from" json:"from,omitempty"`
	To                   int64    `protobuf:"varint,2,opt,name=to" json:"to,omitempty"`
	Msg                  []byte   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_P2P) Reset()         { *m = SC_P2P{} }
func (m *SC_P2P) String() string { return proto.CompactTextString(m) }
func (*SC_P2P) ProtoMessage()    {}
func (*SC_P2P) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{19}
}
func (m *SC_P2P) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_P2P.Unmarshal(m, b)
}
func (m *SC_P2P) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_P2P.Marshal(b, m, deterministic)
}
func (dst *SC_P2P) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_P2P.Merge(dst, src)
}
func (m *SC_P2P) XXX_Size() int {
	return xxx_messageInfo_SC_P2P.Size(m)
}
func (m *SC_P2P) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_P2P.DiscardUnknown(m)
}

var xxx_messageInfo_SC_P2P proto.InternalMessageInfo

func (m *SC_P2P) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *SC_P2P) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *SC_P2P) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type CS_None struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_None) Reset()         { *m = CS_None{} }
func (m *CS_None) String() string { return proto.CompactTextString(m) }
func (*CS_None) ProtoMessage()    {}
func (*CS_None) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_b18e7dfcc50cde1b, []int{20}
}
func (m *CS_None) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_None.Unmarshal(m, b)
}
func (m *CS_None) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_None.Marshal(b, m, deterministic)
}
func (dst *CS_None) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_None.Merge(dst, src)
}
func (m *CS_None) XXX_Size() int {
	return xxx_messageInfo_CS_None.Size(m)
}
func (m *CS_None) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_None.DiscardUnknown(m)
}

var xxx_messageInfo_CS_None proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CS_Tick)(nil), "netproto.CS_Tick")
	proto.RegisterType((*SC_Tick)(nil), "netproto.SC_Tick")
	proto.RegisterType((*CS_Login)(nil), "netproto.CS_Login")
	proto.RegisterType((*SC_Login)(nil), "netproto.SC_Login")
	proto.RegisterType((*CS_ShopAddDiamond)(nil), "netproto.CS_ShopAddDiamond")
	proto.RegisterType((*SC_ShopAddDiamond)(nil), "netproto.SC_ShopAddDiamond")
	proto.RegisterType((*CS_ShopBuyGold)(nil), "netproto.CS_ShopBuyGold")
	proto.RegisterType((*SC_ShopBuyGold)(nil), "netproto.SC_ShopBuyGold")
	proto.RegisterType((*CS_HeroDraw)(nil), "netproto.CS_HeroDraw")
	proto.RegisterType((*SC_HeroDraw)(nil), "netproto.SC_HeroDraw")
	proto.RegisterType((*CS_HeroUplevel)(nil), "netproto.CS_HeroUplevel")
	proto.RegisterType((*SC_HeroUpLevel)(nil), "netproto.SC_HeroUpLevel")
	proto.RegisterType((*CS_HeroEquip)(nil), "netproto.CS_HeroEquip")
	proto.RegisterType((*SC_HeroEquip)(nil), "netproto.SC_HeroEquip")
	proto.RegisterType((*CS_ChapterGetPrize)(nil), "netproto.CS_ChapterGetPrize")
	proto.RegisterType((*SC_ChapterGetPrize)(nil), "netproto.SC_ChapterGetPrize")
	proto.RegisterType((*CS_ChapterChallenge)(nil), "netproto.CS_ChapterChallenge")
	proto.RegisterType((*SC_ChapterChallenge)(nil), "netproto.SC_ChapterChallenge")
	proto.RegisterType((*CS_P2P)(nil), "netproto.CS_P2P")
	proto.RegisterType((*SC_P2P)(nil), "netproto.SC_P2P")
	proto.RegisterType((*CS_None)(nil), "netproto.CS_None")
	proto.RegisterEnum("netproto.E_Code", E_Code_name, E_Code_value)
	proto.RegisterEnum("netproto.EMsgIds", EMsgIds_name, EMsgIds_value)
}

func init() { proto.RegisterFile("logic.proto", fileDescriptor_logic_b18e7dfcc50cde1b) }

var fileDescriptor_logic_b18e7dfcc50cde1b = []byte{
	// 745 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdb, 0x6a, 0xf3, 0x46,
	0x10, 0xfe, 0x65, 0x3b, 0x3e, 0x8c, 0x0f, 0x19, 0xaf, 0x9d, 0xc4, 0xa1, 0x17, 0x0d, 0xa2, 0x85,
	0x90, 0x8b, 0x50, 0x92, 0x52, 0x68, 0x0b, 0xa5, 0xea, 0x7a, 0x71, 0x44, 0x64, 0x49, 0x68, 0x95,
	0xb4, 0xb9, 0x5a, 0xdc, 0x48, 0xb1, 0x4d, 0x6c, 0xad, 0xeb, 0xd8, 0x3d, 0xbd, 0x49, 0x1f, 0xa3,
	0x8f, 0xd1, 0xf3, 0xf9, 0xf8, 0x34, 0x45, 0x5a, 0xf9, 0x90, 0x38, 0x05, 0xc3, 0x7f, 0x37, 0xf3,
	0xcd, 0x37, 0xdf, 0x0c, 0xbb, 0xcc, 0x07, 0xe5, 0x91, 0xec, 0x0f, 0x6f, 0x4f, 0x27, 0x53, 0x39,
	0x93, 0xa4, 0x18, 0x85, 0xb3, 0x24, 0xd2, 0x4b, 0x50, 0xa0, 0x5c, 0xf8, 0xc3, 0xdb, 0xfb, 0x38,
	0xe4, 0x54, 0x85, 0x47, 0x50, 0xa4, 0x5c, 0x58, 0xb2, 0x3f, 0x8c, 0x48, 0x13, 0x76, 0x66, 0xf2,
	0x3e, 0x8c, 0x5a, 0xda, 0x91, 0x76, 0x5c, 0xf2, 0x54, 0xa2, 0xbf, 0x01, 0x45, 0x4e, 0x53, 0xc6,
	0x6b, 0x90, 0xbb, 0x95, 0x41, 0x98, 0x10, 0x6a, 0x67, 0x78, 0xba, 0x10, 0x3f, 0x65, 0x82, 0xca,
	0x20, 0xf4, 0x92, 0xaa, 0xfe, 0x3a, 0xd4, 0x29, 0x17, 0x7c, 0x20, 0x27, 0x46, 0x10, 0xb4, 0x87,
	0xbd, 0xb1, 0x8c, 0x02, 0x82, 0x90, 0x1d, 0x06, 0x9f, 0x27, 0x9d, 0x3b, 0x5e, 0x1c, 0xea, 0x6f,
	0x43, 0x9d, 0xd3, 0xa7, 0xb4, 0xed, 0x26, 0x20, 0xd4, 0xd2, 0x09, 0x1f, 0xcc, 0xbf, 0xe8, 0xc8,
	0x51, 0xa0, 0xbf, 0x05, 0xb5, 0x54, 0x2c, 0x45, 0xb6, 0x54, 0x7a, 0x15, 0xca, 0x94, 0x8b, 0x8b,
	0x70, 0x2a, 0xdb, 0xd3, 0xde, 0x67, 0xf1, 0x96, 0xd1, 0x7c, 0xbc, 0xd8, 0x32, 0x9a, 0x8f, 0xf5,
	0x73, 0x28, 0x73, 0xba, 0x22, 0x6c, 0xa7, 0x7a, 0x9c, 0xec, 0x17, 0x37, 0x5d, 0x4d, 0x46, 0xe1,
	0xa7, 0xe1, 0x88, 0xec, 0x43, 0x7e, 0x10, 0x4e, 0xa5, 0x19, 0xa4, 0xda, 0x69, 0x96, 0xee, 0xad,
	0x98, 0x56, 0xc2, 0xdc, 0x6e, 0xc2, 0xfb, 0x50, 0x49, 0x27, 0xb0, 0x4f, 0xe6, 0xc3, 0xc9, 0xff,
	0xe9, 0x93, 0x16, 0x14, 0xc2, 0x98, 0x60, 0x06, 0xad, 0x4c, 0x52, 0x58, 0xa4, 0xfa, 0x9b, 0x50,
	0x49, 0x27, 0x2b, 0x85, 0xed, 0xe6, 0x36, 0x81, 0x50, 0x2e, 0xe8, 0xa0, 0x37, 0x99, 0x85, 0xd3,
	0x4e, 0x38, 0x73, 0xa7, 0xc3, 0x2f, 0x43, 0xfd, 0x1d, 0x20, 0x9c, 0x3e, 0x45, 0xb7, 0x54, 0xdc,
	0x83, 0xc6, 0x4a, 0x91, 0x0e, 0x7a, 0xa3, 0x51, 0x18, 0xf5, 0x43, 0xfd, 0x5d, 0x68, 0xac, 0x24,
	0x97, 0xf0, 0x96, 0x9a, 0xef, 0x41, 0x9e, 0x72, 0xe1, 0x9e, 0xb9, 0x84, 0x40, 0xee, 0x6e, 0x2a,
	0xd5, 0x8f, 0x66, 0xbd, 0x24, 0x26, 0x35, 0xc8, 0xcc, 0x64, 0xf2, 0x1c, 0x59, 0x2f, 0x33, 0x93,
	0xf1, 0xa7, 0x8f, 0x1f, 0xfa, 0xad, 0xec, 0x91, 0x76, 0x5c, 0xf1, 0xe2, 0x30, 0xee, 0xe7, 0xf4,
	0x25, 0xfa, 0xd5, 0xad, 0xd9, 0x32, 0x0a, 0x4f, 0xbe, 0xca, 0x40, 0x5e, 0xed, 0x46, 0xca, 0x50,
	0x60, 0x82, 0x79, 0x9e, 0xe3, 0xe1, 0x0b, 0x52, 0x84, 0x1c, 0x13, 0xce, 0x25, 0x6a, 0xa4, 0x0a,
	0x25, 0x26, 0x3c, 0x66, 0x39, 0x1d, 0xd3, 0xc6, 0x0c, 0x41, 0xa8, 0x30, 0x61, 0x3b, 0x36, 0x13,
	0xec, 0x23, 0x93, 0xfb, 0x98, 0x55, 0x84, 0x2b, 0xfb, 0xd2, 0x76, 0x3e, 0xb4, 0x31, 0x47, 0x5e,
	0x81, 0x03, 0x26, 0x38, 0xf3, 0xae, 0x99, 0x27, 0x4c, 0xdb, 0x67, 0x9e, 0x6d, 0x58, 0xa9, 0xec,
	0x0e, 0x69, 0xc0, 0x2e, 0x13, 0xa6, 0x7d, 0x6d, 0x58, 0x66, 0x5b, 0xb8, 0x86, 0x67, 0x74, 0x31,
	0x4f, 0xea, 0x50, 0x5d, 0x81, 0x8e, 0xeb, 0x63, 0x41, 0x4d, 0xb1, 0xcc, 0xae, 0xe9, 0x8b, 0x8e,
	0x63, 0xb5, 0xb1, 0xa9, 0x3a, 0x15, 0xd2, 0x36, 0x8d, 0xae, 0x63, 0xb7, 0x71, 0x4f, 0x75, 0x2a,
	0xd0, 0x37, 0xbb, 0x8c, 0xe3, 0x3e, 0x39, 0x84, 0xbd, 0x05, 0x44, 0x2f, 0x0c, 0xd7, 0x67, 0x9e,
	0xb0, 0xd8, 0x35, 0xb3, 0xf0, 0x80, 0xb4, 0xa0, 0xb9, 0x28, 0xb9, 0x96, 0x71, 0xb3, 0xac, 0xb4,
	0x48, 0x13, 0x70, 0xd9, 0x64, 0xb8, 0x06, 0x35, 0xfd, 0x1b, 0x3c, 0x3c, 0xf9, 0x3a, 0x0b, 0x05,
	0xd6, 0x7d, 0xe8, 0x9b, 0xc1, 0x03, 0xa9, 0x40, 0x91, 0xa5, 0x6f, 0x86, 0x2f, 0x92, 0xa7, 0x52,
	0x3f, 0x88, 0xea, 0xdd, 0xd4, 0x77, 0xe0, 0x1d, 0xa9, 0x2a, 0x5e, 0x6c, 0x5e, 0xf8, 0x8d, 0x96,
	0xa4, 0xa9, 0x97, 0xe1, 0xb7, 0x1a, 0xa9, 0x41, 0x89, 0x2d, 0xfc, 0x0c, 0xbf, 0x53, 0xf9, 0xc2,
	0xbd, 0xf0, 0x7b, 0x8d, 0x1c, 0x00, 0x61, 0x1b, 0xde, 0x84, 0x3f, 0xa8, 0xc2, 0x86, 0x1b, 0xe1,
	0x8f, 0x1a, 0x69, 0xc2, 0x2e, 0x7b, 0xec, 0x35, 0xf8, 0x93, 0x42, 0x1f, 0xfb, 0x0d, 0xfe, 0xac,
	0x91, 0x3a, 0x54, 0xd8, 0x9a, 0x9b, 0xe0, 0x2f, 0x0a, 0x5a, 0xf3, 0x0f, 0xfc, 0x75, 0xa9, 0xb8,
	0xe6, 0x0e, 0xf8, 0xdb, 0x52, 0x71, 0xcd, 0x09, 0xf0, 0x77, 0x8d, 0x10, 0xa8, 0xb2, 0xf5, 0x3b,
	0xc7, 0x3f, 0x14, 0xb6, 0x7e, 0xb9, 0xf8, 0xa7, 0x46, 0x5a, 0xd0, 0x60, 0x9b, 0x77, 0x89, 0x7f,
	0xa9, 0xca, 0xe6, 0x6d, 0xe2, 0xdf, 0x1a, 0x39, 0x84, 0x26, 0x7b, 0xe6, 0xf2, 0xf0, 0x1f, 0x55,
	0x7a, 0xe6, 0xfa, 0xf0, 0x5f, 0xed, 0xe3, 0x7c, 0x72, 0x6e, 0xe7, 0xff, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x5a, 0xff, 0x1b, 0x04, 0x66, 0x06, 0x00, 0x00,
}
