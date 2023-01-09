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
  type searchInfo struct {
    RecipeName string
    Material string
    Category int
    Free string
  }
  type request struct {
    Offset int
    SearchInfo searchInfo
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

  // dbから取得
  tx := dbHandle.Begin()
  recipes := []model.Recipe{}
  query := tx.Model(&model.Recipe{}).Preload("Recipe_materials").Preload("Recipe_categories").Preload("Recipe_materials.Food").Limit(20).Offset(requestData.Offset)
  // レシピ名検索
  if (requestData.SearchInfo.RecipeName != "") {
    query = query.Where("name LIKE ?", "%"+requestData.SearchInfo.RecipeName+"%")
  }

  // レシピカテゴリ検索
  if (requestData.SearchInfo.Category != 0) {
    query = query.Where("dish_category = ?", requestData.SearchInfo.Category)
  }

  // 材料検索
  if (requestData.SearchInfo.Material != "") {
    query = query.Joins("JOIN recipe_materials ON recipe_materials.recipe_id = recipes.id").Joins("JOIN foods ON foods.id = recipe_materials.food_id").Where("foods.name LIKE ? OR foods.hiragana_name LIKE ?", "%"+requestData.SearchInfo.Material+"%", "%"+requestData.SearchInfo.Material+"%")
  }

  // レシピタグ検索
  if (requestData.SearchInfo.Free != "") {
    query = query.Joins("JOIN recipe_categories ON recipe_categories.recipe_id = recipes.id").Where("recipe_categories.category_name LIKE ? OR recipe_categories.hiragana_name LIKE ?", "%"+requestData.SearchInfo.Free+"%", "%"+requestData.SearchInfo.Free+"%")
  }

  query.Limit(20).Offset(requestData.Offset)
  err := query.Find(&recipes).Error
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

func MyRecipeDataPeopleNumUpdate(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // リクエストボディ取得
  type data struct {
    MyRecipeID int
    PeopleNum int
  }
  type request struct {
    Data []data
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

  // dbを更新
  tx := dbHandle.Begin()
  for i := 0; i < len(requestData.Data); i++ {
    err := tx.Model(&model.My_recipe{}).Where("ID = ?", requestData.Data[i].MyRecipeID).Update("people_num", requestData.Data[i].PeopleNum).Error;
    if err != nil {
      log.Print(err)
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

func RecipeDataAdd(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // リクエストボディ取得
  requestData := model.Recipe{}
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
  RecipeErr := tx.Create(&requestData).Error
  if RecipeErr != nil {
    log.Print(RecipeErr)
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

func AutoCreateRecipeSettingsGet(c *gin.Context) {
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
  settings := []model.Auto_create_recipe_settings{}
  err := tx.Model(&model.Auto_create_recipe_settings{}).Where("user_id = ?", requestData.UserID).Limit(1).Find(&settings).Error
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
    "data": settings,
  })

  return
}

func AutoCreateRecipeSettingsSave(c *gin.Context) {
  // db接続
  dbHandle := db.Init()
  defer dbHandle.Close()

  // ログオープン
  file := logFile.LogStart()
  defer file.Close()
  log.SetOutput(file)

  // リクエストボディ取得
  requestData := model.Auto_create_recipe_settings{}
  requestErr := c.ShouldBindJSON(&requestData)
  if requestErr != nil {
    log.Print(requestErr)
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
    })
    return
  }

  tx := dbHandle.Begin()
  settings := []model.Auto_create_recipe_settings{}
  err := tx.Model(&model.Auto_create_recipe_settings{}).Where("user_id = ?", requestData.UserID).Find(&settings).Error
  if err != nil {
    log.Print(err)
    tx.Rollback()
    c.JSON(http.StatusBadRequest, gin.H{
      "success": false,
    })
    return
  }

  // dbに保存
  if (len(settings) == 0) {
    // 追加処理
    RecipeErr := tx.Create(&requestData).Error
    if RecipeErr != nil {
      log.Print(RecipeErr)
      tx.Rollback()
      c.JSON(http.StatusBadRequest, gin.H{
        "success": false,
      })
      return
    }
  } else {
    // 更新処理
    err := tx.Model(&model.Auto_create_recipe_settings{}).Where("UserID = ?", requestData.UserID).Updates(requestData).Error
    if err != nil {
      log.Print(err)
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