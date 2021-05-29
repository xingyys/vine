// Code generated by proto-gen-gogo. DO NOT EDIT.
// source: github.com/lack-io/vine/proto/services/network/dns/dns.proto

package dns

import (
	context "context"
	ebinary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _ = ebinary.BigEndian

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Define a message to register a DNS record
type Record struct {
	// e.g. example.com
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// e.g. 127.0.0.1
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// e.g. A, AAAA, TXT, SRV, ...
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// MX and SRV records have priority
	Priority uint32 `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	// TTL
	Ttl uint32 `protobuf:"varint,5,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_2425ca65c0dfa28b, []int{0}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Record.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return m.Size()
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

type AdvertiseRequest struct {
	// Send an arbitrary number of records to advertise
	Records []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (m *AdvertiseRequest) Reset()         { *m = AdvertiseRequest{} }
func (m *AdvertiseRequest) String() string { return proto.CompactTextString(m) }
func (*AdvertiseRequest) ProtoMessage()    {}
func (*AdvertiseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2425ca65c0dfa28b, []int{1}
}
func (m *AdvertiseRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AdvertiseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AdvertiseRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AdvertiseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdvertiseRequest.Merge(m, src)
}
func (m *AdvertiseRequest) XXX_Size() int {
	return m.Size()
}
func (m *AdvertiseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AdvertiseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AdvertiseRequest proto.InternalMessageInfo

type AdvertiseResponse struct {
}

func (m *AdvertiseResponse) Reset()         { *m = AdvertiseResponse{} }
func (m *AdvertiseResponse) String() string { return proto.CompactTextString(m) }
func (*AdvertiseResponse) ProtoMessage()    {}
func (*AdvertiseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2425ca65c0dfa28b, []int{2}
}
func (m *AdvertiseResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AdvertiseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AdvertiseResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AdvertiseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdvertiseResponse.Merge(m, src)
}
func (m *AdvertiseResponse) XXX_Size() int {
	return m.Size()
}
func (m *AdvertiseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AdvertiseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AdvertiseResponse proto.InternalMessageInfo

type RemoveRequest struct {
	// Send an arbitrary number of records to remove
	Records []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2425ca65c0dfa28b, []int{3}
}
func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return m.Size()
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

type RemoveResponse struct {
}

func (m *RemoveResponse) Reset()         { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()    {}
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2425ca65c0dfa28b, []int{4}
}
func (m *RemoveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveResponse.Merge(m, src)
}
func (m *RemoveResponse) XXX_Size() int {
	return m.Size()
}
func (m *RemoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveResponse proto.InternalMessageInfo

type ResolveRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *ResolveRequest) Reset()         { *m = ResolveRequest{} }
func (m *ResolveRequest) String() string { return proto.CompactTextString(m) }
func (*ResolveRequest) ProtoMessage()    {}
func (*ResolveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2425ca65c0dfa28b, []int{5}
}
func (m *ResolveRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResolveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResolveRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResolveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveRequest.Merge(m, src)
}
func (m *ResolveRequest) XXX_Size() int {
	return m.Size()
}
func (m *ResolveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveRequest proto.InternalMessageInfo

type ResolveResponse struct {
	// Return any matching records
	Records []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (m *ResolveResponse) Reset()         { *m = ResolveResponse{} }
func (m *ResolveResponse) String() string { return proto.CompactTextString(m) }
func (*ResolveResponse) ProtoMessage()    {}
func (*ResolveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2425ca65c0dfa28b, []int{6}
}
func (m *ResolveResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResolveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResolveResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResolveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveResponse.Merge(m, src)
}
func (m *ResolveResponse) XXX_Size() int {
	return m.Size()
}
func (m *ResolveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Record)(nil), "dns.Record")
	proto.RegisterType((*AdvertiseRequest)(nil), "dns.AdvertiseRequest")
	proto.RegisterType((*AdvertiseResponse)(nil), "dns.AdvertiseResponse")
	proto.RegisterType((*RemoveRequest)(nil), "dns.RemoveRequest")
	proto.RegisterType((*RemoveResponse)(nil), "dns.RemoveResponse")
	proto.RegisterType((*ResolveRequest)(nil), "dns.ResolveRequest")
	proto.RegisterType((*ResolveResponse)(nil), "dns.ResolveResponse")
}

func init() {
	proto.RegisterFile("github.com/lack-io/vine/proto/services/network/dns/dns.proto", fileDescriptor_2425ca65c0dfa28b)
}

var fileDescriptor_2425ca65c0dfa28b = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0xed, 0x50, 0x7e, 0x3e, 0x2e, 0xe1, 0xfb, 0xf8, 0x06, 0x34, 0x4d, 0x17, 0x0d, 0x69, 0x62,
	0xc2, 0x46, 0x1a, 0xd1, 0x10, 0xfc, 0xd9, 0x68, 0x7c, 0x82, 0xc6, 0x95, 0x3b, 0x68, 0x27, 0xda,
	0x50, 0x66, 0xea, 0xcc, 0x50, 0xc3, 0x5b, 0xf8, 0x1c, 0x3e, 0x09, 0x4b, 0x96, 0x2e, 0x15, 0x5e,
	0xc4, 0x4c, 0x7f, 0xf8, 0x8b, 0x0b, 0x59, 0x34, 0xb9, 0x73, 0xee, 0x3d, 0x67, 0xce, 0x3d, 0x1d,
	0xb8, 0x79, 0x0a, 0xe4, 0xf3, 0x74, 0xd4, 0xf5, 0xd8, 0xc4, 0x09, 0x87, 0xde, 0xf8, 0x34, 0x60,
	0x4e, 0x1c, 0x50, 0xe2, 0x44, 0x9c, 0x49, 0xe6, 0x08, 0xc2, 0xe3, 0xc0, 0x23, 0xc2, 0xa1, 0x44,
	0xbe, 0x32, 0x3e, 0x76, 0x7c, 0x2a, 0xd4, 0xd7, 0x4d, 0xda, 0x58, 0xf7, 0xa9, 0xb0, 0x25, 0x94,
	0x5d, 0xe2, 0x31, 0xee, 0x63, 0x0c, 0x45, 0x3a, 0x9c, 0x10, 0x03, 0xb5, 0x51, 0xa7, 0xea, 0x26,
	0x35, 0x6e, 0x41, 0x29, 0x1e, 0x86, 0x53, 0x62, 0x14, 0x12, 0x30, 0x3d, 0xa8, 0x49, 0x39, 0x8b,
	0x88, 0xa1, 0xa7, 0x93, 0xaa, 0xc6, 0x26, 0xfc, 0x89, 0x78, 0xc0, 0x78, 0x20, 0x67, 0x46, 0xb1,
	0x8d, 0x3a, 0x75, 0x77, 0x7d, 0xc6, 0x0d, 0xd0, 0xa5, 0x0c, 0x8d, 0x52, 0x02, 0xab, 0xd2, 0xbe,
	0x84, 0xc6, 0xad, 0x1f, 0x13, 0x2e, 0x03, 0x41, 0x5c, 0xf2, 0x32, 0x25, 0x42, 0xe2, 0x13, 0xa8,
	0xf0, 0xc4, 0x89, 0x30, 0x50, 0x5b, 0xef, 0xd4, 0x7a, 0xb5, 0xae, 0xf2, 0x9a, 0xba, 0x73, 0xf3,
	0x9e, 0xdd, 0x84, 0xff, 0x5b, 0x54, 0x11, 0x31, 0x2a, 0x88, 0xdd, 0x87, 0xba, 0x4b, 0x26, 0x2c,
	0x3e, 0x54, 0xac, 0x01, 0x7f, 0x73, 0x5e, 0xa6, 0x34, 0x50, 0x88, 0x60, 0xe1, 0x46, 0xea, 0xa7,
	0x5c, 0xf2, 0x04, 0x0a, 0x9b, 0x04, 0xec, 0x01, 0xfc, 0x5b, 0x33, 0x53, 0xb1, 0x5f, 0xba, 0xe8,
	0xbd, 0x23, 0xd0, 0xef, 0xa9, 0xc0, 0x57, 0x50, 0x5d, 0xaf, 0x86, 0x8f, 0x92, 0xd1, 0xfd, 0x94,
	0xcc, 0xe3, 0x7d, 0x38, 0xbb, 0xea, 0x4c, 0xfd, 0x47, 0xb5, 0x09, 0xc6, 0xd9, 0x1d, 0x5b, 0x71,
	0x98, 0xcd, 0x1d, 0x2c, 0xa3, 0x5c, 0x40, 0x25, 0x33, 0x8c, 0xf3, 0xfe, 0xf6, 0xe2, 0x66, 0x6b,
	0x17, 0x4c, 0x59, 0x77, 0x0f, 0xf3, 0x2f, 0x4b, 0x9b, 0x2f, 0x2d, 0xb4, 0x58, 0x5a, 0xe8, 0x73,
	0x69, 0xa1, 0xb7, 0x95, 0xa5, 0x2d, 0x56, 0x96, 0xf6, 0xb1, 0xb2, 0xb4, 0xc7, 0xfe, 0xe1, 0x2f,
	0xf2, 0xda, 0xa7, 0x62, 0x54, 0x4e, 0xfa, 0xe7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x37, 0x8c,
	0x85, 0xe8, 0xd2, 0x02, 0x00, 0x00,
}

func (m *Record) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDns(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovDns(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovDns(uint64(l))
	}
	if m.Priority != 0 {
		n += 1 + sovDns(uint64(m.Priority))
	}
	if m.Ttl != 0 {
		n += 1 + sovDns(uint64(m.Ttl))
	}
	return n
}

func (m *AdvertiseRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovDns(uint64(l))
		}
	}
	return n
}

func (m *AdvertiseResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *RemoveRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovDns(uint64(l))
		}
	}
	return n
}

