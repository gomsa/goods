package hander

import (
	"context"

	pb "github.com/gomsa/goods/proto/category"
	"github.com/gomsa/goods/service"
)

// Category 分类结构
type Category struct {
	Repo service.CategoryRepository
}

// All 获取所有分类
func (srv *Category) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	categorys, err := srv.Repo.All(req.Category)
	if err != nil {
		return err
	}
	res.Categorys = categorys
	return err
}

// List 获取所有分类
func (srv *Category) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	categorys, err := srv.Repo.List(req.Category)
	if err != nil {
		return err
	}
	res.Categorys = categorys
	return err
}

// Get 获取分类
func (srv *Category) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	category, err := srv.Repo.Get(req.Category)
	if err != nil {
		return err
	}
	res.Category = category
	return err
}

// Create 创建分类
func (srv *Category) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	category, err := srv.Repo.Create(req.Category)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Category = category
	res.Valid = true
	return err
}

// Update 更新分类
func (srv *Category) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req.Category)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}

// Delete 删除分类分类
func (srv *Category) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.Category)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}
