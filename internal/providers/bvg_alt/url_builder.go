package bvg_alt

import (
	"fmt"
	"github.com/mboufous/berlin-departure-board/internal"
	"net/url"
)

type URLBuilder struct {
	baseURL   string
	stationID string
	endpoint  string
	params    url.Values
}

// NewURLBuilder initializes a new Builder instance.
func NewURLBuilder(baseURL, stationID string) *URLBuilder {
	return &URLBuilder{
		baseURL:   baseURL,
		stationID: stationID,
		params:    url.Values{},
	}
}

// WithEndpoint sets the API endpoint.
func (b *URLBuilder) WithEndpoint(endpoint string) *URLBuilder {
	b.endpoint = endpoint
	return b
}

// WithDefaultParam adds a parameter with a default value if not provided.
func (b *URLBuilder) WithDefaultParam(key, value, defaultValue string) *URLBuilder {
	if value == "" {
		b.params.Set(key, defaultValue)
	} else {
		b.params.Set(key, value)
	}
	return b
}

// WithParam adds a parameter.
func (b *URLBuilder) WithParam(key, value string) *URLBuilder {
	if value != "" {
		b.params.Set(key, value)
	}
	return b
}

// WithBooleanParam adds a boolean parameter if enabled.
func (b *URLBuilder) WithBooleanParam(key string, enabled bool) *URLBuilder {
	if enabled {
		b.params.Set(key, "true")
	}
	return b
}

// WithProductFilters appends the selected products to the query parameters.
func (b *URLBuilder) WithProductFilters(products []string) *URLBuilder {
	allProducts := map[string]string{
		internal.ProductSuburban: "false",
		internal.ProductSubway:   "false",
		internal.ProductTram:     "false",
		internal.ProductBus:      "false",
		internal.ProductFerry:    "false",
		internal.ProductExpress:  "false",
		internal.ProductRegional: "false",
	}

	// Set specified products to "true"
	for _, product := range products {
		if _, exists := allProducts[product]; exists {
			allProducts[product] = "true"
		}
	}

	// Add all products as parameters
	for product, value := range allProducts {
		b.params.Set(product, value)
	}

	return b
}

// Build constructs the final URL string.
func (b *URLBuilder) Build() (string, error) {
	u, err := url.Parse(fmt.Sprintf("%s/stops/%s%s", b.baseURL, b.stationID, b.endpoint))
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}
	u.RawQuery = b.params.Encode()
	return u.String(), nil
}
