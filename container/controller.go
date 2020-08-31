package container

import (
	"github.com/rudiarta/refactory_test/controller"
	"go.uber.org/dig"
)

func InjectController(c *dig.Container) *dig.Container {
	if err := c.Provide(controller.NewUserController); err != nil {
		panic(err)
	}

	return c
}
