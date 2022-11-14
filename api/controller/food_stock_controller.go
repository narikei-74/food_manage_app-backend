package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func food_stock_get(c *gin.Context) {
  c.String(http.StatusOK, "残り食材取得API")
}

func food_stock_save(c *gin.Context) {
  c.String(http.StatusOK, "残り食材保存API")
}

func receipt_analysis(c *gin.Context) {
  c.String(http.StatusOK, "レシート解析API")
}