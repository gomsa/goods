package hander

import (
	"context"

	pb "github.com/gomsa/goods/proto/goods"
	"github.com/gomsa/goods/service"
)

// Goods 商品结构
type Goods struct {
	Repo service.GoodRepository
}

// List 获取所有商品
func (srv *Goods) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return err
}

// Get 获取商品
func (srv *Goods) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return err
}

// Create 创建商品
func (srv *Goods) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
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
