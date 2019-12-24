package models

import "github/netsaj/petshop-backend/internal/utils"

type Usuario struct {
	Base
	Nombres   string `gorm:"size:225, not null" json:"nombres"`
	Apellidos string `gorm:"size:225, not null" json:"apellidos" validate:"alpha"`
	Username  string `gorm:"size:100, not null" json:"nombre_usuario"`
	Password  string `gorm:"not null" json:"-" validate:"gte=8"`
	Rol       string `gorm:"default:'operativo'" json:"rol"`
}

func (Usuario) TableName() string {
	return "usuarios"
}

func (u Usuario) CheckPassword(password string) bool {
	return utils.CheckPasswordHash(password, u.Password)
}

func (u Usuario) IsAdmin() bool {
	return u.Rol == "Administrador"
}
