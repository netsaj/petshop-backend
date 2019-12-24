package clientes_v1

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
	"net/http"
)

func Update(c echo.Context) error {
	var cliente models.Cliente
	c.Bind(&cliente)
	if err := c.Validate(&cliente); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorToMap(err))
	}
	db := database.GetConnection()
	defer db.Close()
	spew.Print(&cliente)
	if result := db.Model(&cliente).Where("id = ?", cliente.ID).Update(&cliente).First(&cliente, "id = ?", cliente.ID); result.Error != nil {
		code, body := utils.ErrorHandler(result.Error)
		return c.JSON(code, body)
	}
	// resultado OK
	return c.JSON(http.StatusOK, map[string]interface{}{
		"cliente": &cliente,
	})

}
