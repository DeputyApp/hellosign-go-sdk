package model

type CreateApiAppRequest struct {
	Name                 string   `json:"name" form_field:"name"`
	Domains              []string `json:"domain" form_field:"domains"`
	CallbackURL          string   `json:"callback_url" form_field:"callback_url"`
	CustomLogoFile       string   `json:"custom_logo_file" form_field:"custom_logo_file"`
	WhiteLabelingOptions string   `json:"white_labeling_options" form_field:"white_labeling_options"`
}

// GetName returns Name
func (a *CreateApiAppRequest) GetName() string {
	if a != nil {
		return a.Name
	}
	return ""
}

// GetDomains returns Domains
func (a *CreateApiAppRequest) GetDomains() []string {
	if a != nil {
		return a.Domains
	}
	return nil
}

// GetCallbackURL returns CallbackURL
func (a *CreateApiAppRequest) GetCallbackURL() string {
	if a != nil {
		return a.CallbackURL
	}
	return ""
}

// GetCustomLogoFile returns CustomLogoFile
func (a *CreateApiAppRequest) GetCustomLogoFile() string {
	if a != nil {
		return a.CustomLogoFile
	}
	return ""
}

// GetWhiteLabelingOptions returns WhiteLabelingOptions
func (a *CreateApiAppRequest) GetWhiteLabelingOptions() string {
	if a != nil {
		return a.WhiteLabelingOptions
	}
	return ""
}
