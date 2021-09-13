package model

type TemplateDocumentFormField struct {
	APIId       string       `json:"api_id"`
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	X           int          `json:"x"`
	Y           int          `json:"y"`
	Width       int          `json:"width"`
	Height      int          `json:"height"`
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

// GetX returns X
func (d *TemplateDocumentFormField) GetX() int {
	if d != nil {
		return d.X
	}
	return 0
}

// GetY returns Y
func (d *TemplateDocumentFormField) GetY() int {
	if d != nil {
		return d.Y
	}
	return 0
}

// GetWidth returns Width
func (d *TemplateDocumentFormField) GetWidth() int {
	if d != nil {
		return d.Width
	}
	return 0
}

// GetHeight returns Height
func (d *TemplateDocumentFormField) GetHeight() int {
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
