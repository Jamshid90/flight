package validation

import (
	"errors"
	"reflect"
	"strings"

	errpkg "github.com/Jamshid90/flight/internal/errors"
	english "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validatorEn "github.com/go-playground/validator/v10/translations/en"
)

func Validator(s interface{}) error {
	var (
		eng      = english.New()
		uni      = ut.New(eng, eng)
		validate = validator.New()
	)

	trans, found := uni.GetTranslator("en")
	if !found {
		return errors.New("Validator translator not found")
	}

	if err := validatorEn.RegisterDefaultTranslations(validate, trans); err != nil {
		return err
	}

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	if errs, ok := err.(validator.ValidationErrors); ok {
		errValidation := errpkg.NewErrValidation()
		errValidation.Err = err
		for _, e := range errs {
			errValidation.Errors[e.Field()] = strings.Replace(e.Translate(trans), e.Field(), "", 1)
		}
		return errValidation
	}
	return nil
}
