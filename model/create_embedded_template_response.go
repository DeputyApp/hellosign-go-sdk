package model

type CreateEmbeddedTemplateResponse struct {
	Template *Template `json:"template"`
}

// GetTemplate returns SignatureRequest
func (cr *CreateEmbeddedTemplateResponse) GetTemplate() *Template {
	if cr != nil {
		return cr.Template
	}
	return nil
}
