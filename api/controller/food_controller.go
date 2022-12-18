package controller

import (
  "net/http"
  "log"

  "github.com/gin-gonic/gin"
  "github.com/narikei-74/food_manage_app-backend/api/db"
  "github.com/narikei-74/food_manage_app-backend/api/logFile"
  "github.com/narikei-74/food_manage_app-backend/api/model"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func FoodsGet(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // dbから取得
  tx := dbHandle.Begin()
  foods := []model.Food{}
  err := tx.Find(&foods).Error
  if (err != nil) {
    log.Print(err)
    tx.Rollback()
    c.JSON(http.StatusOK, gin.H{
      "success": false,
    })
    return
  }

  tx.Commit()

  // レスポンス
  c.JSON(http.StatusOK, gin.H{
    "success": true,
    "data": foods,
  })
}

func FoodStockGet(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // リクエストボディ取得
  type request struct {UserID int}
  requestData := request{}
  requestErr := c.ShouldBindJSON(&requestData)
  if requestErr != nil {
    log.Print(requestErr)
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
    })
    return
  }

  // 残り食材データ取得
  tx := dbHandle.Begin()
  food_stocks := []model.My_food_stock{}
  err := tx.Model(&model.My_food_stock{}).Preload("Food").Where("user_id = ?", requestData.UserID).Find(&food_stocks).Error
  if (err != nil) {
    log.Print(err)
    tx.Rollback()
    c.JSON(http.StatusOK, gin.H{
      "success": false,
    })
    return
  }

  tx.Commit()

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

  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // リクエストボディ取得
  type request struct {
    Creates []model.My_food_stock
    Updates []model.My_food_stock
  }
  requestData := request{}
  requestErr := c.ShouldBindJSON(&requestData)
  if requestErr != nil {
    log.Print(requestErr)
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
    })
    return
  }

  // dbに保存
  tx := dbHandle.Begin()
  // 作成
  for i := 0; i < len(requestData.Creates); i++ {
    err := tx.Create(&requestData.Creates[i]).Error;
    if err != nil {
      log.Print(err)
      tx.Rollback()
      c.JSON(http.StatusBadRequest, gin.H{
        "success": false,
      })
      return
    }
  }

  // 更新
  for i := 0; i < len(requestData.Updates); i++ {
    err := tx.Model(&model.My_food_stock{}).Where("ID = ?", requestData.Updates[i].ID).Updates(requestData.Updates[i]).Error;
    if err != nil {
      log.Print(err)
      tx.Rollback()
      c.JSON(http.StatusBadRequest, gin.H{
        "success": false,
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
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // リクエストボディ取得
  type request struct {My_food_stock_ID int}
  requestData := request{}
  requestErr := c.ShouldBindJSON(&requestData)
  if requestErr != nil {
    log.Print(requestErr)
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
    })
    return
  }

  // 削除
  tx := dbHandle.Begin()
  err := tx.Unscoped().Delete(&model.My_food_stock{}, requestData.My_food_stock_ID).Error
  if err != nil {
    log.Print(err)
    tx.Rollback()
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
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