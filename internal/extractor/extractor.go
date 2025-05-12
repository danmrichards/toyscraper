package extractor

import (
	"context"
	"fmt"

	"github.com/danmrichards/sandbox/toyscraper/internal/config"
	"google.golang.org/genai"
)

// Extractor is a struct that holds the GenAI client.
type Extractor struct {
	client *genai.Client
}

// NewExtractor creates a new Extractor instance with the provided API key.
func NewExtractor(ctx context.Context, apiKey string) (*Extractor, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create GenAI client: %w", err)
	}

	return &Extractor{client: client}, nil
}

// ExtractContent extracts content from the provided raw content using the GenAI client.
func (e *Extractor) ExtractContent(ctx context.Context, model, schema, url, rawContent string) (string, error) {
	result, err := e.client.Models.GenerateContent(
		ctx,
		model,
		genai.Text(fmt.Sprintf(config.ExtractionPrompt, url, rawContent, schema)),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	return result.Text(), nil
}
