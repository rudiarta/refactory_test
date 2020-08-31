package cart

import "github.com/rudiarta/refactory_test/model"

func (srv *cartService) GetByUserID(id int) (error, []model.Cart) {
	err, result := srv.repo.GetByUserID(id)
	if err != nil {
		return err, nil
	}

	return nil, result
}
