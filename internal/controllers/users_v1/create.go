package users_v1

import (
	"fmt"
	"github.com/labstack/echo"
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
)

func Create(c echo.Context) error {
	var user models.Usuario
	if err := c.Bind(user); err != nil {
		return echo.ErrBadRequest
	}
	db := database.GetConnection()
	defer db.Close()
	if result := db.Model(user).Save(&user); result.Error != nil {
		for _, e := range result.GetErrors() {
			fmt.Printf("error: %s", e)
		}
		return echo.ErrBadRequest
	}
	return c.JSON(201, map[string]interface{}{
		"user": &user,
	})

}
