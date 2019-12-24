package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/locales"
	ES "github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	es_translation "github/netsaj/petshop-backend/internal/utils/translations/es"
	"gopkg.in/go-playground/validator.v9"
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

		return errors.New(string(e))
	}
	return nil
}
