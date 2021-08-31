package model

type SignURLResponse struct {
	SignURL   string `json:"sign_url"`   // URL of the signature page to display in the embedded iFrame.
	ExpiresAt int    `json:"expires_at"` // When the link expires.
}

// GetSignUrl returns SignURL
func (s *SignURLResponse) GetSignUrl() string {
	if s != nil {
		return s.SignURL
	}
	return ""
}

// GetExpiresAt returns ExpiresAt
func (s *SignURLResponse) GetExpiresAt() int {
	if s != nil {
		return s.ExpiresAt
	}
	return 0
}
