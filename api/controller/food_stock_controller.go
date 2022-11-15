package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func FoodStockGet(c *gin.Context) {
  c.String(http.StatusOK, "残り食材取得API")
}

func FoodStockSave(c *gin.Context) {
  c.String(http.StatusOK, "残り食材保存API")
}

func ReceiptAnalysis(c *gin.Context) {
  c.String(http.StatusOK, "レシート解析API")
}