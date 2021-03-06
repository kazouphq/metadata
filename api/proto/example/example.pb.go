// Code generated by protoc-gen-go.
// source: github.com/kazouphq/metadata/api/proto/example/example.proto
// DO NOT EDIT!

/*
Package go_micro_api_metadata is a generated protocol buffer package.

It is generated from these files:
	github.com/kazouphq/metadata/api/proto/example/example.proto

It has these top-level messages:
*/
package go_micro_api_metadata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_micro_api "github.com/micro/micro/api/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

// Client API for Metadata service

type MetadataClient interface {
	Send(ctx context.Context, in *go_micro_api.Request, opts ...client.CallOption) (*go_micro_api.Response, error)
}

type metadataClient struct {
	c           client.Client
	serviceName string
}

func NewMetadataClient(serviceName string, c client.Client) MetadataClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.api.metadata"
	}
	return &metadataClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *metadataClient) Send(ctx context.Context, in *go_micro_api.Request, opts ...client.CallOption) (*go_micro_api.Response, error) {
	req := c.c.NewRequest(c.serviceName, "Metadata.Send", in)
	out := new(go_micro_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Metadata service

type MetadataHandler interface {
	Send(context.Context, *go_micro_api.Request, *go_micro_api.Response) error
}

func RegisterMetadataHandler(s server.Server, hdlr MetadataHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Metadata{hdlr}, opts...))
}

type Metadata struct {
	MetadataHandler
}

func (h *Metadata) Send(ctx context.Context, in *go_micro_api.Request, out *go_micro_api.Response) error {
	return h.MetadataHandler.Send(ctx, in, out)
}

func init() {
	proto.RegisterFile("github.com/kazouphq/metadata/api/proto/example/example.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x4e, 0xac, 0xca, 0x2f, 0x2d, 0xc8, 0x28, 0xd4,
	0xcf, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0xd4, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x28, 0xca,
	0x2f, 0xc9, 0xd7, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x85, 0xd1, 0x7a, 0x60, 0x51, 0x21,
	0xd1, 0xf4, 0x7c, 0xbd, 0xdc, 0xcc, 0xe4, 0xa2, 0x7c, 0x3d, 0xa0, 0x4a, 0x3d, 0x98, 0x36, 0x29,
	0x2d, 0x24, 0x43, 0xc1, 0xd2, 0x50, 0x12, 0x61, 0x1c, 0x48, 0x39, 0x98, 0x65, 0xe4, 0xcc, 0xc5,
	0xe1, 0x0b, 0xd5, 0x27, 0x64, 0xce, 0xc5, 0x12, 0x9c, 0x9a, 0x97, 0x22, 0x24, 0xaa, 0x87, 0x62,
	0x6e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x94, 0x18, 0xba, 0x70, 0x71, 0x41, 0x7e, 0x5e,
	0x71, 0xaa, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x2c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb7,
	0xd6, 0x8d, 0xcf, 0xce, 0x00, 0x00, 0x00,
}
