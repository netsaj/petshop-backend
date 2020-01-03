package models

import "github/netsaj/petshop-backend/internal/utils"

type Usuario struct {
	Base
	Nombres   string `gorm:"size:225;not null" json:"nombres"`
	Apellidos string `gorm:"size:225;not null" json:"apellidos"`
	Username  string `gorm:"size:100;not null;unique_index:username_uk" json:"nombre_usuario"`
	Password  string `gorm:"not null" json:"password"`
	Rol       string `gorm:"default:'Operativo'" json:"rol"`
}

func (Usuario) TableName() string {
	return "usuarios"
}

func (u Usuario) CheckPassword(password string) bool {
	return utils.CheckPasswordHash(password, u.Password)
}
func (u *Usuario) SetPassword(password string) bool {
	var err error
	u.Password, err = utils.HashPassword(password)
	return err == nil
}

func (u Usuario) IsAdmin() bool {
	return u.Rol == "Administrador"
}
