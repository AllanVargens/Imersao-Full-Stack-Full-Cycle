package service

import (
	"github.com/devfullcycle/imersao17/goapi/internal/database"
	"github.com/devfullcycle/imersao17/goapi/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

// Metodos

func (pd *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := pd.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (pd *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := pd.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pd *ProductService) GetProductsByCategoryID(categoryID string) ([]*entity.Product, error) {
	product, err := pd.ProductDB.GetProductByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pd *ProductService) CreateProduct(name, description, category_id, image_url string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, category_id, image_url, price)
	_, err := pd.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
