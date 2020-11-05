package main

import (
	"fmt"
	ES "github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github/netsaj/petshop-backend/internal/database/migrations"
	"github/netsaj/petshop-backend/internal/routes"
	"github/netsaj/petshop-backend/internal/utils"
	es_translation "github/netsaj/petshop-backend/internal/utils/translations/es"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func autoBackup() {
	os.Setenv("PGPASSWORD", "linux")
	currentTime := time.Now().Format(time.RFC3339)
	currentTime = strings.Replace(currentTime, ":", "_", -1)
	backupFile := filepath.FromSlash(fmt.Sprintf("backups/copia-seguridad-%s.backup",currentTime))
	cmd := exec.Command("pg_dump", "--file ",
		backupFile,
		" --host \"localhost\" --port \"5432\" --username \"postgres\" --no-password --verbose --format=c --blobs --section=pre-data --section=data --section=post-data \"petshop\"\n")
	fmt.Println(cmd.String())
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "PGPASSWORD=linux")
	err := cmd.Run()
	if err != nil {
		bytes, _ := cmd.StdoutPipe()
		fmt.Println("BACKUP FAILED --- ", err, "... ", bytes)
	}
}

func main() {
	autoBackup()
	migrations.Run()
	e := echo.New()
	// validator request
	esCo := ES.New()
	uni = ut.New(esCo, esCo)
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	es_translation.RegisterDefaultTranslations(validate, trans)
	e.Validator = utils.New()
	// === routes ====
	routes.V1(e)
	// start
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} \n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Server is running...")
	})
	e.Logger.Fatal(e.Start("0.0.0.0:3000"))
}
