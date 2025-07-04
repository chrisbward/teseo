package opengraph

import (
	"context"
	"html/template"
	"io"

	"github.com/a-h/templ"
	"github.com/indaco/teseo"
)

// Restaurant represents the Open Graph restaurant metadata.
// For more details about the meaning of the properties see: https://ogp.me/#metadata
//
// Example usage:
//
// Pure struct usage:
//
//	// Create a restaurant using pure struct
//	restaurant := &opengraph.Restaurant{
//		OpenGraphObject: opengraph.OpenGraphObject{
//			Title:       "Example Restaurant",
//			URL:         "https://www.example.com/restaurant/example-restaurant",
//			Description: "This is an example restaurant description.",
//			Image:       "https://www.example.com/images/restaurant.jpg",
//		},
//		StreetAddress: "123 Food Street",
//		Locality:      "Gourmet City",
//		Region:        "CA",
//		PostalCode:    "12345",
//		Country:       "USA",
//		Phone:         "+1-800-FOOD-123",
//		MenuURL:       "https://www.example.com/menu",
//		ReservationURL: "https://www.example.com/reservations",
//	}
//
// Factory method usage:
//
//	// Create a restaurant using the factory method
//	restaurant := opengraph.NewRestaurant(
//		"Example Restaurant",
//		"https://www.example.com/restaurant/example-restaurant",
//		"This is an example restaurant description.",
//		"https://www.example.com/images/restaurant.jpg",
//		"123 Food Street",
//		"Gourmet City",
//		"CA",
//		"12345",
//		"USA",
//		"+1-800-FOOD-123",
//		"https://www.example.com/menu",
//		"https://www.example.com/reservations",
//	)
//
// // Rendering the HTML meta tags using templ:
//
//	templ Page() {
//		@restaurant.ToMetaTgs()
//	}
//
// // Rendering the HTML meta tags as `template.HTML` value:
//
//	metaTagsHtml := restaurant.ToGoHTMLMetaTgs()
//
// Expected output:
//
//	<meta property="og:type" content="restaurant"/>
//	<meta property="og:title" content="Example Restaurant"/>
//	<meta property="og:url" content="https://www.example.com/restaurant/example-restaurant"/>
//	<meta property="og:description" content="This is an example restaurant description."/>
//	<meta property="og:image" content="https://www.example.com/images/restaurant.jpg"/>
//	<meta property="place:contact_data:street_address" content="123 Food Street"/>
//	<meta property="place:contact_data:locality" content="Gourmet City"/>
//	<meta property="place:contact_data:region" content="CA"/>
//	<meta property="place:contact_data:postal_code" content="12345"/>
//	<meta property="place:contact_data:country_name" content="USA"/>
//	<meta property="place:contact_data:phone_number" content="+1-800-FOOD-123"/>
//	<meta property="restaurant:menu" content="https://www.example.com/menu"/>
//	<meta property="restaurant:reservation" content="https://www.example.com/reservations"/>
type Restaurant struct {
	OpenGraphObject
	StreetAddress  string // place:contact_data:street_address, street address of the restaurant
	Locality       string // place:contact_data:locality, locality or city of the restaurant
	Region         string // place:contact_data:region, region or state of the restaurant
	PostalCode     string // place:contact_data:postal_code, postal code of the restaurant
	Country        string // place:contact_data:country_name, country of the restaurant
	Phone          string // place:contact_data:phone_number, phone number of the restaurant
	MenuURL        string // restaurant:menu, URL to the restaurant's menu
	ReservationURL string // restaurant:reservation, URL to the reservation page
}

// NewRestaurant initializes a Restaurant with the default type "restaurant".
func NewRestaurant(title, url, description, image, streetAddress, locality, region, postalCode, country, phone, menuURL, reservationURL string) *Restaurant {
	restaurant := &Restaurant{
		OpenGraphObject: OpenGraphObject{
			Title:       title,
			URL:         url,
			Description: description,
			Image:       image,
		},
		StreetAddress:  streetAddress,
		Locality:       locality,
		Region:         region,
		PostalCode:     postalCode,
		Country:        country,
		Phone:          phone,
		MenuURL:        menuURL,
		ReservationURL: reservationURL,
	}
	restaurant.ensureDefaults()
	return restaurant
}

// ToMetaTags generates the HTML meta tags for the Open Graph Restaurant as templ.Component.
func (restaurant *Restaurant) ToMetaTags() templ.Component {
	restaurant.ensureDefaults()
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		for _, tag := range restaurant.metaTags() {
			if tag.content != "" {
				if err := teseo.WriteMetaTag(w, tag.property, tag.content); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

// ToGoHTMLMetaTags generates the HTML meta tags for the Open Graph Restaurant as `template.HTML` value for Go's `html/template`.
func (restaurant *Restaurant) ToGoHTMLMetaTags() (template.HTML, error) {
	return teseo.RenderToHTML(restaurant.ToMetaTags())
}

// ensureDefaults sets default values for Restaurant.
func (restaurant *Restaurant) ensureDefaults() {
	restaurant.OpenGraphObject.ensureDefaults("restaurant")
}

// metaTags returns all meta tags for the Restaurant object, including OpenGraphObject fields and restaurant-specific ones.
func (r *Restaurant) metaTags() []metaTag {
	tags := []metaTag{
		{"og:type", "restaurant"},
		{"og:title", r.Title},
		{"og:url", r.URL},
		{"og:description", r.Description},
		{"og:image", r.Image},
	}

	if r.StreetAddress != "" {
		tags = append(tags, metaTag{"place:contact_data:street_address", r.StreetAddress})
	}
	if r.Locality != "" {
		tags = append(tags, metaTag{"place:contact_data:locality", r.Locality})
	}
	if r.Region != "" {
		tags = append(tags, metaTag{"place:contact_data:region", r.Region})
	}
	if r.PostalCode != "" {
		tags = append(tags, metaTag{"place:contact_data:postal_code", r.PostalCode})
	}
	if r.Country != "" {
		tags = append(tags, metaTag{"place:contact_data:country_name", r.Country})
	}
	if r.Phone != "" {
		tags = append(tags, metaTag{"place:contact_data:phone_number", r.Phone})
	}
	if r.MenuURL != "" {
		tags = append(tags, metaTag{"restaurant:menu", r.MenuURL})
	}
	if r.ReservationURL != "" {
		tags = append(tags, metaTag{"restaurant:reservation", r.ReservationURL})
	}

	return tags
}
