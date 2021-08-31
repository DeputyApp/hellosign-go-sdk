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