// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: util/metric/metric.proto

package metric

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
	_go "github.com/prometheus/client_model/go"
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

// DisplayUnit describes how the metric's units should be displayed in charts.
type Unit int32

const (
	// UNSET expresses that the metric's DisplayUnit wasn't explicitly set.
	Unit_UNSET Unit = 0
	// BYTES expresses that the metric's measurement is in bytes.
	Unit_BYTES Unit = 1
	// CONST expresses that the metric's measurement is a constant value.
	Unit_CONST Unit = 2
	// COUNT expresses that the metric's measurement is a count.
	Unit_COUNT Unit = 3
	// NANOSECONDS expresses that the metric's measurement is in nanoseconds.
	Unit_NANOSECONDS Unit = 4
	// PERCENT expresses that the metric's measurement is a percentage value.
	Unit_PERCENT Unit = 5
	// SECONDS expresses that the metric's measurement is in seconds.
	Unit_SECONDS Unit = 6
	// TIMESTAMP_NS expresses that the metric's measurement is a time since the
	// Unix epoch in nanoseconds.
	Unit_TIMESTAMP_NS Unit = 7
	// TIMESTAMP_SEC expresses that the metric's measurement is a time since the
	// Unix epoch in seconds.
	Unit_TIMESTAMP_SEC Unit = 8
)

var Unit_name = map[int32]string{
	0: "UNSET",
	1: "BYTES",
	2: "CONST",
	3: "COUNT",
	4: "NANOSECONDS",
	5: "PERCENT",
	6: "SECONDS",
	7: "TIMESTAMP_NS",
	8: "TIMESTAMP_SEC",
}

var Unit_value = map[string]int32{
	"UNSET":         0,
	"BYTES":         1,
	"CONST":         2,
	"COUNT":         3,
	"NANOSECONDS":   4,
	"PERCENT":       5,
	"SECONDS":       6,
	"TIMESTAMP_NS":  7,
	"TIMESTAMP_SEC": 8,
}

func (x Unit) Enum() *Unit {
	p := new(Unit)
	*p = x
	return p
}

func (x Unit) String() string {
	return proto.EnumName(Unit_name, int32(x))
}

func (x *Unit) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Unit_value, data, "Unit")
	if err != nil {
		return err
	}
	*x = Unit(value)
	return nil
}

func (Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d59dd1524693214, []int{0}
}

// metric.LabelPair is a proxy for io.prometheus.client.LabelPair.
// io.prometheus.client.LabelPair doesn't support gogoproto.marshaler
// and gogoproto.unmarshaler which are required by gRPC. metric.LabelPair
// stores information that is similarly structured, supports the requisite
// gogoproto options, and is convertible to io.prometheus.client.LabelPair
// to satisfy PrometheusExportable's GetLabels method.
type LabelPair struct {
	Name  *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *LabelPair) Reset()         { *m = LabelPair{} }
func (m *LabelPair) String() string { return proto.CompactTextString(m) }
func (*LabelPair) ProtoMessage()    {}
func (*LabelPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d59dd1524693214, []int{0}
}
func (m *LabelPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LabelPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *LabelPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelPair.Merge(m, src)
}
func (m *LabelPair) XXX_Size() int {
	return m.Size()
}
func (m *LabelPair) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelPair.DiscardUnknown(m)
}

var xxx_messageInfo_LabelPair proto.InternalMessageInfo

