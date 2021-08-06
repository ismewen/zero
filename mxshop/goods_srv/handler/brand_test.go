package handler

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"os"
	"testing"
	"zero/mxshop/goods_srv/initialize"
	"zero/mxshop/goods_srv/model"
	"zero/mxshop/goods_srv/proto"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important

)

var GC proto.GoodsClient
var DB *gorm.DB

func setup() {
	//serviceName := "goods-rpc"
	//host := "127.0.0.1"
	//port := 8500
	address := "consul://127.0.0.1:8500/goods-rpc?wait=14s"
	fmt.Printf("address: %s\n", address)

	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		panic("初始化失败")
	}

	client := proto.NewGoodsClient(conn)
	GC = client

	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	dsn = fmt.Sprintf(dsn, "root", "ismewen", "127.0.0.1", 3306, "mx_goods_srv")
	DB = initialize.CreateDB(dsn)

}

func TestGoodsServer_BrandList(t *testing.T) {
	res, err := GC.BrandList(context.Background(), &proto.BrandFilterRequest{
		Pages:       1,
		PagePerNums: 2,
	})
	if err != nil {
		fmt.Sprintln(err.Error())
		t.Fatal()
	}
	fmt.Println(res)

}

func TestGoodsServer_CreateBrand(t *testing.T) {
	res, err := GC.CreateBrand(context.Background(), &proto.BrandRequest{
		Name: "test-name-rand-3",
		Logo: "test-logo",
	})
	if err != nil {
		panic("创建brand失败" + err.Error())
	}
	fmt.Printf("%+v\n", res)
}

func TestGoodsServer_DeleteBrand(t *testing.T) {
	brand := model.Brands{}
	queryRes := DB.First(&brand)
	if queryRes.Error != nil {
		panic(queryRes.Error.Error())
	}
	res, err := GC.DeleteBrand(context.Background(), &proto.BrandRequest{
		Id: brand.ID,
	})
	if err != nil {
		panic("删除brand失败," + err.Error())
	}
	fmt.Printf("%+v\n", res)
}

func TestGoodsServer_UpdateBrand(t *testing.T) {
	brand := model.Brands{}
	queryRes := DB.Find(&brand)
	if queryRes.Error != nil {
		panic("数据错误")
	}

	_, err := GC.UpdateBrand(context.Background(), &proto.BrandRequest{
		Id: brand.ID,
		Name: "newBrandName",
		Logo: "newLogo",
	})
	if err != nil {
		t.Error("修改失败" + err.Error())
	}
	queryRes = DB.Find(&brand)
	if queryRes.Error != nil {
		panic("数据错误")
	}
	if brand.Logo != "newLogo" {
		t.Error("logo 不符合预期")
	}

	if brand.Name != "newBrandName" {
		t.Error("name 不符合预期")
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
