package product

import "github.com/rudiarta/refactory_test/model"

type ProductService interface {
	Create(entity model.Product) error
	Fetch() (error, []model.Product)
	FetchByID(id int) (error, *model.Product)
}
