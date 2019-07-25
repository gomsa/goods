package client

import (
	"os"

	"github.com/gomsa/tools/k8s/client"

	brandPB "github.com/gomsa/goods/proto/brand"
	categoryPB "github.com/gomsa/goods/proto/category"
	firmPB "github.com/gomsa/goods/proto/firm"
	goodsPB "github.com/gomsa/goods/proto/goods"
	taxcodePB "github.com/gomsa/goods/proto/taxcode"
	unspscPB "github.com/gomsa/goods/proto/unspsc"
)

var (
	// Category 客户端
	Category categoryPB.CategorysClient
	// Brand 客户端
	Brand brandPB.BrandsClient
	// Firm 客户端
	Firm firmPB.FirmsClient
	// Goods 客户端
	Goods goodsPB.GoodsClient
	// Taxcode 客户端
	Taxcode taxcodePB.TaxcodesClient
	// Unspsc 客户端
	Unspsc unspscPB.UnspscsClient
)

func init() {
	SrvName := os.Getenv("GOODS_NAME")
	Category = categoryPB.NewCategorysClient(SrvName, client.DefaultClient)
	Brand = brandPB.NewBrandsClient(SrvName, client.DefaultClient)
	Firm = firmPB.NewFirmsClient(SrvName, client.DefaultClient)
	Goods = goodsPB.NewGoodsClient(SrvName, client.DefaultClient)
	Taxcode = taxcodePB.NewTaxcodesClient(SrvName, client.DefaultClient)
	Unspsc = unspscPB.NewUnspscsClient(SrvName, client.DefaultClient)
}
