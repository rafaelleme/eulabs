package service

import (
	"api/app/entity"
	"api/app/repository"
	"fmt"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo}
}

func (s *ProductService) CreateProduct(name string, price float64) error {
	product := &entity.Product{
		Name:  name,
		Price: price,
	}

	return s.repo.CreateProduct(*product)
}

func (s *ProductService) GetProducts(filter string, page int, limit int) ([]entity.Product, error) {
	return s.repo.GetProducts(filter, page, limit)
}

func (s *ProductService) GetProductByID(id int) (entity.Product, error) {
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		fmt.Println("Err on get product by ID:", id, err)
		return entity.Product{}, err
	}
	return product, nil
}

func (s *ProductService) UpdateProduct(productID int, updatedProduct entity.Product) error {
	product, err := s.GetProductByID(productID)
	if err != nil {
		return err
	}

	if updatedProduct.Name != product.Name {
		product.Name = updatedProduct.Name
	}

	if updatedProduct.Price != product.Price {
		product.Price = updatedProduct.Price
	}

	return s.repo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	product, err := s.GetProductByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteProduct(product)
}
