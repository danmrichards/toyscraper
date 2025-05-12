// Package cleaner provides functionality to sanitize and optimize HTML for content extraction
package cleaner

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/danmrichards/sandbox/toyscraper/internal/config"
	"golang.org/x/net/html"
)

// HTML sanitizes and optimizes HTML for content extraction
func HTML(rawHTML string) (string, error) {
	// Check content length
	if len(rawHTML) > config.MaxContentLength {
		return "", fmt.Errorf("HTML content exceeds maximum allowed length (%d bytes)", config.MaxContentLength)
	}

	// Parse HTML
	doc, err := html.Parse(strings.NewReader(rawHTML))
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %v", err)
	}

	// Clean the HTML tree
	cleanNode(doc)

	// Remove empty elements
	removeEmptyNodes(doc)

	// Render the cleaned HTML
	var buf bytes.Buffer
	if err := html.Render(&buf, doc); err != nil {
		return "", fmt.Errorf("failed to render HTML: %v", err)
	}

	// Remove empty lines and normalize whitespace
	return removeEmptyLines(buf.String()), nil
}

// removeEmptyLines removes consecutive empty lines from HTML content
func removeEmptyLines(content string) string {
	// Split content into lines
	lines := strings.Split(content, "\n")

	var result []string
	var prevLineEmpty bool

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		isLineEmpty := trimmedLine == ""

		// Skip consecutive empty lines
		if isLineEmpty && prevLineEmpty {
			continue
		}

		result = append(result, line)
		prevLineEmpty = isLineEmpty
	}

	return strings.Join(result, "\n")
}

// cleanNode recursively cleans an HTML node and its children
func cleanNode(n *html.Node) {
	// Process children first (before potentially removing them)
	var next *html.Node
	for c := n.FirstChild; c != nil; c = next {
		next = c.NextSibling
		cleanNode(c)
	}

	// Remove comment nodes
	if n.Type == html.CommentNode {
		removeNode(n)
		return
	}

	// Skip if not an element node
	if n.Type != html.ElementNode {
		return
	}

	// Check if this is an unwanted element
	if config.UnwantedElements[n.Data] {
		// If this is the root node, just remove all its children
		if n.Parent == nil {
			n.FirstChild = nil
			n.LastChild = nil
		} else {
			// Remove this node from its parent
			removeNode(n)
		}
		return
	}

	// Process images and other media
	if n.Data == "img" {
		processImageNode(n)
	}

	// Remove unwanted attributes
	cleanAttributes(n)
}

// removeNode removes a node from its parent
func removeNode(n *html.Node) {
	if n.Parent == nil {
		return
	}

	if n.Parent.FirstChild == n {
		n.Parent.FirstChild = n.NextSibling
	}
	if n.Parent.LastChild == n {
		n.Parent.LastChild = n.PrevSibling
	}
	if n.PrevSibling != nil {
		n.PrevSibling.NextSibling = n.NextSibling
	}
	if n.NextSibling != nil {
		n.NextSibling.PrevSibling = n.PrevSibling
	}
}

// cleanAttributes removes unwanted attributes from a node
func cleanAttributes(n *html.Node) {
	// Build a new attribute list with only the attributes we want to keep
	var newAttrs []html.Attribute
	for _, attr := range n.Attr {
		if config.KeepAttributes[attr.Key] {
			newAttrs = append(newAttrs, attr)
		}
	}
	n.Attr = newAttrs
}

// processImageNode processes an image node, replacing it with its alt text if available
func processImageNode(n *html.Node) {
	var alt string
	for _, attr := range n.Attr {
		if attr.Key == "alt" && attr.Val != "" {
			alt = attr.Val
			break
		}
	}

	// If there's alt text, keep the image but ensure it's marked
	if alt != "" {
		// Keep the alt attribute as is
		return
	}

	// No useful alt text, remove the image
	removeNode(n)
}

// removeEmptyNodes removes nodes with no content
func removeEmptyNodes(n *html.Node) bool {
	if n.Type == html.TextNode {
		return len(strings.TrimSpace(n.Data)) > 0
	}

	if n.Type != html.ElementNode {
		return false
	}

	// Elements that shouldn't be removed even if empty
	if n.Data == "br" || n.Data == "hr" || n.Data == "img" {
		return true
	}

	hasContent := false
	var next *html.Node
	for c := n.FirstChild; c != nil; c = next {
		next = c.NextSibling
		if removeEmptyNodes(c) {
			hasContent = true
		} else {
			removeNode(c)
		}
	}

	return hasContent
}
