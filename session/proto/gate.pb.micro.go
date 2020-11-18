// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: session/proto/gate.proto

package gate

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for McbApp service

func NewMcbAppEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for McbApp service

type McbAppService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type mcbAppService struct {
	c    client.Client
	name string
}

func NewMcbAppService(name string, c client.Client) McbAppService {
	return &mcbAppService{
		c:    c,
		name: name,
	}
}

func (c *mcbAppService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "McbApp.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for McbApp service

type McbAppHandler interface {
	Call(context.Context, *Request, *Response) error
}

func RegisterMcbAppHandler(s server.Server, hdlr McbAppHandler, opts ...server.HandlerOption) error {
	type mcbApp interface {
		Call(ctx context.Context, in *Request, out *Response) error
	}
	type McbApp struct {
		mcbApp
	}
	h := &mcbAppHandler{hdlr}
	return s.Handle(s.NewHandler(&McbApp{h}, opts...))
}

type mcbAppHandler struct {
	McbAppHandler
}

func (h *mcbAppHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.McbAppHandler.Call(ctx, in, out)
}

// Api Endpoints for McbGate service

func NewMcbGateEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for McbGate service

type McbGateService interface {
	Push(ctx context.Context, in *PushMsg, opts ...client.CallOption) (*Response, error)
	PushSession(ctx context.Context, in *Session, opts ...client.CallOption) (*Response, error)
	Bind(ctx context.Context, in *Session, opts ...client.CallOption) (*Response, error)
	Kick(ctx context.Context, in *KickMsg, opts ...client.CallOption) (*KickAnswer, error)
}

type mcbGateService struct {
	c    client.Client
	name string
}

func NewMcbGateService(name string, c client.Client) McbGateService {
	return &mcbGateService{
		c:    c,
		name: name,
	}
}

func (c *mcbGateService) Push(ctx context.Context, in *PushMsg, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "McbGate.Push", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcbGateService) PushSession(ctx context.Context, in *Session, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "McbGate.PushSession", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcbGateService) Bind(ctx context.Context, in *Session, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "McbGate.Bind", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mcbGateService) Kick(ctx context.Context, in *KickMsg, opts ...client.CallOption) (*KickAnswer, error) {
	req := c.c.NewRequest(c.name, "McbGate.Kick", in)
	out := new(KickAnswer)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for McbGate service

type McbGateHandler interface {
	Push(context.Context, *PushMsg, *Response) error
	PushSession(context.Context, *Session, *Response) error
	Bind(context.Context, *Session, *Response) error
	Kick(context.Context, *KickMsg, *KickAnswer) error
}

func RegisterMcbGateHandler(s server.Server, hdlr McbGateHandler, opts ...server.HandlerOption) error {
	type mcbGate interface {
		Push(ctx context.Context, in *PushMsg, out *Response) error
		PushSession(ctx context.Context, in *Session, out *Response) error
		Bind(ctx context.Context, in *Session, out *Response) error
		Kick(ctx context.Context, in *KickMsg, out *KickAnswer) error
	}
	type McbGate struct {
		mcbGate
	}
	h := &mcbGateHandler{hdlr}
	return s.Handle(s.NewHandler(&McbGate{h}, opts...))
}

type mcbGateHandler struct {
	McbGateHandler
}

func (h *mcbGateHandler) Push(ctx context.Context, in *PushMsg, out *Response) error {
	return h.McbGateHandler.Push(ctx, in, out)
}

func (h *mcbGateHandler) PushSession(ctx context.Context, in *Session, out *Response) error {
	return h.McbGateHandler.PushSession(ctx, in, out)
}

func (h *mcbGateHandler) Bind(ctx context.Context, in *Session, out *Response) error {
	return h.McbGateHandler.Bind(ctx, in, out)
}

func (h *mcbGateHandler) Kick(ctx context.Context, in *KickMsg, out *KickAnswer) error {
	return h.McbGateHandler.Kick(ctx, in, out)
}