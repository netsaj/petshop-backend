package migrations

import (
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
)

func createTables() {
	dbClient := database.GetConnection()
	defer dbClient.Close()
	dbClient.AutoMigrate(
		&models.Barrio{},
		&models.Usuario{},
		&models.Cliente{},
		&models.Mascota{},
	)
	createIndex()
}

func createIndex() {
	db := database.GetConnection()
	defer db.Close()
	// usuario
	db.Model(models.Usuario{}).AddUniqueIndex("username", "username")
	// clientes
	db.Model(models.Cliente{}).AddUniqueIndex("cedula", "cedula")

}
