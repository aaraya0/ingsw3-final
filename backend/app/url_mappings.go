package app

import (
	router "backend/app"
	productController "backend/controller"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping

	router.GET("/products", productController.GetProducts)
	router.POST("/product", productController.ProductInsert)

	log.Info("Finishing mappings configurations")

}
