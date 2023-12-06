package db

import (
	prodCliente "backend/client"

	"backend/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {

	dsn := "root:@tcp(127.0.0.1:3306)/fastfood?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Info("Connection fail")
		log.Fatal(err)
	} else {
		log.Info("Connection success")
	}
	prodCliente.Db = db
}
func StartDbEngine() {
	db.AutoMigrate(&model.Product{})
	log.Info("Finishing migration database tables")
}
