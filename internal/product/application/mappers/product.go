package mappers

import (
	"github.com/ldez/go-monorepo-example/product/application/dtos"
	"github.com/ldez/go-monorepo-example/product/domain/entities"
)

func FromProductDomainToDTO(product *entities.Product) *dtos.Product {
	return &dtos.Product{
		Id:          product.Id,
		Price:       product.Price.FormatString(),
		Name:        product.Name,
		Description: product.Description,
	}
}
