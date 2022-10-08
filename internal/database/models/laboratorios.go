package models

import uuid "github.com/satori/go.uuid"

type ExamenLaboratorio struct {
	Base
	Observaciones       string                `gorm:"not null" json:"observaciones"`
	Examen              string                `gorm:"not null" json:"examen"`
	Total               float64               `gorm:"not null" json:"total" validate:"gte=0"`
	Abono               float64               `gorm:"not null" json:"abono" validate:"gte=0"`
	DocumentoID         uuid.UUID             `gorm:"not null" json:"documento_id"`
	Terminado           bool                  `gorm:"not null;default:false" json:"terminado"`
	ArchivosLaboratorio []ArchivosLaboratorio `gorm:"foreignkey:ExamenLaboratorioID,association_foreignkey:ID" json:"adjuntos"`
}

func (ExamenLaboratorio) TableName() string {
	return "examenes_laboratorio"
}

type ArchivosLaboratorio struct {
	ID                  string            `gorm:"primary_key;" json:"id"`
	ArchivoID           uuid.UUID         `gorm:"not null;" json:"archivo_id"`
	ExamenLaboratorioID uuid.UUID         `gorm:"not null;" json:"examen_laboratorio_id"`
	ExamenLaboratorio   ExamenLaboratorio `gorm:"foreignkey:ExamenLaboratorioID,association_foreignkey:ID" json:"examen_laboratorio"`
	Archivo             Archivo           `gorm:"foreignkey:ArchivoID,association_foreignkey:ID" json:"archivo"`
}
