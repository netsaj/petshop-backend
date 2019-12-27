package routes

import (
	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/controllers/auth_v1"
	"github/netsaj/petshop-backend/internal/controllers/barrios_v1"
	"github/netsaj/petshop-backend/internal/controllers/calendario_v1"
	"github/netsaj/petshop-backend/internal/controllers/clientes_v1"
	"github/netsaj/petshop-backend/internal/controllers/mascotas_v1"
	"github/netsaj/petshop-backend/internal/controllers/servicios_v1"
	"github/netsaj/petshop-backend/internal/controllers/vacunas_v1"
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

	// vacunas
	vacunas := v1.Group("/vacunas")
	vacunas.POST("", vacunas_v1.CrearVacuna)
	// vacunas - rutas para los grupos de vacunas
	gruposVacunas := vacunas.Group("/grupos")
	gruposVacunas.GET("", vacunas_v1.ListarGrupoVacunas)
	gruposVacunas.POST("", vacunas_v1.CrearGrupoVacuna)

	// documentos
	documentos := v1.Group("/documentos")
	// servicios
	servicios := documentos.Group("/servicios")
	servicios.POST("", servicios_v1.NuevoServicio)
	servicios.PUT("", servicios_v1.EditarUnServicio)
	servicios.GET("", servicios_v1.ListarServiciosActivos)
	servicios.GET("/:id", servicios_v1.FindByID)
	servicios.GET("/contar", servicios_v1.ContarServicios)

	// calendario
	calendario := v1.Group("/calendarios")
	calendario.GET("/pendientes", calendario_v1.ConsultarPendientes)
}
