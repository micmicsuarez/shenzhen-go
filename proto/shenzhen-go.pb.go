//+build !js
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shenzhen-go.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	shenzhen-go.proto

It has these top-level messages:
	Empty
	NodeConfig
	CreateChannelRequest
	CreateNodeRequest
	ConnectPinRequest
	DeleteChannelRequest
	DeleteNodeRequest
	DisconnectPinRequest
	SaveRequest
	SetGraphPropertiesRequest
	SetNodePropertiesRequest
	SetPositionRequest
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto1.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type NodeConfig struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Enabled      bool   `protobuf:"varint,2,opt,name=enabled" json:"enabled,omitempty"`
	Multiplicity uint32 `protobuf:"varint,3,opt,name=multiplicity" json:"multiplicity,omitempty"`
	Wait         bool   `protobuf:"varint,4,opt,name=wait" json:"wait,omitempty"`
	PartCfg      []byte `protobuf:"bytes,5,opt,name=part_cfg,json=partCfg,proto3" json:"part_cfg,omitempty"`
	PartType     string `protobuf:"bytes,6,opt,name=part_type,json=partType" json:"part_type,omitempty"`
	X            int64  `protobuf:"varint,7,opt,name=x" json:"x,omitempty"`
	Y            int64  `protobuf:"varint,8,opt,name=y" json:"y,omitempty"`
}

func (m *NodeConfig) Reset()                    { *m = NodeConfig{} }
func (m *NodeConfig) String() string            { return proto1.CompactTextString(m) }
func (*NodeConfig) ProtoMessage()               {}
func (*NodeConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NodeConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeConfig) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *NodeConfig) GetMultiplicity() uint32 {
	if m != nil {
		return m.Multiplicity
	}
	return 0
}

func (m *NodeConfig) GetWait() bool {
	if m != nil {
		return m.Wait
	}
	return false
}

func (m *NodeConfig) GetPartCfg() []byte {
	if m != nil {
		return m.PartCfg
	}
	return nil
}

func (m *NodeConfig) GetPartType() string {
	if m != nil {
		return m.PartType
	}
	return ""
}

func (m *NodeConfig) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *NodeConfig) GetY() int64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type CreateChannelRequest struct {
	Graph string `protobuf:"bytes,1,opt,name=graph" json:"graph,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type  string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Anon  bool   `protobuf:"varint,4,opt,name=anon" json:"anon,omitempty"`
	Cap   uint64 `protobuf:"varint,5,opt,name=cap" json:"cap,omitempty"`
	// Also connect two pins together, to reduce chatter (1 Create RPC instead of [Create, Connect, Connect])
	Node1 string `protobuf:"bytes,6,opt,name=node1" json:"node1,omitempty"`
	Pin1  string `protobuf:"bytes,7,opt,name=pin1" json:"pin1,omitempty"`
	Node2 string `protobuf:"bytes,8,opt,name=node2" json:"node2,omitempty"`
	Pin2  string `protobuf:"bytes,9,opt,name=pin2" json:"pin2,omitempty"`
}

func (m *CreateChannelRequest) Reset()                    { *m = CreateChannelRequest{} }
func (m *CreateChannelRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateChannelRequest) ProtoMessage()               {}
func (*CreateChannelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateChannelRequest) GetGraph() string {
	if m != nil {
		return m.Graph
	}
	return ""
}

func (m *CreateChannelRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateChannelRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateChannelRequest) GetAnon() bool {
	if m != nil {
		return m.Anon
	}
	return false
}

func (m *CreateChannelRequest) GetCap() uint64 {
	if m != nil {
		return m.Cap
	}
	return 0
}

func (m *CreateChannelRequest) GetNode1() string {
	if m != nil {
		return m.Node1
	}
	return ""
}

func (m *CreateChannelRequest) GetPin1() string {
	if m != nil {
		return m.Pin1
	}
	return ""
}

func (m *CreateChannelRequest) GetNode2() string {
	if m != nil {
		return m.Node2
	}
	return ""
}

func (m *CreateChannelRequest) GetPin2() string {
	if m != nil {
		return m.Pin2
	}
	return ""
}

type CreateNodeRequest struct {
	Graph string      `protobuf:"bytes,1,opt,name=graph" json:"graph,omitempty"`
	Props *NodeConfig `protobuf:"bytes,2,opt,name=props" json:"props,omitempty"`
}

func (m *CreateNodeRequest) Reset()                    { *m = CreateNodeRequest{} }
func (m *CreateNodeRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateNodeRequest) ProtoMessage()               {}
func (*CreateNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateNodeRequest) GetGraph() string {
	if m != nil {
		return m.Graph
	}
	return ""
}

func (m *CreateNodeRequest) GetProps() *NodeConfig {
	if m != nil {
		return m.Props
	}
	return nil
}

type ConnectPinRequest struct {
	Graph   string `protobuf:"bytes,1,opt,name=graph" json:"graph,omitempty"`
	Node    string `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
	Pin     string `protobuf:"bytes,3,opt,name=pin" json:"pin,omitempty"`
	Channel string `protobuf:"bytes,4,opt,name=channel" json:"channel,omitempty"`
}

