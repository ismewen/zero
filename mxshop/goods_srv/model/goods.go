package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type GormList []string

func (g GormList) Value() (driver.Value, error) {
	// 保存的时候如何保存
	return json.Marshal(g)
}

func (g *GormList) Scan(value interface{}) error {
	// 从数据库中读取一个数据的时候，如何处理
	return json.Unmarshal(value.([]byte), &g)
}

// 不用默认的 gorm.Model, 而是自己定义，加深理解
type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`
}

type Category struct {
	// 表
	BaseModel
	Name             string `gorm:"type:varchar(20);not null"`
	ParentCategoryID int32
	ParentCategory   *Category
	Level            int32 `gorm:"type:int;not null;default:1"`
	IsTab            bool  `gorm:"default:false;not null"`
}

type Brands struct {
	// 品牌表
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null"`
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:index_category_brand,unique"`
	Category   Category
	BrandsID    int32 `gorm:"type:int;index:index_category_brand,unique"`
	Brands     Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `gorm:"type:int;default:1;not null"`
}

type Goods struct {
	BaseModel

	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;not null"`
	Brands     Brands

	IsOnSale   bool `gorm:"default:false;not null"`
	IsShipFree bool `gorm:"default:false;not null"`
	IsHot      bool `gorm:"default:false;not null"`

	Name     string `gorm:"type:varchar(50);not null"`
	GoodsSn  string `gorm:"type:varchar(50);not null"`
	ClickNum int32  `gorm:"type:int;default:0;not null"`
	SoldNum  int32  `gorm:"type:int;default:0;not null"`
	FavNum   int32  `gorm:"type:int;default:0;not null"`

	MarketPrice float32 `gorm:"not null"`
	ShopPrice   float32 `gorm:"not null"`
	GoodsBrief  string  `gorm:"type:varchar(100);not null"`

	Images          GormList `gorm:"type:varchar(1000);not null"`
	DescImages      GormList `gorm:"type:varchar(1000);not null"`
	GoodsFrontImage string   `gorm:"type:varchar(50);not null"`
}

type GoodsImages struct {
	GoodsID int
	Image   string
}
