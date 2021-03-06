// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dependency.proto

/*
	Package thunderpb is a generated protocol buffer package.

	It is generated from these files:
		dependency.proto
		federation.proto
		testcustomexecutor.proto

	It has these top-level messages:
		Field
		SQLFilter
		ExpirationTime
		Selection
		Fragment
		SelectionSet
		Query
		ExecuteRequest
		ExecuteResponse
		CustomExecutorRequest
		CustomExecutorResponse
*/
package thunderpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import time "time"
import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import binary "encoding/binary"
import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FieldKind int32

const (
	FieldKind_Unknown FieldKind = 0
	FieldKind_Null    FieldKind = 1
	FieldKind_Bool    FieldKind = 2
	FieldKind_Int     FieldKind = 3
	FieldKind_Uint    FieldKind = 4
	FieldKind_String  FieldKind = 5
	FieldKind_Bytes   FieldKind = 6
	FieldKind_Float64 FieldKind = 7
	FieldKind_Time    FieldKind = 8
)

var FieldKind_name = map[int32]string{
	0: "Unknown",
	1: "Null",
	2: "Bool",
	3: "Int",
	4: "Uint",
	5: "String",
	6: "Bytes",
	7: "Float64",
	8: "Time",
}
var FieldKind_value = map[string]int32{
	"Unknown": 0,
	"Null":    1,
	"Bool":    2,
	"Int":     3,
	"Uint":    4,
	"String":  5,
	"Bytes":   6,
	"Float64": 7,
	"Time":    8,
}

func (x FieldKind) String() string {
	return proto.EnumName(FieldKind_name, int32(x))
}
func (FieldKind) EnumDescriptor() ([]byte, []int) { return fileDescriptorDependency, []int{0} }

type Field struct {
	Kind FieldKind `protobuf:"varint,1,opt,name=kind,proto3,enum=thunderpb.FieldKind" json:"kind,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Field_Bool
	//	*Field_Int
	//	*Field_Uint
	//	*Field_String_
	//	*Field_Bytes
	//	*Field_Float64
	//	*Field_Time
	Value isField_Value `protobuf_oneof:"value"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptorDependency, []int{0} }

type isField_Value interface {
	isField_Value()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Field_Bool struct {
	Bool bool `protobuf:"varint,2,opt,name=bool,proto3,oneof"`
}
type Field_Int struct {
	Int int64 `protobuf:"zigzag64,3,opt,name=int,proto3,oneof"`
}
type Field_Uint struct {
	Uint uint64 `protobuf:"varint,4,opt,name=uint,proto3,oneof"`
}
type Field_String_ struct {
	String_ string `protobuf:"bytes,5,opt,name=string,proto3,oneof"`
}
type Field_Bytes struct {
	Bytes []byte `protobuf:"bytes,6,opt,name=bytes,proto3,oneof"`
}
type Field_Float64 struct {
	Float64 float64 `protobuf:"fixed64,7,opt,name=float64,proto3,oneof"`
}
type Field_Time struct {
	Time *time.Time `protobuf:"bytes,8,opt,name=time,oneof,stdtime"`
}

func (*Field_Bool) isField_Value()    {}
func (*Field_Int) isField_Value()     {}
func (*Field_Uint) isField_Value()    {}
func (*Field_String_) isField_Value() {}
func (*Field_Bytes) isField_Value()   {}
func (*Field_Float64) isField_Value() {}
func (*Field_Time) isField_Value()    {}

func (m *Field) GetValue() isField_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Field) GetKind() FieldKind {
	if m != nil {
		return m.Kind
	}
	return FieldKind_Unknown
}

func (m *Field) GetBool() bool {
	if x, ok := m.GetValue().(*Field_Bool); ok {
		return x.Bool
	}
	return false
}

func (m *Field) GetInt() int64 {
	if x, ok := m.GetValue().(*Field_Int); ok {
		return x.Int
	}
	return 0
}

func (m *Field) GetUint() uint64 {
	if x, ok := m.GetValue().(*Field_Uint); ok {
		return x.Uint
	}
	return 0
}

func (m *Field) GetString_() string {
	if x, ok := m.GetValue().(*Field_String_); ok {
		return x.String_
	}
	return ""
}

func (m *Field) GetBytes() []byte {
	if x, ok := m.GetValue().(*Field_Bytes); ok {
		return x.Bytes
	}
	return nil
}

