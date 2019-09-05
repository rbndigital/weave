// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: orm/codec.proto

package orm

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// MultiRef contains a list of references to pks
type MultiRef struct {
	Refs [][]byte `protobuf:"bytes,1,rep,name=refs,proto3" json:"refs,omitempty"`
}

func (m *MultiRef) Reset()         { *m = MultiRef{} }
func (m *MultiRef) String() string { return proto.CompactTextString(m) }
func (*MultiRef) ProtoMessage()    {}
func (*MultiRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aef1e59ada91b17, []int{0}
}
func (m *MultiRef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MultiRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MultiRef.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MultiRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiRef.Merge(m, src)
}
func (m *MultiRef) XXX_Size() int {
	return m.Size()
}
func (m *MultiRef) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiRef.DiscardUnknown(m)
}

var xxx_messageInfo_MultiRef proto.InternalMessageInfo

func (m *MultiRef) GetRefs() [][]byte {
	if m != nil {
		return m.Refs
	}
	return nil
}

// Counter could be used for sequence, but mainly just for test
type Counter struct {
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *Counter) Reset()         { *m = Counter{} }
func (m *Counter) String() string { return proto.CompactTextString(m) }
func (*Counter) ProtoMessage()    {}
func (*Counter) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aef1e59ada91b17, []int{1}
}
func (m *Counter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Counter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Counter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Counter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Counter.Merge(m, src)
}
func (m *Counter) XXX_Size() int {
	return m.Size()
}
func (m *Counter) XXX_DiscardUnknown() {
	xxx_messageInfo_Counter.DiscardUnknown(m)
}

var xxx_messageInfo_Counter proto.InternalMessageInfo

func (m *Counter) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// VersionedID is the combination of document ID and version number.
type VersionedIDRef struct {
	// Unique identifier
	ID []byte `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	// Document version, starting with 1.
	Version uint32 `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *VersionedIDRef) Reset()         { *m = VersionedIDRef{} }
func (m *VersionedIDRef) String() string { return proto.CompactTextString(m) }
func (*VersionedIDRef) ProtoMessage()    {}
func (*VersionedIDRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aef1e59ada91b17, []int{2}
}
func (m *VersionedIDRef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VersionedIDRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VersionedIDRef.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VersionedIDRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionedIDRef.Merge(m, src)
}
func (m *VersionedIDRef) XXX_Size() int {
	return m.Size()
}
func (m *VersionedIDRef) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionedIDRef.DiscardUnknown(m)
}

var xxx_messageInfo_VersionedIDRef proto.InternalMessageInfo

func (m *VersionedIDRef) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *VersionedIDRef) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

// CounterWithID could be used for sequence, but mainly just for test
type CounterWithID struct {
	ID    []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Count int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *CounterWithID) Reset()         { *m = CounterWithID{} }
func (m *CounterWithID) String() string { return proto.CompactTextString(m) }
func (*CounterWithID) ProtoMessage()    {}
func (*CounterWithID) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aef1e59ada91b17, []int{3}
}
func (m *CounterWithID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CounterWithID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CounterWithID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CounterWithID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CounterWithID.Merge(m, src)
}
func (m *CounterWithID) XXX_Size() int {
	return m.Size()
}
func (m *CounterWithID) XXX_DiscardUnknown() {
	xxx_messageInfo_CounterWithID.DiscardUnknown(m)
}

var xxx_messageInfo_CounterWithID proto.InternalMessageInfo

func (m *CounterWithID) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *CounterWithID) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*MultiRef)(nil), "orm.MultiRef")
	proto.RegisterType((*Counter)(nil), "orm.Counter")
	proto.RegisterType((*VersionedIDRef)(nil), "orm.VersionedIDRef")
	proto.RegisterType((*CounterWithID)(nil), "orm.CounterWithID")
}

func init() { proto.RegisterFile("orm/codec.proto", fileDescriptor_4aef1e59ada91b17) }

var fileDescriptor_4aef1e59ada91b17 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x2f, 0xca, 0xd5,
	0x4f, 0xce, 0x4f, 0x49, 0x4d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2f, 0xca,
	0x95, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xf3, 0xf5, 0x41, 0x2c, 0x88, 0x94, 0x92, 0x1c, 0x17,
	0x87, 0x6f, 0x69, 0x4e, 0x49, 0x66, 0x50, 0x6a, 0x9a, 0x90, 0x10, 0x17, 0x4b, 0x51, 0x6a, 0x5a,
	0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0x4f, 0x10, 0x98, 0xad, 0x24, 0xcf, 0xc5, 0xee, 0x9c, 0x5f,
	0x9a, 0x57, 0x92, 0x5a, 0x24, 0x24, 0xc2, 0xc5, 0x9a, 0x0c, 0x62, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x30, 0x07, 0x41, 0x38, 0x4a, 0x4e, 0x5c, 0x7c, 0x61, 0xa9, 0x45, 0xc5, 0x99, 0xf9, 0x79, 0xa9,
	0x29, 0x9e, 0x2e, 0x20, 0x63, 0xc4, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x58, 0x14, 0x18, 0x35, 0x78,
	0x9c, 0xd8, 0x1e, 0xdd, 0x93, 0x67, 0xf2, 0x74, 0x09, 0x62, 0xca, 0x4c, 0x11, 0x92, 0xe0, 0x62,
	0x2f, 0x83, 0xa8, 0x94, 0x60, 0x55, 0x60, 0xd4, 0xe0, 0x0d, 0x82, 0x71, 0x95, 0x6c, 0xb9, 0x78,
	0xa1, 0x96, 0x84, 0x67, 0x96, 0x64, 0x78, 0xba, 0x40, 0x8d, 0x60, 0xc4, 0x30, 0x02, 0xee, 0x04,
	0x26, 0x24, 0x27, 0x38, 0x49, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47,
	0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x12,
	0x1b, 0xd8, 0x93, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0x95, 0xd0, 0xbd, 0x12, 0x01,
	0x00, 0x00,
}

func (m *MultiRef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MultiRef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Refs) > 0 {
		for _, b := range m.Refs {
			dAtA[i] = 0xa
			i++
			i = encodeVarintCodec(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	return i, nil
}

func (m *Counter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Counter) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Count))
	}
	return i, nil
}

func (m *VersionedIDRef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionedIDRef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.Version != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Version))
	}
	return i, nil
}

func (m *CounterWithID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CounterWithID) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.Count != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Count))
	}
	return i, nil
}

func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MultiRef) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Refs) > 0 {
		for _, b := range m.Refs {
			l = len(b)
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	return n
}

func (m *Counter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Count != 0 {
		n += 1 + sovCodec(uint64(m.Count))
	}
	return n
}

func (m *VersionedIDRef) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovCodec(uint64(m.Version))
	}
	return n
}

func (m *CounterWithID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovCodec(uint64(m.Count))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MultiRef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: MultiRef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MultiRef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Refs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Refs = append(m.Refs, make([]byte, postIndex-iNdEx))
			copy(m.Refs[len(m.Refs)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func (m *Counter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: Counter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Counter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func (m *VersionedIDRef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: VersionedIDRef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionedIDRef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = append(m.ID[:0], dAtA[iNdEx:postIndex]...)
			if m.ID == nil {
				m.ID = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func (m *CounterWithID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: CounterWithID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CounterWithID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = append(m.ID[:0], dAtA[iNdEx:postIndex]...)
			if m.ID == nil {
				m.ID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCodec
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
				return 0, ErrInvalidLengthCodec
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCodec
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCodec
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCodec(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCodec
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCodec = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec   = fmt.Errorf("proto: integer overflow")
)
