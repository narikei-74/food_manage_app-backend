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
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // リクエストボディ取得
  type request struct {UserID int}
  requestData := request{}
  err := c.ShouldBindJSON(&requestData)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": err,
    })
    return
  }

  // 残り食材データ取得
  food_stocks := []model.My_food_stock{}
  foodStocksGetErr := dbHandle.Model(&model.My_food_stock{}).Preload("Food").Where("user_id = ?", requestData.UserID).Find(&food_stocks).Error

  // エラーの場合
  if (foodStocksGetErr != nil) {
    c.JSON(http.StatusOK, gin.H{
      "success": false,
      "error": foodStocksGetErr,
    })

    return
  }

  c.JSON(http.StatusOK, gin.H{
    "success": true,
    "data": food_stocks,
  })

  return

}

func FoodStockSave(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // リクエストボディ取得
  type request struct {
    Creates []model.My_food_stock
    Updates []model.My_food_stock
  }
  requestData := request{}
  requestErr := c.ShouldBindJSON(&requestData)
  if requestErr != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": requestErr,
    })
    return
  }

  // dbに保存
  tx := dbHandle.Begin()
  // 作成
  for i := 0; i < len(requestData.Creates); i++ {
    if err := tx.Create(&requestData.Creates[i]).Error; err != nil {
      tx.Rollback()
      c.JSON(http.StatusBadRequest, gin.H{
        "success": false,
        "error": err,
      })
      return
    }
  }
  // 更新
  for i := 0; i < len(requestData.Updates); i++ {
    if err := tx.Model(&model.My_food_stock{}).Where("ID = ?", requestData.Updates[i].ID).Updates(requestData.Updates[i]).Error; err != nil {
      tx.Rollback()

      c.JSON(http.StatusBadRequest, gin.H{
        "success": false,
        "error": err,
      })

      return
    }
  }

  tx.Commit()

  // レスポンス
  c.JSON(http.StatusOK, gin.H{
    "success": true,
  })

  return
}

func FoodStockDelete(c *gin.Context) {
  dbHandle := db.Init()
  defer dbHandle.Close()

  // リクエストボディ取得
  type request struct {My_food_stock_ID int}
  requestData := request{}
  requestErr := c.ShouldBindJSON(&requestData)
  if requestErr != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": requestErr,
    })
    return
  }

  // 削除
  tx := dbHandle.Begin()
  err := tx.Unscoped().Delete(&model.My_food_stock{}, requestData.My_food_stock_ID).Error
  if err != nil {
    tx.Rollback()
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": err,
    })
    return
  }

  tx.Commit()

  // レスポンス
  c.JSON(http.StatusOK, gin.H{
    "success": true,
  })

  return
}