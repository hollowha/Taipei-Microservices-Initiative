package controllers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateEvent handles creating a new event
func CreateEvent(c *gin.Context) {
	var event models.Activity

	// 將請求的 JSON 數據綁定到 event 變數
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 使用 SQL 查詢插入數據
	query := `INSERT INTO Activity VALUES (?, ?, ?, ?, ?)`
	err := db.Exec(query, event.Title, event.Description, event.Location, event.Time, event.ImageURL).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	// 返回成功訊息和插入的活動數據
	c.JSON(http.StatusOK, event)
}

// GetEvents handles retrieving all events
func GetEvents(c *gin.Context) {
	var events []models.Activity

	if err := db.Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}

	c.JSON(http.StatusOK, events)
}

// GetAllActivitys 取得所有活動
func GetAllActivitys(c *gin.Context) {
	var activities []models.Activity

	// Retrieve all activities from the database
	if err := db.Find(&activities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve activities"})
		return
	}

	// Return the activities in JSON format
	c.JSON(http.StatusOK, activities)
}
