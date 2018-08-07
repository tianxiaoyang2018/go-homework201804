// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contact/contact.service.proto

package contact

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetBidirectionSecretCrushUserIdsRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *GetBidirectionSecretCrushUserIdsRequest) Reset() {
	*m = GetBidirectionSecretCrushUserIdsRequest{}
}
func (m *GetBidirectionSecretCrushUserIdsRequest) String() string { return proto.CompactTextString(m) }
func (*GetBidirectionSecretCrushUserIdsRequest) ProtoMessage()    {}
func (*GetBidirectionSecretCrushUserIdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0}
}

func (m *GetBidirectionSecretCrushUserIdsRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetBidirectionSecretCrushUserIdsReply struct {
	UserIds []string `protobuf:"bytes,1,rep,name=userIds" json:"userIds,omitempty"`
}

func (m *GetBidirectionSecretCrushUserIdsReply) Reset()         { *m = GetBidirectionSecretCrushUserIdsReply{} }
func (m *GetBidirectionSecretCrushUserIdsReply) String() string { return proto.CompactTextString(m) }
func (*GetBidirectionSecretCrushUserIdsReply) ProtoMessage()    {}
func (*GetBidirectionSecretCrushUserIdsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1}
}

func (m *GetBidirectionSecretCrushUserIdsReply) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type UncrushCrushesRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *UncrushCrushesRequest) Reset()                    { *m = UncrushCrushesRequest{} }
func (m *UncrushCrushesRequest) String() string            { return proto.CompactTextString(m) }
func (*UncrushCrushesRequest) ProtoMessage()               {}
func (*UncrushCrushesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *UncrushCrushesRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type UncrushCrushesReply struct {
}

func (m *UncrushCrushesReply) Reset()                    { *m = UncrushCrushesReply{} }
func (m *UncrushCrushesReply) String() string            { return proto.CompactTextString(m) }
func (*UncrushCrushesReply) ProtoMessage()               {}
func (*UncrushCrushesReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type UpsertRequest struct {
	Contacts []*Contact `protobuf:"bytes,1,rep,name=contacts" json:"contacts,omitempty"`
}

func (m *UpsertRequest) Reset()                    { *m = UpsertRequest{} }
func (m *UpsertRequest) String() string            { return proto.CompactTextString(m) }
func (*UpsertRequest) ProtoMessage()               {}
func (*UpsertRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *UpsertRequest) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type UpsertReply struct {
	Contacts []*Contact `protobuf:"bytes,1,rep,name=contacts" json:"contacts,omitempty"`
}

func (m *UpsertReply) Reset()                    { *m = UpsertReply{} }
func (m *UpsertReply) String() string            { return proto.CompactTextString(m) }
func (*UpsertReply) ProtoMessage()               {}
func (*UpsertReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *UpsertReply) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type SelectMobileContactHashesCountRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *SelectMobileContactHashesCountRequest) Reset()         { *m = SelectMobileContactHashesCountRequest{} }
func (m *SelectMobileContactHashesCountRequest) String() string { return proto.CompactTextString(m) }
func (*SelectMobileContactHashesCountRequest) ProtoMessage()    {}
func (*SelectMobileContactHashesCountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{6}
}

func (m *SelectMobileContactHashesCountRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type SelectMobileContactHashesCountReply struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *SelectMobileContactHashesCountReply) Reset()         { *m = SelectMobileContactHashesCountReply{} }
func (m *SelectMobileContactHashesCountReply) String() string { return proto.CompactTextString(m) }
func (*SelectMobileContactHashesCountReply) ProtoMessage()    {}
func (*SelectMobileContactHashesCountReply) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{7}
}

