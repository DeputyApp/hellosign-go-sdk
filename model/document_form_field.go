package model

type DocumentFormField struct {
	APIId    string  `json:"api_id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	Required bool    `json:"required"`
	Signer   int     `json:"signer"`
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
func (d *DocumentFormField) GetX() float64 {
	if d != nil {
		return d.X
	}
	return 0
}

// GetY returns Y
func (d *DocumentFormField) GetY() float64 {
	if d != nil {
		return d.Y
	}
	return 0
}

// GetWidth returns Width
func (d *DocumentFormField) GetWidth() float64 {
	if d != nil {
		return d.Width
	}
	return 0
}

// GetHeight returns Height
func (d *DocumentFormField) GetHeight() float64 {
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
