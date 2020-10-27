package desparasitantes_v1

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func Update(c echo.Context) error {
	var desparasitante models.Desparasitante
	if err := c.Bind(&desparasitante); err != nil {
		return utils.ReturnError(err, c)
	}
	db := database.GetConnection()
	defer db.Close()
	for j := 0; j < len(desparasitante.GruposDesparasitante); j++ {
		id := desparasitante.GruposDesparasitante[j].ID
		db.First(&desparasitante.GruposDesparasitante[j], "id = ?", id)
	}
	db.Exec(fmt.Sprintf("delete from grupo_desparasitante_desparasitantes where desparasitante_id = %d",desparasitante.ID))
	if result := db.Save(&desparasitante); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"desparasitante": desparasitante,
	})
}

func Copy(c echo.Context) error {
	var desparasitante models.Desparasitante
	if err := c.Bind(&desparasitante); err != nil {
		return utils.ReturnError(err, c)
	}
	db := database.GetConnection()
	defer db.Close()
	for j := 0; j < len(desparasitante.GruposDesparasitante); j++ {
		id := desparasitante.GruposDesparasitante[j].ID
		db.First(&desparasitante.GruposDesparasitante[j], "id = ?", id)
	}
	if result := db.Save(&desparasitante); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"desparasitante": desparasitante,
	})
}
