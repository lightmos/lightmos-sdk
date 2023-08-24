// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lightmos/restaking/validator_token.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	types "github.com/lightmos/lightmos-sdk/types"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ValidatorToken struct {
	Address   string      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Total     *types.Coin `protobuf:"bytes,2,opt,name=total,proto3" json:"total,omitempty"`
	Retire    *types.Coin `protobuf:"bytes,3,opt,name=retire,proto3" json:"retire,omitempty"`
	Available *types.Coin `protobuf:"bytes,4,opt,name=available,proto3" json:"available,omitempty"`
}

func (m *ValidatorToken) Reset()         { *m = ValidatorToken{} }
func (m *ValidatorToken) String() string { return proto.CompactTextString(m) }
func (*ValidatorToken) ProtoMessage()    {}
func (*ValidatorToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_8512cd1427318303, []int{0}
}
func (m *ValidatorToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorToken.Merge(m, src)
}
func (m *ValidatorToken) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorToken) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorToken.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorToken proto.InternalMessageInfo

func (m *ValidatorToken) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ValidatorToken) GetTotal() *types.Coin {
	if m != nil {
		return m.Total
	}
	return nil
}

func (m *ValidatorToken) GetRetire() *types.Coin {
	if m != nil {
		return m.Retire
	}
	return nil
}

func (m *ValidatorToken) GetAvailable() *types.Coin {
	if m != nil {
		return m.Available
	}
	return nil
}

func init() {
	proto.RegisterType((*ValidatorToken)(nil), "lightmos.restaking.ValidatorToken")
}

func init() {
	proto.RegisterFile("lightmos/restaking/validator_token.proto", fileDescriptor_8512cd1427318303)
}

var fileDescriptor_8512cd1427318303 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc8, 0xc9, 0x4c, 0xcf,
	0x28, 0xc9, 0xcd, 0x2f, 0xd6, 0x2f, 0x4a, 0x2d, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0x2f,
	0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x8a, 0x2f, 0xc9, 0xcf, 0x4e, 0xcd, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x82, 0xa9, 0xd4, 0x83, 0xab, 0x94, 0x92, 0x4c, 0xce, 0x2f,
	0xce, 0xcd, 0x2f, 0x8e, 0x07, 0xab, 0xd0, 0x87, 0x70, 0x20, 0xca, 0xa5, 0x44, 0xd2, 0xf3, 0xd3,
	0xf3, 0x21, 0xe2, 0x20, 0x16, 0x54, 0x54, 0x02, 0x6e, 0x5d, 0x52, 0x62, 0x71, 0xaa, 0x7e, 0x72,
	0x7e, 0x26, 0xd4, 0x78, 0xa5, 0xad, 0x8c, 0x5c, 0x7c, 0x61, 0x30, 0x8b, 0x43, 0x40, 0xf6, 0x0a,
	0x49, 0x70, 0xb1, 0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0xc1, 0xb8, 0x42, 0x9a, 0x5c, 0xac, 0x25, 0xf9, 0x25, 0x89, 0x39, 0x12, 0x4c, 0x0a, 0x8c,
	0x1a, 0xdc, 0x46, 0xc2, 0x7a, 0x70, 0xb7, 0x81, 0x8c, 0xd5, 0x73, 0xce, 0xcf, 0xcc, 0x0b, 0x82,
	0xa8, 0x10, 0xd2, 0xe6, 0x62, 0x2b, 0x4a, 0x2d, 0xc9, 0x2c, 0x4a, 0x95, 0x60, 0xc6, 0xad, 0x16,
	0xaa, 0x44, 0xc8, 0x90, 0x8b, 0x33, 0xb1, 0x2c, 0x31, 0x33, 0x27, 0x31, 0x29, 0x27, 0x55, 0x82,
	0x05, 0xb7, 0x7a, 0x84, 0x2a, 0x27, 0x93, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63,
	0x88, 0x92, 0x82, 0xfb, 0xb5, 0x02, 0x29, 0x70, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0,
	0x9e, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x71, 0x3e, 0x0e, 0x7f, 0x01, 0x00, 0x00,
}

func (m *ValidatorToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Available != nil {
		{
			size, err := m.Available.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintValidatorToken(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Retire != nil {
		{
			size, err := m.Retire.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintValidatorToken(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Total != nil {
		{
			size, err := m.Total.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintValidatorToken(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintValidatorToken(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidatorToken(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidatorToken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ValidatorToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovValidatorToken(uint64(l))
	}
	if m.Total != nil {
		l = m.Total.Size()
		n += 1 + l + sovValidatorToken(uint64(l))
	}
	if m.Retire != nil {
		l = m.Retire.Size()
		n += 1 + l + sovValidatorToken(uint64(l))
	}
	if m.Available != nil {
		l = m.Available.Size()
		n += 1 + l + sovValidatorToken(uint64(l))
	}
	return n
}

func sovValidatorToken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidatorToken(x uint64) (n int) {
	return sovValidatorToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidatorToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidatorToken
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
			return fmt.Errorf("proto: ValidatorToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorToken
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
				return ErrInvalidLengthValidatorToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidatorToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorToken
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
				return ErrInvalidLengthValidatorToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidatorToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Total == nil {
				m.Total = &types.Coin{}
			}
			if err := m.Total.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Retire", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorToken
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
				return ErrInvalidLengthValidatorToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidatorToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Retire == nil {
				m.Retire = &types.Coin{}
			}
			if err := m.Retire.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Available", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidatorToken
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
				return ErrInvalidLengthValidatorToken
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidatorToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Available == nil {
				m.Available = &types.Coin{}
			}
			if err := m.Available.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidatorToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidatorToken
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
func skipValidatorToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidatorToken
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
					return 0, ErrIntOverflowValidatorToken
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
					return 0, ErrIntOverflowValidatorToken
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
				return 0, ErrInvalidLengthValidatorToken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidatorToken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidatorToken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidatorToken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidatorToken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidatorToken = fmt.Errorf("proto: unexpected end of group")
)
