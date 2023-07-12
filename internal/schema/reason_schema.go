package schema

import (
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/segmentfault/pacman/i18n"
)

type ReasonItem struct {
	ReasonType  int    `json:"reason_type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ContentType string `json:"content_type"`
	Placeholder string `json:"placeholder"`
}

type ReasonReq struct {
	// ObjectType
	ObjectType string `validate:"required" form:"object_type" json:"object_type"`
	// Action
	Action string `validate:"required" form:"action" json:"action"`
}

func (r *ReasonItem) Translate(keyPrefix string, lang i18n.Language) {
	trField := func(fieldName, fieldData string) string {
		// If fieldData is empty, means no need to translate
		if len(fieldData) == 0 {
			return fieldData
		}
		key := keyPrefix + "." + fieldName
		fieldTr := translator.Tr(lang, key)
		if fieldTr != key {
			// If i18n key exists, return i18n value
			return fieldTr
		}
		// If i18n key not exists, return fieldData original value
		return fieldData + "没翻译"
	}

	r.Name = trField("name", r.Name)
	r.Description = trField("desc", r.Description)
	r.Placeholder = trField("placeholder", r.Placeholder)
}
