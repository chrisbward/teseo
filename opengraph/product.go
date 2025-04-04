package opengraph

import (
	"context"
	"html/template"
	"io"
	"log"

	"github.com/a-h/templ"
	"github.com/indaco/teseo"
)

// Product represents the Open Graph product metadata.
// For more details about the meaning of the properties see: https://ogp.me/#metadata
//
// Example usage:
//
// Pure struct usage:
//
//	// Create a product using pure struct
//	product := &opengraph.Product{
//		OpenGraphObject: opengraph.OpenGraphObject{
//			Title:       "Example Product",
//			URL:         "https://www.example.com/product/example-product",
//			Description: "This is an example product description.",
//			Image:       "https://www.example.com/images/product.jpg",
//		},
//		Price:        "29.99",
//		PriceCurrency: "USD",
//	}
//
// Factory method usage:
//
//	// Create a product
//	product := opengraph.NewProduct(
//		"Example Product",
//		"https://www.example.com/product/example-product",
//		"This is an example product description.",
//		"https://www.example.com/images/product.jpg",
//		"29.99",
//		"USD",
//	)
//
// // Rendering the HTML meta tags using templ:
//
//	templ Page() {
//		@product.ToMetaTgs()
//	}
//
// // Rendering the HTML meta tags as `template.HTML` value:
//
//	metaTagsHtml := product.ToGoHTMLMetaTgs()
//
// Expected output:
//
//	<meta property="og:type" content="product"/>
//	<meta property="og:title" content="Example Product"/>
//	<meta property="og:url" content="https://www.example.com/product/example-product"/>
//	<meta property="og:description" content="This is an example product description."/>
//	<meta property="og:image" content="https://www.example.com/images/product.jpg"/>
//	<meta property="product:price:amount" content="29.99"/>
//	<meta property="product:price:currency" content="USD"/>
type Product struct {
	OpenGraphObject
	Price         string // product:price:amount, price of the product
	PriceCurrency string // product:price:currency, currency of the price
}

// NewProduct initializes a Product with the default type "product".
func NewProduct(title, url, description, image, price, priceCurrency string) *Product {
	product := &Product{
		OpenGraphObject: OpenGraphObject{
			Title:       title,
			URL:         url,
			Description: description,
			Image:       image,
		},
		Price:         price,
		PriceCurrency: priceCurrency,
	}
	product.ensureDefaults()
	return product
}

// ToMetaTags generates the HTML meta tags for the Open Graph Product as templ.Component.
func (p *Product) ToMetaTags() templ.Component {
	p.ensureDefaults()
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		for _, tag := range p.metaTags() {
			if tag.content != "" {
				if err := teseo.WriteMetaTag(w, tag.property, tag.content); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

// ToGoHTMLMetaTags generates the HTML meta tags for the Open Graph Product as `template.HTML` value for Go's `html/template`.
func (p *Product) ToGoHTMLMetaTags() (template.HTML, error) {
	html, err := templ.ToGoHTML(context.Background(), p.ToMetaTags())
	if err != nil {
		log.Printf("failed to convert to html: %v", err)
		return "", err
	}
	return html, nil
}

// ensureDefaults sets default values for Product.
func (p *Product) ensureDefaults() {
	p.OpenGraphObject.ensureDefaults("product")
}

// metaTags returns all meta tags for the Product object, including OpenGraphObject fields and product-specific ones.
func (p *Product) metaTags() []metaTag {
	tags := []metaTag{
		{"og:type", "product"},
		{"og:title", p.Title},
		{"og:url", p.URL},
		{"og:description", p.Description},
		{"og:image", p.Image},
	}

	if p.Price != "" {
		tags = append(tags, metaTag{"product:price:amount", p.Price})
	}
	if p.PriceCurrency != "" {
		tags = append(tags, metaTag{"product:price:currency", p.PriceCurrency})
	}

	return tags
}
