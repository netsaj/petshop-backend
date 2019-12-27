package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	"github/netsaj/petshop-backend/internal/database"
)

type DocumentoError struct {
	error
	Code  uint
	Error map[string]interface{} `json:"error"`
}

type Documento struct {
	Base
	PrefijoID uuid.UUID `gorm:"not null" json:"prefijo_id"`
	Numero    uint32    `gorm:"not null" json:"numero" validate:"gte=0"`
	TerceroID uuid.UUID `gorm:"not null" json:"tercero_id"`
	MascotaID uuid.UUID `gorm:"not null" json:"mascota_id"`
	UsuarioID uuid.UUID `gorm:"not null" json:"usuario_id"`
	Tipo      string    `gorm:"not null", json:"tipo"`
	Subtipo   string    `gorm:"not null" json:"subtipo"`
	Total     float64   `gorm:"not null;default:0" json:"total"`

	// agregaciones
	Prefijo    Prefijo    `validate:"-" gorm:"foreignkey:PrefijoID" json:"prefijo,omitempty"`
	Peluqueria Peluqueria `validate:"-" gorm:"foreignkey:DocumentoID" json:"peluqueria,omitempty"`
	Vacunacion Vacunacion `validate:"-" gorm:"foreignkey:DocumentoID" json:"vacunacion,omitempty"`
	Usuario    Usuario    `validate:"-" gorm:"foreignkey:UsuarioID" json:"usuario,omitempty"`
	Tercero    Tercero    `validate:"-" gorm:"foreignkey:TerceroID" json:"tercero,omitempty"`
	Mascota    Mascota    `validate:"-" gorm:"foreignkey:MascotaID" json:"mascota,omitempty"`
}

func (Documento) TableName() string {
	return "documentos"
}

func (d *Documento) calcularTotal() {
	if !reflect.DeepEqual(d.Peluqueria, Peluqueria{}) {
		d.Total += d.Peluqueria.Total
	}
	if !reflect.DeepEqual(d.Vacunacion, Vacunacion{}) {
		d.Total += d.Vacunacion.Total
	}
}

func (d *Documento) CrearDocumentoServicio() error {
	db := database.GetConnection()
	defer db.Close()

	if d.PrefijoID == uuid.Nil {
		d.PrefijoID = uuid.FromStringOrNil(PREFIJO_DEFAULT)
	}
	fmt.Printf("buscando prefijo : %v \n", d.PrefijoID)
	var prefijo Prefijo // Consultamos el prefijo para obtener el numero actual
	if db.First(&prefijo, "id = ?", d.PrefijoID).RecordNotFound() {
		return gorm.ErrRecordNotFound
	}
	d.Tipo = "venta"       // el documento sera de tipo venta
	d.Subtipo = "servicio" // subtipo servicio para asociar peluquería, vacunación y desparasitacion
	esNuevo := d.ID == uuid.Nil
	if esNuevo {
		d.Numero = prefijo.Actual
	}
	d.calcularTotal()
	if result := db.Save(&d); result.Error != nil {
		return result.Error
	} else if esNuevo {
		prefijo.Incrementar()
	}
	if reflect.DeepEqual(d.Peluqueria, Peluqueria{}) {
		fmt.Println("ignorando Peluqueria, estructura vacía")
	} else {
		db.Save(&d.Peluqueria)
		var calendario = new(Calendario)
		calendario.FechaAgendada = time.Now().Add(time.Duration(time.Hour * 24 * 60))
		calendario.Tipo = "Peluquería"
		calendario.TerceroID = d.TerceroID
		calendario.MascotaID = d.MascotaID
		calendario.Terminado = false
		calendario.ObservacionesAbierto = "Observaciones en servicio anterior: " + d.Peluqueria.Observaciones
		db.Save(&calendario)
	}
	if reflect.DeepEqual(d.Vacunacion, Vacunacion{}) {
		fmt.Println("ignorando Vacunacion, estructura vacía")
	} else {
		db.Save(&d.Vacunacion).Preload("Vacuna").Find(&d.Vacunacion)
		var calendario = new(Calendario)
		calendario.FechaAgendada = d.Vacunacion.Revacunacion
		calendario.Tipo = "Vacunación"
		calendario.TerceroID = d.TerceroID
		calendario.MascotaID = d.MascotaID
		calendario.Terminado = false
		calendario.ObservacionesAbierto = fmt.Sprintf("Vacuna anterior: %v", d.Vacunacion.Vacuna.Nombre)
		calendario.DocumentoID = d.ID
		db.Save(&calendario)
	}
	spew.Println(&d)
	return nil
}
