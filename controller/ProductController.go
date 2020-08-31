package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rudiarta/refactory_test/model"
	"github.com/rudiarta/refactory_test/service/product"
)

type ProductController interface {
	CreateProduct(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

func NewProductController(srv product.ProductService) ProductController {
	return &productController{srv}
}

type productController struct {
	srv product.ProductService
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	name := ctx.PostForm("name")
	variant := ctx.PostForm("variant")
	merchantID := ctx.PostForm("merchant_id")
	price := ctx.PostForm("price")
	status := false
	if ctx.PostForm("status") == "true" {
		status = true
	}

	data := model.Product{}
	data.Name = name
	data.Variant = variant
	data.MerchantID, _ = strconv.Atoi(merchantID)
	data.Price, _ = strconv.Atoi(price)
	data.Status = status

	if err := c.srv.Create(data); err != nil {
		ctx.JSON(422, gin.H{
			"status": "failed",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status": "succcess",
	})
	return
}

func (c *productController) GetAll(ctx *gin.Context) {
	err, result := c.srv.Fetch()
	if err != nil {
		ctx.JSON(422, gin.H{
			"status": "failed",
		})

		return
	}

	ctx.JSON(200, gin.H{
		"data": result,
	})
	return
}
