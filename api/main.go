package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  // ログインAPI
  r.POST("/user/login", func(c *gin.Context) {
    c.String(http.StatusOK, "ログインAPI")
  })

  // 会員登録API
  r.POST("/user/register", func(c *gin.Context) {
    c.String(http.StatusOK, "会員登録API")
  })

  //  ユーザー情報取得API
  r.POST("/user/info/get", func(c *gin.Context) {
    c.String(http.StatusOK, "ユーザー情報取得API")
  })

  // ユーザー情報保存API
  r.POST("/user/info/save", func(c *gin.Context) {
    c.String(http.StatusOK, "ユーザー情報保存API")
  })

  // 献立一覧取得API
  r.POST("/recipedata/get", func(c *gin.Context) {
    c.String(http.StatusOK, "献立一覧取得API")
  })

  // my献立保存API
  r.POST("/myrecipedata/save", func(c *gin.Context) {
    c.String(http.StatusOK, "my献立保存API")
  })

  // my献立取得API
  r.POST("/myrecipedata/get", func(c *gin.Context) {
    c.String(http.StatusOK, "my献立取得API")
  })

  // 献立自動作成条件取得API
  r.POST("/recipe_create_setting/get", func(c *gin.Context) {
    c.String(http.StatusOK, "献立自動作成条件取得API")
  })

  // 献立自動作成条件保存API
  r.POST("/recipe_create_setting/save", func(c *gin.Context) {
    c.String(http.StatusOK, "献立自動作成条件保存API")
  })

  // 残り食材取得API
  r.POST("/food_scock/get", func(c *gin.Context) {
    c.String(http.StatusOK, "残り食材取得API")
  })

  // 残り食材保存API
  r.POST("/food_scock/save", func(c *gin.Context) {
    c.String(http.StatusOK, "残り食材保存API")
  })

  // 食材市場データ取得API
  r.POST("/food_market/get", func(c *gin.Context) {
    c.String(http.StatusOK, "食材市場データ取得API")
  })

  // レシート解析API
  r.POST("/receipt_analysis", func(c *gin.Context) {
    c.String(http.StatusOK, "レシート解析API")
  })

  r.Run(":8080")
}