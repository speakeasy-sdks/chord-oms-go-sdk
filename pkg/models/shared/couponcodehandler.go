package shared

type CouponCodeHandler struct {
	Error      *string `json:"error,omitempty"`
	StatusCode *string `json:"status_code,omitempty"`
	Success    *string `json:"success,omitempty"`
	Successful *bool   `json:"successful,omitempty"`
}
