package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Event struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Time        time.Time `json:"time"`
	ImageURL    string    `json:"image_url"`
}

// GetAllActivitys 取得所有活動
func GetAllActivitys(c *gin.Context) {
	c.JSON(http.StatusOK, "to be done......")
}
