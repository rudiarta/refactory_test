package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rudiarta/refactory_test/model"
	"github.com/rudiarta/refactory_test/repository/repocontract"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repocontract.UserRepoContract {
	return &userRepository{db}
}

func (repo *userRepository) Save(entity model.User) error {
	result := repo.db.Create(&entity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *userRepository) Fetch() (error, []model.User) {
	result := make([]model.User, 0)
	repo.db.Find(&result)
	return nil, result
}

func (repo *userRepository) FindByID(id int) (error, *model.User) {

	result := new(model.User)

	repo.db.First(&result, id)

	return nil, result
}
