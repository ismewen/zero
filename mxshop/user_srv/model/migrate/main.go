package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"zero/mxshop/user_srv/model"
)

func CreateUsers(db *gorm.DB){
	u := model.User{}
	for i:= 0; i<10; i++ {
		pwd, _ := u.GetMd5Str(fmt.Sprintf("%d", i))
		user := model.User{
			NickName: fmt.Sprintf("Somebody-%d", i),
			Mobile: fmt.Sprintf("9999%d", i),
			Password: pwd,
		}
		db.Save(&user)

	}
}

func main() {
	dsn := "root:ismewen@tcp(127.0.0.1:3306)/mx_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
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

	_ = db.AutoMigrate(&model.User{})
	CreateUsers(db)

}
