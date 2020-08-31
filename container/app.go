package container

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	c := dig.New()
	c = InjectConfigurator(c)
	c = InjectDatabase(c)
	c = InjectRepository(c)
	c = InjectService(c)
	c = InjectController(c)

	return c
}

func InjectConfigurator(c *dig.Container) *dig.Container {
	if err := c.Provide(func() *gin.Engine {
		return gin.Default()
	}); err != nil {
		panic(err)
	}

	return c
}
