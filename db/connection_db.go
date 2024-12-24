package db

import (
	"api-product-manager/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)


func ConnectDB() (*sql.DB, error) {
	// Configuração da string de conexão
	psqInfo:= config.GetDBConfig()

	// Abrindo conexão com o postgres
	db, err := sql.Open("postgres", psqInfo)
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados: %v", err)
	}

	// teste de ping
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao fazer ping no banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida!")
	return db, nil
}


