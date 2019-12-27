package models

import (
	"github.com/jinzhu/gorm"
)

/**
Agrupamos las vacunas
*/
type GrupoDesparacitante struct {
	gorm.Model
	Nombre  string   `gorm:"not null;unique_index:grupos_desparacitantes_uk" validate:"gte=1" json:"nombre"`
	Especie string   `gorm:"not null;unique_index:grupos_desparacitantes_uk" validate:"gte=1" json:"especie"`
	Desparacitantes []Desparacitante `gorm:"many2many:grupos_vacuna_vacunas;" json:"vacunas" validate:"-"`
}

func (GrupoDesparacitante) TableName() string {
	return "grupos_desparacitantes"
}

type Desparacitante struct {
	gorm.Model
	Nombre       string `gorm:"not null" validate:"gte=1" json:"nombre"`
	Descripcion  string `gorm:"not null" json:"descripcion"`
	Composicion  string `gorm:"not null" json:"composicion"`
	ParaAdulto   bool   `gorm:"not null;default:false" json:"para_adulto"`
	ParaCachorro bool   `gorm:"not null;default:false" json:"para_cachorro"`

	GrupoVacuna []GrupoVacuna `gorm:"many2many:grupos_vacuna_vacunas;" json:"grupo_vacuna" validate:"-"`
}

func (Desparacitante) TableName() string {
	return "desparacitantes"
}
