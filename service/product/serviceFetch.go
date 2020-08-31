package product

import "github.com/rudiarta/refactory_test/model"

func (srv *productService) Fetch() (error, []model.Product) {
	err, result := srv.repo.GetAll()
	if err != nil {
		return err, nil
	}

	return nil, result
}
