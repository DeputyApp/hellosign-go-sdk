package model

type TemplateDocumentFormField struct {
	APIId       string       `json:"api_id"`
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	Width       float64      `json:"width"`
	Height      float64      `json:"height"`
	Required    bool         `json:"required"`
	SignerRoles []SignerRole `json:"signer_roles"`
}

// GetAPIId returns APIId
func (d *TemplateDocumentFormField) GetAPIId() string {
	if d != nil {
		return d.APIId
	}
	return ""
}

// GetName returns Name
func (d *TemplateDocumentFormField) GetName() string {
	if d != nil {
		return d.Name
	}
	return ""
}

// GetType returns Type
func (d *TemplateDocumentFormField) GetType() string {
	if d != nil {
		return d.Type
	}
	return ""
}

// GetWidth returns Width
func (d *TemplateDocumentFormField) GetWidth() float64 {
	if d != nil {
		return d.Width
	}
	return 0
}

// GetHeight returns Height
func (d *TemplateDocumentFormField) GetHeight() float64 {
	if d != nil {
		return d.Height
	}
	return 0
}

// GetRequired returns Required
func (d *TemplateDocumentFormField) GetRequired() bool {
	if d != nil {
		return d.Required
	}
	return false
}

// GetSignerRoles returns SignerRoles
func (d *TemplateDocumentFormField) GetSignerRoles() []SignerRole {
	if d != nil {
		return d.SignerRoles
	}
	return nil
}
