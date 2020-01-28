package clientes_v1

import (
	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
)

func SearchByRfidID(c echo.Context) error {
	// RfidID := c.Param("rfid")
	db := database.GetConnection()
	defer db.Close()
	return nil

}
