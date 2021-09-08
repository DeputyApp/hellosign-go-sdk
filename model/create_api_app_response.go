package model

type CreateAPIAppResponse struct {
	APIApp *APIApp `json:"api_app"`
}

// GetAPIApp returns ApiApp
func (cr *CreateAPIAppResponse) GetAPIApp() *APIApp {
	if cr != nil {
		return cr.APIApp
	}
	return nil
}
