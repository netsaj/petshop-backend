package models

import uuid "github.com/satori/go.uuid"

type Peluqueria struct {
	Base
	DocumentoID   uuid.UUID `gorm:"not null" json:"documento_id"`
	Terminado     bool      `gorm:"not null;default:false" json:"terminado"`
	Observaciones string    `gorm:"not null" json:"observaciones"`
	Total         float64   `gorm:"not null" json:"total" validate:"gte=0"`
	Abono         float64   `gorm:"not null" json:"abono" validate:"gte=0"`
}

func (Peluqueria) TableName() string {
	return "peluqueadas"
}
