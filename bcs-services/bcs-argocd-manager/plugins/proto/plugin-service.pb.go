// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugin-service.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PluginRenderParam struct {
	Data                 *string  `protobuf:"bytes,1,req,name=data" json:"data,omitempty"`
	Env                  []string `protobuf:"bytes,2,rep,name=env" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginRenderParam) Reset()         { *m = PluginRenderParam{} }
func (m *PluginRenderParam) String() string { return proto.CompactTextString(m) }
func (*PluginRenderParam) ProtoMessage()    {}
func (*PluginRenderParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_88652b3ba4ad62f0, []int{0}
}

func (m *PluginRenderParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginRenderParam.Unmarshal(m, b)
}
func (m *PluginRenderParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginRenderParam.Marshal(b, m, deterministic)
}
func (m *PluginRenderParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginRenderParam.Merge(m, src)
}
func (m *PluginRenderParam) XXX_Size() int {
	return xxx_messageInfo_PluginRenderParam.Size(m)
}
func (m *PluginRenderParam) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginRenderParam.DiscardUnknown(m)
}

var xxx_messageInfo_PluginRenderParam proto.InternalMessageInfo

func (m *PluginRenderParam) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func (m *PluginRenderParam) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

type PluginRenderResp struct {
	Code                 *int32   `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Message              *string  `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	Data                 *string  `protobuf:"bytes,3,req,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginRenderResp) Reset()         { *m = PluginRenderResp{} }
func (m *PluginRenderResp) String() string { return proto.CompactTextString(m) }
func (*PluginRenderResp) ProtoMessage()    {}
func (*PluginRenderResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_88652b3ba4ad62f0, []int{1}
}

func (m *PluginRenderResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginRenderResp.Unmarshal(m, b)
}
func (m *PluginRenderResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginRenderResp.Marshal(b, m, deterministic)
}
func (m *PluginRenderResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginRenderResp.Merge(m, src)
}
func (m *PluginRenderResp) XXX_Size() int {
	return xxx_messageInfo_PluginRenderResp.Size(m)
}
func (m *PluginRenderResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginRenderResp.DiscardUnknown(m)
}

var xxx_messageInfo_PluginRenderResp proto.InternalMessageInfo

func (m *PluginRenderResp) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *PluginRenderResp) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *PluginRenderResp) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*PluginRenderParam)(nil), "proto.PluginRenderParam")
	proto.RegisterType((*PluginRenderResp)(nil), "proto.PluginRenderResp")
}

func init() { proto.RegisterFile("plugin-service.proto", fileDescriptor_88652b3ba4ad62f0) }

var fileDescriptor_88652b3ba4ad62f0 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0xc8, 0x29, 0x4d,
	0xcf, 0xcc, 0xd3, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x05, 0x53, 0x4a, 0x96, 0x5c, 0x82, 0x01, 0x60, 0xe9, 0xa0, 0xd4, 0xbc, 0x94, 0xd4,
	0xa2, 0x80, 0xc4, 0xa2, 0xc4, 0x5c, 0x21, 0x21, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0x09, 0x46,
	0x05, 0x26, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x80, 0x8b, 0x39, 0x35, 0xaf, 0x4c, 0x82, 0x49,
	0x81, 0x59, 0x83, 0x33, 0x08, 0xc4, 0x54, 0x0a, 0xe1, 0x12, 0x40, 0xd6, 0x1a, 0x94, 0x5a, 0x5c,
	0x00, 0xd2, 0x99, 0x9c, 0x9f, 0x92, 0x0a, 0xd6, 0xc9, 0x1a, 0x04, 0x66, 0x0b, 0x49, 0x70, 0xb1,
	0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x81, 0x0d, 0x84, 0x71, 0xe1, 0xf6, 0x30,
	0x23, 0xec, 0x71, 0x62, 0x8f, 0x82, 0xb8, 0x0c, 0x10, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xe5, 0x7d,
	0x40, 0xb7, 0x00, 0x00, 0x00,
}
