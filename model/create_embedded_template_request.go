package model

// CreateEmbeddedTemplateRequest contains the request parameters for creating an embedded template draft
// Note: skip_me_now, cc_roles, attachments, allow_reassign, allow_css, editor_options, field_options, merge_field are unused and thus excluded
type CreateEmbeddedTemplateRequest struct {
	TestMode    bool              `form_field:"test_mode"`
	ClientID    string            `form_field:"client_id"`
	FileURL     []string          `form_field:"file_url"`
	File        []string          `form_field:"file"`
	Title       string            `form_field:"title"`
	Subject     string            `form_field:"subject"`
	Message     string            `form_field:"message"`
	SignerRoles []SignerRole      `form_field:"signer_roles"`
	Metadata    map[string]string `form_field:"metadata"`
	ShowPreview bool              `form_field:"show_preview"`
}

// GetTestMode returns TestMode
func (e *CreateEmbeddedTemplateRequest) GetTestMode() bool {
	if e != nil {
		return e.TestMode
	}
	return false
}

// GetClientID returns ClientID
func (e *CreateEmbeddedTemplateRequest) GetClientID() string {
	if e != nil {
		return e.ClientID
	}
	return ""
}

// GetFileURL returns FileURL
func (e *CreateEmbeddedTemplateRequest) GetFileURL() []string {
	if e != nil {
		return e.FileURL
	}
	return nil
}

// GetFile returns File
func (e *CreateEmbeddedTemplateRequest) GetFile() []string {
	if e != nil {
		return e.File
	}
	return nil
}

// GetTitle returns Title
func (e *CreateEmbeddedTemplateRequest) GetTitle() string {
	if e != nil {
		return e.Title
	}
	return ""
}

// GetSubject returns Subject
func (e *CreateEmbeddedTemplateRequest) GetSubject() string {
	if e != nil {
		return e.Subject
	}
	return ""
}

// GetMessage returns Message
func (e *CreateEmbeddedTemplateRequest) GetMessage() string {
	if e != nil {
		return e.Message
	}
	return ""
}

// GetSignerRoles returns Signers
func (e *CreateEmbeddedTemplateRequest) GetSignerRoles() []SignerRole {
	if e != nil {
		return e.SignerRoles
	}
	return nil
}

// GetMetadata returns Metadata
func (e *CreateEmbeddedTemplateRequest) GetMetadata() map[string]string {
	if e != nil {
		return e.Metadata
	}
	return nil
}

func (e *CreateEmbeddedTemplateRequest) IsShowingPreview() bool {
	if e != nil {
		return e.ShowPreview
	}
	return false
}
