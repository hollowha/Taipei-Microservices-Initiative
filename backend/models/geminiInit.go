package models

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var GeminiClient *genai.Client

func InitGeminiClient() {

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyAsrQ5ObzfFJfGjwAIaMcRxp4gW-PMiELg"))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	GeminiClient = client
}
