package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() *gorm.DB{
	db, err := gorm.Open("mysql", "narita:narita1005@tcp(localhost:3306)/food_manage")

  if err != nil {
    panic(err.Error())
  }

	return db
}