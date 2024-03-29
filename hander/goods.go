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

// GoodsByBarcode 根据条形码查询商品
func (srv *Goods) GoodsByBarcode(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Barcode == nil {
		return fmt.Errorf("请求参数错误")
	}
	// 获取条形码商品 ID
	barcode, err := srv.Repo.GetBarcode(req.Barcode)
	if err != nil {
		return err
	}
	// 通过商品 ID 查询商品信息
	good, err := srv.Repo.Get(&pb.Good{
		Id: barcode.GoodId,
	})
	if err != nil {
		return err
	}
	res.Good = good
	return err
}

// Exist 查询条码是否存在
func (srv *Goods) Exist(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Good == nil {
		return fmt.Errorf("请求参数错误")
	}
	valid, err := srv.Repo.Exist(req.Good)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}

// DeleteBarcode 删除条码
func (srv *Goods) DeleteBarcode(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Barcode == nil {
		return fmt.Errorf("请求参数错误")
	}
	valid, err := srv.Repo.DeleteBarcode(req.Barcode)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}

// List 获取所有商品
func (srv *Goods) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.ListQuery == nil {
		req.ListQuery = &pb.ListQuery{}
	}
	if req.Good == nil {
		req.Good = &pb.Good{}
	}
	goods, err := srv.Repo.List(req.ListQuery, req.Good)
	if err != nil {
		return err
	}
	total, err := srv.Repo.Total(req.Good)
	if err != nil {
		return err
	}
	res.Goods = goods
	res.Total = total
	return err
}

// Get 获取商品
func (srv *Goods) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Good == nil {
		return fmt.Errorf("请求参数错误")
	}
	good, err := srv.Repo.Get(req.Good)
	if err != nil {
		return err
	}
	res.Good = good
	return err
}

// Create 创建商品
func (srv *Goods) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Good == nil {
		return fmt.Errorf("请求参数错误")
	}
	good, err := srv.Repo.Create(req.Good)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Good = good
	res.Valid = true
	return err
}

// Update 更新商品
func (srv *Goods) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Good == nil {
		return fmt.Errorf("请求参数错误")
	}
	valid, err := srv.Repo.Update(req.Good)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}

// Delete 删除商品商品
func (srv *Goods) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Good == nil {
		return fmt.Errorf("请求参数错误")
	}
	valid, err := srv.Repo.Delete(req.Good)
	if err != nil {
		res.Valid = false
		return err
	}
	if valid {
		srv.Repo.DeleteBarcodeByGoodID(&pb.Barcode{
			GoodId: req.Good.Id,
		})
	}
	res.Valid = valid
	return err
}
