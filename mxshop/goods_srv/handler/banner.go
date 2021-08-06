package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero/mxshop/goods_srv/global"
	"zero/mxshop/goods_srv/model"
	"zero/mxshop/goods_srv/proto"
)

func (gs *GoodsServer) BannerList(ctx context.Context, req *empty.Empty) (*proto.BannerListResponse, error) {
	response := &proto.BannerListResponse{}
	var banners []model.Banner
	queryRes := global.DB.Find(&banners)
	if queryRes.Error != nil {
		return nil, status.Errorf(codes.Internal, queryRes.Error.Error())
	}
	var bannerResponse []*proto.BannerResponse
	for _, banner := range banners {
		bannerRes := &proto.BannerResponse{
			Id:    banner.ID,
			Image: banner.Image,
			Index: banner.Index,
			Url:   banner.Url,
		}
		bannerResponse = append(bannerResponse, bannerRes)
	}
	return response, nil
}

func (gs *GoodsServer) CreateBanner(ctx context.Context, req *proto.BannerResponse) (*proto.BannerResponse, error) {
	banner := model.Banner{
		Image: req.Image,
		Index: req.Index,
		Url:   req.Url,
	}
	queryRes := global.DB.Save(&banner)
	if queryRes.Error != nil {
		return nil, status.Error(codes.Internal, queryRes.Error.Error())
	}
	return &proto.BannerResponse{
		Id: banner.ID,
	}, nil
}

func (gs *GoodsServer) DeleteBanner(ctx context.Context, req *proto.BannerRequest) (*empty.Empty, error) {
	var banner model.Banner
	if result := global.DB.First(&banner, req.Id); result.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "轮播图不存在")
	}
	return &empty.Empty{}, nil
}

func (gs *GoodsServer) UpdateBanner(ctx context.Context, req *proto.BannerRequest) (*empty.Empty, error) {
	var banner model.Banner
	queryRes := global.DB.Where("id = ?", req.Id).Find(&banner)
	if queryRes.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "Banner 不存在")
	}
	banner.Index = req.Index
	banner.Image = req.Image

	queryRes = global.DB.Save(&banner)
	if queryRes.Error != nil {
		return nil, status.Error(codes.Internal, "save Banner失败, error"+queryRes.Error.Error())
	}
	return &empty.Empty{}, nil
}
