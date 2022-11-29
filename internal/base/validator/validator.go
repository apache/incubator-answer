package validator

import (
	"errors"
	"reflect"
	"strings"

	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/go-playground/locales"
	english "github.com/go-playground/locales/en"
	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/go-playground/validator/v10/translations/zh"
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

// FormErrorField indicates the current form error content. which field is error and error message.
type FormErrorField struct {
	ErrorField string `json:"error_field"`
	ErrorMsg   string `json:"error_msg"`
}

// GlobalValidatorMapping is a mapping from validator to translator used
var GlobalValidatorMapping = make(map[string]*MyValidator, 0)

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
func (m *MyValidator) Check(value interface{}) (errFields []*FormErrorField, err error) {
	err = m.Validate.Struct(value)
	if err != nil {
		var valErrors validator.ValidationErrors
		if !errors.As(err, &valErrors) {
			log.Error(err)
			return nil, errors.New("validate check exception")
		}

		for _, fieldError := range valErrors {
			errField := &FormErrorField{
				ErrorField: fieldError.Field(),
				ErrorMsg:   fieldError.Translate(m.Tran),
			}

			// get original tag name from value for set err field key.
			structNamespace := fieldError.StructNamespace()
			_, fieldName, found := strings.Cut(structNamespace, ".")
			if found {
				originalTag := getObjectTagByFieldName(value, fieldName)
				if len(originalTag) > 0 {
					errField.ErrorField = originalTag
				}
			}
			errFields = append(errFields, errField)
		}
		if len(errFields) > 0 {
			errMsg := ""
			if len(errFields) == 1 {
				errMsg = errFields[0].ErrorMsg
			}
			return errFields, myErrors.BadRequest(reason.RequestFormatError).WithMsg(errMsg)
		}
	}

	if v, ok := value.(Checker); ok {
		errFields, err = v.Check()
		if err != nil {
			return errFields, err
		}
	}
	return nil, nil
}

// Checker .
type Checker interface {
	Check() (errField []*FormErrorField, err error)
}

func getObjectTagByFieldName(obj interface{}, fieldName string) (tag string) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()

	objT := reflect.TypeOf(obj)
	objT = objT.Elem()

	structField, exists := objT.FieldByName(fieldName)
	if !exists {
		return ""
	}
	tag = structField.Tag.Get("json")
	if len(tag) == 0 {
		return structField.Tag.Get("form")
	}
	return tag
}
