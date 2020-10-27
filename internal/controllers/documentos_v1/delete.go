package documentos_v1

import (
	"fmt"
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

	fmt.Printf("Eliminando documento ID: %v", ID)
	var documento models.Documento
	if result := db.Find(&documento, "id = ?", ID); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}

	db.Exec("DELETE FROM vacunaciones WHERE documento_id = ?", ID)
	db.Exec("DELETE FROM peluqueadas WHERE documento_id = ?", ID)
	db.Exec("DELETE FROM desparasitaciones WHERE documento_id = ?", ID)
	db.Exec("DELETE FROM calendarios WHERE documento_id = ?", ID)
	db.Exec("DELETE FROM examenes_laboratorio WHERE documento_id = ?", ID)

	if result := db.Delete(&documento, "id = ?", ID); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"documento": documento,
	})

}
