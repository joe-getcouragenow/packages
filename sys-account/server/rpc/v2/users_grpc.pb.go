// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	NewAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*Account, error)
	ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error)
	SearchAccounts(ctx context.Context, in *SearchAccountsRequest, opts ...grpc.CallOption) (*SearchAccountsResponse, error)
	AssignAccountToRole(ctx context.Context, in *AssignAccountToRoleRequest, opts ...grpc.CallOption) (*Account, error)
	UpdateAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error)
	DisableAccount(ctx context.Context, in *DisableAccountRequest, opts ...grpc.CallOption) (*Account, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

var accountServiceNewAccountStreamDesc = &grpc.StreamDesc{
	StreamName: "NewAccount",
}

func (c *accountServiceClient) NewAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/getcouragenow.v2.sys_account.AccountService/NewAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var accountServiceGetAccountStreamDesc = &grpc.StreamDesc{
	StreamName: "GetAccount",
}

func (c *accountServiceClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/getcouragenow.v2.sys_account.AccountService/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var accountServiceListAccountsStreamDesc = &grpc.StreamDesc{
	StreamName: "ListAccounts",
}

func (c *accountServiceClient) ListAccounts(ctx context.Context, in *ListAccountsRequest, opts ...grpc.CallOption) (*ListAccountsResponse, error) {
	out := new(ListAccountsResponse)
	err := c.cc.Invoke(ctx, "/getcouragenow.v2.sys_account.AccountService/ListAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var accountServiceSearchAccountsStreamDesc = &grpc.StreamDesc{
	StreamName: "SearchAccounts",
}

func (c *accountServiceClient) SearchAccounts(ctx context.Context, in *SearchAccountsRequest, opts ...grpc.CallOption) (*SearchAccountsResponse, error) {
	out := new(SearchAccountsResponse)
	err := c.cc.Invoke(ctx, "/getcouragenow.v2.sys_account.AccountService/SearchAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var accountServiceAssignAccountToRoleStreamDesc = &grpc.StreamDesc{
	StreamName: "AssignAccountToRole",
}

func (c *accountServiceClient) AssignAccountToRole(ctx context.Context, in *AssignAccountToRoleRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/getcouragenow.v2.sys_account.AccountService/AssignAccountToRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var accountServiceUpdateAccountStreamDesc = &grpc.StreamDesc{
	StreamName: "UpdateAccount",
}

func (c *accountServiceClient) UpdateAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/getcouragenow.v2.sys_account.AccountService/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var accountServiceDisableAccountStreamDesc = &grpc.StreamDesc{
	StreamName: "DisableAccount",
}

func (c *accountServiceClient) DisableAccount(ctx context.Context, in *DisableAccountRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/getcouragenow.v2.sys_account.AccountService/DisableAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceService is the service API for AccountService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterAccountServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type AccountServiceService struct {
	NewAccount          func(context.Context, *Account) (*Account, error)
	GetAccount          func(context.Context, *GetAccountRequest) (*Account, error)
	ListAccounts        func(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error)
	SearchAccounts      func(context.Context, *SearchAccountsRequest) (*SearchAccountsResponse, error)
	AssignAccountToRole func(context.Context, *AssignAccountToRoleRequest) (*Account, error)
	UpdateAccount       func(context.Context, *Account) (*Account, error)
	DisableAccount      func(context.Context, *DisableAccountRequest) (*Account, error)
}

func (s *AccountServiceService) newAccount(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.NewAccount == nil {
		return nil, status.Errorf(codes.Unimplemented, "method NewAccount not implemented")
	}
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.NewAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/getcouragenow.v2.sys_account.AccountService/NewAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.NewAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AccountServiceService) getAccount(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.GetAccount == nil {
		return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
	}
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/getcouragenow.v2.sys_account.AccountService/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AccountServiceService) listAccounts(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.ListAccounts == nil {
		return nil, status.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
	}
	in := new(ListAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/getcouragenow.v2.sys_account.AccountService/ListAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListAccounts(ctx, req.(*ListAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AccountServiceService) searchAccounts(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.SearchAccounts == nil {
		return nil, status.Errorf(codes.Unimplemented, "method SearchAccounts not implemented")
	}
	in := new(SearchAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.SearchAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/getcouragenow.v2.sys_account.AccountService/SearchAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.SearchAccounts(ctx, req.(*SearchAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AccountServiceService) assignAccountToRole(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.AssignAccountToRole == nil {
		return nil, status.Errorf(codes.Unimplemented, "method AssignAccountToRole not implemented")
	}
	in := new(AssignAccountToRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.AssignAccountToRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/getcouragenow.v2.sys_account.AccountService/AssignAccountToRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.AssignAccountToRole(ctx, req.(*AssignAccountToRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AccountServiceService) updateAccount(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.UpdateAccount == nil {
		return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
	}
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/getcouragenow.v2.sys_account.AccountService/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AccountServiceService) disableAccount(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.DisableAccount == nil {
		return nil, status.Errorf(codes.Unimplemented, "method DisableAccount not implemented")
	}
	in := new(DisableAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.DisableAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/getcouragenow.v2.sys_account.AccountService/DisableAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DisableAccount(ctx, req.(*DisableAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterAccountServiceService registers a service implementation with a gRPC server.
func RegisterAccountServiceService(s grpc.ServiceRegistrar, srv *AccountServiceService) {
	sd := grpc.ServiceDesc{
		ServiceName: "getcouragenow.v2.sys_account.AccountService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "NewAccount",
				Handler:    srv.newAccount,
			},
			{
				MethodName: "GetAccount",
				Handler:    srv.getAccount,
			},
			{
				MethodName: "ListAccounts",
				Handler:    srv.listAccounts,
			},
			{
				MethodName: "SearchAccounts",
				Handler:    srv.searchAccounts,
			},
			{
				MethodName: "AssignAccountToRole",
				Handler:    srv.assignAccountToRole,
			},
			{
				MethodName: "UpdateAccount",
				Handler:    srv.updateAccount,
			},
			{
				MethodName: "DisableAccount",
				Handler:    srv.disableAccount,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "users.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewAccountServiceService creates a new AccountServiceService containing the
// implemented methods of the AccountService service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewAccountServiceService(s interface{}) *AccountServiceService {
	ns := &AccountServiceService{}
	if h, ok := s.(interface {
		NewAccount(context.Context, *Account) (*Account, error)
	}); ok {
		ns.NewAccount = h.NewAccount
	}
	if h, ok := s.(interface {
		GetAccount(context.Context, *GetAccountRequest) (*Account, error)
	}); ok {
		ns.GetAccount = h.GetAccount
	}
	if h, ok := s.(interface {
		ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error)
	}); ok {
		ns.ListAccounts = h.ListAccounts
	}
	if h, ok := s.(interface {
		SearchAccounts(context.Context, *SearchAccountsRequest) (*SearchAccountsResponse, error)
	}); ok {
		ns.SearchAccounts = h.SearchAccounts
	}
	if h, ok := s.(interface {
		AssignAccountToRole(context.Context, *AssignAccountToRoleRequest) (*Account, error)
	}); ok {
		ns.AssignAccountToRole = h.AssignAccountToRole
	}
	if h, ok := s.(interface {
		UpdateAccount(context.Context, *Account) (*Account, error)
	}); ok {
		ns.UpdateAccount = h.UpdateAccount
	}
	if h, ok := s.(interface {
		DisableAccount(context.Context, *DisableAccountRequest) (*Account, error)
	}); ok {
		ns.DisableAccount = h.DisableAccount
	}
	return ns
}

// UnstableAccountServiceService is the service API for AccountService service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableAccountServiceService interface {
	NewAccount(context.Context, *Account) (*Account, error)
	GetAccount(context.Context, *GetAccountRequest) (*Account, error)
	ListAccounts(context.Context, *ListAccountsRequest) (*ListAccountsResponse, error)
	SearchAccounts(context.Context, *SearchAccountsRequest) (*SearchAccountsResponse, error)
	AssignAccountToRole(context.Context, *AssignAccountToRoleRequest) (*Account, error)
	UpdateAccount(context.Context, *Account) (*Account, error)
	DisableAccount(context.Context, *DisableAccountRequest) (*Account, error)
}
