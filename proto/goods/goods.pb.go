// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/goods/goods.proto

package goods

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// 数据库 goods
type Good struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	EngName              string     `protobuf:"bytes,3,opt,name=eng_name,json=engName,proto3" json:"eng_name,omitempty"`
	Description          string     `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Status               int64      `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	BrandId              int64      `protobuf:"varint,6,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	CategoryId           int64      `protobuf:"varint,7,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	DepartmentId         int64      `protobuf:"varint,8,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
	FirmId               int64      `protobuf:"varint,9,opt,name=firm_id,json=firmId,proto3" json:"firm_id,omitempty"`
	UnspscId             int64      `protobuf:"varint,10,opt,name=unspsc_id,json=unspscId,proto3" json:"unspsc_id,omitempty"`
	TaxcodeId            int64      `protobuf:"varint,11,opt,name=taxcode_id,json=taxcodeId,proto3" json:"taxcode_id,omitempty"`
	Cess                 int64      `protobuf:"varint,12,opt,name=cess,proto3" json:"cess,omitempty"`
	CreatedAt            string     `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string     `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Barcodes             []*Barcode `protobuf:"bytes,15,rep,name=Barcodes,proto3" json:"Barcodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Good) Reset()         { *m = Good{} }
func (m *Good) String() string { return proto.CompactTextString(m) }
func (*Good) ProtoMessage()    {}
func (*Good) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ff7bf58be59df1, []int{0}
}

func (m *Good) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Good.Unmarshal(m, b)
}
func (m *Good) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Good.Marshal(b, m, deterministic)
}
func (m *Good) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Good.Merge(m, src)
}
func (m *Good) XXX_Size() int {
	return xxx_messageInfo_Good.Size(m)
}
func (m *Good) XXX_DiscardUnknown() {
	xxx_messageInfo_Good.DiscardUnknown(m)
}

var xxx_messageInfo_Good proto.InternalMessageInfo

func (m *Good) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Good) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Good) GetEngName() string {
	if m != nil {
		return m.EngName
	}
	return ""
}

func (m *Good) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Good) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Good) GetBrandId() int64 {
	if m != nil {
		return m.BrandId
	}
	return 0
}

func (m *Good) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *Good) GetDepartmentId() int64 {
	if m != nil {
		return m.DepartmentId
	}
	return 0
}

func (m *Good) GetFirmId() int64 {
	if m != nil {
		return m.FirmId
	}
	return 0
}

func (m *Good) GetUnspscId() int64 {
	if m != nil {
		return m.UnspscId
	}
	return 0
}

func (m *Good) GetTaxcodeId() int64 {
	if m != nil {
		return m.TaxcodeId
	}
	return 0
}

func (m *Good) GetCess() int64 {
	if m != nil {
		return m.Cess
	}
	return 0
}

func (m *Good) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Good) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Good) GetBarcodes() []*Barcode {
	if m != nil {
		return m.Barcodes
	}
	return nil
}

// 数据库 barcodes
type Barcode struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GoodsId              string   `protobuf:"bytes,2,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
	StockId              string   `protobuf:"bytes,3,opt,name=stock_id,json=stockId,proto3" json:"stock_id,omitempty"`
	Price                int64    `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Width                int64    `protobuf:"varint,5,opt,name=width,proto3" json:"width,omitempty"`
	Height               int64    `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
	Depth                int64    `protobuf:"varint,7,opt,name=depth,proto3" json:"depth,omitempty"`
	Unit                 string   `protobuf:"bytes,8,opt,name=unit,proto3" json:"unit,omitempty"`
	Spec                 string   `protobuf:"bytes,9,opt,name=spec,proto3" json:"spec,omitempty"`
	Grossweight          int64    `protobuf:"varint,10,opt,name=grossweight,proto3" json:"grossweight,omitempty"`
	Netweight            int64    `protobuf:"varint,11,opt,name=netweight,proto3" json:"netweight,omitempty"`
	Img                  string   `protobuf:"bytes,12,opt,name=img,proto3" json:"img,omitempty"`
	BarcodeImg           string   `protobuf:"bytes,13,opt,name=barcode_img,json=barcodeImg,proto3" json:"barcode_img,omitempty"`
	CreatedAt            string   `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Barcode) Reset()         { *m = Barcode{} }
func (m *Barcode) String() string { return proto.CompactTextString(m) }
func (*Barcode) ProtoMessage()    {}
func (*Barcode) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ff7bf58be59df1, []int{1}
}

func (m *Barcode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Barcode.Unmarshal(m, b)
}
func (m *Barcode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Barcode.Marshal(b, m, deterministic)
}
func (m *Barcode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Barcode.Merge(m, src)
}
func (m *Barcode) XXX_Size() int {
	return xxx_messageInfo_Barcode.Size(m)
}
func (m *Barcode) XXX_DiscardUnknown() {
	xxx_messageInfo_Barcode.DiscardUnknown(m)
}

var xxx_messageInfo_Barcode proto.InternalMessageInfo

func (m *Barcode) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Barcode) GetGoodsId() string {
	if m != nil {
		return m.GoodsId
	}
	return ""
}

func (m *Barcode) GetStockId() string {
	if m != nil {
		return m.StockId
	}
	return ""
}

func (m *Barcode) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Barcode) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Barcode) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Barcode) GetDepth() int64 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *Barcode) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Barcode) GetSpec() string {
	if m != nil {
		return m.Spec
	}
	return ""
}

func (m *Barcode) GetGrossweight() int64 {
	if m != nil {
		return m.Grossweight
	}
	return 0
}

func (m *Barcode) GetNetweight() int64 {
	if m != nil {
		return m.Netweight
	}
	return 0
}

func (m *Barcode) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *Barcode) GetBarcodeImg() string {
	if m != nil {
		return m.BarcodeImg
	}
	return ""
}

func (m *Barcode) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Barcode) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

// 商品信息 抽象对象不对应数据库(有用查询数据返回)
type GoodsInfo struct {
	Good                 *Good    `protobuf:"bytes,1,opt,name=good,proto3" json:"good,omitempty"`
	Brand                string   `protobuf:"bytes,3,opt,name=brand,proto3" json:"brand,omitempty"`
	Category             string   `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Department           string   `protobuf:"bytes,5,opt,name=department,proto3" json:"department,omitempty"`
	Firm                 string   `protobuf:"bytes,6,opt,name=firm,proto3" json:"firm,omitempty"`
	Unspsc               string   `protobuf:"bytes,7,opt,name=unspsc,proto3" json:"unspsc,omitempty"`
	Taxcoe               string   `protobuf:"bytes,8,opt,name=taxcoe,proto3" json:"taxcoe,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodsInfo) Reset()         { *m = GoodsInfo{} }
func (m *GoodsInfo) String() string { return proto.CompactTextString(m) }
func (*GoodsInfo) ProtoMessage()    {}
func (*GoodsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ff7bf58be59df1, []int{2}
}

func (m *GoodsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodsInfo.Unmarshal(m, b)
}
func (m *GoodsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodsInfo.Marshal(b, m, deterministic)
}
func (m *GoodsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsInfo.Merge(m, src)
}
func (m *GoodsInfo) XXX_Size() int {
	return xxx_messageInfo_GoodsInfo.Size(m)
}
func (m *GoodsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsInfo proto.InternalMessageInfo

func (m *GoodsInfo) GetGood() *Good {
	if m != nil {
		return m.Good
	}
	return nil
}

func (m *GoodsInfo) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *GoodsInfo) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *GoodsInfo) GetDepartment() string {
	if m != nil {
		return m.Department
	}
	return ""
}

func (m *GoodsInfo) GetFirm() string {
	if m != nil {
		return m.Firm
	}
	return ""
}

func (m *GoodsInfo) GetUnspsc() string {
	if m != nil {
		return m.Unspsc
	}
	return ""
}

func (m *GoodsInfo) GetTaxcoe() string {
	if m != nil {
		return m.Taxcoe
	}
	return ""
}

type ListQuery struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort                 string   `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListQuery) Reset()         { *m = ListQuery{} }
func (m *ListQuery) String() string { return proto.CompactTextString(m) }
func (*ListQuery) ProtoMessage()    {}
func (*ListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ff7bf58be59df1, []int{3}
}

