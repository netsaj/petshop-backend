package barrios_v1

import (
	"net/http"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func Search(c echo.Context) error {

	db := database.GetConnection()
	defer db.Close()
	var barrios []models.Barrio
	if result := db.Order("municipio asc").Order("nombre asc").Find(&barrios); result.Error != nil {
		code, body := utils.ErrorHandler(result.Error)
		return c.JSON(code, body)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"barrios": barrios,
	})

}
