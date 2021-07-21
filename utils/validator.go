package utils

import (
	"fmt"
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type ValidatorInterface interface {
	Struct(s interface{}) error
}

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	v := &Validator{
		validate: validator.New(),
	}
	v.validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	return v
}

func (v *Validator) Struct(s interface{}) error {
	errs := v.validate.Struct(s)

	if errs != nil {
		fieldErrors, _ := errs.(validator.ValidationErrors)
		for _, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				return fmt.Errorf("%s is a required field", err.Field())
			case "max":
				return fmt.Errorf("%s must be maximum of %s in length", err.Field(), err.Param())
			case "min":
				return fmt.Errorf("%s must be minimum of %s in length", err.Field(), err.Param())
			case "alphanum":
				return fmt.Errorf("%s must containt only alphanumeric characters", err.Field())
			case "containsany":
				return fmt.Errorf("%s must containt at least one of %s characters", err.Field(), err.Param())
			default:
				return fmt.Errorf("something wrong on %s; %s", err.Field(), err.Tag())
			}
		}
	}

	return nil
}
