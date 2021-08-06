package handler

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero/mxshop/common/utils"
	"zero/mxshop/goods_srv/global"
	"zero/mxshop/goods_srv/model"
	"zero/mxshop/goods_srv/proto"
)

func (gs *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	response := &proto.BrandListResponse{}
	var brands []model.Brands
	res := global.DB.Scopes(utils.Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)
	if res.Error != nil {
		zap.S().Info("[BrandList] Failed" + res.Error.Error())
		return nil, res.Error
	}
	var total int64
	global.DB.Model(&model.Brands{}).Count(&total)

	var brandResponses []*proto.BrandInfoResponse

	for _, brand := range brands {
		brandResponseItem := &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		}
		brandResponses = append(brandResponses, brandResponseItem)
	}

	response.Data = brandResponses
	response.Total = int32(total)

	return response, nil
}

func (gs *GoodsServer) CreateBrand(ctx context.Context, req *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	b := &model.Brands{}
	res := global.DB.Where("name = ?", req.Name).First(b)
	fmt.Println("RowAffected:",res.RowsAffected)
	if res.RowsAffected >= 1 {
		return nil, status.Errorf(codes.AlreadyExists, "品牌已经存在")
	}
	brand := model.Brands{
		Name: req.Name,
		Logo: req.Logo,
	}
	res = global.DB.Save(&brand)
	if res.Error != nil {
		return nil, status.Errorf(codes.Internal, res.Error.Error())
	}
	response := proto.BrandInfoResponse{
		Id:   brand.ID,
		Name: brand.Name,
		Logo: brand.Logo,
	}
	return &response, nil
}

func (gs *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandRequest) (*empty.Empty, error) {
	instance := model.Brands{}
	queryRes := global.DB.Where("id = ?", req.Id).Find(&instance)
	if queryRes.RowsAffected != 1 {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}
	instance.Logo = req.Logo
	instance.Name = req.Name
	queryRes = global.DB.Save(&instance)
	if queryRes.Error != nil {
		return nil, status.Errorf(codes.Internal, queryRes.Error.Error())
	}
	return &empty.Empty{}, nil
}

func (gs *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandRequest) (*empty.Empty, error) {
	fmt.Println("Delete brand: ", req)
	if result := global.DB.Delete(&model.Brands{}, req.Id); result.RowsAffected	== 0{
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}
	return &empty.Empty{}, nil
}
