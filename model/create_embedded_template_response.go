package model

type CreateEmbeddedTemplateResponse struct {
	Template *EmbeddedTemplate `json:"template"`
}

// GetTemplate returns Template
func (cr *CreateEmbeddedTemplateResponse) GetTemplate() *EmbeddedTemplate {
	if cr != nil {
		return cr.Template
	}
	return nil
}
