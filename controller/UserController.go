package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rudiarta/refactory_test/model"
	"github.com/rudiarta/refactory_test/service/user"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	GetUserList(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
}

func NewUserController(srv user.UserService) UserController {
	return &userController{srv}
}

type userController struct {
	srv user.UserService
}

func (c *userController) CreateUser(ctx *gin.Context) {
	if err := c.srv.CreateUser(model.User{FullName: "rudi"}); err != nil {
		log.Fatalf(err.Error())
	}

	ctx.JSON(200, gin.H{
		"message": "success",
	})
}

func (c *userController) GetUserList(ctx *gin.Context) {

}

func (c *userController) GetUserByID(ctx *gin.Context) {

}
