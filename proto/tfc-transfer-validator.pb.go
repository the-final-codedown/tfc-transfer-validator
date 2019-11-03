// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/tfc-transfer-validator.proto

package tfc_transfer_validator

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Transfer_TransferType int32

const (
	Transfer_CARD Transfer_TransferType = 0
	Transfer_WIRE Transfer_TransferType = 1
)

var Transfer_TransferType_name = map[int32]string{
	0: "CARD",
	1: "WIRE",
}

var Transfer_TransferType_value = map[string]int32{
	"CARD": 0,
	"WIRE": 1,
}

func (x Transfer_TransferType) String() string {
	return proto.EnumName(Transfer_TransferType_name, int32(x))
}

func (Transfer_TransferType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_90d5e49b340eb586, []int{0, 0}
}

type Transfer struct {
	Origin               int64                 `protobuf:"varint,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Destination          int64                 `protobuf:"varint,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Amount               int32                 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Type                 Transfer_TransferType `protobuf:"varint,4,opt,name=type,proto3,enum=tfc.transfer.validator.Transfer_TransferType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Transfer) Reset()         { *m = Transfer{} }
func (m *Transfer) String() string { return proto.CompactTextString(m) }
func (*Transfer) ProtoMessage()    {}
func (*Transfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_90d5e49b340eb586, []int{0}
}
func (m *Transfer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Transfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Transfer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Transfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transfer.Merge(m, src)
}
func (m *Transfer) XXX_Size() int {
	return m.Size()
}
func (m *Transfer) XXX_DiscardUnknown() {
	xxx_messageInfo_Transfer.DiscardUnknown(m)
}

var xxx_messageInfo_Transfer proto.InternalMessageInfo

func (m *Transfer) GetOrigin() int64 {
	if m != nil {
		return m.Origin
	}
	return 0
}

func (m *Transfer) GetDestination() int64 {
	if m != nil {
		return m.Destination
	}
	return 0
}

func (m *Transfer) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Transfer) GetType() Transfer_TransferType {
	if m != nil {
		return m.Type
	}
	return Transfer_CARD
}

