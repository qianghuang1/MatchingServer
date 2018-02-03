// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gamers.proto

/*
Package gamers is a generated protocol buffer package.

It is generated from these files:
	gamers.proto

It has these top-level messages:
	Gamer
	Gamers
*/
package gamers

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

// Gammer Model
type Gamer struct {
	Id       string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	DomainId string `protobuf:"bytes,2,opt,name=DomainId" json:"DomainId,omitempty"`
	Rank     int32  `protobuf:"varint,3,opt,name=Rank" json:"Rank,omitempty"`
	Priority int32  `protobuf:"varint,4,opt,name=Priority" json:"Priority,omitempty"`
}

func (m *Gamer) Reset()                    { *m = Gamer{} }
func (m *Gamer) String() string            { return proto.CompactTextString(m) }
func (*Gamer) ProtoMessage()               {}
func (*Gamer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Gamer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Gamer) GetDomainId() string {
	if m != nil {
		return m.DomainId
	}
	return ""
}

func (m *Gamer) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *Gamer) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

// Gammers Model
type Gamers struct {
	Count int32    `protobuf:"varint,1,opt,name=Count" json:"Count,omitempty"`
	Items []*Gamer `protobuf:"bytes,2,rep,name=Items" json:"Items,omitempty"`
}

func (m *Gamers) Reset()                    { *m = Gamers{} }
func (m *Gamers) String() string            { return proto.CompactTextString(m) }
func (*Gamers) ProtoMessage()               {}
func (*Gamers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Gamers) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Gamers) GetItems() []*Gamer {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Gamer)(nil), "gamers.Gamer")
	proto.RegisterType((*Gamers)(nil), "gamers.Gamers")
}

func init() { proto.RegisterFile("gamers.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xcf, 0x4e, 0xc4, 0x20,
	0x10, 0xc6, 0x03, 0xbb, 0x34, 0x3a, 0xab, 0x1e, 0x26, 0x1e, 0x88, 0xa7, 0x66, 0xf7, 0xd2, 0x13,
	0x07, 0x7d, 0x83, 0xd6, 0xc4, 0x70, 0x30, 0x69, 0x78, 0x03, 0x2c, 0x4d, 0x4b, 0x14, 0x50, 0xc0,
	0x83, 0x6f, 0x6f, 0x0a, 0x9b, 0xde, 0xbe, 0xdf, 0x37, 0x7f, 0xbe, 0x19, 0xb8, 0x5b, 0xb4, 0x9b,
	0x63, 0x12, 0xdf, 0x31, 0xe4, 0x80, 0x4d, 0xa5, 0xf3, 0x04, 0xec, 0x6d, 0x53, 0xf8, 0x00, 0x54,
	0x1a, 0x4e, 0x5a, 0xd2, 0xdd, 0x2a, 0x2a, 0x0d, 0x3e, 0xc1, 0xcd, 0x6b, 0x70, 0xda, 0x7a, 0x69,
	0x38, 0x2d, 0xee, 0xce, 0x88, 0x70, 0x54, 0xda, 0x7f, 0xf2, 0x43, 0x4b, 0x3a, 0xa6, 0x8a, 0xde,
	0xfa, 0xc7, 0x68, 0x43, 0xb4, 0xf9, 0x8f, 0x1f, 0x8b, 0xbf, 0xf3, 0x79, 0x80, 0xa6, 0x84, 0x24,
	0x7c, 0x04, 0x36, 0x84, 0x5f, 0x9f, 0x4b, 0x10, 0x53, 0x15, 0xf0, 0x02, 0x4c, 0xe6, 0xd9, 0x25,
	0x4e, 0xdb, 0x43, 0x77, 0x7a, 0xbe, 0x17, 0xd7, 0x53, 0xcb, 0x90, 0xaa, 0xb5, 0xfe, 0x02, 0xb8,
	0xfe, 0x08, 0xa7, 0xf3, 0xb4, 0x5a, 0xbf, 0x08, 0x17, 0xcc, 0xfc, 0x95, 0xfa, 0x53, 0x5d, 0xfc,
	0xbe, 0xd1, 0x48, 0x3e, 0x9a, 0xf2, 0xdd, 0xcb, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x35,
	0x8d, 0x9b, 0xed, 0x00, 0x00, 0x00,
}
