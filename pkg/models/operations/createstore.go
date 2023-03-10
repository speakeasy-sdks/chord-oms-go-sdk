package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateStoreStoreInput struct {
	AvailableLocales  []string `json:"available_locales,omitempty"`
	CartTaxCountryIso *string  `json:"cart_tax_country_iso,omitempty"`
	Code              *string  `json:"code,omitempty"`
	DefaultCurrency   *string  `json:"default_currency,omitempty"`
	MailFromAddress   *string  `json:"mail_from_address,omitempty"`
	MetaDescription   *string  `json:"meta_description,omitempty"`
	MetaKeywords      *string  `json:"meta_keywords,omitempty"`
	Name              *string  `json:"name,omitempty"`
	SeoTitle          *string  `json:"seo_title,omitempty"`
	URL               *string  `json:"url,omitempty"`
}

type CreateStoreSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateStore401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateStore422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateStoreRequest struct {
	Request  CreateStoreStoreInput `request:"mediaType=application/json"`
	Security CreateStoreSecurity
}

type CreateStoreResponse struct {
	ContentType                         string
	StatusCode                          int64
	CreateStore401ApplicationJSONObject *CreateStore401ApplicationJSON
	CreateStore422ApplicationJSONObject *CreateStore422ApplicationJSON
	Store                               *shared.Store
}
