package model

type Signature struct {
	SignatureID        string  `json:"signature_id"`         // Signature identifier.
	SignerEmailAddress string  `json:"signer_email_address"` // The email address of the signer.
	SignerName         string  `json:"signer_name"`          // The name of the signer.
	Order              int     `json:"order"`                // If signer order is assigned this is the 0-based index for this signer.
	StatusCode         string  `json:"status_code"`          // The current status of the signature. eg: awaiting_signature, signed, declined
	DeclineReason      string  `json:"decline_reason"`       // The reason provided by the signer for declining the request.
	SignedAt           int     `json:"signed_at"`            // Time that the document was signed or null.
	LastViewedAt       int     `json:"last_viewed_at"`       // The time that the document was last viewed by this signer or null.
	LastRemindedAt     int     `json:"last_reminded_at"`     // The time the last reminder email was sent to the signer or null.
	HasPin             bool    `json:"has_pin"`              // Boolean to indicate whether this signature requires a PIN to access.
	ReassignedBy       string  `json:"reassigned_by"`        // Email address of original signer who reassigned to this signer.
	ReassignmentReason string  `json:"reassignment_reason"`  // Reason provided by original signer who reassigned to this signer.
	Error              *string `json:"error"`                // Error message pertaining to this signer, or null.
}

// GetSignatureID returns SignatureID
func (s *Signature) GetSignatureID() string {
	if s != nil {
		return s.SignatureID
	}
	return ""
}

// GetSignerEmailAddress returns SignerEmailAddress
func (s *Signature) GetSignerEmailAddress() string {
	if s != nil {
		return s.SignerEmailAddress
	}
	return ""
}

// GetSignerName returns SignerName
func (s *Signature) GetSignerName() string {
	if s != nil {
		return s.SignerName
	}
	return ""
}

// GetOrder returns Order
func (s *Signature) GetOrder() int {
	if s != nil {
		return s.Order
	}
	return 0
}

// GetStatusCode returns StatusCode
func (s *Signature) GetStatusCode() string {
	if s != nil {
		return s.StatusCode
	}
	return ""
}

// GetDeclineReason returns DeclineReason
func (s *Signature) GetDeclineReason() string {
	if s != nil {
		return s.DeclineReason
	}
	return ""
}

// GetSignedAt returns SignedAt
func (s *Signature) GetSignedAt() int {
	if s != nil {
		return s.SignedAt
	}
	return 0
}

// GetLastViewedAt returns LastViewedAt
func (s *Signature) GetLastViewedAt() int {
	if s != nil {
		return s.LastViewedAt
	}
	return 0
}

// GetLastRemindedAt returns LastRemindedAt
func (s *Signature) GetLastRemindedAt() int {
	if s != nil {
		return s.LastRemindedAt
	}
	return 0
}

// GetHasPin returns HasPin
func (s *Signature) GetHasPin() bool {
	if s != nil {
		return s.HasPin
	}
	return false
}

// GetReassignedBy returns ReassignedBy
func (s *Signature) GetReassignedBy() string {
	if s != nil {
		return s.ReassignedBy
	}
	return ""
}

// GetReassignmentReason returns ReassignmentReason
func (s *Signature) GetReassignmentReason() string {
	if s != nil {
		return s.ReassignmentReason
	}
	return ""
}

// GetError returns Error
func (s *Signature) GetError() *string {
	if s != nil {
		return s.Error
	}
	return nil
}