package user

import (
	"github.com/rudiarta/refactory_test/model"
	"github.com/rudiarta/refactory_test/repository/repocontract"
)

type UserService interface {
	CreateUser(entity model.User) error
	GetUser() (error, []model.User)
	FindByID(id int) (error, *model.User)
}

type userService struct {
	repo repocontract.UserRepoContract
}

func NewUserService(repo repocontract.UserRepoContract) UserService {
	return &userService{repo}
}
