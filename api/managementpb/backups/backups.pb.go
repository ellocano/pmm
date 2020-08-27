// Code generated by protoc-gen-go. DO NOT EDIT.
// source: managementpb/backups/backups.proto

package backupsv1beta1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ListBackupsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBackupsRequest) Reset()         { *m = ListBackupsRequest{} }
func (m *ListBackupsRequest) String() string { return proto.CompactTextString(m) }
func (*ListBackupsRequest) ProtoMessage()    {}
func (*ListBackupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1343b6655f36147, []int{0}
}

func (m *ListBackupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBackupsRequest.Unmarshal(m, b)
}
func (m *ListBackupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBackupsRequest.Marshal(b, m, deterministic)
}
func (m *ListBackupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBackupsRequest.Merge(m, src)
}
func (m *ListBackupsRequest) XXX_Size() int {
	return xxx_messageInfo_ListBackupsRequest.Size(m)
}
func (m *ListBackupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBackupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBackupsRequest proto.InternalMessageInfo

type ListBackupsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBackupsResponse) Reset()         { *m = ListBackupsResponse{} }
func (m *ListBackupsResponse) String() string { return proto.CompactTextString(m) }
func (*ListBackupsResponse) ProtoMessage()    {}
func (*ListBackupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1343b6655f36147, []int{1}
}

func (m *ListBackupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBackupsResponse.Unmarshal(m, b)
}
func (m *ListBackupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBackupsResponse.Marshal(b, m, deterministic)
}
func (m *ListBackupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBackupsResponse.Merge(m, src)
}
func (m *ListBackupsResponse) XXX_Size() int {
	return xxx_messageInfo_ListBackupsResponse.Size(m)
}
func (m *ListBackupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBackupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBackupsResponse proto.InternalMessageInfo

type StartBackupRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartBackupRequest) Reset()         { *m = StartBackupRequest{} }
func (m *StartBackupRequest) String() string { return proto.CompactTextString(m) }
func (*StartBackupRequest) ProtoMessage()    {}
func (*StartBackupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1343b6655f36147, []int{2}
}

func (m *StartBackupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartBackupRequest.Unmarshal(m, b)
}
func (m *StartBackupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartBackupRequest.Marshal(b, m, deterministic)
}
func (m *StartBackupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartBackupRequest.Merge(m, src)
}
func (m *StartBackupRequest) XXX_Size() int {
	return xxx_messageInfo_StartBackupRequest.Size(m)
}
func (m *StartBackupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartBackupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartBackupRequest proto.InternalMessageInfo

type StartBackupResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartBackupResponse) Reset()         { *m = StartBackupResponse{} }
func (m *StartBackupResponse) String() string { return proto.CompactTextString(m) }
func (*StartBackupResponse) ProtoMessage()    {}
func (*StartBackupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1343b6655f36147, []int{3}
}

func (m *StartBackupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartBackupResponse.Unmarshal(m, b)
}
func (m *StartBackupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartBackupResponse.Marshal(b, m, deterministic)
}
func (m *StartBackupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartBackupResponse.Merge(m, src)
}
func (m *StartBackupResponse) XXX_Size() int {
	return xxx_messageInfo_StartBackupResponse.Size(m)
}
func (m *StartBackupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartBackupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartBackupResponse proto.InternalMessageInfo

type RestoreBackupRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestoreBackupRequest) Reset()         { *m = RestoreBackupRequest{} }
func (m *RestoreBackupRequest) String() string { return proto.CompactTextString(m) }
func (*RestoreBackupRequest) ProtoMessage()    {}
func (*RestoreBackupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1343b6655f36147, []int{4}
}

func (m *RestoreBackupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestoreBackupRequest.Unmarshal(m, b)
}
func (m *RestoreBackupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestoreBackupRequest.Marshal(b, m, deterministic)
}
func (m *RestoreBackupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestoreBackupRequest.Merge(m, src)
}
func (m *RestoreBackupRequest) XXX_Size() int {
	return xxx_messageInfo_RestoreBackupRequest.Size(m)
}
func (m *RestoreBackupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RestoreBackupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RestoreBackupRequest proto.InternalMessageInfo

type RestoreBackupResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestoreBackupResponse) Reset()         { *m = RestoreBackupResponse{} }
func (m *RestoreBackupResponse) String() string { return proto.CompactTextString(m) }
func (*RestoreBackupResponse) ProtoMessage()    {}
func (*RestoreBackupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1343b6655f36147, []int{5}
}

