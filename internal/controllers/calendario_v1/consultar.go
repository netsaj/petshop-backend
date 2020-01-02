package calendario_v1

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func ConsultarPendientes(c echo.Context) error {
	db := database.GetConnection()
	defer db.Close()
	var calendarios []models.Calendario
	if result := db.Model(&calendarios).
		Preload("Tercero").
		Preload("Mascota").
		Preload("Documento").
		Preload("UsuarioCierre").
		Find(&calendarios, "terminado = ? and fecha_agendada <= ?", false, time.Now()); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"calendarios": calendarios,
	})
}

func ConsultarHistorial(c echo.Context) error {
	db := database.GetConnection()
	defer db.Close()
	var calendarios []models.Calendario

	fechaInicio := c.QueryParam("fecha_inicio") // fecha de inicio para el filtrado
	fechaFinal := c.QueryParam("fecha_fin")     // fecha final para el filtrado

	query := db.Model(&calendarios).
		Joins("left join terceros on terceros.id = calendarios.tercero_id").
		Joins("left join mascotas on mascotas.id = calendarios.mascota_id").
		Joins("left join documentos on documentos.id = calendarios.documento_id").
		Preload("Tercero").
		Preload("Mascota").
		Preload("Documento").
		Preload("UsuarioCierre")

	criteria := c.QueryParam("q")
	where := "" // wuere adicional para filtrar si llega un criterio de busqueda.
	for _, q := range strings.Split(criteria, " ") {
		if len(where) > 0 {
			where += " AND "
		}
		where += "(calendarios.tipo ILIKE  '%" + q + "%' OR terceros.nombre ILIKE  '%" + q + "%' OR CAST(terceros.cedula as TEXT) ILIKE  '%" + q + "%' or terceros.telefono ILIKE  '%" + q + "%' or terceros.celular ILIKE  '%" + q + "%' or terceros.direccion ILIKE  '%" + q + "%' or terceros.barrio ILIKE  '%" + q + "%' or terceros.email ILIKE  '%" + q + "%' )"
	}
	if where != "" {
		query = query.Where(where)
	}

	if fechaInicio != "" { // agrego la fecha de inicio si esta definida, en la consulta
		fechaI := strings.Split(fechaInicio, "T")[0]
		query = query.Where(" DATE(calendarios.fecha_agendada) >= ? ", fechaI)
	}
	if fechaFinal != "" { // agrego la fecha de fin si esta definida, en la consulta.
		fechaF := strings.Split(fechaFinal, "T")[0]
		query = query.Where(" DATE(calendarios.fecha_agendada) <= ? ", fechaF)
	}

	if result := query.
		Find(&calendarios); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"calendarios": calendarios,
	})
}
