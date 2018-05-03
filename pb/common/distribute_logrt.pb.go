// Code generated by protoc-gen-go.
// source: distribute_logrt.proto
// DO NOT EDIT!

package Carrier

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DistributeLogRt struct {
	// imei
	Imei string `protobuf:"bytes,1,opt,name=imei" json:"imei,omitempty"`
	// time 格式为YYMMDD-HHMMSS
	// 例如170918-091800表示
	// 2017年09月18日09时18分00秒
	// 勿忘国耻!
	Time string `protobuf:"bytes,2,opt,name=time" json:"time,omitempty"`
	// 1 登录成功
	// 2 非法用户
	Result string `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
	// 默认参数 (terminal version > 106)
	// key 的可选值 : 0 <- 定位频度 单位 分
	//              : 1 <- 低电告警
	Settings map[string]string `protobuf:"bytes,4,rep,name=settings" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DistributeLogRt) Reset()                    { *m = DistributeLogRt{} }
func (m *DistributeLogRt) String() string            { return proto.CompactTextString(m) }
func (*DistributeLogRt) ProtoMessage()               {}
func (*DistributeLogRt) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *DistributeLogRt) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *DistributeLogRt) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *DistributeLogRt) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *DistributeLogRt) GetSettings() map[string]string {
	if m != nil {
		return m.Settings
	}
	return nil
}

func init() {
	proto.RegisterType((*DistributeLogRt)(nil), "Carrier.DistributeLogRt")
}

func init() { proto.RegisterFile("distribute_logrt.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xc9, 0x2c, 0x2e,
	0x29, 0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x8d, 0xcf, 0xc9, 0x4f, 0x2f, 0x2a, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x77, 0x4e, 0x2c, 0x2a, 0xca, 0x4c, 0x2d, 0x52, 0x5a, 0xc1, 0xc8, 0xc5,
	0xef, 0x02, 0x57, 0xe3, 0x93, 0x9f, 0x1e, 0x54, 0x22, 0xc4, 0xc3, 0xc5, 0x92, 0x99, 0x9b, 0x9a,
	0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x09, 0xe2, 0x95, 0x64, 0xe6, 0xa6, 0x4a, 0x30, 0x81, 0x79,
	0x7c, 0x5c, 0x6c, 0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x12, 0xcc, 0x60, 0xbe, 0x05, 0x17, 0x47,
	0x71, 0x6a, 0x49, 0x49, 0x66, 0x5e, 0x7a, 0xb1, 0x04, 0x8b, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x9a,
	0x1e, 0xd4, 0x6c, 0x3d, 0x34, 0x73, 0xf5, 0x82, 0xa1, 0x0a, 0x5d, 0xf3, 0x4a, 0x8a, 0x2a, 0xa5,
	0xf4, 0xb9, 0x78, 0x51, 0x04, 0x84, 0xb8, 0xb9, 0x98, 0xb3, 0x53, 0x2b, 0xa1, 0xb6, 0xf2, 0x72,
	0xb1, 0x96, 0x25, 0xe6, 0x94, 0x42, 0xad, 0xb5, 0x62, 0xb2, 0x60, 0x4c, 0x62, 0x03, 0x3b, 0xdd,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xbd, 0x11, 0xed, 0xd4, 0x00, 0x00, 0x00,
}