func (m *Field) GetFloat64() float64 {
	if x, ok := m.GetValue().(*Field_Float64); ok {
		return x.Float64
	}
	return 0
}

func (m *Field) GetTime() *time.Time {
	if x, ok := m.GetValue().(*Field_Time); ok {
		return x.Time
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Field) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Field_OneofMarshaler, _Field_OneofUnmarshaler, _Field_OneofSizer, []interface{}{
		(*Field_Bool)(nil),
		(*Field_Int)(nil),
		(*Field_Uint)(nil),
		(*Field_String_)(nil),
		(*Field_Bytes)(nil),
		(*Field_Float64)(nil),
		(*Field_Time)(nil),
	}
}

func _Field_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Field)
	// value
	switch x := m.Value.(type) {
	case *Field_Bool:
		t := uint64(0)
		if x.Bool {
			t = 1
		}
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(t)
	case *Field_Int:
		_ = b.EncodeVarint(3<<3 | proto.WireVarint)
		_ = b.EncodeZigzag64(uint64(x.Int))
	case *Field_Uint:
		_ = b.EncodeVarint(4<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Uint))
	case *Field_String_:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.String_)
	case *Field_Bytes:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.Bytes)
	case *Field_Float64:
		_ = b.EncodeVarint(7<<3 | proto.WireFixed64)
		_ = b.EncodeFixed64(math.Float64bits(x.Float64))
	case *Field_Time:
		_ = b.EncodeVarint(8<<3 | proto.WireBytes)
		dAtA, err := github_com_gogo_protobuf_types.StdTimeMarshal(*x.Time)
		if err != nil {
			return err
		}
		if err := b.EncodeRawBytes(dAtA); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Field.Value has unexpected type %T", x)
	}
	return nil
}

func _Field_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Field)
	switch tag {
	case 2: // value.bool
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Field_Bool{x != 0}
		return true, err
	case 3: // value.int
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeZigzag64()
		m.Value = &Field_Int{int64(x)}
		return true, err
	case 4: // value.uint
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Field_Uint{x}
		return true, err
	case 5: // value.string
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &Field_String_{x}
		return true, err
	case 6: // value.bytes
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Value = &Field_Bytes{x}
		return true, err
	case 7: // value.float64
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &Field_Float64{math.Float64frombits(x)}
		return true, err
	case 8: // value.time
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		if err != nil {
			return true, err
		}
		c := new(time.Time)
		if err2 := github_com_gogo_protobuf_types.StdTimeUnmarshal(c, x); err2 != nil {
			return true, err
		}
		m.Value = &Field_Time{c}
		return true, err
	default:
		return false, nil
	}
}

