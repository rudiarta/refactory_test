package user

import "github.com/rudiarta/refactory_test/model"

func (srv *userService) FindByID(id int) (error, *model.User) {
	result := new(model.User)
	_, result = srv.repo.FindByID(id)
	return nil, result
}
