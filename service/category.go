package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"

	categoryPB "github.com/gomsa/goods/proto/category"
	pb "github.com/gomsa/goods/proto/category"
)

// Category 分类仓库失效
type Category struct {
	DB *gorm.DB
}

// All 获取所有分类信息
func (repo *Category) All(req *pb.Category) (categorys []*pb.Category, err error) {
	if err := repo.DB.Find(&categorys).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return categorys, nil
}

// List 获取所有分类信息
func (repo *Category) List(req *categoryPB.Category) (categorys []*pb.Category, err error) {
	db := repo.DB
	// 查询条件
	if req.Id != 0 {
		db = db.Where("id = ?", req.Id)
	}
	if err := repo.DB.Where("parent = ?", req.Parent).Find(&categorys).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return categorys, nil
}

// Get 获取分类信息
func (repo *Category) Get(category *pb.Category) (*pb.Category, error) {
	if err := repo.DB.Model(&category).Find(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

// Create 创建分类
func (repo *Category) Create(category *pb.Category) (*pb.Category, error) {
	err := repo.DB.Create(category).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return category, fmt.Errorf("添加分类失败")
	}
	return category, nil
}

// Update 更新分类
func (repo *Category) Update(category *pb.Category) (bool, error) {
	if category.Id == 0 {
		return false, fmt.Errorf("请传入操作id")
	}
	id := &pb.Category{
		Id: category.Id,
	}
	err := repo.DB.Model(id).Updates(category).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除分类
func (repo *Category) Delete(category *pb.Category) (bool, error) {
	if category.Id == 0 {
		return false, fmt.Errorf("请传入操作id")
	}
	id := &pb.Category{
		Id: category.Id,
	}
	if err := repo.DB.Delete(id).Error; err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
