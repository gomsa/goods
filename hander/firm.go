package hander

import (
	"context"

	pb "github.com/gomsa/goods/proto/firm"
	"github.com/gomsa/goods/service"
)

// Firm 商品公司结构
type Firm struct {
	Repo service.FirmRepository
}

// Exist 获取所有商品公司
func (srv *Firm) Exist(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Exist(req.Firm)
	if err != nil {
		return err
	}
	res.Valid = valid
	return err
}

// All 获取所有商品公司
func (srv *Firm) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	firms, err := srv.Repo.All(req.Firm)
	if err != nil {
		return err
	}
	res.Firms = firms
	return err
}

// List 获取所有商品公司
func (srv *Firm) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	firms, err := srv.Repo.List(req.ListQuery, req.Firm)
	if err != nil {
		return err
	}
	res.Firms = firms
	return err
}

// Get 获取商品公司
func (srv *Firm) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	firm, err := srv.Repo.Get(req.Firm)
	if err != nil {
		return err
	}
	res.Firm = firm
	return err
}

// Create 创建商品公司
func (srv *Firm) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	firm, err := srv.Repo.Create(req.Firm)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Firm = firm
	res.Valid = true
	return err
}

// Update 更新商品公司
func (srv *Firm) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req.Firm)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}

// Delete 删除商品公司商品公司
func (srv *Firm) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.Firm)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}
