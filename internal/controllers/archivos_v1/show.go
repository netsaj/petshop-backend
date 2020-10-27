package archivos_v1

import (
	"path/filepath"

	"github.com/labstack/echo"
)

func ShowFile(c echo.Context) error {
	filename := c.Param("filename")
	return c.File(filepath.FromSlash(STORAGE_PATH + filename))
}

func ShowLogo(c echo.Context) error {
	return c.File(filepath.FromSlash("./logo.png"))
}
