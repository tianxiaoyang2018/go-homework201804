// Code generated by protoc-gen-go. DO NOT EDIT.
// source: membership/membership.service.proto

package membership

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

type RevokeUserPrivilegeReply_CODE int32

const (
	RevokeUserPrivilegeReply_SuccCode       RevokeUserPrivilegeReply_CODE = 0
	RevokeUserPrivilegeReply_UnknownErrCode RevokeUserPrivilegeReply_CODE = -100
)

var RevokeUserPrivilegeReply_CODE_name = map[int32]string{
	0:    "SuccCode",
	-100: "UnknownErrCode",
}
var RevokeUserPrivilegeReply_CODE_value = map[string]int32{
	"SuccCode":       0,
	"UnknownErrCode": -100,
}

func (x RevokeUserPrivilegeReply_CODE) String() string {
	return proto.EnumName(RevokeUserPrivilegeReply_CODE_name, int32(x))
}
func (RevokeUserPrivilegeReply_CODE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 0}
}

type TmTecTranReply_CODE int32

const (
	TmTecTranReply_SuccCode               TmTecTranReply_CODE = 0
	TmTecTranReply_UnknownErrCode         TmTecTranReply_CODE = -100
	TmTecTranReply_ErrClosedCode          TmTecTranReply_CODE = -101
	TmTecTranReply_ErrNotCompleteCode     TmTecTranReply_CODE = -102
	TmTecTranReply_ErrAlreadyAcceptedCode TmTecTranReply_CODE = -103
	TmTecTranReply_ErrCannotGetLockCode   TmTecTranReply_CODE = -104
)

var TmTecTranReply_CODE_name = map[int32]string{
	0:    "SuccCode",
	-100: "UnknownErrCode",
	-101: "ErrClosedCode",
	-102: "ErrNotCompleteCode",
	-103: "ErrAlreadyAcceptedCode",
	-104: "ErrCannotGetLockCode",
}
var TmTecTranReply_CODE_value = map[string]int32{
	"SuccCode":               0,
	"UnknownErrCode":         -100,
	"ErrClosedCode":          -101,
	"ErrNotCompleteCode":     -102,
	"ErrAlreadyAcceptedCode": -103,
	"ErrCannotGetLockCode":   -104,
}

func (x TmTecTranReply_CODE) String() string {
	return proto.EnumName(TmTecTranReply_CODE_name, int32(x))
}
func (TmTecTranReply_CODE) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{11, 0} }

type RevokeUserPrivilegeRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *RevokeUserPrivilegeRequest) Reset()                    { *m = RevokeUserPrivilegeRequest{} }
func (m *RevokeUserPrivilegeRequest) String() string            { return proto.CompactTextString(m) }
func (*RevokeUserPrivilegeRequest) ProtoMessage()               {}
func (*RevokeUserPrivilegeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *RevokeUserPrivilegeRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type RevokeUserPrivilegeReply struct {
	Code RevokeUserPrivilegeReply_CODE `protobuf:"varint,1,opt,name=code,enum=membership.RevokeUserPrivilegeReply_CODE" json:"code,omitempty"`
}

func (m *RevokeUserPrivilegeReply) Reset()                    { *m = RevokeUserPrivilegeReply{} }
func (m *RevokeUserPrivilegeReply) String() string            { return proto.CompactTextString(m) }
func (*RevokeUserPrivilegeReply) ProtoMessage()               {}
func (*RevokeUserPrivilegeReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *RevokeUserPrivilegeReply) GetCode() RevokeUserPrivilegeReply_CODE {
	if m != nil {
		return m.Code
	}
	return RevokeUserPrivilegeReply_SuccCode
}

type UpdateUserProductPrivilegeRequest struct {
	UserID     string             `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	Privileges *ProductPrivileges `protobuf:"bytes,2,opt,name=privileges" json:"privileges,omitempty"`
	Expire     int64              `protobuf:"varint,3,opt,name=expire" json:"expire,omitempty"`
}

func (m *UpdateUserProductPrivilegeRequest) Reset()         { *m = UpdateUserProductPrivilegeRequest{} }
func (m *UpdateUserProductPrivilegeRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserProductPrivilegeRequest) ProtoMessage()    {}
func (*UpdateUserProductPrivilegeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{2}
}

func (m *UpdateUserProductPrivilegeRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UpdateUserProductPrivilegeRequest) GetPrivileges() *ProductPrivileges {
	if m != nil {
		return m.Privileges
	}
	return nil
}

func (m *UpdateUserProductPrivilegeRequest) GetExpire() int64 {
	if m != nil {
		return m.Expire
	}
	return 0
}

type ProductPrivilegesReply struct {
	Privileges *ProductPrivileges `protobuf:"bytes,1,opt,name=privileges" json:"privileges,omitempty"`
}

func (m *ProductPrivilegesReply) Reset()                    { *m = ProductPrivilegesReply{} }
func (m *ProductPrivilegesReply) String() string            { return proto.CompactTextString(m) }
func (*ProductPrivilegesReply) ProtoMessage()               {}
func (*ProductPrivilegesReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ProductPrivilegesReply) GetPrivileges() *ProductPrivileges {
	if m != nil {
		return m.Privileges
	}
	return nil
}

type UserPrivilegesReply struct {
	UserPrivilege []*UserPrivilege `protobuf:"bytes,1,rep,name=userPrivilege" json:"userPrivilege,omitempty"`
}

func (m *UserPrivilegesReply) Reset()                    { *m = UserPrivilegesReply{} }
func (m *UserPrivilegesReply) String() string            { return proto.CompactTextString(m) }
func (*UserPrivilegesReply) ProtoMessage()               {}
func (*UserPrivilegesReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *UserPrivilegesReply) GetUserPrivilege() []*UserPrivilege {
	if m != nil {
		return m.UserPrivilege
	}
	return nil
}

type UserPrivilegeReply struct {
	UserPrivilege *UserPrivilege `protobuf:"bytes,1,opt,name=userPrivilege" json:"userPrivilege,omitempty"`
}

func (m *UserPrivilegeReply) Reset()                    { *m = UserPrivilegeReply{} }
func (m *UserPrivilegeReply) String() string            { return proto.CompactTextString(m) }
func (*UserPrivilegeReply) ProtoMessage()               {}
func (*UserPrivilegeReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *UserPrivilegeReply) GetUserPrivilege() *UserPrivilege {
	if m != nil {
		return m.UserPrivilege
	}
	return nil
}

type FindUserPrivilegesByUserIdsRequest struct {
	Params *FindUserPrivilegesByUserIdsParams `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
}

func (m *FindUserPrivilegesByUserIdsRequest) Reset()         { *m = FindUserPrivilegesByUserIdsRequest{} }
func (m *FindUserPrivilegesByUserIdsRequest) String() string { return proto.CompactTextString(m) }
func (*FindUserPrivilegesByUserIdsRequest) ProtoMessage()    {}
func (*FindUserPrivilegesByUserIdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{6}
}

func (m *FindUserPrivilegesByUserIdsRequest) GetParams() *FindUserPrivilegesByUserIdsParams {
	if m != nil {
		return m.Params
	}
	return nil
}

type FindUserPrivilegesByUserIdsParams struct {
	Ids []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *FindUserPrivilegesByUserIdsParams) Reset()         { *m = FindUserPrivilegesByUserIdsParams{} }
func (m *FindUserPrivilegesByUserIdsParams) String() string { return proto.CompactTextString(m) }
func (*FindUserPrivilegesByUserIdsParams) ProtoMessage()    {}
func (*FindUserPrivilegesByUserIdsParams) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{7}
}

func (m *FindUserPrivilegesByUserIdsParams) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type OrderRequest struct {
	Order *Order `protobuf:"bytes,2,opt,name=order" json:"order,omitempty"`
}

