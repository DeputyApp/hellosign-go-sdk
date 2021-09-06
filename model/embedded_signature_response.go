package model

type EmbeddedSignatureResponse struct {
	Embedded *SignURLResponse `json:"embedded"`
}

// GetEmbedded returns Embedded
func (e *EmbeddedSignatureResponse) GetEmbedded() *SignURLResponse {
	if e != nil {
		return e.Embedded
	}
	return nil
}