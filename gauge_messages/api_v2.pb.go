// Code generated by protoc-gen-go.
// source: api_v2.proto
// DO NOT EDIT!

package gauge_messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScenarioExecutionResult_Status int32

const (
	ScenarioExecutionResult_PASSED  ScenarioExecutionResult_Status = 0
	ScenarioExecutionResult_FAILED  ScenarioExecutionResult_Status = 1
	ScenarioExecutionResult_SKIPPED ScenarioExecutionResult_Status = 2
)

var ScenarioExecutionResult_Status_name = map[int32]string{
	0: "PASSED",
	1: "FAILED",
	2: "SKIPPED",
}
var ScenarioExecutionResult_Status_value = map[string]int32{
	"PASSED":  0,
	"FAILED":  1,
	"SKIPPED": 2,
}

func (x ScenarioExecutionResult_Status) String() string {
	return proto.EnumName(ScenarioExecutionResult_Status_name, int32(x))
}
func (ScenarioExecutionResult_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 0}
}

type ExecutionRequest struct {
	Specs []string `protobuf:"bytes,1,rep,name=specs" json:"specs,omitempty"`
	Args  []string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
}

func (m *ExecutionRequest) Reset()                    { *m = ExecutionRequest{} }
func (m *ExecutionRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecutionRequest) ProtoMessage()               {}
func (*ExecutionRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ScenarioExecutionResult struct {
	ID            string                         `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Status        ScenarioExecutionResult_Status `protobuf:"varint,2,opt,name=status,enum=gauge.messages.ScenarioExecutionResult_Status" json:"status,omitempty"`
	ExecutionTime int64                          `protobuf:"varint,3,opt,name=executionTime" json:"executionTime,omitempty"`
	Message       string                         `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	StackTrace    string                         `protobuf:"bytes,5,opt,name=stackTrace" json:"stackTrace,omitempty"`
}

func (m *ScenarioExecutionResult) Reset()                    { *m = ScenarioExecutionResult{} }
func (m *ScenarioExecutionResult) String() string            { return proto.CompactTextString(m) }
func (*ScenarioExecutionResult) ProtoMessage()               {}
func (*ScenarioExecutionResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*ExecutionRequest)(nil), "gauge.messages.ExecutionRequest")
	proto.RegisterType((*ScenarioExecutionResult)(nil), "gauge.messages.ScenarioExecutionResult")
	proto.RegisterEnum("gauge.messages.ScenarioExecutionResult_Status", ScenarioExecutionResult_Status_name, ScenarioExecutionResult_Status_value)
}

var fileDescriptor1 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x97, 0x76, 0x6b, 0xd9, 0x73, 0x9b, 0x23, 0x20, 0x06, 0x4f, 0xa3, 0x17, 0x77, 0x31,
	0x4a, 0xbd, 0x0b, 0x83, 0x4d, 0x28, 0x7a, 0x28, 0x66, 0x07, 0x6f, 0x12, 0xc3, 0xa3, 0x14, 0xdd,
	0x5a, 0xfb, 0x12, 0xf1, 0x63, 0xfa, 0x91, 0x4c, 0xeb, 0x14, 0x1d, 0x08, 0xde, 0xde, 0x4b, 0xf8,
	0xfd, 0x7f, 0x2f, 0x79, 0x30, 0xd2, 0x75, 0xf9, 0xf0, 0x9a, 0xca, 0xba, 0xa9, 0x6c, 0xc5, 0x27,
	0x85, 0x76, 0x05, 0xca, 0x0d, 0x12, 0xe9, 0x02, 0x29, 0x39, 0x87, 0xe9, 0xea, 0x0d, 0x8d, 0xb3,
	0x65, 0xb5, 0xbd, 0xc3, 0x17, 0x87, 0x64, 0xf9, 0x18, 0x06, 0x54, 0xa3, 0x21, 0xc1, 0x66, 0xe1,
	0x7c, 0xc8, 0x47, 0xd0, 0xd7, 0x4d, 0x41, 0x22, 0x68, 0xbb, 0xe4, 0x9d, 0xc1, 0xb1, 0x32, 0xb8,
	0xd5, 0x4d, 0x59, 0xfd, 0x20, 0xc9, 0x3d, 0x5b, 0x0e, 0x10, 0x64, 0x4b, 0x4f, 0x31, 0x4f, 0x5d,
	0x41, 0x44, 0x56, 0x5b, 0xd7, 0x72, 0x6c, 0x3e, 0x49, 0xa5, 0xfc, 0x6d, 0x96, 0x7f, 0x84, 0x48,
	0xd5, 0x51, 0xfc, 0x08, 0xc6, 0xf8, 0x75, 0xb3, 0x2e, 0x37, 0x28, 0x42, 0x1f, 0x13, 0xf2, 0x43,
	0x88, 0x77, 0x09, 0xa2, 0xdf, 0x79, 0xbc, 0xd4, 0x7b, 0xcc, 0xd3, 0xba, 0xd1, 0x06, 0xc5, 0xa0,
	0x3d, 0x4b, 0xce, 0x20, 0xda, 0xa5, 0x00, 0x44, 0xf9, 0x42, 0xa9, 0xd5, 0x72, 0xda, 0x6b, 0xeb,
	0xeb, 0x45, 0x76, 0xeb, 0x6b, 0xc6, 0x0f, 0x20, 0x56, 0x37, 0x59, 0x9e, 0xfb, 0x26, 0x48, 0x11,
	0x86, 0xdf, 0x43, 0xf0, 0x7b, 0x88, 0x3f, 0xbd, 0xc8, 0x67, 0xfb, 0x23, 0xef, 0xff, 0xd4, 0xc9,
	0xe9, 0x3f, 0x1f, 0x95, 0xf4, 0x2e, 0xd8, 0x63, 0xd4, 0x6d, 0xe0, 0xf2, 0x23, 0x00, 0x00, 0xff,
	0xff, 0xba, 0x44, 0x4a, 0xf4, 0x91, 0x01, 0x00, 0x00,
}
