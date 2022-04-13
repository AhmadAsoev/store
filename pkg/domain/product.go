package domain

import "errors"

type Store struct {
	Name     string `json:"name"`
	Seller   string `json:"seller"`
	Product  string `json:"product"`
	Customer string `json:"customer"`
}

func (s Store) Validate() error {
	if s.Name == "" {
		return errors.New("name must not empty")
	}

	if len(s.Customer) > 15 {
		return errors.New("customer len's more than 15")
	}
	if len(s.Product) > 10 && len(s.Product) < 15 {
		return errors.New("must not be less then 10 more then 15")
	}
	if len(s.Seller) > 10 {
		return errors.New("must not be more then 10")
	}
	return nil
}
