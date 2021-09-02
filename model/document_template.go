package model

// Template contains information about the templates
// Note: we are leaving out cc_roles, accounts
type Template struct {
	TemplateID  string            `json:"template_id"`  // A Template unique identifier.
	Title       string            `json:"title"`        // The title of the template.
	Message     string            `json:"message"`      // The default message that will be sent to signers
	Metadata    map[string]string `json:"metadata"`     // The metadata attached to the template.
	SignerRoles []SignerRole      `json:"signer_roles"` // The current status of the signature. eg: awaiting_signature, signed, declined
	Documents   []Document        `json:"documents"`    // A collection of document that is associated with this template
	IsCreator   bool              `json:"is_creator"`
	IsEmbedded  bool              `json:"is_embedded"`  // True if the template was created using an embedded flow
	CanEdit     bool              `json:"can_edit"`
	IsLocked    bool              `json:"is_locked"`
}

// GetTemplateID returns TemplateID
func (t *Template) GetTemplateID() string {
	if t != nil {
		return t.TemplateID
	}
	return ""
}

// GetTitle returns Title
func (t *Template) GetTitle() string {
	if t != nil {
		return t.Title
	}
	return ""
}

// GetMessage returns Message
func (t *Template) GetMessage() string {
	if t != nil {
		return t.Message
	}
	return ""
}

// GetMetadata returns Metadata
func (t *Template) GetMetadata() map[string]string {
	if t != nil {
		return t.Metadata
	}
	return nil
}

// GetSignerRoles returns SignerRoles
func (t *Template) GetSignerRoles() []SignerRole {
	if t != nil {
		return t.SignerRoles
	}
	return nil
}

// GetDocuments returns Documents
func (t *Template) GetDocuments() []Document {
	if t != nil {
		return t.Documents
	}
	return nil
}

// GetIsCreator returns IsCreator
func (t *Template) GetIsCreator() bool {
	if t != nil {
		return t.IsCreator
	}
	return false
}

// GetIsEmbedded returns IsEmbedded
func (t *Template) GetIsEmbedded() bool {
	if t != nil {
		return t.IsEmbedded
	}
	return false
}

// GetCanEdit returns CanEdit
func (t *Template) GetCanEdit() bool {
	if t != nil {
		return t.CanEdit
	}
	return false
}

// GetIsLocked returns IsLocked
func (t *Template) GetIsLocked() bool {
	if t != nil {
		return t.IsLocked
	}
	return false
}
