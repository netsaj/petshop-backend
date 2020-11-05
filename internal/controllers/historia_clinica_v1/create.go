package historia_clinica_v1

import (
	"fmt"
	"github.com/labstack/echo"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func Create(c echo.Context) error {
	var historia models.HistoriaClinica
	if err := c.Bind(&historia); err != nil {
		fmt.Println("Error", err)
		code, body := utils.ErrorHandler(err)
		return c.JSON(code, body)
	}
	return c.JSON(200, map[string]interface{}{
		"historia": historia,
	})
}
