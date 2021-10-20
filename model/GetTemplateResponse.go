package model

type GetTemplateResponse struct {
	Template *Template `json:"template"`
}

// GetTemplate returns Template
func (t *GetTemplateResponse) GetTemplate() *Template {
	if t != nil {
		return t.Template
	}
	return nil
}