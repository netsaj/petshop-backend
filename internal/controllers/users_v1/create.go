package users_v1

import (
	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func Create(c echo.Context) error {
	var user models.Usuario
	if err := c.Bind(&user); err != nil {
		return utils.ReturnError(err, c)
	}
	db := database.GetConnection()
	defer db.Close()
	user.SetPassword(user.Password)
	if result := db.Save(&user); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	return c.JSON(201, map[string]interface{}{
		"user": &user,
	})

}
