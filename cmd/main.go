package main

import "api-product-manager/config"

func main() {
	config.Init()
	
	server, db, pool := initApp()
	defer db.Close()
	defer pool.Close()

	// Iniciando o servidor na porta 8000
	server.Run(":8000")
}
