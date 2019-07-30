package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"

	pb "github.com/gomsa/goods/proto/brand"
	"github.com/gomsa/tools/uitl"
)

// Brand 品牌仓库失效
type Brand struct {
	DB *gorm.DB
}

// Exist 查询品牌是否存在
func (repo *Brand) Exist(brand *pb.Brand) (bool, error) {
	var count int
	if brand.Id != 0 {
		repo.DB.Model(&brand).Where("id = ?", brand.Id).Count(&count)
		if count > 0 {
			return true, fmt.Errorf("%s 品牌已存在", string(brand.Id))
		}
	}
	if brand.Name != "" {
		repo.DB.Model(&brand).Where("name = ?", brand.Name).Count(&count)
		if count > 0 {
			return true, fmt.Errorf("%s 品牌已存在", brand.Name)
		}
	}
	return false, nil
}

// All 获取所有品牌信息
func (repo *Brand) All(req *pb.Brand) (brands []*pb.Brand, err error) {
	if err := repo.DB.Find(&brands).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return brands, nil
}

// List 获取所有品牌信息
func (repo *Brand) List(req *pb.Request) (brands []*pb.Brand, err error) {
	db := repo.DB.Model(&req.Brand)
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
	if req.Brand.Name != "" {
		db = db.Where("name like ?", "%"+req.Brand.Name+"%")
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&brands).Error; err != nil {
		log.Log(err)
		return brands, err
	}
	return brands, err
}

// Get 获取品牌信息
func (repo *Brand) Get(brand *pb.Brand) (*pb.Brand, error) {
	if err := repo.DB.Model(&brand).Find(&brand).Error; err != nil {
		return brand, err
	}
	return brand, nil
}

// Create 创建品牌
func (repo *Brand) Create(brand *pb.Brand) (*pb.Brand, error) {
	if valid, err := repo.Exist(brand); valid {
		return brand, err
	}
	err := repo.DB.Create(brand).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return brand, fmt.Errorf("添加品牌失败")
	}
	return brand, nil
}

// Update 更新品牌
func (repo *Brand) Update(brand *pb.Brand) (bool, error) {
	if brand.Id == 0 {
		return false, fmt.Errorf("请传入操作id")
	}
	if valid, _ := repo.Exist(brand); valid {
		id := &pb.Brand{
			Id: brand.Id,
		}
		err := repo.DB.Model(id).Updates(brand).Error
		if err != nil {
			log.Log(err)
			return false, err
		}
	} else {
		return false, fmt.Errorf("更新品牌不存在")
	}
	return true, nil
}

// Delete 删除品牌
func (repo *Brand) Delete(brand *pb.Brand) (bool, error) {
	if brand.Id == 0 {
		return false, fmt.Errorf("请传入操作id")
	}
	id := &pb.Brand{
		Id: brand.Id,
	}
	if err := repo.DB.Delete(id).Error; err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
