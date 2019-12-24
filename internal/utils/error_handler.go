package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/lib/pq"
	"net/http"
)

var c echo.Context

func ErrorHandler(err error) (int, map[string]interface{}) {
	switch err.(type) {
	// si el error es de tipo pq (Postgres)
	case *pq.Error:
		{
			fmt.Printf("POSTGRESQL ERROR CODE: %v : %s \n", err.(*pq.Error).Code.Name(), err.(*pq.Error).Message)
			fmt.Println(err.(*pq.Error).Code.Name(), err.(*pq.Error).Where)
			if err.(*pq.Error).Code.Name() == "unique_violation" {
				return http.StatusConflict, map[string]interface{}{
					"error": err.Error(),
				}
			} else {
				return http.StatusUnprocessableEntity, map[string]interface{}{
					"error": err.Error(),
				}
			}
		}
	// si el error es de gorm (a nivel de ORM)
	case *gorm.Errors:
		return http.StatusPreconditionFailed, map[string]interface{}{
			"error": err.Error(),
		}
	default:
		return http.StatusInternalServerError, nil
	}
}
