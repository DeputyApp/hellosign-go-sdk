package model

type ListInfo struct {
	NumPages   int `json:"num_pages"`   // Total number of pages available
	NumResults int `json:"num_results"` // Total number of objects available
	Page       int `json:"page"`        // Number of the page being returned
	PageSize   int `json:"page_size"`   // Objects returned per page
}

// GetNumPages returns NumPages
func (l *ListInfo) GetNumPages() int {
	if l != nil {
		return l.NumPages
	}
	return 0
}

// GetNumResults returns NumResults
func (l *ListInfo) GetNumResults() int {
	if l != nil {
		return l.NumResults
	}
	return 0
}

// GetPage returns Page
func (l *ListInfo) GetPage() int {
	if l != nil {
		return l.Page
	}
	return 0
}

// GetPageSize returns PageSize
func (l *ListInfo) GetPageSize() int {
	if l != nil {
		return l.PageSize
	}
	return 0
}