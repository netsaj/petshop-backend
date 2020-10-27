package archivos_v1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func UploadFileLaboratorio(c echo.Context) error {
	// Source
	var Archivo models.Archivo
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return utils.ReturnError(err, c)
	}
	ExamenLaboratorioID, err := uuid.FromString(c.FormValue("examen_id"))
	if err != nil {
		fmt.Println(err)
		return utils.ReturnError(err, c)
	}
	src, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return utils.ReturnError(err, c)
	}
	defer src.Close()
	// Destination
	db := database.GetConnection()
	defer db.Close()
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	Archivo.Nombre = file.Filename
	Archivo.Ruta = STORAGE_PATH + filename
	Archivo.ContentType = file.Header.Get("Content-Type")
	split := strings.Split(file.Filename, ".")
	Archivo.Extension = split[len(split)-1]
	Archivo.Tamaño = file.Size
	dst, err := os.Create(filepath.FromSlash(Archivo.Ruta))
	if err != nil {
		fmt.Println(err)
		return utils.ReturnError(err, c)
	}

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println(err)
		return utils.ReturnError(err, c)
	}
	if result := db.Save(&Archivo); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	var ArchivoLab models.ArchivosLaboratorio
	ArchivoLab.ArchivoID = Archivo.ID
	ArchivoLab.ExamenLaboratorioID = ExamenLaboratorioID
	if result := db.Save(&ArchivoLab); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}
	db.Preload("Archivo").Find(&ArchivoLab)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"archivo": ArchivoLab,
	})

}
