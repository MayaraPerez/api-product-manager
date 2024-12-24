package routes

import (
	"api-product-manager/controller"

	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	r		   *gin.Engine
	controller controller.ProductController
}

func NewProductRoutes(r *gin.Engine, controller controller.ProductController) *ProductRoutes {
	return &ProductRoutes{
		r: 			r,
		controller: controller,
	}
}

func (pr *ProductRoutes) SetupRoutes() {
	pr.r.GET("/products", pr.controller.GetProducts)
	pr.r.POST("/product", pr.controller.CreateProduct)
	pr.r.GET("/product/:product_id", pr.controller.GetProductById)
	pr.r.PUT("/product/:product_id", pr.controller.UpdateProduct)
	pr.r.DELETE("/product/:product_id", pr.controller.DeleteProduct)
}