type TransferValidation struct {
	Transfer             *Transfer `protobuf:"bytes,1,opt,name=transfer,proto3" json:"transfer,omitempty"`
	Validated            bool      `protobuf:"varint,2,opt,name=validated,proto3" json:"validated,omitempty"`
	Reason               string    `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TransferValidation) Reset()         { *m = TransferValidation{} }
func (m *TransferValidation) String() string { return proto.CompactTextString(m) }
func (*TransferValidation) ProtoMessage()    {}
func (*TransferValidation) Descriptor() ([]byte, []int) {
	return fileDescriptor_90d5e49b340eb586, []int{1}
}
func (m *TransferValidation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransferValidation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransferValidation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransferValidation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferValidation.Merge(m, src)
}
func (m *TransferValidation) XXX_Size() int {
	return m.Size()
}
func (m *TransferValidation) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferValidation.DiscardUnknown(m)
}

var xxx_messageInfo_TransferValidation proto.InternalMessageInfo

func (m *TransferValidation) GetTransfer() *Transfer {
	if m != nil {
		return m.Transfer
	}
	return nil
}

func (m *TransferValidation) GetValidated() bool {
	if m != nil {
		return m.Validated
	}
	return false
}

func (m *TransferValidation) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterEnum("tfc.transfer.validator.Transfer_TransferType", Transfer_TransferType_name, Transfer_TransferType_value)
	proto.RegisterType((*Transfer)(nil), "tfc.transfer.validator.Transfer")
	proto.RegisterType((*TransferValidation)(nil), "tfc.transfer.validator.TransferValidation")
}

func init() { proto.RegisterFile("proto/tfc-transfer-validator.proto", fileDescriptor_90d5e49b340eb586) }

var fileDescriptor_90d5e49b340eb586 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x49, 0x4b, 0xd6, 0x2d, 0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0x4b, 0x2d, 0xd2, 0x2d,
	0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x03, 0x4b, 0x0a, 0x89, 0x95, 0xa4, 0x25,
	0xeb, 0xc1, 0x64, 0xf5, 0xe0, 0xb2, 0x4a, 0x87, 0x19, 0xb9, 0x38, 0x42, 0xa0, 0xc2, 0x42, 0x62,
	0x5c, 0x6c, 0xf9, 0x45, 0x99, 0xe9, 0x99, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x50,
	0x9e, 0x90, 0x02, 0x17, 0x77, 0x4a, 0x6a, 0x71, 0x49, 0x66, 0x5e, 0x62, 0x49, 0x66, 0x7e, 0x9e,
	0x04, 0x13, 0x58, 0x12, 0x59, 0x08, 0xa4, 0x33, 0x31, 0x37, 0xbf, 0x34, 0xaf, 0x44, 0x82, 0x59,
	0x81, 0x51, 0x83, 0x35, 0x08, 0xca, 0x13, 0x72, 0xe4, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60,
	0x51, 0x60, 0xd4, 0xe0, 0x33, 0xd2, 0xd5, 0xc3, 0xee, 0x0a, 0x3d, 0x98, 0x0b, 0xe0, 0x8c, 0x90,
	0xca, 0x82, 0xd4, 0x20, 0xb0, 0x56, 0x25, 0x25, 0x2e, 0x1e, 0x64, 0x51, 0x21, 0x0e, 0x2e, 0x16,
	0x67, 0xc7, 0x20, 0x17, 0x01, 0x06, 0x10, 0x2b, 0xdc, 0x33, 0xc8, 0x55, 0x80, 0x51, 0xa9, 0x83,
	0x91, 0x4b, 0x08, 0xa6, 0x28, 0x0c, 0x62, 0x2a, 0xc8, 0x55, 0x36, 0x5c, 0x1c, 0x30, 0xcb, 0xc0,
	0x3e, 0xe2, 0x36, 0x52, 0x20, 0xe4, 0x82, 0x20, 0xb8, 0x0e, 0x21, 0x19, 0x2e, 0x4e, 0xa8, 0x7c,
	0x6a, 0x0a, 0xd8, 0xcf, 0x1c, 0x41, 0x08, 0x01, 0x90, 0x8f, 0x8b, 0x52, 0x13, 0x8b, 0xf3, 0xf3,
	0xc0, 0x3e, 0xe6, 0x0c, 0x82, 0xf2, 0x8c, 0x0a, 0xb9, 0x24, 0xd0, 0x5c, 0x92, 0x5f, 0x14, 0x9c,
	0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x14, 0xca, 0xc5, 0x1c, 0x90, 0x58, 0x29, 0x44, 0xd0, 0x11,
	0x52, 0x5a, 0x84, 0x54, 0x20, 0x3c, 0xa9, 0xc4, 0xe0, 0x24, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0x8e, 0x74,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x80, 0x7c, 0xba, 0x87, 0x1a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransferValidatorServiceClient is the client API for TransferValidatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransferValidatorServiceClient interface {
	Pay(ctx context.Context, in *Transfer, opts ...grpc.CallOption) (*TransferValidation, error)
}

type transferValidatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewTransferValidatorServiceClient(cc *grpc.ClientConn) TransferValidatorServiceClient {
	return &transferValidatorServiceClient{cc}
}

func (c *transferValidatorServiceClient) Pay(ctx context.Context, in *Transfer, opts ...grpc.CallOption) (*TransferValidation, error) {
	out := new(TransferValidation)
	err := c.cc.Invoke(ctx, "/tfc.transfer.validator.TransferValidatorService/Pay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransferValidatorServiceServer is the server API for TransferValidatorService service.
type TransferValidatorServiceServer interface {
	Pay(context.Context, *Transfer) (*TransferValidation, error)
}

// UnimplementedTransferValidatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTransferValidatorServiceServer struct {
}

func (*UnimplementedTransferValidatorServiceServer) Pay(ctx context.Context, req *Transfer) (*TransferValidation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pay not implemented")
}

func RegisterTransferValidatorServiceServer(s *grpc.Server, srv TransferValidatorServiceServer) {
	s.RegisterService(&_TransferValidatorService_serviceDesc, srv)
}

func _TransferValidatorService_Pay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferValidatorServiceServer).Pay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfc.transfer.validator.TransferValidatorService/Pay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferValidatorServiceServer).Pay(ctx, req.(*Transfer))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransferValidatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tfc.transfer.validator.TransferValidatorService",
	HandlerType: (*TransferValidatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pay",
			Handler:    _TransferValidatorService_Pay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tfc-transfer-validator.proto",
}

func (m *Transfer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transfer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Transfer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Type != 0 {
		i = encodeVarintTfcTransferValidator(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x20
	}
	if m.Amount != 0 {
		i = encodeVarintTfcTransferValidator(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if m.Destination != 0 {
		i = encodeVarintTfcTransferValidator(dAtA, i, uint64(m.Destination))
		i--
		dAtA[i] = 0x10
	}
	if m.Origin != 0 {
		i = encodeVarintTfcTransferValidator(dAtA, i, uint64(m.Origin))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TransferValidation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransferValidation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransferValidation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintTfcTransferValidator(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Validated {
		i--
		if m.Validated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Transfer != nil {
		{
			size, err := m.Transfer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTfcTransferValidator(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTfcTransferValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovTfcTransferValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Transfer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Origin != 0 {
		n += 1 + sovTfcTransferValidator(uint64(m.Origin))
	}
	if m.Destination != 0 {
		n += 1 + sovTfcTransferValidator(uint64(m.Destination))
	}
	if m.Amount != 0 {
		n += 1 + sovTfcTransferValidator(uint64(m.Amount))
	}
	if m.Type != 0 {
		n += 1 + sovTfcTransferValidator(uint64(m.Type))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TransferValidation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Transfer != nil {
		l = m.Transfer.Size()
		n += 1 + l + sovTfcTransferValidator(uint64(l))
	}
	if m.Validated {
		n += 2
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovTfcTransferValidator(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTfcTransferValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTfcTransferValidator(x uint64) (n int) {
	return sovTfcTransferValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Transfer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTfcTransferValidator
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
			return fmt.Errorf("proto: Transfer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transfer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			m.Origin = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTfcTransferValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Origin |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destination", wireType)
			}
			m.Destination = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTfcTransferValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Destination |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTfcTransferValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTfcTransferValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Transfer_TransferType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTfcTransferValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTfcTransferValidator
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTfcTransferValidator
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
func (m *TransferValidation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTfcTransferValidator
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
			return fmt.Errorf("proto: TransferValidation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransferValidation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transfer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTfcTransferValidator
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
				return ErrInvalidLengthTfcTransferValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTfcTransferValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Transfer == nil {
				m.Transfer = &Transfer{}
			}
			if err := m.Transfer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTfcTransferValidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Validated = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTfcTransferValidator
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
				return ErrInvalidLengthTfcTransferValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTfcTransferValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTfcTransferValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTfcTransferValidator
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTfcTransferValidator
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
func skipTfcTransferValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTfcTransferValidator
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
					return 0, ErrIntOverflowTfcTransferValidator
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
					return 0, ErrIntOverflowTfcTransferValidator
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
				return 0, ErrInvalidLengthTfcTransferValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTfcTransferValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTfcTransferValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTfcTransferValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTfcTransferValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTfcTransferValidator = fmt.Errorf("proto: unexpected end of group")
)