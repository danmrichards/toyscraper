// Package scraper provides functionality to scrape web content using browser automation
package scraper

import (
	"fmt"
	"time"

	"github.com/danmrichards/sandbox/toyscraper/internal/config"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// GetHTML fetches the HTML content of a specified URL
func GetHTML(url string, timeoutSeconds int) (string, error) {
	// Validate timeout value
	if timeoutSeconds <= 0 {
		timeoutSeconds = config.DefaultTimeout
	} else if timeoutSeconds > config.MaxTimeout {
		timeoutSeconds = config.MaxTimeout
	}

	// Create a new browser launcher
	l := launcher.New().Headless(true)

	// Launch a new browser
	browser := rod.New().ControlURL(l.MustLaunch()).MustConnect()
	defer browser.MustClose()

	// Create a new page
	page := browser.MustPage(url)

	// Set viewport
	page.MustSetViewport(config.DefaultViewportWidth, config.DefaultViewportHeight, 1, false)

	// Set timeout
	page.Timeout(time.Duration(timeoutSeconds) * time.Second)

	// Wait for the page to load
	page.MustWaitLoad()

	// Get the HTML content of the entire page
	content, err := page.HTML()
	if err != nil {
		return "", fmt.Errorf("failed to get page content: %v", err)
	}

	return content, nil
}
