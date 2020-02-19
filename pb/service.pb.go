// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

package pb

import (
	fmt "fmt"
	_ "github.com/bold-commerce/protoc-gen-struct-transformer/options"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type Inner struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *Inner) Reset()         { *m = Inner{} }
func (m *Inner) String() string { return proto.CompactTextString(m) }
func (*Inner) ProtoMessage()    {}
func (*Inner) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}
func (m *Inner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Inner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Inner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Inner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inner.Merge(m, src)
}
func (m *Inner) XXX_Size() int {
	return m.Size()
}
func (m *Inner) XXX_DiscardUnknown() {
	xxx_messageInfo_Inner.DiscardUnknown(m)
}

var xxx_messageInfo_Inner proto.InternalMessageInfo

func (m *Inner) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Outer struct {
	Id     int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Inners []Inner `protobuf:"bytes,2,rep,name=inners,proto3" json:"inners"`
}

func (m *Outer) Reset()         { *m = Outer{} }
func (m *Outer) String() string { return proto.CompactTextString(m) }
func (*Outer) ProtoMessage()    {}
func (*Outer) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}
func (m *Outer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Outer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Outer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Outer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Outer.Merge(m, src)
}
func (m *Outer) XXX_Size() int {
	return m.Size()
}
func (m *Outer) XXX_DiscardUnknown() {
	xxx_messageInfo_Outer.DiscardUnknown(m)
}

var xxx_messageInfo_Outer proto.InternalMessageInfo

func (m *Outer) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Outer) GetInners() []Inner {
	if m != nil {
		return m.Inners
	}
	return nil
}

func init() {
	proto.RegisterType((*Inner)(nil), "mypackage.Inner")
	proto.RegisterType((*Outer)(nil), "mypackage.Outer")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0x3f, 0x4b, 0xc4, 0x40,
	0x10, 0xc5, 0x93, 0xbd, 0x3f, 0x98, 0x15, 0x45, 0x52, 0x1d, 0x07, 0xae, 0xc7, 0x15, 0x72, 0x4d,
	0x12, 0xd0, 0x42, 0xd0, 0xee, 0x3a, 0x2b, 0x21, 0xa5, 0xdd, 0x26, 0x99, 0x0b, 0x41, 0xb3, 0xbb,
	0xec, 0x4e, 0x44, 0xbf, 0x82, 0x95, 0x58, 0xfa, 0x61, 0xec, 0x02, 0x57, 0x5e, 0x69, 0x25, 0x92,
	0x7c, 0x11, 0x71, 0x22, 0x57, 0xcd, 0x9b, 0x79, 0x6f, 0x7e, 0x3c, 0x7e, 0xe4, 0xc0, 0x3e, 0x55,
	0x39, 0xc4, 0xc6, 0x6a, 0xd4, 0x61, 0x50, 0xbf, 0x18, 0x99, 0x3f, 0xc8, 0x12, 0xe6, 0x57, 0x74,
	0xc9, 0xa3, 0x12, 0x54, 0xe4, 0xd0, 0x36, 0x39, 0x46, 0x68, 0xa5, 0x72, 0x1b, 0x6d, 0x6b, 0xb0,
	0x89, 0x36, 0x58, 0x69, 0xe5, 0x12, 0xa9, 0x94, 0x46, 0x49, 0x7a, 0x60, 0xcc, 0x4f, 0x69, 0x64,
	0xcd, 0x26, 0x29, 0x75, 0xa9, 0x69, 0x21, 0x35, 0xd8, 0xcb, 0x73, 0x3e, 0xb9, 0x55, 0x0a, 0x6c,
	0x18, 0xf2, 0x31, 0xc2, 0x33, 0xce, 0xfc, 0x85, 0xbf, 0x0a, 0x52, 0xd2, 0xd7, 0x41, 0xf7, 0xc9,
	0x06, 0x7b, 0x99, 0xf2, 0xc9, 0x5d, 0x83, 0x60, 0xc3, 0x63, 0xce, 0xaa, 0x82, 0x52, 0xa3, 0x94,
	0x55, 0x45, 0x18, 0xf3, 0x69, 0xf5, 0x97, 0x70, 0x33, 0xb6, 0x18, 0xad, 0x0e, 0x2f, 0x4e, 0xe2,
	0x7d, 0xe9, 0x98, 0x5e, 0xd7, 0xe3, 0xed, 0xf7, 0x99, 0x97, 0xfe, 0xa7, 0x06, 0x26, 0xa1, 0xd6,
	0x37, 0xaf, 0x2d, 0x0b, 0x6a, 0x5d, 0xc0, 0xa3, 0x8b, 0x4b, 0xfd, 0xde, 0xb2, 0x03, 0x0b, 0x06,
	0x24, 0x42, 0xf1, 0xd1, 0x32, 0x66, 0xb2, 0x6d, 0x27, 0xfc, 0x5d, 0x27, 0xfc, 0x9f, 0x4e, 0xf8,
	0x6f, 0xbd, 0xf0, 0x76, 0xbd, 0xf0, 0xbe, 0x7a, 0xe1, 0xdd, 0x33, 0x93, 0x65, 0x53, 0xea, 0x7f,
	0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x2c, 0x5d, 0x16, 0x33, 0x01, 0x00, 0x00,
}

func (m *Inner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Inner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Inner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintService(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Outer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Outer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Outer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Inners) > 0 {
		for iNdEx := len(m.Inners) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Inners[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Id != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Inner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *Outer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovService(uint64(m.Id))
	}
	if len(m.Inners) > 0 {
		for _, e := range m.Inners {
			l = e.Size()
			n += 1 + l + sovService(uint64(l))
		}
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Inner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: Inner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Inner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *Outer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: Outer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Outer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inners", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Inners = append(m.Inners, Inner{})
			if err := m.Inners[len(m.Inners)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
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
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
