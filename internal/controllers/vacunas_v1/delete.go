package vacunas_v1

import (
	"net/http"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func Delete(c echo.Context) error {
	ID := c.Param("id")
	db := database.GetConnection()
	defer db.Close()
	var vacuna models.Vacuna
	if result := db.Find(&vacuna, "id = ?", ID); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	if result := db.Delete(&vacuna); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"deleted": true,
		"vacuna":  vacuna,
	})
}
