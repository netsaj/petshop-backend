package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Mascota struct {
	Base
	TerceroID       uuid.UUID `gorm:"not null;" validate:"required" json:"tercero_id"`
	Nombre          string    `gorm:"size:255;not null" validate:"required,gte=2" json:"nombre"`
	Especie         string    `gorm:"size:10;not null" validate:"required,gte=2,alpha" json:"especie"`
	Raza            string    `gorm:"size:50;not null" validate:"required,gte=2" json:"raza"`
	Peso            float32   `gorm:"not null;" validate:"gte=0.001" json:"peso"`
	Color           string    `gorm:"size:50;not null" validate:"required,gte=2" json:"color"`
	Sexo            string    `gorm:"size:10;not null" validate:"required,gte=2,alpha" json:"sexo"`
	Edad            float32   `gorm:"not null;'" validate:"gte=0" json:"edad"`
	FechaNacimiento time.Time `gorm:"not null;" json:"fecha_nacimiento" `
	RfidCardID      string    `json:"rfid_card_id"`
	HistoriaClinica string    `json:"historia_clinica"`

	Tercero Tercero `gorm:"foreignkey:TerceroID;association_foreignkey:ID" json:"tercero" validate:"-"`
}

func (Mascota) TableName() string {
	return "mascotas"
}
