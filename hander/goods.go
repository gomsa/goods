package hander

import (
	"context"
	"fmt"

	pb "github.com/gomsa/goods/proto/goods"
	"github.com/gomsa/goods/service"
)

// Goods 部门结构
type Goods struct {
	Repo service.Repository
}

// All 获取所有权限
func (srv *Goods) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	goodss, err := srv.Repo.All(req)
	if err != nil {
		return err
	}
	res.Goodss = goodss
	return err
}

// List 获取所有部门
func (srv *Goods) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	goodss, err := srv.Repo.List(req)
	if err != nil {
		return err
	}
	res.Goodss = goodss
	return err
}

// Get 获取部门
func (srv *Goods) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	goods, err := srv.Repo.Get(req)
	if err != nil {
		return err
	}
	res.Goods = goods
	return err
}

// Create 创建部门
func (srv *Goods) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("创建部门失败")
	}
	res.Valid = true
	return err
}

// Update 更新部门
func (srv *Goods) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新部门失败")
	}
	res.Valid = valid
	return err
}

// Delete 删除部门部门
func (srv *Goods) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除部门失败")
	}
	res.Valid = valid
	return err
}
