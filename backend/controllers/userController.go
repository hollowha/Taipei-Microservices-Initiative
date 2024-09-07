package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// User struct 模擬一個 User 資料結構
type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Imgurl   string `json:"imgurl"`
	IDNumber string `json:"idnumber"`
}

// 全域變數保存資料庫連接
var db *gorm.DB

// SetDB 設定資料庫實例
func SetDB(database *gorm.DB) {
	db = database
}

// var users = []User{
// 	{ID: 1, Name: "John Doe", Age: 25},
// 	{ID: 2, Name: "Jane Doe", Age: 30},
// }

// GetAllUsers 取得所有使用者
func GetAllUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GetUserByID 根據 ID 取得單一使用者
func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var user User
	if result := db.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByID 根據 ID 取得單一使用者
func GetUserIMGByFileName(c *gin.Context) {
	filename := c.Param("filename")
	// get the image file path based on the filename
	imgFilePath := "./uploads/" + filename
	fmt.Printf("imgFilePath: %s\n", imgFilePath)

	// Open the image file
	imgFile, err := os.Open(imgFilePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to open image file")
		return
	}
	defer imgFile.Close()

	// Set the appropriate headers for the file download
	c.Header("Content-Type", "image/jpeg")
	c.Header("Content-Disposition", "attachment; filename="+filename)

	// Copy the file content to the response writer
	if _, err := io.Copy(c.Writer, imgFile); err != nil {
		c.String(http.StatusInternalServerError, "Failed to send image file")
	}

	// c.JSON(http.StatusOK, user)

}

// CreateUser 建立新使用者
func CreateUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&newUser)
	c.JSON(http.StatusCreated, newUser)
}

// UpdateUserByID 根據 ID 更新使用者資料
func UpdateUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	fmt.Println(id)
	// print content of request body
	fmt.Println(c.Request.Body)

	var user User
	if result := db.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	fmt.Println(user)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUserByID 根據 ID 刪除使用者
func DeleteUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var user User
	if result := db.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
