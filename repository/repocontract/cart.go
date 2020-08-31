package repocontract

import "github.com/rudiarta/refactory_test/model"

type CartRepoContract interface {
	Save(entity model.Cart) error
	GetByUserID(id int) (error, []model.Cart)
}
