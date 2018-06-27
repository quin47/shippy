// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

package go_micro_srv_consignment

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

type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Consignment          []*Container `protobuf:"bytes,4,rep,name=consignment,proto3" json:"consignment,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_43898a1c8b7346fb, []int{0}
}
func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (dst *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(dst, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetConsignment() []*Container {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_43898a1c8b7346fb, []int{1}
}
func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (dst *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(dst, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consigments          []*Consignment `protobuf:"bytes,3,rep,name=consigments,proto3" json:"consigments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_43898a1c8b7346fb, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsigments() []*Consignment {
	if m != nil {
		return m.Consigments
	}
	return nil
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_43898a1c8b7346fb, []int{3}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Consignment)(nil), "go.micro.srv.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "go.micro.srv.consignment.Container")
	proto.RegisterType((*Response)(nil), "go.micro.srv.consignment.Response")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.consignment.GetRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShippingServiceClient is the client API for ShippingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error)
	GetAllConsignment(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	cc *grpc.ClientConn
}

func NewShippingServiceClient(cc *grpc.ClientConn) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.consignment.ShippingService/CreateConsignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetAllConsignment(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.consignment.ShippingService/GetAllConsignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServiceServer is the server API for ShippingService service.
type ShippingServiceServer interface {
	CreateConsignment(context.Context, *Consignment) (*Response, error)
	GetAllConsignment(context.Context, *GetRequest) (*Response, error)
}

func RegisterShippingServiceServer(s *grpc.Server, srv ShippingServiceServer) {
	s.RegisterService(&_ShippingService_serviceDesc, srv)
}

func _ShippingService_CreateConsignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Consignment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.consignment.ShippingService/CreateConsignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, req.(*Consignment))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShippingService_GetAllConsignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).GetAllConsignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.consignment.ShippingService/GetAllConsignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).GetAllConsignment(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShippingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.consignment.ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConsignment",
			Handler:    _ShippingService_CreateConsignment_Handler,
		},
		{
			MethodName: "GetAllConsignment",
			Handler:    _ShippingService_GetAllConsignment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consignment/consignment.proto",
}

func init() {
	proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor_consignment_43898a1c8b7346fb)
}

var fileDescriptor_consignment_43898a1c8b7346fb = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0xb6, 0xfc, 0x77, 0x6a, 0x34, 0xec, 0x41, 0x37, 0x7a, 0xb0, 0x29, 0x9a, 0x70, 0xaa, 0x09,
	0x3e, 0x81, 0x21, 0x86, 0x70, 0x2d, 0x0f, 0x80, 0xd8, 0x4e, 0xca, 0x24, 0x74, 0xb7, 0xee, 0x2e,
	0xf8, 0x66, 0x5e, 0x7c, 0x0d, 0x1f, 0xc8, 0x74, 0x4b, 0xa1, 0xc4, 0x60, 0xf0, 0xb6, 0xdf, 0x7c,
	0xfb, 0xcd, 0xf7, 0xcd, 0x64, 0x60, 0x90, 0x2b, 0x69, 0xe4, 0x63, 0x2c, 0x85, 0xa6, 0x54, 0x64,
	0x28, 0x4c, 0xfd, 0x1d, 0x5a, 0x96, 0xf1, 0x54, 0x86, 0x19, 0xc5, 0x4a, 0x86, 0x5a, 0x6d, 0xc2,
	0x1a, 0x1f, 0x7c, 0x39, 0xe0, 0x8d, 0xf7, 0x98, 0x5d, 0x40, 0x83, 0x12, 0xee, 0xf8, 0xce, 0xd0,
	0x8d, 0x1a, 0x94, 0x30, 0x1f, 0xbc, 0x04, 0x75, 0xac, 0x28, 0x37, 0x24, 0x05, 0x6f, 0x58, 0xa2,
	0x5e, 0x62, 0x57, 0xd0, 0xf9, 0x40, 0x4a, 0x97, 0x86, 0x37, 0x7d, 0x67, 0xd8, 0x8e, 0xb6, 0x88,
	0xbd, 0x80, 0x57, 0x33, 0xe2, 0x2d, 0xbf, 0x39, 0xf4, 0x46, 0x83, 0xf0, 0x58, 0x92, 0x70, 0x2c,
	0x85, 0x59, 0x90, 0x40, 0x15, 0xd5, 0x75, 0xec, 0x16, 0xdc, 0x0d, 0x6a, 0x8d, 0xab, 0x39, 0x25,
	0xbc, 0x6d, 0xed, 0x7b, 0x65, 0x61, 0x9a, 0x04, 0x19, 0xb8, 0x3b, 0xd9, 0xaf, 0xe8, 0x77, 0xe0,
	0xc5, 0x6b, 0x6d, 0x64, 0x86, 0xaa, 0xd0, 0x96, 0xd1, 0xa1, 0x2a, 0x4d, 0x93, 0x22, 0xb9, 0x54,
	0x94, 0x92, 0xb0, 0xc9, 0xdd, 0x68, 0x8b, 0xd8, 0x35, 0x74, 0xd7, 0xba, 0x14, 0xb5, 0x4a, 0xa2,
	0x80, 0xd3, 0x24, 0xf8, 0x74, 0xa0, 0x17, 0xa1, 0xce, 0xa5, 0xd0, 0xc8, 0x38, 0x74, 0x63, 0x85,
	0x0b, 0x83, 0xa5, 0x67, 0x2f, 0xaa, 0x20, 0x9b, 0x1c, 0x4e, 0x5e, 0x18, 0x7b, 0xa3, 0x87, 0x3f,
	0x27, 0xaf, 0xde, 0x87, 0xb3, 0xef, 0x1a, 0x15, 0x48, 0xf3, 0xa6, 0x5d, 0xe1, 0xff, 0x1a, 0x59,
	0x65, 0x70, 0x0e, 0x30, 0x41, 0x13, 0xe1, 0xfb, 0x1a, 0xb5, 0x19, 0x7d, 0x3b, 0x70, 0x39, 0x5b,
	0x52, 0x9e, 0x93, 0x48, 0x67, 0xa8, 0x36, 0x14, 0x23, 0x7b, 0x85, 0xfe, 0xd8, 0xc6, 0xaf, 0x1f,
	0xc3, 0x69, 0x56, 0x37, 0xc1, 0xf1, 0x6f, 0xd5, 0xb6, 0x82, 0x33, 0x36, 0x87, 0xfe, 0x04, 0xcd,
	0xf3, 0x6a, 0x55, 0x77, 0xb8, 0x3f, 0x2e, 0xdd, 0x07, 0x3e, 0xcd, 0xe0, 0xad, 0x63, 0x6f, 0xfd,
	0xe9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x73, 0x57, 0x14, 0x2e, 0x12, 0x03, 0x00, 0x00,
}