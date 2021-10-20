// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service/protos/queryPrinter.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_06913f648ed2eace, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type Response struct {
	Param1 string `protobuf:"bytes,1,opt,name=param1,proto3" json:"param1,omitempty"`
	Param2 string `protobuf:"bytes,2,opt,name=param2,proto3" json:"param2,omitempty"`
	Param3 string `protobuf:"bytes,3,opt,name=param3,proto3" json:"param3,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_06913f648ed2eace, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return m.Size()
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetParam1() string {
	if m != nil {
		return m.Param1
	}
	return ""
}

func (m *Response) GetParam2() string {
	if m != nil {
		return m.Param2
	}
	return ""
}

func (m *Response) GetParam3() string {
	if m != nil {
		return m.Param3
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "protos.Request")
	proto.RegisterType((*Response)(nil), "protos.Response")
}

func init() { proto.RegisterFile("service/protos/queryPrinter.proto", fileDescriptor_06913f648ed2eace) }

var fileDescriptor_06913f648ed2eace = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa,
	0x0c, 0x28, 0xca, 0xcc, 0x2b, 0x49, 0x2d, 0xd2, 0x03, 0x8b, 0x09, 0xb1, 0x41, 0xa4, 0x94, 0xe4,
	0xb9, 0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0xc1, 0x0a, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0xa5, 0x20, 0x2e, 0x8e, 0xa0, 0xd4, 0xe2, 0x82,
	0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0x82, 0xc4, 0xa2, 0xc4, 0x5c, 0x43, 0xa8, 0x12,
	0x28, 0x0f, 0x2e, 0x6e, 0x24, 0xc1, 0x84, 0x24, 0x6e, 0x04, 0x17, 0x37, 0x96, 0x60, 0x46, 0x12,
	0x37, 0x36, 0xf2, 0xe0, 0x12, 0x08, 0x44, 0x72, 0x92, 0x7b, 0x50, 0x80, 0xb3, 0x90, 0x09, 0x17,
	0x3f, 0x98, 0x1b, 0x00, 0x52, 0x92, 0x5a, 0x92, 0x5a, 0x54, 0x2c, 0xc4, 0x0f, 0x71, 0x6b, 0xb1,
	0x1e, 0xd4, 0x85, 0x52, 0x02, 0x08, 0x01, 0x88, 0x8b, 0x9c, 0x24, 0x4e, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63,
	0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x09, 0xe2, 0x41, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5,
	0xe4, 0xe7, 0x02, 0x0c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryPrinterGRPCClient is the client API for QueryPrinterGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryPrinterGRPCClient interface {
	PrintParameters(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type queryPrinterGRPCClient struct {
	cc *grpc.ClientConn
}

func NewQueryPrinterGRPCClient(cc *grpc.ClientConn) QueryPrinterGRPCClient {
	return &queryPrinterGRPCClient{cc}
}

func (c *queryPrinterGRPCClient) PrintParameters(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protos.QueryPrinterGRPC/PrintParameters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryPrinterGRPCServer is the server API for QueryPrinterGRPC service.
type QueryPrinterGRPCServer interface {
	PrintParameters(context.Context, *Request) (*Response, error)
}

// UnimplementedQueryPrinterGRPCServer can be embedded to have forward compatible implementations.
type UnimplementedQueryPrinterGRPCServer struct {
}

func (*UnimplementedQueryPrinterGRPCServer) PrintParameters(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrintParameters not implemented")
}

func RegisterQueryPrinterGRPCServer(s *grpc.Server, srv QueryPrinterGRPCServer) {
	s.RegisterService(&_QueryPrinterGRPC_serviceDesc, srv)
}

func _QueryPrinterGRPC_PrintParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryPrinterGRPCServer).PrintParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.QueryPrinterGRPC/PrintParameters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryPrinterGRPCServer).PrintParameters(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryPrinterGRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.QueryPrinterGRPC",
	HandlerType: (*QueryPrinterGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrintParameters",
			Handler:    _QueryPrinterGRPC_PrintParameters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/protos/queryPrinter.proto",
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Query) > 0 {
		i -= len(m.Query)
		copy(dAtA[i:], m.Query)
		i = encodeVarintQueryPrinter(dAtA, i, uint64(len(m.Query)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Param3) > 0 {
		i -= len(m.Param3)
		copy(dAtA[i:], m.Param3)
		i = encodeVarintQueryPrinter(dAtA, i, uint64(len(m.Param3)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Param2) > 0 {
		i -= len(m.Param2)
		copy(dAtA[i:], m.Param2)
		i = encodeVarintQueryPrinter(dAtA, i, uint64(len(m.Param2)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Param1) > 0 {
		i -= len(m.Param1)
		copy(dAtA[i:], m.Param1)
		i = encodeVarintQueryPrinter(dAtA, i, uint64(len(m.Param1)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryPrinter(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryPrinter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovQueryPrinter(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Param1)
	if l > 0 {
		n += 1 + l + sovQueryPrinter(uint64(l))
	}
	l = len(m.Param2)
	if l > 0 {
		n += 1 + l + sovQueryPrinter(uint64(l))
	}
	l = len(m.Param3)
	if l > 0 {
		n += 1 + l + sovQueryPrinter(uint64(l))
	}
	return n
}

func sovQueryPrinter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryPrinter(x uint64) (n int) {
	return sovQueryPrinter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryPrinter
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryPrinter
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
				return ErrInvalidLengthQueryPrinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryPrinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryPrinter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryPrinter
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
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryPrinter
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Param1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryPrinter
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
				return ErrInvalidLengthQueryPrinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryPrinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Param1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Param2", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryPrinter
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
				return ErrInvalidLengthQueryPrinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryPrinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Param2 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Param3", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryPrinter
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
				return ErrInvalidLengthQueryPrinter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryPrinter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Param3 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryPrinter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryPrinter
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
func skipQueryPrinter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryPrinter
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
					return 0, ErrIntOverflowQueryPrinter
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
					return 0, ErrIntOverflowQueryPrinter
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
				return 0, ErrInvalidLengthQueryPrinter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryPrinter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryPrinter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryPrinter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryPrinter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryPrinter = fmt.Errorf("proto: unexpected end of group")
)
