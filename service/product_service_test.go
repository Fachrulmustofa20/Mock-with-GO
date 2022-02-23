package service

import (
	"mock/entity"
	"mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetProductNotFound(t *testing.T) {
	// mock.on
	// parameter 1: menerima nama dari method
	// yang akan di mocking
	// parameter 2: Menerima argument yang
	// pada method
	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")

}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := entity.Product{
		Id:   "2",
		Name: "Kacamata",
	}

	productRepository.Mock.On("FindById", "2").Return(product)

	result, err := productService.GetOneProduct("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, product.Id, result.Id, "result has to be '2'")
	assert.Equal(t, product.Name, result.Name, "result has to be 'Kacamata'")
	assert.Equal(t, &product, result, "result has to be a product data with id '2'")

}
