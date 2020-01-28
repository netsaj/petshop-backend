package examenes_v1

import (
	"net/http"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func List(c echo.Context) error {
	var examenes []models.Examenes
	db := database.GetConnection()
	defer db.Close()
	if result := db.Order("nombre asc").Find(&examenes); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"examenes": examenes,
	})
}
