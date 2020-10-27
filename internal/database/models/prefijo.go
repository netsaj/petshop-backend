package models

import (
	"github/netsaj/petshop-backend/internal/database"
)

const PREFIJO_DEFAULT = "a484b2af-7a3a-411b-8967-ab2a2caba22b"

type Prefijo struct {
	Base
	Codigo      string `gorm:"not null;unique_index" json:"codigo"`
	Nombre      string `gorm:"not null;" json:"nombre"`
	Descripcion string `gorm:"not null" json:"descripcion"`
	Tipo        string `gorm:"not null" json:"tipo"`
	Inicio      uint32   `gorm:"not null" json:"inicio"`
	Fin         uint32   `gorm:"not null" json:"fin"`
	Actual      uint32   `gorm:"not null" json:"actual"`
}

func (Prefijo) TableName() string {
	return "prefijos"
}

func (u Prefijo) Incrementar() error {
	u.Actual++
	db := database.GetConnection()
	defer db.Close()
	if result := db.Save(&u); result.Error != nil {
		return result.Error
	}
	return nil
}
