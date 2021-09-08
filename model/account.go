package model

// Account contains information about an account and its settings
// Note: we ignore is_paid_hs, is_paid_hf, quotas, role_code
type Account struct {
	AccountID    string `json:"account_id"`
	EmailAddress string `json:"email_address"`
}

// GetAccountID returns AccountID
func (a *Account) GetAccountID() string {
	if a != nil {
		return a.AccountID
	}
	return ""
}

// GetEmailAddress returns EmailAddress
func (a *Account) GetEmailAddress() string {
	if a != nil {
		return a.EmailAddress
	}
	return ""
}