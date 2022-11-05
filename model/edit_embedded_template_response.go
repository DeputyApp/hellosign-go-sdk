package model

// EditEmbeddedTemplateResponse response object
type EditEmbeddedTemplateResponse struct {
	Template *EmbeddedTemplate `json:"template"`
}

// GetTemplate returns Template
func (cr *EditEmbeddedTemplateResponse) GetTemplate() *EmbeddedTemplate {
	if cr != nil {
		return cr.Template
	}
	return nil
}
