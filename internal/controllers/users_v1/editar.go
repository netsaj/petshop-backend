package users_v1

import (
	"net/http"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func Update(c echo.Context) error {
	var usuario models.Usuario
	if err := c.Bind(&usuario); err != nil {
		return utils.ReturnError(err, c)
	}
	db := database.GetConnection()
	defer db.Close()
	if result := db.Save(&usuario); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"usuario": usuario,
	})
}
