package mascotas_v1

import (
	"fmt"
	"github.com/labstack/echo"
	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"strings"
)

func Search(c echo.Context) error {
	// construyo la query
	criteria := c.QueryParam("q")
	query := ""
	for _, q := range strings.Split(criteria, " ") {
		if len(query) > 0 {
			query += " AND "
		}
		query += "( nombre ILIKE '%" + q + "%' or especie ILIKE '%" + q + "%' or raza ILIKE '%" + q + "%' or color ILIKE '%" + q + "%' or sexo ILIKE '%" + q + "%' or CAST(edad as TEXT) ILIKE '%" + q + "%' or CAST(fecha_nacimiento as TEXT)  ILIKE '%" + q + "%')"
	}
	db := database.GetConnection()
	defer db.Close()
	var mascotas []models.Mascota
	fmt.Println(db.Model(&models.Mascota{}).Preload("Tercero").Find(&mascotas, query).QueryExpr())
	return c.JSON(200, map[string]interface{}{
		"count":    len(mascotas),
		"mascotas": &mascotas,
	})
}
