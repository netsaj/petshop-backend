package calendario_v1

import (
	"net/http"
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
