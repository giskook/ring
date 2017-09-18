// Code generated by protoc-gen-go.
// source: report.proto
// DO NOT EDIT!

package Carrier

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Report_Protocol int32

const (
	Report_UNKNOWN Report_Protocol = 0
	Report_LOGIN   Report_Protocol = 1
)

var Report_Protocol_name = map[int32]string{
	0: "UNKNOWN",
	1: "LOGIN",
}
var Report_Protocol_value = map[string]int32{
	"UNKNOWN": 0,
	"LOGIN":   1,
}

func (x Report_Protocol) String() string {
	return proto.EnumName(Report_Protocol_name, int32(x))
}
func (Report_Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

type Report struct {
	Header   *Header         `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Protocol Report_Protocol `protobuf:"varint,2,opt,name=protocol,enum=Carrier.Report_Protocol" json:"protocol,omitempty"`
	Login    *ReportLogin    `protobuf:"bytes,3,opt,name=login" json:"login,omitempty"`
}

func (m *Report) Reset()                    { *m = Report{} }
func (m *Report) String() string            { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()               {}
func (*Report) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Report) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Report) GetProtocol() Report_Protocol {
	if m != nil {
		return m.Protocol
	}
	return Report_UNKNOWN
}

func (m *Report) GetLogin() *ReportLogin {
	if m != nil {
		return m.Login
	}
	return nil
}

func init() {
	proto.RegisterType((*Report)(nil), "Carrier.Report")
	proto.RegisterEnum("Carrier.Report_Protocol", Report_Protocol_name, Report_Protocol_value)
}

func init() { proto.RegisterFile("report.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4a, 0x2d, 0xc8,
	0x2f, 0x2a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x77, 0x4e, 0x2c, 0x2a, 0xca, 0x4c,
	0x2d, 0x92, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e, 0x85, 0x08, 0x4a, 0x09, 0x41, 0x94, 0xc4, 0xe7, 0xe4,
	0xa7, 0x67, 0xe6, 0x41, 0xc4, 0x94, 0xe6, 0x31, 0x72, 0xb1, 0x05, 0x81, 0x85, 0x85, 0xe4, 0xb9,
	0xd8, 0x32, 0x52, 0x13, 0x53, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xf8, 0xf5,
	0xa0, 0x86, 0xe8, 0x79, 0x80, 0x85, 0x85, 0xb4, 0xb8, 0x38, 0xc0, 0x9a, 0x92, 0xf3, 0x73, 0x24,
	0x98, 0x14, 0x18, 0x35, 0xf8, 0x8c, 0x24, 0xe0, 0x4a, 0x20, 0x66, 0xe8, 0x05, 0x40, 0xe5, 0x85,
	0x94, 0xb9, 0x58, 0xc1, 0xd6, 0x48, 0x30, 0x83, 0xcd, 0x12, 0x41, 0x53, 0xe8, 0x03, 0x92, 0x53,
	0x52, 0xe2, 0xe2, 0x80, 0x6b, 0xe0, 0xe6, 0x62, 0x0f, 0xf5, 0xf3, 0xf6, 0xf3, 0x0f, 0xf7, 0x13,
	0x60, 0x10, 0xe2, 0xe4, 0x62, 0xf5, 0xf1, 0x77, 0xf7, 0xf4, 0x13, 0x60, 0x4c, 0x62, 0x03, 0x5b,
	0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x90, 0x25, 0xc1, 0xe0, 0x00, 0x00, 0x00,
}
