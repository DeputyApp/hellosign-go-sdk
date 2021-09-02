package model

type ListTemplatesResponse struct {
	ListInfo  *ListInfo   `json:"list_info"`
	Templates []*Template `json:"templates"`
}

// GetListInfo returns ListInfo
func (l *ListTemplatesResponse) GetListInfo() *ListInfo {
	if l != nil {
		return l.ListInfo
	}
	return nil
}

// GetTemplates returns Templates
func (l *ListTemplatesResponse) GetTemplates() []*Template {
	if l != nil {
		return l.Templates
	}
	return nil
}
