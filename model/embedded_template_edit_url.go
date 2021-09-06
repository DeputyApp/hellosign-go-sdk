package model

type EmbeddedTemplateEditURL struct {
	EditURL   string `json:"edit_url"`   // URL of the template to display in the embedded iFrame.
	ExpiresAt int    `json:"expires_at"` // When the link expires.
}

// GetEditURL returns EditURL
func (t *EmbeddedTemplateEditURL) GetEditURL() string {
	if t != nil {
		return t.EditURL
	}
	return ""
}

// GetExpiresAt returns ExpiresAt
func (t *EmbeddedTemplateEditURL) GetExpiresAt() int {
	if t != nil {
		return t.ExpiresAt
	}
	return 0
}
