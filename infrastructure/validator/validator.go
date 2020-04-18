package validator

import (
	"github.com/dmitrymatviets/myhood/pkg"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
	trans    ut.Translator
}

func (v *Validator) ValidateStruct(str interface{}) error {
	err := v.validate.Struct(str)
	if err == nil {
		return nil
	}

	validationErr := err.(validator.ValidationErrors)[0]
	return pkg.NewValidationErr(validationErr.Translate(v.trans), err)
}

func NewValidator() *Validator {
	validate := validator.New()

	// Пока есть баги в русской локали. Может паниковать :'(
	/*
		locale := ru_RU.New()
		uni := ut.New(locale, locale)
		trans, _ := uni.GetTranslator(locale.Locale())
		_ = ru.RegisterDefaultTranslations(validate, trans)
	*/

	return &Validator{validate, nil}
}
