package clientes_v1

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
	"net/http"
)

func Create(c echo.Context) error {
	var cliente models.Tercero
	c.Bind(&cliente)
	if err := c.Validate(&cliente); err != nil {
		return utils.ReturnError(err, c)
	}
	db := database.GetConnection()
	cliente.IsCliente = true
	defer db.Close()
	spew.Print(&cliente)
	if result := db.Create(&cliente); result.Error != nil {
		code, body := utils.ErrorHandler(result.Error)
		return c.JSON(code, body)
	}
	// resultado OK
	return c.JSON(http.StatusOK, map[string]interface{}{
		"cliente": &cliente,
	})

}
