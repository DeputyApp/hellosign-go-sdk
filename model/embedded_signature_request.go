package model

// EmbeddedSignatureRequest contains the request parameters for create_embedded
type EmbeddedSignatureRequest struct {
	TestMode              bool                  `form_field:"test_mode"`
	ClientID              string                `form_field:"client_id"`
	FileURL               []string              `form_field:"file_url"`
	File                  []string              `form_field:"file"`
	Title                 string                `form_field:"title"`
	Subject               string                `form_field:"subject"`
	Message               string                `form_field:"message"`
	SigningRedirectURL    string                `form_field:"signing_redirect_url"`
	Signers               []Signer              `form_field:"signers"`
	CustomFields          []CustomField         `form_field:"custom_fields"`
	CCEmailAddresses      []string              `form_field:"cc_email_addresses"`
	UseTextTags           bool                  `form_field:"use_text_tags"`
	HideTextTags          bool                  `form_field:"hide_text_tags"`
	Metadata              map[string]string     `form_field:"metadata"`
	FormFieldsPerDocument [][]DocumentFormField `form_field:"form_fields_per_document"`
}

// GetTestMode returns TestMode
func (e *EmbeddedSignatureRequest) GetTestMode() bool {
	if e != nil {
		return e.TestMode
	}
	return false
}

// GetClientID returns ClientID
func (e *EmbeddedSignatureRequest) GetClientID() string {
	if e != nil {
		return e.ClientID
	}
	return ""
}

// GetFileURL returns FileURL
func (e *EmbeddedSignatureRequest) GetFileURL() []string {
	if e != nil {
		return e.FileURL
	}
	return nil
}

// GetFile returns File
func (e *EmbeddedSignatureRequest) GetFile() []string {
	if e != nil {
		return e.File
	}
	return nil
}

// GetTitle returns Title
func (e *EmbeddedSignatureRequest) GetTitle() string {
	if e != nil {
		return e.Title
	}
	return ""
}

// GetSubject returns Subject
func (e *EmbeddedSignatureRequest) GetSubject() string {
	if e != nil {
		return e.Subject
	}
	return ""
}

// GetMessage returns Message
func (e *EmbeddedSignatureRequest) GetMessage() string {
	if e != nil {
		return e.Message
	}
	return ""
}

// GetSigningRedirectURL returns SigningRedirectURL
func (e *EmbeddedSignatureRequest) GetSigningRedirectURL() string {
	if e != nil {
		return e.SigningRedirectURL
	}
	return ""
}

// GetSigners returns Signers
func (e *EmbeddedSignatureRequest) GetSigners() []Signer {
	if e != nil {
		return e.Signers
	}
	return nil
}

// GetCustomFields returns CustomFields
func (e *EmbeddedSignatureRequest) GetCustomFields() []CustomField {
	if e != nil {
		return e.CustomFields
	}
	return nil
}

// GetCCEmailAddresses returns CCEmailAddresses
func (e *EmbeddedSignatureRequest) GetCCEmailAddresses() []string {
	if e != nil {
		return e.CCEmailAddresses
	}
	return nil
}

// GetUseTextTags returns UseTextTags
func (e *EmbeddedSignatureRequest) GetUseTextTags() bool {
	if e != nil {
		return e.UseTextTags
	}
	return false
}

// GetHideTextTags returns HideTextTags
func (e *EmbeddedSignatureRequest) GetHideTextTags() bool {
	if e != nil {
		return e.HideTextTags
	}
	return false
}

// GetMetadata returns Metadata
func (e *EmbeddedSignatureRequest) GetMetadata() map[string]string {
	if e != nil {
		return e.Metadata
	}
	return nil
}

// GetFormFieldsPerDocument returns FormFieldsPerDocument
func (e *EmbeddedSignatureRequest) GetFormFieldsPerDocument() [][]DocumentFormField {
	if e != nil {
		return e.FormFieldsPerDocument
	}
	return nil
}
