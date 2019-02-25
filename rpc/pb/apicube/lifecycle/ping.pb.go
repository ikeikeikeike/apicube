// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: apicube/lifecycle/ping.proto

package lifecycle

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Ping struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_290f9eea345e8e30, []int{0}
}
func (m *Ping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return m.Size()
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Ping)(nil), "apicube.lifecycle.Ping")
}

func init() { proto.RegisterFile("apicube/lifecycle/ping.proto", fileDescriptor_290f9eea345e8e30) }

var fileDescriptor_290f9eea345e8e30 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4a, 0xf4, 0x30,
	0x1c, 0xc5, 0xc9, 0x7c, 0x1f, 0x23, 0x93, 0xd1, 0x85, 0x55, 0x70, 0x28, 0x43, 0x1c, 0x66, 0xe5,
	0x66, 0x1a, 0x50, 0xf0, 0x00, 0xdd, 0xb8, 0x2d, 0xba, 0x50, 0xdc, 0x48, 0x1a, 0x33, 0x31, 0xd8,
	0xf6, 0x1f, 0xd2, 0x74, 0x8a, 0x88, 0x1b, 0xaf, 0xe0, 0x25, 0x3c, 0x82, 0x4b, 0x97, 0x2e, 0x05,
	0x0f, 0xa0, 0x06, 0x0f, 0xe1, 0x52, 0x9a, 0xa9, 0x75, 0x21, 0xee, 0x5e, 0xde, 0xef, 0xbd, 0x47,
	0x12, 0x3c, 0x66, 0x5a, 0xf1, 0x2a, 0x15, 0x34, 0x53, 0x73, 0xc1, 0xaf, 0x78, 0x26, 0xa8, 0x56,
	0x85, 0x8c, 0xb4, 0x01, 0x0b, 0xc1, 0x7a, 0x4b, 0xa3, 0x8e, 0x86, 0x63, 0x09, 0x20, 0x33, 0x41,
	0x99, 0x56, 0x94, 0x15, 0x05, 0x58, 0x66, 0x15, 0x14, 0xe5, 0xb2, 0x10, 0x92, 0x96, 0xfa, 0x53,
	0x5a, 0xcd, 0x69, 0x6d, 0x98, 0xd6, 0xc2, 0x7c, 0xf3, 0x4d, 0x09, 0x12, 0xbc, 0xa4, 0x8d, 0x6a,
	0xdd, 0x7d, 0xa9, 0xec, 0x45, 0x95, 0x46, 0x1c, 0x72, 0x9a, 0xd7, 0xca, 0x5e, 0x42, 0x4d, 0x25,
	0xcc, 0x3c, 0x9c, 0x2d, 0x58, 0xa6, 0xce, 0x99, 0x05, 0x53, 0xd2, 0x4e, 0x2e, 0x7b, 0xd3, 0x09,
	0xfe, 0x9f, 0xa8, 0x42, 0x06, 0x23, 0xfc, 0x2f, 0x2f, 0xe5, 0x08, 0x4d, 0xd0, 0xce, 0x20, 0xee,
	0xbb, 0xd7, 0xed, 0xde, 0x09, 0x3a, 0x6c, 0xac, 0xdd, 0x33, 0x3c, 0x6c, 0x12, 0x47, 0xc2, 0x2c,
	0x14, 0x17, 0x41, 0x82, 0x57, 0x0e, 0x84, 0xf5, 0x9d, 0xad, 0xe8, 0xd7, 0xdb, 0xa2, 0x06, 0x84,
	0x7f, 0x81, 0xe9, 0xc6, 0xed, 0xcb, 0xc7, 0x5d, 0x6f, 0x2d, 0x18, 0xfa, 0x1f, 0xa2, 0xd7, 0x79,
	0x29, 0x6f, 0xe2, 0xe3, 0xcf, 0x77, 0x82, 0xee, 0x1d, 0x41, 0x0f, 0x8e, 0xa0, 0x47, 0x47, 0xd0,
	0x93, 0x23, 0xe8, 0xd9, 0x11, 0xf4, 0xe6, 0x08, 0xc2, 0x04, 0x8c, 0xec, 0xe6, 0xfc, 0x9d, 0x39,
	0x64, 0x3f, 0xbb, 0xf1, 0x6a, 0x33, 0x9c, 0xb4, 0x7e, 0x82, 0x4e, 0x07, 0x1d, 0x4a, 0xfb, 0x3e,
	0xbe, 0xf7, 0x15, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x61, 0x56, 0x24, 0xa1, 0x01, 0x00, 0x00,
}

func (this *Ping) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Ping)
	if !ok {
		that2, ok := that.(Ping)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Msg != that1.Msg {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PingServiceClient is the client API for PingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingServiceClient interface {
	GetPing(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Ping, error)
}

type pingServiceClient struct {
	cc *grpc.ClientConn
}

func NewPingServiceClient(cc *grpc.ClientConn) PingServiceClient {
	return &pingServiceClient{cc}
}

func (c *pingServiceClient) GetPing(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Ping, error) {
	out := new(Ping)
	err := c.cc.Invoke(ctx, "/apicube.lifecycle.PingService/GetPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingServiceServer is the server API for PingService service.
type PingServiceServer interface {
	GetPing(context.Context, *Ping) (*Ping, error)
}

func RegisterPingServiceServer(s *grpc.Server, srv PingServiceServer) {
	s.RegisterService(&_PingService_serviceDesc, srv)
}

func _PingService_GetPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServiceServer).GetPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apicube.lifecycle.PingService/GetPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServiceServer).GetPing(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apicube.lifecycle.PingService",
	HandlerType: (*PingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPing",
			Handler:    _PingService_GetPing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apicube/lifecycle/ping.proto",
}

func (m *Ping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ping) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Msg) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPing(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintPing(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedPing(r randyPing, easy bool) *Ping {
	this := &Ping{}
	this.Msg = string(randStringPing(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedPing(r, 2)
	}
	return this
}

type randyPing interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RunePing(r randyPing) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringPing(r randyPing) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RunePing(r)
	}
	return string(tmps)
}
func randUnrecognizedPing(r randyPing, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldPing(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldPing(dAtA []byte, r randyPing, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulatePing(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulatePing(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulatePing(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulatePing(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulatePing(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulatePing(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulatePing(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Ping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovPing(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPing(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPing(x uint64) (n int) {
	return sovPing(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Ping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPing
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
			return fmt.Errorf("proto: Ping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPing
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
				return ErrInvalidLengthPing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPing
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPing
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPing(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPing
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
					return 0, ErrIntOverflowPing
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPing
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
				return 0, ErrInvalidLengthPing
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthPing
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPing
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPing(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthPing
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPing = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPing   = fmt.Errorf("proto: integer overflow")
)
