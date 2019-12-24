package models

type Cliente struct {
	Base
	Nombre    string    `gorm:"size: 255;not null" json:"nombre" validate:"required,gte=2"`
	Cedula    uint64    `gorm:"not null" sql:"not null;" validate:"required,gte=999" json:"cedula"`
	Telefono  string    `json:"telefono" gorm:"size:50"`
	Celular   string    `json:"celular" gorm:"size:50"`
	Direccion string    `json:"direccion" gorm:"size: 255" validate:"required,gte=5"`
	Barrio    string    `json:"barrio" gorm:"size: 255",validate:"required,gte=2"`
	Email     string    `json:"email,omitempty" validate:"omitempty,email" `
	Mascotas  []Mascota `gorm:"foreignkey:ClienteID,association_foreignkey:ID" json:"mascotas"`
}

func (Cliente) TableName() string {
	return "clientes"
}
