// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/auth/auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	proto/auth/auth.proto

It has these top-level messages:
	User
	Request
	ResponseUser
	ResponseRole
	ResponseMenu
	ResponseToken
	Token
	Error
	Role
	Menu
	Badge
	Wrapper
	Atributes
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseUser, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseToken, error)
	GetUserMenus(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseMenu, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*ResponseToken, error)
	CreateRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error)
	GetRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error)
	GetAllRoles(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseRole, error)
	CreateMenu(ctx context.Context, in *Menu, opts ...client.CallOption) (*ResponseMenu, error)
	GetMenu(ctx context.Context, in *Menu, opts ...client.CallOption) (*ResponseMenu, error)
	GetAllMenues(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseMenu, error)
}

type authClient struct {
	c           client.Client
	serviceName string
}

func NewAuthClient(serviceName string, c client.Client) AuthClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Create", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Get", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseUser, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetAll", in)
	out := new(ResponseUser)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseToken, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Auth", in)
	out := new(ResponseToken)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetUserMenus(ctx context.Context, in *User, opts ...client.CallOption) (*ResponseMenu, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetUserMenus", in)
	out := new(ResponseMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*ResponseToken, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.ValidateToken", in)
	out := new(ResponseToken)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.CreateRole", in)
	out := new(ResponseRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetRole(ctx context.Context, in *Role, opts ...client.CallOption) (*ResponseRole, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetRole", in)
	out := new(ResponseRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAllRoles(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseRole, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetAllRoles", in)
	out := new(ResponseRole)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CreateMenu(ctx context.Context, in *Menu, opts ...client.CallOption) (*ResponseMenu, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.CreateMenu", in)
	out := new(ResponseMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetMenu(ctx context.Context, in *Menu, opts ...client.CallOption) (*ResponseMenu, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetMenu", in)
	out := new(ResponseMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetAllMenues(ctx context.Context, in *Request, opts ...client.CallOption) (*ResponseMenu, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.GetAllMenues", in)
	out := new(ResponseMenu)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	Create(context.Context, *User, *ResponseUser) error
	Get(context.Context, *User, *ResponseUser) error
	GetAll(context.Context, *Request, *ResponseUser) error
	Auth(context.Context, *User, *ResponseToken) error
	GetUserMenus(context.Context, *User, *ResponseMenu) error
	ValidateToken(context.Context, *Token, *ResponseToken) error
	CreateRole(context.Context, *Role, *ResponseRole) error
	GetRole(context.Context, *Role, *ResponseRole) error
	GetAllRoles(context.Context, *Request, *ResponseRole) error
	CreateMenu(context.Context, *Menu, *ResponseMenu) error
	GetMenu(context.Context, *Menu, *ResponseMenu) error
	GetAllMenues(context.Context, *Request, *ResponseMenu) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Auth{hdlr}, opts...))
}

type Auth struct {
	AuthHandler
}

func (h *Auth) Create(ctx context.Context, in *User, out *ResponseUser) error {
	return h.AuthHandler.Create(ctx, in, out)
}

func (h *Auth) Get(ctx context.Context, in *User, out *ResponseUser) error {
	return h.AuthHandler.Get(ctx, in, out)
}

func (h *Auth) GetAll(ctx context.Context, in *Request, out *ResponseUser) error {
	return h.AuthHandler.GetAll(ctx, in, out)
}

func (h *Auth) Auth(ctx context.Context, in *User, out *ResponseToken) error {
	return h.AuthHandler.Auth(ctx, in, out)
}

func (h *Auth) GetUserMenus(ctx context.Context, in *User, out *ResponseMenu) error {
	return h.AuthHandler.GetUserMenus(ctx, in, out)
}

func (h *Auth) ValidateToken(ctx context.Context, in *Token, out *ResponseToken) error {
	return h.AuthHandler.ValidateToken(ctx, in, out)
}

func (h *Auth) CreateRole(ctx context.Context, in *Role, out *ResponseRole) error {
	return h.AuthHandler.CreateRole(ctx, in, out)
}

func (h *Auth) GetRole(ctx context.Context, in *Role, out *ResponseRole) error {
	return h.AuthHandler.GetRole(ctx, in, out)
}

func (h *Auth) GetAllRoles(ctx context.Context, in *Request, out *ResponseRole) error {
	return h.AuthHandler.GetAllRoles(ctx, in, out)
}

func (h *Auth) CreateMenu(ctx context.Context, in *Menu, out *ResponseMenu) error {
	return h.AuthHandler.CreateMenu(ctx, in, out)
}

func (h *Auth) GetMenu(ctx context.Context, in *Menu, out *ResponseMenu) error {
	return h.AuthHandler.GetMenu(ctx, in, out)
}

func (h *Auth) GetAllMenues(ctx context.Context, in *Request, out *ResponseMenu) error {
	return h.AuthHandler.GetAllMenues(ctx, in, out)
}
