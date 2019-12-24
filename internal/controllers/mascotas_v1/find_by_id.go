package mascotas_v1

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/lib/pq"
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"net/http"
)

func FindById(c echo.Context) error {
	ID := c.Param("id")
	db := database.GetConnection()
	defer db.Close()
	var mascota models.Mascota
	if result := db.Model(&mascota).Where("id  ?", ID).Find(&mascota); result.RecordNotFound() {
		return echo.ErrNotFound
	} else if result.Error != nil {
		pqErr := (result.Error).(*pq.Error)
		fmt.Printf("code :%s , : message: %s , type: %s", pqErr.Code, pqErr.Message, pqErr.Code.Class().Name())
	}
	var cliente models.Cliente
	db.Model(&mascota).Related(&cliente)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"mascota": &mascota,
		//"cliente": cliente,
	})
}
