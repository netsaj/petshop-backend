package vacunas_v1

import (
	"github.com/labstack/echo"
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
	"net/http"
)

func CrearVacuna(c echo.Context) error {
	var vacuna models.Vacuna
	if err := c.Bind(&vacuna); err != nil {
		return utils.ReturnError(err, c)
	}
	if err := c.Validate(&vacuna); err != nil {
		return utils.ReturnError(err, c)
	}
	db := database.GetConnection()
	defer db.Close()
	for j := 0; j < len(vacuna.GrupoVacuna); j++ {
		id := vacuna.GrupoVacuna[j].ID
		db.First(&vacuna.GrupoVacuna[j], "id = ?", id)
	}
	if result := db.Preload("GrupoVacuna").Save(&vacuna); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"vacuna": vacuna,
	})
}
