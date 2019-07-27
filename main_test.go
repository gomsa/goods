package main

import (
	"fmt"
	"testing"

	goodsPB "github.com/gomsa/goods/proto/goods"
)

func TestAddGoods(t *testing.T) {
	goods := goodsPB.Good{
		Name:        `测试商品`,
		EngName:     `goods`,
		Description: `描述`,
		Cess:        0,
		Barcodes: []*goodsPB.Barcode{
			{
				Id:      6152340211,
				StockId: `asdas1s23123`,
				Price:   1000,
			},
			{
				Id:      6152340212,
				StockId: `asdas1s23123`,
				Price:   1000,
			},
		},
	}
	fmt.Println(goods)
	t.Log(t)
}
