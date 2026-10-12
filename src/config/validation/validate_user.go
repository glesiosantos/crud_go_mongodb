package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")

		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidatorUserError(validation_err error) *errorsc.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return errorsc.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {

		errorCauses := []errorsc.Cause{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := errorsc.Cause{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}
		return errorsc.NewValidationError("Some fields are invalid", errorCauses)
	} else {
		return errorsc.NewBadRequestError("Error trying to convert fields")
	}

}
