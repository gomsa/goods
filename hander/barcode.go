package hander

import (
	"context"
	"fmt"

	"github.com/bigrocs/barcode/drives"

	pb "github.com/gomsa/goods/proto/barcode"
)

// Barcode 条码信息结构
type Barcode struct {
}

// Get 获取条码信息
func (srv *Barcode) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Goods == nil {
		return fmt.Errorf("请求参数错误")
	}
	chinatrace := &drives.Chinatrace{
		BaseHost: "http://webapi.chinatrace.org",
		Key:      "V7N3Xpm4jpRon/WsZ8X/63G8oMeGdUkA8Luxs1CenTY=",
	}
	goods, err := chinatrace.Get(req.Goods.Barcode)
	if err != nil {
		return err
	}
	res.Goods = &pb.Goods{}
	res.Goods.Barcode = req.Goods.Barcode
	res.Goods.Name = goods.Name
	res.Goods.EnName = goods.EnName
	res.Goods.Image = goods.Image
	res.Goods.BrandName = goods.BrandName
	res.Goods.Specification = goods.Specification
	res.Goods.Unit = goods.Unit
	res.Goods.Width = goods.Width
	res.Goods.Height = goods.Height
	res.Goods.Depth = goods.Depth
	res.Goods.NetWeight = goods.NetWeight
	res.Goods.GrossWeight = goods.GrossWeight
	res.Goods.Unspsc = goods.Unspsc
	res.Goods.UnspscName = goods.UnspscName
	res.Goods.Source = goods.Source
	res.Goods.Place = goods.Place
	res.Goods.Country = goods.Country
	res.Goods.FirmName = goods.FirmName
	res.Goods.FirmAddress = goods.FirmAddress
	res.Goods.FirmStatus = goods.FirmStatus
	return err
}