func (m *RestoreBackupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestoreBackupResponse.Unmarshal(m, b)
}
func (m *RestoreBackupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestoreBackupResponse.Marshal(b, m, deterministic)
}
func (m *RestoreBackupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestoreBackupResponse.Merge(m, src)
}
func (m *RestoreBackupResponse) XXX_Size() int {
	return xxx_messageInfo_RestoreBackupResponse.Size(m)
}
func (m *RestoreBackupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RestoreBackupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RestoreBackupResponse proto.InternalMessageInfo

type RemoveBackupRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveBackupRequest) Reset()         { *m = RemoveBackupRequest{} }
func (m *RemoveBackupRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveBackupRequest) ProtoMessage()    {}
func (*RemoveBackupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1343b6655f36147, []int{6}
}

func (m *RemoveBackupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveBackupRequest.Unmarshal(m, b)
}
func (m *RemoveBackupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveBackupRequest.Marshal(b, m, deterministic)
}
func (m *RemoveBackupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveBackupRequest.Merge(m, src)
}
func (m *RemoveBackupRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveBackupRequest.Size(m)
}
func (m *RemoveBackupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveBackupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveBackupRequest proto.InternalMessageInfo

type RemoveBackupResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveBackupResponse) Reset()         { *m = RemoveBackupResponse{} }
func (m *RemoveBackupResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveBackupResponse) ProtoMessage()    {}
func (*RemoveBackupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1343b6655f36147, []int{7}
}

func (m *RemoveBackupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveBackupResponse.Unmarshal(m, b)
}
func (m *RemoveBackupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveBackupResponse.Marshal(b, m, deterministic)
}
func (m *RemoveBackupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveBackupResponse.Merge(m, src)
}
func (m *RemoveBackupResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveBackupResponse.Size(m)
}
func (m *RemoveBackupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveBackupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveBackupResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListBackupsRequest)(nil), "backups.v1beta1.ListBackupsRequest")
	proto.RegisterType((*ListBackupsResponse)(nil), "backups.v1beta1.ListBackupsResponse")
	proto.RegisterType((*StartBackupRequest)(nil), "backups.v1beta1.StartBackupRequest")
	proto.RegisterType((*StartBackupResponse)(nil), "backups.v1beta1.StartBackupResponse")
	proto.RegisterType((*RestoreBackupRequest)(nil), "backups.v1beta1.RestoreBackupRequest")
	proto.RegisterType((*RestoreBackupResponse)(nil), "backups.v1beta1.RestoreBackupResponse")
	proto.RegisterType((*RemoveBackupRequest)(nil), "backups.v1beta1.RemoveBackupRequest")
	proto.RegisterType((*RemoveBackupResponse)(nil), "backups.v1beta1.RemoveBackupResponse")
}

func init() {
	proto.RegisterFile("managementpb/backups/backups.proto", fileDescriptor_f1343b6655f36147)
}

