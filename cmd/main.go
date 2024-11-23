package main

import (
	"api-product-manager/controller"
	"api-product-manager/db"
	"api-product-manager/repository"
	usecase "api-product-manager/useCase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	// Rota de exemplo para verificar se o servidor está rodando
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Conectando ao banco de dados
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer dbConnection.Close()

	// instanciando o repository do produto
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Instanciando o useCase do produto
	ProductUseCase := usecase.NewProductUserCase(ProductRepository)

	// Instanciando o controlador de produto
	ProductController := controller.NewProductController(ProductUseCase)

	// Registrando a rota `/products`
	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:product_id", ProductController.GetProductById)

	// Iniciando o servidor na porta 8000
	server.Run(":8000")
}
