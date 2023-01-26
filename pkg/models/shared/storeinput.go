package shared

type StoreInput struct {
	AvailableLocales  []string `json:"available_locales,omitempty"`
	CartTaxCountryIso *string  `json:"cart_tax_country_iso,omitempty"`
	DefaultCurrency   *string  `json:"default_currency,omitempty"`
	MailFromAddress   *string  `json:"mail_from_address,omitempty"`
	MetaDescription   *string  `json:"meta_description,omitempty"`
	MetaKeywords      *string  `json:"meta_keywords,omitempty"`
	Name              *string  `json:"name,omitempty"`
	SeoTitle          *string  `json:"seo_title,omitempty"`
	URL               *string  `json:"url,omitempty"`
}
