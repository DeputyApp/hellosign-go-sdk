package model

type ListResponse struct {
	ListInfo          *ListInfo           `json:"list_info"`
	SignatureRequests []*SignatureRequest `json:"signature_requests"`
}
