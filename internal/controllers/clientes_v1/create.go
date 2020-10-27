package clientes_v1

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func Create(c echo.Context) error {
	var cliente models.Tercero
	c.Bind(&cliente)
	if err := c.Validate(&cliente); err != nil {
		return utils.ReturnError(err, c)
	}
	db := database.GetConnection()
	var temp models.Tercero
	if !db.Where("documento = ?", cliente.Cedula).Find(&temp).RecordNotFound() {
		cliente.ID = temp.ID
	}

	cliente.IsCliente = true
	defer db.Close()
	spew.Print(&cliente)
	if result := db.Save(&cliente); result.Error != nil {
		code, body := utils.ErrorHandler(result.Error)
		return c.JSON(code, body)
	}
	// resultado OK
	return c.JSON(http.StatusOK, map[string]interface{}{
		"cliente": &cliente,
	})

}
