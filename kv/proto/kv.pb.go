// Code generated by protoc-gen-go.
// source: github.com/micro/go-os/kv/proto/kv.proto
// DO NOT EDIT!

/*
Package go_micro_os_kv is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/go-os/kv/proto/kv.proto

It has these top-level messages:
	Item
	GetRequest
	GetResponse
	PutRequest
	PutResponse
	DelRequest
	DelResponse
*/
package go_micro_os_kv

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Item struct {
	Key        string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value      []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Expiration int64  `protobuf:"varint,3,opt,name=expiration" json:"expiration,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetResponse struct {
	Item *Item `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetResponse) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type PutRequest struct {
	Item *Item `protobuf:"bytes,1,opt,name=item" json:"item,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PutRequest) GetItem() *Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type PutResponse struct {
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type DelRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *DelRequest) Reset()                    { *m = DelRequest{} }
func (m *DelRequest) String() string            { return proto.CompactTextString(m) }
func (*DelRequest) ProtoMessage()               {}
func (*DelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type DelResponse struct {
}

func (m *DelResponse) Reset()                    { *m = DelResponse{} }
func (m *DelResponse) String() string            { return proto.CompactTextString(m) }
func (*DelResponse) ProtoMessage()               {}
func (*DelResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*Item)(nil), "go.micro.os.kv.Item")
	proto.RegisterType((*GetRequest)(nil), "go.micro.os.kv.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "go.micro.os.kv.GetResponse")
	proto.RegisterType((*PutRequest)(nil), "go.micro.os.kv.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "go.micro.os.kv.PutResponse")
	proto.RegisterType((*DelRequest)(nil), "go.micro.os.kv.DelRequest")
	proto.RegisterType((*DelResponse)(nil), "go.micro.os.kv.DelResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for KV service

type KVClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...client.CallOption) (*PutResponse, error)
	Del(ctx context.Context, in *DelRequest, opts ...client.CallOption) (*DelResponse, error)
}

type kVClient struct {
	c           client.Client
	serviceName string
}

func NewKVClient(serviceName string, c client.Client) KVClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.os.kv"
	}
	return &kVClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *kVClient) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.serviceName, "KV.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Put(ctx context.Context, in *PutRequest, opts ...client.CallOption) (*PutResponse, error) {
	req := c.c.NewRequest(c.serviceName, "KV.Put", in)
	out := new(PutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Del(ctx context.Context, in *DelRequest, opts ...client.CallOption) (*DelResponse, error) {
	req := c.c.NewRequest(c.serviceName, "KV.Del", in)
	out := new(DelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KV service

type KVHandler interface {
	Get(context.Context, *GetRequest, *GetResponse) error
	Put(context.Context, *PutRequest, *PutResponse) error
	Del(context.Context, *DelRequest, *DelResponse) error
}

func RegisterKVHandler(s server.Server, hdlr KVHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&KV{hdlr}, opts...))
}

type KV struct {
	KVHandler
}

func (h *KV) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.KVHandler.Get(ctx, in, out)
}

func (h *KV) Put(ctx context.Context, in *PutRequest, out *PutResponse) error {
	return h.KVHandler.Put(ctx, in, out)
}

func (h *KV) Del(ctx context.Context, in *DelRequest, out *DelResponse) error {
	return h.KVHandler.Del(ctx, in, out)
}

func init() { proto.RegisterFile("github.com/micro/go-os/kv/proto/kv.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x9b, 0xa6, 0x0a, 0x4e, 0x54, 0x64, 0xe9, 0x21, 0x54, 0x10, 0xd9, 0x53, 0x2e, 0x6e,
	0xa0, 0x82, 0x5e, 0x3d, 0x08, 0x22, 0x82, 0xc8, 0x1e, 0xbc, 0xb7, 0x65, 0x89, 0x4b, 0x9a, 0x4c,
	0xcc, 0x6e, 0x82, 0xfe, 0x45, 0x7f, 0x95, 0x9b, 0x89, 0x98, 0xe8, 0x2a, 0xf4, 0x36, 0xd9, 0xf7,
	0xe6, 0xcd, 0xf7, 0x20, 0x90, 0x64, 0xda, 0xbe, 0x34, 0x6b, 0xb1, 0xc1, 0x22, 0x2d, 0xf4, 0xa6,
	0xc6, 0x34, 0xc3, 0x0b, 0x34, 0x69, 0xde, 0xa6, 0x55, 0x8d, 0x16, 0xdd, 0x20, 0x68, 0x60, 0xc7,
	0x19, 0x0a, 0x72, 0x08, 0x34, 0x22, 0x6f, 0xf9, 0x23, 0xcc, 0xee, 0xad, 0x2a, 0xd8, 0x09, 0x84,
	0xb9, 0x7a, 0x8f, 0x83, 0xf3, 0x20, 0x39, 0x90, 0xdd, 0xc8, 0xe6, 0xb0, 0xd7, 0xae, 0xb6, 0x8d,
	0x8a, 0xa7, 0xee, 0xed, 0x50, 0xf6, 0x1f, 0xec, 0x0c, 0x40, 0xbd, 0x55, 0xba, 0x5e, 0x59, 0x8d,
	0x65, 0x1c, 0x3a, 0x29, 0x94, 0xa3, 0x17, 0xee, 0xf4, 0x3b, 0x65, 0xa5, 0x7a, 0x6d, 0x94, 0xb1,
	0x7e, 0x2a, 0xbf, 0x86, 0x88, 0x74, 0x53, 0x61, 0x69, 0x14, 0x4b, 0x60, 0xa6, 0xdd, 0x79, 0x72,
	0x44, 0xcb, 0xb9, 0xf8, 0x49, 0x27, 0x3a, 0x34, 0x49, 0x0e, 0x7e, 0x05, 0xf0, 0xd4, 0x7c, 0x07,
	0xef, 0xbe, 0x77, 0x04, 0x11, 0xed, 0xf5, 0x07, 0x3b, 0xbe, 0x5b, 0xb5, 0xfd, 0x9f, 0xcf, 0xd9,
	0x49, 0xef, 0xed, 0xcb, 0x8f, 0x00, 0xa6, 0x0f, 0xcf, 0xec, 0x06, 0x42, 0x47, 0xcd, 0x16, 0xbf,
	0xef, 0x0c, 0x55, 0x17, 0xa7, 0x7f, 0x6a, 0x5f, 0x57, 0x27, 0x5d, 0x82, 0xc3, 0xf0, 0x13, 0x86,
	0x4e, 0x7e, 0xc2, 0x98, 0x9b, 0x12, 0x1c, 0x99, 0x9f, 0x30, 0xd4, 0xf1, 0x13, 0x46, 0x55, 0xf8,
	0x64, 0xbd, 0x4f, 0xbf, 0xc0, 0xe5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x1b, 0x8c, 0x55,
	0x2e, 0x02, 0x00, 0x00,
}
