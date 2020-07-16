// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/menger/menger.proto

package wm_sx_svc_menger

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Menger service

type MengerService interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error)
	GetMenger(ctx context.Context, in *GetMengerRequest, opts ...client.CallOption) (*GetMengerResponse, error)
	GetMengerList(ctx context.Context, in *GetMengerListRequest, opts ...client.CallOption) (*GetMengerListResponse, error)
}

type mengerService struct {
	c    client.Client
	name string
}

func NewMengerService(name string, c client.Client) MengerService {
	return &mengerService{
		c:    c,
		name: name,
	}
}

func (c *mengerService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.name, "Menger.Register", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mengerService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Menger.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mengerService) Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error) {
	req := c.c.NewRequest(c.name, "Menger.Logout", in)
	out := new(LogoutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mengerService) GetMenger(ctx context.Context, in *GetMengerRequest, opts ...client.CallOption) (*GetMengerResponse, error) {
	req := c.c.NewRequest(c.name, "Menger.GetMenger", in)
	out := new(GetMengerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mengerService) GetMengerList(ctx context.Context, in *GetMengerListRequest, opts ...client.CallOption) (*GetMengerListResponse, error) {
	req := c.c.NewRequest(c.name, "Menger.GetMengerList", in)
	out := new(GetMengerListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Menger service

type MengerHandler interface {
	Register(context.Context, *RegisterRequest, *RegisterResponse) error
	Login(context.Context, *LoginRequest, *LoginResponse) error
	Logout(context.Context, *LogoutRequest, *LogoutResponse) error
	GetMenger(context.Context, *GetMengerRequest, *GetMengerResponse) error
	GetMengerList(context.Context, *GetMengerListRequest, *GetMengerListResponse) error
}

func RegisterMengerHandler(s server.Server, hdlr MengerHandler, opts ...server.HandlerOption) error {
	type menger interface {
		Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		Logout(ctx context.Context, in *LogoutRequest, out *LogoutResponse) error
		GetMenger(ctx context.Context, in *GetMengerRequest, out *GetMengerResponse) error
		GetMengerList(ctx context.Context, in *GetMengerListRequest, out *GetMengerListResponse) error
	}
	type Menger struct {
		menger
	}
	h := &mengerHandler{hdlr}
	return s.Handle(s.NewHandler(&Menger{h}, opts...))
}

type mengerHandler struct {
	MengerHandler
}

func (h *mengerHandler) Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error {
	return h.MengerHandler.Register(ctx, in, out)
}

func (h *mengerHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.MengerHandler.Login(ctx, in, out)
}

func (h *mengerHandler) Logout(ctx context.Context, in *LogoutRequest, out *LogoutResponse) error {
	return h.MengerHandler.Logout(ctx, in, out)
}

func (h *mengerHandler) GetMenger(ctx context.Context, in *GetMengerRequest, out *GetMengerResponse) error {
	return h.MengerHandler.GetMenger(ctx, in, out)
}

func (h *mengerHandler) GetMengerList(ctx context.Context, in *GetMengerListRequest, out *GetMengerListResponse) error {
	return h.MengerHandler.GetMengerList(ctx, in, out)
}
