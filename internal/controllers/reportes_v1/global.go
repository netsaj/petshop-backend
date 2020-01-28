package reportes_v1

import (
	"strings"
	"time"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func ServiciosGlobal(c echo.Context) error {
	fechaInicio := c.QueryParam("fecha_inicio") // fecha de inicio para el filtrado
	fechaFinal := c.QueryParam("fecha_fin")     // fecha final para el filtrado
	if fechaInicio != "" {                      // agrego la fecha de inicio si esta definida, en la consulta
		fechaInicio = strings.Split(fechaInicio, "T")[0]
	} else {
		fechaInicio = "2019-01-01"
	}
	if fechaFinal != "" { // agrego la fecha de fin si esta definida, en la consulta.
		fechaFinal = strings.Split(fechaFinal, "T")[0]
	} else {
		fechaFinal = strings.Split(time.Now().String(), "T")[0]
	}
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
		Where(" DATE(documentos.created_at) >= ? and DATE(documentos.created_at) <= ?", fechaInicio, fechaFinal)
	if result := query.Find(&documentos); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}

	var countPeluqueria uint
	var countVacunacion uint
	var countDesparasitacion uint
	var countLaboratorio uint

	db.Model(models.Peluqueria{}).Joins("left join documentos on peluqueadas.documento_id = documentos.id ").Where(" DATE(documentos.created_at) >= ? and DATE(documentos.created_at) <= ?", fechaInicio, fechaFinal).
		Count(&countPeluqueria)
	db.Model(models.Vacunacion{}).Joins("left join documentos on vacunaciones.documento_id = documentos.id ").Where(" DATE(documentos.created_at) >= ? and DATE(documentos.created_at) <= ?", fechaInicio, fechaFinal).
		Count(&countVacunacion)
	db.Model(models.Desparasitacion{}).Joins("left join documentos on desparasitaciones.documento_id = documentos.id ").Where(" DATE(documentos.created_at) >= ? and DATE(documentos.created_at) <= ?", fechaInicio, fechaFinal).
		Count(&countDesparasitacion)
	db.Model(models.ExamenLaboratorio{}).Joins("left join documentos on examenes_laboratorio.documento_id = documentos.id ").Where(" DATE(documentos.created_at) >= ? and DATE(documentos.created_at) <= ?", fechaInicio, fechaFinal).
		Count(&countLaboratorio)

	return c.JSON(200, map[string]interface{}{
		"documentos":            documentos,
		"total_peluqueria":      countPeluqueria,
		"total_vacunacion":      countVacunacion,
		"total_desparasitacion": countDesparasitacion,
		"total_laboratorio":     countLaboratorio,
	})
}
