package model

// EditEmbeddedTemplateRequest contains the request parameters for creating an embedded template draft
// Note: skip_me_now, cc_roles, attachments, allow_reassign, allow_css, editor_options, field_options, merge_field are unused and thus excluded
type EditEmbeddedTemplateRequest struct {
	TestMode     bool              `form_field:"test_mode"`
	ClientID     string            `form_field:"client_id"`
	FileURL      []string          `form_field:"file_url"`
	File         []string          `form_field:"file"`
	Title        string            `form_field:"title"`
	Subject      string            `form_field:"subject"`
	Message      string            `form_field:"message"`
	SignerRoles  []SignerRole      `form_field:"signer_roles"`
	Metadata     map[string]string `form_field:"metadata"`
	ShowPreview  bool              `form_field:"show_preview"`
	PreviewOnly  bool              `form_field:"preview_only"`
	CustomFields string            `form_field:"merge_fields"`
}

// GetTestMode returns TestMode
func (e *EditEmbeddedTemplateRequest) GetTestMode() bool {
	if e != nil {
		return e.TestMode
	}
	return false
}

// GetClientID returns ClientID
func (e *EditEmbeddedTemplateRequest) GetClientID() string {
	if e != nil {
		return e.ClientID
	}
	return ""
}

// GetFileURL returns FileURL
func (e *EditEmbeddedTemplateRequest) GetFileURL() []string {
	if e != nil {
		return e.FileURL
	}
	return nil
}

// GetFile returns File
func (e *EditEmbeddedTemplateRequest) GetFile() []string {
	if e != nil {
		return e.File
	}
	return nil
}

// GetTitle returns Title
func (e *EditEmbeddedTemplateRequest) GetTitle() string {
	if e != nil {
		return e.Title
	}
	return ""
}

// GetSubject returns Subject
func (e *EditEmbeddedTemplateRequest) GetSubject() string {
	if e != nil {
		return e.Subject
	}
	return ""
}

// GetMessage returns Message
func (e *EditEmbeddedTemplateRequest) GetMessage() string {
	if e != nil {
		return e.Message
	}
	return ""
}

// GetSignerRoles returns Signers
func (e *EditEmbeddedTemplateRequest) GetSignerRoles() []SignerRole {
	if e != nil {
		return e.SignerRoles
	}
	return nil
}

// GetMetadata returns Metadata
func (e *EditEmbeddedTemplateRequest) GetMetadata() map[string]string {
	if e != nil {
		return e.Metadata
	}
	return nil
}

// IsShowingPreview returns ShowPreview
func (e *EditEmbeddedTemplateRequest) IsShowingPreview() bool {
	if e != nil {
		return e.ShowPreview
	}
	return false
}

// GetCustomFields returns CustomFields
func (e *EditEmbeddedTemplateRequest) GetCustomFields() string {
	if e != nil {
		return e.CustomFields
	}
	return ""
}

// IsPreviewOnly returns PreviewOnly
func (e *EditEmbeddedTemplateRequest) IsPreviewOnly() bool {
	if e != nil {
		return e.PreviewOnly
	}
	return false
}
