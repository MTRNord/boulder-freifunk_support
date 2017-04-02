// Code generated by protoc-gen-go.
// source: ra/proto/ra.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	ra/proto/ra.proto

It has these top-level messages:
	NewAuthorizationRequest
	NewCertificateRequest
	UpdateRegistrationRequest
	UpdateAuthorizationRequest
	RevokeCertificateWithRegRequest
	AdministrativelyRevokeCertificateRequest
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/letsencrypt/boulder/core/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type NewAuthorizationRequest struct {
	Authz            *core.Authorization `protobuf:"bytes,1,opt,name=authz" json:"authz,omitempty"`
	RegID            *int64              `protobuf:"varint,2,opt,name=regID" json:"regID,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *NewAuthorizationRequest) Reset()                    { *m = NewAuthorizationRequest{} }
func (m *NewAuthorizationRequest) String() string            { return proto1.CompactTextString(m) }
func (*NewAuthorizationRequest) ProtoMessage()               {}
func (*NewAuthorizationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewAuthorizationRequest) GetAuthz() *core.Authorization {
	if m != nil {
		return m.Authz
	}
	return nil
}

func (m *NewAuthorizationRequest) GetRegID() int64 {
	if m != nil && m.RegID != nil {
		return *m.RegID
	}
	return 0
}

type NewCertificateRequest struct {
	Csr              []byte `protobuf:"bytes,1,opt,name=csr" json:"csr,omitempty"`
	RegID            *int64 `protobuf:"varint,2,opt,name=regID" json:"regID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *NewCertificateRequest) Reset()                    { *m = NewCertificateRequest{} }
func (m *NewCertificateRequest) String() string            { return proto1.CompactTextString(m) }
func (*NewCertificateRequest) ProtoMessage()               {}
func (*NewCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NewCertificateRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

func (m *NewCertificateRequest) GetRegID() int64 {
	if m != nil && m.RegID != nil {
		return *m.RegID
	}
	return 0
}

type UpdateRegistrationRequest struct {
	Base             *core.Registration `protobuf:"bytes,1,opt,name=base" json:"base,omitempty"`
	Update           *core.Registration `protobuf:"bytes,2,opt,name=update" json:"update,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *UpdateRegistrationRequest) Reset()                    { *m = UpdateRegistrationRequest{} }
func (m *UpdateRegistrationRequest) String() string            { return proto1.CompactTextString(m) }
func (*UpdateRegistrationRequest) ProtoMessage()               {}
func (*UpdateRegistrationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateRegistrationRequest) GetBase() *core.Registration {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *UpdateRegistrationRequest) GetUpdate() *core.Registration {
	if m != nil {
		return m.Update
	}
	return nil
}

type UpdateAuthorizationRequest struct {
	Authz            *core.Authorization `protobuf:"bytes,1,opt,name=authz" json:"authz,omitempty"`
	ChallengeIndex   *int64              `protobuf:"varint,2,opt,name=challengeIndex" json:"challengeIndex,omitempty"`
	Response         *core.Challenge     `protobuf:"bytes,3,opt,name=response" json:"response,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *UpdateAuthorizationRequest) Reset()                    { *m = UpdateAuthorizationRequest{} }
func (m *UpdateAuthorizationRequest) String() string            { return proto1.CompactTextString(m) }
func (*UpdateAuthorizationRequest) ProtoMessage()               {}
func (*UpdateAuthorizationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateAuthorizationRequest) GetAuthz() *core.Authorization {
	if m != nil {
		return m.Authz
	}
	return nil
}

func (m *UpdateAuthorizationRequest) GetChallengeIndex() int64 {
	if m != nil && m.ChallengeIndex != nil {
		return *m.ChallengeIndex
	}
	return 0
}

func (m *UpdateAuthorizationRequest) GetResponse() *core.Challenge {
	if m != nil {
		return m.Response
	}
	return nil
}

type RevokeCertificateWithRegRequest struct {
	Cert             []byte `protobuf:"bytes,1,opt,name=cert" json:"cert,omitempty"`
	Code             *int64 `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	RegID            *int64 `protobuf:"varint,3,opt,name=regID" json:"regID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RevokeCertificateWithRegRequest) Reset()                    { *m = RevokeCertificateWithRegRequest{} }
func (m *RevokeCertificateWithRegRequest) String() string            { return proto1.CompactTextString(m) }
func (*RevokeCertificateWithRegRequest) ProtoMessage()               {}
func (*RevokeCertificateWithRegRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RevokeCertificateWithRegRequest) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

func (m *RevokeCertificateWithRegRequest) GetCode() int64 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *RevokeCertificateWithRegRequest) GetRegID() int64 {
	if m != nil && m.RegID != nil {
		return *m.RegID
	}
	return 0
}

type AdministrativelyRevokeCertificateRequest struct {
	Cert             []byte  `protobuf:"bytes,1,opt,name=cert" json:"cert,omitempty"`
	Code             *int64  `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	AdminName        *string `protobuf:"bytes,3,opt,name=adminName" json:"adminName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AdministrativelyRevokeCertificateRequest) Reset() {
	*m = AdministrativelyRevokeCertificateRequest{}
}
func (m *AdministrativelyRevokeCertificateRequest) String() string { return proto1.CompactTextString(m) }
func (*AdministrativelyRevokeCertificateRequest) ProtoMessage()    {}
func (*AdministrativelyRevokeCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5}
}

func (m *AdministrativelyRevokeCertificateRequest) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

func (m *AdministrativelyRevokeCertificateRequest) GetCode() int64 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *AdministrativelyRevokeCertificateRequest) GetAdminName() string {
	if m != nil && m.AdminName != nil {
		return *m.AdminName
	}
	return ""
}

func init() {
	proto1.RegisterType((*NewAuthorizationRequest)(nil), "ra.NewAuthorizationRequest")
	proto1.RegisterType((*NewCertificateRequest)(nil), "ra.NewCertificateRequest")
	proto1.RegisterType((*UpdateRegistrationRequest)(nil), "ra.UpdateRegistrationRequest")
	proto1.RegisterType((*UpdateAuthorizationRequest)(nil), "ra.UpdateAuthorizationRequest")
	proto1.RegisterType((*RevokeCertificateWithRegRequest)(nil), "ra.RevokeCertificateWithRegRequest")
	proto1.RegisterType((*AdministrativelyRevokeCertificateRequest)(nil), "ra.AdministrativelyRevokeCertificateRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for RegistrationAuthority service

type RegistrationAuthorityClient interface {
	NewRegistration(ctx context.Context, in *core.Registration, opts ...grpc.CallOption) (*core.Registration, error)
	NewAuthorization(ctx context.Context, in *NewAuthorizationRequest, opts ...grpc.CallOption) (*core.Authorization, error)
	NewCertificate(ctx context.Context, in *NewCertificateRequest, opts ...grpc.CallOption) (*core.Certificate, error)
	UpdateRegistration(ctx context.Context, in *UpdateRegistrationRequest, opts ...grpc.CallOption) (*core.Registration, error)
	UpdateAuthorization(ctx context.Context, in *UpdateAuthorizationRequest, opts ...grpc.CallOption) (*core.Authorization, error)
	RevokeCertificateWithReg(ctx context.Context, in *RevokeCertificateWithRegRequest, opts ...grpc.CallOption) (*core.Empty, error)
	DeactivateRegistration(ctx context.Context, in *core.Registration, opts ...grpc.CallOption) (*core.Empty, error)
	DeactivateAuthorization(ctx context.Context, in *core.Authorization, opts ...grpc.CallOption) (*core.Empty, error)
	AdministrativelyRevokeCertificate(ctx context.Context, in *AdministrativelyRevokeCertificateRequest, opts ...grpc.CallOption) (*core.Empty, error)
}

type registrationAuthorityClient struct {
	cc *grpc.ClientConn
}

func NewRegistrationAuthorityClient(cc *grpc.ClientConn) RegistrationAuthorityClient {
	return &registrationAuthorityClient{cc}
}

func (c *registrationAuthorityClient) NewRegistration(ctx context.Context, in *core.Registration, opts ...grpc.CallOption) (*core.Registration, error) {
	out := new(core.Registration)
	err := grpc.Invoke(ctx, "/ra.RegistrationAuthority/NewRegistration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) NewAuthorization(ctx context.Context, in *NewAuthorizationRequest, opts ...grpc.CallOption) (*core.Authorization, error) {
	out := new(core.Authorization)
	err := grpc.Invoke(ctx, "/ra.RegistrationAuthority/NewAuthorization", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) NewCertificate(ctx context.Context, in *NewCertificateRequest, opts ...grpc.CallOption) (*core.Certificate, error) {
	out := new(core.Certificate)
	err := grpc.Invoke(ctx, "/ra.RegistrationAuthority/NewCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) UpdateRegistration(ctx context.Context, in *UpdateRegistrationRequest, opts ...grpc.CallOption) (*core.Registration, error) {
	out := new(core.Registration)
	err := grpc.Invoke(ctx, "/ra.RegistrationAuthority/UpdateRegistration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) UpdateAuthorization(ctx context.Context, in *UpdateAuthorizationRequest, opts ...grpc.CallOption) (*core.Authorization, error) {
	out := new(core.Authorization)
	err := grpc.Invoke(ctx, "/ra.RegistrationAuthority/UpdateAuthorization", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) RevokeCertificateWithReg(ctx context.Context, in *RevokeCertificateWithRegRequest, opts ...grpc.CallOption) (*core.Empty, error) {
	out := new(core.Empty)
	err := grpc.Invoke(ctx, "/ra.RegistrationAuthority/RevokeCertificateWithReg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) DeactivateRegistration(ctx context.Context, in *core.Registration, opts ...grpc.CallOption) (*core.Empty, error) {
	out := new(core.Empty)
	err := grpc.Invoke(ctx, "/ra.RegistrationAuthority/DeactivateRegistration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) DeactivateAuthorization(ctx context.Context, in *core.Authorization, opts ...grpc.CallOption) (*core.Empty, error) {
	out := new(core.Empty)
	err := grpc.Invoke(ctx, "/ra.RegistrationAuthority/DeactivateAuthorization", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) AdministrativelyRevokeCertificate(ctx context.Context, in *AdministrativelyRevokeCertificateRequest, opts ...grpc.CallOption) (*core.Empty, error) {
	out := new(core.Empty)
	err := grpc.Invoke(ctx, "/ra.RegistrationAuthority/AdministrativelyRevokeCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RegistrationAuthority service

type RegistrationAuthorityServer interface {
	NewRegistration(context.Context, *core.Registration) (*core.Registration, error)
	NewAuthorization(context.Context, *NewAuthorizationRequest) (*core.Authorization, error)
	NewCertificate(context.Context, *NewCertificateRequest) (*core.Certificate, error)
	UpdateRegistration(context.Context, *UpdateRegistrationRequest) (*core.Registration, error)
	UpdateAuthorization(context.Context, *UpdateAuthorizationRequest) (*core.Authorization, error)
	RevokeCertificateWithReg(context.Context, *RevokeCertificateWithRegRequest) (*core.Empty, error)
	DeactivateRegistration(context.Context, *core.Registration) (*core.Empty, error)
	DeactivateAuthorization(context.Context, *core.Authorization) (*core.Empty, error)
	AdministrativelyRevokeCertificate(context.Context, *AdministrativelyRevokeCertificateRequest) (*core.Empty, error)
}

func RegisterRegistrationAuthorityServer(s *grpc.Server, srv RegistrationAuthorityServer) {
	s.RegisterService(&_RegistrationAuthority_serviceDesc, srv)
}

func _RegistrationAuthority_NewRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Registration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).NewRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/NewRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).NewRegistration(ctx, req.(*core.Registration))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_NewAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAuthorizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).NewAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/NewAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).NewAuthorization(ctx, req.(*NewAuthorizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_NewCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).NewCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/NewCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).NewCertificate(ctx, req.(*NewCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_UpdateRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).UpdateRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/UpdateRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).UpdateRegistration(ctx, req.(*UpdateRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_UpdateAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthorizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).UpdateAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/UpdateAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).UpdateAuthorization(ctx, req.(*UpdateAuthorizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_RevokeCertificateWithReg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeCertificateWithRegRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).RevokeCertificateWithReg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/RevokeCertificateWithReg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).RevokeCertificateWithReg(ctx, req.(*RevokeCertificateWithRegRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_DeactivateRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Registration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).DeactivateRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/DeactivateRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).DeactivateRegistration(ctx, req.(*core.Registration))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_DeactivateAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Authorization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).DeactivateAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/DeactivateAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).DeactivateAuthorization(ctx, req.(*core.Authorization))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_AdministrativelyRevokeCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdministrativelyRevokeCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).AdministrativelyRevokeCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/AdministrativelyRevokeCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).AdministrativelyRevokeCertificate(ctx, req.(*AdministrativelyRevokeCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegistrationAuthority_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ra.RegistrationAuthority",
	HandlerType: (*RegistrationAuthorityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewRegistration",
			Handler:    _RegistrationAuthority_NewRegistration_Handler,
		},
		{
			MethodName: "NewAuthorization",
			Handler:    _RegistrationAuthority_NewAuthorization_Handler,
		},
		{
			MethodName: "NewCertificate",
			Handler:    _RegistrationAuthority_NewCertificate_Handler,
		},
		{
			MethodName: "UpdateRegistration",
			Handler:    _RegistrationAuthority_UpdateRegistration_Handler,
		},
		{
			MethodName: "UpdateAuthorization",
			Handler:    _RegistrationAuthority_UpdateAuthorization_Handler,
		},
		{
			MethodName: "RevokeCertificateWithReg",
			Handler:    _RegistrationAuthority_RevokeCertificateWithReg_Handler,
		},
		{
			MethodName: "DeactivateRegistration",
			Handler:    _RegistrationAuthority_DeactivateRegistration_Handler,
		},
		{
			MethodName: "DeactivateAuthorization",
			Handler:    _RegistrationAuthority_DeactivateAuthorization_Handler,
		},
		{
			MethodName: "AdministrativelyRevokeCertificate",
			Handler:    _RegistrationAuthority_AdministrativelyRevokeCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto1.RegisterFile("ra/proto/ra.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x14, 0x4c, 0x9a, 0x06, 0xe8, 0x4b, 0x69, 0xc9, 0xab, 0xd2, 0xa6, 0x46, 0x40, 0xba, 0x5c, 0x72,
	0x40, 0xa9, 0x54, 0x8e, 0x15, 0x12, 0xa5, 0x05, 0x29, 0x52, 0xe4, 0x43, 0x24, 0x84, 0xe0, 0xc4,
	0xe2, 0x3c, 0xe2, 0x15, 0x89, 0xd7, 0xac, 0xd7, 0x09, 0x29, 0xdf, 0xc3, 0x7f, 0x22, 0xaf, 0x37,
	0x24, 0x76, 0x6c, 0x35, 0xf4, 0xb6, 0xd6, 0xbe, 0x99, 0x9d, 0x99, 0x37, 0x32, 0x34, 0x15, 0x3f,
	0x0f, 0x95, 0xd4, 0xf2, 0x5c, 0xf1, 0x9e, 0x39, 0xe0, 0x8e, 0xe2, 0x4e, 0xcb, 0x93, 0x8a, 0xec,
	0x45, 0x72, 0x4c, 0xaf, 0xd8, 0x00, 0x4e, 0x5c, 0x9a, 0x5f, 0xc5, 0xda, 0x97, 0x4a, 0xdc, 0x72,
	0x2d, 0x64, 0x30, 0xa4, 0x9f, 0x31, 0x45, 0x1a, 0x19, 0xd4, 0x79, 0xac, 0xfd, 0xdb, 0x76, 0xb5,
	0x53, 0xed, 0x36, 0x2e, 0x8e, 0x7a, 0x06, 0x96, 0x19, 0xc5, 0xc7, 0x50, 0x57, 0x34, 0xee, 0xdf,
	0xb4, 0x77, 0x3a, 0xd5, 0x6e, 0x8d, 0xbd, 0x86, 0x96, 0x4b, 0xf3, 0x6b, 0x52, 0x5a, 0x7c, 0x17,
	0x1e, 0xd7, 0xb4, 0xe4, 0x6a, 0x40, 0xcd, 0x8b, 0x94, 0x61, 0xda, 0xcf, 0x83, 0x38, 0x9c, 0x7e,
	0x0c, 0x47, 0x66, 0x78, 0x2c, 0x22, 0xad, 0x32, 0x22, 0x3a, 0xb0, 0xfb, 0x8d, 0x47, 0x64, 0x35,
	0x60, 0xaa, 0x61, 0x7d, 0x10, 0x19, 0x3c, 0x88, 0x0d, 0xdc, 0xd0, 0x15, 0xce, 0xb0, 0xdf, 0xe0,
	0xa4, 0x4f, 0xdc, 0xdb, 0xe8, 0x31, 0x1c, 0x78, 0x3e, 0x9f, 0x4c, 0x28, 0x18, 0x53, 0x3f, 0x18,
	0xd1, 0xaf, 0x54, 0x3c, 0x9e, 0xc1, 0x23, 0x45, 0x51, 0x28, 0x83, 0x88, 0xda, 0x35, 0x03, 0x3f,
	0x4c, 0xe1, 0xd7, 0xcb, 0x69, 0x36, 0x80, 0x17, 0x43, 0x9a, 0xc9, 0x1f, 0xb4, 0x96, 0xcb, 0x27,
	0xa1, 0xfd, 0x21, 0x8d, 0x97, 0x0a, 0xf6, 0x61, 0xd7, 0x23, 0xa5, 0x6d, 0x3e, 0xc9, 0x97, 0x1c,
	0x91, 0x7d, 0xe1, 0x5f, 0x5a, 0x35, 0x93, 0xd6, 0x67, 0xe8, 0x5e, 0x8d, 0xa6, 0x22, 0xb0, 0xe6,
	0x66, 0x34, 0x59, 0x6c, 0xb0, 0x6f, 0x43, 0xdb, 0x84, 0x3d, 0x9e, 0xf0, 0xb8, 0x7c, 0x9a, 0x2a,
	0xdf, 0xbb, 0xf8, 0x53, 0x87, 0xd6, 0x7a, 0x6c, 0x36, 0x01, 0xbd, 0xc0, 0x4b, 0x38, 0x74, 0x69,
	0x9e, 0x89, 0xbd, 0x20, 0x66, 0xa7, 0x28, 0xfa, 0x0a, 0x7e, 0x80, 0x27, 0xf9, 0x8a, 0xe1, 0xd3,
	0x9e, 0xe2, 0xbd, 0x92, 0xe2, 0x39, 0x45, 0x0b, 0x60, 0x15, 0x7c, 0x0b, 0x07, 0xd9, 0x72, 0xe1,
	0xa9, 0x65, 0xd9, 0xb4, 0xee, 0x34, 0xed, 0x16, 0x56, 0x37, 0xac, 0x82, 0x7d, 0xc0, 0xcd, 0xa6,
	0xe1, 0xb3, 0x84, 0xa5, 0xb4, 0x81, 0x25, 0xa6, 0x06, 0x70, 0x54, 0xd0, 0x28, 0x7c, 0xbe, 0xe2,
	0xfa, 0x1f, 0x6b, 0x2e, 0xb4, 0xcb, 0x2a, 0x82, 0x2f, 0x13, 0xca, 0x3b, 0x0a, 0xe4, 0x34, 0x52,
	0xde, 0xf7, 0xd3, 0x50, 0x2f, 0x58, 0x05, 0x2f, 0xe1, 0xf8, 0x86, 0xb8, 0xa7, 0xc5, 0x2c, 0x6f,
	0xb6, 0x68, 0x6d, 0x39, 0xf0, 0x1b, 0x38, 0x59, 0x81, 0xb3, 0xf6, 0x8a, 0xe4, 0xe7, 0xe1, 0x5f,
	0xe1, 0xec, 0xce, 0x82, 0xe2, 0xab, 0xc4, 0xd4, 0xb6, 0x3d, 0xce, 0xbd, 0xf0, 0xee, 0xe1, 0x97,
	0xba, 0xf9, 0x79, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x14, 0xc6, 0x23, 0xeb, 0x04, 0x00,
	0x00,
}
