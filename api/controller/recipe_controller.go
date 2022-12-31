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

func RecipeDataGet(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()
  
  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // リクエストボディ取得
  type request struct {Offset int}
  requestData := request{}
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
  recipes := []model.Recipe{}
  err := tx.Model(&model.Recipe{}).Preload("Recipe_materials").Preload("Recipe_categories").Preload("Recipe_materials.Food").Limit(20).Offset(requestData.Offset).Find(&recipes).Error
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
    "data": recipes,
  })

  return
}

func MyRecipeDataCreate(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // リクエストボディ取得
  requestData := model.My_recipe{}
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

func MyRecipeDataUpdate(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // リクエストボディ取得
  requestData := model.My_recipe{}
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
  err := tx.Model(&model.My_recipe{}).Where("ID = ?", requestData.ID).Updates(requestData).Error
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

func MyRecipeDataDelete(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // リクエストボディ取得
  type request struct {RecipeID int}
  requestData := request{}
  requestErr := c.ShouldBindJSON(&requestData)
  if requestErr != nil {
    log.Print(requestErr)
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
    })
    return
  }

  // dbから削除
  tx := dbHandle.Begin()
  err := tx.Unscoped().Delete(&model.My_recipe{}, requestData.RecipeID).Error
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

func MyRecipeDataGet(c *gin.Context) {
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

  // dbから取得
  tx := dbHandle.Begin()
  myRecipes := []model.My_recipe{}
  err := tx.Model(&model.My_recipe{}).Preload("Recipe").Preload("Recipe.Recipe_materials").Preload("Recipe.Recipe_categories").Preload("Recipe.Recipe_materials.Food").Where("user_id = ?", requestData.UserID).Order("index").Find(&myRecipes).Error
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

func RecipeDataAdd(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // リクエストボディ取得
  type request struct {
    Image string
    Data model.Recipe
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
  // レシピ
  tx := dbHandle.Begin()
  RecipeErr := tx.Create(&requestData.Data).Error
  if RecipeErr != nil {
    log.Print(RecipeErr)
    tx.Rollback()
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
    })
    return
  }

  // レシピの材料
  for i := 0; i < len(requestData.Data.Recipe_materials); i++ {
    RecipeMaterialsErr := tx.Create(&requestData.Data.Recipe_materials[i]).Error
    if RecipeMaterialsErr != nil {
      log.Print(RecipeMaterialsErr)
      tx.Rollback()
      c.JSON(http.StatusBadRequest, gin.H{
        "success": false,
      })
      return
    }
  }

  // レシピのタグ
  for i :=0; i < len(requestData.Data.Recipe_categories); i++ {
    RecipeCategoriesErr := tx.Create(&requestData.Data.Recipe_categories[i]).Error
    if RecipeCategoriesErr != nil {
      log.Print(RecipeCategoriesErr)
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