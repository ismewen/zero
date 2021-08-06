package handler

import "zero/mxshop/goods_srv/proto"

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

func (gs *GoodsServer) SayHello() {

}
