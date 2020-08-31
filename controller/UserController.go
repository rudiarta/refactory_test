package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rudiarta/refactory_test/model"
	"github.com/rudiarta/refactory_test/service/cart"
	"github.com/rudiarta/refactory_test/service/product"
	"github.com/rudiarta/refactory_test/service/user"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	GetUserList(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	AddProductToCartByUser(ctx *gin.Context)
}

func NewUserController(srv user.UserService, srvCart cart.CartService, srvProd product.ProductService) UserController {
	return &userController{srv, srvCart, srvProd}
}

type userController struct {
	srv     user.UserService
	srvCart cart.CartService
	srvProd product.ProductService
}

func (c *userController) CreateUser(ctx *gin.Context) {
	name := ctx.PostForm("fullname")
	if err := c.srv.CreateUser(model.User{FullName: name}); err != nil {
		ctx.JSON(200, gin.H{
			"message": "failed",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"message": "success",
	})

	return
}

func (c *userController) GetUserList(ctx *gin.Context) {
	_, result := c.srv.GetUser()

	ctx.JSON(200, gin.H{
		"data": result,
	})

	return
}

func (c *userController) GetUserByID(ctx *gin.Context) {
	id := ctx.PostForm("id")
	idInt, _ := strconv.Atoi(id)
	_, result := c.srv.FindByID(idInt)

	err, cart := c.srvCart.GetByUserID(result.ID)
	data := make([]model.GetCart, 0)
	d := model.GetCart{}
	d.User = *result
	if err == nil {
		for _, y := range cart {
			_, x := c.srvProd.FetchByID(y.ProductID)
			d.Product = append(d.Product, *x)
		}
	}

	data = append(data, d)

	ctx.JSON(200, gin.H{
		"data": data,
	})

	return
}

func (c *userController) AddProductToCartByUser(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.PostForm("user_id"))
	productID, _ := strconv.Atoi(ctx.PostForm("product_id"))
	fmt.Println(userID, productID)
	data := model.Cart{}
	data.ProductID = productID
	data.UserID = userID

	if err := c.srvCart.Create(data); err != nil {
		ctx.JSON(200, gin.H{
			"message": "failed",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"message": "success",
	})

	return
}
