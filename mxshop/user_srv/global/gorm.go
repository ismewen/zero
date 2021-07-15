package global

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

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
func init() {
	dsn := "root:ismewen@tcp(127.0.0.1:3306)/mx_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	DB = CreateDB(dsn)
}
