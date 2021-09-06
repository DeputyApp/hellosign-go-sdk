package model

type EmbeddedTemplateResponse struct {
	Embedded *EmbeddedTemplateEditURL `json:"embedded"`
}

// GetEmbedded returns Embedded
func (e *EmbeddedTemplateResponse) GetEmbedded() *EmbeddedTemplateEditURL {
	if e != nil {
		return e.Embedded
	}
	return nil
}
