package service

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"

	pb "github.com/gomsa/goods/proto/firm"
	"github.com/gomsa/tools/uitl"
)

// Firm 商品公司仓库失效
type Firm struct {
	DB *gorm.DB
}

// Exist 查询商品公司是否存在
func (repo *Firm) Exist(firm *pb.Firm) bool {
	var count int
	if firm.Id != 0 {
		repo.DB.Model(&firm).Where("id = ?", firm.Id).Count(&count)
		if count > 0 {
			return true
		}
	}
	if firm.Name != "" {
		repo.DB.Model(&firm).Where("name = ?", firm.Name).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// All 获取所有商品公司信息
func (repo *Firm) All(req *pb.Firm) (firms []*pb.Firm, err error) {
	if err := repo.DB.Find(&firms).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return firms, nil
}

// List 获取所有商品公司信息
func (repo *Firm) List(listQuery *pb.ListQuery, firm *pb.Firm) (firms []*pb.Firm, err error) {
	db := repo.DB.Model(&firm)
	// 计算分页
	limit, offset := uitl.Page(listQuery.Limit, listQuery.Page)
	// 排序
	var sort string
	if listQuery.Sort != "" {
		sort = listQuery.Sort
	} else {
		sort = "created_at desc"
	}
	// 查询条件
	if firm.Name != "" {
		db = db.Where("name like ?", "%"+firm.Name+"%")
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&firms).Error; err != nil {
		log.Log(err)
		return firms, err
	}
	return firms, err
}

// Get 获取商品公司信息
func (repo *Firm) Get(firm *pb.Firm) (*pb.Firm, error) {
	if err := repo.DB.Model(&firm).Find(&firm).Error; err != nil {
		return firm, err
	}
	return firm, nil
}

// Create 创建商品公司
func (repo *Firm) Create(firm *pb.Firm) (*pb.Firm, error) {
	if valid := repo.Exist(firm); valid {
		return firm, fmt.Errorf("公司已存在")
	}
	err := repo.DB.Create(firm).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return firm, fmt.Errorf("添加商品公司失败")
	}
	return firm, nil
}

// Update 更新商品公司
func (repo *Firm) Update(firm *pb.Firm) (bool, error) {
	if firm.Id == 0 {
		return false, fmt.Errorf("请传入操作id")
	}
	if valid := repo.Exist(firm); valid {
		id := &pb.Firm{
			Id: firm.Id,
		}
		err := repo.DB.Model(id).Updates(firm).Error
		if err != nil {
			log.Log(err)
			return false, err
		}
	} else {
		return false, fmt.Errorf("更新商品公司不存在")
	}
	return true, nil
}

// Delete 删除商品公司
func (repo *Firm) Delete(firm *pb.Firm) (bool, error) {
	if firm.Id == 0 {
		return false, fmt.Errorf("请传入操作id")
	}
	id := &pb.Firm{
		Id: firm.Id,
	}
	if err := repo.DB.Delete(id).Error; err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
