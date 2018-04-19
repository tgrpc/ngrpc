// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ngrpc.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	ngrpc.proto

It has these top-level messages:
	ListReq
	Lang
	Version
	ListResp
	GetLangReq
*/
package helloworld

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

type ListReq struct {
	Offset int32 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *ListReq) Reset()                    { *m = ListReq{} }
func (m *ListReq) String() string            { return proto.CompactTextString(m) }
func (*ListReq) ProtoMessage()               {}
func (*ListReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ListReq) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListReq) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type Lang struct {
	Name     string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Birthday int64      `protobuf:"varint,2,opt,name=birthday" json:"birthday,omitempty"`
	Versions []*Version `protobuf:"bytes,3,rep,name=versions" json:"versions,omitempty"`
}

func (m *Lang) Reset()                    { *m = Lang{} }
func (m *Lang) String() string            { return proto.CompactTextString(m) }
func (*Lang) ProtoMessage()               {}
func (*Lang) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Lang) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Lang) GetBirthday() int64 {
	if m != nil {
		return m.Birthday
	}
	return 0
}

func (m *Lang) GetVersions() []*Version {
	if m != nil {
		return m.Versions
	}
	return nil
}

type Version struct {
	Id      int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Desc    string `protobuf:"bytes,3,opt,name=desc" json:"desc,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Version) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Version) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Version) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type ListResp struct {
	Langs      []*Lang `protobuf:"bytes,1,rep,name=langs" json:"langs,omitempty"`
	TotalCount int32   `protobuf:"varint,2,opt,name=totalCount" json:"totalCount,omitempty"`
}

func (m *ListResp) Reset()                    { *m = ListResp{} }
func (m *ListResp) String() string            { return proto.CompactTextString(m) }
func (*ListResp) ProtoMessage()               {}
func (*ListResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListResp) GetLangs() []*Lang {
	if m != nil {
		return m.Langs
	}
	return nil
}

func (m *ListResp) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type GetLangReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetLangReq) Reset()                    { *m = GetLangReq{} }
func (m *GetLangReq) String() string            { return proto.CompactTextString(m) }
func (*GetLangReq) ProtoMessage()               {}
func (*GetLangReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetLangReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ListReq)(nil), "helloworld.ListReq")
	proto.RegisterType((*Lang)(nil), "helloworld.Lang")
	proto.RegisterType((*Version)(nil), "helloworld.Version")
	proto.RegisterType((*ListResp)(nil), "helloworld.ListResp")
	proto.RegisterType((*GetLangReq)(nil), "helloworld.GetLangReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LangService service

type LangServiceClient interface {
	List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error)
	GetLang(ctx context.Context, in *GetLangReq, opts ...grpc.CallOption) (*Lang, error)
}

type langServiceClient struct {
	cc *grpc.ClientConn
}

func NewLangServiceClient(cc *grpc.ClientConn) LangServiceClient {
	return &langServiceClient{cc}
}

func (c *langServiceClient) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error) {
	out := new(ListResp)
	err := grpc.Invoke(ctx, "/helloworld.LangService/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *langServiceClient) GetLang(ctx context.Context, in *GetLangReq, opts ...grpc.CallOption) (*Lang, error) {
	out := new(Lang)
	err := grpc.Invoke(ctx, "/helloworld.LangService/GetLang", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LangService service

type LangServiceServer interface {
	List(context.Context, *ListReq) (*ListResp, error)
	GetLang(context.Context, *GetLangReq) (*Lang, error)
}

func RegisterLangServiceServer(s *grpc.Server, srv LangServiceServer) {
	s.RegisterService(&_LangService_serviceDesc, srv)
}

func _LangService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LangServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.LangService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LangServiceServer).List(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LangService_GetLang_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLangReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LangServiceServer).GetLang(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.LangService/GetLang",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LangServiceServer).GetLang(ctx, req.(*GetLangReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _LangService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.LangService",
	HandlerType: (*LangServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _LangService_List_Handler,
		},
		{
			MethodName: "GetLang",
			Handler:    _LangService_GetLang_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ngrpc.proto",
}

func init() { proto.RegisterFile("ngrpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x6e, 0xc2, 0x30,
	0x0c, 0x86, 0x05, 0x05, 0x5a, 0x5c, 0x69, 0x9a, 0x3c, 0x84, 0xa2, 0x1e, 0xa6, 0xaa, 0x87, 0x89,
	0x53, 0xa7, 0xc1, 0x61, 0x0f, 0xb0, 0x03, 0x17, 0x4e, 0x99, 0xb4, 0x7b, 0x69, 0x43, 0x89, 0x14,
	0x92, 0x92, 0x04, 0xa6, 0xbd, 0xfd, 0xd4, 0x34, 0x40, 0x27, 0x76, 0xf3, 0xff, 0xdb, 0x96, 0x3f,
	0xdb, 0x10, 0xcb, 0x5a, 0x37, 0x65, 0xde, 0x68, 0x65, 0x15, 0xc2, 0x9e, 0x09, 0xa1, 0xbe, 0x95,
	0x16, 0x55, 0xf6, 0x0e, 0xe1, 0x86, 0x1b, 0x4b, 0xd9, 0x11, 0xe7, 0x30, 0x51, 0xbb, 0x9d, 0x61,
	0x96, 0x0c, 0xd2, 0xc1, 0x62, 0x4c, 0xbd, 0xc2, 0x19, 0x8c, 0x05, 0x3f, 0x70, 0x4b, 0x86, 0xce,
	0xee, 0x44, 0x56, 0xc3, 0x68, 0x53, 0xc8, 0x1a, 0x11, 0x46, 0xb2, 0x38, 0x30, 0xd7, 0x33, 0xa5,
	0x2e, 0xc6, 0x04, 0xa2, 0x2d, 0xd7, 0x76, 0x5f, 0x15, 0x3f, 0xae, 0x29, 0xa0, 0x57, 0x8d, 0xaf,
	0x10, 0x9d, 0x99, 0x36, 0x5c, 0x49, 0x43, 0x82, 0x34, 0x58, 0xc4, 0xcb, 0xa7, 0xfc, 0xc6, 0x93,
	0x7f, 0x75, 0x39, 0x7a, 0x2d, 0xca, 0xd6, 0x10, 0x7a, 0x13, 0x1f, 0x60, 0xc8, 0x2b, 0x4f, 0x37,
	0xe4, 0x15, 0x12, 0x08, 0x7d, 0x99, 0x1b, 0x33, 0xa5, 0x17, 0xd9, 0x52, 0x55, 0xcc, 0x94, 0x24,
	0xe8, 0xa8, 0xda, 0x38, 0xa3, 0x10, 0x75, 0xab, 0x9a, 0x06, 0x5f, 0x60, 0x2c, 0x0a, 0x59, 0x1b,
	0x32, 0x70, 0x08, 0x8f, 0x7d, 0x84, 0x76, 0x2d, 0xda, 0xa5, 0xf1, 0x19, 0xc0, 0x2a, 0x5b, 0x88,
	0x0f, 0x75, 0x92, 0x97, 0x03, 0xf4, 0x9c, 0x2c, 0x05, 0x58, 0x33, 0xeb, 0x3a, 0xd8, 0xf1, 0xbf,
	0x5b, 0x2c, 0x4f, 0x10, 0xb7, 0xe9, 0x4f, 0xa6, 0xcf, 0xbc, 0x64, 0xf8, 0x06, 0xa3, 0x16, 0x02,
	0xff, 0x2c, 0xed, 0x3f, 0x90, 0xcc, 0xee, 0x4d, 0xd3, 0xe0, 0x0a, 0x42, 0x3f, 0x03, 0xe7, 0xfd,
	0x82, 0xdb, 0xe0, 0xe4, 0x8e, 0x7f, 0x3b, 0x71, 0xaf, 0x5e, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x6c, 0xfb, 0xfd, 0xb8, 0xf9, 0x01, 0x00, 0x00,
}
