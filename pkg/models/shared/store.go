package shared

type Store struct {
	AvailableLocales []string `json:"available_locales,omitempty"`
	Code             *string  `json:"code,omitempty"`
	Default          *bool    `json:"default,omitempty"`
	DefaultCurrency  *string  `json:"default_currency,omitempty"`
	ID               *int64   `json:"id,omitempty"`
	MailFromAddress  *string  `json:"mail_from_address,omitempty"`
	MetaDescription  *string  `json:"meta_description,omitempty"`
	MetaKeywords     *string  `json:"meta_keywords,omitempty"`
	Name             *string  `json:"name,omitempty"`
	SeoTitle         *string  `json:"seo_title,omitempty"`
	URL              *string  `json:"url,omitempty"`
}
