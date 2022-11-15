package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
  c.String(http.StatusOK, "ログインAPI")
}

func UserRegister(c *gin.Context) {
  c.String(http.StatusOK, "会員登録API")
}

func UserInfoGet(c *gin.Context) {
  c.String(http.StatusOK, "ユーザー情報取得API")
}

func UserInfoSave(c *gin.Context) {
  c.String(http.StatusOK, "ユーザー情報保存API")
}

