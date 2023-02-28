package shared

type SchemeAPIKey struct {
	Authorization string `security:"name=Authorization"`
}

type SchemeStorefrontLogin struct {
	APIKey string `security:"name=Authorization"`
}

type SchemeOrderToken struct {
	APIKey string `security:"name=X-Spree-Order-Token"`
}
