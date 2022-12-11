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

// ログイン
func UserLogin(c *gin.Context) {
  c.String(http.StatusOK, "ログインAPI")
}

// 会員登録
func UserRegister(c *gin.Context) {
  c.String(http.StatusOK, "会員登録API")
}

// ゲスト登録
func UserRegisterGuest(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // リクエストボディ
  requestData := model.User{}
  requestErr := c.ShouldBindJSON(&requestData);
  if requestErr != nil {
    log.Print(requestErr)
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
    })
    return
  }

  // dbに保存
  tx := dbHandle.Begin()
  user := model.User{}
  err := tx.Create(&user).Error
  if err != nil {
    tx.Rollback()
    log.Print(err)
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
    })
    return
  }

  tx.Commit()

  // レスポンス
  c.JSON(http.StatusOK, gin.H{
    "success": true,
    "userId": user.ID,
  })

  return
}

// ユーザー情報取得
func UserInfoGet(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // リクエストボディ取得
  requestData := model.User{}
  requestErr := c.ShouldBindJSON(&requestData)
  if requestErr != nil {
    log.Print(requestErr)
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
    })
    return
  }

  // dbから取得
  tx := dbHandle.Begin()
  user := model.User{}
  err := tx.Model(&model.User{}).Preload("User_family_infos").First(&user, user.ID).Error
  if err != nil {
    tx.Rollback()
    log.Print(err)
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
    })
    return
  }

  tx.Commit()

  // レスポンス
  c.JSON(http.StatusOK, gin.H{
    "success": true,
    "data": user,
  })

  return
}

// ユーザー情報登録
func UserInfoCreate(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // リクエストボディ取得
  requestData := model.User_family_info{}
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
  err := tx.Create(&requestData).Error
  if err != nil {
    tx.Rollback()
    log.Print(err)
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

// ユーザー情報編集
func UserInfoUpdate(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // リクエストボディ取得
  requestData := model.User_family_info{}
  requestErr := c.ShouldBindJSON(&requestData)
  if requestErr != nil {
    log.Print(requestErr)
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
    })
    return
  }

  // dbを更新
  tx := dbHandle.Begin()
  err := tx.Model(&model.User_family_info{}).Where("ID = ?", requestData.ID).Updates(requestData).Error
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