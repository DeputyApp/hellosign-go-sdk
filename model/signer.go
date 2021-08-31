package model

type Signer struct {
	Name  string `field:"name"`
	Email string `field:"email_address"`
	Order int    `field:"order"`
	Pin   string `field:"pin"`
}