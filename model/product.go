package model

type Product struct {
	ID    int     `json:"id_product"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}


func UpdateVerifyProduct(productUpdade Product, product Product) Product {
	if productUpdade.Name != "" {
		product.Name = productUpdade.Name
	}
	if productUpdade.Price != 0 {
		product.Price = productUpdade.Price
	}
	return product
}