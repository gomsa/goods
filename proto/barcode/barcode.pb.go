// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/barcode/barcode.proto

package barcodes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Goods
type Goods struct {
	Barcode              string   `protobuf:"bytes,1,opt,name=barcode,proto3" json:"barcode,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	EnName               string   `protobuf:"bytes,3,opt,name=en_name,json=enName,proto3" json:"en_name,omitempty"`
	Images               []string `protobuf:"bytes,4,rep,name=images,proto3" json:"images,omitempty"`
	BrandName            string   `protobuf:"bytes,5,opt,name=brand_name,json=brandName,proto3" json:"brand_name,omitempty"`
	Specification        string   `protobuf:"bytes,6,opt,name=specification,proto3" json:"specification,omitempty"`
	Unit                 string   `protobuf:"bytes,7,opt,name=unit,proto3" json:"unit,omitempty"`
	Width                int64    `protobuf:"varint,8,opt,name=width,proto3" json:"width,omitempty"`
	Height               int64    `protobuf:"varint,9,opt,name=height,proto3" json:"height,omitempty"`
	Depth                int64    `protobuf:"varint,10,opt,name=depth,proto3" json:"depth,omitempty"`
	NetWeight            int64    `protobuf:"varint,11,opt,name=net_weight,json=netWeight,proto3" json:"net_weight,omitempty"`
	GrossWeight          int64    `protobuf:"varint,12,opt,name=gross_weight,json=grossWeight,proto3" json:"gross_weight,omitempty"`
	Unspsc               int64    `protobuf:"varint,13,opt,name=unspsc,proto3" json:"unspsc,omitempty"`
	UnspscName           string   `protobuf:"bytes,14,opt,name=unspsc_name,json=unspscName,proto3" json:"unspsc_name,omitempty"`
	Source               string   `protobuf:"bytes,15,opt,name=source,proto3" json:"source,omitempty"`
	Place                string   `protobuf:"bytes,16,opt,name=place,proto3" json:"place,omitempty"`
	Country              string   `protobuf:"bytes,17,opt,name=country,proto3" json:"country,omitempty"`
	FirmName             string   `protobuf:"bytes,18,opt,name=firm_name,json=firmName,proto3" json:"firm_name,omitempty"`
	FirmAddress          string   `protobuf:"bytes,19,opt,name=firm_address,json=firmAddress,proto3" json:"firm_address,omitempty"`
	FirmStatus           string   `protobuf:"bytes,20,opt,name=firm_status,json=firmStatus,proto3" json:"firm_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Goods) Reset()         { *m = Goods{} }
func (m *Goods) String() string { return proto.CompactTextString(m) }
func (*Goods) ProtoMessage()    {}
func (*Goods) Descriptor() ([]byte, []int) {
	return fileDescriptor_630139e7c1c8d2b3, []int{0}
}

func (m *Goods) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Goods.Unmarshal(m, b)
}
func (m *Goods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Goods.Marshal(b, m, deterministic)
}
func (m *Goods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Goods.Merge(m, src)
}
func (m *Goods) XXX_Size() int {
	return xxx_messageInfo_Goods.Size(m)
}
func (m *Goods) XXX_DiscardUnknown() {
	xxx_messageInfo_Goods.DiscardUnknown(m)
}

var xxx_messageInfo_Goods proto.InternalMessageInfo

func (m *Goods) GetBarcode() string {
	if m != nil {
		return m.Barcode
	}
	return ""
}

func (m *Goods) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Goods) GetEnName() string {
	if m != nil {
		return m.EnName
	}
	return ""
}

func (m *Goods) GetImages() []string {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *Goods) GetBrandName() string {
	if m != nil {
		return m.BrandName
	}
	return ""
}

func (m *Goods) GetSpecification() string {
	if m != nil {
		return m.Specification
	}
	return ""
}

func (m *Goods) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Goods) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Goods) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Goods) GetDepth() int64 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *Goods) GetNetWeight() int64 {
	if m != nil {
		return m.NetWeight
	}
	return 0
}

func (m *Goods) GetGrossWeight() int64 {
	if m != nil {
		return m.GrossWeight
	}
	return 0
}

