package barrios_v1

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)


func Create(c echo.Context) error {
	var barrio models.Barrio
	if err := c.Bind(&barrio); err != nil {
		return utils.ReturnError(err, c)
	}
	db := database.GetConnection()
	defer db.Close()
	barrio.ID = uint(int32(time.Now().Unix()))
	if result := db.Save(&barrio); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"barrio": barrio,
	})
}
