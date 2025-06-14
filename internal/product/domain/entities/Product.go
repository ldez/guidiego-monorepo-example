// Package entities hold entities for Product Domain
package entities

import (
	"errors"
	"fmt"

	"github.com/ldez/go-monorepo-example/product/domain/valueobjects"
)

type Product struct {
	Id          string
	Price       *valueobjects.Money
	Name        string
	Description string
}

func emptyFieldError(field string) error {
	return errors.New(
		fmt.Sprintf("Cannot accept %s as empty string", field),
	)
}

func (product *Product) Validate() error {
	if product.Id == "" {
		return emptyFieldError("id")
	}

	if product.Description == "" {
		return emptyFieldError("description")
	}

	if product.Name == "" {
		return emptyFieldError("name")
	}

	if priceError := product.Price.Validate(); priceError != nil {
		return priceError
	}

	return nil
}
