package servicios_v1

import (
	"net/http"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func FindByID(c echo.Context) error {
	id := c.Param("id")
	db := database.GetConnection()
	defer db.Close()
	query := db.Model(models.Documento{}).
		Joins("left join peluqueadas on peluqueadas.documento_id = documentos.id ").
		Joins("left join vacunaciones on vacunaciones.documento_id = documentos.id ").
		Joins("left join desparasitaciones on desparasitaciones.documento_id = documentos.id ").
		Joins("left join examenes_laboratorio on examenes_laboratorio.documento_id = documentos.id ").
		Preload("ExamenLaboratorio").
		Preload("ExamenLaboratorio.ArchivosLaboratorio").
		Preload("ExamenLaboratorio.ArchivosLaboratorio.Archivo").
		Preload("Peluqueria").
		Preload("Vacunacion").
		Preload("Vacunacion.Vacuna").
		Preload("Vacunacion.GrupoVacuna").
		Preload("Desparasitacion").
		Preload("Desparasitacion.Desparasitante").
		Preload("Desparasitacion.GrupoDesparasitante").
		Preload("Usuario").
		Preload("Tercero").
		Preload("Mascota").
		Preload("Prefijo").
		Where("documentos.tipo = 'venta' and documentos.subtipo = 'servicio' and " +
			"(" +
			"(peluqueadas.id != '00000000-0000-0000-0000-000000000000') OR " +
			"(vacunaciones.id != '00000000-0000-0000-0000-000000000000') OR " +
			"(desparasitaciones.id != '00000000-0000-0000-0000-000000000000') OR " +
			"(examenes_laboratorio.id != '00000000-0000-0000-0000-000000000000')" +
			") ")

	var documento models.Documento
	if result := query.First(&documento, "documentos.id =?", id); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"documento": documento,
	})
}
