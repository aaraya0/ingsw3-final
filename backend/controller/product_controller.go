package controllers

import (
	"net/http"

	"backend/dto"
	"backend/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetProducts(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetProducts()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)

}

func ProductInsert(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "POST")
	var productDto dto.ProductDto
	err := c.BindJSON(&productDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	productDto, er := service.ProductService.InsertProduct(productDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, productDto)
}
