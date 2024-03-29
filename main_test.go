package main

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/gomsa/goods/hander"
	brandPB "github.com/gomsa/goods/proto/brand"
	categoryPB "github.com/gomsa/goods/proto/category"
	goodsPB "github.com/gomsa/goods/proto/goods"
	unspscPB "github.com/gomsa/goods/proto/unspsc"

	db "github.com/gomsa/goods/providers/database"
	"github.com/gomsa/goods/service"
)

func TestAddGoods(t *testing.T) {
	images, _ := json.Marshal([]string{
		"http://www.xxx.com/userfile/uploada/gra/1608153471/06917878036526/06917878036526.7.jpg",
		"http://www.xxx.com/userfile/uploada/gra/1608153471/06917878036526/06917878036526.8.jpg",
	})
	req := &goodsPB.Request{
		Good: &goodsPB.Good{
			Code:        `10084`,
			Name:        `测试商品2`,
			EnName:      `goods`,
			Description: `描述1`,
			Cess:        0,
			Barcodes: []*goodsPB.Barcode{
				{
					Id:      `61523402101211`,
					StockId: `asdas1s23123`,
					Price:   11000,
				},
				{
					Id:      `615234021202312`,
					StockId: `asdas1s23123`,
					Price:   11000,
					Images:  string(images),
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
		// ListQuery: &goodsPB.ListQuery{},
		Good: &goodsPB.Good{},
	}
	res := &goodsPB.Response{}
	err := h.List(context.TODO(), req, res)
	// fmt.Println(err, res)
	t.Log(t, err)
}

func TestExistGoods(t *testing.T) {
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
	err := h.Exist(context.TODO(), req, res)
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
	// fmt.Println(err, res)
	t.Log(t, err)
}

func TestUpdateGoods(t *testing.T) {
	req := &goodsPB.Request{
		Good: &goodsPB.Good{
			Id:          `9cf86abf-ad72-45e2-ab84-248eb70daf59`,
			Name:        `测试商品3`,
			EnName:      `goods`,
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

func TestAddCategory(t *testing.T) {
	req := &categoryPB.Request{
		Category: &categoryPB.Category{
			Name:   `洗化-1`,
			Parent: 2,
		},
	}

	res := &categoryPB.Response{}
	h := &hander.Category{&service.Category{db.DB}}
	err := h.Create(context.TODO(), req, res)
	// fmt.Println(err, res)
	t.Log(t, err)
}
func TestAllCategory(t *testing.T) {
	req := &categoryPB.Request{
		Category: &categoryPB.Category{},
	}

	res := &categoryPB.Response{}
	h := &hander.Category{&service.Category{db.DB}}
	err := h.All(context.TODO(), req, res)
	// fmt.Println(err, res)
	t.Log(t, err)
}

func TestListCategory(t *testing.T) {
	req := &categoryPB.Request{
		Category: &categoryPB.Category{
			Parent: 2,
		},
	}

	res := &categoryPB.Response{}
	h := &hander.Category{&service.Category{db.DB}}
	err := h.List(context.TODO(), req, res)
	// fmt.Println(err, res)
	t.Log(t, err)
}

func TestGetCategory(t *testing.T) {
	req := &categoryPB.Request{
		Category: &categoryPB.Category{
			Id: 2,
		},
	}

	res := &categoryPB.Response{}
	h := &hander.Category{&service.Category{db.DB}}
	err := h.Get(context.TODO(), req, res)
	// fmt.Println(err, res)
	t.Log(t, err)
}

func TestUpdateCategory(t *testing.T) {
	req := &categoryPB.Request{
		Category: &categoryPB.Category{
			Id:   1,
			Name: `牙膏`,
		},
	}

	res := &categoryPB.Response{}
	h := &hander.Category{&service.Category{db.DB}}
	// err := h.Update(context.TODO(), req, res)
	// fmt.Println(err, res)
	t.Log(t, h, res, req)
}

func TestDeleteCategory(t *testing.T) {
	req := &categoryPB.Request{
		Category: &categoryPB.Category{
			Id: 3,
		},
	}

	res := &categoryPB.Response{}
	h := &hander.Category{&service.Category{db.DB}}
	err := h.Delete(context.TODO(), req, res)
	// fmt.Println(err, res)
	t.Log(t, err)
}

// func TestBarcode(t *testing.T) {
// 	req := &barcodePB.Request{
// 		Goods: &barcodePB.Goods{
// 			Barcode: `6917878036526`,
// 		},
// 	}

// 	res := &barcodePB.Response{}
// 	h := &hander.Barcode{}
// 	err := h.Get(context.TODO(), req, res)
// 	// fmt.Println(res.Goods, err)
// 	t.Log(t, err)
// }

func TestUnspscCheckCreate(t *testing.T) {
	req := &unspscPB.Request{
		Unspsc: &unspscPB.Unspsc{
			Id:   52152010,
			Name: `家用电器和日用电子产品>>家用厨具>>家用器皿、盛器和存储容器>>家用真空瓶`,
		},
	}

	res := &unspscPB.Response{}
	h := &hander.Unspsc{&service.Unspsc{db.DB}}
	err := h.CheckCreate(context.TODO(), req, res)
	// fmt.Println(res, err)
	t.Log(t, err)
}

func TestUnspscExist(t *testing.T) {
	req := &unspscPB.Request{
		Unspsc: &unspscPB.Unspsc{
			Id: 52152010,
		},
	}

	res := &unspscPB.Response{}
	h := &hander.Unspsc{&service.Unspsc{db.DB}}
	err := h.Exist(context.TODO(), req, res)
	fmt.Println(res, err)
	t.Log(t, err)
}
