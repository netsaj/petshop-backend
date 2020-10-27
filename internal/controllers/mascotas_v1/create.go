package mascotas_v1

import (
	"github.com/labstack/echo"
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
	"net/http"
)

func Create(c echo.Context) error {
	var mascota models.Mascota
	c.Bind(&mascota)
	if err := c.Validate(&mascota); err != nil {
		return utils.ReturnError(err, c)
	}
	db := database.GetConnection()
	defer db.Close()

	if result := db.Save(&mascota).Preload("Tercero").Find(&mascota); result.Error != nil {
		code, body := utils.ErrorHandler(result.Error)
		return c.JSON(code, body)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"mascota": &mascota,
	})

}
