// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/feegrant/v1beta1/feegrant.proto

package feegrant

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	types1 "github.com/lightmos/lightmos-sdk/codec/types"
	github_com_lightmos_lightmos_sdk_types "github.com/lightmos/lightmos-sdk/types"
	types "github.com/lightmos/lightmos-sdk/types"
	_ "github.com/lightmos/lightmos-sdk/types/tx/amino"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BasicAllowance implements Allowance with a one-time grant of coins
// that optionally expires. The grantee can use up to SpendLimit to cover fees.
type BasicAllowance struct {
	// spend_limit specifies the maximum amount of coins that can be spent
	// by this allowance and will be updated as coins are spent. If it is
	// empty, there is no spend limit and any amount of coins can be spent.
	SpendLimit github_com_lightmos_lightmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=spend_limit,json=spendLimit,proto3,castrepeated=github.com/lightmos/lightmos-sdk/types.Coins" json:"spend_limit"`
	// expiration specifies an optional time when this allowance expires
	Expiration *time.Time `protobuf:"bytes,2,opt,name=expiration,proto3,stdtime" json:"expiration,omitempty"`
}

func (m *BasicAllowance) Reset()         { *m = BasicAllowance{} }
func (m *BasicAllowance) String() string { return proto.CompactTextString(m) }
func (*BasicAllowance) ProtoMessage()    {}
func (*BasicAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_7279582900c30aea, []int{0}
}
func (m *BasicAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasicAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasicAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasicAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasicAllowance.Merge(m, src)
}
func (m *BasicAllowance) XXX_Size() int {
	return m.Size()
}
func (m *BasicAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_BasicAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_BasicAllowance proto.InternalMessageInfo

func (m *BasicAllowance) GetSpendLimit() github_com_lightmos_lightmos_sdk_types.Coins {
	if m != nil {
		return m.SpendLimit
	}
	return nil
}

func (m *BasicAllowance) GetExpiration() *time.Time {
	if m != nil {
		return m.Expiration
	}
	return nil
}

// PeriodicAllowance extends Allowance to allow for both a maximum cap,
// as well as a limit per time period.
type PeriodicAllowance struct {
	// basic specifies a struct of `BasicAllowance`
	Basic BasicAllowance `protobuf:"bytes,1,opt,name=basic,proto3" json:"basic"`
	// period specifies the time duration in which period_spend_limit coins can
	// be spent before that allowance is reset
	Period time.Duration `protobuf:"bytes,2,opt,name=period,proto3,stdduration" json:"period"`
	// period_spend_limit specifies the maximum number of coins that can be spent
	// in the period
	PeriodSpendLimit github_com_lightmos_lightmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=period_spend_limit,json=periodSpendLimit,proto3,castrepeated=github.com/lightmos/lightmos-sdk/types.Coins" json:"period_spend_limit"`
	// period_can_spend is the number of coins left to be spent before the period_reset time
	PeriodCanSpend github_com_lightmos_lightmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=period_can_spend,json=periodCanSpend,proto3,castrepeated=github.com/lightmos/lightmos-sdk/types.Coins" json:"period_can_spend"`
	// period_reset is the time at which this period resets and a new one begins,
	// it is calculated from the start time of the first transaction after the
	// last period ended
	PeriodReset time.Time `protobuf:"bytes,5,opt,name=period_reset,json=periodReset,proto3,stdtime" json:"period_reset"`
}

func (m *PeriodicAllowance) Reset()         { *m = PeriodicAllowance{} }
func (m *PeriodicAllowance) String() string { return proto.CompactTextString(m) }
func (*PeriodicAllowance) ProtoMessage()    {}
func (*PeriodicAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_7279582900c30aea, []int{1}
}
func (m *PeriodicAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PeriodicAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PeriodicAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PeriodicAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeriodicAllowance.Merge(m, src)
}
func (m *PeriodicAllowance) XXX_Size() int {
	return m.Size()
}
func (m *PeriodicAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_PeriodicAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_PeriodicAllowance proto.InternalMessageInfo

func (m *PeriodicAllowance) GetBasic() BasicAllowance {
	if m != nil {
		return m.Basic
	}
	return BasicAllowance{}
}

func (m *PeriodicAllowance) GetPeriod() time.Duration {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *PeriodicAllowance) GetPeriodSpendLimit() github_com_lightmos_lightmos_sdk_types.Coins {
	if m != nil {
		return m.PeriodSpendLimit
	}
	return nil
}

func (m *PeriodicAllowance) GetPeriodCanSpend() github_com_lightmos_lightmos_sdk_types.Coins {
	if m != nil {
		return m.PeriodCanSpend
	}
	return nil
}

func (m *PeriodicAllowance) GetPeriodReset() time.Time {
	if m != nil {
		return m.PeriodReset
	}
	return time.Time{}
}

// AllowedMsgAllowance creates allowance only for specified message types.
type AllowedMsgAllowance struct {
	// allowance can be any of basic and periodic fee allowance.
	Allowance *types1.Any `protobuf:"bytes,1,opt,name=allowance,proto3" json:"allowance,omitempty"`
	// allowed_messages are the messages for which the grantee has the access.
	AllowedMessages []string `protobuf:"bytes,2,rep,name=allowed_messages,json=allowedMessages,proto3" json:"allowed_messages,omitempty"`
}

func (m *AllowedMsgAllowance) Reset()         { *m = AllowedMsgAllowance{} }
func (m *AllowedMsgAllowance) String() string { return proto.CompactTextString(m) }
func (*AllowedMsgAllowance) ProtoMessage()    {}
func (*AllowedMsgAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_7279582900c30aea, []int{2}
}
func (m *AllowedMsgAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedMsgAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedMsgAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedMsgAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedMsgAllowance.Merge(m, src)
}
func (m *AllowedMsgAllowance) XXX_Size() int {
	return m.Size()
}
func (m *AllowedMsgAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedMsgAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedMsgAllowance proto.InternalMessageInfo

// Grant is stored in the KVStore to record a grant with full context
type Grant struct {
	// granter is the address of the user granting an allowance of their funds.
	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	// grantee is the address of the user being granted an allowance of another user's funds.
	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// allowance can be any of basic, periodic, allowed fee allowance.
	Allowance *types1.Any `protobuf:"bytes,3,opt,name=allowance,proto3" json:"allowance,omitempty"`
}

func (m *Grant) Reset()         { *m = Grant{} }
func (m *Grant) String() string { return proto.CompactTextString(m) }
func (*Grant) ProtoMessage()    {}
func (*Grant) Descriptor() ([]byte, []int) {
	return fileDescriptor_7279582900c30aea, []int{3}
}
func (m *Grant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Grant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Grant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Grant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Grant.Merge(m, src)
}
func (m *Grant) XXX_Size() int {
	return m.Size()
}
func (m *Grant) XXX_DiscardUnknown() {
	xxx_messageInfo_Grant.DiscardUnknown(m)
}

var xxx_messageInfo_Grant proto.InternalMessageInfo

func (m *Grant) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *Grant) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *Grant) GetAllowance() *types1.Any {
	if m != nil {
		return m.Allowance
	}
	return nil
}

func init() {
	proto.RegisterType((*BasicAllowance)(nil), "cosmos.feegrant.v1beta1.BasicAllowance")
	proto.RegisterType((*PeriodicAllowance)(nil), "cosmos.feegrant.v1beta1.PeriodicAllowance")
	proto.RegisterType((*AllowedMsgAllowance)(nil), "cosmos.feegrant.v1beta1.AllowedMsgAllowance")
	proto.RegisterType((*Grant)(nil), "cosmos.feegrant.v1beta1.Grant")
}

func init() {
	proto.RegisterFile("cosmos/feegrant/v1beta1/feegrant.proto", fileDescriptor_7279582900c30aea)
}

var fileDescriptor_7279582900c30aea = []byte{
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x3f, 0x4f, 0x14, 0x41,
	0x14, 0xbf, 0xe1, 0x00, 0x73, 0x73, 0x8a, 0xb0, 0x92, 0xb8, 0x10, 0xb3, 0x4b, 0x2e, 0x51, 0x01,
	0x65, 0x37, 0x60, 0x25, 0x15, 0x2c, 0x2a, 0x6a, 0x20, 0x21, 0x8b, 0x95, 0x89, 0xb9, 0xcc, 0xee,
	0x0e, 0xcb, 0xc4, 0xdb, 0x9d, 0xcd, 0xce, 0x9c, 0x42, 0x63, 0x61, 0x65, 0xac, 0x88, 0x95, 0xa5,
	0xa5, 0xb1, 0xa2, 0xa0, 0xb6, 0x26, 0x16, 0x86, 0x58, 0x59, 0x89, 0x81, 0x82, 0xda, 0x6f, 0x60,
	0x76, 0x66, 0x76, 0xef, 0xb8, 0x93, 0x00, 0x89, 0x34, 0x97, 0x99, 0x37, 0xef, 0xfd, 0xfe, 0xbc,
	0xf7, 0x2e, 0x0b, 0x6f, 0xf9, 0x94, 0x45, 0x94, 0xd9, 0x6b, 0x18, 0x87, 0x29, 0x8a, 0xb9, 0xfd,
	0x6a, 0xda, 0xc3, 0x1c, 0x4d, 0x17, 0x01, 0x2b, 0x49, 0x29, 0xa7, 0xda, 0x75, 0x99, 0x67, 0x15,
	0x61, 0x95, 0x37, 0x3a, 0x1c, 0xd2, 0x90, 0x8a, 0x1c, 0x3b, 0x3b, 0xc9, 0xf4, 0xd1, 0x91, 0x90,
	0xd2, 0xb0, 0x81, 0x6d, 0x71, 0xf3, 0x9a, 0x6b, 0x36, 0x8a, 0x37, 0xf3, 0x27, 0x89, 0x54, 0x97,
	0x35, 0x0a, 0x56, 0x3e, 0x19, 0x4a, 0x8c, 0x87, 0x18, 0x2e, 0x84, 0xf8, 0x94, 0xc4, 0xea, 0x7d,
	0x08, 0x45, 0x24, 0xa6, 0xb6, 0xf8, 0x55, 0x21, 0xb3, 0x93, 0x88, 0x93, 0x08, 0x33, 0x8e, 0xa2,
	0x24, 0xc7, 0xec, 0x4c, 0x08, 0x9a, 0x29, 0xe2, 0x84, 0x2a, 0xcc, 0xda, 0x87, 0x1e, 0x38, 0xe0,
	0x20, 0x46, 0xfc, 0xf9, 0x46, 0x83, 0xbe, 0x46, 0xb1, 0x8f, 0xb5, 0x26, 0xac, 0xb2, 0x04, 0xc7,
	0x41, 0xbd, 0x41, 0x22, 0xc2, 0x75, 0x30, 0x56, 0x1e, 0xaf, 0xce, 0x8c, 0x58, 0x4a, 0x6a, 0x26,
	0x2e, 0x77, 0x6f, 0x2d, 0x50, 0x12, 0x3b, 0xf7, 0x77, 0x7f, 0x99, 0xa5, 0x2f, 0xfb, 0xe6, 0xdd,
	0x90, 0xf0, 0xf5, 0xa6, 0x67, 0xf9, 0x34, 0xb2, 0x1b, 0x24, 0x5c, 0xe7, 0x99, 0x97, 0xfc, 0x30,
	0xc5, 0x82, 0x97, 0x36, 0xdf, 0x4c, 0x30, 0x13, 0x45, 0xec, 0xf3, 0xd1, 0xf6, 0x24, 0x70, 0xa1,
	0x20, 0x5a, 0xca, 0x78, 0xb4, 0x39, 0x08, 0xf1, 0x46, 0x42, 0xa4, 0x3a, 0xbd, 0x67, 0x0c, 0x8c,
	0x57, 0x67, 0x46, 0x2d, 0x29, 0xdf, 0xca, 0xe5, 0x5b, 0xcf, 0x72, 0x7f, 0x4e, 0xef, 0xd6, 0xbe,
	0x09, 0xdc, 0xb6, 0x9a, 0xd9, 0xc5, 0x6f, 0x3b, 0x53, 0x37, 0x4f, 0x18, 0x94, 0xf5, 0x08, 0xe3,
	0xc2, 0xe2, 0x93, 0xf7, 0x47, 0xdb, 0x93, 0x6a, 0x10, 0x42, 0xd8, 0xf1, 0x0e, 0xd4, 0xbe, 0xf6,
	0xc2, 0xa1, 0x15, 0x9c, 0x12, 0x1a, 0xb4, 0xf7, 0xe5, 0x31, 0xec, 0xf3, 0xb2, 0x3c, 0x1d, 0x08,
	0x6d, 0xb7, 0xad, 0x93, 0xa8, 0x8e, 0xa3, 0x39, 0x95, 0xac, 0x3f, 0xd2, 0xaf, 0x04, 0xd0, 0xe6,
	0x60, 0x7f, 0x22, 0xe0, 0x95, 0xcd, 0x91, 0x2e, 0x9b, 0x0f, 0xd4, 0x94, 0x9c, 0x2b, 0x59, 0xf1,
	0xc7, 0x7d, 0x13, 0x48, 0x00, 0x55, 0xa7, 0xbd, 0x05, 0x50, 0x93, 0xc7, 0x7a, 0xfb, 0xac, 0xca,
	0x17, 0x38, 0xab, 0x41, 0xc9, 0xb7, 0xda, 0x9a, 0xd8, 0x1b, 0xa8, 0x62, 0x75, 0x1f, 0xc5, 0x52,
	0x87, 0xde, 0x7b, 0x81, 0x0a, 0x06, 0x24, 0xdb, 0x02, 0x8a, 0x85, 0x08, 0x6d, 0x09, 0x5e, 0x56,
	0xfc, 0x29, 0x66, 0x98, 0xeb, 0x7d, 0xa7, 0xee, 0x8c, 0xe8, 0xe6, 0x56, 0xd1, 0xcd, 0xaa, 0x2c,
	0x77, 0xb3, 0xea, 0xd9, 0xa7, 0xe7, 0xda, 0x9e, 0x1b, 0x6d, 0xdb, 0xd3, 0xb5, 0x2a, 0xb5, 0x3f,
	0x00, 0x5e, 0x13, 0x37, 0x1c, 0x2c, 0xb3, 0xb0, 0xb5, 0x42, 0x2f, 0x60, 0x05, 0xe5, 0x17, 0xb5,
	0x46, 0xc3, 0x5d, 0x72, 0xe7, 0xe3, 0x4d, 0x67, 0xe2, 0xcc, 0x62, 0xdc, 0x16, 0xa2, 0x36, 0x01,
	0x07, 0x91, 0x64, 0xad, 0x47, 0x98, 0x31, 0x14, 0x62, 0xa6, 0xf7, 0x8c, 0x95, 0xc7, 0x2b, 0xee,
	0x55, 0x15, 0x5f, 0x56, 0xe1, 0xd9, 0x95, 0x77, 0x9f, 0xcc, 0xd2, 0xb9, 0x1c, 0x1b, 0x6d, 0x8e,
	0xff, 0xe1, 0xad, 0xf6, 0x1d, 0xc0, 0xbe, 0xc5, 0x0c, 0x42, 0x9b, 0x81, 0x97, 0x04, 0x16, 0x4e,
	0x85, 0xc7, 0x8a, 0xa3, 0xff, 0xd8, 0x99, 0x1a, 0x56, 0x44, 0xf3, 0x41, 0x90, 0x62, 0xc6, 0x56,
	0x79, 0x4a, 0xe2, 0xd0, 0xcd, 0x13, 0x5b, 0x35, 0x58, 0xfc, 0x27, 0xce, 0x50, 0xd3, 0xd1, 0xcd,
	0xf2, 0xff, 0xee, 0xa6, 0xf3, 0x70, 0xf7, 0xc0, 0x00, 0x7b, 0x07, 0x06, 0xf8, 0x7d, 0x60, 0x80,
	0xad, 0x43, 0xa3, 0xb4, 0x77, 0x68, 0x94, 0x7e, 0x1e, 0x1a, 0xa5, 0xe7, 0x77, 0x4e, 0xdd, 0xdd,
	0x8d, 0xe2, 0x03, 0xe2, 0xf5, 0x0b, 0x29, 0xf7, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x82, 0x34,
	0x40, 0xe4, 0x6b, 0x06, 0x00, 0x00,
}

func (m *BasicAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasicAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasicAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Expiration != nil {
		n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.Expiration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Expiration):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintFeegrant(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SpendLimit) > 0 {
		for iNdEx := len(m.SpendLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpendLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFeegrant(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PeriodicAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PeriodicAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PeriodicAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.PeriodReset, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.PeriodReset):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintFeegrant(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	if len(m.PeriodCanSpend) > 0 {
		for iNdEx := len(m.PeriodCanSpend) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PeriodCanSpend[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFeegrant(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.PeriodSpendLimit) > 0 {
		for iNdEx := len(m.PeriodSpendLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PeriodSpendLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFeegrant(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	n3, err3 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.Period, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Period):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintFeegrant(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Basic.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFeegrant(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *AllowedMsgAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedMsgAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedMsgAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AllowedMessages) > 0 {
		for iNdEx := len(m.AllowedMessages) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedMessages[iNdEx])
			copy(dAtA[i:], m.AllowedMessages[iNdEx])
			i = encodeVarintFeegrant(dAtA, i, uint64(len(m.AllowedMessages[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Allowance != nil {
		{
			size, err := m.Allowance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFeegrant(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Grant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Grant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Grant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Allowance != nil {
		{
			size, err := m.Allowance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFeegrant(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintFeegrant(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintFeegrant(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFeegrant(dAtA []byte, offset int, v uint64) int {
	offset -= sovFeegrant(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BasicAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SpendLimit) > 0 {
		for _, e := range m.SpendLimit {
			l = e.Size()
			n += 1 + l + sovFeegrant(uint64(l))
		}
	}
	if m.Expiration != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Expiration)
		n += 1 + l + sovFeegrant(uint64(l))
	}
	return n
}

func (m *PeriodicAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Basic.Size()
	n += 1 + l + sovFeegrant(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Period)
	n += 1 + l + sovFeegrant(uint64(l))
	if len(m.PeriodSpendLimit) > 0 {
		for _, e := range m.PeriodSpendLimit {
			l = e.Size()
			n += 1 + l + sovFeegrant(uint64(l))
		}
	}
	if len(m.PeriodCanSpend) > 0 {
		for _, e := range m.PeriodCanSpend {
			l = e.Size()
			n += 1 + l + sovFeegrant(uint64(l))
		}
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.PeriodReset)
	n += 1 + l + sovFeegrant(uint64(l))
	return n
}

func (m *AllowedMsgAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Allowance != nil {
		l = m.Allowance.Size()
		n += 1 + l + sovFeegrant(uint64(l))
	}
	if len(m.AllowedMessages) > 0 {
		for _, s := range m.AllowedMessages {
			l = len(s)
			n += 1 + l + sovFeegrant(uint64(l))
		}
	}
	return n
}

func (m *Grant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovFeegrant(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovFeegrant(uint64(l))
	}
	if m.Allowance != nil {
		l = m.Allowance.Size()
		n += 1 + l + sovFeegrant(uint64(l))
	}
	return n
}

func sovFeegrant(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFeegrant(x uint64) (n int) {
	return sovFeegrant(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BasicAllowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeegrant
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BasicAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasicAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpendLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpendLimit = append(m.SpendLimit, types.Coin{})
			if err := m.SpendLimit[len(m.SpendLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expiration == nil {
				m.Expiration = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.Expiration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeegrant
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PeriodicAllowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeegrant
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PeriodicAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PeriodicAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Basic", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Basic.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.Period, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodSpendLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeriodSpendLimit = append(m.PeriodSpendLimit, types.Coin{})
			if err := m.PeriodSpendLimit[len(m.PeriodSpendLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodCanSpend", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeriodCanSpend = append(m.PeriodCanSpend, types.Coin{})
			if err := m.PeriodCanSpend[len(m.PeriodCanSpend)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeriodReset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.PeriodReset, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeegrant
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AllowedMsgAllowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeegrant
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AllowedMsgAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedMsgAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Allowance == nil {
				m.Allowance = &types1.Any{}
			}
			if err := m.Allowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedMessages", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedMessages = append(m.AllowedMessages, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeegrant
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Grant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeegrant
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Grant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Grant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Allowance == nil {
				m.Allowance = &types1.Any{}
			}
			if err := m.Allowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeegrant
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFeegrant(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFeegrant
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFeegrant
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthFeegrant
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFeegrant
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFeegrant
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFeegrant        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFeegrant          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFeegrant = fmt.Errorf("proto: unexpected end of group")
)
