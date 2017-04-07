// Code generated by protoc-gen-go.
// source: egress.proto
// DO NOT EDIT!

/*
Package loggregator_v2 is a generated protocol buffer package.

It is generated from these files:
	egress.proto
	envelope.proto
	ingress.proto

It has these top-level messages:
	EgressRequest
	Filter
	LogFilter
	Envelope
	Value
	Log
	Counter
	Gauge
	GaugeValue
	Timer
	IngressResponse
*/
package loggregator_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EgressRequest struct {
	ShardId string  `protobuf:"bytes,1,opt,name=shard_id,json=shardId" json:"shard_id,omitempty"`
	Filter  *Filter `protobuf:"bytes,2,opt,name=filter" json:"filter,omitempty"`
}

func (m *EgressRequest) Reset()                    { *m = EgressRequest{} }
func (m *EgressRequest) String() string            { return proto.CompactTextString(m) }
func (*EgressRequest) ProtoMessage()               {}
func (*EgressRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EgressRequest) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type Filter struct {
	SourceId string `protobuf:"bytes,1,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	// Types that are valid to be assigned to Message:
	//	*Filter_Log
	Message isFilter_Message `protobuf_oneof:"Message"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isFilter_Message interface {
	isFilter_Message()
}

type Filter_Log struct {
	Log *LogFilter `protobuf:"bytes,2,opt,name=log,oneof"`
}

func (*Filter_Log) isFilter_Message() {}

func (m *Filter) GetMessage() isFilter_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Filter) GetLog() *LogFilter {
	if x, ok := m.GetMessage().(*Filter_Log); ok {
		return x.Log
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Filter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Filter_OneofMarshaler, _Filter_OneofUnmarshaler, _Filter_OneofSizer, []interface{}{
		(*Filter_Log)(nil),
	}
}

func _Filter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Filter)
	// Message
	switch x := m.Message.(type) {
	case *Filter_Log:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Log); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Filter.Message has unexpected type %T", x)
	}
	return nil
}

func _Filter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Filter)
	switch tag {
	case 2: // Message.log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogFilter)
		err := b.DecodeMessage(msg)
		m.Message = &Filter_Log{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Filter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Filter)
	// Message
	switch x := m.Message.(type) {
	case *Filter_Log:
		s := proto.Size(x.Log)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LogFilter struct {
}

func (m *LogFilter) Reset()                    { *m = LogFilter{} }
func (m *LogFilter) String() string            { return proto.CompactTextString(m) }
func (*LogFilter) ProtoMessage()               {}
func (*LogFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*EgressRequest)(nil), "loggregator.v2.EgressRequest")
	proto.RegisterType((*Filter)(nil), "loggregator.v2.Filter")
	proto.RegisterType((*LogFilter)(nil), "loggregator.v2.LogFilter")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Egress service

type EgressClient interface {
	Receiver(ctx context.Context, in *EgressRequest, opts ...grpc.CallOption) (Egress_ReceiverClient, error)
}

type egressClient struct {
	cc *grpc.ClientConn
}

func NewEgressClient(cc *grpc.ClientConn) EgressClient {
	return &egressClient{cc}
}

func (c *egressClient) Receiver(ctx context.Context, in *EgressRequest, opts ...grpc.CallOption) (Egress_ReceiverClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Egress_serviceDesc.Streams[0], c.cc, "/loggregator.v2.Egress/Receiver", opts...)
	if err != nil {
		return nil, err
	}
	x := &egressReceiverClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Egress_ReceiverClient interface {
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type egressReceiverClient struct {
	grpc.ClientStream
}

func (x *egressReceiverClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Egress service

type EgressServer interface {
	Receiver(*EgressRequest, Egress_ReceiverServer) error
}

func RegisterEgressServer(s *grpc.Server, srv EgressServer) {
	s.RegisterService(&_Egress_serviceDesc, srv)
}

func _Egress_Receiver_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EgressRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EgressServer).Receiver(m, &egressReceiverServer{stream})
}

type Egress_ReceiverServer interface {
	Send(*Envelope) error
	grpc.ServerStream
}

type egressReceiverServer struct {
	grpc.ServerStream
}

func (x *egressReceiverServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

var _Egress_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loggregator.v2.Egress",
	HandlerType: (*EgressServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Receiver",
			Handler:       _Egress_Receiver_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("egress.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xbb, 0x0a, 0xdb, 0xcd, 0x54, 0x7b, 0xc8, 0x41, 0xb6, 0x15, 0xa1, 0xe4, 0xd4, 0x8b,
	0x41, 0xd6, 0x6f, 0x20, 0xf8, 0xa7, 0xa0, 0x07, 0x73, 0xf4, 0x52, 0xd6, 0x66, 0x8c, 0x0b, 0xc1,
	0xa9, 0x93, 0x74, 0x3f, 0xbf, 0x98, 0x2c, 0x4a, 0xdb, 0xe3, 0xcc, 0xef, 0xf1, 0x83, 0xf7, 0xe0,
	0x0c, 0x1d, 0x63, 0x08, 0x7a, 0xcb, 0x14, 0x49, 0x4e, 0x3d, 0x39, 0xc7, 0xe8, 0xda, 0x48, 0xac,
	0xfb, 0x66, 0x3e, 0xc5, 0xaf, 0x1e, 0x3d, 0x6d, 0x31, 0x73, 0xf5, 0x06, 0xe7, 0xf7, 0x29, 0x6f,
	0xf0, 0x7b, 0x87, 0x21, 0xca, 0x19, 0x54, 0xe1, 0xb3, 0x65, 0xbb, 0xee, 0x6c, 0x5d, 0x2c, 0x8a,
	0xa5, 0x30, 0xe3, 0x74, 0xaf, 0xac, 0xd4, 0x50, 0x7e, 0x74, 0x3e, 0x22, 0xd7, 0x27, 0x8b, 0x62,
	0x39, 0x69, 0x2e, 0xf4, 0xbe, 0x5c, 0x3f, 0x24, 0x6a, 0x86, 0x94, 0x5a, 0x43, 0x99, 0x3f, 0xf2,
	0x12, 0x44, 0xa0, 0x1d, 0x6f, 0xf0, 0xdf, 0x5a, 0xe5, 0xc7, 0xca, 0xca, 0x6b, 0x38, 0xf5, 0xe4,
	0x06, 0xe7, 0xec, 0xd0, 0xf9, 0x4c, 0x2e, 0x4b, 0x9e, 0x46, 0xe6, 0x37, 0x77, 0x27, 0x60, 0xfc,
	0x82, 0x21, 0xb4, 0x0e, 0xd5, 0x04, 0xc4, 0x1f, 0x6e, 0x5e, 0xa1, 0xcc, 0x4d, 0xe4, 0x23, 0x54,
	0x06, 0x37, 0xd8, 0xf5, 0xc8, 0xf2, 0xea, 0xd0, 0xb7, 0xd7, 0x76, 0x5e, 0x1f, 0xe1, 0x61, 0x1e,
	0x35, 0xba, 0x29, 0xde, 0xcb, 0xb4, 0xd1, 0xed, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0x5e,
	0x62, 0xcc, 0x53, 0x01, 0x00, 0x00,
}
