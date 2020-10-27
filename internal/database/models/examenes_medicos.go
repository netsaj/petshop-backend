package models

type Examenes struct {
	Base
	Nombre      string `gorm:"not null;unique_index:examen_nombre_uk" json:"nombre"`
	Descripcion string `gorm:"not null" json:"descripcion"`
}

func (Examenes) TableName() string {
	return "examenes"
}