func (m *ListQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListQuery.Unmarshal(m, b)
}
func (m *ListQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListQuery.Marshal(b, m, deterministic)
}
func (m *ListQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListQuery.Merge(m, src)
}
func (m *ListQuery) XXX_Size() int {
	return xxx_messageInfo_ListQuery.Size(m)
}
func (m *ListQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ListQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ListQuery proto.InternalMessageInfo

func (m *ListQuery) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListQuery) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListQuery) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

type Request struct {
	ListQuery            *ListQuery `protobuf:"bytes,1,opt,name=list_query,json=listQuery,proto3" json:"list_query,omitempty"`
	Good                 *Good      `protobuf:"bytes,2,opt,name=good,proto3" json:"good,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ff7bf58be59df1, []int{4}
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

func (m *Request) GetListQuery() *ListQuery {
	if m != nil {
		return m.ListQuery
	}
	return nil
}

func (m *Request) GetGood() *Good {
	if m != nil {
		return m.Good
	}
	return nil
}

type Response struct {
	GoodsInfo            *GoodsInfo   `protobuf:"bytes,1,opt,name=goods_info,json=goodsInfo,proto3" json:"goods_info,omitempty"`
	GoodsInfos           []*GoodsInfo `protobuf:"bytes,2,rep,name=goods_infos,json=goodsInfos,proto3" json:"goods_infos,omitempty"`
	Valid                bool         `protobuf:"varint,3,opt,name=valid,proto3" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ff7bf58be59df1, []int{5}
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

func (m *Response) GetGoodsInfo() *GoodsInfo {
	if m != nil {
		return m.GoodsInfo
	}
	return nil
}

func (m *Response) GetGoodsInfos() []*GoodsInfo {
	if m != nil {
		return m.GoodsInfos
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
	proto.RegisterType((*Good)(nil), "goods.Good")
	proto.RegisterType((*Barcode)(nil), "goods.Barcode")
	proto.RegisterType((*GoodsInfo)(nil), "goods.GoodsInfo")
	proto.RegisterType((*ListQuery)(nil), "goods.ListQuery")
	proto.RegisterType((*Request)(nil), "goods.Request")
	proto.RegisterType((*Response)(nil), "goods.Response")
}

func init() { proto.RegisterFile("proto/goods/goods.proto", fileDescriptor_42ff7bf58be59df1) }

var fileDescriptor_42ff7bf58be59df1 = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0x7f, 0x6d, 0xfa, 0x27, 0x3e, 0xdd, 0xba, 0xc9, 0xfa, 0x89, 0x65, 0xe3, 0xcf, 0xaa,
	0x70, 0x53, 0x40, 0xda, 0xc4, 0x78, 0x82, 0x01, 0xd2, 0x14, 0x09, 0x21, 0x61, 0x89, 0x2b, 0x2e,
	0xaa, 0x2c, 0xf6, 0x52, 0x8b, 0x36, 0xce, 0x62, 0x97, 0xb1, 0x5b, 0xc4, 0x1b, 0xf1, 0x06, 0xbc,
	0x0a, 0x2f, 0x82, 0xce, 0xb1, 0x9b, 0x55, 0x43, 0xa0, 0xde, 0x4c, 0xfe, 0x7e, 0xbf, 0xc7, 0x4e,
	0x72, 0x3e, 0xc7, 0x2b, 0x1c, 0xd4, 0x8d, 0x71, 0xe6, 0xb4, 0x34, 0x46, 0x5a, 0xff, 0xf7, 0x84,
	0x1c, 0xde, 0x27, 0x91, 0xfe, 0x88, 0xa0, 0x77, 0x61, 0x8c, 0xe4, 0x63, 0xe8, 0x6a, 0x99, 0x74,
	0x26, 0x9d, 0x29, 0x13, 0x5d, 0x2d, 0x39, 0x87, 0x5e, 0x95, 0x2f, 0x55, 0xd2, 0x25, 0x87, 0xd6,
	0xfc, 0x10, 0x62, 0x55, 0x95, 0x33, 0xf2, 0x23, 0xf2, 0x87, 0xaa, 0x2a, 0xdf, 0x63, 0x34, 0x81,
	0x91, 0x54, 0xb6, 0x68, 0x74, 0xed, 0xb4, 0xa9, 0x92, 0x1e, 0xa5, 0x9b, 0x16, 0x7f, 0x00, 0x03,
	0xeb, 0x72, 0xb7, 0xb2, 0x49, 0x7f, 0xd2, 0x99, 0x46, 0x22, 0x28, 0x3c, 0xf4, 0xb2, 0xc9, 0x2b,
	0x39, 0xd3, 0x32, 0x19, 0x50, 0x32, 0x24, 0x9d, 0x49, 0x7e, 0x0c, 0xa3, 0x22, 0x77, 0xaa, 0x34,
	0xcd, 0x2d, 0xa6, 0x43, 0x4a, 0x61, 0x6d, 0x65, 0x92, 0x3f, 0x85, 0x5d, 0xa9, 0xea, 0xbc, 0x71,
	0x4b, 0x55, 0x39, 0x2c, 0x89, 0xa9, 0x64, 0xe7, 0xce, 0xcc, 0x24, 0x3f, 0x80, 0xe1, 0x95, 0x6e,
	0x96, 0x18, 0x33, 0xff, 0x64, 0x94, 0x99, 0xe4, 0x0f, 0x81, 0xad, 0x2a, 0x5b, 0xdb, 0x02, 0x23,
	0xa0, 0x28, 0xf6, 0x46, 0x26, 0xf9, 0x63, 0x00, 0x97, 0x7f, 0x2d, 0x8c, 0x54, 0x98, 0x8e, 0x28,
	0x65, 0xc1, 0xc9, 0xa8, 0x3d, 0x85, 0xb2, 0x36, 0xd9, 0xa1, 0x80, 0xd6, 0xb8, 0xa5, 0x68, 0x54,
	0xee, 0x94, 0x9c, 0xe5, 0x2e, 0xd9, 0xa5, 0x16, 0xb0, 0xe0, 0x9c, 0x3b, 0x8c, 0x57, 0xb5, 0x5c,
	0xc7, 0x63, 0x1f, 0x07, 0xe7, 0xdc, 0xf1, 0xe7, 0x10, 0xbf, 0xce, 0x1b, 0x3c, 0xde, 0x26, 0x7b,
	0x93, 0x68, 0x3a, 0x3a, 0x1b, 0x9f, 0x78, 0x60, 0xc1, 0x16, 0x6d, 0x9e, 0x7e, 0x8b, 0x60, 0x18,
	0xc4, 0x06, 0xb8, 0x88, 0xc0, 0x1d, 0x42, 0x4c, 0xdb, 0xf0, 0xb5, 0x3d, 0xbc, 0x21, 0xe9, 0x8c,
	0x22, 0xeb, 0x4c, 0xf1, 0x19, 0xa3, 0xc0, 0x8f, 0x74, 0x26, 0xf9, 0xff, 0xd0, 0xaf, 0x1b, 0x5d,
	0x28, 0x22, 0x17, 0x09, 0x2f, 0xd0, 0xbd, 0xd1, 0xd2, 0xcd, 0x03, 0x32, 0x2f, 0x90, 0xe4, 0x5c,
	0xe9, 0x72, 0xee, 0x02, 0xaf, 0xa0, 0xb0, 0x5a, 0xaa, 0xda, 0xcd, 0x03, 0x28, 0x2f, 0xb0, 0x53,
	0xab, 0x4a, 0x3b, 0x42, 0xc3, 0x04, 0xad, 0xd1, 0xb3, 0xb5, 0x2a, 0x88, 0x07, 0x13, 0xb4, 0xc6,
	0x09, 0x2a, 0x1b, 0x63, 0xed, 0x8d, 0x3f, 0xda, 0xf3, 0xd8, 0xb4, 0xf8, 0x23, 0x60, 0x95, 0x72,
	0x21, 0x0f, 0x44, 0x5a, 0x83, 0xef, 0x43, 0xa4, 0x97, 0x25, 0x01, 0x61, 0x02, 0x97, 0x38, 0x3e,
	0x97, 0xbe, 0x49, 0x33, 0x4c, 0x3c, 0x10, 0x08, 0x56, 0xb6, 0x2c, 0xef, 0x01, 0x1b, 0xff, 0x1b,
	0xd8, 0xde, 0x3d, 0x60, 0xe9, 0xcf, 0x0e, 0xb0, 0x0b, 0xea, 0x6c, 0x75, 0x65, 0xf8, 0x31, 0xf4,
	0xb0, 0xcd, 0x04, 0x62, 0x74, 0x36, 0x0a, 0xe8, 0x30, 0x17, 0x14, 0x60, 0x77, 0x68, 0xae, 0x43,
	0xe7, 0xbd, 0xe0, 0x47, 0x10, 0xaf, 0xe7, 0x39, 0x5c, 0x9a, 0x56, 0xf3, 0x27, 0x00, 0x77, 0x83,
	0x4c, 0x08, 0x98, 0xd8, 0x70, 0xb0, 0x8b, 0x38, 0xc9, 0x44, 0x81, 0x09, 0x5a, 0x23, 0x1b, 0x3f,
	0xc2, 0x04, 0x81, 0x89, 0xa0, 0xd0, 0xa7, 0xe1, 0x55, 0x81, 0x43, 0x50, 0x69, 0x06, 0xec, 0x9d,
	0xb6, 0xee, 0xc3, 0x4a, 0x35, 0xb7, 0xf8, 0x8a, 0x0b, 0xbd, 0xd4, 0x2e, 0x4c, 0x93, 0x17, 0xf8,
	0x98, 0x3a, 0x2f, 0xfd, 0x7f, 0x82, 0x48, 0xd0, 0x9a, 0x00, 0x9a, 0xc6, 0x85, 0x6f, 0xa1, 0x75,
	0xfa, 0x09, 0x86, 0x42, 0x5d, 0xaf, 0x94, 0x75, 0xfc, 0x14, 0x60, 0xa1, 0xad, 0x9b, 0x5d, 0xe3,
	0xb1, 0xa1, 0x25, 0xfb, 0xa1, 0x25, 0xed, 0xe3, 0x04, 0x5b, 0xb4, 0x4f, 0x5e, 0x77, 0xaf, 0xfb,
	0x97, 0xee, 0xa5, 0xdf, 0x3b, 0x10, 0x0b, 0x65, 0x6b, 0x53, 0x59, 0x85, 0xc7, 0x87, 0x11, 0xaf,
	0xae, 0xcc, 0xbd, 0xe3, 0x5b, 0x22, 0x82, 0x95, 0x2d, 0x9c, 0x97, 0x30, 0xba, 0xdb, 0x60, 0x93,
	0x2e, 0x5d, 0xaf, 0x3f, 0x77, 0x40, 0xbb, 0xc3, 0x62, 0x2f, 0xbe, 0xe4, 0x8b, 0x70, 0x51, 0x62,
	0xe1, 0xc5, 0xd9, 0xaf, 0x0e, 0xf4, 0xa9, 0x9e, 0x3f, 0x83, 0x1e, 0x7e, 0x09, 0x5f, 0x5f, 0xd2,
	0xf0, 0xe9, 0x47, 0x7b, 0xad, 0xf6, 0x2f, 0x9b, 0xfe, 0xc7, 0xa7, 0x10, 0x5d, 0xa8, 0xad, 0x2a,
	0x5f, 0xc0, 0xe0, 0x0d, 0x8d, 0xdf, 0x96, 0xc5, 0x1f, 0x69, 0x18, 0xb7, 0x2c, 0x7e, 0xab, 0x16,
	0x6a, 0xab, 0xe2, 0xcb, 0x01, 0xfd, 0x44, 0xbc, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x6c, 0xf0,
	0x18, 0x5a, 0x3d, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Goods service

type GoodsClient interface {
	// 获取商品列表
	List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 根据 唯一 获取商品信息
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 创建商品
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 更新商品
	Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 删除商品
	Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type goodsClient struct {
	c           client.Client
	serviceName string
}

func NewGoodsClient(serviceName string, c client.Client) GoodsClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "goods"
	}
	return &goodsClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *goodsClient) List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Goods.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Goods.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Goods.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Goods.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Goods.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Goods service

type GoodsHandler interface {
	// 获取商品列表
	List(context.Context, *Request, *Response) error
	// 根据 唯一 获取商品信息
	Get(context.Context, *Request, *Response) error
	// 创建商品
	Create(context.Context, *Request, *Response) error
	// 更新商品
	Update(context.Context, *Request, *Response) error
	// 删除商品
	Delete(context.Context, *Request, *Response) error
}

func RegisterGoodsHandler(s server.Server, hdlr GoodsHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Goods{hdlr}, opts...))
}

type Goods struct {
	GoodsHandler
}

func (h *Goods) List(ctx context.Context, in *Request, out *Response) error {
	return h.GoodsHandler.List(ctx, in, out)
}

func (h *Goods) Get(ctx context.Context, in *Request, out *Response) error {
	return h.GoodsHandler.Get(ctx, in, out)
}

func (h *Goods) Create(ctx context.Context, in *Request, out *Response) error {
	return h.GoodsHandler.Create(ctx, in, out)
}

func (h *Goods) Update(ctx context.Context, in *Request, out *Response) error {
	return h.GoodsHandler.Update(ctx, in, out)
}

func (h *Goods) Delete(ctx context.Context, in *Request, out *Response) error {
	return h.GoodsHandler.Delete(ctx, in, out)
}
