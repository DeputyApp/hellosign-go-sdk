package model

type CustomField struct {
	Name     string      `json:"name"`     // The name of the Custom Field.
	Type     string      `json:"type"`     // The type of this Custom Field. Only 'text' and 'checkbox' are currently supported.
	Value    interface{} `json:"value"`    // A text string for text fields or true/false for checkbox fields
	Required bool        `json:"required"` // A boolean value denoting if this field is required.
	ApiID    string      `json:"api_id"`   // The unique ID for this field.
	Editor   *string     `json:"editor"`   // The name of the Role that is able to edit this field.
}

// GetName returns Name
func (c *CustomField) GetName() string {
	if c != nil {
		return c.Name
	}
	return ""
}

// GetType returns Type
func (c *CustomField) GetType() string {
	if c != nil {
		return c.Type
	}
	return ""
}

// GetValue returns Value
func (c *CustomField) GetValue() interface{} {
	if c != nil {
		return c.Value
	}
	return ""
}

// GetRequired returns Required
func (c *CustomField) GetRequired() bool {
	if c != nil {
		return c.Required
	}
	return false
}

// GetApiID returns ApiID
func (c *CustomField) GetApiID() string {
	if c != nil {
		return c.ApiID
	}
	return ""
}

// GetEditor returns Editor
func (c *CustomField) GetEditor() *string {
	if c != nil {
		return c.Editor
	}
	return nil
}
