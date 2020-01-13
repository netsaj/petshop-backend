package clientes_v1

import (
	"net/http"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func FindByID(c echo.Context) error {
	ID := c.Param("id")
	var cliente models.Tercero
	db := database.GetConnection()
	defer db.Close()
	if result := db.Model(&cliente).
		Where("id = ?", &ID).
		Preload("Mascotas").
		Find(&cliente); result.Error != nil {
		println(result.Error)
		return c.JSON(http.StatusNotFound, utils.ErrorToMap(result.Error))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"cliente": &cliente,
	})
}
