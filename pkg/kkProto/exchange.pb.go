// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exchange.proto

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
// Structure representing address for various coin types (BTC, LTC, and etc).
// @used in ExchangeResponse
type ExchangeAddress struct {
	CoinType             *string  `protobuf:"bytes,1,opt,name=coin_type,json=coinType" json:"coin_type,omitempty"`
	Address              *string  `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	DestTag              *string  `protobuf:"bytes,3,opt,name=dest_tag,json=destTag" json:"dest_tag,omitempty"`
	RsAddress            *string  `protobuf:"bytes,4,opt,name=rs_address,json=rsAddress" json:"rs_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExchangeAddress) Reset()         { *m = ExchangeAddress{} }
func (m *ExchangeAddress) String() string { return proto.CompactTextString(m) }
func (*ExchangeAddress) ProtoMessage()    {}
func (*ExchangeAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_04885ab31447f0fd, []int{0}
}
func (m *ExchangeAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeAddress.Unmarshal(m, b)
}
func (m *ExchangeAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeAddress.Marshal(b, m, deterministic)
}
func (dst *ExchangeAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeAddress.Merge(dst, src)
}
func (m *ExchangeAddress) XXX_Size() int {
	return xxx_messageInfo_ExchangeAddress.Size(m)
}
func (m *ExchangeAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeAddress.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeAddress proto.InternalMessageInfo

func (m *ExchangeAddress) GetCoinType() string {
	if m != nil && m.CoinType != nil {
		return *m.CoinType
	}
	return ""
}

func (m *ExchangeAddress) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *ExchangeAddress) GetDestTag() string {
	if m != nil && m.DestTag != nil {
		return *m.DestTag
	}
	return ""
}

func (m *ExchangeAddress) GetRsAddress() string {
	if m != nil && m.RsAddress != nil {
		return *m.RsAddress
	}
	return ""
}

// *
// Structure representing exchange response version 2
type ExchangeResponseV2 struct {
	DepositAddress       *ExchangeAddress `protobuf:"bytes,1,opt,name=deposit_address,json=depositAddress" json:"deposit_address,omitempty"`
	DepositAmount        []byte           `protobuf:"bytes,2,opt,name=deposit_amount,json=depositAmount" json:"deposit_amount,omitempty"`
	Expiration           *int64           `protobuf:"varint,3,opt,name=expiration" json:"expiration,omitempty"`
	QuotedRate           []byte           `protobuf:"bytes,4,opt,name=quoted_rate,json=quotedRate" json:"quoted_rate,omitempty"`
	WithdrawalAddress    *ExchangeAddress `protobuf:"bytes,5,opt,name=withdrawal_address,json=withdrawalAddress" json:"withdrawal_address,omitempty"`
	WithdrawalAmount     []byte           `protobuf:"bytes,6,opt,name=withdrawal_amount,json=withdrawalAmount" json:"withdrawal_amount,omitempty"`
	ReturnAddress        *ExchangeAddress `protobuf:"bytes,7,opt,name=return_address,json=returnAddress" json:"return_address,omitempty"`
	ApiKey               []byte           `protobuf:"bytes,8,opt,name=api_key,json=apiKey" json:"api_key,omitempty"`
	MinerFee             []byte           `protobuf:"bytes,9,opt,name=miner_fee,json=minerFee" json:"miner_fee,omitempty"`
	OrderId              []byte           `protobuf:"bytes,10,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ExchangeResponseV2) Reset()         { *m = ExchangeResponseV2{} }
func (m *ExchangeResponseV2) String() string { return proto.CompactTextString(m) }
func (*ExchangeResponseV2) ProtoMessage()    {}
func (*ExchangeResponseV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_04885ab31447f0fd, []int{1}
}
func (m *ExchangeResponseV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeResponseV2.Unmarshal(m, b)
}
func (m *ExchangeResponseV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeResponseV2.Marshal(b, m, deterministic)
}
func (dst *ExchangeResponseV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeResponseV2.Merge(dst, src)
}
func (m *ExchangeResponseV2) XXX_Size() int {
	return xxx_messageInfo_ExchangeResponseV2.Size(m)
}
func (m *ExchangeResponseV2) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeResponseV2.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeResponseV2 proto.InternalMessageInfo

func (m *ExchangeResponseV2) GetDepositAddress() *ExchangeAddress {
	if m != nil {
		return m.DepositAddress
	}
	return nil
}

func (m *ExchangeResponseV2) GetDepositAmount() []byte {
	if m != nil {
		return m.DepositAmount
	}
	return nil
}

func (m *ExchangeResponseV2) GetExpiration() int64 {
	if m != nil && m.Expiration != nil {
		return *m.Expiration
	}
	return 0
}

func (m *ExchangeResponseV2) GetQuotedRate() []byte {
	if m != nil {
		return m.QuotedRate
	}
	return nil
}

func (m *ExchangeResponseV2) GetWithdrawalAddress() *ExchangeAddress {
	if m != nil {
		return m.WithdrawalAddress
	}
	return nil
}

func (m *ExchangeResponseV2) GetWithdrawalAmount() []byte {
	if m != nil {
		return m.WithdrawalAmount
	}
	return nil
}

func (m *ExchangeResponseV2) GetReturnAddress() *ExchangeAddress {
	if m != nil {
		return m.ReturnAddress
	}
	return nil
}

func (m *ExchangeResponseV2) GetApiKey() []byte {
	if m != nil {
		return m.ApiKey
	}
	return nil
}

func (m *ExchangeResponseV2) GetMinerFee() []byte {
	if m != nil {
		return m.MinerFee
	}
	return nil
}

func (m *ExchangeResponseV2) GetOrderId() []byte {
	if m != nil {
		return m.OrderId
	}
	return nil
}

type SignedExchangeResponse struct {
	Response             *ExchangeResponse   `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	Signature            []byte              `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	ResponseV2           *ExchangeResponseV2 `protobuf:"bytes,3,opt,name=responseV2" json:"responseV2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SignedExchangeResponse) Reset()         { *m = SignedExchangeResponse{} }
func (m *SignedExchangeResponse) String() string { return proto.CompactTextString(m) }
func (*SignedExchangeResponse) ProtoMessage()    {}
func (*SignedExchangeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_04885ab31447f0fd, []int{2}
}
func (m *SignedExchangeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedExchangeResponse.Unmarshal(m, b)
}
func (m *SignedExchangeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedExchangeResponse.Marshal(b, m, deterministic)
}
func (dst *SignedExchangeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedExchangeResponse.Merge(dst, src)
}
func (m *SignedExchangeResponse) XXX_Size() int {
	return xxx_messageInfo_SignedExchangeResponse.Size(m)
}
func (m *SignedExchangeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedExchangeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignedExchangeResponse proto.InternalMessageInfo

func (m *SignedExchangeResponse) GetResponse() *ExchangeResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *SignedExchangeResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignedExchangeResponse) GetResponseV2() *ExchangeResponseV2 {
	if m != nil {
		return m.ResponseV2
	}
	return nil
}

// *
// Structure representing exchange response version 1 (deprecated)
type ExchangeResponse struct {
	DepositAddress       *ExchangeAddress `protobuf:"bytes,1,opt,name=deposit_address,json=depositAddress" json:"deposit_address,omitempty"`
	DepositAmount        *uint64          `protobuf:"varint,2,opt,name=deposit_amount,json=depositAmount" json:"deposit_amount,omitempty"`
	Expiration           *int64           `protobuf:"varint,3,opt,name=expiration" json:"expiration,omitempty"`
	QuotedRate           *uint64          `protobuf:"varint,4,opt,name=quoted_rate,json=quotedRate" json:"quoted_rate,omitempty"`
	WithdrawalAddress    *ExchangeAddress `protobuf:"bytes,5,opt,name=withdrawal_address,json=withdrawalAddress" json:"withdrawal_address,omitempty"`
	WithdrawalAmount     *uint64          `protobuf:"varint,6,opt,name=withdrawal_amount,json=withdrawalAmount" json:"withdrawal_amount,omitempty"`
	ReturnAddress        *ExchangeAddress `protobuf:"bytes,7,opt,name=return_address,json=returnAddress" json:"return_address,omitempty"`
	ApiKey               []byte           `protobuf:"bytes,8,opt,name=api_key,json=apiKey" json:"api_key,omitempty"`
	MinerFee             *uint64          `protobuf:"varint,9,opt,name=miner_fee,json=minerFee" json:"miner_fee,omitempty"`
	OrderId              []byte           `protobuf:"bytes,10,opt,name=order_id,json=orderId" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ExchangeResponse) Reset()         { *m = ExchangeResponse{} }
func (m *ExchangeResponse) String() string { return proto.CompactTextString(m) }
func (*ExchangeResponse) ProtoMessage()    {}
func (*ExchangeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_exchange_04885ab31447f0fd, []int{3}
}
func (m *ExchangeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeResponse.Unmarshal(m, b)
}
func (m *ExchangeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeResponse.Marshal(b, m, deterministic)
}
func (dst *ExchangeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeResponse.Merge(dst, src)
}
func (m *ExchangeResponse) XXX_Size() int {
	return xxx_messageInfo_ExchangeResponse.Size(m)
}
func (m *ExchangeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeResponse proto.InternalMessageInfo

func (m *ExchangeResponse) GetDepositAddress() *ExchangeAddress {
	if m != nil {
		return m.DepositAddress
	}
	return nil
}

func (m *ExchangeResponse) GetDepositAmount() uint64 {
	if m != nil && m.DepositAmount != nil {
		return *m.DepositAmount
	}
	return 0
}

func (m *ExchangeResponse) GetExpiration() int64 {
	if m != nil && m.Expiration != nil {
		return *m.Expiration
	}
	return 0
}

func (m *ExchangeResponse) GetQuotedRate() uint64 {
	if m != nil && m.QuotedRate != nil {
		return *m.QuotedRate
	}
	return 0
}

func (m *ExchangeResponse) GetWithdrawalAddress() *ExchangeAddress {
	if m != nil {
		return m.WithdrawalAddress
	}
	return nil
}

func (m *ExchangeResponse) GetWithdrawalAmount() uint64 {
	if m != nil && m.WithdrawalAmount != nil {
		return *m.WithdrawalAmount
	}
	return 0
}

func (m *ExchangeResponse) GetReturnAddress() *ExchangeAddress {
	if m != nil {
		return m.ReturnAddress
	}
	return nil
}

func (m *ExchangeResponse) GetApiKey() []byte {
	if m != nil {
		return m.ApiKey
	}
	return nil
}

func (m *ExchangeResponse) GetMinerFee() uint64 {
	if m != nil && m.MinerFee != nil {
		return *m.MinerFee
	}
	return 0
}

func (m *ExchangeResponse) GetOrderId() []byte {
	if m != nil {
		return m.OrderId
	}
	return nil
}

func init() {
	proto.RegisterType((*ExchangeAddress)(nil), "ExchangeAddress")
	proto.RegisterType((*ExchangeResponseV2)(nil), "ExchangeResponseV2")
	proto.RegisterType((*SignedExchangeResponse)(nil), "SignedExchangeResponse")
	proto.RegisterType((*ExchangeResponse)(nil), "ExchangeResponse")
}

func init() { proto.RegisterFile("exchange.proto", fileDescriptor_exchange_04885ab31447f0fd) }

var fileDescriptor_exchange_04885ab31447f0fd = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0x4f, 0x6b, 0x14, 0x31,
	0x18, 0xc6, 0x59, 0x77, 0xec, 0xce, 0xbc, 0xdb, 0xee, 0x9f, 0x08, 0x3a, 0x52, 0xff, 0xb1, 0x20,
	0x08, 0xd2, 0x39, 0xac, 0x07, 0xf1, 0x24, 0x16, 0x14, 0x64, 0x6f, 0xb1, 0x78, 0x1d, 0xc2, 0xe4,
	0x75, 0x1a, 0xb6, 0x9b, 0xc4, 0x24, 0x6b, 0x3b, 0x57, 0x3f, 0x86, 0xdf, 0x41, 0xfc, 0x8a, 0x32,
	0x99, 0xc9, 0x74, 0xd8, 0x52, 0xf0, 0xb0, 0xd8, 0xdb, 0xe6, 0x79, 0x9f, 0xe4, 0xfd, 0x2d, 0xcf,
	0xc3, 0xc0, 0x04, 0xaf, 0x8a, 0x73, 0x26, 0x4b, 0xcc, 0xb4, 0x51, 0x4e, 0x2d, 0x7e, 0x0e, 0x60,
	0xfa, 0xb1, 0x95, 0x3e, 0x70, 0x6e, 0xd0, 0x5a, 0x72, 0x0c, 0x49, 0xa1, 0x84, 0xcc, 0x5d, 0xa5,
	0x31, 0x1d, 0xbc, 0x18, 0xbc, 0x4a, 0x68, 0x5c, 0x0b, 0x67, 0x95, 0x46, 0x92, 0xc2, 0x88, 0x35,
	0xbe, 0xf4, 0x9e, 0x1f, 0x85, 0x23, 0x79, 0x0c, 0x31, 0x47, 0xeb, 0x72, 0xc7, 0xca, 0x74, 0xd8,
	0x8c, 0xea, 0xf3, 0x19, 0x2b, 0xc9, 0x53, 0x00, 0x63, 0xf3, 0x70, 0x2f, 0xf2, 0xc3, 0xc4, 0xd8,
	0x76, 0xe1, 0xe2, 0xcf, 0x10, 0x48, 0x80, 0xa0, 0x68, 0xb5, 0x92, 0x16, 0xbf, 0x2e, 0xc9, 0x3b,
	0x98, 0x72, 0xd4, 0xca, 0x0a, 0xd7, 0x5d, 0xad, 0x69, 0xc6, 0xcb, 0x59, 0xb6, 0x83, 0x4c, 0x27,
	0xad, 0x31, 0xfc, 0x85, 0x97, 0x30, 0xe9, 0xae, 0x6e, 0xd4, 0x56, 0x3a, 0x0f, 0x7b, 0x48, 0x8f,
	0x82, 0xcf, 0x8b, 0xe4, 0x19, 0x00, 0x5e, 0x69, 0x61, 0x98, 0x13, 0x4a, 0x7a, 0xe8, 0x21, 0xed,
	0x29, 0xe4, 0x39, 0x8c, 0xbf, 0x6f, 0x95, 0x43, 0x9e, 0x1b, 0xe6, 0xd0, 0x83, 0x1f, 0x52, 0x68,
	0x24, 0xca, 0x1c, 0x92, 0xf7, 0x40, 0x2e, 0x85, 0x3b, 0xe7, 0x86, 0x5d, 0xb2, 0x8b, 0x8e, 0xf2,
	0xfe, 0x2d, 0x94, 0xf3, 0x6b, 0x6f, 0x00, 0x7d, 0x0d, 0xf3, 0xfe, 0x03, 0x0d, 0xeb, 0x81, 0xdf,
	0x33, 0xeb, 0xb9, 0x1b, 0xdc, 0xb7, 0x30, 0x31, 0xe8, 0xb6, 0x46, 0x76, 0x9b, 0x46, 0xb7, 0x6c,
	0x3a, 0x6a, 0x7c, 0x61, 0xcb, 0x23, 0x18, 0x31, 0x2d, 0xf2, 0x35, 0x56, 0x69, 0xec, 0xdf, 0x3e,
	0x60, 0x5a, 0xac, 0xb0, 0xaa, 0xa3, 0xde, 0x08, 0x89, 0x26, 0xff, 0x86, 0x98, 0x26, 0x7e, 0x14,
	0x7b, 0xe1, 0x13, 0x62, 0x1d, 0xa8, 0x32, 0x1c, 0x4d, 0x2e, 0x78, 0x0a, 0x7e, 0x36, 0xf2, 0xe7,
	0xcf, 0x7c, 0xf1, 0x6b, 0x00, 0x0f, 0xbf, 0x88, 0x52, 0x22, 0xdf, 0xcd, 0x8d, 0x9c, 0x40, 0x6c,
	0xda, 0xdf, 0x6d, 0x5c, 0xf3, 0x6c, 0xd7, 0x44, 0x3b, 0x0b, 0x79, 0x02, 0x89, 0x15, 0xa5, 0x64,
	0x6e, 0x6b, 0xb0, 0x0d, 0xe9, 0x5a, 0x20, 0x6f, 0x00, 0x4c, 0x57, 0x08, 0x1f, 0xd0, 0x78, 0xf9,
	0x20, 0xbb, 0xd9, 0x15, 0xda, 0xb3, 0x2d, 0x7e, 0x0f, 0x61, 0x76, 0x03, 0x6b, 0xef, 0x65, 0x8a,
	0xf6, 0x50, 0xa6, 0xe8, 0x3f, 0x95, 0x29, 0xba, 0x9b, 0x32, 0x45, 0xff, 0x54, 0xa6, 0xd3, 0x0c,
	0x8e, 0x0b, 0xb5, 0xc9, 0xd6, 0x88, 0x7a, 0x8d, 0x55, 0xc6, 0xf1, 0x87, 0x28, 0xf0, 0xc4, 0x7f,
	0x9f, 0x0a, 0x75, 0x71, 0x3a, 0x5d, 0x21, 0xea, 0x15, 0x56, 0x01, 0xeb, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x8a, 0x52, 0x6d, 0x36, 0xc4, 0x04, 0x00, 0x00,
}
