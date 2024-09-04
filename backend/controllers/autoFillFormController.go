package controllers

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func FillFormbyID(c *gin.Context) {
	// Get the form file path based on the ID
	formFilePath := "./formfile/simpleform.pdf"

	// Open the form file
	formFile, err := os.Open(formFilePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to open form file")
		return
	}
	defer formFile.Close()

	// Set the appropriate headers for the file download
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=simpleform.pdf")

	// Copy the file content to the response writer
	if _, err := io.Copy(c.Writer, formFile); err != nil {
		c.String(http.StatusInternalServerError, "Failed to send form file")
	}
}
