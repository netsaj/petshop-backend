package models

type Tercero struct {
	Base
	Nombre    string    `gorm:"size: 255;not null" json:"nombre" validate:"required,gte=2"`
	Cedula    uint64    `gorm:"not null" sql:"not null;" validate:"required,gte=0" json:"cedula"`
	Telefono  string    `json:"telefono" gorm:"size:50"`
	Celular   string    `json:"celular" gorm:"size:50"`
	Direccion string    `json:"direccion" gorm:"size: 255" validate:"required,gte=0"`
	Barrio    string    `json:"barrio" gorm:"size: 255",validate:"required,gte=2"`
	Email     string    `json:"email,omitempty" validate:"omitempty,email" `
	Mascotas  []Mascota `gorm:"foreignkey:TerceroID;association_foreignkey:ID" json:"mascotas"`
	IsCliente bool      `gorm:"not null" json:"is_cliente"`
}

func (Tercero) TableName() string {
	return "terceros"
}
