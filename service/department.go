package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"

	pb "github.com/gomsa/goods/proto/department"
)

// Department 部门仓库失效
type Department struct {
	DB *gorm.DB
}

// All 获取所有部门信息
func (repo *Department) All(req *pb.Department) (departments []*pb.Department, err error) {
	if err := repo.DB.Find(&departments).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return departments, nil
}

// List 获取所有部门信息
func (repo *Department) List(req *pb.Department) (departments []*pb.Department, err error) {
	if err := repo.DB.Where("parent = ?", req.Parent).Find(&departments).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return departments, nil
}

// Get 获取部门信息
func (repo *Department) Get(department *pb.Department) (*pb.Department, error) {
	if err := repo.DB.Model(&department).Find(&department).Error; err != nil {
		return department, err
	}
	return department, nil
}

// Create 创建部门
func (repo *Department) Create(department *pb.Department) (*pb.Department, error) {
	err := repo.DB.Create(department).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return department, fmt.Errorf("添加部门失败")
	}
	return department, nil
}

// Update 更新部门
func (repo *Department) Update(department *pb.Department) (bool, error) {
	if department.Id == 0 {
		return false, fmt.Errorf("请传入操作id")
	}
	id := &pb.Department{
		Id: department.Id,
	}
	err := repo.DB.Model(id).Updates(department).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除部门
func (repo *Department) Delete(department *pb.Department) (bool, error) {
	if department.Id == 0 {
		return false, fmt.Errorf("请传入操作id")
	}
	id := &pb.Department{
		Id: department.Id,
	}
	if err := repo.DB.Delete(id).Error; err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
