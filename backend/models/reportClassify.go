package models

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func ReportClassify(data string) {
	ctx := context.Background()

	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		log.Fatalf("API_KEY not found")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	model.SetTemperature(1)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(`你將負責幫助輸入的資訊進行分類，輸入的資訊將是json格式的內容，
		包含title、type、description等，\ntype會有兩種可能：問題、建議。針對問題，你需要依據緊急、重大程度，
		極可能造成的危害、影響去判斷是否要優先處理，輸出將會是1~3，3代表最緊急、應該要被立即處理的反之則是1，
		2代表在兩者之間。\n如果type是建議，你可以分類為1或2，你可以通過判斷提交這項建議的使用者是否是認真的、或是根據內容決定，
		如果只是一些沒有意義的、或者是說開玩笑的內容則歸類為1，如果是認真的改善建議，則分類為2\n根據傳入的資料，
		你只需要回傳一個json格式的輸出，範例如下：\n{ type: \"建議\", class: 1 }\n{ type: \"問題\", class: 3}`)},
	}

	session := model.StartChat()

	resp, err := session.SendMessage(ctx, genai.Text("INSERT_INPUT_HERE"))
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	for _, part := range resp.Candidates[0].Content.Parts {
		fmt.Printf("%v\n", part)
	}
}
