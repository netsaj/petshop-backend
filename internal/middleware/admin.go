package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"time"
)

/*
 */
func ValidateAdminMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			fmt.Print(claims)
			isAdmin := claims["admin"].(bool)
			if isAdmin && int64(claims["exp"].(float64)) > time.Now().Unix() {
				println("is user admins")
				return next(c)
			}
			return echo.ErrUnauthorized
		}
	}
}
func ValidateStandardUserMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			if int64(claims["exp"].(float64)) > time.Now().Unix() {
				return next(c)
			}
			return echo.ErrUnauthorized
		}
	}
}
