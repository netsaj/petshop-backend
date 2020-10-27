package models

/*
es una estructura base usada para definir las propiedades basicas de un archivo subido al servidor
*/
type Archivo struct {
	Base
	Nombre      string `gorm:"not null" json:"nombre"`
	Extension   string `gorm:"not null" json:"extension"`
	ContentType string `gorm:"not null" json:"content_type"`
	Tamaño      int64  `gorm:"not null;default:0" json:"tamaño"`
	Ruta        string `gorm:"not null" json:"ruta"`
}

func (Archivo) TableName() string {
	return "archivos"
}
