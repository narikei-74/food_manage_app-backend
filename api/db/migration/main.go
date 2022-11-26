package main

import (
	"github.com/narikei-74/food_manage_app-backend/api/db"
	"github.com/narikei-74/food_manage_app-backend/api/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := db.Connection()
  defer db.Close()

  db.AutoMigrate(&model.User{})
}