func (m *Goods) GetUnspsc() int64 {
	if m != nil {
		return m.Unspsc
	}
	return 0
}

func (m *Goods) GetUnspscName() string {
	if m != nil {
		return m.UnspscName
	}
	return ""
}

func (m *Goods) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Goods) GetPlace() string {
	if m != nil {
		return m.Place
	}
	return ""
}

func (m *Goods) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Goods) GetFirmName() string {
	if m != nil {
		return m.FirmName
	}
	return ""
}

func (m *Goods) GetFirmAddress() string {
	if m != nil {
		return m.FirmAddress
	}
	return ""
}

func (m *Goods) GetFirmStatus() string {
	if m != nil {
		return m.FirmStatus
	}
	return ""
}

type Request struct {
	Goods                *Goods   `protobuf:"bytes,1,opt,name=goods,proto3" json:"goods,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_630139e7c1c8d2b3, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetGoods() *Goods {
	if m != nil {
		return m.Goods
	}
	return nil
}

type Response struct {
	Goods                *Goods   `protobuf:"bytes,1,opt,name=goods,proto3" json:"goods,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_630139e7c1c8d2b3, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetGoods() *Goods {
	if m != nil {
		return m.Goods
	}
	return nil
}

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*Goods)(nil), "barcodes.Goods")
	proto.RegisterType((*Request)(nil), "barcodes.Request")
	proto.RegisterType((*Response)(nil), "barcodes.Response")
}

func init() { proto.RegisterFile("proto/barcode/barcode.proto", fileDescriptor_630139e7c1c8d2b3) }

