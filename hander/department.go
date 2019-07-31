package hander

import (
	"context"

	pb "github.com/gomsa/goods/proto/department"
	"github.com/gomsa/goods/service"
)

// Department 部门结构
type Department struct {
	Repo service.DepartmentRepository
}

// All 获取所有部门
func (srv *Department) All(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	departments, err := srv.Repo.All(req.Department)
	if err != nil {
		return err
	}
	res.Departments = departments
	return err
}

// List 获取所有部门
func (srv *Department) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	departments, err := srv.Repo.List(req.Department)
	if err != nil {
		return err
	}
	res.Departments = departments
	return err
}

// Get 获取部门
func (srv *Department) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	department, err := srv.Repo.Get(req.Department)
	if err != nil {
		return err
	}
	res.Department = department
	return err
}

// Create 创建部门
func (srv *Department) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	department, err := srv.Repo.Create(req.Department)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Department = department
	res.Valid = true
	return err
}

// Update 更新部门
func (srv *Department) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req.Department)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}

// Delete 删除部门部门
func (srv *Department) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req.Department)
	if err != nil {
		res.Valid = false
		return err
	}
	res.Valid = valid
	return err
}
