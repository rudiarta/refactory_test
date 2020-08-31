package user

import "github.com/rudiarta/refactory_test/model"

func (srv *userService) GetUser() (error, []model.User) {
	result := make([]model.User, 0)
	return nil, result
}
