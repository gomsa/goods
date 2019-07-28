package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"

	pb "github.com/gomsa/goods/proto/goods"
	"github.com/gomsa/tools/uitl"
)

// Goods 商品仓库失效
type Goods struct {
	DB *gorm.DB
}

// IsBarcode 查询条码是否存在
func (repo *Goods) IsBarcode(good *pb.Good) (bool, error) {
	var count int
	for _, barcode := range good.Barcodes {
		if err := repo.DB.Model(barcode).Where("id = ?", barcode.Id).Count(&count).Error; err != nil {
			return false, err
		}
		if count > 0 {
			return true, fmt.Errorf("%s 条形码商品已存在", barcode.Id)
		}
	}
	return false, nil
}

// List 获取所有商品信息
func (repo *Goods) List(req *pb.Request) (goods []*pb.Good, err error) {
	db := repo.DB.Model(&req.Good)
	// 计算分页
	limit, offset := uitl.Page(req.ListQuery.Limit, req.ListQuery.Page)
	// 排序
	var sort string
	if req.ListQuery.Sort != "" {
		sort = req.ListQuery.Sort
	} else {
		sort = "created_at desc"
	}
	// 查询条件
	if req.Good.Name != "" {
		db = db.Where("name like ?", "%"+req.Good.Name+"%")
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&goods).Error; err != nil {
		log.Log(err)
		return goods, err
	}
	// 查询关联
	for _, good := range goods {
		repo.Related(good)
	}
	return goods, err
}

// Get 获取商品信息
func (repo *Goods) Get(good *pb.Good) (*pb.Good, error) {
	if err := repo.DB.Model(&good).Find(&good).Error; err != nil {
		return good, err
	}
	// 查询关联
	repo.Related(good)
	return good, nil
}

// Create 创建商品
func (repo *Goods) Create(good *pb.Good) (*pb.Good, error) {
	if exist, message := repo.IsBarcode(good); exist {
		return good, message
	}
	err := repo.DB.Create(good).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return good, fmt.Errorf("添加商品失败")
	}
	return good, nil
}

// Update 更新商品
func (repo *Goods) Update(goods *pb.Good) (bool, error) {
	return true, nil
}

// Delete 删除商品
func (repo *Goods) Delete(goods *pb.Good) (bool, error) {
	return true, nil
}

// Related 关联处理
func (repo *Goods) Related(good *pb.Good) (*pb.Good, error) {
	if err := repo.DB.Model(&good).Related(&good.Barcodes).Error; err != nil {
		if err.Error() != "record not found" {
			return good, err
		}
	}
	Brand := &pb.Brand{}
	if err := repo.DB.Model(&good).Related(Brand).Error; err != nil {
		if err.Error() != "record not found" {
			return good, err
		}
	}
	Category := &pb.Category{}
	if err := repo.DB.Model(&good).Related(Category).Error; err != nil {
		if err.Error() != "record not found" {
			return good, err
		}
	}

	Firm := &pb.Firm{}
	if err := repo.DB.Model(&good).Related(Firm).Error; err != nil {
		if err.Error() != "record not found" {
			return good, err
		}
	}

	Unspsc := &pb.Unspsc{}
	if err := repo.DB.Model(&good).Related(Unspsc).Error; err != nil {
		if err.Error() != "record not found" {
			return good, err
		}
	}

	Taxcode := &pb.Taxcode{}
	if err := repo.DB.Model(&good).Related(Taxcode).Error; err != nil {
		if err.Error() != "record not found" {
			return good, err
		}
	}

	good.Brand = Brand
	good.Category = Category
	good.Firm = Firm
	good.Unspsc = Unspsc
	good.Taxcode = Taxcode
	return good, nil
}
