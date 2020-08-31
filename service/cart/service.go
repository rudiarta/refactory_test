package cart

import (
	"github.com/rudiarta/refactory_test/model"
	"github.com/rudiarta/refactory_test/repository/repocontract"
)

type CartService interface {
	Create(entity model.Cart) error
	GetByUserID(id int) (error, []model.Cart)
}

type cartService struct {
	repo repocontract.CartRepoContract
}

func NewCartService(repo repocontract.CartRepoContract) CartService {
	return &cartService{repo}
}
