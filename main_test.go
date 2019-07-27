package main

import (
	"context"
	"fmt"
	"testing"

	goodsPB "github.com/gomsa/goods/proto/goods"

	"github.com/gomsa/goods/hander"
	db "github.com/gomsa/goods/providers/database"
	"github.com/gomsa/goods/service"
)

func TestAddGoods(t *testing.T) {
	// req := &goodsPB.Request{
	// 	Good: &goodsPB.Good{
	// 		Name:        `测试商品1`,
	// 		EngName:     `goods`,
	// 		Description: `描述1`,
	// 		Cess:        0,
	// 		Barcodes: []*goodsPB.Barcode{
	// 			{
	// 				Id:      61523402111,
	// 				StockId: `asdas1s23123`,
	// 				Price:   11000,
	// 			},
	// 			{
	// 				Id:      61523402121,
	// 				StockId: `asdas1s23123`,
	// 				Price:   11000,
	// 			},
	// 		},
	// 	},
	// }

	// res := &goodsPB.Response{}
	// repo := &service.Goods{db.DB}
	// h := &hander.Goods{repo}
	// err := h.Create(context.TODO(), req, res)
	// // fmt.Println(err, res, req)
	// t.Log(t, err)
}

func TestGetGoods(t *testing.T) {
	repo := &service.Goods{db.DB}
	h := &hander.Goods{repo}

	req := &goodsPB.Request{
		Good: &goodsPB.Good{
			Id: `a0da2378-1ba9-4285-9c74-b6eee9ede39d`,
		},
	}
	res := &goodsPB.Response{}
	err := h.Get(context.TODO(), req, res)
	fmt.Println(err, res)
	t.Log(t, err)
}
