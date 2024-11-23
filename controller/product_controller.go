package controller

import (
	"api-product-manager/model"
	usecase "api-product-manager/useCase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUserCase usecase.ProductUserCase
}

func NewProductController(usecase usecase.ProductUserCase) ProductController {
	return ProductController{
		ProductUserCase: usecase,
	}
}

func (pc *ProductController) GetProducts(ctx *gin.Context) {
	product, err := pc.ProductUserCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, product)
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	//instancio o model.product
	var product model.Product
	// tranformo JSON em STRUCT
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	// chamo o userCase e passo product e trato possivel erro
	insertProduct, err := pc.ProductUserCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return 
	}

	ctx.JSON(http.StatusOK, insertProduct)

}

func (pc *ProductController) GetProductById(ctx *gin.Context) {
	//pego id da requisicao
	id_product := ctx.Param("product_id")

	if id_product == "" {
		message := model.Response {
			Message: "id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, message)
		return 
	}

	productId, err := strconv.Atoi(id_product)
	if err != nil {
		message := model.Response {
			Message: "id do produto precisa ser um numero valido",
		}
		ctx.JSON(http.StatusBadRequest, message)
		return 
	}

	product, err := pc.ProductUserCase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return 
	}

	if product == nil {
		response := model.Response {
			Message: "produto não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return 
	}

	ctx.JSON(http.StatusOK, product)
	
}