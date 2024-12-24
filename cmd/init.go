package main

import (
	"api-product-manager/config"
	"api-product-manager/controller"
	"api-product-manager/repository"
	"api-product-manager/routes"
	"api-product-manager/useCase"
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4/stdlib"
)

// Função para inicializar a conexão com o banco de dados
func initDB() (*sql.DB, *pgxpool.Pool) {
	dbConfig := config.GetDBConfig()
	pool, err := pgxpool.Connect(context.Background(), dbConfig)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	fmt.Println("Conectado ao banco de dados com sucesso!")

	db := stdlib.OpenDB(*pool.Config().ConnConfig)

	return db, pool
}


func initApp() (*gin.Engine, *sql.DB, *pgxpool.Pool) {
	db, pool := initDB()

	// Cria os repositórios
	productRepository := repository.NewProductRepository(db)

	// Cria os casos de uso
	productUseCase := usecase.NewProductUserCase(productRepository)

	// Cria os controladores
	productController := controller.NewProductController(productUseCase)

	// Configura as rotas
	server := gin.Default()
	router := routes.NewProductRoutes(server, productController)
	router.SetupRoutes()

	return server, db, pool
}
