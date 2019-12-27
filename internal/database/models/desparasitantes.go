package models

import "github.com/jinzhu/gorm"

type GrupoDesparasitante struct {
	gorm.Model
	Nombre          string           `gorm:"not null;unique_index:grupo_desparasitante_uk" json:"nombre" validate:"gte=1"`
	Desparacitantes []Desparasitante `gorm:"many2many:grupo_desparasitante_desparasitantes;" json:"desparasitantes"`
}

func (GrupoDesparasitante) TableName() string {
	return "grupos_desparasitantes"
}

type Desparasitante struct {
	gorm.Model
	Nombre       string `gorm:"not null" validate:"gte=1" json:"nombre"`
	Despcripcion string `gorm:"not null" json:"despcripcion"`
	Composicion  string `gorm:"not null" json:"composicion"`
	ParaAdulto   bool   `gorm:"not null;default:false" json:"para_adulto"`
	ParaCachorro bool   `gorm:"not null;default:false" json:"para_cachorro"`
	Tipo string     `gorm:"not null;default:'Liquido'" json:"tipo"`
	GruposDesparasitante []GrupoDesparacitante `gorm:"many2many:grupo_desparasitante_desparasitantes" json:"grupos_desparasitante"`
}

func (Desparasitante) TableName() string {
	return "desparasitantes"
}
