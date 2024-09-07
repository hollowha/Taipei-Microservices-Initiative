package controllers

import (
	"backend/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Report struct {
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Time        string `json:"time"`
	ID          int    `json:"id"`
}

func CreateReport(c *gin.Context) {
	var report Report
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// do query
	query := `INSERT INTO reports (title, type, description, time) VALUES (?, ?, ?, ?)`
	err := db.Exec(query, report.Title, report.Type, report.Description, report.Time).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert report"})
		return
	}

	jsonData, err := json.Marshal(report)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal report"})
		return
	}
	models.ReportClassify(string(jsonData))

	c.JSON(http.StatusOK, report)
}

func GetAllReports(c *gin.Context) {
	var reports []Report
	query := `SELECT * FROM reports`
	err := db.Raw(query).Scan(&reports).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve reports"})
		return
	}
	c.JSON(http.StatusOK, reports)
}
func DeleteReportByID(c *gin.Context) {
	id := c.Param("id")
	query := `DELETE FROM reports WHERE id = ?`
	err := db.Exec(query, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete report"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Report deleted successfully"})
}
