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

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

const STORAGE_PATH = "storage/"

func UploadFile(c echo.Context) error {
	// Source
	var archivo models.Archivo
	file, err := c.FormFile("file")
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
	db := database.GetConnection()
	defer db.Close()
	// Destination
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	archivo.Nombre = file.Filename
	archivo.Ruta =  STORAGE_PATH + filename
	archivo.ContentType = file.Header.Get("Content-Type")
	split := strings.Split(file.Filename, ".")
	archivo.Extension = split[len(split)-1]
	archivo.Tamaño = file.Size
	dst, err := os.Create(filepath.FromSlash(archivo.Ruta))
	if err != nil {
		fmt.Println(err)
		return utils.ReturnError(err, c)
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println(err)
		return utils.ReturnError(err, c)
	}
	db.Save(&archivo)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"archivo": archivo,
	})

}
