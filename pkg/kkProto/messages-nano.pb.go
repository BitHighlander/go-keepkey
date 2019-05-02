// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages-nano.proto

package kkproto

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

// *
// Request: Ask device for Nano address corresponding to address_n path
// @next PassphraseRequest
// @next NanoAddress
// @next Failure
type NanoGetAddress struct {
	AddressN             []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	CoinName             *string  `protobuf:"bytes,2,opt,name=coin_name,json=coinName,def=Nano" json:"coin_name,omitempty"`
	ShowDisplay          *bool    `protobuf:"varint,3,opt,name=show_display,json=showDisplay" json:"show_display,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NanoGetAddress) Reset()         { *m = NanoGetAddress{} }
func (m *NanoGetAddress) String() string { return proto.CompactTextString(m) }
func (*NanoGetAddress) ProtoMessage()    {}
func (*NanoGetAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_nano_c37393c7f260464d, []int{0}
}
func (m *NanoGetAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NanoGetAddress.Unmarshal(m, b)
}
func (m *NanoGetAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NanoGetAddress.Marshal(b, m, deterministic)
}
func (dst *NanoGetAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NanoGetAddress.Merge(dst, src)
}
func (m *NanoGetAddress) XXX_Size() int {
	return xxx_messageInfo_NanoGetAddress.Size(m)
}
func (m *NanoGetAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_NanoGetAddress.DiscardUnknown(m)
}

var xxx_messageInfo_NanoGetAddress proto.InternalMessageInfo

const Default_NanoGetAddress_CoinName string = "Nano"

func (m *NanoGetAddress) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *NanoGetAddress) GetCoinName() string {
	if m != nil && m.CoinName != nil {
		return *m.CoinName
	}
	return Default_NanoGetAddress_CoinName
}

func (m *NanoGetAddress) GetShowDisplay() bool {
	if m != nil && m.ShowDisplay != nil {
		return *m.ShowDisplay
	}
	return false
}

// *
// Response: Contains a Nano address derived from device private seed
// @prev NanoGetAddress
type NanoAddress struct {
	Address              *string  `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NanoAddress) Reset()         { *m = NanoAddress{} }
func (m *NanoAddress) String() string { return proto.CompactTextString(m) }
func (*NanoAddress) ProtoMessage()    {}
func (*NanoAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_nano_c37393c7f260464d, []int{1}
}
func (m *NanoAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NanoAddress.Unmarshal(m, b)
}
func (m *NanoAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NanoAddress.Marshal(b, m, deterministic)
}
func (dst *NanoAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NanoAddress.Merge(dst, src)
}
func (m *NanoAddress) XXX_Size() int {
	return xxx_messageInfo_NanoAddress.Size(m)
}
func (m *NanoAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_NanoAddress.DiscardUnknown(m)
}

var xxx_messageInfo_NanoAddress proto.InternalMessageInfo

func (m *NanoAddress) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

