package migrations

import (
	"fmt"
	"math"
	"path/filepath"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/davecgh/go-spew/spew"
	uuid "github.com/satori/go.uuid"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func createAdminIfNotExist() {
	dbClient := database.GetConnection()
	defer dbClient.Close()

	// add admins user if no exist
	var user models.Usuario
	// find User with username `admins`
	if err := dbClient.Where("username = ?", "admin").First(&user).Error; err != nil {
		user.ID = uuid.FromStringOrNil("87ec255a-bdbb-4b0d-b0e1-ab886f229777")
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
		// struct for load `barrios`
		type Temp struct {
			Barrios []models.Barrio
		}
		// we initialize our Users array
		var data Temp
		if err := utils.LoadFileJSON(filepath.FromSlash(filepath.FromSlash("resources/barrios.json")), &data); err != nil {
			panic(err)
		}
		bar := pb.StartNew(len(data.Barrios))
		defer bar.Finish()
		for i := 0; i < len(data.Barrios); i++ {
			db.Create(&data.Barrios[i])
			bar.Increment()
			time.Sleep(time.Millisecond)
		}
		db.Model(&models.Barrio{}).Count(&count)

	}
	fmt.Printf("Total de barrios: %v \n", count)
}

func agregarGruposVacunas() {
	db := database.GetConnection()
	defer db.Close()
	var count int
	db.Model(&models.GrupoVacuna{}).Count(&count)
	if count == 0 {
		// struct for load `barrios`
		var data []models.GrupoVacuna
		// we initialize our Users array
		if err := utils.LoadFileJSON(filepath.FromSlash(filepath.FromSlash("resources/grupo_vacuna.json")), &data); err != nil {
			panic(err)
		}
		fmt.Printf("Grupos encontrados en el archivo: %v", len(data))
		bar := pb.StartNew(len(data))
		defer bar.Finish()
		for i := 0; i < len(data); i++ {
			db.Create(&data[i])
			bar.Increment()
			time.Sleep(time.Millisecond)
		}
		db.Model(&models.GrupoVacuna{}).Count(&count)

	}
	fmt.Printf("Total de grupos de vacunas: %v \n", count)
}

func agregarVacunas() {
	db := database.GetConnection()
	defer db.Close()
	var count int
	db.Model(&models.Vacuna{}).Count(&count)
	if count == 0 {
		// struct for load `barrios`
		var data []models.Vacuna
		// we initialize our Users array
		if err := utils.LoadFileJSON(filepath.FromSlash(filepath.FromSlash("resources/vacunas.json")), &data); err != nil {
			panic(err)
		}
		fmt.Printf("Vacunas encontradas en el archivo: %v", len(data))
		bar := pb.StartNew(len(data))
		defer bar.Finish()
		for i := 0; i < len(data); i++ {
			for j := 0; j < len(data[i].GrupoVacuna); j++ {
				id := data[i].GrupoVacuna[j].ID
				db.First(&data[i].GrupoVacuna[j], "id = ?", id)
			}
			db.Model(&models.Vacuna{}).Save(&data[i])
			bar.Increment()
			time.Sleep(time.Millisecond)
		}
		db.Model(&models.Vacuna{}).Count(&count)

	}
	fmt.Printf("Total de vacunas: %v \n", count)
}

func prefijoDefault() {
	db := database.GetConnection()
	defer db.Close()
	var prefijo models.Prefijo
	if db.First(&prefijo, "id = ?", models.PREFIJO_DEFAULT).RecordNotFound() {
		prefijo.ID = uuid.FromStringOrNil(models.PREFIJO_DEFAULT)
		prefijo.Nombre = "Prefijo servicios"
		prefijo.Codigo = "SV"
		prefijo.Descripcion = "Prefijo creado por defecto para los servicios."
		prefijo.Inicio = 1
		prefijo.Fin = math.MaxInt32
		prefijo.Actual = 1
		if result := db.Save(&prefijo); result.Error != nil {
			fmt.Println(result.Error)
		}
		fmt.Sprintf("prefijo creado : %s", spew.Sdump(prefijo))
		fmt.Println("prefijo por default creado")

	}
}

func agregarGruposDesparasitantes() {
	db := database.GetConnection()
	defer db.Close()
	var count int
	db.Model(&models.GrupoDesparasitante{}).Count(&count)
	if count == 0 {
		// struct for load `GrupoDesparacitante`
		var data []models.GrupoDesparasitante
		// we initialize our Users array
		if err := utils.LoadFileJSON(filepath.FromSlash(filepath.FromSlash("resources/grupo_desparasitantes.json")), &data); err != nil {
			panic(err)
		}
		fmt.Printf("Grupos encontrados en el archivo: %v", len(data))
		bar := pb.StartNew(len(data))
		defer bar.Finish()
		for i := 0; i < len(data); i++ {
			db.Create(&data[i])
			bar.Increment()
			time.Sleep(time.Millisecond)
		}
		db.Model(&models.GrupoVacuna{}).Count(&count)

	}
	fmt.Printf("Total de grupos de desparasitantes: %v \n", count)
}

func agregarDesparasitantes() {
	db := database.GetConnection()
	defer db.Close()
	var count int
	db.Model(&models.Desparasitante{}).Count(&count)
	if count == 0 {
		// struct for load `barrios`
		var data []models.Desparasitante
		// we initialize our Users array
		if err := utils.LoadFileJSON(filepath.FromSlash(filepath.FromSlash("resources/desparasitantes.json")), &data); err != nil {
			panic(err)
		}
		fmt.Printf("Desparasitante encontradas en el archivo: %v", len(data))
		bar := pb.StartNew(len(data))
		defer bar.Finish()
		for i := 0; i < len(data); i++ {
			for j := 0; j < len(data[i].GruposDesparasitante); j++ {
				id := data[i].GruposDesparasitante[j].ID
				db.First(&data[i].GruposDesparasitante[j], "id = ?", id)
			}
			db.Model(&models.Desparasitante{}).Save(&data[i])
			bar.Increment()
			time.Sleep(time.Millisecond)
		}
		db.Model(&models.Desparasitante{}).Count(&count)

	}
	fmt.Printf("Total de Desparasitante: %v \n", count)
}

func agregarExamenes() {
	db := database.GetConnection()
	defer db.Close()
	var count int
	db.Model(&models.Examenes{}).Count(&count)
	if count == 0 {
		// struct for load `barrios`
		var data []models.Examenes
		// we initialize our Users array
		if err := utils.LoadFileJSON(filepath.FromSlash(filepath.FromSlash("resources/examenes_laboratorio.json")), &data); err != nil {
			panic(err)
		}
		fmt.Printf("Examenes encontrados en el archivo: %v", len(data))
		bar := pb.StartNew(len(data))
		defer bar.Finish()
		for i := 0; i < len(data); i++ {
			db.Model(&models.Examenes{}).Save(&data[i])
			bar.Increment()
			time.Sleep(time.Millisecond)
		}
		defer bar.Finish()
	}
	db.Model(&models.Examenes{}).Count(&count)
	fmt.Printf("Total de Examenes medicos: %v \n", count)
}
