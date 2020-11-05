package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
)

type HistoriaClinica struct {
	Base
	DocumentoID uuid.UUID      `gorm:"not null" json:"documento_id"`
	Total       float64        `gorm:"not null" json:"total" validate:"gte=0"`
	Abono       float64        `gorm:"not null" json:"abono" validate:"gte=0"`
	Terminado   bool           `gorm:"not null;default:false" json:"terminado"`
	Observaciones string `json:"observaciones"`
	Contenido   datatypes.JSON `json:"contenido"`
	Evolucion   datatypes.JSON `json:"evolucion"`
	Timeline    datatypes.JSON `json:"timeline"`
}