func (m *ConnectPinRequest) Reset()                    { *m = ConnectPinRequest{} }
func (m *ConnectPinRequest) String() string            { return proto1.CompactTextString(m) }
func (*ConnectPinRequest) ProtoMessage()               {}
func (*ConnectPinRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ConnectPinRequest) GetGraph() string {
	if m != nil {
		return m.Graph
	}
	return ""
}

func (m *ConnectPinRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *ConnectPinRequest) GetPin() string {
	if m != nil {
		return m.Pin
	}
	return ""
}

func (m *ConnectPinRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

type DeleteChannelRequest struct {
	Graph   string `protobuf:"bytes,1,opt,name=graph" json:"graph,omitempty"`
	Channel string `protobuf:"bytes,2,opt,name=channel" json:"channel,omitempty"`
}

func (m *DeleteChannelRequest) Reset()                    { *m = DeleteChannelRequest{} }
func (m *DeleteChannelRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeleteChannelRequest) ProtoMessage()               {}
func (*DeleteChannelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteChannelRequest) GetGraph() string {
	if m != nil {
		return m.Graph
	}
	return ""
}

func (m *DeleteChannelRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

type DeleteNodeRequest struct {
	Graph string `protobuf:"bytes,1,opt,name=graph" json:"graph,omitempty"`
	Node  string `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
}

func (m *DeleteNodeRequest) Reset()                    { *m = DeleteNodeRequest{} }
func (m *DeleteNodeRequest) String() string            { return proto1.CompactTextString(m) }
func (*DeleteNodeRequest) ProtoMessage()               {}
func (*DeleteNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteNodeRequest) GetGraph() string {
	if m != nil {
		return m.Graph
	}
	return ""
}

func (m *DeleteNodeRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type DisconnectPinRequest struct {
	Graph string `protobuf:"bytes,1,opt,name=graph" json:"graph,omitempty"`
	Node  string `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
	Pin   string `protobuf:"bytes,3,opt,name=pin" json:"pin,omitempty"`
}

func (m *DisconnectPinRequest) Reset()                    { *m = DisconnectPinRequest{} }
func (m *DisconnectPinRequest) String() string            { return proto1.CompactTextString(m) }
func (*DisconnectPinRequest) ProtoMessage()               {}
func (*DisconnectPinRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DisconnectPinRequest) GetGraph() string {
	if m != nil {
		return m.Graph
	}
	return ""
}

func (m *DisconnectPinRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *DisconnectPinRequest) GetPin() string {
	if m != nil {
		return m.Pin
	}
	return ""
}

type SaveRequest struct {
	Graph string `protobuf:"bytes,1,opt,name=graph" json:"graph,omitempty"`
}

func (m *SaveRequest) Reset()                    { *m = SaveRequest{} }
func (m *SaveRequest) String() string            { return proto1.CompactTextString(m) }
func (*SaveRequest) ProtoMessage()               {}
func (*SaveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SaveRequest) GetGraph() string {
	if m != nil {
		return m.Graph
	}
	return ""
}

type SetGraphPropertiesRequest struct {
	Graph       string `protobuf:"bytes,1,opt,name=graph" json:"graph,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	PackagePath string `protobuf:"bytes,3,opt,name=package_path,json=packagePath" json:"package_path,omitempty"`
	IsCommand   bool   `protobuf:"varint,4,opt,name=is_command,json=isCommand" json:"is_command,omitempty"`
}

func (m *SetGraphPropertiesRequest) Reset()                    { *m = SetGraphPropertiesRequest{} }
func (m *SetGraphPropertiesRequest) String() string            { return proto1.CompactTextString(m) }
func (*SetGraphPropertiesRequest) ProtoMessage()               {}
func (*SetGraphPropertiesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *SetGraphPropertiesRequest) GetGraph() string {
	if m != nil {
		return m.Graph
	}
	return ""
}

func (m *SetGraphPropertiesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetGraphPropertiesRequest) GetPackagePath() string {
	if m != nil {
		return m.PackagePath
	}
	return ""
}

func (m *SetGraphPropertiesRequest) GetIsCommand() bool {
	if m != nil {
		return m.IsCommand
	}
	return false
}

type SetNodePropertiesRequest struct {
	Graph string      `protobuf:"bytes,1,opt,name=graph" json:"graph,omitempty"`
	Node  string      `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
	Props *NodeConfig `protobuf:"bytes,3,opt,name=props" json:"props,omitempty"`
}

func (m *SetNodePropertiesRequest) Reset()                    { *m = SetNodePropertiesRequest{} }
func (m *SetNodePropertiesRequest) String() string            { return proto1.CompactTextString(m) }
func (*SetNodePropertiesRequest) ProtoMessage()               {}
func (*SetNodePropertiesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SetNodePropertiesRequest) GetGraph() string {
	if m != nil {
		return m.Graph
	}
	return ""
}

func (m *SetNodePropertiesRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *SetNodePropertiesRequest) GetProps() *NodeConfig {
	if m != nil {
		return m.Props
	}
	return nil
}

type SetPositionRequest struct {
	Graph string `protobuf:"bytes,1,opt,name=graph" json:"graph,omitempty"`
	Node  string `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
	X     int64  `protobuf:"varint,3,opt,name=x" json:"x,omitempty"`
	Y     int64  `protobuf:"varint,4,opt,name=y" json:"y,omitempty"`
}

func (m *SetPositionRequest) Reset()                    { *m = SetPositionRequest{} }
func (m *SetPositionRequest) String() string            { return proto1.CompactTextString(m) }
func (*SetPositionRequest) ProtoMessage()               {}
func (*SetPositionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *SetPositionRequest) GetGraph() string {
	if m != nil {
		return m.Graph
	}
	return ""
}

func (m *SetPositionRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *SetPositionRequest) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *SetPositionRequest) GetY() int64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func init() {
	proto1.RegisterType((*Empty)(nil), "proto.Empty")
	proto1.RegisterType((*NodeConfig)(nil), "proto.NodeConfig")
	proto1.RegisterType((*CreateChannelRequest)(nil), "proto.CreateChannelRequest")
	proto1.RegisterType((*CreateNodeRequest)(nil), "proto.CreateNodeRequest")
	proto1.RegisterType((*ConnectPinRequest)(nil), "proto.ConnectPinRequest")
	proto1.RegisterType((*DeleteChannelRequest)(nil), "proto.DeleteChannelRequest")
	proto1.RegisterType((*DeleteNodeRequest)(nil), "proto.DeleteNodeRequest")
	proto1.RegisterType((*DisconnectPinRequest)(nil), "proto.DisconnectPinRequest")
	proto1.RegisterType((*SaveRequest)(nil), "proto.SaveRequest")
	proto1.RegisterType((*SetGraphPropertiesRequest)(nil), "proto.SetGraphPropertiesRequest")
	proto1.RegisterType((*SetNodePropertiesRequest)(nil), "proto.SetNodePropertiesRequest")
	proto1.RegisterType((*SetPositionRequest)(nil), "proto.SetPositionRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ShenzhenGo service

type ShenzhenGoClient interface {
	// CreateChannel makes a new channel.
	CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*Empty, error)
	// CreateNode makes a new node.
	CreateNode(ctx context.Context, in *CreateNodeRequest, opts ...grpc.CallOption) (*Empty, error)
	// ConnectPin connects a pin to a channel.
	ConnectPin(ctx context.Context, in *ConnectPinRequest, opts ...grpc.CallOption) (*Empty, error)
	// DeleteChannel deletes a channel (and all connections).
	DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*Empty, error)
	// DeleteNode deletes a node (and all connections).
	DeleteNode(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*Empty, error)
	// DisconnectPin deletes the connection from a pin to a channel.
	DisconnectPin(ctx context.Context, in *DisconnectPinRequest, opts ...grpc.CallOption) (*Empty, error)
	// Save saves the graph to disk.
	Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*Empty, error)
	// SetGraphProperties changes metdata such as name and package path.
	SetGraphProperties(ctx context.Context, in *SetGraphPropertiesRequest, opts ...grpc.CallOption) (*Empty, error)
	// SetNodeProperties changes node metadata such as name and multiplicity.
	SetNodeProperties(ctx context.Context, in *SetNodePropertiesRequest, opts ...grpc.CallOption) (*Empty, error)
	// SetPosition changes the node position in the diagram.
	SetPosition(ctx context.Context, in *SetPositionRequest, opts ...grpc.CallOption) (*Empty, error)
}

type shenzhenGoClient struct {
	cc *grpc.ClientConn
}

func NewShenzhenGoClient(cc *grpc.ClientConn) ShenzhenGoClient {
	return &shenzhenGoClient{cc}
}

func (c *shenzhenGoClient) CreateChannel(ctx context.Context, in *CreateChannelRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.ShenzhenGo/CreateChannel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shenzhenGoClient) CreateNode(ctx context.Context, in *CreateNodeRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.ShenzhenGo/CreateNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shenzhenGoClient) ConnectPin(ctx context.Context, in *ConnectPinRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.ShenzhenGo/ConnectPin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shenzhenGoClient) DeleteChannel(ctx context.Context, in *DeleteChannelRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.ShenzhenGo/DeleteChannel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shenzhenGoClient) DeleteNode(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.ShenzhenGo/DeleteNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shenzhenGoClient) DisconnectPin(ctx context.Context, in *DisconnectPinRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.ShenzhenGo/DisconnectPin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shenzhenGoClient) Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.ShenzhenGo/Save", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shenzhenGoClient) SetGraphProperties(ctx context.Context, in *SetGraphPropertiesRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.ShenzhenGo/SetGraphProperties", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shenzhenGoClient) SetNodeProperties(ctx context.Context, in *SetNodePropertiesRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.ShenzhenGo/SetNodeProperties", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shenzhenGoClient) SetPosition(ctx context.Context, in *SetPositionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/proto.ShenzhenGo/SetPosition", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShenzhenGo service

type ShenzhenGoServer interface {
	// CreateChannel makes a new channel.
	CreateChannel(context.Context, *CreateChannelRequest) (*Empty, error)
	// CreateNode makes a new node.
	CreateNode(context.Context, *CreateNodeRequest) (*Empty, error)
	// ConnectPin connects a pin to a channel.
	ConnectPin(context.Context, *ConnectPinRequest) (*Empty, error)
	// DeleteChannel deletes a channel (and all connections).
	DeleteChannel(context.Context, *DeleteChannelRequest) (*Empty, error)
	// DeleteNode deletes a node (and all connections).
	DeleteNode(context.Context, *DeleteNodeRequest) (*Empty, error)
	// DisconnectPin deletes the connection from a pin to a channel.
	DisconnectPin(context.Context, *DisconnectPinRequest) (*Empty, error)
	// Save saves the graph to disk.
	Save(context.Context, *SaveRequest) (*Empty, error)
	// SetGraphProperties changes metdata such as name and package path.
	SetGraphProperties(context.Context, *SetGraphPropertiesRequest) (*Empty, error)
	// SetNodeProperties changes node metadata such as name and multiplicity.
	SetNodeProperties(context.Context, *SetNodePropertiesRequest) (*Empty, error)
	// SetPosition changes the node position in the diagram.
	SetPosition(context.Context, *SetPositionRequest) (*Empty, error)
}

func RegisterShenzhenGoServer(s *grpc.Server, srv ShenzhenGoServer) {
	s.RegisterService(&_ShenzhenGo_serviceDesc, srv)
}

func _ShenzhenGo_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShenzhenGoServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ShenzhenGo/CreateChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShenzhenGoServer).CreateChannel(ctx, req.(*CreateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShenzhenGo_CreateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShenzhenGoServer).CreateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ShenzhenGo/CreateNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShenzhenGoServer).CreateNode(ctx, req.(*CreateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShenzhenGo_ConnectPin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectPinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShenzhenGoServer).ConnectPin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ShenzhenGo/ConnectPin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShenzhenGoServer).ConnectPin(ctx, req.(*ConnectPinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShenzhenGo_DeleteChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShenzhenGoServer).DeleteChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ShenzhenGo/DeleteChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShenzhenGoServer).DeleteChannel(ctx, req.(*DeleteChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShenzhenGo_DeleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShenzhenGoServer).DeleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ShenzhenGo/DeleteNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShenzhenGoServer).DeleteNode(ctx, req.(*DeleteNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShenzhenGo_DisconnectPin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectPinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShenzhenGoServer).DisconnectPin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ShenzhenGo/DisconnectPin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShenzhenGoServer).DisconnectPin(ctx, req.(*DisconnectPinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShenzhenGo_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShenzhenGoServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ShenzhenGo/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShenzhenGoServer).Save(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShenzhenGo_SetGraphProperties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGraphPropertiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShenzhenGoServer).SetGraphProperties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ShenzhenGo/SetGraphProperties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShenzhenGoServer).SetGraphProperties(ctx, req.(*SetGraphPropertiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShenzhenGo_SetNodeProperties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetNodePropertiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShenzhenGoServer).SetNodeProperties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ShenzhenGo/SetNodeProperties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShenzhenGoServer).SetNodeProperties(ctx, req.(*SetNodePropertiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShenzhenGo_SetPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShenzhenGoServer).SetPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ShenzhenGo/SetPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShenzhenGoServer).SetPosition(ctx, req.(*SetPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShenzhenGo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ShenzhenGo",
	HandlerType: (*ShenzhenGoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChannel",
			Handler:    _ShenzhenGo_CreateChannel_Handler,
		},
		{
			MethodName: "CreateNode",
			Handler:    _ShenzhenGo_CreateNode_Handler,
		},
		{
			MethodName: "ConnectPin",
			Handler:    _ShenzhenGo_ConnectPin_Handler,
		},
		{
			MethodName: "DeleteChannel",
			Handler:    _ShenzhenGo_DeleteChannel_Handler,
		},
		{
			MethodName: "DeleteNode",
			Handler:    _ShenzhenGo_DeleteNode_Handler,
		},
		{
			MethodName: "DisconnectPin",
			Handler:    _ShenzhenGo_DisconnectPin_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _ShenzhenGo_Save_Handler,
		},
		{
			MethodName: "SetGraphProperties",
			Handler:    _ShenzhenGo_SetGraphProperties_Handler,
		},
		{
			MethodName: "SetNodeProperties",
			Handler:    _ShenzhenGo_SetNodeProperties_Handler,
		},
		{
			MethodName: "SetPosition",
			Handler:    _ShenzhenGo_SetPosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shenzhen-go.proto",
}

func init() { proto1.RegisterFile("shenzhen-go.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x65, 0x9b, 0x4d, 0xd3, 0x4c, 0x52, 0x89, 0x58, 0x39, 0xb8, 0x54, 0x88, 0xb0, 0x1c, 0x88,
	0x90, 0xa8, 0xd4, 0x20, 0x21, 0x0e, 0xe5, 0x94, 0xd2, 0xde, 0x50, 0xe4, 0x70, 0xe2, 0x52, 0xb9,
	0x9b, 0x69, 0x62, 0x91, 0xd8, 0x26, 0xeb, 0x42, 0xc3, 0x0f, 0xf0, 0x67, 0x1c, 0xf9, 0x08, 0xbe,
	0x04, 0xd9, 0xeb, 0x64, 0x77, 0xb3, 0x4b, 0x4b, 0x11, 0xa7, 0x9d, 0xf1, 0x7a, 0x66, 0xde, 0xcc,
	0xbc, 0x67, 0xe8, 0x24, 0x33, 0x94, 0xdf, 0x66, 0x28, 0x5f, 0x4e, 0xd5, 0x91, 0x5e, 0x2a, 0xa3,
	0x48, 0xdd, 0x7d, 0xa2, 0x06, 0xd4, 0xdf, 0x2d, 0xb4, 0x59, 0x45, 0x3f, 0x02, 0x80, 0xf7, 0x6a,
	0x82, 0x43, 0x25, 0xaf, 0xc4, 0x94, 0x10, 0x08, 0x25, 0x5f, 0x20, 0x0d, 0x7a, 0x41, 0xbf, 0xc9,
	0x9c, 0x4d, 0x28, 0x34, 0x50, 0xf2, 0xcb, 0x39, 0x4e, 0xe8, 0x4e, 0x2f, 0xe8, 0xef, 0xb1, 0xb5,
	0x4b, 0x22, 0x68, 0x2f, 0xae, 0xe7, 0x46, 0xe8, 0xb9, 0x88, 0x85, 0x59, 0xd1, 0x5a, 0x2f, 0xe8,
	0xef, 0xb3, 0xc2, 0x99, 0xcd, 0xf8, 0x95, 0x0b, 0x43, 0x43, 0x17, 0xea, 0x6c, 0x72, 0x00, 0x7b,
	0x9a, 0x2f, 0xcd, 0x45, 0x7c, 0x35, 0xa5, 0xf5, 0x5e, 0xd0, 0x6f, 0xb3, 0x86, 0xf5, 0x87, 0x57,
	0x53, 0x72, 0x08, 0x4d, 0xf7, 0xcb, 0xac, 0x34, 0xd2, 0x5d, 0x87, 0xc2, 0xdd, 0xfd, 0xb0, 0xd2,
	0x48, 0xda, 0x10, 0xdc, 0xd0, 0x46, 0x2f, 0xe8, 0xd7, 0x58, 0x70, 0x63, 0xbd, 0x15, 0xdd, 0x4b,
	0xbd, 0x55, 0xf4, 0x33, 0x80, 0xee, 0x70, 0x89, 0xdc, 0xe0, 0x70, 0xc6, 0xa5, 0xc4, 0x39, 0xc3,
	0xcf, 0xd7, 0x98, 0x18, 0xd2, 0x85, 0xfa, 0x74, 0xc9, 0xf5, 0xcc, 0xf7, 0x94, 0x3a, 0x9b, 0x46,
	0x77, 0x72, 0x8d, 0x12, 0x08, 0x5d, 0xd9, 0x5a, 0x7a, 0x66, 0x6d, 0x7b, 0xc6, 0xa5, 0x92, 0x6b,
	0xf8, 0xd6, 0x26, 0x0f, 0xa1, 0x16, 0x73, 0xed, 0x90, 0x87, 0xcc, 0x9a, 0xb6, 0x86, 0x54, 0x13,
	0x3c, 0xf6, 0x88, 0x53, 0xc7, 0xc6, 0x6a, 0x21, 0x8f, 0x1d, 0xe2, 0x26, 0x73, 0xf6, 0xfa, 0xe6,
	0xc0, 0x01, 0xf7, 0x37, 0x07, 0xfe, 0xe6, 0x80, 0x36, 0x37, 0x37, 0x07, 0x11, 0x83, 0x4e, 0xda,
	0x8f, 0x5d, 0xcf, 0xed, 0xcd, 0x3c, 0x07, 0xbb, 0x56, 0x9d, 0xb8, 0x6e, 0x5a, 0x83, 0x4e, 0xba,
	0xeb, 0xa3, 0x6c, 0xaf, 0x2c, 0xfd, 0x1f, 0x09, 0xe8, 0x0c, 0x95, 0x94, 0x18, 0x9b, 0x91, 0x90,
	0x77, 0x0f, 0x48, 0x4d, 0xb2, 0x01, 0xa9, 0x09, 0xda, 0xc6, 0xb5, 0x90, 0x7e, 0x3e, 0xd6, 0xb4,
	0xdc, 0x88, 0xd3, 0x71, 0xbb, 0x09, 0x35, 0xd9, 0xda, 0x8d, 0xce, 0xa0, 0x7b, 0x8a, 0x73, 0xfc,
	0xcb, 0x75, 0xe4, 0xf2, 0xec, 0x14, 0xf3, 0xbc, 0x85, 0x4e, 0x9a, 0xe7, 0xee, 0x31, 0x54, 0x40,
	0x8e, 0x18, 0x74, 0x4f, 0x45, 0x12, 0xff, 0xcf, 0xa6, 0xa3, 0x67, 0xd0, 0x1a, 0xf3, 0x2f, 0xb7,
	0x83, 0x89, 0xbe, 0x07, 0x70, 0x30, 0x46, 0x73, 0x6e, 0x9d, 0xd1, 0x52, 0x69, 0x5c, 0x1a, 0x81,
	0xc9, 0xfd, 0x49, 0xf9, 0x14, 0xda, 0x9a, 0xc7, 0x9f, 0xf8, 0x14, 0x2f, 0x34, 0x37, 0x33, 0x8f,
	0xa3, 0xe5, 0xcf, 0x46, 0xdc, 0xcc, 0xc8, 0x63, 0x00, 0x91, 0x5c, 0xc4, 0x6a, 0xb1, 0xe0, 0x72,
	0xe2, 0x99, 0xda, 0x14, 0xc9, 0x30, 0x3d, 0x88, 0x16, 0x40, 0xc7, 0x68, 0xec, 0xf8, 0xee, 0x83,
	0x63, 0x7b, 0x0c, 0x1b, 0x8e, 0xd5, 0xee, 0xe0, 0xd8, 0x47, 0x20, 0x63, 0x34, 0x23, 0x95, 0x08,
	0x23, 0xd4, 0x3f, 0xcc, 0xdb, 0x89, 0xbc, 0x56, 0x10, 0x79, 0xe8, 0x45, 0x3e, 0xf8, 0x15, 0x02,
	0x8c, 0xfd, 0x9b, 0x76, 0xae, 0xc8, 0x09, 0xec, 0x17, 0x24, 0x4f, 0x0e, 0x3d, 0xaa, 0xaa, 0x87,
	0xe0, 0x51, 0xdb, 0xff, 0x4c, 0x1f, 0xbe, 0x07, 0xe4, 0x35, 0x40, 0x26, 0x30, 0x42, 0x0b, 0xa1,
	0x39, 0xb2, 0x55, 0xc6, 0x6d, 0xf8, 0x94, 0xc5, 0x6d, 0x53, 0xac, 0x14, 0x77, 0x02, 0xfb, 0x05,
	0x45, 0x6c, 0xd0, 0x56, 0xe9, 0xa4, 0xaa, 0x6a, 0xa6, 0x83, 0x4d, 0xd5, 0x92, 0x34, 0x2a, 0xab,
	0xe6, 0x05, 0x90, 0x55, 0xad, 0x90, 0x45, 0x29, 0xfa, 0x05, 0x84, 0x96, 0xea, 0x84, 0xf8, 0xf3,
	0x1c, 0xef, 0x4b, 0x77, 0xcf, 0xdc, 0xe2, 0xb7, 0x08, 0x4f, 0x7a, 0xeb, 0xc8, 0x3f, 0x69, 0xa1,
	0x94, 0xe7, 0x14, 0x3a, 0x25, 0xbe, 0x92, 0x27, 0x59, 0x9a, 0x4a, 0x26, 0x97, 0xb2, 0xbc, 0x81,
	0x56, 0x8e, 0x86, 0xe4, 0x20, 0x8b, 0xdf, 0xa2, 0xe6, 0x76, 0xe4, 0xe5, 0xae, 0x73, 0x5f, 0xfd,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x96, 0x15, 0xc5, 0x3a, 0x3e, 0x07, 0x00, 0x00,
}
