package main

import (
	// 公共引入
	"github.com/micro/go-log"
	micro "github.com/micro/go-micro"
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
	repo := &service.GoodsRepository{db.DB}
	goodsPB.RegisterGoodssHandler(srv.Server(), &hander.Goods{repo})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
