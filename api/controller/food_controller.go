package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/narikei-74/food_manage_app-backend/api/db"
  "github.com/narikei-74/food_manage_app-backend/api/model"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func FoodsGet(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // 食材データ取得
  foods := []model.Food{}
  foodsGetErr := dbHandle.Where("Spices_flag = ?", 0).Find(&foods).Error

  // エラーの場合
  if (foodsGetErr != nil) {
    c.JSON(http.StatusOK, gin.H{
      "success": false,
      "error": foodsGetErr,
    })

    return
  }

  c.JSON(http.StatusOK, gin.H{
    "success": true,
    "data": foods,
  })
}

func FoodStockGet(c *gin.Context) {
  c.String(http.StatusOK, "残り食材取得API")
}

func FoodStockSave(c *gin.Context) {
  c.String(http.StatusOK, "残り食材保存API")
}

func ReceiptAnalysis(c *gin.Context) {
  c.String(http.StatusOK, "レシート解析API")
}