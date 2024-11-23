package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq" 

)

const (
	db_name = "product-manager"
	port = 5432
	user = "postgres"
	password = "@mayaraPerez1992"
	host = "go_db"
)

func ConnectDB() (*sql.DB, error) {
	// Configuração da string de conexão
	psqInfo:= fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
						   host, port, user, password, db_name)

	// Abrindo conexão com o postgres
	db, err := sql.Open("postgres", psqInfo)
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados:", err)
	}

	// teste de ping
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Conexão com o banco de dados estabelecida! Database: " + db_name)

	return db, nil
}