var fileDescriptor_630139e7c1c8d2b3 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6f, 0x13, 0x31,
	0x10, 0x85, 0x09, 0x49, 0x36, 0xbb, 0xb3, 0x29, 0xa5, 0xa6, 0x2a, 0x16, 0x15, 0x22, 0x8d, 0x40,
	0xea, 0x29, 0xa0, 0x72, 0xe3, 0x06, 0x97, 0xdc, 0x38, 0x2c, 0x07, 0x8e, 0x95, 0xb3, 0x9e, 0x26,
	0x96, 0x1a, 0x7b, 0xf1, 0x78, 0xa9, 0xf8, 0x45, 0xfc, 0x4d, 0xe4, 0x99, 0x5d, 0x01, 0xb7, 0x9e,
	0xe2, 0xf7, 0xcd, 0x9b, 0xcc, 0x5b, 0x7b, 0xe0, 0xb2, 0x8b, 0x21, 0x85, 0xf7, 0x3b, 0x13, 0xdb,
	0x60, 0x71, 0xfc, 0xdd, 0x30, 0x55, 0xe5, 0x20, 0x69, 0xfd, 0x7b, 0x06, 0xf3, 0x6d, 0x08, 0x96,
	0x94, 0x86, 0xc5, 0x40, 0xf5, 0x64, 0x35, 0xb9, 0xae, 0x9a, 0x51, 0x2a, 0x05, 0x33, 0x6f, 0x8e,
	0xa8, 0x9f, 0x32, 0xe6, 0xb3, 0x7a, 0x09, 0x0b, 0xf4, 0xb7, 0x8c, 0xa7, 0x8c, 0x0b, 0xf4, 0x5f,
	0x73, 0xe1, 0x02, 0x0a, 0x77, 0x34, 0x7b, 0x24, 0x3d, 0x5b, 0x4d, 0x33, 0x17, 0xa5, 0x5e, 0x03,
	0xec, 0xa2, 0xf1, 0x56, 0x7a, 0xe6, 0xdc, 0x53, 0x31, 0xe1, 0xb6, 0xb7, 0x70, 0x42, 0x1d, 0xb6,
	0xee, 0xce, 0xb5, 0x26, 0xb9, 0xe0, 0x75, 0xc1, 0x8e, 0xff, 0x61, 0x4e, 0xd2, 0x7b, 0x97, 0xf4,
	0x42, 0x92, 0xe4, 0xb3, 0x3a, 0x87, 0xf9, 0x83, 0xb3, 0xe9, 0xa0, 0xcb, 0xd5, 0xe4, 0x7a, 0xda,
	0x88, 0xc8, 0x31, 0x0e, 0xe8, 0xf6, 0x87, 0xa4, 0x2b, 0xc6, 0x83, 0xca, 0x6e, 0x8b, 0x5d, 0x3a,
	0x68, 0x10, 0x37, 0x8b, 0x1c, 0xce, 0x63, 0xba, 0x7d, 0x90, 0x8e, 0x9a, 0x4b, 0x95, 0xc7, 0xf4,
	0x5d, 0x9a, 0xae, 0x60, 0xb9, 0x8f, 0x81, 0x68, 0x34, 0x2c, 0xd9, 0x50, 0x33, 0x1b, 0x2c, 0x17,
	0x50, 0xf4, 0x9e, 0x3a, 0x6a, 0xf5, 0x89, 0xcc, 0x13, 0xa5, 0xde, 0x40, 0x2d, 0x27, 0xf9, 0xee,
	0x67, 0x1c, 0x1c, 0x04, 0x8d, 0xf7, 0x45, 0xa1, 0x8f, 0x2d, 0xea, 0x53, 0xb9, 0x47, 0x51, 0x39,
	0x68, 0x77, 0x6f, 0x5a, 0xd4, 0xcf, 0x19, 0x8b, 0xc8, 0x8f, 0xd4, 0x86, 0xde, 0xa7, 0xf8, 0x4b,
	0x9f, 0xc9, 0x23, 0x0d, 0x52, 0x5d, 0x42, 0x75, 0xe7, 0xe2, 0x51, 0xc6, 0x28, 0xae, 0x95, 0x19,
	0xf0, 0x90, 0x2b, 0x58, 0x72, 0xd1, 0x58, 0x1b, 0x91, 0x48, 0xbf, 0xe0, 0x7a, 0x9d, 0xd9, 0x67,
	0x41, 0x39, 0x28, 0x5b, 0x28, 0x99, 0xd4, 0x93, 0x3e, 0x97, 0xa0, 0x19, 0x7d, 0x63, 0xb2, 0xfe,
	0x00, 0x8b, 0x06, 0x7f, 0xf4, 0x48, 0x49, 0xbd, 0x83, 0xf9, 0x3e, 0xef, 0x0c, 0x2f, 0x4a, 0x7d,
	0x73, 0xba, 0x19, 0xd7, 0x69, 0xc3, 0xab, 0xd4, 0x48, 0x75, 0xbd, 0x85, 0xb2, 0x41, 0xea, 0x82,
	0x27, 0x7c, 0x64, 0x4b, 0xfe, 0xea, 0x9f, 0xe6, 0xde, 0x59, 0xde, 0xb5, 0xb2, 0x11, 0x71, 0xf3,
	0x09, 0xca, 0x2f, 0x83, 0x5d, 0x6d, 0x60, 0xba, 0xc5, 0xa4, 0xce, 0xfe, 0xfe, 0xc1, 0x90, 0xea,
	0x95, 0xfa, 0x17, 0xc9, 0xd8, 0xf5, 0x93, 0x5d, 0xc1, 0x1b, 0xff, 0xf1, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x82, 0xcf, 0xb0, 0x6a, 0x10, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Barcodes service

type BarcodesClient interface {
	// 通过条形查询商品信息
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type barcodesClient struct {
	c           client.Client
	serviceName string
}

func NewBarcodesClient(serviceName string, c client.Client) BarcodesClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "barcodes"
	}
	return &barcodesClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *barcodesClient) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Barcodes.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Barcodes service

type BarcodesHandler interface {
	// 通过条形查询商品信息
	Get(context.Context, *Request, *Response) error
}

func RegisterBarcodesHandler(s server.Server, hdlr BarcodesHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Barcodes{hdlr}, opts...))
}

type Barcodes struct {
	BarcodesHandler
}

func (h *Barcodes) Get(ctx context.Context, in *Request, out *Response) error {
	return h.BarcodesHandler.Get(ctx, in, out)
}
