package model

type SignatureRequestResponse struct {
	SignatureRequest *SignatureRequest `json:"signature_request"`
}

// GetSignatureRequest returns SignatureRequest
func (sr *SignatureRequestResponse) GetSignatureRequest() *SignatureRequest {
	if sr != nil {
		return sr.SignatureRequest
	}
	return nil
}
