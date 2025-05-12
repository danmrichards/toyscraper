package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/danmrichards/sandbox/toyscraper/internal/cleaner"
	"github.com/danmrichards/sandbox/toyscraper/internal/config"
	"github.com/danmrichards/sandbox/toyscraper/internal/converter"
	"github.com/danmrichards/sandbox/toyscraper/internal/extractor"
	"github.com/danmrichards/sandbox/toyscraper/internal/schema"
	"github.com/danmrichards/sandbox/toyscraper/internal/scraper"
)

func main() {
	var (
		url     string
		timeout int
	)

	flag.StringVar(&url, "url", "", "URL to scrape")
	flag.IntVar(&timeout, "timeout", config.DefaultTimeout, "Timeout in seconds")
	flag.Parse()

	if url == "" {
		log.Fatal("URL is required. Use -url flag to specify the URL to scrape.")
	}

	// Load the extractor API key from environment variables.
	geminiAPIKey := os.Getenv("GEMINI_API_KEY")
	if geminiAPIKey == "" {
		log.Fatal("EXTRACTOR_API_KEY environment variable is required.")
	}

	// Get the HTML content of the page
	content, err := scraper.GetHTML(url, timeout)
	if err != nil {
		log.Fatalf("Failed to scrape URL: %v", err)
	}

	// Clean HTML
	cleanedContent, err := cleaner.HTML(content)
	if err != nil {
		log.Fatalf("Failed to clean HTML: %v", err)
	}

	// Convert HTML to Markdown
	markdown, err := converter.ToMarkdown(cleanedContent)
	if err != nil {
		log.Fatalf("Failed to convert HTML to Markdown: %v", err)
	}

	ext, err := extractor.NewExtractor(context.Background(), geminiAPIKey)
	if err != nil {
		log.Fatalf("Failed to create extractor: %v", err)
	}

	// Use a JSON schema to structure the extracted content.
	jobSchema, err := schema.JSONSchemaString(schema.JobPosting{})
	if err != nil {
		log.Fatalf("Failed to get JSON schema: %v", err)
	}

	// Extract content.
	extractedContent, err := ext.ExtractContent(
		context.Background(),
		config.ExtractionModel,
		jobSchema,
		url,
		markdown,
	)
	if err != nil {
		log.Fatalf("Failed to extract content: %v", err)
	}
	// Print the extracted content
	fmt.Println(extractedContent)
}
