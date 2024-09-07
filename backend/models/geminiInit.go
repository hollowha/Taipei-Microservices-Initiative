package models

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var geminiClient *genai.Client

func initGeminiClient() {
	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		log.Fatalf("API_KEY not found")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	geminiClient = client
}
