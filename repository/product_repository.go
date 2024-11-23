package repository

import (
	"api-product-manager/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}	

func (pr *ProductRepository) GetProducts() ([]model.Product, error){
	query := "SELECT id, product_name, price FROM product "
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productsList []model.Product
	var productObject model.Product

	for rows.Next() {
		err = rows.Scan(&productObject.ID, &productObject.Name, &productObject.Price)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}
		productsList = append(productsList, productObject)
	}

	defer rows.Close()

	return productsList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	//criando a query de create no banco
	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// executando a query passando os paraetros que queremos criar e pegamos o id que sera gerado
	var id int
	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	query.Close()

	return id, nil
	
}

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {
	// query para consultar por um id 
	query, err  := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// rodar a query 
	var product model.Product
	err = query.QueryRow(id_product).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		// verifica se o sql nao encontrou linha compativel com o id
		if err == sql.ErrNoRows {   
			return nil, nil  
		}
		fmt.Println(err)
		return nil, err
	}

	defer query.Close()

	return &product, nil

}

