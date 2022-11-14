package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func food_scock_get(c *gin.Context) {
  c.String(http.StatusOK, "食材市場データ取得API")
}