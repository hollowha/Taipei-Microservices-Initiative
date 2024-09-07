package controllers

import (
	"backend/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Itinerary struct {
	Location    string `json:"location"`
	Description string `json:"description"`
}

func SubmitItinerary(c *gin.Context) {
	var itinerary Itinerary
	if err := c.ShouldBindJSON(&itinerary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 使用 gemini 模型來推薦地點
	jsonData, err := json.Marshal(itinerary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal itinerary"})
		return
	}

	// 假設 gemini 模型會返回推薦的地點
	recommendedLocation := models.SearchLocation(string(jsonData))

	// 回傳推薦地點給前端
	c.JSON(http.StatusOK, gin.H{
		"location":            itinerary.Location,
		"recommendedLocation": recommendedLocation, // 回傳推薦地點
	})
}
