package model

type CreateApiAppRequest struct {
	Name                 string `json:"name" form_field:"name"`
	Domain               string `json:"domain" form_field:"domain"`
	CallbackURL          string `json:"callback_url" form_field:"callback_url"`
	CustomLogoFile       string `json:"custom_logo_file" form_field:"custom_logo_file"`
	WhiteLabelingOptions string `json:"white_labeling_options" form_field:"white_labeling_options"`
}

// GetName returns Name
func (a *CreateApiAppRequest) GetName() string {
	if a != nil {
		return a.Name
	}
	return ""
}

// GetDomain returns Domain
func (a *CreateApiAppRequest) GetDomain() string {
	if a != nil {
		return a.Domain
	}
	return ""
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