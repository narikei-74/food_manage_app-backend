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
  user := model.User{}
  requestErr := c.ShouldBindJSON(&user);
  if requestErr == nil {
    log.Print(requestErr)
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
    })
    return
  }

  // dbに保存
  tx := dbHandle.Begin()
  err := tx.Create(&user).Error
  if err != nil {
    tx.Rollback()
    log.Print(requestErr)
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
  user := model.User{}
  requestErr := c.ShouldBindJSON(&user)
  if requestErr != nil {
    log.Print(requestErr)
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
    })
    return
  }

  // dbから取得
  tx := dbHandle.Begin()
  err := dbHandle.Model(&model.User{}).Preload("User_family_infos").First(&user, user.ID).Error
  if err != nil {
    tx.Rollback()
    log.Print(requestErr)
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

// ユーザー情報保存
func UserInfoSave(c *gin.Context) {
  c.String(http.StatusOK, "ユーザー情報保存API")
}

