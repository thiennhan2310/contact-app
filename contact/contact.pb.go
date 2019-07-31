// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contact/contact.proto

package contact

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type AddContactReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddContactReq) Reset()         { *m = AddContactReq{} }
func (m *AddContactReq) String() string { return proto.CompactTextString(m) }
func (*AddContactReq) ProtoMessage()    {}
func (*AddContactReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f68fe29e5220f006, []int{0}
}

func (m *AddContactReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddContactReq.Unmarshal(m, b)
}
func (m *AddContactReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddContactReq.Marshal(b, m, deterministic)
}
func (m *AddContactReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddContactReq.Merge(m, src)
}
func (m *AddContactReq) XXX_Size() int {
	return xxx_messageInfo_AddContactReq.Size(m)
}
func (m *AddContactReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddContactReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddContactReq proto.InternalMessageInfo

func (m *AddContactReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddContactReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type AddContactRep struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddContactRep) Reset()         { *m = AddContactRep{} }
func (m *AddContactRep) String() string { return proto.CompactTextString(m) }
func (*AddContactRep) ProtoMessage()    {}
func (*AddContactRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_f68fe29e5220f006, []int{1}
}

func (m *AddContactRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddContactRep.Unmarshal(m, b)
}
func (m *AddContactRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddContactRep.Marshal(b, m, deterministic)
}
func (m *AddContactRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddContactRep.Merge(m, src)
}
func (m *AddContactRep) XXX_Size() int {
	return xxx_messageInfo_AddContactRep.Size(m)
}
func (m *AddContactRep) XXX_DiscardUnknown() {
	xxx_messageInfo_AddContactRep.DiscardUnknown(m)
}

var xxx_messageInfo_AddContactRep proto.InternalMessageInfo

func (m *AddContactRep) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetContactReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetContactReq) Reset()         { *m = GetContactReq{} }
func (m *GetContactReq) String() string { return proto.CompactTextString(m) }
func (*GetContactReq) ProtoMessage()    {}
func (*GetContactReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f68fe29e5220f006, []int{2}
}

func (m *GetContactReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContactReq.Unmarshal(m, b)
}
func (m *GetContactReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContactReq.Marshal(b, m, deterministic)
}
func (m *GetContactReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContactReq.Merge(m, src)
}
func (m *GetContactReq) XXX_Size() int {
	return xxx_messageInfo_GetContactReq.Size(m)
}
func (m *GetContactReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContactReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetContactReq proto.InternalMessageInfo

func (m *GetContactReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetContactRep struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetContactRep) Reset()         { *m = GetContactRep{} }
func (m *GetContactRep) String() string { return proto.CompactTextString(m) }
func (*GetContactRep) ProtoMessage()    {}
func (*GetContactRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_f68fe29e5220f006, []int{3}
}

func (m *GetContactRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetContactRep.Unmarshal(m, b)
}
func (m *GetContactRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetContactRep.Marshal(b, m, deterministic)
}
func (m *GetContactRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetContactRep.Merge(m, src)
}
func (m *GetContactRep) XXX_Size() int {
	return xxx_messageInfo_GetContactRep.Size(m)
}
func (m *GetContactRep) XXX_DiscardUnknown() {
	xxx_messageInfo_GetContactRep.DiscardUnknown(m)
}

var xxx_messageInfo_GetContactRep proto.InternalMessageInfo

func (m *GetContactRep) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetContactRep) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetContactRep) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func init() {
	proto.RegisterType((*AddContactReq)(nil), "AddContactReq")
	proto.RegisterType((*AddContactRep)(nil), "AddContactRep")
	proto.RegisterType((*GetContactReq)(nil), "GetContactReq")
	proto.RegisterType((*GetContactRep)(nil), "GetContactRep")
}

func init() { proto.RegisterFile("contact/contact.proto", fileDescriptor_f68fe29e5220f006) }

var fileDescriptor_f68fe29e5220f006 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xcf, 0x2b,
	0x49, 0x4c, 0x2e, 0xd1, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x4a, 0x96, 0x5c, 0xbc,
	0x8e, 0x29, 0x29, 0xce, 0x10, 0xb1, 0xa0, 0xd4, 0x42, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc,
	0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x84, 0x8b, 0xb5, 0x20, 0x23,
	0x3f, 0x2f, 0x55, 0x82, 0x09, 0x2c, 0x08, 0xe1, 0x28, 0x69, 0xa2, 0x6a, 0x2d, 0x10, 0x92, 0xe0,
	0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x87, 0xe9, 0x86, 0x71, 0x95, 0xe4, 0xb9, 0x78, 0xdd,
	0x53, 0x4b, 0x90, 0x6c, 0xe1, 0xe3, 0x62, 0xca, 0x4c, 0x01, 0xab, 0x62, 0x0d, 0x62, 0xca, 0x4c,
	0x51, 0xf2, 0x44, 0x55, 0x50, 0x80, 0xa4, 0x80, 0x13, 0xa4, 0x00, 0xee, 0x2c, 0x26, 0x6c, 0xce,
	0x62, 0x46, 0x72, 0x96, 0x51, 0x2a, 0x17, 0x3b, 0xd4, 0x1c, 0x21, 0x1d, 0x2e, 0x2e, 0x84, 0x0b,
	0x85, 0xf8, 0xf4, 0x50, 0x7c, 0x2a, 0x85, 0xca, 0x2f, 0x00, 0xa9, 0x46, 0xb8, 0x41, 0x88, 0x4f,
	0x0f, 0xc5, 0xc5, 0x52, 0xa8, 0xfc, 0x82, 0x24, 0x36, 0x70, 0xf8, 0x19, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x58, 0x4f, 0xe5, 0x04, 0x58, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContactClient is the client API for Contact service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContactClient interface {
	AddContact(ctx context.Context, in *AddContactReq, opts ...grpc.CallOption) (*AddContactRep, error)
	GetContact(ctx context.Context, in *GetContactReq, opts ...grpc.CallOption) (*GetContactRep, error)
}

type contactClient struct {
	cc *grpc.ClientConn
}

func NewContactClient(cc *grpc.ClientConn) ContactClient {
	return &contactClient{cc}
}

func (c *contactClient) AddContact(ctx context.Context, in *AddContactReq, opts ...grpc.CallOption) (*AddContactRep, error) {
	out := new(AddContactRep)
	err := c.cc.Invoke(ctx, "/Contact/AddContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactClient) GetContact(ctx context.Context, in *GetContactReq, opts ...grpc.CallOption) (*GetContactRep, error) {
	out := new(GetContactRep)
	err := c.cc.Invoke(ctx, "/Contact/GetContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactServer is the server API for Contact service.
type ContactServer interface {
	AddContact(context.Context, *AddContactReq) (*AddContactRep, error)
	GetContact(context.Context, *GetContactReq) (*GetContactRep, error)
}

func RegisterContactServer(s *grpc.Server, srv ContactServer) {
	s.RegisterService(&_Contact_serviceDesc, srv)
}

func _Contact_AddContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServer).AddContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Contact/AddContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServer).AddContact(ctx, req.(*AddContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contact_GetContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServer).GetContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Contact/GetContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServer).GetContact(ctx, req.(*GetContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Contact_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Contact",
	HandlerType: (*ContactServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddContact",
			Handler:    _Contact_AddContact_Handler,
		},
		{
			MethodName: "GetContact",
			Handler:    _Contact_GetContact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contact/contact.proto",
}