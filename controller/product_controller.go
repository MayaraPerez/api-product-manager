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

func (p *ProductController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("product_id")
	if id == "" {
		response := model.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um numero valido",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.ProductUserCase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto nao foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (pc *ProductController) UpdateProduct(ctx *gin.Context) {
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

	receiveUpdateProduct := model.Product{}
	err = ctx.ShouldBindJSON(&receiveUpdateProduct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos na requisição"})
	}

	updateProduct, err := pc.ProductUserCase.UpdateProduct(receiveUpdateProduct.Name, receiveUpdateProduct.Price, productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}

	product := model.UpdateVerifyProduct(*updateProduct, receiveUpdateProduct)

	ctx.JSON(http.StatusOK, product)

}
func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	id_product := ctx.Param("product_id")

	if id_product == "" {
		mensage := model.Response {
			Message: "id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, mensage)
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

	deleteProduct, err := pc.ProductUserCase.DeleteProduct(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return 
	}
	
	ctx.JSON(http.StatusOK, deleteProduct)
}
	
	