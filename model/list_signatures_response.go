package model

type ListSignaturesResponse struct {
	ListInfo          *ListInfo           `json:"list_info"`
	SignatureRequests []*SignatureRequest `json:"signature_requests"`
}

// GetListInfo returns ListInfo
func (l *ListSignaturesResponse) GetListInfo() *ListInfo {
	if l != nil {
		return l.ListInfo
	}
	return nil
}

// GetSignatureRequests returns SignatureRequests
func (l *ListSignaturesResponse) GetSignatureRequests() []*SignatureRequest {
	if l != nil {
		return l.SignatureRequests
	}
	return nil
}
