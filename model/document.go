package model

type Document struct {
	Name         string               `json:"name"`
	Index        int                  `json:"index"`
	FieldGroups  []DocumentFieldGroup `json:"field_groups"`
	FormFields   []DocumentFormField  `json:"form_fields"`
	CustomFields []CustomField        `json:"custom_fields"`
}

// GetName returns Name
func (d *Document) GetName() string {
	if d != nil {
		return d.Name
	}

	return ""
}

// GetIndex returns Index
func (d *Document) GetIndex() int {
	if d != nil {
		return d.Index
	}

	return 0
}

// GetFieldGroups returns FieldGroups
func (d *Document) GetFieldGroups() []DocumentFieldGroup {
	if d != nil {
		return d.FieldGroups
	}

	return nil
}

// GetFormFields returns FormFields
func (d *Document) GetFormFields() []DocumentFormField {
	if d != nil {
		return d.FormFields
	}

	return nil
}

// GetCustomFields returns CustomFields
func (d *Document) GetCustomFields() []CustomField {
	if d != nil {
		return d.CustomFields
	}

	return nil
}
