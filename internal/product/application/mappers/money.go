package mappers

import (
	"github.com/ldez/go-monorepo-example/product/domain/valueobjects"
)

func FromStringToMoney(priceString string) *valueobjects.Money {
	return &valueobjects.Money{
		Value:    1,
		Currency: valueobjects.Currency{},
	}
}
