package schemaorg

import (
	"encoding/json"
)

// Common type definitions used across multiple JSON-LD entities

type StringOrSlice []string

func (s StringOrSlice) IsZero() bool {
	return len(s) == 0 || (len(s) == 1 && s[0] == "")
}

// UnmarshalJSON handles both string and []string
func (s StringOrSlice) MarshalJSON() ([]byte, error) {
	if len(s) == 0 || (len(s) == 1 && s[0] == "") {
		// Should never get here if IsZero is respected, but safe fallback:
		return []byte("null"), nil
	}
	if len(s) == 1 {
		return json.Marshal(s[0])
	}
	return json.Marshal([]string(s))
}

// ContactPoint represents a Schema.org ContactPoint object
// For more details about the meaning of the properties see: https://schema.org/ContactPoint
type ContactPoint struct {
	Type              string        `json:"@type"`
	Telephone         string        `json:"telephone,omitempty"`
	ContactType       string        `json:"contactType,omitempty"`
	ContactOption     StringOrSlice `json:"contactOption,omitempty"`
	AreaServed        StringOrSlice `json:"areaServed,omitempty"`
	AvailableLanguage string        `json:"availableLanguage,omitempty"`
}

// ImageObject represents a Schema.org ImageObject object
// For more details about the meaning of the properties see: https://schema.org/ImageObject
type ImageObject struct {
	Type string `json:"@type"`
	URL  string `json:"url,omitempty"`
}

// ensureDefaults sets default values for ImageObject if they are not already set.
func (img *ImageObject) ensureDefaults() {
	if img.Type == "" {
		img.Type = "ImageObject"
	}
}

// ListItem represents a Schema.org ListItem object
// For more details about the meaning of the properties see: https://schema.org/ListItem
type ListItem struct {
	Type     string `json:"@type"`
	Position int    `json:"position,omitempty"`
	Name     string `json:"name,omitempty"`
	Item     string `json:"item,omitempty"`
}
