package model

type ListResponse struct {
	ListInfo          *ListInfo           `json:"list_info"`
	SignatureRequests []*SignatureRequest `json:"signature_requests"`
}

// GetListInfo returns ListInfo
func (l *ListResponse) GetListInfo() *ListInfo {
	if l != nil {
		return l.ListInfo
	}
	return nil
}

// GetSignatureRequests returns SignatureRequests
func (l *ListResponse) GetSignatureRequests() []*SignatureRequest {
	if l != nil {
		return l.SignatureRequests
	}
	return nil
}