func (m *OrderRequest) Reset()                    { *m = OrderRequest{} }
func (m *OrderRequest) String() string            { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()               {}
func (*OrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *OrderRequest) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type PrivilegesRequest struct {
	Privilege *Privileges `protobuf:"bytes,2,opt,name=privilege" json:"privilege,omitempty"`
}

func (m *PrivilegesRequest) Reset()                    { *m = PrivilegesRequest{} }
func (m *PrivilegesRequest) String() string            { return proto.CompactTextString(m) }
func (*PrivilegesRequest) ProtoMessage()               {}
func (*PrivilegesRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *PrivilegesRequest) GetPrivilege() *Privileges {
	if m != nil {
		return m.Privilege
	}
	return nil
}

type OrderParams struct {
	Order *Order `protobuf:"bytes,1,opt,name=order" json:"order,omitempty"`
}

func (m *OrderParams) Reset()                    { *m = OrderParams{} }
func (m *OrderParams) String() string            { return proto.CompactTextString(m) }
func (*OrderParams) ProtoMessage()               {}
func (*OrderParams) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *OrderParams) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type TmTecTranReply struct {
	Code TmTecTranReply_CODE `protobuf:"varint,1,opt,name=code,enum=membership.TmTecTranReply_CODE" json:"code,omitempty"`
}

func (m *TmTecTranReply) Reset()                    { *m = TmTecTranReply{} }
func (m *TmTecTranReply) String() string            { return proto.CompactTextString(m) }
func (*TmTecTranReply) ProtoMessage()               {}
func (*TmTecTranReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *TmTecTranReply) GetCode() TmTecTranReply_CODE {
	if m != nil {
		return m.Code
	}
	return TmTecTranReply_SuccCode
}

type UpdateUserPrivilegeRequest struct {
	UserPrivilege *UserPrivilegeUpdate `protobuf:"bytes,2,opt,name=userPrivilege" json:"userPrivilege,omitempty"`
}

func (m *UpdateUserPrivilegeRequest) Reset()                    { *m = UpdateUserPrivilegeRequest{} }
func (m *UpdateUserPrivilegeRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserPrivilegeRequest) ProtoMessage()               {}
func (*UpdateUserPrivilegeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *UpdateUserPrivilegeRequest) GetUserPrivilege() *UserPrivilegeUpdate {
	if m != nil {
		return m.UserPrivilege
	}
	return nil
}

type UpdateUserPrivilegeParams struct {
	UserPrivilege *UserPrivilegeUpdate `protobuf:"bytes,1,opt,name=userPrivilege" json:"userPrivilege,omitempty"`
}

func (m *UpdateUserPrivilegeParams) Reset()                    { *m = UpdateUserPrivilegeParams{} }
func (m *UpdateUserPrivilegeParams) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserPrivilegeParams) ProtoMessage()               {}
func (*UpdateUserPrivilegeParams) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *UpdateUserPrivilegeParams) GetUserPrivilege() *UserPrivilegeUpdate {
	if m != nil {
		return m.UserPrivilege
	}
	return nil
}

