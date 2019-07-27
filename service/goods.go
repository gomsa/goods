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

// List 获取所有商品信息
func (repo *Goods) List(req *pb.ListQuery) (departments []*pb.Good, err error) {
	return nil, nil
}

// Get 获取商品信息
func (repo *Goods) Get(good *pb.Good) (*pb.Good, error) {
	if err := repo.DB.Where(&good).Find(&good).Error; err != nil {
		return good, err
	}
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
