package models

import "github.com/jinzhu/gorm"

type Barrio struct {
	gorm.Model
	Municipio string `gorm:"not null;" json:"municipio"`
	Nombre    string `gorm:"not null;" json:"nombre"`
}

func (Barrio) TableName() string {
	return "barrios"
}
