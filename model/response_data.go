package model

type ResponseData struct {
	ApiID       string `json:"api_id"`       // The unique ID for this field.
	SignatureID string `json:"signature_id"` // The ID of the signature to which this response is linked.
	Name        string `json:"name"`         // The name of the form field.
	Value       string `json:"value"`        // The value of the form field.
	Required    bool   `json:"required"`     // A boolean value denoting if this field is required.
	Type        string `json:"type"`         // The type of this form field. See field types
}

// GetApiID returns ApiID
func (r *ResponseData) GetApiID() string {
	if r != nil {
		return r.ApiID
	}
	return ""
}

// GetSignatureID returns SignatureID
func (r *ResponseData) GetSignatureID() string {
	if r != nil {
		return r.SignatureID
	}
	return ""
}

// GetName returns Name
func (r *ResponseData) GetName() string {
	if r != nil {
		return r.Name
	}
	return ""
}

// GetValue returns Value
func (r *ResponseData) GetValue() string {
	if r != nil {
		return r.Value
	}
	return ""
}

// GetRequired returns Required
func (r *ResponseData) GetRequired() bool {
	if r != nil {
		return r.Required
	}
	return false
}

// GetType returns Type
func (r *ResponseData) GetType() string {
	if r != nil {
		return r.Type
	}
	return ""
}
