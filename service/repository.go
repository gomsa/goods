package service

import (
	brandPB "github.com/gomsa/goods/proto/brand"
	goodsPB "github.com/gomsa/goods/proto/goods"
)

// GoodRepository  商品仓库接口
type GoodRepository interface {
	GetBarcode(*goodsPB.Barcode) (*goodsPB.Barcode, error)
	IsBarcode(*goodsPB.Good) (bool, error)
	DeleteBarcode(*goodsPB.Barcode) (bool, error)
	DeleteBarcodeByGoodID(*goodsPB.Barcode) (bool, error)
	IsGood(*goodsPB.Good) (bool, error)
	Create(*goodsPB.Good) (*goodsPB.Good, error)
	Get(*goodsPB.Good) (*goodsPB.Good, error)
	List(*goodsPB.Request) ([]*goodsPB.Good, error)
	Update(*goodsPB.Good) (bool, error)
	Delete(*goodsPB.Good) (bool, error)
}

// BrandRepository  品牌仓库接口
type BrandRepository interface {
	Exist(*brandPB.Brand) (bool, error)
	All(*brandPB.Brand) ([]*brandPB.Brand, error)
	List(*brandPB.Request) ([]*brandPB.Brand, error)
	Create(*brandPB.Brand) (*brandPB.Brand, error)
	Update(*brandPB.Brand) (bool, error)
	Delete(*brandPB.Brand) (bool, error)
	Get(*brandPB.Brand) (*brandPB.Brand, error)
}
