package container

import (
	"github.com/rudiarta/refactory_test/config"
	"go.uber.org/dig"
)

func InjectDatabase(c *dig.Container) *dig.Container {
	if err := c.Provide(config.NewMysqlConfigurator); err != nil {
		panic(err)
	}

	return c
}
