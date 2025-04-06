package schemaorg

import (
	"fmt"
	"html/template"

	"github.com/a-h/templ"
	"github.com/indaco/teseo"
)

// WebPage represents a Schema.org WebPage object.
// For more details about the meaning of the properties see: https://schema.org/WebPage
//
// Example usage:
//
// Pure struct usage:
//
//	webpage := &schemaorg.WebPage{
//		URL:         "https://www.example.com",
//		Name:        "Example WebPage",
//		Headline:    "Welcome to Example WebPage",
//		Description: "This is an example webpage.",
//		About:       "Something related to the home page",
//		Keywords:    "example, webpage, demo",
//		InLanguage:  "en",
//	}
//
// Factory method usage:
//
//	webpage := schemaorg.NewWebPage(
//		"https://www.example.com",
//		"Example WebPage",
//		"Welcome to Example WebPage",
//		"This is an example webpage",
//		"Something related to the home page",
//		"example, webpage, demo",
//		"en",
//		"",
//		"",
//		"",
//		"",
//		"",
//
// )
//
// // Rendering JSON-LD using templ:
//
//	templ Page() {
//		@webPage.ToJsonLd()
//	}
//
// // Rendering JSON-LD as `template.HTML` value:
//
//	jsonLdHtml := webPage.ToGoHTMLJsonLd()
//
// Expected output:
//
//	{
//		"@context": "https://schema.org",
//		"@type": "WebPage",
//		"url": "https://www.example.com",
//		"name": "Example WebPage",
//		"headline": "Welcome to Example WebPage",
//		"description": "This is an example webpage",
//		"keywords": "example, webpage, demo"
//	}
type WebPage struct {
	Context       string `json:"@context"`
	Type          string `json:"@type"`
	URL           string `json:"url,omitempty"`
	Name          string `json:"name,omitempty"`
	Headline      string `json:"headline,omitempty"`
	Description   string `json:"description,omitempty"`
	About         string `json:"about,omitempty"`
	Keywords      string `json:"keywords,omitempty"`
	InLanguage    string `json:"inLanguage,omitempty"`
	IsPartOf      string `json:"isPartOf,omitempty"`
	LastReviewed  string `json:"lastReviewed,omitempty"`
	PrimaryImage  string `json:"primaryImageOfPage,omitempty"`
	DatePublished string `json:"datePublished,omitempty"`
	DateModified  string `json:"dateModified,omitempty"`
}

func NewWebPage(url string, name string, headline string, description string, about string, keywords string, inLanguage string, isPartOf string, lastReviewed string, primaryImage string, datePublished string, dateModified string) *WebPage {
	webpage := &WebPage{
		URL:           url,
		Name:          name,
		Headline:      headline,
		Description:   description,
		About:         about,
		Keywords:      keywords,
		InLanguage:    inLanguage,
		IsPartOf:      isPartOf,
		LastReviewed:  lastReviewed,
		PrimaryImage:  primaryImage,
		DatePublished: datePublished,
		DateModified:  dateModified,
	}
	webpage.ensureDefaults()
	return webpage
}

func (wp *WebPage) Validate() []string {
	var warnings []string

	if wp.URL == "" {
		warnings = append(warnings, "missing recommended field: url")
	}
	if wp.Name == "" {
		warnings = append(warnings, "missing recommended field: name")
	}
	if wp.Headline == "" {
		warnings = append(warnings, "missing recommended field: headline")
	}
	if wp.Description == "" {
		warnings = append(warnings, "missing recommended field: description")
	}

	return warnings
}

// ToJsonLd converts the WebPage struct to a JSON-LD `templ.Component`.
func (wp *WebPage) ToJsonLd() templ.Component {
	wp.ensureDefaults()
	id := fmt.Sprintf("%s-%s", "webpage", teseo.GenerateUniqueKey())
	return templ.JSONScript(id, wp).WithType("application/ld+json")
}

// ToGoHTMLJsonLd renders the WebSite struct as `template.HTML` value for Go's `html/template`.
func (wp *WebPage) ToGoHTMLJsonLd() (template.HTML, error) {
	return teseo.RenderToHTML(wp.ToJsonLd())
}

func (wp *WebPage) ensureDefaults() {
	if wp.Context == "" {
		wp.Context = "https://schema.org"
	}

	if wp.Type == "" {
		wp.Type = "WebPage"
	}
}
