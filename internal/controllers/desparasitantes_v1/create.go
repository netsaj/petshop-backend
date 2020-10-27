package desparasitantes_v1

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func Create(c echo.Context) error {
	var desparasitante models.Desparasitante
	if err := c.Bind(&desparasitante); err != nil {
		return utils.ReturnError(err, c)
	}
	if err := c.Validate(&desparasitante); err != nil {
		return utils.ReturnError(err, c)
	}
	db := database.GetConnection()
	defer db.Close()
	for j := 0; j < len(desparasitante.GruposDesparasitante); j++ {
		id := desparasitante.GruposDesparasitante[j].ID
		db.First(&desparasitante.GruposDesparasitante[j], "id = ?", id)
	}
	desparasitante.ID = uint(int32(time.Now().Unix()))
	if result := db.Preload("GruposDesparasitante").Save(&desparasitante); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"desparasitante": desparasitante,
	})
}