var fileDescriptor_f1343b6655f36147 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x65, 0x84, 0x40, 0x72, 0xb9, 0xc9, 0x6d, 0x41, 0x0a, 0x05, 0x81, 0x7b, 0x01, 0x3a,
	0xc4, 0x0a, 0x6c, 0x1d, 0x3b, 0x33, 0x85, 0x8d, 0xcd, 0x41, 0x56, 0x14, 0x41, 0xec, 0x10, 0xbb,
	0x99, 0x01, 0x89, 0x81, 0x99, 0x57, 0xe0, 0x8d, 0x78, 0x05, 0x1e, 0x04, 0xc5, 0x3d, 0x15, 0xb9,
	0xd2, 0x29, 0xca, 0xd1, 0xf7, 0xdb, 0x5f, 0xfe, 0x13, 0x4c, 0x63, 0x2e, 0x79, 0x28, 0x62, 0x21,
	0x4d, 0x12, 0xb0, 0x80, 0x3f, 0x3c, 0x2e, 0x12, 0xbd, 0x7a, 0xba, 0x49, 0xaa, 0x8c, 0x22, 0xfb,
	0xab, 0xd7, 0xcc, 0x0b, 0x84, 0xe1, 0x9e, 0x33, 0x08, 0x95, 0x0a, 0x9f, 0x04, 0xe3, 0x49, 0xc4,
	0xb8, 0x94, 0xca, 0x70, 0x13, 0x29, 0x09, 0x38, 0xed, 0x61, 0x72, 0x1b, 0x69, 0x33, 0x5f, 0x86,
	0x7c, 0xf1, 0xbc, 0x10, 0xda, 0xd0, 0x3e, 0xee, 0x96, 0xa6, 0x3a, 0x51, 0x52, 0x8b, 0x1c, 0xbe,
	0x33, 0x3c, 0x85, 0x79, 0x01, 0x2e, 0x4d, 0x01, 0x3e, 0xc4, 0x3d, 0x5f, 0x68, 0xa3, 0x52, 0x51,
	0xc6, 0x8f, 0x70, 0xbf, 0x32, 0x87, 0x40, 0x1f, 0x77, 0x7d, 0x11, 0xab, 0xac, 0xc2, 0xdb, 0x73,
	0x8a, 0xe3, 0x25, 0x7e, 0xfd, 0xb5, 0x89, 0xb7, 0x41, 0x90, 0xbc, 0x20, 0xdc, 0x29, 0x08, 0x93,
	0xa1, 0x5b, 0x69, 0xc1, 0xad, 0x7f, 0xa4, 0x33, 0xfa, 0x1f, 0x02, 0xab, 0xc9, 0xdb, 0xf7, 0xcf,
	0xe7, 0xc6, 0x19, 0x3d, 0x66, 0x99, 0xc7, 0xfe, 0xfa, 0x67, 0xc0, 0xb1, 0x3c, 0x33, 0x43, 0x53,
	0xf2, 0x8a, 0x70, 0xa7, 0x50, 0x43, 0x83, 0x42, 0xbd, 0xba, 0x06, 0x85, 0xa6, 0x26, 0x2f, 0xac,
	0xc2, 0x39, 0x1d, 0xb4, 0x28, 0xd8, 0x4c, 0xee, 0xf0, 0x81, 0xf0, 0x6e, 0xa9, 0x5b, 0x32, 0xae,
	0x5d, 0xd0, 0xb4, 0x13, 0x67, 0xb2, 0x0e, 0x03, 0x93, 0x2b, 0x6b, 0x32, 0xa4, 0xa7, 0x2d, 0x26,
	0x90, 0xca, 0x5d, 0xde, 0x11, 0xde, 0x29, 0xee, 0x8d, 0x8c, 0x1a, 0xee, 0xa8, 0x6d, 0xdb, 0x19,
	0xaf, 0xa1, 0x40, 0xe4, 0xd2, 0x8a, 0x50, 0x7a, 0xd2, 0x2a, 0x92, 0x87, 0x66, 0x68, 0x3a, 0x3f,
	0xb8, 0xdf, 0x83, 0x13, 0xe1, 0xc0, 0x60, 0xcb, 0xfe, 0xf9, 0x37, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x0c, 0xdf, 0xd5, 0x26, 0x4e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BackupsClient is the client API for Backups service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackupsClient interface {
	// ListBackups returns a list of backups jobs and files.
	ListBackups(ctx context.Context, in *ListBackupsRequest, opts ...grpc.CallOption) (*ListBackupsResponse, error)
	// StartBackup TODO.
	StartBackup(ctx context.Context, in *StartBackupRequest, opts ...grpc.CallOption) (*StartBackupResponse, error)
	// RestoreBackup TODO.
	RestoreBackup(ctx context.Context, in *RestoreBackupRequest, opts ...grpc.CallOption) (*RestoreBackupResponse, error)
	// RemoveBackup stops the backup job if it is still running, and deletes the backup file.
	RemoveBackup(ctx context.Context, in *RemoveBackupRequest, opts ...grpc.CallOption) (*RemoveBackupResponse, error)
}

