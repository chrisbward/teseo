package schemaorg

import (
	"encoding/json"
	"testing"
)

// Struct that uses the custom StringOrSlice type
type MyStruct struct {
	AreaServed StringOrSlice `json:"areaServed,omitempty"`
}

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    MyStruct
		expected string
	}{

		{
			name:     "Non-empty slice",
			input:    MyStruct{AreaServed: StringOrSlice{"hello", "world"}},
			expected: `{"areaServed":["hello","world"]}`,
		},
		{
			name:     "String",
			input:    MyStruct{AreaServed: StringOrSlice{"hello"}},
			expected: `{"areaServed":"hello"}`,
		},
		{
			name:     "Empty slice (should be omitted)",
			input:    MyStruct{AreaServed: StringOrSlice{}},
			expected: `{}`,
		},
		{
			name:     "Nil slice (should be omitted)",
			input:    MyStruct{},
			expected: `{}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("unexpected error marshaling: %v", err)
			}

			if string(data) != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, data)
			}
		})
	}
}
