package models

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var GeminiClient *genai.Client

func InitGeminiClient() {
	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		log.Fatalf("API_KEY not found")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	GeminiClient = client
}
