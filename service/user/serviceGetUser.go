package user

import "github.com/rudiarta/refactory_test/model"

func (srv *userService) GetUser() (error, []model.User) {
	_, result := srv.repo.Fetch()
	return nil, result
}
