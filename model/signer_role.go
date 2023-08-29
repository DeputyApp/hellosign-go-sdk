package model

type SignerRole struct {
	Name  string `field:"name"`
	Order int    `field:"order"`
}

// GetName returns the Name of the signer
func (s *SignerRole) GetName() string {
	if s != nil {
		return s.Name
	}
	return ""
}

// GetOrder returns the Order of the signer
func (s *SignerRole) GetOrder() int {
	if s != nil {
		return s.Order
	}
	return 0
}
