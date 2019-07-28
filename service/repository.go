package service

import (
	pb "github.com/gomsa/goods/proto/goods"
)

// GoodRepository  商品仓库接口
type GoodRepository interface {
	IsBarcode(*pb.Good) (bool, error)
	Create(*pb.Good) (*pb.Good, error)
	Get(*pb.Good) (*pb.Good, error)
	List(*pb.ListQuery) ([]*pb.Good, error)
	Update(*pb.Good) (bool, error)
	Delete(*pb.Good) (bool, error)
}
