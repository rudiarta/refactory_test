package repocontract

import "github.com/rudiarta/refactory_test/model"

type UserRepoContract interface {
	Save(entity model.User) error
	Fetch() (error, []model.User)
	FindByID(id int) (error, *model.User)
}
