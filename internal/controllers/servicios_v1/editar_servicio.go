package servicios_v1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

type EditarServicio struct {
	ID                uuid.UUID `json:"id"`
	Prefijo           uuid.UUID `json:"prefijo_id"`
	Numero            uint32    `json:"numero"`
	TerceroID         uuid.UUID `json:"tercero_id" validate:"required"`
	MascotaID         uuid.UUID `json:"mascota_id" validate:"required"`
	UsuarioID         uuid.UUID `json:"usuario_id" validate:"required"`
	ServicioTerminado bool      `json:"servicio_terminado"`
	CreatedAt         time.Time `json:"created_at"`

	Peluqueria        models.Peluqueria        `json:"peluqueria" validate:"omitempty"`
	Vacunacion        models.Vacunacion        `json:"vacunacion" validate:"omitempty"`
	Desparasitacion   models.Desparasitacion   `json:"desparasitacion" validate:"omitempty"`
	ExamenLaboratorio models.ExamenLaboratorio `json:"laboratorio" validate:"omitempty"`
}

func EditarUnServicio(c echo.Context) error {
	var servicio EditarServicio
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
	documento.ID = servicio.ID
	documento.PrefijoID = servicio.Prefijo
	documento.Numero = servicio.Numero
	documento.ServicioTerminado = servicio.ServicioTerminado
	documento.TerceroID = servicio.TerceroID
	documento.MascotaID = servicio.MascotaID
	documento.UsuarioID = servicio.UsuarioID
	documento.Peluqueria = servicio.Peluqueria
	documento.Vacunacion = servicio.Vacunacion
	documento.Desparasitacion = servicio.Desparasitacion
	documento.ExamenLaboratorio = servicio.ExamenLaboratorio
	documento.CreatedAt = servicio.CreatedAt

	if err := documento.CrearDocumentoServicio(); err != nil {
		return utils.ReturnError(err, c)
	}
	spew.Println(documento)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"documento": &documento,
	})
}
