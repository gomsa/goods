package main

import (
	// 公共引入
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	k8s "github.com/micro/kubernetes/go/micro"

	// 执行数据迁移
	_ "github.com/gomsa/goods/database/migrations"

	"github.com/gomsa/goods/hander"
	goodsPB "github.com/gomsa/goods/proto/goods"
	db "github.com/gomsa/goods/providers/database"
	"github.com/gomsa/goods/service"
)

func main() {
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
	)
	srv.Init()

	// 权限服务实现
	repo := &service.Goods{db.DB}
	goodsPB.RegisterGoodsHandler(srv.Server(), &hander.Goods{repo})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
