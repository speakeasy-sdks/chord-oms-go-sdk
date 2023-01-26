package shared

type SchemeAPIKey struct {
	APIKey string `security:"name=Authorization"`
}

type SchemeOrderToken struct {
	APIKey string `security:"name=X-Spree-Order-Token"`
}
