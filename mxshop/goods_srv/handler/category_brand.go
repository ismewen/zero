package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zero/mxshop/common/utils"
	"zero/mxshop/goods_srv/global"
	"zero/mxshop/goods_srv/model"
	"zero/mxshop/goods_srv/proto"
)

func (gs *GoodsServer) CategoryBrandList(ctx context.Context, req *proto.CategoryBrandFilterRequest) (*proto.CategoryBrandListResponse, error) {
	var categoryBrands []model.GoodsCategoryBrand
	categoryBrandListResponse := proto.CategoryBrandListResponse{}

	var total int64

	global.DB.Model(&model.GoodsCategoryBrand{}).Count(&total)

	categoryBrandListResponse.Total = int32(total)

	global.DB.Preload("Category").Preload("Brands").Scopes(utils.Paginate(int(req.Pages), int(req.PagePerNums))).Find(&categoryBrands)

	var categoryResponses []*proto.CategoryBrandResponse

	for _, cb := range categoryBrands {
		cbg := cb.Category
		cbb := cb.Brands
		categoryResponses = append(categoryResponses, &proto.CategoryBrandResponse{
			Category: &proto.CategoryInfoResponse{
				Id:             cbg.ID,
				Name:           cbg.Name,
				Level:          cbg.Level,
				IsTab:          cbg.IsTab,
				ParentCategory: cbg.ParentCategoryID,
			},
			Brand: &proto.BrandInfoResponse{
				Id:   cbb.ID,
				Name: cbb.Name,
				Logo: cbb.Logo,
			},
		})
	}
	categoryBrandListResponse.Data = categoryResponses
	return &categoryBrandListResponse, nil

}

func (gs *GoodsServer) GetCategoryBrandList(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.BrandListResponse, error) {
	brandListResponse := proto.BrandListResponse{}

	var category model.Category

	queryRes := global.DB.Find(&category, req.Id).First(&category)
	if queryRes.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "商品分类不存在")
	}

	var categoryBrands []*model.GoodsCategoryBrand
	queryRes = global.DB.Preload("Brands").Where(&model.GoodsCategoryBrand{CategoryID: req.Id}).Find(&categoryBrands)
	if queryRes.RowsAffected > 0 {
		brandListResponse.Total = int32(queryRes.RowsAffected)
	}

	var brandInfoRespons []*proto.BrandInfoResponse
	for _, cb := range categoryBrands {
		cbb := cb.Brands
		item := &proto.BrandInfoResponse{
			Id:   cbb.ID,
			Name: cbb.Name,
			Logo: cbb.Logo,
		}
		brandInfoRespons = append(brandInfoRespons, item)
	}
	brandListResponse.Data = brandInfoRespons

	return &brandListResponse, nil
}

func (gs *GoodsServer) CreateCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*proto.CategoryBrandResponse, error) {
	var category model.Category
	queryRes := global.DB.First(&category, req.CategoryId)
	if queryRes.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}
	var brand model.Brands
	queryRes = global.DB.First(&brand, req.BrandId)
	if queryRes.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	categoryBrand := model.GoodsCategoryBrand{
		CategoryID: req.CategoryId,
		BrandsID:   req.BrandId,
	}

	queryRes = global.DB.Save(&categoryBrand)
	if queryRes.Error != nil {
		return nil, status.Errorf(codes.Internal, "保存失败， error: "+queryRes.Error.Error())
	}
	return &proto.CategoryBrandResponse{
		Id: categoryBrand.ID,
	}, nil
}

func (gs *GoodsServer) DeleteCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*empty.Empty, error) {
	queryRes := global.DB.Delete(&model.GoodsCategoryBrand{}, req.Id)
	if queryRes.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, "分类不存在")
	}
	return &empty.Empty{}, nil
}

func (gs *GoodsServer) UpdateCategoryBrand(ctx context.Context, req *proto.CategoryBrandRequest) (*empty.Empty, error) {
	var categoryBrand model.GoodsCategoryBrand
	queryRes := global.DB.First(&categoryBrand, req.Id)
	if queryRes.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, "品牌分类不存在")
	}

	var category model.Category

	queryRes = global.DB.First(&category, req.CategoryId)
	if queryRes.RowsAffected == 0 {
		return nil, status.Error(codes.InvalidArgument, "品牌不存在")
	}

	categoryBrand.CategoryID = req.CategoryId
	categoryBrand.BrandsID = req.BrandId

	global.DB.Save(&category)

	return &empty.Empty{}, nil
}



