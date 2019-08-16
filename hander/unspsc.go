package hander

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	pb "github.com/gomsa/goods/proto/unspsc"
	"github.com/gomsa/goods/service"
)

// Unspsc 国际商品及服务编码结构
type Unspsc struct {
	Repo service.UnspscRepository
}

// CheckCreate 检查批量创建
func (srv *Unspsc) CheckCreate(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Unspsc == nil {
		return fmt.Errorf("请求参数错误")
	}
	if req.Unspsc.Id == 0 {
		return fmt.Errorf("请求参数 id 为空")
	}
	if req.Unspsc.Name == `` {
		return fmt.Errorf("请求参数 name 为空")
	}
	valid := srv.Repo.Exist(req.Unspsc)
	if valid {
		return fmt.Errorf("%s 国际商品及服务编码已存在", strconv.FormatInt(req.Unspsc.Id, 10))
	}
	if !valid {
		names := strings.Split(req.Unspsc.Name, ">>")
		for key, name := range names {
			// 计算 ID 父 parent
			// 补位
			spell := ``
			switch key {
			case 0:
				spell = `000000`
			case 1:
				spell = `0000`
			case 2:
				spell = `00`
			}
			s := strconv.FormatInt(req.Unspsc.Id, 10)
			id, _ := strconv.ParseInt(s[0:2+(2*key)]+spell, 10, 64)
			parent, _ := strconv.ParseInt(s[0:(2*key)]+spell+`00`, 10, 64)
			unspsc := &pb.Unspsc{
				Id:     id,
				Parent: parent,
				Name:   name,
			}
			srv.Repo.Create(unspsc)
		}
		valid = true
	}
	res.Valid = valid
	return err
}

// Exist 检查国际商品及服务编码
func (srv *Unspsc) Exist(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Unspsc == nil {
		return fmt.Errorf("请求参数错误")
	}
	valid := srv.Repo.Exist(req.Unspsc)
	res.Valid = valid
	return err
}

// All 获取所有国际商品及服务编码
func (srv *Unspsc) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Unspsc == nil {
		req.Unspsc = &pb.Unspsc{}
	}
	unspscs, err := srv.Repo.All(req.Unspsc)
	if err != nil {
		return err
	}
	res.Unspscs = unspscs
	return err
}

// List 获取所有国际商品及服务编码
func (srv *Unspsc) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Unspsc == nil {
		req.Unspsc = &pb.Unspsc{}
	}
	unspscs, err := srv.Repo.List(req.Unspsc)
	if err != nil {
		return err
	}
	res.Unspscs = unspscs
	return err
}

// Get 获取国际商品及服务编码
func (srv *Unspsc) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Unspsc == nil {
		return fmt.Errorf("请求参数错误")
	}
	unspsc, err := srv.Repo.Get(req.Unspsc)
	if err != nil {
		return err
	}
	res.Unspsc = unspsc
	return err
}

// Create 创建国际商品及服务编码
func (srv *Unspsc) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	// 防止空指针报错
	if req.Unspsc == nil {
		return fmt.Errorf("请求参数错误")
	}
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
	// 防止空指针报错
	if req.Unspsc == nil {
		return fmt.Errorf("请求参数错误")
	}
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
	// 防止空指针报错
	if req.Unspsc == nil {
		return fmt.Errorf("请求参数错误")
	}
	valid, err := srv.Repo.Delete(req.Unspsc)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}
