package utils

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-playground/locales"
	ES "github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"

	es_translation "github/netsaj/petshop-backend/internal/utils/translations/es"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
	esCo     locales.Translator
)

func init() {

}

type CustomValidator struct {
	Validator *validator.Validate
}

func New() *CustomValidator {
	esCo = ES.New()
	uni = ut.New(esCo, esCo)
	// this is usually know or extracted from http 'Accept-Language' header
	trans, _ = uni.GetTranslator("es")
	validate = validator.New()
	es_translation.RegisterDefaultTranslations(validate, trans)
	return &CustomValidator{Validator: validate}
}

type ValidatorError struct {
	error
	Message string
}

func (ve ValidatorError) Error() string {
	return spew.Sprint(ve.Message)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)
		translateErrors := make(map[string]string)
		for _, e := range errs {
			key := e.Field()
			translateErrors[fmt.Sprint(key)] = e.Translate(trans)
		}
		e, _ := json.Marshal(translateErrors)
		fmt.Println(string(e))
		var message map[string]string
		json.Unmarshal([]byte(e), &message)
		msg := ""
		for key, value := range message {
			msg += fmt.Sprintf("%s : %v \n", key, value)
		}
		return ValidatorError{Message: msg}
	}
	return nil
}
