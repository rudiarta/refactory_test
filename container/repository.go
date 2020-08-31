package container

import (
	"github.com/rudiarta/refactory_test/repository"
	"go.uber.org/dig"
)

func InjectRepository(c *dig.Container) *dig.Container {

	if err := c.Provide(repository.NewUserRepository); err != nil {
		panic(err)
	}

	if err := c.Provide(repository.NewCartRepository); err != nil {
		panic(err)
	}

	if err := c.Provide(repository.NewProductRepository); err != nil {
		panic(err)
	}

	return c
}
