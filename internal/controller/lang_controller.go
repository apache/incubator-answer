package controller

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/base/handler"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/pacman/i18n"
)

type LangController struct {
	translator i18n.Translator
}

// NewLangController new language controller.
func NewLangController(tr i18n.Translator) *LangController {
	return &LangController{translator: tr}
}

// GetLangMapping get language config mapping
// @Summary get language config mapping
// @Description get language config mapping
// @Tags Lang
// @Param Accept-Language header string true "Accept-Language"
// @Produce json
// @Success 200 {object} handler.RespBody{}
// @Router /answer/api/v1/language/config [get]
func (u *LangController) GetLangMapping(ctx *gin.Context) {
	data, _ := u.translator.Dump(handler.GetLang(ctx))
	var resp map[string]any
	_ = json.Unmarshal(data, &resp)
	handler.HandleResponse(ctx, nil, resp)
}

// GetLangOptions Get language options
// @Summary Get language options
// @Description Get language options
// @Security ApiKeyAuth
// @Tags Lang
// @Produce json
// @Success 200 {object} handler.RespBody{}
// @Router /answer/api/v1/language/options [get]
// @Router /answer/admin/api/language/options [get]
func (u *LangController) GetLangOptions(ctx *gin.Context) {
	handler.HandleResponse(ctx, nil, schema.GetLangOptions)
}
