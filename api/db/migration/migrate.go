package main

import (
	"github.com/narikei-74/food_manage_app-backend/api/db"
	"github.com/narikei-74/food_manage_app-backend/api/model"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := db.Init()
  defer db.Close()

  db.AutoMigrate(&model.User{})
  db.AutoMigrate(&model.User_family_info{})
  db.AutoMigrate(&model.Recipe{})
  db.AutoMigrate(&model.Recipe_material{})
  db.AutoMigrate(&model.Recipe_category{})
  db.AutoMigrate(&model.My_recipe{})
  db.AutoMigrate(&model.Food{})
}