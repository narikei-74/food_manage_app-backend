package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() *gorm.DB{
	db, err := gorm.Open("mysql", "narita:narita1005@tcp(mysql)/food_manage?parseTime=true")

  if err != nil {
    panic(err.Error())
  }

	return db
}