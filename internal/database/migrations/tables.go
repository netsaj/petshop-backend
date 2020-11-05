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
		&models.Tercero{},
		&models.Mascota{},

		// listado de vacunas
		&models.GrupoVacuna{},
		&models.Vacuna{},
		// listado de desparasitante
		&models.GrupoDesparasitante{},
		&models.Desparasitante{},
		// examenes medicos
		&models.Examenes{},

		// manejo de documentos generados (servicios,...)
		&models.Prefijo{},
		&models.Documento{},
		&models.Peluqueria{},
		&models.Vacunacion{},
		&models.Desparasitacion{},
		&models.ExamenLaboratorio{},
		&models.ArchivosLaboratorio{},
		&models.Archivo{},

		// Historia clinica
		&models.HistoriaClinica{},


		// CRM (calendario, ...)
		&models.Calendario{},

	)
	createIndex()
}

func createIndex() {
	db := database.GetConnection()
	defer db.Close()
	// usuario
	db.Model(models.Usuario{}).AddUniqueIndex("username", "username")
	// clientes
	db.Model(models.Tercero{}).AddUniqueIndex("cedula", "cedula")

}
