package model

type EmbeddedTemplate struct {
	TemplateID string `json:"template_id"`
	EditURL    string `json:"edit_url"`
	ExpiresAt  int32  `json:"expires_at"`
}

// GetTemplateID returns TemplateID
func (e *EmbeddedTemplate) GetTemplateID() string {
	if e != nil {
		return e.TemplateID
	}
	return ""
}

// GetEditURL returns EditURL
func (e *EmbeddedTemplate) GetEditURL() string {
	if e != nil {
		return e.EditURL
	}
	return ""
}

// GetExpiresAt returns ExpiresAt
func (e *EmbeddedTemplate) GetExpiresAt() int32 {
	if e != nil {
		return e.ExpiresAt
	}
	return 0
}
