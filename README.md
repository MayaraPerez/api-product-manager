# api-product-manager

A RESTful API developed with Go, using the Gin Gonic library, integrated with a PostgreSQL/MySQL database.<br>
The project follows a Clean Architecture
And uses Docker to manage the execution environment
both the application and the database are executed in Docker containers.


## Project Structure
```
│ └── main.go                 # Initialize the server and configure API routing
├── config/
│ └── config.go               # Configurations and management
├── routes/
│ └── routes.go               # Define the application's routes
├── controller/
│ └── product_controller.go   # Handle HTTP requests and basic validations
├── useCase/
│ └── product_useCase.go      # Process data and apply business rules
├── repository/
│ └── product_repository.go   # Query and manipulate data in the database data
├── db/
│ └── connection_db.go        # Configure the database connection
├── model/
│ ├── product.go              # Define the product structure
│ └── response.go             # Define response structures
├── Dockerfile                # Define the container environment
├── docker-compose.yml        # Configure the application and database in Docker
├── .env.example              # Example of environment variables
└── go.mod                    # Project dependencies

## Tecnologias Usadas 

- Linguagem de programação: Go
- Framework para rotas: Gin Gonic
- Banco de dados: PostgreSQL ou MySQL
- Gerenciamento de containers: Docker e Docker Compose
- Autenticação: JWT (EM ANDAMENTO)
```


## Pré-requisitos

Certifique-se de ter as seguintes ferramentas instaladas no seu ambiente de desenvolvimento:

- Docker e Docker Compose
- Git
- [Go](https://go.dev/doc/install) 1.20


## Installing dependencies:
Após a clonagem, entre no diretório do projeto e instale as dependências necessárias. Instale as dependências do Go: Se você não tiver o Go instalado Baixe as dependências do projeto: Dentro do diretório do projeto, execute:

Go 1.20
Rodando o projeto:

Clone este repositório:

Clone este repositório:
````
git clone https://github.com/MayaraPerez/api-product-manager.git
````
Navegue até o diretório do projeto:
````
cd api-ProdManager
````

Execute o programa:
````
go run cmd/main.go
````


