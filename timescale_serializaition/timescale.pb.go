// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: timescale.proto

/*
	Package timescale_serialization is a generated protocol buffer package.

	It is generated from these files:
		timescale.proto

	It has these top-level messages:
		FlatPoint
*/
package timescale_serialization

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FlatPoint_ValueType int32

const (
	FlatPoint_INTEGER FlatPoint_ValueType = 0
	FlatPoint_FLOAT   FlatPoint_ValueType = 1
	FlatPoint_STRING  FlatPoint_ValueType = 2
)

var FlatPoint_ValueType_name = map[int32]string{
	0: "INTEGER",
	1: "FLOAT",
	2: "STRING",
}
var FlatPoint_ValueType_value = map[string]int32{
	"INTEGER": 0,
	"FLOAT":   1,
	"STRING":  2,
}

func (x FlatPoint_ValueType) String() string {
	return proto.EnumName(FlatPoint_ValueType_name, int32(x))
}
func (FlatPoint_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorTimescale, []int{0, 0}
}

type FlatPoint struct {
	MeasurementName string                      `protobuf:"bytes,1,opt,name=measurementName,proto3" json:"measurementName,omitempty"`
	Columns         []string                    `protobuf:"bytes,2,rep,name=columns" json:"columns,omitempty"`
	Values          []*FlatPoint_FlatPointValue `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
}

func (m *FlatPoint) Reset()                    { *m = FlatPoint{} }
func (m *FlatPoint) String() string            { return proto.CompactTextString(m) }
func (*FlatPoint) ProtoMessage()               {}
func (*FlatPoint) Descriptor() ([]byte, []int) { return fileDescriptorTimescale, []int{0} }

func (m *FlatPoint) GetMeasurementName() string {
	if m != nil {
		return m.MeasurementName
	}
	return ""
}

func (m *FlatPoint) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *FlatPoint) GetValues() []*FlatPoint_FlatPointValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type FlatPoint_FlatPointValue struct {
	Type      FlatPoint_ValueType `protobuf:"varint,1,opt,name=type,proto3,enum=timescale_serialization.FlatPoint_ValueType" json:"type,omitempty"`
	IntVal    int64               `protobuf:"varint,2,opt,name=intVal,proto3" json:"intVal,omitempty"`
	DoubleVal float64             `protobuf:"fixed64,3,opt,name=doubleVal,proto3" json:"doubleVal,omitempty"`
	StringVal string              `protobuf:"bytes,4,opt,name=stringVal,proto3" json:"stringVal,omitempty"`
}

func (m *FlatPoint_FlatPointValue) Reset()         { *m = FlatPoint_FlatPointValue{} }
func (m *FlatPoint_FlatPointValue) String() string { return proto.CompactTextString(m) }
func (*FlatPoint_FlatPointValue) ProtoMessage()    {}
func (*FlatPoint_FlatPointValue) Descriptor() ([]byte, []int) {
	return fileDescriptorTimescale, []int{0, 0}
}

func (m *FlatPoint_FlatPointValue) GetType() FlatPoint_ValueType {
	if m != nil {
		return m.Type
	}
	return FlatPoint_INTEGER
}

func (m *FlatPoint_FlatPointValue) GetIntVal() int64 {
	if m != nil {
		return m.IntVal
	}
	return 0
}

func (m *FlatPoint_FlatPointValue) GetDoubleVal() float64 {
	if m != nil {
		return m.DoubleVal
	}
	return 0
}

func (m *FlatPoint_FlatPointValue) GetStringVal() string {
	if m != nil {
		return m.StringVal
	}
	return ""
}

func init() {
	proto.RegisterType((*FlatPoint)(nil), "timescale_serialization.FlatPoint")
	proto.RegisterType((*FlatPoint_FlatPointValue)(nil), "timescale_serialization.FlatPoint.FlatPointValue")
	proto.RegisterEnum("timescale_serialization.FlatPoint_ValueType", FlatPoint_ValueType_name, FlatPoint_ValueType_value)
}
func (m *FlatPoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FlatPoint) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.MeasurementName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTimescale(dAtA, i, uint64(len(m.MeasurementName)))
		i += copy(dAtA[i:], m.MeasurementName)
	}
	if len(m.Columns) > 0 {
		for _, s := range m.Columns {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Values) > 0 {
		for _, msg := range m.Values {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintTimescale(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *FlatPoint_FlatPointValue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FlatPoint_FlatPointValue) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTimescale(dAtA, i, uint64(m.Type))
	}
	if m.IntVal != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTimescale(dAtA, i, uint64(m.IntVal))
	}
	if m.DoubleVal != 0 {
		dAtA[i] = 0x19
		i++
		i = encodeFixed64Timescale(dAtA, i, uint64(math.Float64bits(float64(m.DoubleVal))))
	}
	if len(m.StringVal) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTimescale(dAtA, i, uint64(len(m.StringVal)))
		i += copy(dAtA[i:], m.StringVal)
	}
	return i, nil
}

func encodeFixed64Timescale(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Timescale(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTimescale(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FlatPoint) Size() (n int) {
	var l int
	_ = l
	l = len(m.MeasurementName)
	if l > 0 {
		n += 1 + l + sovTimescale(uint64(l))
	}
	if len(m.Columns) > 0 {
		for _, s := range m.Columns {
			l = len(s)
			n += 1 + l + sovTimescale(uint64(l))
		}
	}
	if len(m.Values) > 0 {
		for _, e := range m.Values {
			l = e.Size()
			n += 1 + l + sovTimescale(uint64(l))
		}
	}
	return n
}

func (m *FlatPoint_FlatPointValue) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovTimescale(uint64(m.Type))
	}
	if m.IntVal != 0 {
		n += 1 + sovTimescale(uint64(m.IntVal))
	}
	if m.DoubleVal != 0 {
		n += 9
	}
	l = len(m.StringVal)
	if l > 0 {
		n += 1 + l + sovTimescale(uint64(l))
	}
	return n
}

func sovTimescale(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTimescale(x uint64) (n int) {
	return sovTimescale(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FlatPoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimescale
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FlatPoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FlatPoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MeasurementName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimescale
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTimescale
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MeasurementName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimescale
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTimescale
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimescale
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTimescale
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, &FlatPoint_FlatPointValue{})
			if err := m.Values[len(m.Values)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimescale(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimescale
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
func (m *FlatPoint_FlatPointValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimescale
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FlatPointValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FlatPointValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimescale
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (FlatPoint_ValueType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntVal", wireType)
			}
			m.IntVal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimescale
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IntVal |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field DoubleVal", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.DoubleVal = float64(math.Float64frombits(v))
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringVal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimescale
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTimescale
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StringVal = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimescale(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimescale
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
func skipTimescale(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimescale
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
					return 0, ErrIntOverflowTimescale
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
					return 0, ErrIntOverflowTimescale
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTimescale
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTimescale
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
				next, err := skipTimescale(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthTimescale = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimescale   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("timescale.proto", fileDescriptorTimescale) }

var fileDescriptorTimescale = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xc9, 0xcc, 0x4d,
	0x2d, 0x4e, 0x4e, 0xcc, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x87, 0x0b, 0xc4,
	0x17, 0xa7, 0x16, 0x65, 0x26, 0xe6, 0x64, 0x56, 0x25, 0x96, 0x64, 0xe6, 0xe7, 0x29, 0x7d, 0x61,
	0xe2, 0xe2, 0x74, 0xcb, 0x49, 0x2c, 0x09, 0xc8, 0xcf, 0xcc, 0x2b, 0x11, 0xd2, 0xe0, 0xe2, 0xcf,
	0x4d, 0x4d, 0x2c, 0x2e, 0x2d, 0x4a, 0xcd, 0x4d, 0xcd, 0x2b, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x17, 0x16, 0x92, 0xe0, 0x62, 0x4f, 0xce, 0xcf, 0x29, 0xcd,
	0xcd, 0x2b, 0x96, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0x3c, 0xb9, 0xd8, 0xca,
	0x12, 0x73, 0x4a, 0x53, 0x8b, 0x25, 0x98, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x0c, 0xf5, 0x70, 0xd8,
	0xad, 0x07, 0xb7, 0x17, 0xc1, 0x0a, 0x03, 0xe9, 0x0c, 0x82, 0x1a, 0x20, 0xb5, 0x8c, 0x91, 0x8b,
	0x0f, 0x55, 0x4a, 0xc8, 0x81, 0x8b, 0xa5, 0xa4, 0xb2, 0x00, 0xe2, 0x2c, 0x3e, 0x23, 0x1d, 0x22,
	0xcc, 0x06, 0xeb, 0x0b, 0xa9, 0x2c, 0x48, 0x0d, 0x02, 0xeb, 0x14, 0x12, 0xe3, 0x62, 0x83, 0x98,
	0x26, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1c, 0x04, 0xe5, 0x09, 0xc9, 0x70, 0x71, 0xa6, 0xe4, 0x97,
	0x26, 0xe5, 0xa4, 0x82, 0xa4, 0x98, 0x15, 0x18, 0x35, 0x18, 0x83, 0x10, 0x02, 0x20, 0xd9, 0xe2,
	0x92, 0xa2, 0xcc, 0xbc, 0x74, 0x90, 0x2c, 0x0b, 0x38, 0x4c, 0x10, 0x02, 0x4a, 0xfa, 0x5c, 0x9c,
	0x70, 0x6b, 0x84, 0xb8, 0xb9, 0xd8, 0x3d, 0xfd, 0x42, 0x5c, 0xdd, 0x5d, 0x83, 0x04, 0x18, 0x84,
	0x38, 0xb9, 0x58, 0xdd, 0x7c, 0xfc, 0x1d, 0x43, 0x04, 0x18, 0x85, 0xb8, 0xb8, 0xd8, 0x82, 0x43,
	0x82, 0x3c, 0xfd, 0xdc, 0x05, 0x98, 0x9c, 0x04, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e,
	0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x11, 0x65, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0xeb, 0xe5, 0x54, 0xce, 0xbb, 0x01, 0x00, 0x00,
}
