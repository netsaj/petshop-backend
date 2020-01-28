package reportes_v1

import (
	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func ReporteMascota(c echo.Context) error {
	mascotaID := c.Param("id")
	// fechaInicio := c.QueryParam("fecha_inicio") // fecha de inicio para el filtrado
	// fechaFinal := c.QueryParam("fecha_fin")     // fecha final para el filtrado
	db := database.GetConnection()
	defer db.Close()
	var documentos []models.Documento
	query := db.Model(models.Documento{}).
		Joins("left join peluqueadas on peluqueadas.documento_id = documentos.id ").
		Joins("left join vacunaciones on vacunaciones.documento_id = documentos.id ").
		Joins("left join desparasitaciones on desparasitaciones.documento_id = documentos.id ").
		Joins("left join terceros on terceros.id = documentos.tercero_id").
		Joins("left join mascotas on mascotas.id = documentos.mascota_id").
		Joins("left join prefijos on prefijos.id = documentos.prefijo_id").
		Joins("left join examenes_laboratorio on examenes_laboratorio.documento_id = documentos.id ").
		Preload("ExamenLaboratorio").
		Preload("Peluqueria").
		Preload("ExamenLaboratorio").
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
		Order("documentos.created_at desc").
		Group("documentos.id").
		// Where("documentos.created_at >= ?", fechaInicio).
		// Where("documentos.created_at <= ?", fechaFinal).
		Where("documentos.mascota_id = ?", mascotaID)
	if result := query.Find(&documentos); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}

	var countPeluqueria uint
	var countVacunacion uint
	var countDesparasitacion uint
	var countLaboratorio uint

	db.Model(models.Peluqueria{}).Joins("left join documentos on peluqueadas.documento_id = documentos.id ").Where("documentos.mascota_id = ?", mascotaID).
		Count(&countPeluqueria)
	db.Model(models.Vacunacion{}).Joins("left join documentos on vacunaciones.documento_id = documentos.id ").Where("documentos.mascota_id = ?", mascotaID).
		Count(&countVacunacion)
	db.Model(models.Desparasitacion{}).Joins("left join documentos on desparasitaciones.documento_id = documentos.id ").Where("documentos.mascota_id = ?", mascotaID).
		Count(&countDesparasitacion)
	db.Model(models.ExamenLaboratorio{}).Joins("left join documentos on examenes_laboratorio.documento_id = documentos.id ").Where("documentos.mascota_id = ?", mascotaID).
		Count(&countLaboratorio)

	return c.JSON(200, map[string]interface{}{
		"documentos":            documentos,
		"total_peluqueria":      countPeluqueria,
		"total_vacunacion":      countVacunacion,
		"total_desparasitacion": countDesparasitacion,
		"total_laboratorio":     countLaboratorio,
	})

}
