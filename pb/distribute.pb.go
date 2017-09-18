// Code generated by protoc-gen-go.
// source: distribute.proto
// DO NOT EDIT!

package Carrier

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Distribute_Protocol int32

const (
	Distribute_UNKNOWN Distribute_Protocol = 0
	Distribute_LOGRT   Distribute_Protocol = 1
)

var Distribute_Protocol_name = map[int32]string{
	0: "UNKNOWN",
	1: "LOGRT",
}
var Distribute_Protocol_value = map[string]int32{
	"UNKNOWN": 0,
	"LOGRT":   1,
}

func (x Distribute_Protocol) String() string {
	return proto.EnumName(Distribute_Protocol_name, int32(x))
}
func (Distribute_Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

type Distribute struct {
	Header   *Header             `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Protocol Distribute_Protocol `protobuf:"varint,2,opt,name=protocol,enum=Carrier.Distribute_Protocol" json:"protocol,omitempty"`
	Logrt    *DistributeLogRt    `protobuf:"bytes,3,opt,name=logrt" json:"logrt,omitempty"`
}

func (m *Distribute) Reset()                    { *m = Distribute{} }
func (m *Distribute) String() string            { return proto.CompactTextString(m) }
func (*Distribute) ProtoMessage()               {}
func (*Distribute) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Distribute) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Distribute) GetProtocol() Distribute_Protocol {
	if m != nil {
		return m.Protocol
	}
	return Distribute_UNKNOWN
}

func (m *Distribute) GetLogrt() *DistributeLogRt {
	if m != nil {
		return m.Logrt
	}
	return nil
}

func init() {
	proto.RegisterType((*Distribute)(nil), "Carrier.Distribute")
	proto.RegisterEnum("Carrier.Distribute_Protocol", Distribute_Protocol_name, Distribute_Protocol_value)
}

func init() { proto.RegisterFile("distribute.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xc9, 0x2c, 0x2e,
	0x29, 0xca, 0x4c, 0x2a, 0x2d, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x77, 0x4e,
	0x2c, 0x2a, 0xca, 0x4c, 0x2d, 0x92, 0xe2, 0x4a, 0x4a, 0x2c, 0x86, 0x0a, 0x4a, 0x89, 0x21, 0x94,
	0xc5, 0xe7, 0xe4, 0xa7, 0x17, 0x95, 0x40, 0xc4, 0x95, 0x56, 0x31, 0x72, 0x71, 0xb9, 0xc0, 0xa5,
	0x84, 0xe4, 0xb9, 0xd8, 0x32, 0x52, 0x13, 0x53, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8,
	0x8d, 0xf8, 0xf5, 0xa0, 0x86, 0xe9, 0x79, 0x80, 0x85, 0x85, 0xf4, 0xb8, 0x38, 0xc0, 0x1a, 0x93,
	0xf3, 0x73, 0x24, 0x98, 0x14, 0x18, 0x35, 0xf8, 0x8c, 0x64, 0xe0, 0x4a, 0x10, 0xe6, 0xe8, 0x05,
	0x40, 0xd5, 0x08, 0xa9, 0x73, 0xb1, 0x82, 0xad, 0x93, 0x60, 0x06, 0x9b, 0x27, 0x81, 0x45, 0xb1,
	0x4f, 0x7e, 0x7a, 0x50, 0x89, 0x92, 0x12, 0x17, 0x07, 0x5c, 0x13, 0x37, 0x17, 0x7b, 0xa8, 0x9f,
	0xb7, 0x9f, 0x7f, 0xb8, 0x9f, 0x00, 0x83, 0x10, 0x27, 0x17, 0xab, 0x8f, 0xbf, 0x7b, 0x50, 0x88,
	0x00, 0x63, 0x12, 0x1b, 0xd8, 0x6a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xc8, 0x9f,
	0xa4, 0xf4, 0x00, 0x00, 0x00,
}
