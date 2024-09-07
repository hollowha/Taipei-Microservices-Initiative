package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Event struct {
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

// 獲取所有活動及其選項
func GetRecord() ([]Event, error) {
	var activities []Event

	err := db.Preload("Options").Find(&activities).Error
	if err != nil {
		return nil, err
	}

	return activities, nil
}

// 更新選項的投票數
func UpdateVote(optionID int) error {
	err := db.Model(&Option{}).Where("id = ?", optionID).Update("votes", gorm.Expr("votes + ?", 1)).Error
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
