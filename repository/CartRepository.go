package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rudiarta/refactory_test/model"
	"github.com/rudiarta/refactory_test/repository/repocontract"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) repocontract.CartRepoContract {
	return &cartRepository{db}
}

func (repo *cartRepository) Save(entity model.Cart) error {
	result := repo.db.Create(&entity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *cartRepository) GetByUserID(id int) (error, []model.Cart) {
	data := make([]model.Cart, 0)
	result := repo.db.Where("user_id = ?", id).Find(&data)
	if result.Error != nil {
		return result.Error, nil
	}

	return nil, data
}
