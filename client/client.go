package client

import (
	"os"

	"github.com/gomsa/tools/k8s/client"

	goodsPB "github.com/gomsa/goods/proto/goods"
)

var (
	// Goods 权限客户端
	Goods goodsPB.GoodsClient
)

func init() {
	SrvName := os.Getenv("GOODS_NAME")
	Goods = goodsPB.NewGoodsClient(SrvName, client.DefaultClient)
}
