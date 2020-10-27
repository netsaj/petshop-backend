package archivos_v1

import (
	"fmt"
	"github.com/labstack/echo"
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
	"net/http"
	"os"
	"path/filepath"
)

func DeleteFile(c echo.Context) error {

	ArchivoID := c.Param("archivo_id")
	db := database.GetConnection()

	var Archivo models.Archivo

	if result := db.Find(&Archivo, "id = ?", ArchivoID); result.Error != nil {
		fmt.Print("error encontrando archivo")
		return echo.ErrNotFound
	}

	var ArchivoLab models.ArchivosLaboratorio
	// borro los archivos del examen de laboratorio
	db.Delete(&ArchivoLab, "archivo_id = ?", ArchivoID)
	// borro los archivos de la base de datos
	if result := db.Delete(&Archivo, "id = ?", ArchivoID); result.Error != nil {
		fmt.Print("error eliminando archivo de la base de datos")
		return utils.ReturnError(result.Error, c)
	}
	// borro los archivos del disco
	if err := os.Remove(filepath.FromSlash(Archivo.Ruta)); err != nil{
		fmt.Print("error eliminando archivo del dicho")
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"result":  "deleted",
	})

}
