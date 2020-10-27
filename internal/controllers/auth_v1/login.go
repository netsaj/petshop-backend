package auth_v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"net/http"
	"time"
)

func Login(c echo.Context) (err error) {

	type Params struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
	}

	u := new(Params)

	if err = c.Bind(u); err != nil {
		print(err)
		return echo.ErrBadRequest
	}
	DB := database.GetConnection()
	defer DB.Close()

	var user models.Usuario
	if err := DB.Where("username = ?", u.Username).First(&user).Error; err != nil {
		print(err)
	}
	if user.CheckPassword(u.Password) {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		// Set custom claims
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = user.ID
		claims["username"] = user.Username
		claims["admin"] = user.IsAdmin()
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		// Create token with claims
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}
