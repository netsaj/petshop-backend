package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Vacunacion struct {
	Base
	DocumentoID   uuid.UUID `gorm:"not null" json:"documento_id"`
	GrupoVacunaID uint      `gorm:"not null" json:"grupo_vacuna_id"`
	VacunaID      uint      `gorm:"not null" json:"vacuna_id"`
	Revacunacion  time.Time `gorm:"not null" json:"revacunacion"`
	Total         float64   `gorm:"not null" json:"total" validate:"gte=0"`
	Abono         float64   `gorm:"not null" json:"abono" validate:"gte=0"`
	Terminado     bool      `gorm:"not null;default:false" json:"terminado"`

	GrupoVacuna GrupoVacuna `validate:"-" gorm:"foreignkey:GrupoVacunaID" json:"grupo_vacuna,omitempty"`
	Vacuna Vacuna `validate:"-" gorm:"foreignkey:VacunaID" json:"vacuna,omitempty"`
}


func (Vacunacion) TableName() string {
	return "vacunaciones"
}
