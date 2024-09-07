package models

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
)

func SearchLocation(data string) string {
	ctx := context.Background()

	// 确保 gemini client 已经初始化
	if GeminiClient == nil {
		InitGeminiClient()
	}

	model := GeminiClient.GenerativeModel("gemini-1.5-flash")

	model.SetTemperature(1)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(`你將做為一個導遊，為使用者推薦地點。使用者會給你一個地點的名稱，以及行程的描述，你需要為他們推薦該地點附近、符合行程描述的一個地點的名稱，可能是餐廳、景點等等`)},
	}

	session := model.StartChat()

	resp, err := session.SendMessage(ctx, genai.Text(data))
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		str := fmt.Sprintf("%v", part)
		return str
	}
	return "" // Add a return statement here
}
