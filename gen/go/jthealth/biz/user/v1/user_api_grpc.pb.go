// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package userv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserAPIClient is the client API for UserAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAPIClient interface {
	//---------------------经通天下------------------
	//手机号验证码登录
	JingTongSignInByPhoneCode(ctx context.Context, in *JingTongSignInByPhoneCodeRequest, opts ...grpc.CallOption) (*JingTongSignInByPhoneCodeResponse, error)
	//微信授权登录
	JingTongSignInByWechatMiniProgram(ctx context.Context, in *JingTongSignInByWechatMiniProgramRequest, opts ...grpc.CallOption) (*JingTongSignInByWechatMiniProgramResponse, error)
	//获取用户信息
	JingTongGetUserProfile(ctx context.Context, in *JingTongGetUserProfileRequest, opts ...grpc.CallOption) (*JingTongGetUserProfileResponse, error)
	//更新用户信息
	JingTongUpdateUserProfile(ctx context.Context, in *JingTongUpdateUserProfileRequest, opts ...grpc.CallOption) (*JingTongUpdateUserProfileResponse, error)
	//修改手机号
	JingTongUpdatePhone(ctx context.Context, in *JingTongUpdatePhoneRequest, opts ...grpc.CallOption) (*JingTongUpdatePhoneResponse, error)
	//后台用户登录
	JingTongBgLogin(ctx context.Context, in *JingTongBgLoginRequest, opts ...grpc.CallOption) (*JingTongBgLoginResponse, error)
}

type userAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAPIClient(cc grpc.ClientConnInterface) UserAPIClient {
	return &userAPIClient{cc}
}

