package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/narikei-74/food_manage_app-backend/api/db"
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
  dbHandle := db.Init()
  defer dbHandle.Close()

  user := model.User{}
  err := c.ShouldBindJSON(&user);
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": err.Error(),
    })
    return
  }

  result := dbHandle.Create(&user)

  if result.Error != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": result.Error,
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "success": true,
    "userId": user.ID,
  })
}

// ユーザー情報取得
func UserInfoGet(c *gin.Context) {
  c.String(http.StatusOK, "ユーザー情報取得API")
}

// ユーザー情報保存
func UserInfoSave(c *gin.Context) {
  c.String(http.StatusOK, "ユーザー情報保存API")
}

