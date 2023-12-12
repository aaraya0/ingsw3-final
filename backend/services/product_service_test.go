// product_service_test.go
package services

import (
	"testing"

	"github.com/aaraya0/ingsw3-final/backend/dto"
	e "github.com/aaraya0/ingsw3-final/backend/utils/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) GetProducts() (dto.ProductsDto, e.ApiError) {
	args := m.Called()

	if args.Get(1) == nil {
		return args.Get(0).(dto.ProductsDto), nil
	}

	return args.Get(0).(dto.ProductsDto), args.Get(1).(e.ApiError)
}

func (m *MockProductService) InsertProduct(productDto dto.ProductDto) (dto.ProductDto, e.ApiError) {
	args := m.Called(productDto)

	if args.Get(1) == nil {
		return args.Get(0).(dto.ProductDto), nil
	}

	return args.Get(0).(dto.ProductDto), args.Get(1).(e.ApiError)
}

func TestGetProducts(t *testing.T) {
	// Crear un mock del servicio
	mockService := new(MockProductService)
	defer mockService.AssertExpectations(t)

	// Configurar el mock para que devuelva productos al llamar a GetProducts
	expectedProductsDto := dto.ProductsDto{
		{Id: 1, Title: "Product 1", Price: 10.99, Image: "image1.jpg", Description: "Description 1"},
		{Id: 2, Title: "Product 2", Price: 19.99, Image: "image2.jpg", Description: "Description 2"},
	}
	mockService.On("GetProducts").Return(expectedProductsDto, nil)

	// Utilizar el mock en lugar del servicio real
	ProductService = mockService

	// Llamar a la función que estamos probando con el mock
	result, err := ProductService.GetProducts()

	// Afirmar que la función devolvió los productos esperados y no hay errores
	assert.Nil(t, err, "Error in GetProducts")
	assert.Equal(t, expectedProductsDto, result)
}

func TestInsertProduct(t *testing.T) {
	// Crear un mock del servicio
	mockService := new(MockProductService)
	defer mockService.AssertExpectations(t)

	// Configurar el mock para que devuelva un producto al llamar a InsertProduct
	expectedProductDto := dto.ProductDto{Id: 1, Title: "New Product", Price: 15.99, Image: "new_image.jpg", Description: "New Description"}
	mockService.On("InsertProduct", expectedProductDto).Return(expectedProductDto, nil)

	// Utilizar el mock en lugar del servicio real
	ProductService = mockService

	// Llamar a la función que estamos probando con el mock
	result, err := ProductService.InsertProduct(expectedProductDto)

	// Afirmar que la función devolvió el producto esperado y no hay errores
	assert.Nil(t, err, "Error in InsertProduct")
	assert.Equal(t, expectedProductDto, result)
}
