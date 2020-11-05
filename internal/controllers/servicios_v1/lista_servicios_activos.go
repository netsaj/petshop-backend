package servicios_v1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func ListarServiciosActivos(c echo.Context) error {
	db := database.GetConnection()
	defer db.Close()
	var documentos []models.Documento
	query := db.Model(models.Documento{}).
		Joins("left join peluqueadas on peluqueadas.documento_id = documentos.id ").
		Joins("left join vacunaciones on vacunaciones.documento_id = documentos.id ").
		Joins("left join desparasitaciones on desparasitaciones.documento_id = documentos.id ").
		Joins("left join examenes_laboratorio on examenes_laboratorio.documento_id = documentos.id ").
		Joins("left join historia_clinicas on historia_clinicas.documento_id = documentos.id ").
		Preload("HistoriaClinica").
		Preload("ExamenLaboratorio").
		Preload("ExamenLaboratorio.ArchivosLaboratorio").
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
		Order("documentos.created_at desc").
		Group("documentos.id").
		Where("documentos.tipo = 'venta' and documentos.subtipo = 'servicio' and " +
			"(" +
			"(peluqueadas.id != '00000000-0000-0000-0000-000000000000' AND peluqueadas.terminado = false) OR " +
			"(vacunaciones.id != '00000000-0000-0000-0000-000000000000' and vacunaciones.terminado = false) OR " +
			"(desparasitaciones.id != '00000000-0000-0000-0000-000000000000' and desparasitaciones.terminado = false) OR " +
			"(historia_clinicas.id != '00000000-0000-0000-0000-000000000000' and historia_clinicas.terminado = false) OR " +
			"(examenes_laboratorio.id != '00000000-0000-0000-0000-000000000000' and examenes_laboratorio.terminado = false)" +
			") ")

	var count uint
	if result := query.Count(&count); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}

	if result := query.
		Find(&documentos);
		result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}

	fmt.Printf("total servicios : %v \n", count)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"hola":       "hola",
		"total":      count,
		"documentos": documentos,
	})
}