func _Field_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Field)
	// value
	switch x := m.Value.(type) {
	case *Field_Bool:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += 1
	case *Field_Int:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(uint64(x.Int<<1) ^ uint64((int64(x.Int) >> 63))))
	case *Field_Uint:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Uint))
	case *Field_String_:
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.String_)))
		n += len(x.String_)
	case *Field_Bytes:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Bytes)))
		n += len(x.Bytes)
	case *Field_Float64:
		n += proto.SizeVarint(7<<3 | proto.WireFixed64)
		n += 8
	case *Field_Time:
		s := github_com_gogo_protobuf_types.SizeOfStdTime(*x.Time)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SQLFilter struct {
	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	// Map of column name to field value.
	Fields map[string]*Field `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SQLFilter) Reset()                    { *m = SQLFilter{} }
func (m *SQLFilter) String() string            { return proto.CompactTextString(m) }
func (*SQLFilter) ProtoMessage()               {}
func (*SQLFilter) Descriptor() ([]byte, []int) { return fileDescriptorDependency, []int{1} }

func (m *SQLFilter) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *SQLFilter) GetFields() map[string]*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

type ExpirationTime struct {
	Time time.Time `protobuf:"bytes,1,opt,name=time,stdtime" json:"time"`
}

func (m *ExpirationTime) Reset()                    { *m = ExpirationTime{} }
func (m *ExpirationTime) String() string            { return proto.CompactTextString(m) }
func (*ExpirationTime) ProtoMessage()               {}
func (*ExpirationTime) Descriptor() ([]byte, []int) { return fileDescriptorDependency, []int{2} }

func (m *ExpirationTime) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Field)(nil), "thunderpb.Field")
	proto.RegisterType((*SQLFilter)(nil), "thunderpb.SQLFilter")
	proto.RegisterType((*ExpirationTime)(nil), "thunderpb.ExpirationTime")
	proto.RegisterEnum("thunderpb.FieldKind", FieldKind_name, FieldKind_value)
}
func (m *Field) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Field) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Kind != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDependency(dAtA, i, uint64(m.Kind))
	}
	if m.Value != nil {
		nn1, err := m.Value.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *Field_Bool) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x10
	i++
	if m.Bool {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}
func (m *Field_Int) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x18
	i++
	i = encodeVarintDependency(dAtA, i, uint64((uint64(m.Int)<<1)^uint64((m.Int>>63))))
	return i, nil
}
func (m *Field_Uint) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x20
	i++
	i = encodeVarintDependency(dAtA, i, uint64(m.Uint))
	return i, nil
}
func (m *Field_String_) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x2a
	i++
	i = encodeVarintDependency(dAtA, i, uint64(len(m.String_)))
	i += copy(dAtA[i:], m.String_)
	return i, nil
}
func (m *Field_Bytes) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Bytes != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintDependency(dAtA, i, uint64(len(m.Bytes)))
		i += copy(dAtA[i:], m.Bytes)
	}
	return i, nil
}
func (m *Field_Float64) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x39
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Float64))))
	i += 8
	return i, nil
}
func (m *Field_Time) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Time != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintDependency(dAtA, i, uint64(types.SizeOfStdTime(*m.Time)))
		n2, err := types.StdTimeMarshalTo(*m.Time, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *SQLFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SQLFilter) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Table) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDependency(dAtA, i, uint64(len(m.Table)))
		i += copy(dAtA[i:], m.Table)
	}
	if len(m.Fields) > 0 {
		for k, _ := range m.Fields {
			dAtA[i] = 0x12
			i++
			v := m.Fields[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovDependency(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovDependency(uint64(len(k))) + msgSize
			i = encodeVarintDependency(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintDependency(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintDependency(dAtA, i, uint64(v.Size()))
				n3, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n3
			}
		}
	}
	return i, nil
}

func (m *ExpirationTime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExpirationTime) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintDependency(dAtA, i, uint64(types.SizeOfStdTime(m.Time)))
	n4, err := types.StdTimeMarshalTo(m.Time, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func encodeVarintDependency(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Field) Size() (n int) {
	var l int
	_ = l
	if m.Kind != 0 {
		n += 1 + sovDependency(uint64(m.Kind))
	}
	if m.Value != nil {
		n += m.Value.Size()
	}
	return n
}

func (m *Field_Bool) Size() (n int) {
	var l int
	_ = l
	n += 2
	return n
}
func (m *Field_Int) Size() (n int) {
	var l int
	_ = l
	n += 1 + sozDependency(uint64(m.Int))
	return n
}
func (m *Field_Uint) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovDependency(uint64(m.Uint))
	return n
}
func (m *Field_String_) Size() (n int) {
	var l int
	_ = l
	l = len(m.String_)
	n += 1 + l + sovDependency(uint64(l))
	return n
}
func (m *Field_Bytes) Size() (n int) {
	var l int
	_ = l
	if m.Bytes != nil {
		l = len(m.Bytes)
		n += 1 + l + sovDependency(uint64(l))
	}
	return n
}
func (m *Field_Float64) Size() (n int) {
	var l int
	_ = l
	n += 9
	return n
}
func (m *Field_Time) Size() (n int) {
	var l int
	_ = l
	if m.Time != nil {
		l = types.SizeOfStdTime(*m.Time)
		n += 1 + l + sovDependency(uint64(l))
	}
	return n
}
func (m *SQLFilter) Size() (n int) {
	var l int
	_ = l
	l = len(m.Table)
	if l > 0 {
		n += 1 + l + sovDependency(uint64(l))
	}
	if len(m.Fields) > 0 {
		for k, v := range m.Fields {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovDependency(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovDependency(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovDependency(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *ExpirationTime) Size() (n int) {
	var l int
	_ = l
	l = types.SizeOfStdTime(m.Time)
	n += 1 + l + sovDependency(uint64(l))
	return n
}

func sovDependency(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDependency(x uint64) (n int) {
	return sovDependency(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Field) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDependency
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
			return fmt.Errorf("proto: Field: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Field: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			m.Kind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDependency
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kind |= (FieldKind(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bool", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDependency
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Value = &Field_Bool{b}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDependency
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.Value = &Field_Int{int64(v)}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uint", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDependency
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Value = &Field_Uint{v}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field String_", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDependency
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
				return ErrInvalidLengthDependency
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = &Field_String_{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDependency
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDependency
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Value = &Field_Bytes{v}
			iNdEx = postIndex
		case 7:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Float64", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = &Field_Float64{float64(math.Float64frombits(v))}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDependency
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
				return ErrInvalidLengthDependency
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := new(time.Time)
			if err := types.StdTimeUnmarshal(v, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Value = &Field_Time{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDependency(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDependency
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
func (m *SQLFilter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDependency
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
			return fmt.Errorf("proto: SQLFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SQLFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Table", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDependency
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
				return ErrInvalidLengthDependency
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Table = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fields", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDependency
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
				return ErrInvalidLengthDependency
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fields == nil {
				m.Fields = make(map[string]*Field)
			}
			var mapkey string
			var mapvalue *Field
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDependency
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDependency
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthDependency
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDependency
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthDependency
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthDependency
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Field{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipDependency(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthDependency
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Fields[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDependency(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDependency
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
func (m *ExpirationTime) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDependency
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
			return fmt.Errorf("proto: ExpirationTime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExpirationTime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDependency
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
				return ErrInvalidLengthDependency
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDependency(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDependency
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
func skipDependency(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDependency
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
					return 0, ErrIntOverflowDependency
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
					return 0, ErrIntOverflowDependency
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
				return 0, ErrInvalidLengthDependency
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDependency
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
				next, err := skipDependency(dAtA[start:])
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
	ErrInvalidLengthDependency = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDependency   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("dependency.proto", fileDescriptorDependency) }

var fileDescriptorDependency = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x33, 0xcd, 0xbf, 0xe6, 0x54, 0x96, 0x61, 0x28, 0x12, 0x7a, 0xd1, 0x86, 0x5e, 0x2c,
	0x41, 0x30, 0x85, 0xba, 0x2c, 0xc5, 0xcb, 0xc2, 0x96, 0xea, 0x8a, 0x60, 0xd6, 0xbd, 0xf1, 0x2e,
	0xd9, 0x4c, 0xd3, 0xa1, 0xe9, 0x4c, 0x4d, 0x26, 0x6a, 0xef, 0x7c, 0x04, 0x1f, 0xc3, 0x47, 0xd9,
	0x4b, 0x9f, 0x40, 0xa5, 0x4f, 0x22, 0x33, 0x49, 0x6b, 0xf1, 0xc6, 0xbb, 0xf3, 0x9d, 0xf9, 0xce,
	0x64, 0x7e, 0xdf, 0x09, 0xe0, 0x8c, 0xee, 0x28, 0xcf, 0x28, 0x7f, 0xd8, 0x47, 0xbb, 0x52, 0x48,
	0x41, 0x3c, 0xb9, 0xae, 0x79, 0x46, 0xcb, 0x5d, 0x3a, 0x78, 0x9e, 0x33, 0xb9, 0xae, 0xd3, 0xe8,
	0x41, 0x6c, 0x27, 0xb9, 0xc8, 0xc5, 0x44, 0x3b, 0xd2, 0x7a, 0xa5, 0x95, 0x16, 0xba, 0x6a, 0x26,
	0x07, 0xa3, 0x5c, 0x88, 0xbc, 0xa0, 0x7f, 0x5d, 0x92, 0x6d, 0x69, 0x25, 0x93, 0xed, 0xae, 0x31,
	0x8c, 0xbf, 0x76, 0xc0, 0x5e, 0x30, 0x5a, 0x64, 0x24, 0x04, 0x6b, 0xc3, 0x78, 0xe6, 0xa3, 0x00,
	0x85, 0x17, 0xd3, 0x7e, 0x74, 0xfa, 0x66, 0xa4, 0xcf, 0x6f, 0x19, 0xcf, 0x62, 0xed, 0x20, 0x7d,
	0xb0, 0x52, 0x21, 0x0a, 0xbf, 0x13, 0xa0, 0xb0, 0xbb, 0x34, 0x62, 0xad, 0x08, 0x01, 0x93, 0x71,
	0xe9, 0x9b, 0x01, 0x0a, 0xc9, 0xd2, 0x88, 0x95, 0x50, 0xce, 0x5a, 0x35, 0xad, 0x00, 0x85, 0x96,
	0x72, 0x2a, 0x45, 0x7c, 0x70, 0x2a, 0x59, 0x32, 0x9e, 0xfb, 0x76, 0x80, 0x42, 0x6f, 0x69, 0xc4,
	0xad, 0x26, 0x4f, 0xc1, 0x4e, 0xf7, 0x92, 0x56, 0xbe, 0x13, 0xa0, 0xf0, 0xc9, 0xd2, 0x88, 0x1b,
	0x49, 0x06, 0xe0, 0xae, 0x0a, 0x91, 0xc8, 0xeb, 0x2b, 0xdf, 0x0d, 0x50, 0x88, 0x96, 0x46, 0x7c,
	0x6c, 0x90, 0x6b, 0xb0, 0x14, 0x94, 0xdf, 0x0d, 0x50, 0xd8, 0x9b, 0x0e, 0xa2, 0x86, 0x38, 0x3a,
	0x12, 0x47, 0xef, 0x8f, 0xc4, 0x73, 0xeb, 0xdb, 0xaf, 0x91, 0x1a, 0xd5, 0xfe, 0xb9, 0x0b, 0xf6,
	0xa7, 0xa4, 0xa8, 0xe9, 0xf8, 0x3b, 0x02, 0xef, 0xee, 0xdd, 0x9b, 0x05, 0x2b, 0x24, 0x2d, 0x49,
	0x1f, 0x6c, 0x99, 0xa4, 0x05, 0xd5, 0x39, 0x78, 0x71, 0x23, 0xc8, 0x0c, 0x9c, 0x95, 0x4a, 0xa1,
	0xf2, 0x3b, 0x81, 0x19, 0xf6, 0xa6, 0xc1, 0x59, 0x3c, 0xa7, 0xd9, 0x26, 0xa8, 0xea, 0x86, 0xcb,
	0x72, 0x1f, 0xb7, 0xfe, 0xc1, 0x2d, 0xf4, 0xce, 0xda, 0x04, 0x83, 0xb9, 0xa1, 0xfb, 0xf6, 0x72,
	0x55, 0x92, 0xcb, 0xf6, 0x1d, 0x3a, 0xce, 0xde, 0x14, 0xff, 0x1b, 0x7c, 0xdc, 0x1c, 0xbf, 0xec,
	0xcc, 0xd0, 0xf8, 0x35, 0x5c, 0xdc, 0x7c, 0xd9, 0xb1, 0x32, 0x91, 0x4c, 0x70, 0x05, 0x46, 0x66,
	0x2d, 0x3d, 0xfa, 0x2f, 0x7d, 0xf7, 0xf1, 0xe7, 0xc8, 0x50, 0x09, 0x34, 0xfc, 0xcf, 0xb6, 0xe0,
	0x9d, 0x16, 0x4b, 0x7a, 0xe0, 0xde, 0xf3, 0x0d, 0x17, 0x9f, 0x39, 0x36, 0x48, 0x17, 0xac, 0xb7,
	0x75, 0x51, 0x60, 0xa4, 0xaa, 0xb9, 0x10, 0x05, 0xee, 0x10, 0x17, 0xcc, 0x57, 0x5c, 0x62, 0x53,
	0xb5, 0xee, 0x19, 0x97, 0xd8, 0x22, 0x00, 0xce, 0x9d, 0x5e, 0x1b, 0xb6, 0x89, 0x07, 0xf6, 0x5c,
	0x6d, 0x0a, 0x3b, 0xea, 0xaa, 0x45, 0xb3, 0x1a, 0xec, 0x2a, 0xb7, 0x7a, 0x01, 0xee, 0xce, 0xaf,
	0x1e, 0x0f, 0x43, 0xf4, 0xe3, 0x30, 0x44, 0xbf, 0x0f, 0x43, 0xf4, 0xe1, 0xf2, 0xec, 0x37, 0xae,
	0x92, 0x6d, 0x95, 0x94, 0xc9, 0xfa, 0xe3, 0xa4, 0x25, 0x9f, 0x9c, 0x12, 0x48, 0x1d, 0x0d, 0xf2,
	0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0xa6, 0x56, 0xfc, 0x14, 0x03, 0x00, 0x00,
}