type InvalidateTokenRequest struct {
	Params *InvalidateTokenParams `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
}

func (m *InvalidateTokenRequest) Reset()                    { *m = InvalidateTokenRequest{} }
func (m *InvalidateTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*InvalidateTokenRequest) ProtoMessage()               {}
func (*InvalidateTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *InvalidateTokenRequest) GetParams() *InvalidateTokenParams {
	if m != nil {
		return m.Params
	}
	return nil
}

type InvalidateTokenParams struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *InvalidateTokenParams) Reset()                    { *m = InvalidateTokenParams{} }
func (m *InvalidateTokenParams) String() string            { return proto.CompactTextString(m) }
func (*InvalidateTokenParams) ProtoMessage()               {}
func (*InvalidateTokenParams) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *InvalidateTokenParams) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type InvalidateTokenReply struct {
}

func (m *InvalidateTokenReply) Reset()                    { *m = InvalidateTokenReply{} }
func (m *InvalidateTokenReply) String() string            { return proto.CompactTextString(m) }
func (*InvalidateTokenReply) ProtoMessage()               {}
func (*InvalidateTokenReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{16} }

func init() {
	proto.RegisterType((*RevokeUserPrivilegeRequest)(nil), "membership.RevokeUserPrivilegeRequest")
	proto.RegisterType((*RevokeUserPrivilegeReply)(nil), "membership.RevokeUserPrivilegeReply")
	proto.RegisterType((*UpdateUserProductPrivilegeRequest)(nil), "membership.UpdateUserProductPrivilegeRequest")
	proto.RegisterType((*ProductPrivilegesReply)(nil), "membership.ProductPrivilegesReply")
	proto.RegisterType((*UserPrivilegesReply)(nil), "membership.UserPrivilegesReply")
	proto.RegisterType((*UserPrivilegeReply)(nil), "membership.UserPrivilegeReply")
	proto.RegisterType((*FindUserPrivilegesByUserIdsRequest)(nil), "membership.FindUserPrivilegesByUserIdsRequest")
	proto.RegisterType((*FindUserPrivilegesByUserIdsParams)(nil), "membership.FindUserPrivilegesByUserIdsParams")
	proto.RegisterType((*OrderRequest)(nil), "membership.OrderRequest")
	proto.RegisterType((*PrivilegesRequest)(nil), "membership.PrivilegesRequest")
	proto.RegisterType((*OrderParams)(nil), "membership.OrderParams")
	proto.RegisterType((*TmTecTranReply)(nil), "membership.TmTecTranReply")
	proto.RegisterType((*UpdateUserPrivilegeRequest)(nil), "membership.UpdateUserPrivilegeRequest")
	proto.RegisterType((*UpdateUserPrivilegeParams)(nil), "membership.UpdateUserPrivilegeParams")
	proto.RegisterType((*InvalidateTokenRequest)(nil), "membership.InvalidateTokenRequest")
	proto.RegisterType((*InvalidateTokenParams)(nil), "membership.InvalidateTokenParams")
	proto.RegisterType((*InvalidateTokenReply)(nil), "membership.InvalidateTokenReply")
	proto.RegisterEnum("membership.RevokeUserPrivilegeReply_CODE", RevokeUserPrivilegeReply_CODE_name, RevokeUserPrivilegeReply_CODE_value)
	proto.RegisterEnum("membership.TmTecTranReply_CODE", TmTecTranReply_CODE_name, TmTecTranReply_CODE_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MembershipService service

type MembershipServiceClient interface {
	FindUserPrivilegesByUserIds(ctx context.Context, in *FindUserPrivilegesByUserIdsRequest, opts ...grpc.CallOption) (*UserPrivilegesReply, error)
	UpdateUserProductPrivileges(ctx context.Context, in *UpdateUserProductPrivilegeRequest, opts ...grpc.CallOption) (*ProductPrivilegesReply, error)
	RevokeUserPrivilege(ctx context.Context, in *RevokeUserPrivilegeRequest, opts ...grpc.CallOption) (*RevokeUserPrivilegeReply, error)
	TmUpsertUserPrivilegesByUserIds(ctx context.Context, in *UpdateUserPrivilegeRequest, opts ...grpc.CallOption) (*UserPrivilegeReply, error)
	BuildUpdatePrivilegesTransaction(ctx context.Context, in *PrivilegesRequest, opts ...grpc.CallOption) (*TmTecTranReply, error)
	UpdateProductPrivileges(ctx context.Context, in *PrivilegesRequest, opts ...grpc.CallOption) (*TmTecTranReply, error)
}

type membershipServiceClient struct {
	cc *grpc.ClientConn
}

func NewMembershipServiceClient(cc *grpc.ClientConn) MembershipServiceClient {
	return &membershipServiceClient{cc}
}

func (c *membershipServiceClient) FindUserPrivilegesByUserIds(ctx context.Context, in *FindUserPrivilegesByUserIdsRequest, opts ...grpc.CallOption) (*UserPrivilegesReply, error) {
	out := new(UserPrivilegesReply)
	err := grpc.Invoke(ctx, "/membership.MembershipService/FindUserPrivilegesByUserIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membershipServiceClient) UpdateUserProductPrivileges(ctx context.Context, in *UpdateUserProductPrivilegeRequest, opts ...grpc.CallOption) (*ProductPrivilegesReply, error) {
	out := new(ProductPrivilegesReply)
	err := grpc.Invoke(ctx, "/membership.MembershipService/UpdateUserProductPrivileges", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membershipServiceClient) RevokeUserPrivilege(ctx context.Context, in *RevokeUserPrivilegeRequest, opts ...grpc.CallOption) (*RevokeUserPrivilegeReply, error) {
	out := new(RevokeUserPrivilegeReply)
	err := grpc.Invoke(ctx, "/membership.MembershipService/RevokeUserPrivilege", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membershipServiceClient) TmUpsertUserPrivilegesByUserIds(ctx context.Context, in *UpdateUserPrivilegeRequest, opts ...grpc.CallOption) (*UserPrivilegeReply, error) {
	out := new(UserPrivilegeReply)
	err := grpc.Invoke(ctx, "/membership.MembershipService/TmUpsertUserPrivilegesByUserIds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membershipServiceClient) BuildUpdatePrivilegesTransaction(ctx context.Context, in *PrivilegesRequest, opts ...grpc.CallOption) (*TmTecTranReply, error) {
	out := new(TmTecTranReply)
	err := grpc.Invoke(ctx, "/membership.MembershipService/BuildUpdatePrivilegesTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *membershipServiceClient) UpdateProductPrivileges(ctx context.Context, in *PrivilegesRequest, opts ...grpc.CallOption) (*TmTecTranReply, error) {
	out := new(TmTecTranReply)
	err := grpc.Invoke(ctx, "/membership.MembershipService/UpdateProductPrivileges", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MembershipService service

type MembershipServiceServer interface {
	FindUserPrivilegesByUserIds(context.Context, *FindUserPrivilegesByUserIdsRequest) (*UserPrivilegesReply, error)
	UpdateUserProductPrivileges(context.Context, *UpdateUserProductPrivilegeRequest) (*ProductPrivilegesReply, error)
	RevokeUserPrivilege(context.Context, *RevokeUserPrivilegeRequest) (*RevokeUserPrivilegeReply, error)
	TmUpsertUserPrivilegesByUserIds(context.Context, *UpdateUserPrivilegeRequest) (*UserPrivilegeReply, error)
	BuildUpdatePrivilegesTransaction(context.Context, *PrivilegesRequest) (*TmTecTranReply, error)
	UpdateProductPrivileges(context.Context, *PrivilegesRequest) (*TmTecTranReply, error)
}

func RegisterMembershipServiceServer(s *grpc.Server, srv MembershipServiceServer) {
	s.RegisterService(&_MembershipService_serviceDesc, srv)
}

func _MembershipService_FindUserPrivilegesByUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindUserPrivilegesByUserIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServiceServer).FindUserPrivilegesByUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/membership.MembershipService/FindUserPrivilegesByUserIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServiceServer).FindUserPrivilegesByUserIds(ctx, req.(*FindUserPrivilegesByUserIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MembershipService_UpdateUserProductPrivileges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserProductPrivilegeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServiceServer).UpdateUserProductPrivileges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/membership.MembershipService/UpdateUserProductPrivileges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServiceServer).UpdateUserProductPrivileges(ctx, req.(*UpdateUserProductPrivilegeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MembershipService_RevokeUserPrivilege_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeUserPrivilegeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServiceServer).RevokeUserPrivilege(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/membership.MembershipService/RevokeUserPrivilege",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServiceServer).RevokeUserPrivilege(ctx, req.(*RevokeUserPrivilegeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MembershipService_TmUpsertUserPrivilegesByUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPrivilegeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServiceServer).TmUpsertUserPrivilegesByUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/membership.MembershipService/TmUpsertUserPrivilegesByUserIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServiceServer).TmUpsertUserPrivilegesByUserIds(ctx, req.(*UpdateUserPrivilegeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MembershipService_BuildUpdatePrivilegesTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivilegesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServiceServer).BuildUpdatePrivilegesTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/membership.MembershipService/BuildUpdatePrivilegesTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServiceServer).BuildUpdatePrivilegesTransaction(ctx, req.(*PrivilegesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MembershipService_UpdateProductPrivileges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivilegesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MembershipServiceServer).UpdateProductPrivileges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/membership.MembershipService/UpdateProductPrivileges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MembershipServiceServer).UpdateProductPrivileges(ctx, req.(*PrivilegesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MembershipService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "membership.MembershipService",
	HandlerType: (*MembershipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindUserPrivilegesByUserIds",
			Handler:    _MembershipService_FindUserPrivilegesByUserIds_Handler,
		},
		{
			MethodName: "UpdateUserProductPrivileges",
			Handler:    _MembershipService_UpdateUserProductPrivileges_Handler,
		},
		{
			MethodName: "RevokeUserPrivilege",
			Handler:    _MembershipService_RevokeUserPrivilege_Handler,
		},
		{
			MethodName: "TmUpsertUserPrivilegesByUserIds",
			Handler:    _MembershipService_TmUpsertUserPrivilegesByUserIds_Handler,
		},
		{
			MethodName: "BuildUpdatePrivilegesTransaction",
			Handler:    _MembershipService_BuildUpdatePrivilegesTransaction_Handler,
		},
		{
			MethodName: "UpdateProductPrivileges",
			Handler:    _MembershipService_UpdateProductPrivileges_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "membership/membership.service.proto",
}

func init() { proto.RegisterFile("membership/membership.service.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 789 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x5b, 0x6f, 0xd3, 0x4a,
	0x10, 0x8e, 0x4f, 0xda, 0xea, 0x74, 0x7a, 0x51, 0xbb, 0xed, 0xc9, 0x49, 0x13, 0x95, 0x24, 0x5b,
	0x54, 0xc2, 0x43, 0x12, 0x35, 0x85, 0x02, 0x0f, 0x15, 0xea, 0x25, 0xa0, 0x48, 0x40, 0x2b, 0x37,
	0x29, 0x12, 0x2f, 0xc8, 0xb1, 0x57, 0xa9, 0x15, 0xdb, 0x6b, 0xd6, 0x76, 0x20, 0x7f, 0x80, 0x27,
	0x9e, 0x78, 0xe2, 0xf6, 0x3b, 0xf8, 0x79, 0x80, 0x6c, 0x6f, 0xe2, 0x4b, 0xed, 0xa6, 0x11, 0x79,
	0xf2, 0xec, 0xce, 0xf7, 0x7d, 0xb3, 0x33, 0x3b, 0xb3, 0x81, 0x1d, 0x9d, 0xe8, 0x3d, 0xc2, 0xac,
	0x2b, 0xd5, 0x6c, 0x04, 0x9f, 0x75, 0x8b, 0xb0, 0xa1, 0x2a, 0x93, 0xba, 0xc9, 0xa8, 0x4d, 0x11,
	0x04, 0x3b, 0x85, 0x62, 0x32, 0xc0, 0x73, 0xc4, 0x75, 0x28, 0x88, 0x64, 0x48, 0x07, 0xa4, 0x6b,
	0x11, 0x76, 0xce, 0xd4, 0xa1, 0xaa, 0x91, 0x3e, 0x11, 0xc9, 0x3b, 0x87, 0x58, 0x36, 0x5a, 0x83,
	0xac, 0xa3, 0x2a, 0x79, 0xa1, 0x2c, 0x54, 0x17, 0x45, 0xf7, 0x13, 0x7f, 0x12, 0x20, 0x9f, 0x08,
	0x30, 0xb5, 0x11, 0x3a, 0x84, 0x39, 0x99, 0x2a, 0xc4, 0xf3, 0x5f, 0x6d, 0xde, 0xaf, 0x87, 0xd4,
	0xd2, 0x30, 0xf5, 0x93, 0xb3, 0xd3, 0x96, 0xe8, 0xc1, 0xf0, 0x1e, 0xcc, 0xb9, 0x16, 0x5a, 0x86,
	0x7f, 0x2f, 0x1c, 0x59, 0x3e, 0xa1, 0x0a, 0x59, 0xcb, 0xa0, 0x22, 0xac, 0x76, 0x8d, 0x81, 0x41,
	0xdf, 0x1b, 0x2d, 0xc6, 0xbc, 0xb5, 0x1f, 0xbf, 0xf9, 0x4f, 0xc0, 0x9f, 0x05, 0xa8, 0x74, 0x4d,
	0x45, 0xb2, 0x39, 0x35, 0x55, 0x1c, 0xd9, 0xbe, 0x76, 0x8c, 0x1c, 0x2c, 0x38, 0x16, 0x61, 0xed,
	0x53, 0x7e, 0x12, 0x6e, 0xa1, 0x43, 0x00, 0x73, 0xec, 0x6b, 0xe5, 0xff, 0x29, 0x0b, 0xd5, 0xa5,
	0xe6, 0x76, 0x38, 0xea, 0x38, 0xa1, 0x25, 0x86, 0x00, 0x2e, 0x2d, 0xf9, 0x60, 0xaa, 0x8c, 0xe4,
	0xb3, 0x65, 0xa1, 0x9a, 0x15, 0xb9, 0x85, 0x5f, 0x43, 0xee, 0x3a, 0x90, 0x27, 0x28, 0x2c, 0x28,
	0xcc, 0x28, 0x88, 0x2f, 0x61, 0x23, 0x92, 0x41, 0xce, 0xfa, 0x14, 0x56, 0x9c, 0xf0, 0x72, 0x5e,
	0x28, 0x67, 0xab, 0x4b, 0xcd, 0xad, 0x30, 0x71, 0x34, 0xf3, 0x51, 0x7f, 0xdc, 0x05, 0x94, 0x50,
	0xcd, 0x04, 0x5a, 0x61, 0x26, 0xda, 0x01, 0xe0, 0x67, 0xaa, 0xa1, 0x44, 0x43, 0x3e, 0x1e, 0xb9,
	0x76, 0x5b, 0xb1, 0xc6, 0xc5, 0x69, 0xc1, 0x82, 0x29, 0x31, 0x49, 0x1f, 0xe7, 0xa3, 0x16, 0xe6,
	0xbf, 0x01, 0x7f, 0xee, 0x81, 0x44, 0x0e, 0xc6, 0x0f, 0xa1, 0x32, 0xd5, 0xd9, 0xbd, 0xcf, 0xaa,
	0x62, 0x79, 0xf9, 0x59, 0x14, 0xdd, 0x4f, 0xfc, 0x08, 0x96, 0xcf, 0x98, 0x42, 0xd8, 0x38, 0x9a,
	0x7b, 0x30, 0x4f, 0x5d, 0x9b, 0xdf, 0x86, 0xf5, 0x70, 0x30, 0xbe, 0xa3, 0xbf, 0x8f, 0xdb, 0xb0,
	0x1e, 0xae, 0x83, 0x8f, 0x7e, 0x00, 0x8b, 0x93, 0x72, 0x71, 0x86, 0x5c, 0xb4, 0xbc, 0x13, 0x44,
	0xe0, 0x88, 0x0f, 0x60, 0xc9, 0xa3, 0xe6, 0x41, 0x4e, 0x42, 0x10, 0xa6, 0x84, 0xf0, 0x4b, 0x80,
	0xd5, 0x8e, 0xde, 0x21, 0x72, 0x87, 0x49, 0x86, 0x5f, 0xb3, 0xfd, 0x48, 0x07, 0x96, 0xc2, 0xd0,
	0xa8, 0x67, 0xb8, 0xef, 0x7e, 0x0a, 0x33, 0x37, 0x1e, 0x2a, 0xc0, 0x8a, 0xbb, 0xaa, 0x51, 0x8b,
	0x28, 0xde, 0xde, 0xf7, 0x60, 0xaf, 0x04, 0xa8, 0xc5, 0xd8, 0x2b, 0x6a, 0x9f, 0x50, 0xdd, 0xd4,
	0x88, 0x4d, 0x3c, 0x87, 0x6f, 0x81, 0xc3, 0x0e, 0xe4, 0x5a, 0x8c, 0x1d, 0x69, 0x8c, 0x48, 0xca,
	0xe8, 0x48, 0x96, 0x89, 0x69, 0x73, 0x96, 0xaf, 0x81, 0x53, 0x05, 0x36, 0x5d, 0x05, 0xc9, 0x30,
	0xa8, 0xfd, 0x9c, 0xd8, 0x2f, 0xa8, 0x3c, 0xf0, 0x5c, 0xbe, 0x04, 0xdd, 0x2f, 0x43, 0x21, 0xdc,
	0xfc, 0xb1, 0xae, 0x6f, 0xc5, 0xef, 0xaf, 0x5f, 0x90, 0x52, 0xea, 0xfd, 0xf5, 0xb9, 0xe2, 0xb7,
	0xb8, 0x07, 0x5b, 0x09, 0x22, 0xbc, 0x56, 0xad, 0xe4, 0x1e, 0x99, 0x55, 0xe3, 0x02, 0x72, 0x6d,
	0x63, 0x28, 0x69, 0xaa, 0xbb, 0xd9, 0xa1, 0x03, 0x62, 0x8c, 0x0f, 0xf1, 0x24, 0xd6, 0x1d, 0x95,
	0x30, 0x73, 0x0c, 0x13, 0xeb, 0x88, 0x1a, 0xfc, 0x97, 0xe8, 0x80, 0x36, 0x61, 0xde, 0x76, 0x4d,
	0x3e, 0x0d, 0x7d, 0x03, 0xe7, 0x60, 0xf3, 0x5a, 0x0c, 0xa6, 0x36, 0x6a, 0x7e, 0x9c, 0x87, 0xf5,
	0x97, 0x13, 0xcd, 0x0b, 0xff, 0x99, 0x41, 0x26, 0x14, 0x6f, 0x68, 0x37, 0x54, 0xbf, 0x65, 0x13,
	0xf3, 0x63, 0x16, 0xd2, 0x13, 0xe6, 0xcf, 0x38, 0x9c, 0x41, 0x0c, 0x8a, 0xe9, 0x93, 0xde, 0x42,
	0x91, 0xb1, 0x31, 0xf5, 0x49, 0x28, 0xe0, 0x9b, 0xa7, 0x2e, 0xd7, 0x24, 0xb0, 0x91, 0xf0, 0x70,
	0xa1, 0xdd, 0xa9, 0x2f, 0x9b, 0x2f, 0x72, 0xf7, 0x36, 0x2f, 0x20, 0xce, 0xa0, 0x01, 0x94, 0x3a,
	0x7a, 0xd7, 0xb4, 0x08, 0xb3, 0xd3, 0x12, 0xba, 0x9b, 0x76, 0xbc, 0x98, 0xe4, 0x9d, 0xf4, 0xe9,
	0xcc, 0xc5, 0xde, 0x42, 0xf9, 0xd8, 0x51, 0x35, 0xc5, 0x27, 0x09, 0xb4, 0xdc, 0xd1, 0x60, 0x49,
	0xb2, 0xad, 0x52, 0x03, 0x6d, 0xa7, 0x0c, 0x2d, 0x2e, 0x52, 0x48, 0x9f, 0x2b, 0x38, 0x83, 0x2e,
	0xe1, 0xff, 0x31, 0x77, 0xbc, 0x48, 0x7f, 0xc3, 0x7b, 0xfc, 0xf8, 0xcd, 0x41, 0x5f, 0xb5, 0xaf,
	0x9c, 0x5e, 0x5d, 0xa6, 0x7a, 0xc3, 0xdc, 0x93, 0x8d, 0x86, 0x2d, 0x19, 0xb6, 0x64, 0xd4, 0x14,
	0xaa, 0x4b, 0xaa, 0x51, 0xb3, 0xe4, 0x2b, 0xa2, 0x4b, 0x8d, 0x3e, 0xd5, 0x24, 0xa3, 0x1f, 0xfa,
	0xab, 0xd3, 0x5b, 0xf0, 0xfe, 0xeb, 0xec, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x88, 0xda,
	0xd3, 0x3b, 0x09, 0x00, 0x00,
}
