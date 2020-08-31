package user

import (
	"github.com/rudiarta/refactory_test/model"
)

func (srv *userService) CreateUser(entity model.User) error {
	if err := srv.repo.Save(entity); err != nil {
		return err
	}
	return nil
}
