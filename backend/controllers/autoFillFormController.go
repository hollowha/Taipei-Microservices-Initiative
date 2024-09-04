package controllers

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

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
