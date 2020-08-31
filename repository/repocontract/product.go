package repocontract

import "github.com/rudiarta/refactory_test/model"

type ProductRepoContract interface {
	Save(entity model.Product) error
	GetAll() (error, []model.Product)
	GetByID(id int) (error, *model.Product)
}
