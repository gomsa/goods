package service

import (
	"fmt"

	pb "github.com/gomsa/goods/proto/goods"
	"github.com/micro/go-micro/util/log"

	"github.com/jinzhu/gorm"
)

// Goods 商品仓库失效
type Goods struct {
	DB *gorm.DB
}


// IsBarcode 查询条码是否存在
func (repo *Goods) IsBarcode(good *pb.Good) (bool, error) {
	var count int
	for _, barcode := range good.Barcodes {
		if err := repo.DB.Model(&pb.Barcode{}).Where("id = ?", barcode.Id).Count(&count).Error; err != nil {
			return false, err
		}
		if count > 0 {
			return true
		}
	}
	return false, nil
}

// List 获取所有商品信息
func (repo *Goods) List(req *pb.ListQuery) (goods []*pb.Good, err error) {
	return nil, nil
}

// Get 获取商品信息
func (repo *Goods) Get(good *pb.Good) (*pb.Good, error) {
	if err := repo.DB.Model(&good).Find(&good).Related(&good.Barcodes).Error; err != nil {
		return good, err
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

// Create 创建商品
func (repo *Goods) Create(good *pb.Good) (*pb.Good, error) {
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
