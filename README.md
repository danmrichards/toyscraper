# Toyscraper

Toyscraper is a command-line web scraping tool that extracts and processes web content with minimal noise. It offers HTML cleaning, Markdown conversion, and optional AI-powered content extraction.

## Features

- **Web Scraping**: Headless browser-based web scraping using Rod (Golang equivalent of Playwright)
- **HTML Cleaning**: Removes unwanted elements, attributes, and comments from HTML content
- **Markdown Conversion**: Converts cleaned HTML to Markdown for better readability
- **AI Content Extraction**: Uses Google's Gemini AI model to extract structured information (optional)
- **JSON Output**: Option to output extracted content in JSON format

## Installation

### Prerequisites

- Go 1.16 or higher
- Gemini API Key (only for AI extraction feature)

### Building from Source

1. Clone the repository:

   ```bash
   git clone https://github.com/danmrichards/sandbox/toyscraper.git
   cd toyscraper
   ```

2. Build the binary:

   ```bash
   go build ./cmd/toyscraper
   ```

3. (Optional) If you want to use the AI extraction feature, set your Gemini API key as an environment variable:
   ```bash
   export GEMINI_API_KEY="your-api-key-here"
   ```

## Usage

### Basic Usage

To scrape a web page and output its cleaned content as Markdown:

```bash
./toyscraper -url="https://example.com"
```

### Available Flags

- `-url`: (Required) URL to scrape
- `-timeout`: (Optional) Timeout in seconds (default: 30)

### Examples

1. Basic scraping:

   ```bash
   ./toyscraper -url="https://example.com"
   ```

2. Scraping with a custom timeout:

   ```bash
   ./toyscraper -url="https://example.com" -timeout=60
   ```

## Project Structure

```
toyscraper/
├── cmd/
│   └── toyscraper/       # Command-line application
│       └── main.go
├── internal/
│   ├── cleaner/          # HTML cleaning functionality
│   ├── config/           # Application configuration
│   ├── converter/        # HTML to Markdown conversion
│   ├── extractor/        # AI-powered content extraction
│   └── scraper/          # Web scraping functionality
```

## Roadmap

- [ ] Implement an AI classifier to determine if the given content matches given labels
- [ ] Add support for deep crawling a webpage and discovering nested content
- [ ] Add support for additional AI models beyond Gemini
- [ ] Implement unit tests for all internal packages
- [ ] Add support for scraping multiple URLs in parallel
- [ ] Improve error handling and logging
- [ ] Create a Dockerfile for easier deployment
- [ ] Write comprehensive user documentation
