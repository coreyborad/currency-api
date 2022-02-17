package validate

import (
	"reflect"

	zh_tw "github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh_tw"
)

var validateHelper *validator.Validate
var trans ut.Translator

func ValidateInit() {
	zh := zh_tw.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh_tw")

	validateHelper = validator.New()
	validateHelper.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
	zh_translations.RegisterDefaultTranslations(validateHelper, trans)
}
func Translate(errs validator.ValidationErrors) []string {
	var errList []string
	for _, e := range errs {
		// can translate each error one at a time.
		errList = append(errList, e.Translate(trans))
	}
	return errList
}
func GetValidate() *validator.Validate {
	return validateHelper
}
