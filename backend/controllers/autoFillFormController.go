package controllers

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Form struct {
	Formname string `json:"formname"`
	Detail   string `json:"detail"`
}

func FillFormbyID(c *gin.Context) {

	formname := c.Param("formname")
	// Get the form file path based on the ID
	formFilePath := "./assets/formfile/" + formname + ".pdf"

	// Open the form file
	formFile, err := os.Open(formFilePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to open form file")
		return
	}
	defer formFile.Close()

	// Set the appropriate headers for the file download
	newFileName := "BE_" + formname + ".pdf"
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename="+newFileName)

	// Copy the file content to the response writer
	if _, err := io.Copy(c.Writer, formFile); err != nil {
		c.String(http.StatusInternalServerError, "Failed to send form file")
	}
}

func GetFormDetail(c *gin.Context) {
	formname := c.Param("formname")
	var form Form

	// Fetch the form details by formname
	if result := db.Where("formname = ?", formname).First(&form); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "form not found"})
		return
	}

	// Respond with the form details
	c.JSON(http.StatusOK, gin.H{
		"formname": form.Formname,
		"detail":   form.Detail,
	})
}
