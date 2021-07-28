package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"zero/mxshop/goods_srv/model"
)

func main() {
	dsn := "root:ismewen@tcp(127.0.0.1:3306)/mx_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	queryLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: queryLogger,
		},
	)
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(
		&model.Category{},
		&model.Brands{},
		&model.GoodsCategoryBrand{},
		&model.Goods{},
	)

	if err != nil {
		panic(err.Error())
	}

}
