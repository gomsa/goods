package main

import (
	// micro 公共引入
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	k8s "github.com/micro/kubernetes/go/micro"

	// 服务引用
	_ "github.com/gomsa/goods/database/migrations"
	"github.com/gomsa/goods/hander"
	db "github.com/gomsa/goods/providers/database"
	"github.com/gomsa/goods/service"

	// 接口引用
	brandPB "github.com/gomsa/goods/proto/brand"
	categoryPB "github.com/gomsa/goods/proto/category"
	departmentPB "github.com/gomsa/goods/proto/department"
	firmPB "github.com/gomsa/goods/proto/firm"
	goodsPB "github.com/gomsa/goods/proto/goods"
	unspscPB "github.com/gomsa/goods/proto/unspsc"
)

func main() {
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
	)
	srv.Init()

	// 服务实现
	goodsPB.RegisterGoodsHandler(srv.Server(), &hander.Goods{&service.Goods{db.DB}})
	brandPB.RegisterBrandsHandler(srv.Server(), &hander.Brand{&service.Brand{db.DB}})
	firmPB.RegisterFirmsHandler(srv.Server(), &hander.Firm{&service.Firm{db.DB}})
	categoryPB.RegisterCategorysHandler(srv.Server(), &hander.Category{&service.Category{db.DB}})
	departmentPB.RegisterDepartmentsHandler(srv.Server(), &hander.Department{&service.Department{db.DB}})
	unspscPB.RegisterUnspscsHandler(srv.Server(), &hander.Unspsc{&service.Unspsc{db.DB}})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
