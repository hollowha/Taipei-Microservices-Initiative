package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllActivitys 取得所有活動
func GetAllActivitys(c *gin.Context) {
	c.JSON(http.StatusOK, "to be done......")
}
