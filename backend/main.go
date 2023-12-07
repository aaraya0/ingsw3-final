package main

import (
	"github.com/aaraya0/ingsw3-final/backend/app"
	"github.com/aaraya0/ingsw3-final/backend/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}
