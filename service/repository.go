package service

import (
	brandPB "github.com/gomsa/goods/proto/brand"
	categoryPB "github.com/gomsa/goods/proto/category"
	departmentPB "github.com/gomsa/goods/proto/department"
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
