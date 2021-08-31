package model

type ResponseData struct {
	ApiID       string `json:"api_id"`       // The unique ID for this field.
	SignatureID string `json:"signature_id"` // The ID of the signature to which this response is linked.
	Name        string `json:"name"`         // The name of the form field.
	Value       string `json:"value"`        // The value of the form field.
	Required    bool   `json:"required"`     // A boolean value denoting if this field is required.
	Type        string `json:"type"`         // The type of this form field. See field types
}