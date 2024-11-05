package usecase

import (
	"go-api/internal/model"
	"go-api/internal/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return ProductUseCase{repository: repository}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUseCase) GetProductById(id int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUseCase) DelProductById(id int) error {
	err := pu.repository.DelProductById(id)
	return err
}
