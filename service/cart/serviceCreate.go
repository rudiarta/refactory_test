package cart

import "github.com/rudiarta/refactory_test/model"

func (srv *cartService) Create(entity model.Cart) error {
	if err := srv.repo.Save(entity); err != nil {
		return err
	}

	return nil
}
