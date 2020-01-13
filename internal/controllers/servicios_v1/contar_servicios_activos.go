package servicios_v1

import (
	"fmt"
	"github.com/labstack/echo"
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"net/http"
)

func ContarServicios(c echo.Context) error {
	db := database.GetConnection()
	defer db.Close()
	var peluqueria uint
	var vacunacion uint
	var desparasitacion uint
	var laboratorio uint

	if result := db.Model(models.Peluqueria{}).Where("terminado = false").Count(&peluqueria); result.Error != nil {
		peluqueria = 0
	}
	if result := db.Model(models.Vacunacion{}).Where("terminado = false").Count(&vacunacion); result.Error != nil {
		fmt.Print("error al contar vacunacion")
		vacunacion = 0
	}
	if result := db.Model(models.Desparasitacion{}).Where("terminado = false").Count(&desparasitacion); result.Error != nil {
		fmt.Print("error al contar Desparasitacion")
		desparasitacion = 0
	}
	if result := db.Model(models.ExamenLaboratorio{}).Where("terminado = false").Count(&laboratorio); result.Error != nil {
		fmt.Print("error al contar ExamenLaboratorio")
		laboratorio = 0
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"peluqueria": peluqueria,
		"vacunacion": vacunacion,
		"desparasitacion":desparasitacion,
		"laboratorio":laboratorio,
	})
}
