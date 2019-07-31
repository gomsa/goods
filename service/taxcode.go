package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"

	pb "github.com/gomsa/goods/proto/taxcode"
)

// Taxcode 国家税收分类编码仓库失效
type Taxcode struct {
	DB *gorm.DB
}

// All 获取所有国家税收分类编码信息
func (repo *Taxcode) All(req *pb.Taxcode) (taxcodes []*pb.Taxcode, err error) {
	if err := repo.DB.Find(&taxcodes).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return taxcodes, nil
}

// List 获取所有国家税收分类编码信息
func (repo *Taxcode) List(req *pb.Taxcode) (taxcodes []*pb.Taxcode, err error) {
	if err := repo.DB.Where("parent = ?", req.Parent).Find(&taxcodes).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return taxcodes, nil
}

// Get 获取国家税收分类编码信息
func (repo *Taxcode) Get(taxcode *pb.Taxcode) (*pb.Taxcode, error) {
	if err := repo.DB.Model(&taxcode).Find(&taxcode).Error; err != nil {
		return taxcode, err
	}
	return taxcode, nil
}

// Create 创建国家税收分类编码
func (repo *Taxcode) Create(taxcode *pb.Taxcode) (*pb.Taxcode, error) {
	err := repo.DB.Create(taxcode).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return taxcode, fmt.Errorf("添加国家税收分类编码失败")
	}
	return taxcode, nil
}

// Update 更新国家税收分类编码
func (repo *Taxcode) Update(taxcode *pb.Taxcode) (bool, error) {
	if taxcode.Id == 0 {
		return false, fmt.Errorf("请传入操作id")
	}
	id := &pb.Taxcode{
		Id: taxcode.Id,
	}
	err := repo.DB.Model(id).Updates(taxcode).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除国家税收分类编码
func (repo *Taxcode) Delete(taxcode *pb.Taxcode) (bool, error) {
	if taxcode.Id == 0 {
		return false, fmt.Errorf("请传入操作id")
	}
	id := &pb.Taxcode{
		Id: taxcode.Id,
	}
	if err := repo.DB.Delete(id).Error; err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
