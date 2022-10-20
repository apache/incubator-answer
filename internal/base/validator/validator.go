package validator

import (
	"errors"
	"reflect"

	"github.com/go-playground/locales"
	english "github.com/go-playground/locales/en"
	zhongwen "github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/go-playground/validator/v10/translations/zh"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/base/translator"
	myErrors "github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/i18n"
	"github.com/segmentfault/pacman/log"
)

// MyValidator my validator
type MyValidator struct {
	Validate *validator.Validate
	Tran     ut.Translator
	Lang     i18n.Language
}

// ErrorField error field
type ErrorField struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var (
	// GlobalValidatorMapping is a mapping from validator to translator used
	GlobalValidatorMapping = make(map[string]*MyValidator, 0)
)

func init() {
	zhTran, zhVal := getTran(zhongwen.New(), i18n.LanguageChinese.Abbr()), createDefaultValidator(i18n.LanguageChinese)
	if err := zh.RegisterDefaultTranslations(zhVal, zhTran); err != nil {
		panic(err)
	}
	GlobalValidatorMapping[i18n.LanguageChinese.Abbr()] = &MyValidator{Validate: zhVal, Tran: zhTran, Lang: i18n.LanguageChinese}

	enTran, enVal := getTran(english.New(), i18n.LanguageEnglish.Abbr()), createDefaultValidator(i18n.LanguageEnglish)
	if err := en.RegisterDefaultTranslations(enVal, enTran); err != nil {
		panic(err)
	}
	GlobalValidatorMapping[i18n.LanguageEnglish.Abbr()] = &MyValidator{Validate: enVal, Tran: enTran, Lang: i18n.LanguageEnglish}
}

func getTran(lo locales.Translator, la string) ut.Translator {
	tran, ok := ut.New(lo, lo).GetTranslator(la)
	if !ok {
		panic(ok)
	}
	return tran
}

func createDefaultValidator(la i18n.Language) *validator.Validate {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) (res string) {
		defer func() {
			if len(res) > 0 {
				res = translator.GlobalTrans.Tr(la, res)
			}
		}()
		if jsonTag := fld.Tag.Get("json"); len(jsonTag) > 0 {
			if jsonTag == "-" {
				return ""
			}
			return jsonTag
		}
		if formTag := fld.Tag.Get("form"); len(formTag) > 0 {
			return formTag
		}
		return fld.Name
	})
	return validate
}

func GetValidatorByLang(la string) *MyValidator {
	if GlobalValidatorMapping[la] != nil {
		return GlobalValidatorMapping[la]
	}
	return GlobalValidatorMapping[i18n.DefaultLang.Abbr()]
}

// Check /
func (m *MyValidator) Check(value interface{}) (errField *ErrorField, err error) {
	err = m.Validate.Struct(value)
	if err != nil {
		var valErrors validator.ValidationErrors
		if !errors.As(err, &valErrors) {
			log.Error(err)
			return nil, errors.New("validate check exception")
		}

		for _, fieldError := range valErrors {
			errField = &ErrorField{
				Key:   translator.GlobalTrans.Tr(m.Lang, fieldError.Field()),
				Value: fieldError.Translate(m.Tran),
			}
			return errField, myErrors.BadRequest(reason.RequestFormatError).WithMsg(fieldError.Translate(m.Tran))
		}
	}

	if v, ok := value.(Checker); ok {
		errField, err = v.Check()
		if err != nil {
			return errField, err
		}
	}
	return nil, nil
}

// Checker .
type Checker interface {
	Check() (errField *ErrorField, err error)
}
