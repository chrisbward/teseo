package opengraph

import (
	"context"
	"html/template"
	"io"
	"log"

	"github.com/a-h/templ"
	"github.com/indaco/teseo"
)

// WebSite represents the Open Graph website metadata.
// For more details about the meaning of the properties see: https://ogp.me/#type_website
//
// Example usage:
//
// Pure struct usage:
//
//	// Create a website using pure struct
//	website := &opengraph.WebSite{
//		OpenGraphObject: opengraph.OpenGraphObject{
//			Title:       "Example Website",
//			URL:         "https://www.example.com",
//			Description: "This is an example website description.",
//			Image:       "https://www.example.com/images/logo.jpg",
//		},
//	}
//
// Factory method usage:
//
//	// Create a website using the factory method
//	website := opengraph.NewWebSite(
//		"Example Website",
//		"https://www.example.com",
//		"This is an example website description.",
//		"https://www.example.com/images/logo.jpg",
//	)
//
// // Rendering the HTML meta tags using templ:
//
//	templ Page() {
//		@website.ToMetaTgs()
//	}
//
// // Rendering the HTML meta tags as `template.HTML` value:
//
//	metaTagsHtml := website.ToGoHTMLMetaTgs()
//
// Expected output:
//
//	<meta property="og:type" content="website"/>
//	<meta property="og:title" content="Example Website"/>
//	<meta property="og:url" content="https://www.example.com"/>
//	<meta property="og:description" content="This is an example website description."/>
//	<meta property="og:image" content="https://www.example.com/images/logo.jpg"/>
type WebSite struct {
	OpenGraphObject
}

// NewWebSite initializes a WebSite with the default type "website".
func NewWebSite(title, url, description, image string) *WebSite {
	website := &WebSite{
		OpenGraphObject: OpenGraphObject{
			Title:       title,
			URL:         url,
			Description: description,
			Image:       image,
		},
	}
	website.ensureDefaults()
	return website
}

// ToMetaTags generates the HTML meta tags for the Open Graph WebSite using templ.Component.
func (ws *WebSite) ToMetaTags() templ.Component {
	ws.ensureDefaults()
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		for _, tag := range ws.metaTags() {
			if tag.content != "" {
				if err := teseo.WriteMetaTag(w, tag.property, tag.content); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

// ToGoHTMLMetaTags generates the HTML meta tags for the Open Graph WebSite as `template.HTML` value for Go's `html/template`.
func (ws *WebSite) ToGoHTMLMetaTags() (template.HTML, error) {
	html, err := templ.ToGoHTML(context.Background(), ws.ToMetaTags())
	if err != nil {
		log.Printf("failed to convert to html: %v", err)
		return "", err
	}
	return html, nil
}

// ensureDefaults sets default values for WebSite.
func (ws *WebSite) ensureDefaults() {
	ws.OpenGraphObject.ensureDefaults("website")
}

// metaTags returns the meta tags for the WebSite as a slice of property-content pairs.
func (ws *WebSite) metaTags() []metaTag {
	return []metaTag{
		{"og:type", "website"},
		{"og:title", ws.Title},
		{"og:url", ws.URL},
		{"og:description", ws.Description},
		{"og:image", ws.Image},
	}
}
