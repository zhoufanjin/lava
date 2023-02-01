// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pairing/provider_payment_storage.proto

package types

import (
	fmt "fmt"
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

type ProviderPaymentStorage struct {
	Index                                  string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Epoch                                  uint64   `protobuf:"varint,3,opt,name=epoch,proto3" json:"epoch,omitempty"`
	UnresponsivenessComplaints             []string `protobuf:"bytes,4,rep,name=unresponsiveness_complaints,json=unresponsivenessComplaints,proto3" json:"unresponsiveness_complaints,omitempty"`
	UniquePaymentStorageClientProviderKeys []string `protobuf:"bytes,5,rep,name=uniquePaymentStorageClientProviderKeys,proto3" json:"uniquePaymentStorageClientProviderKeys,omitempty"`
}

func (m *ProviderPaymentStorage) Reset()         { *m = ProviderPaymentStorage{} }
func (m *ProviderPaymentStorage) String() string { return proto.CompactTextString(m) }
func (*ProviderPaymentStorage) ProtoMessage()    {}
func (*ProviderPaymentStorage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f1d2e8d774659ae, []int{0}
}
func (m *ProviderPaymentStorage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProviderPaymentStorage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProviderPaymentStorage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProviderPaymentStorage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderPaymentStorage.Merge(m, src)
}
func (m *ProviderPaymentStorage) XXX_Size() int {
	return m.Size()
}
func (m *ProviderPaymentStorage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderPaymentStorage.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderPaymentStorage proto.InternalMessageInfo

func (m *ProviderPaymentStorage) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ProviderPaymentStorage) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *ProviderPaymentStorage) GetUnresponsivenessComplaints() []string {
	if m != nil {
		return m.UnresponsivenessComplaints
	}
	return nil
}

func (m *ProviderPaymentStorage) GetUniquePaymentStorageClientProviderKeys() []string {
	if m != nil {
		return m.UniquePaymentStorageClientProviderKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*ProviderPaymentStorage)(nil), "lavanet.lava.pairing.ProviderPaymentStorage")
}

func init() {
	proto.RegisterFile("pairing/provider_payment_storage.proto", fileDescriptor_4f1d2e8d774659ae)
}

var fileDescriptor_4f1d2e8d774659ae = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4a, 0xf4, 0x40,
	0x14, 0x85, 0x33, 0x7f, 0xb2, 0x3f, 0x6e, 0x2a, 0x09, 0x8b, 0x84, 0x15, 0x86, 0x60, 0xb1, 0xa6,
	0x4a, 0x0a, 0xed, 0x45, 0xb7, 0xd3, 0x66, 0x89, 0x60, 0x61, 0x13, 0x66, 0xb3, 0x97, 0xec, 0x40,
	0x32, 0x33, 0xce, 0x4c, 0xc2, 0xe6, 0x2d, 0x7c, 0x2c, 0xcb, 0x2d, 0x2d, 0x25, 0x79, 0x11, 0x49,
	0x26, 0x11, 0xdc, 0xca, 0xea, 0x72, 0x2e, 0xe7, 0x3b, 0x1c, 0x8e, 0xbb, 0x12, 0x84, 0x4a, 0xca,
	0xf2, 0x58, 0x48, 0x5e, 0xd3, 0x1d, 0xc8, 0x54, 0x90, 0xa6, 0x04, 0xa6, 0x53, 0xa5, 0xb9, 0x24,
	0x39, 0x44, 0x42, 0x72, 0xcd, 0xbd, 0x45, 0x41, 0x6a, 0xc2, 0x40, 0x47, 0xfd, 0x8d, 0x46, 0x68,
	0x79, 0x3b, 0xd1, 0x15, 0xa3, 0x6f, 0x15, 0x9c, 0xb2, 0x69, 0x56, 0xd0, 0x5e, 0x4e, 0xd9, 0x26,
	0xeb, 0xaa, 0x43, 0xee, 0xc5, 0x66, 0x7c, 0x6d, 0x0c, 0xf1, 0x6c, 0x00, 0x6f, 0xe1, 0xce, 0x28,
	0xdb, 0xc1, 0xc1, 0x47, 0x01, 0x0a, 0xe7, 0x89, 0x11, 0xfd, 0x17, 0x04, 0xcf, 0xf6, 0xbe, 0x1d,
	0xa0, 0xd0, 0x49, 0x8c, 0xf0, 0xee, 0xdc, 0xcb, 0x8a, 0x49, 0x50, 0x82, 0x33, 0x45, 0x6b, 0x60,
	0xa0, 0x54, 0x9a, 0xf1, 0x52, 0x14, 0x84, 0x32, 0xad, 0x7c, 0x27, 0xb0, 0xc3, 0x79, 0xb2, 0x3c,
	0xb5, 0xac, 0x7f, 0x1c, 0xde, 0x8b, 0xbb, 0x32, 0xbd, 0x7f, 0x97, 0x58, 0x0f, 0xa5, 0xa7, 0x82,
	0x4f, 0xd0, 0x28, 0x7f, 0x36, 0x64, 0xfd, 0xd1, 0xfd, 0xe8, 0x9c, 0xfd, 0x3b, 0xb7, 0x1f, 0xee,
	0x3f, 0x5a, 0x8c, 0x8e, 0x2d, 0x46, 0x5f, 0x2d, 0x46, 0xef, 0x1d, 0xb6, 0x8e, 0x1d, 0xb6, 0x3e,
	0x3b, 0x6c, 0xbd, 0x5e, 0xe7, 0x54, 0xef, 0xab, 0x6d, 0x94, 0xf1, 0x32, 0x1e, 0x67, 0x1d, 0x6e,
	0x7c, 0x88, 0xa7, 0x3d, 0x75, 0x23, 0x40, 0x6d, 0xff, 0x0f, 0x7b, 0xdd, 0x7c, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xb5, 0x8b, 0x82, 0xb1, 0xa5, 0x01, 0x00, 0x00,
}

