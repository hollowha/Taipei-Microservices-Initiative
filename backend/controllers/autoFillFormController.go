package controllers

import (
	"fmt"
	"io"
	"io/ioutil"
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

func SavePDF(c *gin.Context) {
	formname := c.Param("formname")
	fmt.Println("formname", formname)
	// Get the form file path based on the ID
	filepath := "./assets/filledPdf/" + formname + ".pdf"

	if formname == "" {
		fmt.Println("formname is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Form name is required"})
		return
	}

	// Read the raw PDF bytes from the request body
	pdfBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	outFile, err := os.Create(filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
		return
	}
	defer outFile.Close()

	// Write the raw PDF bytes to the file
	_, err = outFile.Write(pdfBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file"})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{"message": "File saved successfully"})
}

func ShowForm(c *gin.Context) {
	fmt.Println("ShowForm")
	formname := c.Param("formname")
	c.File("./assets/filledPdf/" + formname)
}
