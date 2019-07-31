package hander

import (
	"context"

	pb "github.com/gomsa/goods/proto/taxcode"
	"github.com/gomsa/goods/service"
)

// Taxcode 国家税收分类编码结构
type Taxcode struct {
	Repo service.TaxcodeRepository
}

// All 获取所有国家税收分类编码
func (srv *Taxcode) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	taxcodes, err := srv.Repo.All(req.Taxcode)
	if err != nil {
		return err
	}
	res.Taxcodes = taxcodes
	return err
}

// List 获取所有国家税收分类编码
func (srv *Taxcode) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	taxcodes, err := srv.Repo.List(req.Taxcode)
	if err != nil {
		return err
	}
	res.Taxcodes = taxcodes
	return err
}

// Get 获取国家税收分类编码
func (srv *Taxcode) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	taxcode, err := srv.Repo.Get(req.Taxcode)
	if err != nil {
		return err
	}
	res.Taxcode = taxcode
	return err
}

// Create 创建国家税收分类编码
func (srv *Taxcode) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	taxcode, err := srv.Repo.Create(req.Taxcode)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Taxcode = taxcode
	res.Valid = true
	return err
}

// Update 更新国家税收分类编码
func (srv *Taxcode) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req.Taxcode)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}

// Delete 删除国家税收分类编码国家税收分类编码
func (srv *Taxcode) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.Taxcode)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}
