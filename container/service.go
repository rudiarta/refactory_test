package container

import (
	"github.com/rudiarta/refactory_test/service/user"
	"go.uber.org/dig"
)

func InjectService(c *dig.Container) *dig.Container {
	if err := c.Provide(user.NewUserService); err != nil {
		panic(err)
	}

	return c
}
