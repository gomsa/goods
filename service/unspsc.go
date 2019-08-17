package service

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"

	pb "github.com/gomsa/goods/proto/unspsc"
)

// Unspsc 国际商品及服务编码仓库失效
type Unspsc struct {
	DB *gorm.DB
}

// Exist 查询国际商品及服务编码是否存在
func (repo *Unspsc) Exist(unspsc *pb.Unspsc) (bool, error) {
	var count int
	if unspsc.Id != 0 {
		repo.DB.Model(&unspsc).Where("id = ?", unspsc.Id).Count(&count)
		if count > 0 {
			return true, fmt.Errorf("%s 国际商品及服务编码已存在", strconv.FormatInt(unspsc.Id, 10))
		}
	}
	if unspsc.Name != `` {
		repo.DB.Model(&unspsc).Where("name = ?", unspsc.Name).Count(&count)
		if count > 0 {
			return true, fmt.Errorf("%s 国际商品及服务编码已存在", unspsc.Name)
		}
	}
	log.Log(unspsc, count)
	return false, nil
}

// All 获取所有国际商品及服务编码信息
func (repo *Unspsc) All(req *pb.Unspsc) (unspscs []*pb.Unspsc, err error) {
	if err := repo.DB.Find(&unspscs).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return unspscs, nil
}

// List 获取所有国际商品及服务编码信息
func (repo *Unspsc) List(req *pb.Unspsc) (unspscs []*pb.Unspsc, err error) {
	if err := repo.DB.Where("parent = ?", req.Parent).Find(&unspscs).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return unspscs, nil
}

// Get 获取国际商品及服务编码信息
func (repo *Unspsc) Get(unspsc *pb.Unspsc) (*pb.Unspsc, error) {
	if err := repo.DB.Model(&unspsc).Find(&unspsc).Error; err != nil {
		return unspsc, err
	}
	return unspsc, nil
}

// Create 创建国际商品及服务编码
func (repo *Unspsc) Create(unspsc *pb.Unspsc) (*pb.Unspsc, error) {
	if valid, err := repo.Exist(unspsc); valid {
		return unspsc, err
	}
	err := repo.DB.Create(unspsc).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return unspsc, fmt.Errorf("添加国际商品及服务编码失败")
	}
	return unspsc, nil
}

// Update 更新国际商品及服务编码
func (repo *Unspsc) Update(unspsc *pb.Unspsc) (bool, error) {
	if unspsc.Id == 0 {
		return false, fmt.Errorf("请传入操作id")
	}
	id := &pb.Unspsc{
		Id: unspsc.Id,
	}
	err := repo.DB.Model(id).Updates(unspsc).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除国际商品及服务编码
func (repo *Unspsc) Delete(unspsc *pb.Unspsc) (bool, error) {
	if unspsc.Id == 0 {
		return false, fmt.Errorf("请传入操作id")
	}
	id := &pb.Unspsc{
		Id: unspsc.Id,
	}
	if err := repo.DB.Delete(id).Error; err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
