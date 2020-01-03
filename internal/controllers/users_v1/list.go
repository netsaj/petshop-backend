package users_v1

import (
	"net/http"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
)

func List(c echo.Context) error {
	var users []models.Usuario
	db := database.GetConnection()
	defer db.Close()

	db.Find(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"usuarios": users,
	})
}
