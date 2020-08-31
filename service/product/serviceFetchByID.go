package product

import "github.com/rudiarta/refactory_test/model"

func (srv *productService) FetchByID(id int) (error, *model.Product) {
	err, result := srv.repo.GetByID(id)
	if err != nil {
		return err, nil
	}

	return nil, result
}
