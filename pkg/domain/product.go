package domain

import (
	"errors"
)

type Store struct {
	Name     string `json:"name" gorm:"column:name"`
	Seller   string `json:"seller" gorm:"column:seller"`
	Product  string `json:"product" gorm:"column:product"`
	Customer string `json:"customer" gorm:"column:customer"`
}

func (Store) TableName() string {
	return "products"
}

func (s Store) Validate() error {

	if len(s.Customer) > 15 {
		return errors.New("customer len's more than 15")
	}
	if len(s.Product) > 10 {
		return errors.New("must not be less then 10 more then 15")
	}
	if len(s.Seller) > 10 {
		return errors.New("must not be more then 10")
	}
	return nil
}
