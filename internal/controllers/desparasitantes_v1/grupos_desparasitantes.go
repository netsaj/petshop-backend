package desparasitantes_v1

import (
	"net/http"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func ListarGrupos(c echo.Context) error {
	db := database.GetConnection()
	defer db.Close()
	var gruposDesparasitantes []models.GrupoDesparasitante

	if result := db.Model(&gruposDesparasitantes).
		Preload("Desparasitantes").
		Order("nombre asc").
		Find(&gruposDesparasitantes); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"grupos": gruposDesparasitantes,
	})
}
