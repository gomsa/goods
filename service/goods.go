package service

import (
	pb "github.com/gomsa/goods/proto/goods"

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
func (repo *Goods) Get(goods *pb.Good) (*pb.Good, error) {
	return goods, nil
}

// Create 创建商品
func (repo *Goods) Create(goods *pb.Good) (*pb.Good, error) {
	return goods, nil
}

// Update 更新商品
func (repo *Goods) Update(goods *pb.Good) (bool, error) {
	return true, nil
}

// Delete 删除商品
func (repo *Goods) Delete(goods *pb.Good) (bool, error) {
	return true, nil
}
