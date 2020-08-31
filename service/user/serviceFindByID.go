package user

import "github.com/rudiarta/refactory_test/model"

func (srv *userService) FindByID(id int) (error, *model.User) {
	result := new(model.User)
	return nil, result
}