func (m *ProviderPaymentStorage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProviderPaymentStorage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProviderPaymentStorage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UniquePaymentStorageClientProviderKeys) > 0 {
		for iNdEx := len(m.UniquePaymentStorageClientProviderKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.UniquePaymentStorageClientProviderKeys[iNdEx])
			copy(dAtA[i:], m.UniquePaymentStorageClientProviderKeys[iNdEx])
			i = encodeVarintProviderPaymentStorage(dAtA, i, uint64(len(m.UniquePaymentStorageClientProviderKeys[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.UnresponsivenessComplaints) > 0 {
		for iNdEx := len(m.UnresponsivenessComplaints) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.UnresponsivenessComplaints[iNdEx])
			copy(dAtA[i:], m.UnresponsivenessComplaints[iNdEx])
			i = encodeVarintProviderPaymentStorage(dAtA, i, uint64(len(m.UnresponsivenessComplaints[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Epoch != 0 {
		i = encodeVarintProviderPaymentStorage(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintProviderPaymentStorage(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProviderPaymentStorage(dAtA []byte, offset int, v uint64) int {
	offset -= sovProviderPaymentStorage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProviderPaymentStorage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovProviderPaymentStorage(uint64(l))
	}
	if m.Epoch != 0 {
		n += 1 + sovProviderPaymentStorage(uint64(m.Epoch))
	}
	if len(m.UnresponsivenessComplaints) > 0 {
		for _, s := range m.UnresponsivenessComplaints {
			l = len(s)
			n += 1 + l + sovProviderPaymentStorage(uint64(l))
		}
	}
	if len(m.UniquePaymentStorageClientProviderKeys) > 0 {
		for _, s := range m.UniquePaymentStorageClientProviderKeys {
			l = len(s)
			n += 1 + l + sovProviderPaymentStorage(uint64(l))
		}
	}
	return n
}

func sovProviderPaymentStorage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProviderPaymentStorage(x uint64) (n int) {
	return sovProviderPaymentStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProviderPaymentStorage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProviderPaymentStorage
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
			return fmt.Errorf("proto: ProviderPaymentStorage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProviderPaymentStorage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderPaymentStorage
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
				return ErrInvalidLengthProviderPaymentStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProviderPaymentStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderPaymentStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnresponsivenessComplaints", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderPaymentStorage
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
				return ErrInvalidLengthProviderPaymentStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProviderPaymentStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnresponsivenessComplaints = append(m.UnresponsivenessComplaints, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UniquePaymentStorageClientProviderKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderPaymentStorage
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
				return ErrInvalidLengthProviderPaymentStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProviderPaymentStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UniquePaymentStorageClientProviderKeys = append(m.UniquePaymentStorageClientProviderKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProviderPaymentStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProviderPaymentStorage
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
func skipProviderPaymentStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProviderPaymentStorage
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
					return 0, ErrIntOverflowProviderPaymentStorage
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
					return 0, ErrIntOverflowProviderPaymentStorage
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
				return 0, ErrInvalidLengthProviderPaymentStorage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProviderPaymentStorage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProviderPaymentStorage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProviderPaymentStorage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProviderPaymentStorage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProviderPaymentStorage = fmt.Errorf("proto: unexpected end of group")
)
