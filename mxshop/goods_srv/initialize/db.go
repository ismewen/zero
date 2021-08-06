package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"zero/mxshop/goods_srv/global"
)

func CreateDB(dsn string) *gorm.DB {

	queryLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Error,
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
		panic("Init db failed")
	}

	return db
}

func InitDB() {
	mysqlConfig := global.ServerConfig.MySqlConfig
	fmt.Println(mysqlConfig.GetGormDsn())
	global.DB = CreateDB(mysqlConfig.GetGormDsn())
}
