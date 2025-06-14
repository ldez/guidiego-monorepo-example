package repositories

import "github.com/ldez/go-monorepo-example/product/domain/entities"

type IProductRepository interface {
	List() []*entities.Product
	Save(product *entities.Product)
}