func (m *SelectMobileContactHashesCountReply) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type FindSecretCrushesRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *FindSecretCrushesRequest) Reset()                    { *m = FindSecretCrushesRequest{} }
func (m *FindSecretCrushesRequest) String() string            { return proto.CompactTextString(m) }
func (*FindSecretCrushesRequest) ProtoMessage()               {}
func (*FindSecretCrushesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *FindSecretCrushesRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type FindSecretCrushesReply struct {
	Contacts []*Contact `protobuf:"bytes,1,rep,name=contacts" json:"contacts,omitempty"`
}

func (m *FindSecretCrushesReply) Reset()                    { *m = FindSecretCrushesReply{} }
func (m *FindSecretCrushesReply) String() string            { return proto.CompactTextString(m) }
func (*FindSecretCrushesReply) ProtoMessage()               {}
func (*FindSecretCrushesReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *FindSecretCrushesReply) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type FindAsUserIdsRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *FindAsUserIdsRequest) Reset()                    { *m = FindAsUserIdsRequest{} }
func (m *FindAsUserIdsRequest) String() string            { return proto.CompactTextString(m) }
func (*FindAsUserIdsRequest) ProtoMessage()               {}
func (*FindAsUserIdsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *FindAsUserIdsRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type FindAsUserIdsReply struct {
	UserIds []string `protobuf:"bytes,1,rep,name=userIds" json:"userIds,omitempty"`
}

func (m *FindAsUserIdsReply) Reset()                    { *m = FindAsUserIdsReply{} }
func (m *FindAsUserIdsReply) String() string            { return proto.CompactTextString(m) }
func (*FindAsUserIdsReply) ProtoMessage()               {}
func (*FindAsUserIdsReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *FindAsUserIdsReply) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type FindInversedAsUserIdsRequest struct {
	UserId string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *FindInversedAsUserIdsRequest) Reset()                    { *m = FindInversedAsUserIdsRequest{} }
func (m *FindInversedAsUserIdsRequest) String() string            { return proto.CompactTextString(m) }
func (*FindInversedAsUserIdsRequest) ProtoMessage()               {}
func (*FindInversedAsUserIdsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *FindInversedAsUserIdsRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type FindInversedAsUserIdsReply struct {
	UserIds []string `protobuf:"bytes,1,rep,name=userIds" json:"userIds,omitempty"`
}

func (m *FindInversedAsUserIdsReply) Reset()                    { *m = FindInversedAsUserIdsReply{} }
func (m *FindInversedAsUserIdsReply) String() string            { return proto.CompactTextString(m) }
func (*FindInversedAsUserIdsReply) ProtoMessage()               {}
func (*FindInversedAsUserIdsReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *FindInversedAsUserIdsReply) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

type FindByIdsRequest struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *FindByIdsRequest) Reset()                    { *m = FindByIdsRequest{} }
func (m *FindByIdsRequest) String() string            { return proto.CompactTextString(m) }
func (*FindByIdsRequest) ProtoMessage()               {}
func (*FindByIdsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *FindByIdsRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type FindByIdsReply struct {
	Contacts []*Contact `protobuf:"bytes,1,rep,name=contacts" json:"contacts,omitempty"`
}

func (m *FindByIdsReply) Reset()                    { *m = FindByIdsReply{} }
func (m *FindByIdsReply) String() string            { return proto.CompactTextString(m) }
func (*FindByIdsReply) ProtoMessage()               {}
func (*FindByIdsReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *FindByIdsReply) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type FindMobileHomeLocationRequest struct {
	Phone string `protobuf:"bytes,1,opt,name=phone" json:"phone,omitempty"`
}

func (m *FindMobileHomeLocationRequest) Reset()                    { *m = FindMobileHomeLocationRequest{} }
func (m *FindMobileHomeLocationRequest) String() string            { return proto.CompactTextString(m) }
func (*FindMobileHomeLocationRequest) ProtoMessage()               {}
func (*FindMobileHomeLocationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{16} }

func (m *FindMobileHomeLocationRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type FindMobileHomeLocationReply struct {
	MobileArea string `protobuf:"bytes,1,opt,name=mobileArea" json:"mobileArea,omitempty"`
}

func (m *FindMobileHomeLocationReply) Reset()                    { *m = FindMobileHomeLocationReply{} }
func (m *FindMobileHomeLocationReply) String() string            { return proto.CompactTextString(m) }
func (*FindMobileHomeLocationReply) ProtoMessage()               {}
func (*FindMobileHomeLocationReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{17} }

func (m *FindMobileHomeLocationReply) GetMobileArea() string {
	if m != nil {
		return m.MobileArea
	}
	return ""
}

type FindContactUserIdsRequest struct {
	Hashes []string `protobuf:"bytes,1,rep,name=hashes" json:"hashes,omitempty"`
}

func (m *FindContactUserIdsRequest) Reset()                    { *m = FindContactUserIdsRequest{} }
func (m *FindContactUserIdsRequest) String() string            { return proto.CompactTextString(m) }
func (*FindContactUserIdsRequest) ProtoMessage()               {}
func (*FindContactUserIdsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{18} }

func (m *FindContactUserIdsRequest) GetHashes() []string {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type FindContactUserIdsReply struct {
	UserIds []string `protobuf:"bytes,1,rep,name=userIds" json:"userIds,omitempty"`
}

func (m *FindContactUserIdsReply) Reset()                    { *m = FindContactUserIdsReply{} }
func (m *FindContactUserIdsReply) String() string            { return proto.CompactTextString(m) }
func (*FindContactUserIdsReply) ProtoMessage()               {}
func (*FindContactUserIdsReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{19} }

func (m *FindContactUserIdsReply) GetUserIds() []string {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func init() {
	proto.RegisterType((*GetBidirectionSecretCrushUserIdsRequest)(nil), "contact.GetBidirectionSecretCrushUserIdsRequest")
	proto.RegisterType((*GetBidirectionSecretCrushUserIdsReply)(nil), "contact.GetBidirectionSecretCrushUserIdsReply")
	proto.RegisterType((*UncrushCrushesRequest)(nil), "contact.UncrushCrushesRequest")
	proto.RegisterType((*UncrushCrushesReply)(nil), "contact.UncrushCrushesReply")
	proto.RegisterType((*UpsertRequest)(nil), "contact.UpsertRequest")
	proto.RegisterType((*UpsertReply)(nil), "contact.UpsertReply")
	proto.RegisterType((*SelectMobileContactHashesCountRequest)(nil), "contact.SelectMobileContactHashesCountRequest")
	proto.RegisterType((*SelectMobileContactHashesCountReply)(nil), "contact.SelectMobileContactHashesCountReply")
	proto.RegisterType((*FindSecretCrushesRequest)(nil), "contact.FindSecretCrushesRequest")
	proto.RegisterType((*FindSecretCrushesReply)(nil), "contact.FindSecretCrushesReply")
	proto.RegisterType((*FindAsUserIdsRequest)(nil), "contact.FindAsUserIdsRequest")
	proto.RegisterType((*FindAsUserIdsReply)(nil), "contact.FindAsUserIdsReply")
	proto.RegisterType((*FindInversedAsUserIdsRequest)(nil), "contact.FindInversedAsUserIdsRequest")
	proto.RegisterType((*FindInversedAsUserIdsReply)(nil), "contact.FindInversedAsUserIdsReply")
	proto.RegisterType((*FindByIdsRequest)(nil), "contact.FindByIdsRequest")
	proto.RegisterType((*FindByIdsReply)(nil), "contact.FindByIdsReply")
	proto.RegisterType((*FindMobileHomeLocationRequest)(nil), "contact.FindMobileHomeLocationRequest")
	proto.RegisterType((*FindMobileHomeLocationReply)(nil), "contact.FindMobileHomeLocationReply")
	proto.RegisterType((*FindContactUserIdsRequest)(nil), "contact.FindContactUserIdsRequest")
	proto.RegisterType((*FindContactUserIdsReply)(nil), "contact.FindContactUserIdsReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ContactService service

type ContactServiceClient interface {
	GetBidirectionSecretCrushUserIds(ctx context.Context, in *GetBidirectionSecretCrushUserIdsRequest, opts ...grpc.CallOption) (*GetBidirectionSecretCrushUserIdsReply, error)
	UncrushCrushes(ctx context.Context, in *UncrushCrushesRequest, opts ...grpc.CallOption) (*UncrushCrushesReply, error)
	Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*UpsertReply, error)
	SelectMobileContactHashesCount(ctx context.Context, in *SelectMobileContactHashesCountRequest, opts ...grpc.CallOption) (*SelectMobileContactHashesCountReply, error)
	FindSecretCrushes(ctx context.Context, in *FindSecretCrushesRequest, opts ...grpc.CallOption) (*FindSecretCrushesReply, error)
	FindAsUserIds(ctx context.Context, in *FindAsUserIdsRequest, opts ...grpc.CallOption) (*FindAsUserIdsReply, error)
	FindInversedAsUserIds(ctx context.Context, in *FindInversedAsUserIdsRequest, opts ...grpc.CallOption) (*FindInversedAsUserIdsReply, error)
	FindByIds(ctx context.Context, in *FindByIdsRequest, opts ...grpc.CallOption) (*FindByIdsReply, error)
	FindMobileHomeLocation(ctx context.Context, in *FindMobileHomeLocationRequest, opts ...grpc.CallOption) (*FindMobileHomeLocationReply, error)
	FindContactUserIds(ctx context.Context, in *FindContactUserIdsRequest, opts ...grpc.CallOption) (*FindContactUserIdsReply, error)
}

type contactServiceClient struct {
	cc *grpc.ClientConn
}

func NewContactServiceClient(cc *grpc.ClientConn) ContactServiceClient {
	return &contactServiceClient{cc}
}

func (c *contactServiceClient) GetBidirectionSecretCrushUserIds(ctx context.Context, in *GetBidirectionSecretCrushUserIdsRequest, opts ...grpc.CallOption) (*GetBidirectionSecretCrushUserIdsReply, error) {
	out := new(GetBidirectionSecretCrushUserIdsReply)
	err := grpc.Invoke(ctx, "/contact.ContactService/GetBidirectionSecretCrushUserIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) UncrushCrushes(ctx context.Context, in *UncrushCrushesRequest, opts ...grpc.CallOption) (*UncrushCrushesReply, error) {
	out := new(UncrushCrushesReply)
	err := grpc.Invoke(ctx, "/contact.ContactService/UncrushCrushes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*UpsertReply, error) {
	out := new(UpsertReply)
	err := grpc.Invoke(ctx, "/contact.ContactService/Upsert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) SelectMobileContactHashesCount(ctx context.Context, in *SelectMobileContactHashesCountRequest, opts ...grpc.CallOption) (*SelectMobileContactHashesCountReply, error) {
	out := new(SelectMobileContactHashesCountReply)
	err := grpc.Invoke(ctx, "/contact.ContactService/SelectMobileContactHashesCount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) FindSecretCrushes(ctx context.Context, in *FindSecretCrushesRequest, opts ...grpc.CallOption) (*FindSecretCrushesReply, error) {
	out := new(FindSecretCrushesReply)
	err := grpc.Invoke(ctx, "/contact.ContactService/FindSecretCrushes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) FindAsUserIds(ctx context.Context, in *FindAsUserIdsRequest, opts ...grpc.CallOption) (*FindAsUserIdsReply, error) {
	out := new(FindAsUserIdsReply)
	err := grpc.Invoke(ctx, "/contact.ContactService/FindAsUserIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) FindInversedAsUserIds(ctx context.Context, in *FindInversedAsUserIdsRequest, opts ...grpc.CallOption) (*FindInversedAsUserIdsReply, error) {
	out := new(FindInversedAsUserIdsReply)
	err := grpc.Invoke(ctx, "/contact.ContactService/FindInversedAsUserIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) FindByIds(ctx context.Context, in *FindByIdsRequest, opts ...grpc.CallOption) (*FindByIdsReply, error) {
	out := new(FindByIdsReply)
	err := grpc.Invoke(ctx, "/contact.ContactService/FindByIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) FindMobileHomeLocation(ctx context.Context, in *FindMobileHomeLocationRequest, opts ...grpc.CallOption) (*FindMobileHomeLocationReply, error) {
	out := new(FindMobileHomeLocationReply)
	err := grpc.Invoke(ctx, "/contact.ContactService/FindMobileHomeLocation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactServiceClient) FindContactUserIds(ctx context.Context, in *FindContactUserIdsRequest, opts ...grpc.CallOption) (*FindContactUserIdsReply, error) {
	out := new(FindContactUserIdsReply)
	err := grpc.Invoke(ctx, "/contact.ContactService/FindContactUserIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContactService service

type ContactServiceServer interface {
	GetBidirectionSecretCrushUserIds(context.Context, *GetBidirectionSecretCrushUserIdsRequest) (*GetBidirectionSecretCrushUserIdsReply, error)
	UncrushCrushes(context.Context, *UncrushCrushesRequest) (*UncrushCrushesReply, error)
	Upsert(context.Context, *UpsertRequest) (*UpsertReply, error)
	SelectMobileContactHashesCount(context.Context, *SelectMobileContactHashesCountRequest) (*SelectMobileContactHashesCountReply, error)
	FindSecretCrushes(context.Context, *FindSecretCrushesRequest) (*FindSecretCrushesReply, error)
	FindAsUserIds(context.Context, *FindAsUserIdsRequest) (*FindAsUserIdsReply, error)
	FindInversedAsUserIds(context.Context, *FindInversedAsUserIdsRequest) (*FindInversedAsUserIdsReply, error)
	FindByIds(context.Context, *FindByIdsRequest) (*FindByIdsReply, error)
	FindMobileHomeLocation(context.Context, *FindMobileHomeLocationRequest) (*FindMobileHomeLocationReply, error)
	FindContactUserIds(context.Context, *FindContactUserIdsRequest) (*FindContactUserIdsReply, error)
}

func RegisterContactServiceServer(s *grpc.Server, srv ContactServiceServer) {
	s.RegisterService(&_ContactService_serviceDesc, srv)
}

func _ContactService_GetBidirectionSecretCrushUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBidirectionSecretCrushUserIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).GetBidirectionSecretCrushUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/GetBidirectionSecretCrushUserIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).GetBidirectionSecretCrushUserIds(ctx, req.(*GetBidirectionSecretCrushUserIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_UncrushCrushes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UncrushCrushesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).UncrushCrushes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/UncrushCrushes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).UncrushCrushes(ctx, req.(*UncrushCrushesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).Upsert(ctx, req.(*UpsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_SelectMobileContactHashesCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectMobileContactHashesCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).SelectMobileContactHashesCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/SelectMobileContactHashesCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).SelectMobileContactHashesCount(ctx, req.(*SelectMobileContactHashesCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_FindSecretCrushes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindSecretCrushesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).FindSecretCrushes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/FindSecretCrushes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).FindSecretCrushes(ctx, req.(*FindSecretCrushesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_FindAsUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAsUserIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).FindAsUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/FindAsUserIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).FindAsUserIds(ctx, req.(*FindAsUserIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_FindInversedAsUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindInversedAsUserIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).FindInversedAsUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/FindInversedAsUserIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).FindInversedAsUserIds(ctx, req.(*FindInversedAsUserIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_FindByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).FindByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/FindByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).FindByIds(ctx, req.(*FindByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_FindMobileHomeLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindMobileHomeLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).FindMobileHomeLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/FindMobileHomeLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).FindMobileHomeLocation(ctx, req.(*FindMobileHomeLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactService_FindContactUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindContactUserIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServiceServer).FindContactUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contact.ContactService/FindContactUserIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServiceServer).FindContactUserIds(ctx, req.(*FindContactUserIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContactService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contact.ContactService",
	HandlerType: (*ContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBidirectionSecretCrushUserIds",
			Handler:    _ContactService_GetBidirectionSecretCrushUserIds_Handler,
		},
		{
			MethodName: "UncrushCrushes",
			Handler:    _ContactService_UncrushCrushes_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _ContactService_Upsert_Handler,
		},
		{
			MethodName: "SelectMobileContactHashesCount",
			Handler:    _ContactService_SelectMobileContactHashesCount_Handler,
		},
		{
			MethodName: "FindSecretCrushes",
			Handler:    _ContactService_FindSecretCrushes_Handler,
		},
		{
			MethodName: "FindAsUserIds",
			Handler:    _ContactService_FindAsUserIds_Handler,
		},
		{
			MethodName: "FindInversedAsUserIds",
			Handler:    _ContactService_FindInversedAsUserIds_Handler,
		},
		{
			MethodName: "FindByIds",
			Handler:    _ContactService_FindByIds_Handler,
		},
		{
			MethodName: "FindMobileHomeLocation",
			Handler:    _ContactService_FindMobileHomeLocation_Handler,
		},
		{
			MethodName: "FindContactUserIds",
			Handler:    _ContactService_FindContactUserIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contact/contact.service.proto",
}

func init() { proto.RegisterFile("contact/contact.service.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 642 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xcf, 0x4f, 0xdb, 0x30,
	0x14, 0xc7, 0x57, 0x21, 0x60, 0x3c, 0x04, 0x62, 0x5e, 0x0b, 0xc5, 0x40, 0xd7, 0x05, 0xba, 0x71,
	0x80, 0x74, 0x6b, 0x05, 0x9a, 0x84, 0xd8, 0xd4, 0x56, 0x62, 0x20, 0x0d, 0x69, 0x6a, 0xd5, 0xc3,
	0xa6, 0x5d, 0x52, 0xd7, 0x6a, 0x22, 0xb5, 0x76, 0x96, 0xb8, 0x48, 0xbd, 0xed, 0xb4, 0xbf, 0x6c,
	0x7f, 0xd8, 0xe4, 0x38, 0x09, 0x49, 0x9b, 0x1f, 0xed, 0x29, 0xb5, 0xfd, 0xfd, 0xbc, 0xe7, 0xf7,
	0x9a, 0xf7, 0x55, 0xe0, 0x84, 0x70, 0x26, 0x0c, 0x22, 0xea, 0xfe, 0x53, 0x77, 0xa9, 0xf3, 0x64,
	0x11, 0xaa, 0xdb, 0x0e, 0x17, 0x1c, 0x6d, 0xfa, 0xdb, 0xb8, 0x34, 0xaf, 0xf3, 0xce, 0xb5, 0x16,
	0xbc, 0xff, 0x4a, 0x45, 0xdb, 0x1a, 0x5a, 0x0e, 0x25, 0xc2, 0xe2, 0xac, 0x47, 0x89, 0x43, 0x45,
	0xc7, 0x99, 0xba, 0x66, 0xdf, 0xa5, 0xce, 0xc3, 0xd0, 0xed, 0xd2, 0xdf, 0x53, 0xea, 0x0a, 0xb4,
	0x0f, 0x1b, 0x53, 0x6f, 0xa7, 0x5c, 0xa8, 0x16, 0xce, 0xb7, 0xba, 0xfe, 0x4a, 0x6b, 0x41, 0x2d,
	0x3f, 0x84, 0x3d, 0x9e, 0xa1, 0x32, 0x6c, 0x2a, 0xc4, 0x2d, 0x17, 0xaa, 0x6b, 0xe7, 0x5b, 0xdd,
	0x60, 0xa9, 0xd5, 0xa1, 0xd4, 0x67, 0x44, 0x02, 0x1e, 0x45, 0x73, 0x73, 0x96, 0xe0, 0xf5, 0x3c,
	0x60, 0x8f, 0x67, 0xda, 0x2d, 0xec, 0xf4, 0x6d, 0x97, 0x3a, 0x22, 0xe0, 0x2f, 0xe0, 0xa5, 0x5f,
	0xaf, 0xca, 0xb9, 0xdd, 0xd8, 0xd3, 0x83, 0x06, 0x74, 0xd4, 0xb3, 0x1b, 0x2a, 0xb4, 0x1b, 0xd8,
	0x0e, 0x70, 0x79, 0xdf, 0xd5, 0xe0, 0x2f, 0x50, 0xeb, 0xd1, 0x31, 0x25, 0xe2, 0x91, 0x0f, 0xac,
	0x31, 0xf5, 0x05, 0xf7, 0x86, 0xbc, 0x5d, 0x87, 0x4f, 0x99, 0xc8, 0xab, 0xe9, 0x06, 0x4e, 0xf3,
	0x02, 0xc8, 0x5b, 0x15, 0x61, 0x9d, 0xc8, 0x95, 0x47, 0xaf, 0x75, 0xd5, 0x42, 0x6b, 0x40, 0xf9,
	0xce, 0x62, 0xc3, 0x48, 0xeb, 0xf3, 0x9b, 0x78, 0x07, 0xfb, 0x09, 0xcc, 0xea, 0x95, 0xeb, 0x50,
	0x94, 0x71, 0x5a, 0xee, 0x92, 0x2f, 0x8c, 0x0e, 0x68, 0x4e, 0x9f, 0xfd, 0x76, 0x5c, 0xc3, 0xb1,
	0xd4, 0x3f, 0xb0, 0x27, 0xea, 0xb8, 0x74, 0xf9, 0x3c, 0xd7, 0x80, 0x53, 0xb8, 0xec, 0x7c, 0x67,
	0xb0, 0x27, 0xb9, 0xf6, 0x2c, 0x92, 0x63, 0x0f, 0xd6, 0xac, 0x50, 0x29, 0x7f, 0x6a, 0x9f, 0x61,
	0x37, 0xa2, 0x5a, 0xbd, 0x6b, 0x57, 0x70, 0x22, 0x79, 0xf5, 0x67, 0xdf, 0xf3, 0x09, 0xfd, 0xc6,
	0x89, 0x21, 0xc7, 0x27, 0x48, 0x59, 0x84, 0x75, 0xdb, 0xe4, 0x8c, 0xfa, 0x55, 0xa9, 0x85, 0x76,
	0x0b, 0x47, 0x69, 0x98, 0xbc, 0x43, 0x05, 0x60, 0xe2, 0x1d, 0xb5, 0x1c, 0x6a, 0xf8, 0x64, 0x64,
	0x47, 0x6b, 0xc2, 0xa1, 0xc4, 0xfd, 0xeb, 0x2c, 0x36, 0xd2, 0xf4, 0x5e, 0x37, 0xbf, 0x4e, 0x7f,
	0xa5, 0x35, 0xe1, 0x20, 0x09, 0xca, 0xec, 0x62, 0xe3, 0xdf, 0x26, 0xec, 0xfa, 0x44, 0x4f, 0x59,
	0x12, 0xfa, 0x5b, 0x80, 0x6a, 0x9e, 0x55, 0xa0, 0x0f, 0x61, 0xcf, 0x96, 0x34, 0x26, 0xac, 0xaf,
	0x40, 0x48, 0x97, 0x78, 0x81, 0xbe, 0xc3, 0x6e, 0xdc, 0x3e, 0x50, 0x25, 0x8c, 0x91, 0x68, 0x44,
	0xf8, 0x38, 0xf5, 0x5c, 0x45, 0xfc, 0x04, 0x1b, 0xca, 0x3a, 0xd0, 0xfe, 0xb3, 0x32, 0x6a, 0x45,
	0xb8, 0xb8, 0xb0, 0xaf, 0xc8, 0x3f, 0x05, 0xa8, 0x64, 0xcf, 0x3d, 0x7a, 0x2e, 0x70, 0x29, 0x87,
	0xc1, 0x17, 0x4b, 0xeb, 0xd5, 0x15, 0x7e, 0xc0, 0xab, 0x05, 0x23, 0x40, 0x6f, 0xc3, 0x20, 0x69,
	0xc6, 0x82, 0xdf, 0x64, 0x49, 0x54, 0xe8, 0x47, 0xd8, 0x89, 0xcd, 0x3a, 0x3a, 0x89, 0x31, 0xf3,
	0xb3, 0x8c, 0x8f, 0xd2, 0x8e, 0x55, 0x38, 0x0a, 0xa5, 0xc4, 0x91, 0x46, 0xb5, 0x18, 0x97, 0x66,
	0x15, 0xf8, 0x34, 0x4f, 0xa6, 0xd2, 0xb4, 0x60, 0x2b, 0x9c, 0x6d, 0x74, 0x18, 0x63, 0xa2, 0xae,
	0x80, 0x0f, 0x92, 0x8e, 0x54, 0x08, 0x53, 0x99, 0xeb, 0xe2, 0x9c, 0xa2, 0x77, 0x31, 0x28, 0x75,
	0xfe, 0xf1, 0x59, 0xae, 0x4e, 0x65, 0xfa, 0xa5, 0xec, 0x34, 0x3e, 0x9d, 0x48, 0x8b, 0xd1, 0x89,
	0xf3, 0x8e, 0xab, 0x99, 0x1a, 0x2f, 0x7a, 0xfb, 0xea, 0x67, 0x73, 0x64, 0x09, 0x73, 0x3a, 0xd0,
	0x09, 0x9f, 0xd4, 0xed, 0x8f, 0x84, 0xd5, 0x85, 0xc1, 0x84, 0xc1, 0x2e, 0x87, 0x7c, 0x62, 0x58,
	0xec, 0xd2, 0x25, 0x26, 0x9d, 0x18, 0xf5, 0x11, 0x1f, 0x1b, 0x6c, 0x14, 0x7c, 0x5d, 0x0c, 0x36,
	0xbc, 0xcf, 0x8b, 0xe6, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x26, 0x61, 0x57, 0x9f, 0x08,
	0x00, 0x00,
}
