package main

import (
	"context"
	"fmt"
	"testing"

	brandPB "github.com/gomsa/goods/proto/brand"
	goodsPB "github.com/gomsa/goods/proto/goods"

	"github.com/gomsa/goods/hander"
	db "github.com/gomsa/goods/providers/database"
	"github.com/gomsa/goods/service"
)

func TestAddGoods(t *testing.T) {
	req := &goodsPB.Request{
		Good: &goodsPB.Good{
			Name:        `测试商品2`,
			EngName:     `goods`,
			Description: `描述1`,
			Cess:        0,
			Barcodes: []*goodsPB.Barcode{
				{
					Id:      `6152340210121`,
					StockId: `asdas1s23123`,
					Price:   11000,
				},
				{
					Id:      `61523402120231`,
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
	// fmt.Println(err, res, req)
	t.Log(t, err)
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

func TestIsBarcodeGoods(t *testing.T) {
	repo := &service.Goods{db.DB}
	h := &hander.Goods{repo}

	req := &goodsPB.Request{
		Good: &goodsPB.Good{
			Barcodes: []*goodsPB.Barcode{
				{
					Id: `615234021202`,
				},
				{
					Id: `61523402101`,
				},
			},
		},
	}
	res := &goodsPB.Response{}
	err := h.IsBarcode(context.TODO(), req, res)
	// fmt.Println(err, res)
	t.Log(t, err)
}
func TestGoodsByBarcodeGoods(t *testing.T) {
	repo := &service.Goods{db.DB}
	h := &hander.Goods{repo}

	req := &goodsPB.Request{
		Barcode: &goodsPB.Barcode{
			Id: `61523402121`,
		},
	}
	res := &goodsPB.Response{}
	err := h.GoodsByBarcode(context.TODO(), req, res)
	fmt.Println(err, res)
	t.Log(t, err)
}

func TestUpdateGoods(t *testing.T) {
	req := &goodsPB.Request{
		Good: &goodsPB.Good{
			Id:          `9cf86abf-ad72-45e2-ab84-248eb70daf59`,
			Name:        `测试商品3`,
			EngName:     `goods`,
			Description: `描述1`,
			Cess:        0,
			Barcodes: []*goodsPB.Barcode{
				{
					Id:      `61523402104`,
					StockId: `asdas1s23123`,
					Price:   14000,
				},
				{
					Id:      `615234021202`,
					StockId: `asdas1s23123`,
					Price:   12000,
				},
			},
		},
	}

	res := &goodsPB.Response{}
	repo := &service.Goods{db.DB}
	h := &hander.Goods{repo}
	err := h.Update(context.TODO(), req, res)
	// fmt.Println(err, res)

	// 删除商品条码 需要前端单独调用
	reqb := &goodsPB.Request{
		Barcode: &goodsPB.Barcode{
			Id: `61523402101`,
		},
	}
	resb := &goodsPB.Response{}
	err = h.DeleteBarcode(context.TODO(), reqb, resb)
	// fmt.Println(err, resb)
	t.Log(t, err)

}

func TestDeleteGoods(t *testing.T) {
	req := &goodsPB.Request{
		Good: &goodsPB.Good{
			Id: `8668dbe3-4dc0-4ded-96f3-9e0f9522280a`,
		},
	}

	res := &goodsPB.Response{}
	repo := &service.Goods{db.DB}
	h := &hander.Goods{repo}
	err := h.Delete(context.TODO(), req, res)
	// fmt.Println(err, res)
	t.Log(t, err)
}

func TestAddBrand(t *testing.T) {
	req := &brandPB.Request{
		Brand: &brandPB.Brand{
			Name: `伊利`,
		},
	}

	res := &brandPB.Response{}
	repo := &service.Brand{db.DB}
	h := &hander.Brand{repo}
	err := h.Create(context.TODO(), req, res)
	// fmt.Println(err, res)
	t.Log(t, err)
}

func TestGetBrand(t *testing.T) {
	req := &brandPB.Request{
		Brand: &brandPB.Brand{
			Id: 2,
		},
	}

	res := &brandPB.Response{}
	repo := &service.Brand{db.DB}
	h := &hander.Brand{repo}
	err := h.Get(context.TODO(), req, res)
	// fmt.Println(err, res)
	t.Log(t, err)
}

func TestUpdateBrand(t *testing.T) {
	req := &brandPB.Request{
		Brand: &brandPB.Brand{
			Id:   2,
			Name: `伊丽莎白`,
		},
	}
	res := &brandPB.Response{}
	repo := &service.Brand{db.DB}
	h := &hander.Brand{repo}
	err := h.Update(context.TODO(), req, res)
	// fmt.Println(err, res)
	t.Log(t, err)
}

func TestDeleteBrand(t *testing.T) {
	req := &brandPB.Request{
		Brand: &brandPB.Brand{
			Id: 2,
		},
	}
	res := &brandPB.Response{}
	repo := &service.Brand{db.DB}
	h := &hander.Brand{repo}
	err := h.Delete(context.TODO(), req, res)
	// fmt.Println(err, res)
	t.Log(t, err)
}
