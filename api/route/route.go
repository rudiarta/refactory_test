package route

import (
	"github.com/gin-gonic/gin"
	"github.com/rudiarta/refactory_test/controller"
)

func InvokeRoute(
	r *gin.Engine,
	user controller.UserController,
) {

	r.POST("/create-user", user.CreateUser)

	r.Run(":8080")
}
