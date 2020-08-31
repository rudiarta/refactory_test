package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/rudiarta/refactory_test/model"
	"github.com/rudiarta/refactory_test/repository/repocontract"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repocontract.ProductRepoContract {
	return &productRepository{db}
}

func (repo *productRepository) Save(entity model.Product) error {
	result := repo.db.Create(&entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *productRepository) GetAll() (error, []model.Product) {
	data := make([]model.Product, 0)
	result := repo.db.Find(&data)

	if result.Error != nil {
		return result.Error, nil
	}

	return nil, data
}

func (repo *productRepository) GetByID(id int) (error, *model.Product) {
	data := new(model.Product)

	result := repo.db.First(&data, id)

	if result.Error != nil {
		return result.Error, nil
	}

	return nil, data
}
