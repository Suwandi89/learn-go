package service

import (
	"challenge8/entity"
	"challenge8/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {

	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")

}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := entity.Product{
		Id:          "1",
		Title:       "Bantal keren",
		Description: "Bantal keren buatan rumah",
	}

	productRepository.Mock.On("FindById", "1").Return(product)

	result, err := productService.GetOneProduct("1")

	assert.Nil(t, err)

	assert.NotNil(t, result)

	assert.Equal(t, product.Id, result.Id, "result has to be '1'")
	assert.Equal(t, product.Title, result.Title, "result has to be 'Bantal keren'")
	assert.Equal(t, product.Description, result.Description, "result has to be 'Bantal keren buatan rumah'")
}
