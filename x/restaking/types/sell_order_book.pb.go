// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/restaking/v1beta1/sell_order_book.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
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

type SellOrderBook struct {
	Index       string     `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	AmountDenom string     `protobuf:"bytes,2,opt,name=amountDenom,proto3" json:"amountDenom,omitempty"`
	PriceDenom  string     `protobuf:"bytes,3,opt,name=priceDenom,proto3" json:"priceDenom,omitempty"`
	Book        *OrderBook `protobuf:"bytes,4,opt,name=book,proto3" json:"book,omitempty"`
}

func (m *SellOrderBook) Reset()         { *m = SellOrderBook{} }
func (m *SellOrderBook) String() string { return proto.CompactTextString(m) }
func (*SellOrderBook) ProtoMessage()    {}
func (*SellOrderBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e7f48073810b349, []int{0}
}
func (m *SellOrderBook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SellOrderBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SellOrderBook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SellOrderBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SellOrderBook.Merge(m, src)
}
func (m *SellOrderBook) XXX_Size() int {
	return m.Size()
}
func (m *SellOrderBook) XXX_DiscardUnknown() {
	xxx_messageInfo_SellOrderBook.DiscardUnknown(m)
}

var xxx_messageInfo_SellOrderBook proto.InternalMessageInfo

func (m *SellOrderBook) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *SellOrderBook) GetAmountDenom() string {
	if m != nil {
		return m.AmountDenom
	}
	return ""
}

func (m *SellOrderBook) GetPriceDenom() string {
	if m != nil {
		return m.PriceDenom
	}
	return ""
}

func (m *SellOrderBook) GetBook() *OrderBook {
	if m != nil {
		return m.Book
	}
	return nil
}

func init() {
	proto.RegisterType((*SellOrderBook)(nil), "cosmos.restaking.v1beta1.SellOrderBook")
}

func init() {
	proto.RegisterFile("cosmos/restaking/v1beta1/sell_order_book.proto", fileDescriptor_2e7f48073810b349)
}

var fileDescriptor_2e7f48073810b349 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x2f, 0x4a, 0x2d, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0xd7, 0x2f, 0x33, 0x4c,
	0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x4e, 0xcd, 0xc9, 0x89, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x8a,
	0x4f, 0xca, 0xcf, 0xcf, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x80, 0xa8, 0xd7, 0x83,
	0xab, 0xd7, 0x83, 0xaa, 0x97, 0x52, 0xc1, 0x69, 0x12, 0xd8, 0x10, 0x88, 0x7e, 0xa5, 0x05, 0x8c,
	0x5c, 0xbc, 0xc1, 0xa9, 0x39, 0x39, 0xfe, 0x20, 0x31, 0xa7, 0xfc, 0xfc, 0x6c, 0x21, 0x11, 0x2e,
	0xd6, 0xcc, 0xbc, 0x94, 0xd4, 0x0a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48,
	0x81, 0x8b, 0x3b, 0x31, 0x37, 0xbf, 0x34, 0xaf, 0xc4, 0x25, 0x35, 0x2f, 0x3f, 0x57, 0x82, 0x09,
	0x2c, 0x87, 0x2c, 0x24, 0x24, 0xc7, 0xc5, 0x55, 0x50, 0x94, 0x99, 0x9c, 0x0a, 0x51, 0xc0, 0x0c,
	0x56, 0x80, 0x24, 0x22, 0x64, 0xce, 0xc5, 0x02, 0x72, 0xb7, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0xb7,
	0x91, 0xb2, 0x1e, 0x2e, 0x87, 0xeb, 0xc1, 0x9d, 0x12, 0x04, 0xd6, 0xe0, 0xe4, 0x73, 0xe2, 0x91,
	0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1,
	0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x46, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a,
	0xc9, 0xf9, 0xb9, 0xfa, 0x39, 0x99, 0xe9, 0x19, 0x25, 0x20, 0xff, 0xc2, 0x18, 0xba, 0xc5, 0x29,
	0xd9, 0xfa, 0x15, 0x48, 0xde, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xfb, 0xdb, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x27, 0xfb, 0xf5, 0xb5, 0x69, 0x01, 0x00, 0x00,
}

func (m *SellOrderBook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SellOrderBook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SellOrderBook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Book != nil {
		{
			size, err := m.Book.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSellOrderBook(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.PriceDenom) > 0 {
		i -= len(m.PriceDenom)
		copy(dAtA[i:], m.PriceDenom)
		i = encodeVarintSellOrderBook(dAtA, i, uint64(len(m.PriceDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AmountDenom) > 0 {
		i -= len(m.AmountDenom)
		copy(dAtA[i:], m.AmountDenom)
		i = encodeVarintSellOrderBook(dAtA, i, uint64(len(m.AmountDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintSellOrderBook(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSellOrderBook(dAtA []byte, offset int, v uint64) int {
	offset -= sovSellOrderBook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SellOrderBook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovSellOrderBook(uint64(l))
	}
	l = len(m.AmountDenom)
	if l > 0 {
		n += 1 + l + sovSellOrderBook(uint64(l))
	}
	l = len(m.PriceDenom)
	if l > 0 {
		n += 1 + l + sovSellOrderBook(uint64(l))
	}
	if m.Book != nil {
		l = m.Book.Size()
		n += 1 + l + sovSellOrderBook(uint64(l))
	}
	return n
}

func sovSellOrderBook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSellOrderBook(x uint64) (n int) {
	return sovSellOrderBook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SellOrderBook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSellOrderBook
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
			return fmt.Errorf("proto: SellOrderBook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SellOrderBook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOrderBook
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
				return ErrInvalidLengthSellOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSellOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOrderBook
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
				return ErrInvalidLengthSellOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSellOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOrderBook
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
				return ErrInvalidLengthSellOrderBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSellOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Book", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSellOrderBook
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
				return ErrInvalidLengthSellOrderBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSellOrderBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Book == nil {
				m.Book = &OrderBook{}
			}
			if err := m.Book.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSellOrderBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSellOrderBook
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
func skipSellOrderBook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSellOrderBook
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
					return 0, ErrIntOverflowSellOrderBook
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
					return 0, ErrIntOverflowSellOrderBook
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
				return 0, ErrInvalidLengthSellOrderBook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSellOrderBook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSellOrderBook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSellOrderBook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSellOrderBook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSellOrderBook = fmt.Errorf("proto: unexpected end of group")
)
