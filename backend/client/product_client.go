package client

import (
	"log"

	"github.com/aaraya0/ingsw3-final/backend/model"
	"gorm.io/gorm"
)

type Database interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
}

var Db Database

func GetProducts() model.Products {
	var products model.Products
	Db.Find(&products)
	log.Println("Products: ", products)
	return products
}

func InsertProduct(product model.Product) model.Product {
	result := Db.Create(&product)

	if result.Error != nil {
		log.Println(result.Error)
	}
	log.Println("Product submitted: ", product.Id)
	return product
}
