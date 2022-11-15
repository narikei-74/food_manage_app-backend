package main

import (
  "github.com/gin-gonic/gin"
  "github.com/narikei-74/food_manage_app-backend/api/controller"
)

func main() {
  r := gin.Default()

  // ログインAPI
  r.POST("/user/login", controller.UserLogin)
  // 会員登録API
  r.POST("/user/register", controller.UserRegister)
  // ユーザー情報取得API
  r.POST("/user/info/get", controller.UserInfoGet)
  // ユーザー情報保存API
  r.POST("/user/info/save", controller.UserInfoSave)
  // 献立一覧取得API
  r.POST("/recipedata/get", controller.RecipeDataGet)
  // my献立保存API
  r.POST("/myrecipedata/save", controller.MyRecipeDataSave)
  // my献立取得API
  r.POST("/myrecipedata/get", controller.MyRecipeDataGet)
  // 献立自動作成条件取得API
  r.POST("/recipe_create_setting/get", controller.RecipeCreateSettingGet)
  // 献立自動作成条件保存API
  r.POST("/recipe_create_setting/save", controller.RecipeCreateSettingSave)
  // 残り食材取得API
  r.POST("/food_stock/get", controller.FoodStockGet)
  // 残り食材保存API
  r.POST("/food_stock/save", controller.FoodStockSave)
  // 食材市場データ取得API
  r.POST("/food_market/get", controller.FoodMarketGet)
  // レシート解析API
  r.POST("/receipt_analysis", controller.ReceiptAnalysis)

  r.Run(":8080")
}