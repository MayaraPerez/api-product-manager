package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Init()  {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Erro ao carregar o arquivo .env")
		panic(err)
	}
}

func GetDBConfig() string {
	dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",  
						dbHost, dbPort, dbUser, dbPassword, dbName)
}
