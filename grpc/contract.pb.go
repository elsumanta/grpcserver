// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contract.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The request message containing the user's name.
type RegRequest struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegRequest) Reset()         { *m = RegRequest{} }
func (m *RegRequest) String() string { return proto.CompactTextString(m) }
func (*RegRequest) ProtoMessage()    {}
func (*RegRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d19debeba7dea55a, []int{0}
}

func (m *RegRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegRequest.Unmarshal(m, b)
}
func (m *RegRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegRequest.Marshal(b, m, deterministic)
}
func (m *RegRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegRequest.Merge(m, src)
}
func (m *RegRequest) XXX_Size() int {
	return xxx_messageInfo_RegRequest.Size(m)
}
func (m *RegRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegRequest proto.InternalMessageInfo

func (m *RegRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *RegRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *RegRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RegRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

// The response message containing the greetings
type RegResponse struct {
	HttpStatus           int32    `protobuf:"varint,1,opt,name=http_status,json=httpStatus,proto3" json:"http_status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegResponse) Reset()         { *m = RegResponse{} }
func (m *RegResponse) String() string { return proto.CompactTextString(m) }
func (*RegResponse) ProtoMessage()    {}
func (*RegResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d19debeba7dea55a, []int{1}
}

func (m *RegResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegResponse.Unmarshal(m, b)
}
func (m *RegResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegResponse.Marshal(b, m, deterministic)
}
func (m *RegResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegResponse.Merge(m, src)
}
func (m *RegResponse) XXX_Size() int {
	return xxx_messageInfo_RegResponse.Size(m)
}
func (m *RegResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegResponse proto.InternalMessageInfo

func (m *RegResponse) GetHttpStatus() int32 {
	if m != nil {
		return m.HttpStatus
	}
	return 0
}

func (m *RegResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RegRequest)(nil), "grpc.RegRequest")
	proto.RegisterType((*RegResponse)(nil), "grpc.RegResponse")
}

func init() { proto.RegisterFile("contract.proto", fileDescriptor_d19debeba7dea55a) }

var fileDescriptor_d19debeba7dea55a = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3f, 0x4b, 0xf4, 0x40,
	0x10, 0xc6, 0xdf, 0xbc, 0x9e, 0xde, 0x65, 0x22, 0xfe, 0xd9, 0x2a, 0x78, 0x88, 0x9a, 0xca, 0x2a,
	0xa0, 0x82, 0x95, 0xd5, 0x35, 0x5e, 0x75, 0x1c, 0xb1, 0x10, 0x6c, 0xc2, 0x5e, 0x32, 0x26, 0x81,
	0x64, 0x77, 0xdd, 0x99, 0xa0, 0x5f, 0xc0, 0xef, 0x2d, 0x3b, 0x47, 0xd0, 0x6e, 0xe6, 0xf7, 0x2b,
	0xe6, 0x99, 0x07, 0x4e, 0x2a, 0x6b, 0xd8, 0xeb, 0x8a, 0x73, 0xe7, 0x2d, 0x5b, 0x35, 0x6b, 0xbc,
	0xab, 0xb2, 0xef, 0x08, 0xa0, 0xc0, 0xa6, 0xc0, 0x8f, 0x11, 0x89, 0xd5, 0x25, 0xc0, 0x7b, 0xe7,
	0x89, 0x4b, 0xa3, 0x07, 0x4c, 0xa3, 0xeb, 0xe8, 0x36, 0x2e, 0x62, 0x21, 0x1b, 0x3d, 0xa0, 0x5a,
	0x42, 0xdc, 0xeb, 0xc9, 0xfe, 0x17, 0xbb, 0x08, 0x40, 0x64, 0x0a, 0x73, 0x5d, 0xd7, 0x1e, 0x89,
	0xd2, 0x03, 0x51, 0xd3, 0xaa, 0x6e, 0xe0, 0xd8, 0xb5, 0xd6, 0x60, 0x69, 0xc6, 0x61, 0x87, 0x3e,
	0x9d, 0x89, 0x4e, 0x84, 0x6d, 0x04, 0x65, 0x6b, 0x48, 0x24, 0x06, 0x39, 0x6b, 0x08, 0xd5, 0x15,
	0x24, 0x2d, 0xb3, 0x2b, 0x89, 0x35, 0x8f, 0x24, 0x41, 0x0e, 0x0b, 0x08, 0xe8, 0x45, 0x48, 0x38,
	0x36, 0x20, 0x91, 0x6e, 0xa6, 0x1c, 0xd3, 0x7a, 0xff, 0x04, 0xf3, 0x67, 0x8f, 0xc8, 0xe8, 0xd5,
	0x1d, 0x2c, 0x0a, 0x6c, 0x3a, 0x0a, 0xf3, 0x59, 0x1e, 0xfe, 0xcd, 0x7f, 0x7f, 0xbd, 0x38, 0xff,
	0x43, 0xf6, 0x67, 0xb3, 0x7f, 0xab, 0x47, 0x58, 0x76, 0x76, 0x2f, 0xf0, 0x4b, 0x0f, 0xae, 0x47,
	0xca, 0x5b, 0xec, 0x7b, 0xfb, 0x69, 0x7d, 0x5f, 0xaf, 0x4e, 0xd7, 0x61, 0x7e, 0x0d, 0xf3, 0x36,
	0xb4, 0xb8, 0x8d, 0xde, 0xa4, 0xc7, 0xdd, 0x91, 0x94, 0xfa, 0xf0, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x64, 0x88, 0xe6, 0xb6, 0x66, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	Register(ctx context.Context, in *RegRequest, opts ...grpc.CallOption) (*RegResponse, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Register(ctx context.Context, in *RegRequest, opts ...grpc.CallOption) (*RegResponse, error) {
	out := new(RegResponse)
	err := c.cc.Invoke(ctx, "/grpc.Greeter/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	Register(context.Context, *RegRequest) (*RegResponse, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) Register(ctx context.Context, req *RegRequest) (*RegResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Greeter/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Register(ctx, req.(*RegRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Greeter_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract.proto",
}
