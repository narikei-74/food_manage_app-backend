package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func user_login(c *gin.Context) {
  c.String(http.StatusOK, "ログインAPI")
}

func user_register(c *gin.Context) {
  c.String(http.StatusOK, "会員登録API")
}

func user_info_get(c *gin.Context) {
  c.String(http.StatusOK, "ユーザー情報取得API")
}

func user_info_save(c *gin.Context) {
  c.String(http.StatusOK, "ユーザー情報保存API")
}

