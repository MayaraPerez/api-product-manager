package usecase

import (
	"api-product-manager/model"
	"api-product-manager/repository"
)


type ProductUserCase struct {
	Repository repository.ProductRepository
} 

func NewProductUserCase(repo repository.ProductRepository) ProductUserCase {
	return ProductUserCase{
		Repository: repo,
	}
} 

func (pu *ProductUserCase) GetProducts() ([]model.Product, error) {
	product, err := pu.Repository.GetProducts()
	if err != nil {
		return []model.Product{}, err
	}
	return product, nil
}

func (pu *ProductUserCase) CreateProduct(product model.Product) (model.Product, error){
	ProductId, err := pu.Repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}
	 product.ID = ProductId

	 return product, nil

}

func (pu *ProductUserCase) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.Repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
