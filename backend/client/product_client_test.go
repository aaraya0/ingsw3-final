package client

import (
	"testing"

	"github.com/aaraya0/ingsw3-final/backend/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

var resultDetails string

func TestGetProducts(t *testing.T) {
	// Crear un mock de la base de datos
	mockDB := new(MockDB)
	defer mockDB.AssertExpectations(t)

	// Configurar el mock para que devuelva directamente los productos al llamar a Find
	expectedProducts := []model.Product{
		{Id: 1, Title: "Product 1", Price: 10.99, Image: "image1.jpg", Description: "Description 1"},
		{Id: 2, Title: "Product 2", Price: 19.99, Image: "image2.jpg", Description: "Description 2"},
	}
	mockDB.On("Find", mock.AnythingOfType("*model.Products"), mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).(*model.Products)
		*dest = expectedProducts
	}).Return(&gorm.DB{})

	// Utilizar el mock en lugar de la conexión real a la base de datos
	Db = mockDB

	// Llamar a la función que estamos probando con el mock
	result := GetProducts()

	// Convertir model.Products a []model.Product para la comparación
	resultSlice := model.Products(result)

	// Afirmar que la función devolvió los productos esperados
	assert.ElementsMatch(t, expectedProducts, resultSlice)
}

func TestInsertProduct(t *testing.T) {
	// Crear un mock de la base de datos
	mockDB := new(MockDB)
	defer mockDB.AssertExpectations(t)

	// Configurar el mock para que devuelva una instancia de Product al llamar a Create
	expectedProduct := model.Product{Id: 1, Title: "New Product", Price: 15.99, Image: "new_image.jpg", Description: "New Description"}
	mockDB.On("Create", &expectedProduct).Return(&gorm.DB{}, expectedProduct)

	// Utilizar el mock en lugar de la conexión real a la base de datos
	Db = mockDB

	// Llamar a la función que estamos probando con el mock
	result := InsertProduct(expectedProduct)

	// Afirmar que la función devolvió el producto esperado
	assert.Equal(t, expectedProduct, result)
}
