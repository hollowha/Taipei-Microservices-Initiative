package controllers

import (
	"backend/models"
	"database/sql"
	"net/http"
	"path/filepath"

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
	query := `INSERT INTO activities VALUES (?, ?, ?, ?, ?)`
	err := db.Exec(query, event.Title, event.Description, event.Location, event.ImagePath, event.Time).Error

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

func ServeImage(c *gin.Context) {
	imageTitle := c.Param("imageTitle") // 從 URL 參數中取得圖片名稱

	var imagePath string

	// 使用 GORM 的 Raw 方法執行 SELECT 查詢，並將結果掃描到 imagePath 變數中
	err := db.Raw("SELECT image_path FROM activities WHERE title = ?", imageTitle).Scan(&imagePath).Error
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve image path"})
		}
		return
	}

	// 構建圖片的完整路徑
	fullImagePath := filepath.Join("Taipei-Microservices-Initiative/uploads", imagePath)

	// 檢查文件是否存在，並返回給前端
	c.File(fullImagePath)
}