type backupsClient struct {
	cc grpc.ClientConnInterface
}

func NewBackupsClient(cc grpc.ClientConnInterface) BackupsClient {
	return &backupsClient{cc}
}

func (c *backupsClient) ListBackups(ctx context.Context, in *ListBackupsRequest, opts ...grpc.CallOption) (*ListBackupsResponse, error) {
	out := new(ListBackupsResponse)
	err := c.cc.Invoke(ctx, "/backups.v1beta1.Backups/ListBackups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupsClient) StartBackup(ctx context.Context, in *StartBackupRequest, opts ...grpc.CallOption) (*StartBackupResponse, error) {
	out := new(StartBackupResponse)
	err := c.cc.Invoke(ctx, "/backups.v1beta1.Backups/StartBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupsClient) RestoreBackup(ctx context.Context, in *RestoreBackupRequest, opts ...grpc.CallOption) (*RestoreBackupResponse, error) {
	out := new(RestoreBackupResponse)
	err := c.cc.Invoke(ctx, "/backups.v1beta1.Backups/RestoreBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupsClient) RemoveBackup(ctx context.Context, in *RemoveBackupRequest, opts ...grpc.CallOption) (*RemoveBackupResponse, error) {
	out := new(RemoveBackupResponse)
	err := c.cc.Invoke(ctx, "/backups.v1beta1.Backups/RemoveBackup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupsServer is the server API for Backups service.
type BackupsServer interface {
	// ListBackups returns a list of backups jobs and files.
	ListBackups(context.Context, *ListBackupsRequest) (*ListBackupsResponse, error)
	// StartBackup TODO.
	StartBackup(context.Context, *StartBackupRequest) (*StartBackupResponse, error)
	// RestoreBackup TODO.
	RestoreBackup(context.Context, *RestoreBackupRequest) (*RestoreBackupResponse, error)
	// RemoveBackup stops the backup job if it is still running, and deletes the backup file.
	RemoveBackup(context.Context, *RemoveBackupRequest) (*RemoveBackupResponse, error)
}

// UnimplementedBackupsServer can be embedded to have forward compatible implementations.
type UnimplementedBackupsServer struct {
}

func (*UnimplementedBackupsServer) ListBackups(ctx context.Context, req *ListBackupsRequest) (*ListBackupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBackups not implemented")
}
func (*UnimplementedBackupsServer) StartBackup(ctx context.Context, req *StartBackupRequest) (*StartBackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartBackup not implemented")
}
func (*UnimplementedBackupsServer) RestoreBackup(ctx context.Context, req *RestoreBackupRequest) (*RestoreBackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestoreBackup not implemented")
}
func (*UnimplementedBackupsServer) RemoveBackup(ctx context.Context, req *RemoveBackupRequest) (*RemoveBackupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBackup not implemented")
}

func RegisterBackupsServer(s *grpc.Server, srv BackupsServer) {
	s.RegisterService(&_Backups_serviceDesc, srv)
}

func _Backups_ListBackups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBackupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupsServer).ListBackups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backups.v1beta1.Backups/ListBackups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupsServer).ListBackups(ctx, req.(*ListBackupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backups_StartBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupsServer).StartBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backups.v1beta1.Backups/StartBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupsServer).StartBackup(ctx, req.(*StartBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backups_RestoreBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupsServer).RestoreBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backups.v1beta1.Backups/RestoreBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupsServer).RestoreBackup(ctx, req.(*RestoreBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backups_RemoveBackup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupsServer).RemoveBackup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backups.v1beta1.Backups/RemoveBackup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupsServer).RemoveBackup(ctx, req.(*RemoveBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Backups_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backups.v1beta1.Backups",
	HandlerType: (*BackupsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBackups",
			Handler:    _Backups_ListBackups_Handler,
		},
		{
			MethodName: "StartBackup",
			Handler:    _Backups_StartBackup_Handler,
		},
		{
			MethodName: "RestoreBackup",
			Handler:    _Backups_RestoreBackup_Handler,
		},
		{
			MethodName: "RemoveBackup",
			Handler:    _Backups_RemoveBackup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/backups/backups.proto",
}
