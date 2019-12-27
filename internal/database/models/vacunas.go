package models

import (
	"github.com/jinzhu/gorm"
)

/**
Agrupamos las vacunas
*/
type GrupoVacuna struct {
	gorm.Model
	Nombre  string   `gorm:"not null;unique_index:grupo_vacuna_uk" validate:"gte=1" json:"nombre"`
	Especie string   `gorm:"not null;unique_index:grupo_vacuna_uk" validate:"gte=1" json:"especie"`
	Vacunas []Vacuna `gorm:"many2many:grupos_vacuna_vacunas;" json:"vacunas" validate:"-"`
}

func (GrupoVacuna) TableName() string {
	return "grupos_vacunas"
}

type Vacuna struct {
	gorm.Model
	Nombre       string `gorm:"not null" validate:"gte=1" json:"nombre"`
	Descripcion  string `gorm:"not null" json:"descripcion"`
	Composicion  string `gorm:"not null" json:"composicion"`
	ParaAdulto   bool   `gorm:"not null;default:false" json:"para_adulto"`
	ParaCachorro bool   `gorm:"not null;default:false" json:"para_cachorro"`

	GrupoVacuna []GrupoVacuna `gorm:"many2many:grupos_vacuna_vacunas;" json:"grupo_vacuna" validate:"-"`
}

func (Vacuna) TableName() string {
	return "vacunas"
}
