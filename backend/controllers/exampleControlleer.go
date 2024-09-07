package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Example 回應簡單的 JSON 資料
func Example(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, this is an example! 1130 version",
	})
}
