package model

// EmbeddedSignatureWithTemplateRequest contains the request parameters for create_embedded
type EmbeddedSignatureWithTemplateRequest struct {
	TestMode         bool              `form_field:"test_mode"`
	ClientID         string            `form_field:"client_id"`
	Title            string            `form_field:"title"`
	Subject          string            `form_field:"subject"`
	Message          string            `form_field:"message"`
	Signers          []Signer          `form_field:"signers"`
	CustomFields     []CustomField     `form_field:"custom_fields"`
	CCEmailAddresses []string          `form_field:"cc_email_addresses"`
	Metadata         map[string]string `form_field:"metadata"`
	TemplateID       string            `form_field:"template_id"`
}

// GetTestMode returns TestMode
func (e *EmbeddedSignatureWithTemplateRequest) GetTestMode() bool {
	if e != nil {
		return e.TestMode
	}
	return false
}

// GetClientID returns ClientID
func (e *EmbeddedSignatureWithTemplateRequest) GetClientID() string {
	if e != nil {
		return e.ClientID
	}
	return ""
}

// GetTitle returns Title
func (e *EmbeddedSignatureWithTemplateRequest) GetTitle() string {
	if e != nil {
		return e.Title
	}
	return ""
}

// GetSubject returns Subject
func (e *EmbeddedSignatureWithTemplateRequest) GetSubject() string {
	if e != nil {
		return e.Subject
	}
	return ""
}

// GetMessage returns Message
func (e *EmbeddedSignatureWithTemplateRequest) GetMessage() string {
	if e != nil {
		return e.Message
	}
	return ""
}

// GetSigners returns Signers
func (e *EmbeddedSignatureWithTemplateRequest) GetSigners() []Signer {
	if e != nil {
		return e.Signers
	}
	return nil
}

// GetCustomFields returns CustomFields
func (e *EmbeddedSignatureWithTemplateRequest) GetCustomFields() []CustomField {
	if e != nil {
		return e.CustomFields
	}
	return nil
}

// GetCCEmailAddresses returns CCEmailAddresses
func (e *EmbeddedSignatureWithTemplateRequest) GetCCEmailAddresses() []string {
	if e != nil {
		return e.CCEmailAddresses
	}
	return nil
}

// GetMetadata returns Metadata
func (e *EmbeddedSignatureWithTemplateRequest) GetMetadata() map[string]string {
	if e != nil {
		return e.Metadata
	}
	return nil
}

// GetTemplateID returns TemplateID
func (e *EmbeddedSignatureWithTemplateRequest) GetTemplateID() string {
	if e != nil {
		return e.TemplateID
	}
	return ""
}
