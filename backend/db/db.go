package db

import (
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
	// Modifica el DSN para usar el usuario aaraya0, contraseña root y la base de datos fastfood
	dsn := "root:aaraya0@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"

	// Selecciona la base de datos recién creada o existente después de establecer la conexión
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Info("Connection fail")
		log.Fatal(err)
	}

	// Configura para que GORM cree la base de datos si no existe
	err = createDatabaseIfNotExists(db, "fastfood")
	if err != nil {
		log.Info("Database creation fail")
		log.Fatal(err)
	} else {
		log.Info("Database creation success")
	}

	prodCliente.Db = db
}

func StartDbEngine() {
	// AutoMigrate solo crea tablas si no existen
	db.AutoMigrate(&model.Product{})
	log.Info("Finishing migration database tables")
}

// createDatabaseIfNotExists crea la base de datos si no existe
func createDatabaseIfNotExists(db *gorm.DB, dbName string) error {
	// Intenta ejecutar una consulta en la base de datos para verificar si existe
	err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName).Error
	if err != nil {
		return err
	}

	// Selecciona la base de datos recién creada o existente
	return db.Exec("USE " + dbName).Error
}
