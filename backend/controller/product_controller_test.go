package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aaraya0/ingsw3-final/backend/dto"
	"github.com/aaraya0/ingsw3-final/backend/services"
	"github.com/aaraya0/ingsw3-final/backend/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	// Mock del servicio para simular la obtención de productos
	services.ProductService = &MockProductService{}

	// Crear una instancia de Gin para las pruebas
	router := gin.Default()
	router.GET("/products", GetProducts)

	// Realizar una solicitud HTTP GET simulada a /products
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	router.ServeHTTP(w, req)

	// Verificar el código de estado HTTP y el cuerpo de la respuesta
	assert.Equal(t, http.StatusOK, w.Code)

	// Deserializar el cuerpo de la respuesta JSON en un objeto ProductsDto
	var responseDto dto.ProductsDto
	err := json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.Nil(t, err)

	// Verificar que la respuesta contiene datos de productos simulados
	assert.NotEmpty(t, responseDto)
}

// MockProductService es una implementación de ProductService para usar en las pruebas
type MockProductService struct{}

// GetProducts simula la obtención de productos
func (m *MockProductService) GetProducts() (dto.ProductsDto, errors.ApiError) {
	// Simular una lista de productos
	products := []dto.ProductDto{
		{Title: "Product 1", Price: 10.99, Image: "image1.jpg", Description: "Description 1"},
		{Title: "Product 2", Price: 19.99, Image: "image2.jpg", Description: "Description 2"},
	}

	return products, nil
}

// InsertProduct simula la inserción de un producto
func (m *MockProductService) InsertProduct(productDto dto.ProductDto) (dto.ProductDto, errors.ApiError) {
	// Simular la inserción exitosa
	return productDto, nil
}

// TestProductInsert simula la inserción de un producto a través de la API HTTP
func TestProductInsert(t *testing.T) {
	// Mock del servicio para simular la inserción de productos
	services.ProductService = &MockProductService{}

	// Crear una instancia de Gin para las pruebas
	router := gin.Default()
	router.POST("/product", ProductInsert)

	// Crear un objeto ProductDto simulado para la solicitud JSON
	productDto := dto.ProductDto{
		Title:       "Test Product",
		Price:       29.99,
		Image:       "test_image.jpg",
		Description: "Test Description",
	}

	// Convertir el objeto ProductDto en JSON
	payload, err := json.Marshal(productDto)
	assert.Nil(t, err)

	// Realizar una solicitud HTTP POST simulada a /product
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Verificar el código de estado HTTP y el cuerpo de la respuesta
	assert.Equal(t, http.StatusCreated, w.Code)

	// Deserializar el cuerpo de la respuesta JSON en un objeto ProductDto
	var responseDto dto.ProductDto
	err = json.Unmarshal(w.Body.Bytes(), &responseDto)
	assert.Nil(t, err)

	// Verificar que la respuesta contiene datos del producto simulado
	assert.Equal(t, productDto.Title, responseDto.Title)
	assert.Equal(t, productDto.Price, responseDto.Price)
	assert.Equal(t, productDto.Image, responseDto.Image)
	assert.Equal(t, productDto.Description, responseDto.Description)
}
