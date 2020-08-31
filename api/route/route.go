package route

import (
	"github.com/gin-gonic/gin"
	"github.com/rudiarta/refactory_test/controller"
)

func InvokeRoute(
	r *gin.Engine,
	user controller.UserController,
	product controller.ProductController) {

	r.POST("/create-user", user.CreateUser)
	r.GET("/list-user", user.GetUserList)
	r.GET("/get-user-id", user.GetUserByID)

	r.POST("/create-product", product.CreateProduct)
	r.POST("/add-cart", user.AddProductToCartByUser)
	r.GET("/list-product", product.GetAll)

	r.Run(":8080")
}
