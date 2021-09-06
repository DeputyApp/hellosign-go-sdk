package model

type DocumentFormField struct {
	APIId    string `json:"api_id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Required bool   `json:"required"`
	Signer   int    `json:"signer"`
}

// GetAPIId returns APIId
func (d *DocumentFormField) GetAPIId() string {
	if d != nil {
		return d.APIId
	}
	return ""
}

// GetName returns Name
func (d *DocumentFormField) GetName() string {
	if d != nil {
		return d.Name
	}
	return ""
}

// GetType returns Type
func (d *DocumentFormField) GetType() string {
	if d != nil {
		return d.Type
	}
	return ""
}

// GetX returns X
func (d *DocumentFormField) GetX() int {
	if d != nil {
		return d.X
	}
	return 0
}

// GetY returns Y
func (d *DocumentFormField) GetY() int {
	if d != nil {
		return d.Y
	}
	return 0
}

// GetWidth returns Width
func (d *DocumentFormField) GetWidth() int {
	if d != nil {
		return d.Width
	}
	return 0
}

// GetHeight returns Height
func (d *DocumentFormField) GetHeight() int {
	if d != nil {
		return d.Height
	}
	return 0
}

// GetRequired returns Required
func (d *DocumentFormField) GetRequired() bool {
	if d != nil {
		return d.Required
	}
	return false
}

// GetSigner returns Signer
func (d *DocumentFormField) GetSigner() int {
	if d != nil {
		return d.Signer
	}
	return 0
}