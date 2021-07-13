package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:ismewen@/learn_gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		panic("connect mysql failed")
	}
	defer db.Close()
}
