package servicios_v1

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
	"net/http"
)

type Servicio struct {
	Prefijo   uuid.UUID `json:"prefijo"`
	TerceroID uuid.UUID `json:"tercero_id" validate:"required"`
	MascotaID uuid.UUID `json:"mascota_id" validate:"required"`
	UsuarioID uuid.UUID `json:"usuario_id" validate:"required"`

	Peluqueria models.Peluqueria `json:"peluqueria" validate:"omitempty"`
	Vacunacion models.Vacunacion `json:"vacunacion" validate:"omitempty"`
}

func NuevoServicio(c echo.Context) error {
	var servicio Servicio
	if err := c.Bind(&servicio); err != nil {
		fmt.Println(err)
		return utils.ReturnError(err, c)
	}
	if err := c.Validate(&servicio); err != nil {
		return utils.ReturnError(err, c)
	}
	db := database.GetConnection()
	defer db.Close()
	// si el prefijo es vació, lo asignos al default

	var documento models.Documento
	documento.PrefijoID = servicio.Prefijo
	documento.TerceroID = servicio.TerceroID
	documento.MascotaID = servicio.MascotaID
	documento.UsuarioID = servicio.UsuarioID
	documento.Peluqueria = servicio.Peluqueria
	documento.Vacunacion = servicio.Vacunacion

	if err := documento.CrearDocumentoServicio(); err != nil {
		return utils.ReturnError(err, c)
	}
	spew.Println(documento)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"documento": &documento,
	})
}
