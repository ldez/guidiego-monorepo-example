package usecases

import (
	"github.com/ldez/go-monorepo-example/product/application/dtos"
	"github.com/ldez/go-monorepo-example/product/application/mappers"
	"github.com/ldez/go-monorepo-example/product/domain/entities"
	"github.com/ldez/go-monorepo-example/product/domain/repositories"
)

type SaveProductUseCase struct {
	Repository repositories.IProductRepository
}

func (usecase *SaveProductUseCase) Execute(productDto dtos.Product) error {
	product := entities.Product{
		Id:          "",
		Name:        productDto.Name,
		Description: productDto.Description,
		Price:       mappers.FromStringToMoney(productDto.Price),
	}

	err := product.Validate()
	if err != nil {
		return err
	}

	usecase.Repository.Save(&product)
	return nil
}
