// Code generated by protoc-gen-go.
// source: router.proto
// DO NOT EDIT!

/*
Package gproto is a generated protocol buffer package.

It is generated from these files:
	router.proto

It has these top-level messages:
	Router
	Routers
*/
package gproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Router struct {
	Hostname string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	IP       []byte `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
}

func (m *Router) Reset()                    { *m = Router{} }
func (m *Router) String() string            { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()               {}
func (*Router) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Router) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Router) GetIP() []byte {
	if m != nil {
		return m.IP
	}
	return nil
}

type Routers struct {
	Router []*Router `protobuf:"bytes,1,rep,name=router" json:"router,omitempty"`
}

func (m *Routers) Reset()                    { *m = Routers{} }
func (m *Routers) String() string            { return proto.CompactTextString(m) }
func (*Routers) ProtoMessage()               {}
func (*Routers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Routers) GetRouter() []*Router {
	if m != nil {
		return m.Router
	}
	return nil
}

func init() {
	proto.RegisterType((*Router)(nil), "gproto.Router")
	proto.RegisterType((*Routers)(nil), "gproto.Routers")
}

func init() { proto.RegisterFile("router.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x07, 0xd3, 0x4a, 0x26, 0x5c,
	0x6c, 0x41, 0x60, 0x71, 0x21, 0x29, 0x2e, 0x8e, 0x8c, 0xfc, 0xe2, 0x92, 0xbc, 0xc4, 0xdc, 0x54,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38, 0x5f, 0x88, 0x8f, 0x8b, 0xc9, 0x33, 0x40, 0x82,
	0x49, 0x81, 0x51, 0x83, 0x27, 0x88, 0xc9, 0x33, 0x40, 0xc9, 0x90, 0x8b, 0x1d, 0xa2, 0xab, 0x58,
	0x48, 0x8d, 0x8b, 0x0d, 0x62, 0xb0, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0xb7, 0x11, 0x9f, 0x1e, 0xc4,
	0x64, 0x3d, 0x88, 0x82, 0x20, 0xa8, 0x6c, 0x12, 0x1b, 0x58, 0xd4, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xe8, 0x1d, 0x0f, 0xd7, 0x87, 0x00, 0x00, 0x00,
}