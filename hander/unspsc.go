package hander

import (
	"context"

	pb "github.com/gomsa/goods/proto/unspsc"
	"github.com/gomsa/goods/service"
)

// Unspsc 国际商品及服务编码结构
type Unspsc struct {
	Repo service.UnspscRepository
}

// All 获取所有国际商品及服务编码
func (srv *Unspsc) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	unspscs, err := srv.Repo.All(req.Unspsc)
	if err != nil {
		return err
	}
	res.Unspscs = unspscs
	return err
}

// List 获取所有国际商品及服务编码
func (srv *Unspsc) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	unspscs, err := srv.Repo.List(req.Unspsc)
	if err != nil {
		return err
	}
	res.Unspscs = unspscs
	return err
}

// Get 获取国际商品及服务编码
func (srv *Unspsc) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	unspsc, err := srv.Repo.Get(req.Unspsc)
	if err != nil {
		return err
	}
	res.Unspsc = unspsc
	return err
}

// Create 创建国际商品及服务编码
func (srv *Unspsc) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	unspsc, err := srv.Repo.Create(req.Unspsc)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Unspsc = unspsc
	res.Valid = true
	return err
}

// Update 更新国际商品及服务编码
func (srv *Unspsc) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req.Unspsc)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}

// Delete 删除国际商品及服务编码国际商品及服务编码
func (srv *Unspsc) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.Unspsc)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}
