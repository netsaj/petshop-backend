package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Calendario struct {
	Base
	FechaAgendada        time.Time `gorm:"not null" json:"fecha_agendada"`
	FechaCierre          time.Time `gorm:"default:NULL" json:"fecha_cierre"`
	Tipo                 string    `gorm:"not null" json:"tipo"`
	TerceroID            uuid.UUID `gorm:"not null" json:"tercero_id"`
	MascotaID            uuid.UUID `gorm:"not null" json:"mascota_id"`
	Terminado            bool      `gorm:"not null;default:false;" json:"terminado"`
	ObservacionesAbierto string    `gorm:"not null;" json:"observaciones_abierto"`
	ObservacionesCerrado string    `gorm:"not null;" json:"observaciones_cerrado"`
	UsuarioCierreID      uuid.UUID `gorm:"default:NULL" json:"usuario_cierre_id"`
	DocumentoID          uuid.UUID `gorm:"default:NULL" json:"documento_id"`

	Tercero       Tercero   `validate:"-" gorm:"foreignkey:TerceroID" json:"tercero"`
	Mascota       Mascota   `validate:"-" gorm:"foreignkey:MascotaID" json:"mascota"`
	Documento     Documento `validate:"-" gorm:"foreignkey:DocumentoID" json:"documento"`
	UsuarioCierre Usuario   `validate:"-" gorm:"foreignkey:UsuarioCierreID" json:"usuario_cierre"`
}

func (Calendario) TableName() string {
	return "calendarios"
}
