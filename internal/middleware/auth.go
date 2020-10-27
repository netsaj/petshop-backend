package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func CustomJwtMiddleware() echo.MiddlewareFunc {
	return middleware.JWT([]byte("secret"))
}
