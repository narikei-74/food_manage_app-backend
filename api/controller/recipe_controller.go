package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func recipedata_get(c *gin.Context) {
  c.String(http.StatusOK, "献立一覧取得API")
}

func myrecipedata_save(c *gin.Context) {
  c.String(http.StatusOK, "my献立保存API")
}

func myrecipedata_get(c *gin.Context) {
  c.String(http.StatusOK, "my献立取得API")
}

func recipe_create_setting_get(c *gin.Context) {
  c.String(http.StatusOK, "献立自動作成条件取得API")
}

func recipe_create_setting_save(c *gin.Context) {
  c.String(http.StatusOK, "献立自動作成条件保存API")
}