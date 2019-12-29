package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Desparasitacion struct {
	Base
	DocumentoID           uuid.UUID `gorm:"not null" json:"documento_id"`
	GrupoDesparasitanteID uint      `gorm:"not null" json:"grupo_desparasitante_id"`
	DesparasitanteID      uint      `gorm:"not null" json:"desparasitante_id"`
	Redesparacitacion     time.Time `gorm:"not null" json:"redesparacitacion"`
	Dosis                 string    `gorm:"not null" json:"dosis"`
	Total                 float64   `gorm:"not null" json:"total"`
	Abono                 float64   `gorm:"not null" json:"abono"`
	Terminado             bool      `gorm:"not null" json:"terminado"`

	GrupoDesparasitante GrupoDesparasitante `validate:"-" gorm:"foreignkey:GrupoDesparasitanteID" json:"grupo_desparasitante,omitempty"`
	Desparasitante      Desparasitante      `validate:"-" gorm:"foreignkey:DesparasitanteID" json:"desparasitante,omitempty"`
}

func (Desparasitacion) TableName() string {
	return "desparasitaciones"
}
