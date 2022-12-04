package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/narikei-74/food_manage_app-backend/api/db"
  "github.com/narikei-74/food_manage_app-backend/api/model"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func RecipeDataGet(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // リクエストボディ取得
  type request struct {Offset int}
  requestData := request{}
  err := c.ShouldBindJSON(&requestData)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": err,
    })
    return
  }

  // レシピデータ取得
  recipes := []model.Recipe{}
  recipesGetErr := dbHandle.Model(&model.Recipe{}).Preload("Recipe_materials").Preload("Recipe_categories").Preload("Recipe_materials.Food").Limit(20).Offset(requestData.Offset).Find(&recipes).Error

  // エラーの場合
  if (recipesGetErr != nil) {
    c.JSON(http.StatusOK, gin.H{
      "success": false,
      "error": recipesGetErr,
    })

    return
  }

  c.JSON(http.StatusOK, gin.H{
    "success": true,
    "data": recipes,
  })

  return
}

func MyRecipeDataSave(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // リクエストボディ取得
  myRecipe := model.My_recipe{}
  err := c.ShouldBindJSON(&myRecipe)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": err,
    })
    return
  }

  // dbに保存
  result := dbHandle.Create(&myRecipe)
  if result.Error != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
    })
    return
  }

  // レスポンス
  c.JSON(http.StatusOK, gin.H{
    "success": true,
  })

  return
}

func MyRecipeDataGet(c *gin.Context) {
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

  // dbから取得
  myRecipes := []model.My_recipe{}
  result := dbHandle.Model(&model.My_recipe{}).Preload("Recipe").Preload("Recipe.Recipe_materials").Preload("Recipe.Recipe_categories").Where("user_id = ?", requestData.UserID).Order("index").Find(&myRecipes)
  if result.Error != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
      "error": result.Error,
    })
    return
  }

  // レスポンス
  c.JSON(http.StatusOK, gin.H{
    "success": true,
    "data": myRecipes,
  })

  return
}

func RecipeCreateSettingGet(c *gin.Context) {
  c.String(http.StatusOK, "献立自動作成条件取得API")
}

func RecipeCreateSettingSave(c *gin.Context) {
  c.String(http.StatusOK, "献立自動作成条件保存API")
}