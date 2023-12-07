package services

import (
	prodCliente "github.com/aaraya0/ingsw3-final/backend/client"
	"github.com/aaraya0/ingsw3-final/backend/dto"
	"github.com/aaraya0/ingsw3-final/backend/model"
	e "github.com/aaraya0/ingsw3-final/backend/utils/errors"
)

type productService struct{}

type productServiceInterface interface {
	GetProducts() (dto.ProductsDto, e.ApiError)
	InsertProduct(productDto dto.ProductDto) (dto.ProductDto, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}

func (p *productService) GetProducts() (dto.ProductsDto, e.ApiError) {
	var products model.Products = prodCliente.GetProducts()
	var productsDto dto.ProductsDto
	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Title = product.Title
		productDto.Price = product.Price
		productDto.Id = product.Id
		productDto.Image = product.Image
		productDto.Description = product.Description
		productsDto = append(productsDto, productDto)
	}
	return productsDto, nil

}
func (s *productService) InsertProduct(productDto dto.ProductDto) (dto.ProductDto, e.ApiError) {

	var product model.Product

	product.Title = productDto.Title
	product.Price = productDto.Price
	product.Description = productDto.Description
	product.Image = productDto.Image

	product = prodCliente.InsertProduct(product)

	productDto.Id = product.Id

	var productResponseDto dto.ProductDto

	productResponseDto.Id = product.Id
	productResponseDto.Title = product.Title
	productResponseDto.Description = product.Description
	productResponseDto.Price = product.Price
	productResponseDto.Image = product.Image

	return productResponseDto, nil
}
