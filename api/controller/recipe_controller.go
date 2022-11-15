package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func RecipeDataGet(c *gin.Context) {
  c.String(http.StatusOK, "献立一覧取得API")
}

func MyRecipeDataSave(c *gin.Context) {
  c.String(http.StatusOK, "my献立保存API")
}

func MyRecipeDataGet(c *gin.Context) {
  c.String(http.StatusOK, "my献立取得API")
}

func RecipeCreateSettingGet(c *gin.Context) {
  c.String(http.StatusOK, "献立自動作成条件取得API")
}

func RecipeCreateSettingSave(c *gin.Context) {
  c.String(http.StatusOK, "献立自動作成条件保存API")
}