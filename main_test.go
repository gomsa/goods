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
	req := &goodsPB.Request{
		Good: &goodsPB.Good{
			Name:        `测试商品1`,
			EngName:     `goods`,
			Description: `描述1`,
			Cess:        0,
			Barcodes: []*goodsPB.Barcode{
				{
					Id:      `61523402101`,
					StockId: `asdas1s23123`,
					Price:   11000,
				},
				{
					Id:      `615234021202`,
					StockId: `asdas1s23123`,
					Price:   11000,
				},
			},
		},
	}

	res := &goodsPB.Response{}
	repo := &service.Goods{db.DB}
	h := &hander.Goods{repo}
	err := h.Create(context.TODO(), req, res)
	fmt.Println(err, res, req)
	t.Log(t)
}

func TestGetGoods(t *testing.T) {
	repo := &service.Goods{db.DB}
	h := &hander.Goods{repo}

	req := &goodsPB.Request{
		Good: &goodsPB.Good{
			// Id: `9cf86abf-ad72-45e2-ab84-248eb70daf59`,
			Barcodes: []*goodsPB.Barcode{
				{
					Id: `61523402101`,
				},
			},
		},
	}
	res := &goodsPB.Response{}
	err := h.Get(context.TODO(), req, res)
	// fmt.Println(err, res)
	t.Log(t, err)
}

func TestListGoods(t *testing.T) {
	repo := &service.Goods{db.DB}
	h := &hander.Goods{repo}

	req := &goodsPB.Request{
		ListQuery: &goodsPB.ListQuery{},
		Good: &goodsPB.Good{
			Name: `测试商品1`,
		},
	}
	res := &goodsPB.Response{}
	err := h.List(context.TODO(), req, res)
	// fmt.Println(err, res)
	t.Log(t, err)
}
