package model

type SignatureRequest struct {
	TestMode              bool                     `json:"test_mode"`               // Whether this is a test signature request. Test requests have no legal value. Defaults to 0.
	SignatureRequestID    string                   `json:"signature_request_id"`    // The id of the SignatureRequest.
	RequesterEmailAddress string                   `json:"requester_email_address"` // The email address of the initiator of the SignatureRequest.
	Title                 string                   `json:"title"`                   // The title the specified Account uses for the SignatureRequest.
	OriginalTitle         string                   `json:"original_title"`          // Default Label for account.
	Subject               string                   `json:"subject"`                 // The subject in the email that was initially sent to the signers.
	Message               string                   `json:"message"`                 // The custom message in the email that was initially sent to the signers.
	Metadata              map[string]interface{}   `json:"metadata"`                // The metadata attached to the signature request.
	CreatedAt             int                      `json:"created_at"`              // Time the signature request was created.
	IsComplete            bool                     `json:"is_complete"`             // Whether or not the SignatureRequest has been fully executed by all signers.
	IsDeclined            bool                     `json:"is_declined"`             // Whether or not the SignatureRequest has been declined by a signer.
	HasError              bool                     `json:"has_error"`               // Whether or not an error occurred (either during the creation of the SignatureRequest or during one of the signings).
	FilesURL              string                   `json:"files_url"`               // The URL where a copy of the request's documents can be downloaded.
	SigningURL            string                   `json:"signing_url"`             // The URL where a signer, after authenticating, can sign the documents. This should only be used by users with existing HelloSign accounts as they will be required to log in before signing.
	DetailsURL            string                   `json:"details_url"`             // The URL where the requester and the signers can view the current status of the SignatureRequest.
	CCEmailAddress        []*string                `json:"cc_email_addresses"`      // A list of email addresses that were CCed on the SignatureRequest. They will receive a copy of the final PDF once all the signers have signed.
	SigningRedirectURL    string                   `json:"signing_redirect_url"`    // The URL you want the signer redirected to after they successfully sign.
	CustomFields          []map[string]interface{} `json:"custom_fields"`           // An array of Custom Field objects containing the name and type of each custom field.
	ResponseData          []*ResponseData          `json:"response_data"`           // An array of form field objects containing the name, value, and type of each textbox or checkmark field filled in by the signers.
	Signatures            []*Signature             `json:"signatures"`              // An array of signature objects, 1 for each signer.
	Warnings              []*Warning               `json:"warnings"`                // An array of warning objects.
	TemplateIDs           []string                 `json:"template_ids"`
	ClientID              string                   `json:"client_id"`
}

// GetTestMode returns TestMode
func (s *SignatureRequest) GetTestMode() bool {
	if s != nil {
		return s.TestMode
	}
	return false
}

// GetSignatureRequestID returns SignatureRequestID
func (s *SignatureRequest) GetSignatureRequestID() string {
	if s != nil {
		return s.SignatureRequestID
	}
	return ""
}

// GetRequesterEmailAddress returns RequesterEmailAddress
func (s *SignatureRequest) GetRequesterEmailAddress() string {
	if s != nil {
		return s.RequesterEmailAddress
	}
	return ""
}

// GetTitle returns Title
func (s *SignatureRequest) GetTitle() string {
	if s != nil {
		return s.Title
	}
	return ""
}

// GetOriginalTitle returns OriginalTitle
func (s *SignatureRequest) GetOriginalTitle() string {
	if s != nil {
		return s.OriginalTitle
	}
	return ""
}

// GetSubject returns Subject
func (s *SignatureRequest) GetSubject() string {
	if s != nil {
		return s.Subject
	}
	return ""
}

// GetMessage returns Message
func (s *SignatureRequest) GetMessage() string {
	if s != nil {
		return s.Message
	}
	return ""
}

// GetMetadata returns Metadata
func (s *SignatureRequest) GetMetadata() map[string]interface{} {
	if s != nil {
		return s.Metadata
	}
	return nil
}

// GetCreatedAt returns CreatedAt
func (s *SignatureRequest) GetCreatedAt() int {
	if s != nil {
		return s.CreatedAt
	}
	return 0
}

// GetIsComplete returns IsComplete
func (s *SignatureRequest) GetIsComplete() bool {
	if s != nil {
		return s.IsComplete
	}
	return false
}

// GetIsDeclined returns IsDeclined
func (s *SignatureRequest) GetIsDeclined() bool {
	if s != nil {
		return s.IsDeclined
	}
	return false
}

// GetHasError returns HasError
func (s *SignatureRequest) GetHasError() bool {
	if s != nil {
		return s.HasError
	}
	return false
}

// GetFilesURL returns FilesURL
func (s *SignatureRequest) GetFilesURL() string {
	if s != nil {
		return s.FilesURL
	}
	return ""
}

// GetSigningURL returns SigningURL
func (s *SignatureRequest) GetSigningURL() string {
	if s != nil {
		return s.SigningURL
	}
	return ""
}

// GetDetailsURL returns DetailsURL
func (s *SignatureRequest) GetDetailsURL() string {
	if s != nil {
		return s.DetailsURL
	}
	return ""
}

// GetCCEmailAddress returns CCEmailAddress
func (s *SignatureRequest) GetCCEmailAddress() []*string {
	if s != nil {
		return s.CCEmailAddress
	}
	return nil
}

// GetSigningRedirectURL returns SigningRedirectURL
func (s *SignatureRequest) GetSigningRedirectURL() string {
	if s != nil {
		return s.SigningRedirectURL
	}
	return ""
}

// GetCustomFields returns CustomFields
func (s *SignatureRequest) GetCustomFields() []map[string]interface{} {
	if s != nil {
		return s.CustomFields
	}
	return nil
}

// GetResponseData returns ResponseData
func (s *SignatureRequest) GetResponseData() []*ResponseData {
	if s != nil {
		return s.ResponseData
	}
	return nil
}

// GetSignatures returns Signatures
func (s *SignatureRequest) GetSignatures() []*Signature {
	if s != nil {
		return s.Signatures
	}
	return nil
}

// GetWarnings returns Warnings
func (s *SignatureRequest) GetWarnings() []*Warning {
	if s != nil {
		return s.Warnings
	}
	return nil
}

// GetTemplateIDs returns TemplateIDs
func (s *SignatureRequest) GetTemplateIDs() []string {
	if s != nil {
		return s.TemplateIDs
	}
	return nil
}

// GetClientID returns ClientID
func (s *SignatureRequest) GetClientID() string {
	if s != nil {
		return s.ClientID
	}
	return ""
}
