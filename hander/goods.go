package hander

import (
	"context"
	"fmt"

	pb "github.com/gomsa/goods/proto/goods"
	"github.com/gomsa/goods/service"
)

// Goods 商品结构
type Goods struct {
	Repo service.GoodRepository
}

// List 获取所有商品
func (srv *Goods) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	goods, err := srv.Repo.List(req)
	if err != nil {
		return err
	}
	res.Goods = goods
	return err
}

// Get 获取商品
func (srv *Goods) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	good, err := srv.Repo.Get(req.Good)
	if err != nil {
		return err
	}
	res.Good = good
	return err
}

// Create 创建商品
func (srv *Goods) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	good, err := srv.Repo.Create(req.Good)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("创建用户失败")
	}
	res.Good = good
	res.Valid = true
	return err
}

// Update 更新商品
func (srv *Goods) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return err
}

// Delete 删除商品商品
func (srv *Goods) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return err
}
