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
	// FieldOptions map[string]string `form_field:"field_options"``
	// AllowDecline          int                   `form_field:"allow_decline"`
	// AllowReassign         int                   `form_field:"allow_reassign"`
	// Attachments            []Attachment `form_field:"attachments"`
}