package handler

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"zero/mxshop/goods_srv/global"
	"zero/mxshop/goods_srv/model"
	"zero/mxshop/goods_srv/proto"
)

func (gs *GoodsServer) GetAllCategorysList(ctx context.Context, empty *empty.Empty) (*proto.CategoryListResponse, error) {
	var category []model.Category
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&category)
	b, _ := json.Marshal(&category)
	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}

func (gs *GoodsServer) GetSubCategory(ctx context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	categoryListResponse := proto.SubCategoryListResponse{}
	var category model.Category

	queryRes := global.DB.First(&category, req.Id)
	if queryRes.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "商品分类不存在")
	}
	categoryListResponse.Info = &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		Level:          category.Level,
		IsTab:          category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}

	var subcateogries []model.Category
	var subCategoryResponse []*proto.CategoryInfoResponse

	global.DB.Where(&model.Category{ParentCategoryID: req.Id}).Find(&subcateogries)

	for _, subCategory := range subcateogries {
		subCategoryResponse = append(subCategoryResponse, &proto.CategoryInfoResponse{
			Id:             subCategory.ID,
			Name:           subCategory.Name,
			Level:          subCategory.Level,
			IsTab:          subCategory.IsTab,
			ParentCategory: subCategory.ParentCategoryID,
		})
	}

	categoryListResponse.SubCategorys = subCategoryResponse
	return &categoryListResponse, nil
}

func (gs *GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	category := model.Category{}
	cMap := map[string]interface{}{}

	cMap["name"] = req.Name
	cMap["level"] = req.Level
	cMap["is_tab"] = req.IsTab

	if req.Level != 1 {
		cMap["parent_category_id"] = req.ParentCategory
	}
	tx := global.DB.Model(&model.Category{}).Create(cMap)
	if tx.Error != nil {
		return nil, status.Error(codes.Internal, "创建category失败"+tx.Error.Error())
	}
	return &proto.CategoryInfoResponse{
		Id: category.ID,
	}, nil
}

func (gs *GoodsServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryRequest) (*empty.Empty, error) {
	queryRes := global.DB.Delete(&model.Category{}, req.Id)
	if queryRes.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "category 不存在")
	}
	return &empty.Empty{}, nil
}

func (gs *GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*empty.Empty, error) {
	var category model.Category
	queryRes := global.DB.First(&category, req.Id)
	if queryRes.RowsAffected == 0 {
		return nil, status.Error(codes.NotFound, "category 不存在")
	}
	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Level != 0 {
		category.Level = req.Level
	}
	if req.IsTab {
		category.IsTab = req.IsTab
	}

	queryRes = global.DB.Save(&category)
	if queryRes.Error != nil {
		return nil, status.Error(codes.Internal, "保存失败， error: "+queryRes.Error.Error())
	}
	return &empty.Empty{}, nil
}
