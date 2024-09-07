package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Example 回應簡單的 JSON 資料
func Example(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, this is an example??!\nfrom the air",
		"message": "Hello, this is an example! XDDDDDDD",
	})
}