// *
// Request: ask device to sign Nano transaction
// @start
// @next NanoSignedTx
type NanoSignTx struct {
	AddressN             []uint32                `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	CoinName             *string                 `protobuf:"bytes,2,opt,name=coin_name,json=coinName,def=Nano" json:"coin_name,omitempty"`
	ParentBlock          *NanoSignTx_ParentBlock `protobuf:"bytes,3,opt,name=parent_block,json=parentBlock" json:"parent_block,omitempty"`
	LinkHash             []byte                  `protobuf:"bytes,4,opt,name=link_hash,json=linkHash" json:"link_hash,omitempty"`
	LinkRecipient        *string                 `protobuf:"bytes,5,opt,name=link_recipient,json=linkRecipient" json:"link_recipient,omitempty"`
	LinkRecipientN       []uint32                `protobuf:"varint,6,rep,name=link_recipient_n,json=linkRecipientN" json:"link_recipient_n,omitempty"`
	Representative       *string                 `protobuf:"bytes,7,opt,name=representative" json:"representative,omitempty"`
	Balance              []byte                  `protobuf:"bytes,8,opt,name=balance" json:"balance,omitempty"`
	ExchangeType         *ExchangeType           `protobuf:"bytes,9,opt,name=exchange_type,json=exchangeType" json:"exchange_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *NanoSignTx) Reset()         { *m = NanoSignTx{} }
func (m *NanoSignTx) String() string { return proto.CompactTextString(m) }
func (*NanoSignTx) ProtoMessage()    {}
func (*NanoSignTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_nano_c37393c7f260464d, []int{2}
}
func (m *NanoSignTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NanoSignTx.Unmarshal(m, b)
}
func (m *NanoSignTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NanoSignTx.Marshal(b, m, deterministic)
}
func (dst *NanoSignTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NanoSignTx.Merge(dst, src)
}
func (m *NanoSignTx) XXX_Size() int {
	return xxx_messageInfo_NanoSignTx.Size(m)
}
func (m *NanoSignTx) XXX_DiscardUnknown() {
	xxx_messageInfo_NanoSignTx.DiscardUnknown(m)
}

var xxx_messageInfo_NanoSignTx proto.InternalMessageInfo

const Default_NanoSignTx_CoinName string = "Nano"

func (m *NanoSignTx) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *NanoSignTx) GetCoinName() string {
	if m != nil && m.CoinName != nil {
		return *m.CoinName
	}
	return Default_NanoSignTx_CoinName
}

func (m *NanoSignTx) GetParentBlock() *NanoSignTx_ParentBlock {
	if m != nil {
		return m.ParentBlock
	}
	return nil
}

func (m *NanoSignTx) GetLinkHash() []byte {
	if m != nil {
		return m.LinkHash
	}
	return nil
}

func (m *NanoSignTx) GetLinkRecipient() string {
	if m != nil && m.LinkRecipient != nil {
		return *m.LinkRecipient
	}
	return ""
}

func (m *NanoSignTx) GetLinkRecipientN() []uint32 {
	if m != nil {
		return m.LinkRecipientN
	}
	return nil
}

func (m *NanoSignTx) GetRepresentative() string {
	if m != nil && m.Representative != nil {
		return *m.Representative
	}
	return ""
}

func (m *NanoSignTx) GetBalance() []byte {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *NanoSignTx) GetExchangeType() *ExchangeType {
	if m != nil {
		return m.ExchangeType
	}
	return nil
}

