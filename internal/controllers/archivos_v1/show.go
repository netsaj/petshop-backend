package archivos_v1

import (
	"path/filepath"

	"github.com/labstack/echo"
)

func ShowFile(c echo.Context) error {
	filename := c.Param("filename")
	return c.File(filepath.FromSlash(STORAGE_PATH + filename))
}
