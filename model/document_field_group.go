package model

type DocumentFieldGroup struct {
	Name     string `json:"name"`
	Rule     string `json:"rule"`
}

// GetName returns Name
func (d *DocumentFieldGroup) GetName() string {
	if d != nil {
		return d.Name
	}

	return ""
}

// GetRule returns Rule
func (d *DocumentFieldGroup) GetRule() string {
	if d != nil {
		return d.Rule
	}

	return ""
}
