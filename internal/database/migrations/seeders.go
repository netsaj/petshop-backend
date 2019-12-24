package migrations

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/schollz/progressbar/v2"
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

func createAdminIfNotExist() {
	dbClient := database.GetConnection()
	defer dbClient.Close()

	// add admins user if no exist
	var user models.Usuario
	// find User with username `admins`
	if err := dbClient.Where("username = ?", "admin").First(&user).Error; err != nil {
		user.Username = "admin"
		user.Nombres = "admins"
		user.Rol = "Administrador"
		user.Password, _ = utils.HashPassword("admin")
		if err = dbClient.Create(&user).Error; err != nil {
			print("Usuario 'admin' creado")
			spew.Dump(&user)
		}

	}

}

func agregarBarrios() {
	db := database.GetConnection()
	defer db.Close()
	var count int
	db.Model(&models.Barrio{}).Count(&count)
	if count == 0 {
		// Open our jsonFile
		var jsonFile, err = os.Open(filepath.FromSlash("resources/barrios.json"))
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Opened users.json")
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)
		// struct for load `barrios`
		type Barrios struct {
			Barrios []models.Barrio
		}
		// we initialize our Users array
		var data Barrios

		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		if err := json.Unmarshal(byteValue, &data); err != nil {
			panic(err)
		}

		fmt.Printf("barrios encontrados en el archivo: %v", data.Barrios)
		bar := progressbar.New(len(data.Barrios))
		for i := 0; i < len(data.Barrios); i++ {
			db.Create(&data.Barrios[i])
			bar.Add(1)
		}
		db.Model(&models.Barrio{}).Count(&count)

	}
	fmt.Printf("Total de barrios: %v \n", count)
}
