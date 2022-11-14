package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/narikei-74/food_manage_app-backend/api/controller"
)

func main() {
  r := gin.Default()

  // ログインAPI
  r.POST("/user/login", controller.user_login)
  // 会員登録API
  r.POST("/user/register", controller.user_register)
  // ユーザー情報取得API
  r.POST("/user/info/get", controller.user_info_get)
  // ユーザー情報保存API
  r.POST("/user/info/save", controller.user_info_save)
  // 献立一覧取得API
  r.POST("/recipedata/get", controller.recipedata_get)
  // my献立保存API
  r.POST("/myrecipedata/save", controller.myrecipedata_save)
  // my献立取得API
  r.POST("/myrecipedata/get", controller.myrecipedata_get)
  // 献立自動作成条件取得API
  r.POST("/recipe_create_setting/get", controller.recipe_create_setting_get)
  // 献立自動作成条件保存API
  r.POST("/recipe_create_setting/save", controller.recipe_create_setting_save)
  // 残り食材取得API
  r.POST("/food_scock/get", controller.food_stock_get)
  // 残り食材保存API
  r.POST("/food_scock/save", controller.food_stock_save)
  // 食材市場データ取得API
  r.POST("/food_market/get", controller.food_market_get)
  // レシート解析API
  r.POST("/receipt_analysis", controller.receipt_analysis)

  r.Run(":8080")
}