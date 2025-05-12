// Package converter provides functionality to convert HTML content to other formats
package converter

import (
	"fmt"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
)

// ToMarkdown converts HTML content to Markdown format
func ToMarkdown(htmlContent string) (string, error) {
	// Create a new converter
	converter := md.NewConverter("", true, nil)

	// Add GitHub-flavored Markdown plugins
	converter.Use(plugin.GitHubFlavored())

	// Convert HTML to Markdown
	markdown, err := converter.ConvertString(htmlContent)
	if err != nil {
		return "", fmt.Errorf("failed to convert HTML to Markdown: %v", err)
	}

	return markdown, nil
}
