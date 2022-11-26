package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/narikei-74/food_manage_app-backend/api/db"
  "github.com/narikei-74/food_manage_app-backend/api/model"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func RecipeDataGet(c *gin.Context) {
  dbHandle := db.Init()
  defer dbHandle.Close()
  recipes := []model.Recipe{}
  err := dbHandle.Model(&model.Recipe{}).Preload("Recipe_materials").Preload("Recipe_categories").Find(&recipes).Error

  // エラーの場合
  if (err != nil) {
    c.JSON(http.StatusOK, gin.H{
      "success": false,
      "data": "",
    })

    return
  }

  c.JSON(http.StatusOK, gin.H{
    "success": true,
    "data": recipes,
  })

  return
}

func MyRecipeDataSave(c *gin.Context) {
  c.String(http.StatusOK, "my献立保存API")
}

func MyRecipeDataGet(c *gin.Context) {
  c.String(http.StatusOK, "test")
}

func RecipeCreateSettingGet(c *gin.Context) {
  c.String(http.StatusOK, "献立自動作成条件取得API")
}

func RecipeCreateSettingSave(c *gin.Context) {
  c.String(http.StatusOK, "献立自動作成条件保存API")
}