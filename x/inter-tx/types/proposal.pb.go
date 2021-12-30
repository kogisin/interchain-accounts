// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: intertx/proposal.proto

package types

import (
	fmt "fmt"
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

type RegisterInterchainAccountProposal struct {
	Title                    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description              string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ConnectionId             string `protobuf:"bytes,3,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty" yaml:"connection_id"`
	CounterpartyConnectionId string `protobuf:"bytes,4,opt,name=counterparty_connection_id,json=counterpartyConnectionId,proto3" json:"counterparty_connection_id,omitempty" yaml:"counterparty_connection_id"`
}

func (m *RegisterInterchainAccountProposal) Reset()      { *m = RegisterInterchainAccountProposal{} }
func (*RegisterInterchainAccountProposal) ProtoMessage() {}
func (*RegisterInterchainAccountProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_277624c2d82e037d, []int{0}
}
func (m *RegisterInterchainAccountProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterInterchainAccountProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterInterchainAccountProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterInterchainAccountProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterInterchainAccountProposal.Merge(m, src)
}
func (m *RegisterInterchainAccountProposal) XXX_Size() int {
	return m.Size()
}
func (m *RegisterInterchainAccountProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterInterchainAccountProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterInterchainAccountProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterInterchainAccountProposal)(nil), "intertx.RegisterInterchainAccountProposal")
}

func init() { proto.RegisterFile("intertx/proposal.proto", fileDescriptor_277624c2d82e037d) }

var fileDescriptor_277624c2d82e037d = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0xcc, 0x2b, 0x49,
	0x2d, 0x2a, 0xa9, 0xd0, 0x2f, 0x28, 0xca, 0x2f, 0xc8, 0x2f, 0x4e, 0xcc, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x87, 0x8a, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0xc5, 0xf4, 0x41,
	0x2c, 0x88, 0xb4, 0x52, 0x27, 0x13, 0x97, 0x62, 0x50, 0x6a, 0x7a, 0x66, 0x71, 0x49, 0x6a, 0x91,
	0x27, 0x48, 0x65, 0x72, 0x46, 0x62, 0x66, 0x9e, 0x63, 0x72, 0x72, 0x7e, 0x69, 0x5e, 0x49, 0x00,
	0xd4, 0x28, 0x21, 0x11, 0x2e, 0xd6, 0x92, 0xcc, 0x92, 0x9c, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x08, 0x47, 0x48, 0x81, 0x8b, 0x3b, 0x25, 0xb5, 0x38, 0xb9, 0x28, 0xb3, 0xa0, 0x24,
	0x33, 0x3f, 0x4f, 0x82, 0x09, 0x2c, 0x87, 0x2c, 0x24, 0x64, 0xcb, 0xc5, 0x9b, 0x9c, 0x9f, 0x97,
	0x97, 0x9a, 0x0c, 0xe2, 0xc5, 0x67, 0xa6, 0x48, 0x30, 0x83, 0xd4, 0x38, 0x49, 0x7c, 0xba, 0x27,
	0x2f, 0x52, 0x99, 0x98, 0x9b, 0x63, 0xa5, 0x84, 0x22, 0xad, 0x14, 0xc4, 0x83, 0xe0, 0x7b, 0xa6,
	0x08, 0x25, 0x73, 0x49, 0x81, 0xdd, 0x91, 0x5a, 0x54, 0x90, 0x58, 0x54, 0x52, 0x19, 0x8f, 0x6a,
	0x16, 0x0b, 0xd8, 0x2c, 0xd5, 0x4f, 0xf7, 0xe4, 0x15, 0x61, 0x66, 0xe1, 0x52, 0xab, 0x14, 0x24,
	0x81, 0x2c, 0xe9, 0x8c, 0x64, 0x89, 0x15, 0x47, 0xc7, 0x02, 0x79, 0x86, 0x19, 0x0b, 0xe4, 0x19,
	0x9c, 0x02, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09,
	0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x2c, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x39, 0xbf, 0x38, 0x37, 0xbf, 0x58, 0x3f,
	0x13, 0x1e, 0x58, 0xba, 0x89, 0x90, 0xd0, 0x2a, 0xd6, 0xaf, 0x80, 0x88, 0xea, 0x96, 0x54, 0xe8,
	0x97, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x03, 0xd9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xe5, 0x62, 0x7c, 0xff, 0x9d, 0x01, 0x00, 0x00,
}

func (m *RegisterInterchainAccountProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterInterchainAccountProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterInterchainAccountProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CounterpartyConnectionId) > 0 {
		i -= len(m.CounterpartyConnectionId)
		copy(dAtA[i:], m.CounterpartyConnectionId)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.CounterpartyConnectionId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegisterInterchainAccountProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.CounterpartyConnectionId)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisterInterchainAccountProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: RegisterInterchainAccountProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterInterchainAccountProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CounterpartyConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CounterpartyConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
