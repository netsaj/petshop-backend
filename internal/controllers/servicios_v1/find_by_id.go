package servicios_v1

import (
	"github.com/labstack/echo"
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
	"net/http"
)

func FindByID(c echo.Context) error {
	id := c.Param("id")
	db := database.GetConnection()
	defer db.Close()
	query := db.Model(models.Documento{}).
		Joins("left join peluqueadas on peluqueadas.documento_id = documentos.id ").
		Joins("left join vacunaciones on vacunaciones.documento_id = documentos.id ").
		Preload("Peluqueria").
		Preload("Vacunacion").
		Preload("Vacunacion.Vacuna").
		Preload("Vacunacion.GrupoVacuna").
		Preload("Usuario").
		Preload("Tercero").
		Preload("Mascota").
		Preload("Prefijo").
		Where("documentos.id = '" + id + "' and documentos.tipo = 'venta' and documentos.subtipo = 'servicio' and " +
			"((peluqueadas.id != '00000000-0000-0000-0000-000000000000' AND peluqueadas.terminado = false) OR (vacunaciones.id != '00000000-0000-0000-0000-000000000000' and vacunaciones.terminado = false)) ")

	var documento models.Documento
	if result := query.First(&documento); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"documento": documento,
	})
}