func (m *RemoveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ResolveRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDns(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovDns(uint64(l))
	}
	return n
}

func (m *ResolveResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovDns(uint64(l))
		}
	}
	return n
}

func sovDns(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func sozDns(x uint64) (n int) {
	return sovDns(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Record) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ttl != 0 {
		i = encodeVarintDns(dAtA, i, uint64(m.Ttl))
		i--
		dAtA[i] = 0x28
	}
	if m.Priority != 0 {
		i = encodeVarintDns(dAtA, i, uint64(m.Priority))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintDns(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintDns(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDns(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AdvertiseRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdvertiseRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AdvertiseRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDns(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AdvertiseResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AdvertiseResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AdvertiseResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *RemoveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoveRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDns(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RemoveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ResolveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResolveRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResolveRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintDns(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDns(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResolveResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResolveResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResolveResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDns(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintDns(dAtA []byte, offset int, v uint64) int {
	offset -= sovDns(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDns
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ttl", wireType)
			}
			m.Ttl = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ttl |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDns
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AdvertiseRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDns
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AdvertiseRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdvertiseRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDns
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Records = append(m.Records, &Record{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDns
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AdvertiseResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDns
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AdvertiseResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AdvertiseResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDns
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RemoveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDns
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RemoveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDns
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Records = append(m.Records, &Record{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDns
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RemoveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDns
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RemoveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDns
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResolveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDns
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResolveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResolveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDns
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDns
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResolveResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDns
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResolveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResolveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDns
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDns
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDns
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Records = append(m.Records, &Record{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDns(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDns
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDns(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDns
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDns
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDns
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDns
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDns
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDns
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDns        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDns          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDns = fmt.Errorf("proto: unexpected end of group")
)

// DnsClient is the client API for Dns service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DnsClient interface {
	Advertise(ctx context.Context, in *AdvertiseRequest, opts ...grpc.CallOption) (*AdvertiseResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error)
}

type dnsClient struct {
	cc *grpc.ClientConn
}

func NewDnsClient(cc *grpc.ClientConn) DnsClient {
	return &dnsClient{cc}
}

func (c *dnsClient) Advertise(ctx context.Context, in *AdvertiseRequest, opts ...grpc.CallOption) (*AdvertiseResponse, error) {
	out := new(AdvertiseResponse)
	err := c.cc.Invoke(ctx, "/dns.Dns/Advertise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, "/dns.Dns/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsClient) Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error) {
	out := new(ResolveResponse)
	err := c.cc.Invoke(ctx, "/dns.Dns/Resolve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DnsServer is the server API for Dns service.
type DnsServer interface {
	Advertise(context.Context, *AdvertiseRequest) (*AdvertiseResponse, error)
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	Resolve(context.Context, *ResolveRequest) (*ResolveResponse, error)
}

// UnimplementedDnsServer can be embedded to have forward compatible implementations.
type UnimplementedDnsServer struct {
}

func (*UnimplementedDnsServer) Advertise(ctx context.Context, req *AdvertiseRequest) (*AdvertiseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Advertise not implemented")
}
func (*UnimplementedDnsServer) Remove(ctx context.Context, req *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (*UnimplementedDnsServer) Resolve(ctx context.Context, req *ResolveRequest) (*ResolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resolve not implemented")
}

func RegisterDnsServer(s *grpc.Server, srv DnsServer) {
	s.RegisterService(&_Dns_serviceDesc, srv)
}

func _Dns_Advertise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdvertiseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsServer).Advertise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dns.Dns/Advertise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsServer).Advertise(ctx, req.(*AdvertiseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dns_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dns.Dns/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dns_Resolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsServer).Resolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dns.Dns/Resolve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsServer).Resolve(ctx, req.(*ResolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dns_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dns.Dns",
	HandlerType: (*DnsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Advertise",
			Handler:    _Dns_Advertise_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Dns_Remove_Handler,
		},
		{
			MethodName: "Resolve",
			Handler:    _Dns_Resolve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/lack-io/vine/proto/services/network/dns/dns.proto",
}
