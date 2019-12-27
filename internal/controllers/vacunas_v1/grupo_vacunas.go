package vacunas_v1

import (
	"github.com/labstack/echo"
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
	"net/http"
)

func ListarGrupoVacunas(c echo.Context) error {
	db := database.GetConnection()
	defer db.Close()

	var grupos []models.GrupoVacuna
	if result := db.Preload("Vacunas").Find(&grupos); result.Error != nil {
		code, body := utils.ErrorHandler(result.Error)
		return c.JSON(code, body)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"grupos": grupos,
	})
}

func CrearGrupoVacuna(c echo.Context) error {
	db := database.GetConnection()
	defer db.Close()
	var grupo models.GrupoVacuna
	if err := c.Bind(&grupo); err != nil {
		return utils.ReturnError(err, c)
	}
	if err := c.Validate(&grupo); err != nil {
		return utils.ReturnError(err, c)
	}
	if err := db.Save(&grupo); err.Error != nil {
		return utils.ReturnError(err.Error, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"grupo": grupo,
	})
}