type NanoSignTx_ParentBlock struct {
	ParentHash           []byte   `protobuf:"bytes,1,opt,name=parent_hash,json=parentHash" json:"parent_hash,omitempty"`
	Link                 []byte   `protobuf:"bytes,2,opt,name=link" json:"link,omitempty"`
	Representative       *string  `protobuf:"bytes,4,opt,name=representative" json:"representative,omitempty"`
	Balance              []byte   `protobuf:"bytes,5,opt,name=balance" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NanoSignTx_ParentBlock) Reset()         { *m = NanoSignTx_ParentBlock{} }
func (m *NanoSignTx_ParentBlock) String() string { return proto.CompactTextString(m) }
func (*NanoSignTx_ParentBlock) ProtoMessage()    {}
func (*NanoSignTx_ParentBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_nano_c37393c7f260464d, []int{2, 0}
}
func (m *NanoSignTx_ParentBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NanoSignTx_ParentBlock.Unmarshal(m, b)
}
func (m *NanoSignTx_ParentBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NanoSignTx_ParentBlock.Marshal(b, m, deterministic)
}
func (dst *NanoSignTx_ParentBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NanoSignTx_ParentBlock.Merge(dst, src)
}
func (m *NanoSignTx_ParentBlock) XXX_Size() int {
	return xxx_messageInfo_NanoSignTx_ParentBlock.Size(m)
}
func (m *NanoSignTx_ParentBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_NanoSignTx_ParentBlock.DiscardUnknown(m)
}

var xxx_messageInfo_NanoSignTx_ParentBlock proto.InternalMessageInfo

func (m *NanoSignTx_ParentBlock) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *NanoSignTx_ParentBlock) GetLink() []byte {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *NanoSignTx_ParentBlock) GetRepresentative() string {
	if m != nil && m.Representative != nil {
		return *m.Representative
	}
	return ""
}

func (m *NanoSignTx_ParentBlock) GetBalance() []byte {
	if m != nil {
		return m.Balance
	}
	return nil
}

// *
// Response: signature for transaction
// @end
type NanoSignedTx struct {
	Signature            []byte   `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	BlockHash            []byte   `protobuf:"bytes,2,opt,name=block_hash,json=blockHash" json:"block_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NanoSignedTx) Reset()         { *m = NanoSignedTx{} }
func (m *NanoSignedTx) String() string { return proto.CompactTextString(m) }
func (*NanoSignedTx) ProtoMessage()    {}
func (*NanoSignedTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_nano_c37393c7f260464d, []int{3}
}
func (m *NanoSignedTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NanoSignedTx.Unmarshal(m, b)
}
func (m *NanoSignedTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NanoSignedTx.Marshal(b, m, deterministic)
}
func (dst *NanoSignedTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NanoSignedTx.Merge(dst, src)
}
func (m *NanoSignedTx) XXX_Size() int {
	return xxx_messageInfo_NanoSignedTx.Size(m)
}
func (m *NanoSignedTx) XXX_DiscardUnknown() {
	xxx_messageInfo_NanoSignedTx.DiscardUnknown(m)
}

var xxx_messageInfo_NanoSignedTx proto.InternalMessageInfo

func (m *NanoSignedTx) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *NanoSignedTx) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func init() {
	proto.RegisterType((*NanoGetAddress)(nil), "NanoGetAddress")
	proto.RegisterType((*NanoAddress)(nil), "NanoAddress")
	proto.RegisterType((*NanoSignTx)(nil), "NanoSignTx")
	proto.RegisterType((*NanoSignTx_ParentBlock)(nil), "NanoSignTx.ParentBlock")
	proto.RegisterType((*NanoSignedTx)(nil), "NanoSignedTx")
}

func init() { proto.RegisterFile("messages-nano.proto", fileDescriptor_messages_nano_c37393c7f260464d) }

var fileDescriptor_messages_nano_c37393c7f260464d = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x15, 0xda, 0x65, 0x93, 0x49, 0x5a, 0x21, 0x73, 0xc0, 0x2a, 0x20, 0xba, 0x95, 0x80,
	0x5c, 0x88, 0xd0, 0x1e, 0xf7, 0x46, 0x05, 0x02, 0xa9, 0xa2, 0x42, 0xa6, 0xf7, 0xc8, 0x9b, 0x8c,
	0x1a, 0xab, 0x89, 0x6d, 0xc5, 0x61, 0x69, 0xee, 0xbc, 0x31, 0x2f, 0x80, 0x6c, 0x27, 0xea, 0x16,
	0x21, 0x71, 0xd8, 0x9b, 0xe7, 0x9b, 0xb1, 0xff, 0x7f, 0x66, 0x12, 0x78, 0xda, 0xa0, 0x31, 0x7c,
	0x8f, 0xe6, 0x9d, 0xe4, 0x52, 0x65, 0xba, 0x55, 0x9d, 0x5a, 0xc4, 0x5d, 0xaf, 0xd1, 0xf8, 0x60,
	0x65, 0x60, 0xbe, 0xe5, 0x52, 0x7d, 0xc6, 0xee, 0x43, 0x59, 0xb6, 0x68, 0x0c, 0x79, 0x0e, 0x11,
	0xf7, 0xc7, 0x5c, 0xd2, 0x60, 0x39, 0x49, 0x67, 0x2c, 0x1c, 0xc0, 0x96, 0x5c, 0x41, 0x54, 0x28,
	0x21, 0x73, 0xc9, 0x1b, 0xa4, 0x8f, 0x96, 0x41, 0x1a, 0xdd, 0x4c, 0xed, 0x7d, 0x16, 0x5a, 0xbc,
	0xe5, 0x0d, 0x92, 0x2b, 0x48, 0x4c, 0xa5, 0x7e, 0xe6, 0xa5, 0x30, 0xba, 0xe6, 0x3d, 0x9d, 0x2c,
	0x83, 0x34, 0x64, 0xb1, 0x65, 0x1f, 0x3d, 0x5a, 0xbd, 0x85, 0xd8, 0x5e, 0x1a, 0x15, 0x29, 0x5c,
	0x0e, 0x02, 0x34, 0xb0, 0x4f, 0xb2, 0x31, 0x5c, 0xfd, 0x9e, 0x00, 0xd8, 0xca, 0xef, 0x62, 0x2f,
	0x77, 0xc7, 0x07, 0x5b, 0xbb, 0x81, 0x44, 0xf3, 0x16, 0x65, 0x97, 0xdf, 0xd6, 0xaa, 0x38, 0x38,
	0x6b, 0xf1, 0xf5, 0xb3, 0xec, 0x24, 0x91, 0x7d, 0x73, 0xf9, 0xb5, 0x4d, 0xb3, 0x58, 0x9f, 0x02,
	0xab, 0x5d, 0x0b, 0x79, 0xc8, 0x2b, 0x6e, 0x2a, 0x3a, 0x5d, 0x06, 0x69, 0xc2, 0x42, 0x0b, 0xbe,
	0x70, 0x53, 0x91, 0xd7, 0x30, 0x77, 0xc9, 0x16, 0x0b, 0xa1, 0x05, 0xca, 0x8e, 0x5e, 0xb8, 0x46,
	0x66, 0x96, 0xb2, 0x11, 0x92, 0x14, 0x9e, 0x9c, 0x97, 0xe5, 0x92, 0x3e, 0x76, 0x6d, 0xcc, 0xcf,
	0x0a, 0xb7, 0xe4, 0x0d, 0xcc, 0x5b, 0xd4, 0x2d, 0x1a, 0x94, 0x1d, 0xef, 0xc4, 0x1d, 0xd2, 0x4b,
	0xf7, 0xe0, 0x5f, 0xd4, 0x8e, 0xee, 0x96, 0xd7, 0x5c, 0x16, 0x48, 0x43, 0xe7, 0x69, 0x0c, 0xc9,
	0x35, 0xcc, 0xf0, 0x58, 0x54, 0x5c, 0xee, 0x31, 0xb7, 0x0b, 0xa7, 0x91, 0x6b, 0x76, 0x96, 0x7d,
	0x1a, 0xe8, 0xae, 0xd7, 0xc8, 0x12, 0xbc, 0x17, 0x2d, 0x7e, 0x05, 0x10, 0xdf, 0x1b, 0x00, 0x79,
	0x05, 0xc3, 0x08, 0x7c, 0xd7, 0x81, 0x53, 0x00, 0x8f, 0x5c, 0xdf, 0x04, 0xa6, 0xd6, 0xb8, 0x1b,
	0x77, 0xc2, 0xdc, 0xf9, 0x1f, 0xd6, 0xa7, 0xff, 0xb3, 0x7e, 0x71, 0x66, 0x7d, 0xb5, 0x81, 0x64,
	0xdc, 0x08, 0x96, 0xbb, 0x23, 0x79, 0x01, 0x91, 0x11, 0x7b, 0xc9, 0xbb, 0x1f, 0x2d, 0x0e, 0x26,
	0x4e, 0x80, 0xbc, 0x04, 0x70, 0xdb, 0xf4, 0x1e, 0xbd, 0x93, 0xc8, 0x11, 0x6b, 0x71, 0xfd, 0x1e,
	0x16, 0x85, 0x6a, 0xb2, 0x03, 0xa2, 0x3e, 0x60, 0x9f, 0x95, 0x78, 0x27, 0x0a, 0x74, 0xdf, 0x7e,
	0xa1, 0xea, 0x35, 0xd9, 0x20, 0xea, 0x0d, 0xf6, 0x5f, 0xfd, 0x7f, 0x62, 0x65, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x4b, 0xb9, 0x44, 0xe3, 0x35, 0x03, 0x00, 0x00,
}
