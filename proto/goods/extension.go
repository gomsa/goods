package goods

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/jinzhu/gorm"
)

// TimeLayout 转换字符

var (
	dateTime = time.Now().In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04:05")
)

// BeforeCreate 插入前数据处理
func (p *Good) BeforeCreate(scope *gorm.Scope) (err error) {
	uuid := uuid.NewV4()
	err = scope.SetColumn("Id", uuid.String())

	err = scope.SetColumn("CreatedAt", dateTime)
	if err != nil {
		return err
	}
	err = scope.SetColumn("UpdatedAt", dateTime)
	if err != nil {
		return err
	}
	return nil
}

// BeforeUpdate 更新前数据处理
func (p *Good) BeforeUpdate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("UpdatedAt", dateTime)
	if err != nil {
		return err
	}
	return nil
}

// BeforeCreate 插入前数据处理
func (p *Barcode) BeforeCreate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("CreatedAt", dateTime)
	if err != nil {
		return err
	}
	err = scope.SetColumn("UpdatedAt", dateTime)
	if err != nil {
		return err
	}
	return nil
}

// BeforeUpdate 更新前数据处理
func (p *Barcode) BeforeUpdate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("UpdatedAt", dateTime)
	if err != nil {
		return err
	}
	return nil
}
