package model

type CustomField struct {
	Name     string      `json:"name"`     // The name of the Custom Field.
	Type     string      `json:"type"`     // The type of this Custom Field. Only 'text' and 'checkbox' are currently supported.
	Value    interface{} `json:"value"`    // A text string for text fields or true/false for checkbox fields
	Required bool        `json:"required"` // A boolean value denoting if this field is required.
	ApiID    string      `json:"api_id"`   // The unique ID for this field.
	Editor   *string     `json:"editor"`   // The name of the Role that is able to edit this field.
}