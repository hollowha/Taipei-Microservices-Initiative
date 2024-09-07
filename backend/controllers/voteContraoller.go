package controllers

import (
	"backend/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"gorm.io/gorm"
)

type Event struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}
type EventCollection struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Options []Option `json:"options" gorm:"foreignKey:ActivityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Option struct {
	ID         int    `json:"id"`
	ActivityID int    `json:"activity_id"`
	Label      string `json:"label"`
	Votes      int    `json:"votes"`
}

type VoteRequest struct {
	OptionID int `json:"option_id"`
}

func (EventCollection) TableName() string {
	return "events" // 替換為實際的資料表名稱
}

// 獲取所有活動及其選項
func GetRecord() ([]EventCollection, error) {
	var activities []EventCollection

	err := db2.Preload("Options").Find(&activities).Error
	if err != nil {
		return nil, err
	}

	return activities, nil
}

// 更新選項的投票數
func UpdateVote(optionID int) error {
	err := db2.Model(&Option{}).Where("id = ?", optionID).Update("votes", gorm.Expr("votes + ?", 1)).Error
	return err
}

/*
----------------------------------------------------------------------
以下是controller
*/

// GetVoteActivities 返回所有投票活動及其選項
func GetVoteActivities(c *gin.Context) {
	activities, err := GetRecord()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activities)
}

// SubmitVote 接收投票請求並更新選項的投票數
func SubmitVote(c *gin.Context) {
	var voteReq VoteRequest
	if err := c.ShouldBindJSON(&voteReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := UpdateVote(voteReq.OptionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vote submitted successfully"})
}

func CreateVoteActivities(c *gin.Context) {
	var eventCollect EventCollection
	if err := c.ShouldBindJSON(&eventCollect); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無效的輸入資料"})
		return
	}
	// 判斷是否要新增
	jsonEvent, err := json.Marshal(eventCollect)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal report"})
		return
	}
	insertOrNot, err := CheckEvent(string(jsonEvent))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check"})
		return
	}
	fmt.Println("Insert or not: ", insertOrNot)
	if !insertOrNot {
		c.JSON(http.StatusOK, gin.H{
			"message": "活動未審核通過",
		})
		return
	}

	var event Event
	event.ID = 0
	event.Name = eventCollect.Name
	err = db2.Create(&event).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// query := `INSERT INTO options (activity_id, label)`
	for _, option := range eventCollect.Options {
		option.ActivityID = event.ID // 設定關聯的 event_id
		if err := db2.Create(&option).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert option"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "活動已成功新增",
	})
}
func CheckEvent(eventString string) (bool, error) {
	ctx := context.Background()

	// 确保 gemini client 已经初始化
	if models.GeminiClient == nil {
		models.InitGeminiClient()
	}

	model := models.GeminiClient.GenerativeModel("gemini-1.5-flash")

	model.SetTemperature(1)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(`你將作為審查員，審核使用者傳入的內容是否合理。
		輸入的內容如果違背善良風俗，或無意義等，就將它過濾掉。如果內容合理，輸出true，反之則輸出false。並解釋原因`)},
	}

	session := model.StartChat()
	fmt.Println("prompt", eventString)
	resp, err := session.SendMessage(ctx, genai.Text(eventString))
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
		return false, fmt.Errorf("error when sending event")
	}
	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Printf("%v\n", part)
		str := fmt.Sprintf("%v", part)
		if strings.Contains(str, "true") {
			return true, nil
		} else if strings.Contains(str, "false") {
			return false, nil
		}
		// return false, fmt.Errorf("no response")
	}
	return false, fmt.Errorf("error when checking content")
}
