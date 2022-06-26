package utils

import (
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

type IError struct {
    Field string
    Tag   string
    Value string
}

func ValidateAddCompany(w http.ResponseWriter,c *http.Request) error {
    var errors []*IError
	body, err := ioutil.ReadAll(c.Body)

    Validator.Struct(body)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            var el IError
            el.Field = err.Field()
            el.Tag = err.Tag()
            el.Value = err.Param()
            errors = append(errors, &el)
        }

		return err
	}
	return nil
}