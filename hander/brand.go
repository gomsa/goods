package hander

import (
	"context"
	"fmt"

	pb "github.com/gomsa/goods/proto/brand"
	"github.com/gomsa/goods/service"
)

// Brand 品牌结构
type Brand struct {
	Repo service.BrandRepository
}

// Exist 获取所有品牌
func (srv *Brand) Exist(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Brand == nil {
		req.Brand = &pb.Brand{}
	}
	valid := srv.Repo.Exist(req.Brand)
	res.Valid = valid
	return err
}

// All 获取所有品牌
func (srv *Brand) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Brand == nil {
		req.Brand = &pb.Brand{}
	}
	brands, err := srv.Repo.All(req.Brand)
	if err != nil {
		return err
	}
	res.Brands = brands
	return err
}

// List 获取所有品牌
func (srv *Brand) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.ListQuery == nil {
		req.ListQuery = &pb.ListQuery{}
	}
	if req.Brand == nil {
		req.Brand = &pb.Brand{}
	}
	brands, err := srv.Repo.List(req.ListQuery, req.Brand)
	if err != nil {
		return err
	}
	res.Brands = brands
	return err
}

// Get 获取品牌
func (srv *Brand) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Brand == nil {
		return fmt.Errorf("请求参数错误")
	}
	brand, err := srv.Repo.Get(req.Brand)
	if err != nil {
		return err
	}
	res.Brand = brand
	return err
}

// Create 创建品牌
func (srv *Brand) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Brand == nil {
		return fmt.Errorf("请求参数错误")
	}
	brand, err := srv.Repo.Create(req.Brand)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Brand = brand
	res.Valid = true
	return err
}

// Update 更新品牌
func (srv *Brand) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Brand == nil {
		return fmt.Errorf("请求参数错误")
	}
	valid, err := srv.Repo.Update(req.Brand)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}

// Delete 删除品牌品牌
func (srv *Brand) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Brand == nil {
		return fmt.Errorf("请求参数错误")
	}
	valid, err := srv.Repo.Delete(req.Brand)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}