// Metadata holds metadata about a metric. It must be embedded in
// each metric object. It's used to export information about the
// metric to Prometheus and for Admin UI charts.
type Metadata struct {
	Name        string         `protobuf:"bytes,1,req,name=name" json:"name"`
	Help        string         `protobuf:"bytes,2,req,name=help" json:"help"`
	Measurement string         `protobuf:"bytes,3,req,name=measurement" json:"measurement"`
	Unit        Unit           `protobuf:"varint,4,req,name=unit,enum=cockroach.util.metric.Unit" json:"unit"`
	MetricType  _go.MetricType `protobuf:"varint,5,opt,name=metricType,enum=io.prometheus.client.MetricType" json:"metricType"`
	Labels      []*LabelPair   `protobuf:"bytes,6,rep,name=labels" json:"labels,omitempty"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d59dd1524693214, []int{1}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return m.Size()
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("cockroach.util.metric.Unit", Unit_name, Unit_value)
	proto.RegisterType((*LabelPair)(nil), "cockroach.util.metric.LabelPair")
	proto.RegisterType((*Metadata)(nil), "cockroach.util.metric.Metadata")
}

func init() { proto.RegisterFile("util/metric/metric.proto", fileDescriptor_8d59dd1524693214) }

var fileDescriptor_8d59dd1524693214 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x8a, 0x13, 0x41,
	0x10, 0xc6, 0xa7, 0x27, 0x93, 0xec, 0xa6, 0xa2, 0x6b, 0xdb, 0xac, 0xd0, 0x44, 0x18, 0x87, 0x3d,
	0xc8, 0xe0, 0x61, 0x16, 0x02, 0x0b, 0x5e, 0x37, 0x71, 0x04, 0xc1, 0x74, 0x42, 0xba, 0x73, 0xd0,
	0xcb, 0xd2, 0xce, 0x36, 0x9b, 0xc1, 0xf9, 0x13, 0x26, 0x3d, 0x82, 0x57, 0x9f, 0xc0, 0x37, 0xf0,
	0x75, 0x72, 0xdc, 0xe3, 0x9e, 0x44, 0x93, 0x17, 0x91, 0x4e, 0x27, 0x71, 0x0e, 0x7a, 0xea, 0xaf,
	0xeb, 0xfb, 0x55, 0x51, 0xf5, 0x01, 0xad, 0x75, 0x9a, 0x5d, 0xe6, 0x4a, 0x57, 0x69, 0xb2, 0x7f,
	0xa2, 0x65, 0x55, 0xea, 0x92, 0x3c, 0x4b, 0xca, 0xe4, 0x73, 0x55, 0xca, 0x64, 0x11, 0x19, 0x26,
	0xb2, 0x66, 0xff, 0xfc, 0xae, 0xbc, 0x2b, 0x77, 0xc4, 0xa5, 0x51, 0x16, 0xee, 0xf7, 0x93, 0x2c,
	0x55, 0x85, 0xbe, 0xc9, 0xcb, 0x5b, 0x75, 0x18, 0xb7, 0xb2, 0xde, 0xc5, 0x15, 0x74, 0xdf, 0xcb,
	0x4f, 0x2a, 0x9b, 0xca, 0xb4, 0x22, 0x04, 0xbc, 0x42, 0xe6, 0x8a, 0xa2, 0x00, 0x85, 0xdd, 0xd9,
	0x4e, 0x93, 0x73, 0x68, 0x7f, 0x91, 0x59, 0xad, 0xa8, 0xbb, 0x2b, 0xda, 0xcf, 0xc5, 0x0f, 0x17,
	0x4e, 0xc7, 0x4a, 0xcb, 0x5b, 0xa9, 0x25, 0xa1, 0xc7, 0x36, 0x37, 0xec, 0x0e, 0xbd, 0xf5, 0xcf,
	0x17, 0xce, 0xbe, 0x99, 0x82, 0xb7, 0x50, 0xd9, 0x92, 0xba, 0x4d, 0xc7, 0x54, 0xc8, 0x4b, 0xe8,
	0xe5, 0x4a, 0xae, 0xea, 0x4a, 0xe5, 0xaa, 0xd0, 0xb4, 0xd5, 0x00, 0x9a, 0x06, 0xb9, 0x02, 0xaf,
	0x2e, 0x52, 0x4d, 0xbd, 0xc0, 0x0d, 0xcf, 0x06, 0xcf, 0xa3, 0x7f, 0xde, 0x1d, 0xcd, 0x8b, 0x54,
	0x1f, 0xc6, 0x1b, 0x9c, 0xbc, 0x05, 0xb0, 0x96, 0xf8, 0xba, 0x54, 0xb4, 0x1d, 0xa0, 0xf0, 0x6c,
	0x10, 0x44, 0xe9, 0x2e, 0x91, 0x5c, 0xe9, 0x85, 0xaa, 0x57, 0x91, 0x4d, 0x25, 0x1a, 0x1f, 0xb9,
	0xfd, 0x84, 0x46, 0x27, 0x79, 0x0d, 0x9d, 0xcc, 0xc4, 0xb3, 0xa2, 0x9d, 0xa0, 0x15, 0xf6, 0x06,
	0xc1, 0x7f, 0x16, 0x38, 0x66, 0x38, 0xdb, 0xf3, 0xaf, 0xbe, 0x21, 0xf0, 0xcc, 0x5a, 0xa4, 0x0b,
	0xed, 0x39, 0xe3, 0xb1, 0xc0, 0x8e, 0x91, 0xc3, 0x0f, 0x22, 0xe6, 0x18, 0x19, 0x39, 0x9a, 0x30,
	0x2e, 0xb0, 0x6b, 0xe5, 0x9c, 0x09, 0xdc, 0x22, 0x4f, 0xa0, 0xc7, 0xae, 0xd9, 0x84, 0xc7, 0xa3,
	0x09, 0x7b, 0xc3, 0xb1, 0x47, 0x7a, 0x70, 0x32, 0x8d, 0x67, 0xa3, 0x98, 0x09, 0xdc, 0x36, 0x9f,
	0x83, 0xd3, 0x21, 0x18, 0x1e, 0x89, 0x77, 0xe3, 0x98, 0x8b, 0xeb, 0xf1, 0xf4, 0x86, 0x71, 0x7c,
	0x42, 0x9e, 0xc2, 0xe3, 0xbf, 0x15, 0x1e, 0x8f, 0xf0, 0xe9, 0x30, 0x5c, 0xff, 0xf6, 0x9d, 0xf5,
	0xc6, 0x47, 0xf7, 0x1b, 0x1f, 0x3d, 0x6c, 0x7c, 0xf4, 0x6b, 0xe3, 0xa3, 0xef, 0x5b, 0xdf, 0xb9,
	0xdf, 0xfa, 0xce, 0xc3, 0xd6, 0x77, 0x3e, 0x76, 0xec, 0x01, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xa4, 0xfc, 0xd5, 0x6a, 0x6b, 0x02, 0x00, 0x00,
}

func (m *LabelPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LabelPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LabelPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		i -= len(*m.Value)
		copy(dAtA[i:], *m.Value)
		i = encodeVarintMetric(dAtA, i, uint64(len(*m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if m.Name != nil {
		i -= len(*m.Name)
		copy(dAtA[i:], *m.Name)
		i = encodeVarintMetric(dAtA, i, uint64(len(*m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Metadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for iNdEx := len(m.Labels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Labels[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMetric(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	i = encodeVarintMetric(dAtA, i, uint64(m.MetricType))
	i--
	dAtA[i] = 0x28
	i = encodeVarintMetric(dAtA, i, uint64(m.Unit))
	i--
	dAtA[i] = 0x20
	i -= len(m.Measurement)
	copy(dAtA[i:], m.Measurement)
	i = encodeVarintMetric(dAtA, i, uint64(len(m.Measurement)))
	i--
	dAtA[i] = 0x1a
	i -= len(m.Help)
	copy(dAtA[i:], m.Help)
	i = encodeVarintMetric(dAtA, i, uint64(len(m.Help)))
	i--
	dAtA[i] = 0x12
	i -= len(m.Name)
	copy(dAtA[i:], m.Name)
	i = encodeVarintMetric(dAtA, i, uint64(len(m.Name)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMetric(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetric(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LabelPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Name != nil {
		l = len(*m.Name)
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.Value != nil {
		l = len(*m.Value)
		n += 1 + l + sovMetric(uint64(l))
	}
	return n
}

func (m *Metadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovMetric(uint64(l))
	l = len(m.Help)
	n += 1 + l + sovMetric(uint64(l))
	l = len(m.Measurement)
	n += 1 + l + sovMetric(uint64(l))
	n += 1 + sovMetric(uint64(m.Unit))
	n += 1 + sovMetric(uint64(m.MetricType))
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.Size()
			n += 1 + l + sovMetric(uint64(l))
		}
	}
	return n
}

func sovMetric(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetric(x uint64) (n int) {
	return sovMetric(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LabelPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
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
			return fmt.Errorf("proto: LabelPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LabelPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetric
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Name = &s
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetric
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Value = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetric
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
func (m *Metadata) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetric
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Help", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetric
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Help = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Measurement", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetric
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Measurement = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			m.Unit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Unit |= Unit(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000008)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricType", wireType)
			}
			m.MetricType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MetricType |= _go.MetricType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
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
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetric
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = append(m.Labels, &LabelPair{})
			if err := m.Labels[len(m.Labels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMetric
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("name")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("help")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("measurement")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("unit")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMetric(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetric
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
					return 0, ErrIntOverflowMetric
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
					return 0, ErrIntOverflowMetric
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
				return 0, ErrInvalidLengthMetric
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMetric
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMetric
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMetric        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetric          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMetric = fmt.Errorf("proto: unexpected end of group")
)
