package usecases

import (
	"github.com/ldez/go-monorepo-example/product/application/dtos"
	"github.com/ldez/go-monorepo-example/product/application/mappers"
	"github.com/ldez/go-monorepo-example/product/domain/repositories"
)

type ListProductUseCase struct {
	Repository repositories.IProductRepository
}

func (usecase *ListProductUseCase) Execute() []*dtos.Product {
	products := usecase.Repository.List()
	productsDtos := make([]*dtos.Product, 0, len(products))
	for _, product := range products {
		productsDtos = append(productsDtos, mappers.FromProductDomainToDTO(product))
	}

	return productsDtos
}
