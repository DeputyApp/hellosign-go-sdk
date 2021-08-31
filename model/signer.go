package model

type Signer struct {
	Name  string `field:"name"`
	Email string `field:"email_address"`
	Order int    `field:"order"`
	Pin   string `field:"pin"`
}

// GetName returns Signer's Name
func (s *Signer) GetName() string {
	if s != nil {
		return s.Name
	}
	return ""
}

// GetEmail returns Signer's Email
func (s *Signer) GetEmail() string {
	if s != nil {
		return s.Email
	}
	return ""
}

// GetOrder returns Signer's Order
func (s *Signer) GetOrder() int {
	if s != nil {
		return s.Order
	}
	return 0
}

// GetPin returns Signer's Pin
func (s *Signer) GetPin() string {
	if s != nil {
		return s.Pin
	}
	return ""
}