package clientes_v1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
)

type searchResult struct {
	clientes []models.Tercero
}

func Search(c echo.Context) error {
	// construyo la where
	page := 0
	size := 10
	if param := c.QueryParam("page"); param != "" {
		fmt.Printf("page: %v", param)
		if x, e := strconv.Atoi(param); e == nil {
			page = x - 1
		}
	}
	if param := c.QueryParam("size"); param != "" {
		fmt.Printf("page: %v", param)
		if x, e := strconv.Atoi(param); e == nil {
			size = x
		}
	}
	// calculamos el offset a partir de la pagina y el tamaño
	offset := page * size

	where := ""
	criteria := c.QueryParam("q")
	if len(criteria) > 0 {
		for _, q := range strings.Split(criteria, " ") {
			if len(where) > 0 {
				where += " AND "
			}
			where += "( " +
				"mascotas.rfid_card_id = '" + q + "' or terceros.nombre ILIKE  '%" + q + "%' or CAST(terceros.cedula as TEXT) ILIKE  '%" + q + "%' or terceros.telefono ILIKE  '%" + q + "%' or terceros.celular ILIKE  '%" + q + "%' or terceros.direccion ILIKE  '%" + q + "%' or terceros.email ILIKE  '%" + q + "%' " +
				"or mascotas.nombre ILIKE  '%" + q + "%' or CAST(mascotas.raza as TEXT) ILIKE  '%" + q + "%' or mascotas.especie ILIKE  '%" + q + "%' or mascotas.color ILIKE  '%" + q + "%' or mascotas.sexo ILIKE  '%" + q + "%' " +
				")"
		}
	}
	db := database.GetConnection()
	defer db.Close()
	var mascotas []models.Mascota
	var count int
	query := db.Model(mascotas).Joins("left join terceros on terceros.id = mascotas.tercero_id")
	if where != "" {
		query = query.Where(where)
	}
	query.Count(&count)
	fmt.Printf("resultados para la busqueda: %v \n", count)
	fmt.Printf("pagina: %v - Tamaño: %v ", page, size)
	query.Preload("Tercero").Offset(offset).Limit(size).Find(&mascotas)
	return c.JSON(200, map[string]interface{}{
		"count":     len(mascotas),
		"resultado": &mascotas,
		"pagina":    page + 1,
		"tamaño":    size,
		"total":     count,
	})
}
