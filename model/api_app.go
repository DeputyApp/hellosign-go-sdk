package model

// APIApp contains information about an API App
// Note: we ignore options and oauth here
type APIApp struct {
	ClientID     string   `json:"client_id"`
	CreatedAt    int      `json:"created_at"`
	Name         string   `json:"name"`
	Domain       []string `json:"domain"`
	CallbackURL  string   `json:"callback_url"`
	IsApproved   bool     `json:"is_approved"`
	OwnerAccount *Account `json:"owner_account"`
	// WhiteLabelingOptions is an array of elements and values serialized to a string
	WhiteLabelingOptions string `json:"white_labeling_options"`
}

// GetClientID returns ClientID
func (a *APIApp) GetClientID() string {
	if a != nil {
		return a.ClientID
	}
	return ""
}

// GetCreatedAt returns CreatedAt
func (a *APIApp) GetCreatedAt() int {
	if a != nil {
		return a.CreatedAt
	}
	return 0
}

// GetName returns Name
func (a *APIApp) GetName() string {
	if a != nil {
		return a.Name
	}
	return ""
}

// GetDomain returns Domain
func (a *APIApp) GetDomain() []string {
	if a != nil {
		return a.Domain
	}
	return nil
}

// GetCallbackURL returns CallbackURL
func (a *APIApp) GetCallbackURL() string {
	if a != nil {
		return a.CallbackURL
	}
	return ""
}

// GetIsApproved returns IsApproved
func (a *APIApp) GetIsApproved() bool {
	if a != nil {
		return a.IsApproved
	}
	return false
}

// GetOwnerAccount returns OwnerAccount
func (a *APIApp) GetOwnerAccount() *Account {
	if a != nil {
		return a.OwnerAccount
	}
	return nil
}

// GetWhiteLabelingOptions returns WhiteLabelingOptions
func (a *APIApp) GetWhiteLabelingOptions() string {
	if a != nil {
		return a.WhiteLabelingOptions
	}
	return ""
}