func (c *userAPIClient) JingTongSignInByPhoneCode(ctx context.Context, in *JingTongSignInByPhoneCodeRequest, opts ...grpc.CallOption) (*JingTongSignInByPhoneCodeResponse, error) {
	out := new(JingTongSignInByPhoneCodeResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/JingTongSignInByPhoneCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) JingTongSignInByWechatMiniProgram(ctx context.Context, in *JingTongSignInByWechatMiniProgramRequest, opts ...grpc.CallOption) (*JingTongSignInByWechatMiniProgramResponse, error) {
	out := new(JingTongSignInByWechatMiniProgramResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/JingTongSignInByWechatMiniProgram", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) JingTongGetUserProfile(ctx context.Context, in *JingTongGetUserProfileRequest, opts ...grpc.CallOption) (*JingTongGetUserProfileResponse, error) {
	out := new(JingTongGetUserProfileResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/JingTongGetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) JingTongUpdateUserProfile(ctx context.Context, in *JingTongUpdateUserProfileRequest, opts ...grpc.CallOption) (*JingTongUpdateUserProfileResponse, error) {
	out := new(JingTongUpdateUserProfileResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/JingTongUpdateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) JingTongUpdatePhone(ctx context.Context, in *JingTongUpdatePhoneRequest, opts ...grpc.CallOption) (*JingTongUpdatePhoneResponse, error) {
	out := new(JingTongUpdatePhoneResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/JingTongUpdatePhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) JingTongBgLogin(ctx context.Context, in *JingTongBgLoginRequest, opts ...grpc.CallOption) (*JingTongBgLoginResponse, error) {
	out := new(JingTongBgLoginResponse)
	err := c.cc.Invoke(ctx, "/jthealth.biz.user.v1.UserAPI/JingTongBgLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAPIServer is the server API for UserAPI service.
// All implementations must embed UnimplementedUserAPIServer
// for forward compatibility
type UserAPIServer interface {
	//---------------------经通天下------------------
	//手机号验证码登录
	JingTongSignInByPhoneCode(context.Context, *JingTongSignInByPhoneCodeRequest) (*JingTongSignInByPhoneCodeResponse, error)
	//微信授权登录
	JingTongSignInByWechatMiniProgram(context.Context, *JingTongSignInByWechatMiniProgramRequest) (*JingTongSignInByWechatMiniProgramResponse, error)
	//获取用户信息
	JingTongGetUserProfile(context.Context, *JingTongGetUserProfileRequest) (*JingTongGetUserProfileResponse, error)
	//更新用户信息
	JingTongUpdateUserProfile(context.Context, *JingTongUpdateUserProfileRequest) (*JingTongUpdateUserProfileResponse, error)
	//修改手机号
	JingTongUpdatePhone(context.Context, *JingTongUpdatePhoneRequest) (*JingTongUpdatePhoneResponse, error)
	//后台用户登录
	JingTongBgLogin(context.Context, *JingTongBgLoginRequest) (*JingTongBgLoginResponse, error)
	mustEmbedUnimplementedUserAPIServer()
}

// UnimplementedUserAPIServer must be embedded to have forward compatible implementations.
type UnimplementedUserAPIServer struct {
}

func (UnimplementedUserAPIServer) JingTongSignInByPhoneCode(context.Context, *JingTongSignInByPhoneCodeRequest) (*JingTongSignInByPhoneCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JingTongSignInByPhoneCode not implemented")
}
func (UnimplementedUserAPIServer) JingTongSignInByWechatMiniProgram(context.Context, *JingTongSignInByWechatMiniProgramRequest) (*JingTongSignInByWechatMiniProgramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JingTongSignInByWechatMiniProgram not implemented")
}
func (UnimplementedUserAPIServer) JingTongGetUserProfile(context.Context, *JingTongGetUserProfileRequest) (*JingTongGetUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JingTongGetUserProfile not implemented")
}
func (UnimplementedUserAPIServer) JingTongUpdateUserProfile(context.Context, *JingTongUpdateUserProfileRequest) (*JingTongUpdateUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JingTongUpdateUserProfile not implemented")
}
func (UnimplementedUserAPIServer) JingTongUpdatePhone(context.Context, *JingTongUpdatePhoneRequest) (*JingTongUpdatePhoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JingTongUpdatePhone not implemented")
}
func (UnimplementedUserAPIServer) JingTongBgLogin(context.Context, *JingTongBgLoginRequest) (*JingTongBgLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JingTongBgLogin not implemented")
}
func (UnimplementedUserAPIServer) mustEmbedUnimplementedUserAPIServer() {}

// UnsafeUserAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAPIServer will
// result in compilation errors.
type UnsafeUserAPIServer interface {
	mustEmbedUnimplementedUserAPIServer()
}

func RegisterUserAPIServer(s *grpc.Server, srv UserAPIServer) {
	s.RegisterService(&_UserAPI_serviceDesc, srv)
}

func _UserAPI_JingTongSignInByPhoneCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JingTongSignInByPhoneCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).JingTongSignInByPhoneCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/JingTongSignInByPhoneCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).JingTongSignInByPhoneCode(ctx, req.(*JingTongSignInByPhoneCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_JingTongSignInByWechatMiniProgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JingTongSignInByWechatMiniProgramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).JingTongSignInByWechatMiniProgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/JingTongSignInByWechatMiniProgram",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).JingTongSignInByWechatMiniProgram(ctx, req.(*JingTongSignInByWechatMiniProgramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_JingTongGetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JingTongGetUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).JingTongGetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/JingTongGetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).JingTongGetUserProfile(ctx, req.(*JingTongGetUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_JingTongUpdateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JingTongUpdateUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).JingTongUpdateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/JingTongUpdateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).JingTongUpdateUserProfile(ctx, req.(*JingTongUpdateUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_JingTongUpdatePhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JingTongUpdatePhoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).JingTongUpdatePhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/JingTongUpdatePhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).JingTongUpdatePhone(ctx, req.(*JingTongUpdatePhoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_JingTongBgLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JingTongBgLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).JingTongBgLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jthealth.biz.user.v1.UserAPI/JingTongBgLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).JingTongBgLogin(ctx, req.(*JingTongBgLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jthealth.biz.user.v1.UserAPI",
	HandlerType: (*UserAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JingTongSignInByPhoneCode",
			Handler:    _UserAPI_JingTongSignInByPhoneCode_Handler,
		},
		{
			MethodName: "JingTongSignInByWechatMiniProgram",
			Handler:    _UserAPI_JingTongSignInByWechatMiniProgram_Handler,
		},
		{
			MethodName: "JingTongGetUserProfile",
			Handler:    _UserAPI_JingTongGetUserProfile_Handler,
		},
		{
			MethodName: "JingTongUpdateUserProfile",
			Handler:    _UserAPI_JingTongUpdateUserProfile_Handler,
		},
		{
			MethodName: "JingTongUpdatePhone",
			Handler:    _UserAPI_JingTongUpdatePhone_Handler,
		},
		{
			MethodName: "JingTongBgLogin",
			Handler:    _UserAPI_JingTongBgLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jthealth/biz/user/v1/user_api.proto",
}
