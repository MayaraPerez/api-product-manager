# api-product-manager

A RESTful API developed with Go, using the Gin Gonic library, integrated with a PostgreSQL/MySQL database.<br>
The project follows a Clean Architecture
And uses Docker to manage the execution environment
both the application and the database are executed in Docker containers.


## Project Structure
```
├── cmd/
│ ├── .env                     # Environment variables file for local configuration
│ ├── init.go                  # Initial server configuration (dependencies, middlewares)
│ └── main.go                  # Initialize the server and configure API routing
├── config/
│ └── config.go                # Configuration management (e.g., loading .env, global variables)
├── controller/
│ └── product_controller.go    # Handle HTTP requests and perform basic validations
├── db/
│ └── connection_db.go         # Configure the connection to the database data
├── model/
│ ├── product.go               # Defines the product structure
│ └── response.go              # Defines response structures (e.g., error messages, success)
├── repository/
│ └── product_repository.go    # Performs database operations (CRUD)
├── routes/
│ └── routes.go                # Defines the application routes
├── useCase/
│ └── product_useCase.go       # Contains business logic (rules and data processing)
├── .gitignore                 # File to ignore files in Git (e.g., .env, binaries)
├── algo.txt                   # Notes or algorithms used in the project
├── docker-compose.yml         # Docker environment configuration (API and database)
├── Dockerfile                 # Defines the container environment for the application
├── go.mod                     # Go dependency management
├── go.sum                     # Checksums of Go dependencies
└── init.sql                   # Script to initialize the database

```

## Technologies Used

- Programming language: Go
- Routing framework: Gin Gonic
- Database: PostgreSQL or MySQL
- Container management: Docker and Docker Compose
- Authentication: JWT (IN PROGRESS)

##  Installing dependencies:
After cloning, enter the project directory and install the necessary dependencies. Install Go dependencies: If you don't have Go installed Download the project dependencies: Inside the project directory, run:

- Docker e Docker Compose
- Git
- [Go](https://go.dev/doc/install) 1.20


## Running the project:

Clone this repository:

````
git clone https://github.com/MayaraPerez/api-product-manager.git
````
Navigate to the project directory:
````
cd api-ProdManager
````

Run the program:
````
go run cmd/main.go
````


