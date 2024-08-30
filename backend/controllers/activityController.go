package controllers

import (
	"backend/models"
	"database/sql"
	"fmt"
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
	err := db.Exec(query, event.Title, event.Description, event.Location, event.ImageName, event.Time).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	// 返回成功訊息和插入的活動數據
	c.JSON(http.StatusOK, event)
}

func GetEvents(c *gin.Context) {
	title := c.Param("title") // 從查詢參數中取得 title

	var events []models.Activity

	// 根據 title 查詢資料庫中的 activities 表
	if err := db.Table("activities").Where("title = ?", title).Scan(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}

	// 如果沒有找到對應的活動資料，返回 404
	if len(events) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No events found with the given title"})
		return
	}
	fmt.Println(events)
	// 返回找到的活動資料
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

	var image_name string

	// 使用 GORM 的 Raw 方法執行 SELECT 查詢，並將結果掃描到 imagePath 變數中
	err := db.Raw("SELECT image_path FROM activities WHERE title = ?", imageTitle).Scan(&image_name).Error
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve image path"})
		}
		return
	}

	// 構建圖片的完整路徑
	fullImagePath := filepath.Join("uploads", image_name)

	// 檢查文件是否存在，並返回給前端
	c.File(fullImagePath)
}
