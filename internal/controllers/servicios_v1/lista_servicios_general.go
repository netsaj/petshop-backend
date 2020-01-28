package servicios_v1

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"

	"github/netsaj/petshop-backend/internal/database"
	"github/netsaj/petshop-backend/internal/database/models"
	"github/netsaj/petshop-backend/internal/utils"
)

func ListarServicios(c echo.Context) error {
	page := 0
	size := 10
	fechaInicio := c.QueryParam("fecha_inicio")     // fecha de inicio para el filtrado
	fechaFinal := c.QueryParam("fecha_fin")         // fecha final para el filtrado
	if param := c.QueryParam("page"); param != "" { // recupero la pagina de la peticion, si no viene definida dejo la pagina en 1
		fmt.Printf("page: %v", param)
		if x, e := strconv.Atoi(param); e == nil {
			page = x - 1
		}
	}
	if param := c.QueryParam("size"); param != "" { // reviso el numero de elementos por pagina
		fmt.Printf("page: %v", param)
		if x, e := strconv.Atoi(param); e == nil {
			size = x
		}
	}
	// calculamos el offset a partir de la pagina y el tamaño
	offset := page * size // calculo el offset de la paginanacion.
	criteria := c.QueryParam("q")
	where := "" // wuere adicional para filtrar si llega un criterio de busqueda.
	for _, q := range strings.Split(criteria, " ") {
		if len(where) > 0 {
			where += " AND "
		}
		where += "( CONCAT(prefijos.codigo,documentos.numero) ILIKE '%" + q + "%' OR terceros.nombre ILIKE  '%" + q + "%' or CAST(terceros.cedula as TEXT) ILIKE  '%" + q + "%' or terceros.telefono ILIKE  '%" + q + "%' or terceros.celular ILIKE  '%" + q + "%' or terceros.direccion ILIKE  '%" + q + "%' or terceros.barrio ILIKE  '%" + q + "%' or terceros.email ILIKE  '%" + q + "%' )"
	}
	db := database.GetConnection()
	defer db.Close()
	var documentos []models.Documento
	query := db.Model(models.Documento{}).
		Joins("left join peluqueadas on peluqueadas.documento_id = documentos.id ").
		Joins("left join vacunaciones on vacunaciones.documento_id = documentos.id ").
		Joins("left join desparasitaciones on desparasitaciones.documento_id = documentos.id ").
		Joins("left join terceros on terceros.id = documentos.tercero_id").
		Joins("left join mascotas on mascotas.id = documentos.mascota_id").
		Joins("left join prefijos on prefijos.id = documentos.prefijo_id").
		Joins("left join examenes_laboratorio on examenes_laboratorio.documento_id = documentos.id ").
		Preload("ExamenLaboratorio").
		Preload("Peluqueria").
		Preload("ExamenLaboratorio").
		Preload("Vacunacion").
		Preload("Vacunacion.Vacuna").
		Preload("Vacunacion.GrupoVacuna").
		Preload("Desparasitacion").
		Preload("Desparasitacion.Desparasitante").
		Preload("Desparasitacion.GrupoDesparasitante").
		Preload("Usuario").
		Preload("Tercero").
		Preload("Mascota").
		Preload("Prefijo").
		Order("documentos.created_at desc").
		Where(where).
		Group("documentos.id").
		Where("documentos.tipo = 'venta' and documentos.subtipo = 'servicio' and " +
			"(peluqueadas.id != '00000000-0000-0000-0000-000000000000' OR vacunaciones.id != '00000000-0000-0000-0000-000000000000'  OR desparasitaciones.id != '00000000-0000-0000-0000-000000000000'  OR  examenes_laboratorio.id != '00000000-0000-0000-0000-000000000000') ")

	if fechaInicio != "" { // agrego la fecha de inicio si esta definida, en la consulta
		fechaInicio = strings.Split(fechaInicio, "T")[0]
	} else {
		fechaInicio = "2019-01-01"
	}
	if fechaFinal != "" { // agrego la fecha de fin si esta definida, en la consulta.
		fechaFinal = strings.Split(fechaFinal, "T")[0]
	} else {
		fechaFinal = strings.Split(time.Now().String(), "T")[0]
	}
	query = query.Where(" (DATE(documentos.created_at) >= ? AND DATE(documentos.created_at) <= ? ) ", fechaInicio, fechaFinal)
	var count uint
	if result := query.Count(&count); result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}

	if result := query.Limit(size).Offset(offset).
		Find(&documentos);
		result.Error != nil {
		return utils.ReturnError(result.Error, c)
	}

	fmt.Printf("total servicios : %v \n", count)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"page":       page + 1,
		"size":       size,
		"total":      count,
		"documentos": documentos,
	})
}
