package product

import (
	"github.com/rudiarta/refactory_test/model"
	"github.com/rudiarta/refactory_test/repository/repocontract"
)

type productService struct {
	repo repocontract.ProductRepoContract
}

func NewProductService(repo repocontract.ProductRepoContract) ProductService {
	return &productService{repo}
}

func (srv *productService) Create(entity model.Product) error {
	if err := srv.repo.Save(entity); err != nil {
		return err
	}

	return nil
}
