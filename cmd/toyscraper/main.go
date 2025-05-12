package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	// "github.com/danmrichards/sandbox/toyscraper/internal/classifier"

	"github.com/danmrichards/sandbox/toyscraper/internal/classifier"
	"github.com/danmrichards/sandbox/toyscraper/internal/cleaner"
	"github.com/danmrichards/sandbox/toyscraper/internal/config"
	"github.com/danmrichards/sandbox/toyscraper/internal/converter"
	"github.com/danmrichards/sandbox/toyscraper/internal/extractor"
	"github.com/danmrichards/sandbox/toyscraper/internal/schema"
	"github.com/danmrichards/sandbox/toyscraper/internal/scraper"
)

func main() {
	var (
		url                string
		classifierModel    string
		classifierModelDir string
		classify           bool
		timeout            int
	)

	flag.StringVar(&url, "url", "", "URL to scrape")
	flag.BoolVar(&classify, "classify", false, "Classify the content")
	flag.StringVar(&classifierModel, "classifier-model", config.DefaultClassifierModel, "Classifier model to use")
	flag.StringVar(&classifierModelDir, "classifier-model-dir", config.DefaultClassifierModelDir, "Directory for classifier models")
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

	if classify {
		// TODO: A pre-processing step here to remove menus and other junk content
		// would help the classifier.

		// Classify the content.
		//
		// NOTE: we're using a pure-go classifier here, but there is nothing to stop
		// us from passing the markdown to some external model or service.
		zs, err := classifier.NewZeroShot(classifierModelDir, classifierModel)
		if err != nil {
			log.Fatalf("Failed to create zero-shot classifier: %v", err)
		}

		// TODO: this is a dirty hack to workaround the input token limit in the
		// cybertron package (1024 tokens for the default model).

		// Remove all empty lines and use half the content.
		shortenedMarkdown := strings.Replace(markdown, "\n\n", "\n", -1)
		shortenedMarkdown = strings.TrimSpace(shortenedMarkdown)

		classificationResult, err := zs.Classify(
			context.Background(),
			shortenedMarkdown[:len(shortenedMarkdown)/2],
			config.ClassificationLabels,
		)
		if err != nil {
			log.Fatalf("Failed to classify content: %v", err)
		}

		fmt.Printf("Classification result: %v\n", classificationResult)
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

	fmt.Println(extractedContent)
}
