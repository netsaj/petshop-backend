package routes

import (
	"github.com/labstack/echo"
	"github/netsaj/petshop-backend/internal/controllers/auth_v1"
	"github/netsaj/petshop-backend/internal/controllers/barrios_v1"
	"github/netsaj/petshop-backend/internal/controllers/clientes_v1"
	"github/netsaj/petshop-backend/internal/controllers/mascotas_v1"
	"github/netsaj/petshop-backend/internal/middleware"
)

func V1(e *echo.Echo) {
	v1 := e.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/login", auth_v1.Login)
	auth.GET("/user", auth_v1.GetLoggedUser, middleware.CustomJwtMiddleware(), middleware.ValidateAdminMiddleware())

	// barrios
	barrios := v1.Group("/barrios")
	barrios.GET("", barrios_v1.Search)
	// clientes
	clientes := v1.Group("/clientes")
	clientes.GET("", clientes_v1.Search)
	clientes.POST("", clientes_v1.Create)
	clientes.PUT("", clientes_v1.Update)
	clientes.GET("/:id", clientes_v1.FindByID)

	// mascotas
	mascotas := v1.Group("/mascotas")
	mascotas.POST("", mascotas_v1.Create)
	mascotas.PUT("", mascotas_v1.Create)
	mascotas.GET("", mascotas_v1.Search)
	mascotas.GET("/:id", mascotas_v1.FindById)
}
