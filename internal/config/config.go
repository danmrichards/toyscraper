// Package config provides configuration constants for the application
package config

// Default timeout values in seconds
const (
	// DefaultTimeout is the default timeout for scraping operations
	DefaultTimeout = 30

	// MaxTimeout is the maximum allowed timeout for scraping operations
	MaxTimeout = 120
)

// Browser configurations
const (
	// DefaultViewportWidth is the default viewport width for the browser
	DefaultViewportWidth = 1920

	// DefaultViewportHeight is the default viewport height for the browser
	DefaultViewportHeight = 1080
)

// HTML and Markdown configurations
const (
	// MaxContentLength is the maximum allowed length of HTML content to process
	MaxContentLength = 10 * 1024 * 1024 // 10MB
)

// Content extraction configurations.
const (
	// ExtractionModel is the model used for content extraction.
	ExtractionModel = "gemini-2.5-pro-exp-03-25"

	// ExtractionPrompt are the instructions for the content extraction model.
	ExtractionPrompt = `Here is the content from the URL:
<url>%s</url>

<url_content>
%s
</url_content>

The user has made the following request for what information to extract from the above content:

<user_request>
Extract job posting details, using markdown structure to:
1. Identify requirement priorities from headings and subheadings
2. Extract contact info from the page footer or dedicated contact section
3. Parse salary information from specially formatted elements if available
4. Determine application deadline from timestamp or date elements

Use markdown format and structure to enhance extraction accuracy.
</user_request>

<schema_block>
%s
</schema_block>

Please carefully read the URL content and the user's request. If the user provided a desired JSON schema in the <schema_block> above, extract the requested information from the URL content according to that schema. If no schema was provided, infer an appropriate JSON schema based on the user's request that will best capture the key information they are looking for.

Extraction instructions:
Return the extracted information as a list of JSON objects, with each object in the list corresponding to a block of content from the URL, in the same order as it appears on the page.

Quality Reflection:
Before outputting your final answer, double check that the JSON you are returning is complete, containing all the information requested by the user, and is valid JSON that could be parsed by json.loads() with no errors or omissions. The outputted JSON objects should fully match the schema, either provided or inferred.

Quality Score:
After reflecting, score the quality and completeness of the JSON data you are about to return on a scale of 1 to 5. Write the score inside <score> tags.

Avoid Common Mistakes:
- Do NOT add any comments using "//" or "#" in the JSON output. It causes parsing errors.
- Make sure the JSON is properly formatted with curly braces, square brackets, and commas in the right places.
- Do not miss closing </blocks> tag at the end of the JSON output.
- Do not generate the Python code show me how to do the task, this is your task to extract the information and return it in JSON format.

Result
Output the final list of JSON objects. Make sure to close the tag properly.`
)

// List of unwanted element tags to remove during cleaning
var UnwantedElements = map[string]bool{
	"script":   true,
	"style":    true,
	"noscript": true,
	"iframe":   true,
	"svg":      true,
	"canvas":   true,
	"button":   true,
	"meta":     true,
	"link":     true,
	"head":     true,
	"path":     true,
}

// List of HTML attributes to keep during cleaning
var KeepAttributes = map[string]bool{
	"alt":   true,
	"title": true,
	"href":  true,
	"src":   true,
}
