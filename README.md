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


