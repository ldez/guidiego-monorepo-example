package repositories

import "github.com/ldez/go-monorepo-example/product/domain/entities"

type InMemoryProductRepository struct {
	data map[string]*entities.Product
}

func (repo *InMemoryProductRepository) List() []*entities.Product {
	productsSlice := make([]*entities.Product, 0, len(repo.data))
	for _, product := range repo.data {
		productsSlice = append(productsSlice, product)
	}

	return productsSlice
}

func (repo *InMemoryProductRepository) Save(product *entities.Product) {
	repo.data[product.Id] = product
}
