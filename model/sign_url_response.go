package model

type SignURLResponse struct {
	SignURL   string `json:"sign_url"`   // URL of the signature page to display in the embedded iFrame.
	ExpiresAt int    `json:"expires_at"` // When the link expires.
}
