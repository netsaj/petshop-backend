package utils

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/lib/pq"
)

func ReturnError(err error, c echo.Context) error {
	code, body := ErrorHandler(err)
	return c.JSON(code, body)
}

func ErrorHandler(err error) (int, map[string]interface{}) {
	fmt.Println(reflect.TypeOf(err))
	fmt.Println(reflect.TypeOf(err) == reflect.TypeOf(&echo.HTTPError{}))
	fmt.Println(reflect.TypeOf(&echo.HTTPError{}))
	fmt.Println(err)
	switch reflect.TypeOf(err) {
	case reflect.TypeOf(ValidatorError{}):
		{
			print("entre por validateError")
			return http.StatusUnprocessableEntity, map[string]interface{}{
				"error": err.(ValidatorError).Message,
			}
		}
	// si el error es de tipo pq (Postgres)
	case reflect.TypeOf(&pq.Error{}):
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
	case reflect.TypeOf(gorm.Errors{}):
		return http.StatusPreconditionFailed, map[string]interface{}{
			"error": err.Error(),
		}
	default:
		if "*echo.HTTPError" == fmt.Sprint(reflect.TypeOf(err)) {
			httpError := err.(*echo.HTTPError)
			if err.Error() == "" {
				return http.StatusBadRequest , map[string]interface{}{
					"error": "No se ha subido ningún archivo al servidor.",
				}
			}
			return httpError.Code, map[string]interface{}{
				"error": spew.Sdump(httpError.Message),
			}
		}
		return http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		}
	}
}
