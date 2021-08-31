package model

type ListInfo struct {
	NumPages   int `json:"num_pages"`   // Total number of pages available
	NumResults int `json:"num_results"` // Total number of objects available
	Page       int `json:"page"`        // Number of the page being returned
	PageSize   int `json:"page_size"`   // Objects returned per page
}
