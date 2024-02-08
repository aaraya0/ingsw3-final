package db

import (
	"os"

	prodCliente "github.com/aaraya0/ingsw3-final/backend/client"
	"github.com/aaraya0/ingsw3-final/backend/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// ingsw3-final-database-1
	//db, err := sql.Open("mysql", "user:password@tcp("+os.Getenv("DB_HOST")+")/dbname")
	//dsn := "root:aaraya0@tcp(my-app-database:3307)/fastfood?charset=utf8mb4&parseTime=True&loc=Local"
	//DB_HOST := "my-app-database.run.app"
	dsn := "root:aaraya0@tcp(" + os.Getenv("DB_HOST") + ":3306)/fastfood?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:aaraya0@tcp(my-app-database-xho37fneiq-uc.a.run.app:3306)/fastfood?charset=utf8mb4&parseTime=True&loc=Local"
	//DB_HOST := "my-app-database-xho37fneiq-uc.a.run.app/"
	//dsn := "root:aaraya0@tcp(" + DB_HOST + ":3306)/fastfood?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:aaraya0@tcp(ec2-user-database-1:3306)/fastfood?charset=utf8mb4&parseTime=True&loc=Local"
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
