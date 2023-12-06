package product

import (
	"backend/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetProducts() model.Products {
	var products model.Products
	Db.Find(&products)
	log.Debug("Products: ", products)
	return products

}

func InsertProduct(product model.Product) model.Product {
	result := Db.Create(&product)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Product submited: ", product.Id)
	return product
}
