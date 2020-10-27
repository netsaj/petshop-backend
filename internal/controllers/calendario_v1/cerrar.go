package calendario_v1

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

type CerrarCalendarioPayload struct {
	ID                   uuid.UUID `json:"id"`
	UsuarioID            uuid.UUID `json:"usuario_id"`
	ObservacionesCerrado string    `json:"observaciones_cerrado"`
}

func CerrarCalendario(c echo.Context) error {
	var cierre CerrarCalendarioPayload
	if err := c.Bind(&cierre); err != nil {
		return utils.ReturnError(err, c)
	}
	db := database.GetConnection()
	defer db.Close()
	var calendario models.Calendario
	if result := db.Find(&calendario, "id = ?", cierre.ID); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	calendario.ObservacionesCerrado = cierre.ObservacionesCerrado
	calendario.UsuarioCierreID = cierre.UsuarioID
	calendario.Terminado = true
	calendario.FechaCierre = time.Now()
	if result := db.Save(&calendario); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"calendario": calendario,
	})

}
