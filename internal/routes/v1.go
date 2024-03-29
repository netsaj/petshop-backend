package routes

import (
	"github.com/labstack/echo"
	"github/netsaj/petshop-backend/internal/controllers/historia_clinica_v1"

	"github/netsaj/petshop-backend/internal/controllers/archivos_v1"
	"github/netsaj/petshop-backend/internal/controllers/auth_v1"
	"github/netsaj/petshop-backend/internal/controllers/barrios_v1"
	"github/netsaj/petshop-backend/internal/controllers/calendario_v1"
	"github/netsaj/petshop-backend/internal/controllers/clientes_v1"
	"github/netsaj/petshop-backend/internal/controllers/desparasitantes_v1"
	"github/netsaj/petshop-backend/internal/controllers/documentos_v1"
	"github/netsaj/petshop-backend/internal/controllers/examenes_v1"
	"github/netsaj/petshop-backend/internal/controllers/mascotas_v1"
	"github/netsaj/petshop-backend/internal/controllers/reportes_v1"
	"github/netsaj/petshop-backend/internal/controllers/servicios_v1"
	"github/netsaj/petshop-backend/internal/controllers/users_v1"
	"github/netsaj/petshop-backend/internal/controllers/vacunas_v1"
	"github/netsaj/petshop-backend/internal/middleware"
)

func V1(e *echo.Echo) {
	// exploracion de archivos
	e.GET("/storage/:filename", archivos_v1.ShowFile)
	e.GET("/logo", archivos_v1.ShowLogo)

	// api endpoints
	v1 := e.Group("/v1")

	auth := v1.Group("/auth")
	auth.POST("/login", auth_v1.Login)
	auth.GET("/user", auth_v1.GetLoggedUser, middleware.CustomJwtMiddleware(), middleware.ValidateStandardUserMiddleware())

	// barrios
	barrios := v1.Group("/barrios")
	barrios.GET("", barrios_v1.Search)
	barrios.POST("", barrios_v1.Create)
	barrios.PUT("", barrios_v1.Create)
	barrios.DELETE("/:id", barrios_v1.Delete)

	// clientes
	clientes := v1.Group("/clientes")
	clientes.GET("", clientes_v1.Search)
	clientes.POST("", clientes_v1.Create)
	clientes.PUT("", clientes_v1.Update)
	clientes.GET("/:id", clientes_v1.FindByID)
	clientes.GET("/terceros", clientes_v1.SearchAll)

	// mascotas
	mascotas := v1.Group("/mascotas")
	mascotas.POST("", mascotas_v1.Create)
	mascotas.PUT("", mascotas_v1.Create)
	mascotas.GET("", mascotas_v1.Search)
	mascotas.GET("/:id", mascotas_v1.FindById)

	// vacunas
	vacunas := v1.Group("/vacunas")
	vacunas.POST("", vacunas_v1.CrearVacuna)
	vacunas.DELETE("/:id", vacunas_v1.Delete)
	vacunas.PUT("", vacunas_v1.Update)

	// examenes
	examenes := v1.Group("/examenes")
	examenes.GET("", examenes_v1.List)

	// vacunas - rutas para los grupos de vacunas
	gruposVacunas := vacunas.Group("/grupos")
	gruposVacunas.GET("", vacunas_v1.ListarGrupoVacunas)
	gruposVacunas.POST("", vacunas_v1.CrearGrupoVacuna)

	// desparasitantes
	desparasitantes := v1.Group("/desparasitantes")
	desparasitantes.POST("", desparasitantes_v1.Create)
	desparasitantes.DELETE("/:id", desparasitantes_v1.Delete)
	desparasitantes.PUT("", desparasitantes_v1.Update)
	gruposDesparasitantes := desparasitantes.Group("/grupos")
	gruposDesparasitantes.GET("", desparasitantes_v1.ListarGrupos)

	// documentos
	documentos := v1.Group("/documentos")
	documentos.DELETE("/:id", documentos_v1.Delete) // elimina un documento y toda la informacion asociada.

	// servicios
	servicios := documentos.Group("/servicios")
	servicios.POST("", servicios_v1.NuevoServicio)
	servicios.PUT("", servicios_v1.EditarUnServicio)
	servicios.GET("", servicios_v1.ListarServiciosActivos)
	servicios.GET("/all", servicios_v1.ListarServicios)
	servicios.GET("/:id", servicios_v1.FindByID)
	servicios.GET("/contar", servicios_v1.ContarServicios)

	// calendario
	calendario := v1.Group("/calendarios")
	calendario.GET("/pendientes", calendario_v1.ConsultarPendientes)
	calendario.GET("/historial", calendario_v1.ConsultarHistorial)
	calendario.POST("/cerrar", calendario_v1.CerrarCalendario)

	// usuarios
	usuarios := v1.Group("/usuarios")
	usuarios.GET("", users_v1.List)
	usuarios.POST("", users_v1.Create)
	usuarios.DELETE("/:id", users_v1.Delete)
	usuarios.PUT("", users_v1.Update)

	// carga y descarga de archivos
	archivos := v1.Group("/archivos")
	archivos.POST("/cargar", archivos_v1.UploadFile)
	archivos.POST("/laboratorio/cargar", archivos_v1.UploadFileLaboratorio)
	archivos.GET("/laboratorio/borrar/:archivo_id", archivos_v1.DeleteFile)

	reportes := v1.Group("/reportes")
	reportes.GET("/general", reportes_v1.ServiciosGlobal)
	reportes.GET("/mascota/:id", reportes_v1.ReporteMascota)

	// historia clinica
	historia := v1.Group("/historiaclinica")
	historia.POST("", historia_clinica_v1.Create)
}
