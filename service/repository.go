package service

import (
	brandPB "github.com/gomsa/goods/proto/brand"
	categoryPB "github.com/gomsa/goods/proto/category"
	departmentPB "github.com/gomsa/goods/proto/department"
	firmPB "github.com/gomsa/goods/proto/firm"
	goodsPB "github.com/gomsa/goods/proto/goods"
	taxcodePB "github.com/gomsa/goods/proto/taxcode"
	unspscPB "github.com/gomsa/goods/proto/unspsc"
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
	List(*goodsPB.ListQuery, *goodsPB.Good) ([]*goodsPB.Good, error)
	Update(*goodsPB.Good) (bool, error)
	Delete(*goodsPB.Good) (bool, error)
}

// BrandRepository  品牌仓库接口
type BrandRepository interface {
	Exist(*brandPB.Brand) (bool, error)
	All(*brandPB.Brand) ([]*brandPB.Brand, error)
	List(*brandPB.ListQuery, *brandPB.Brand) ([]*brandPB.Brand, error)
	Create(*brandPB.Brand) (*brandPB.Brand, error)
	Update(*brandPB.Brand) (bool, error)
	Delete(*brandPB.Brand) (bool, error)
	Get(*brandPB.Brand) (*brandPB.Brand, error)
}

// CategoryRepository  分类接口
type CategoryRepository interface {
	All(*categoryPB.Category) ([]*categoryPB.Category, error)
	List(*categoryPB.Category) ([]*categoryPB.Category, error)
	Create(*categoryPB.Category) (*categoryPB.Category, error)
	Update(*categoryPB.Category) (bool, error)
	Delete(*categoryPB.Category) (bool, error)
	Get(*categoryPB.Category) (*categoryPB.Category, error)
}

// DepartmentRepository  部门接口
type DepartmentRepository interface {
	All(*departmentPB.Department) ([]*departmentPB.Department, error)
	List(*departmentPB.Department) ([]*departmentPB.Department, error)
	Create(*departmentPB.Department) (*departmentPB.Department, error)
	Update(*departmentPB.Department) (bool, error)
	Delete(*departmentPB.Department) (bool, error)
	Get(*departmentPB.Department) (*departmentPB.Department, error)
}

// UnspscRepository  国际商品及服务编码接口
type UnspscRepository interface {
	All(*unspscPB.Unspsc) ([]*unspscPB.Unspsc, error)
	List(*unspscPB.Unspsc) ([]*unspscPB.Unspsc, error)
	Create(*unspscPB.Unspsc) (*unspscPB.Unspsc, error)
	Update(*unspscPB.Unspsc) (bool, error)
	Delete(*unspscPB.Unspsc) (bool, error)
	Get(*unspscPB.Unspsc) (*unspscPB.Unspsc, error)
}

// TaxcodeRepository  国家税收分类编码接口
type TaxcodeRepository interface {
	All(*taxcodePB.Taxcode) ([]*taxcodePB.Taxcode, error)
	List(*taxcodePB.Taxcode) ([]*taxcodePB.Taxcode, error)
	Create(*taxcodePB.Taxcode) (*taxcodePB.Taxcode, error)
	Update(*taxcodePB.Taxcode) (bool, error)
	Delete(*taxcodePB.Taxcode) (bool, error)
	Get(*taxcodePB.Taxcode) (*taxcodePB.Taxcode, error)
}

// FirmRepository  公司接口
type FirmRepository interface {
	Exist(*firmPB.Firm) (bool, error)
	All(*firmPB.Firm) ([]*firmPB.Firm, error)
	List(*firmPB.ListQuery, *firmPB.Firm) ([]*firmPB.Firm, error)
	Create(*firmPB.Firm) (*firmPB.Firm, error)
	Update(*firmPB.Firm) (bool, error)
	Delete(*firmPB.Firm) (bool, error)
	Get(*firmPB.Firm) (*firmPB.Firm, error)
}
