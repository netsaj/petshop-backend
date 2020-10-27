package auth_v1

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"net/http"
)

func GetLoggedUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	db := database.GetConnection()
	defer db.Close()

	var usuario models.Usuario
	if result := db.First(&usuario,"id = ?", id); result.Error != nil {
		fmt.Print(result.Error.Error())
		return echo.ErrUnauthorized
	}
	fmt.Print(usuario)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": usuario,
	})
}
