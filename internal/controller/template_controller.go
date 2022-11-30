package controller

import (
	"html/template"
	"net/http"
	"regexp"

	"github.com/answerdev/answer/internal/base/handler"
	templaterender "github.com/answerdev/answer/internal/controller/template_render"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/ui"
	"github.com/gin-gonic/gin"
)

type TemplateController struct {
	scriptPath               string
	cssPath                  string
	templateRenderController *templaterender.TemplateRenderController
}

// NewTemplateController new controller
func NewTemplateController(
	templateRenderController *templaterender.TemplateRenderController,
) *TemplateController {
	script, css := GetStyle()
	return &TemplateController{
		scriptPath:               script,
		cssPath:                  css,
		templateRenderController: templateRenderController,
	}
}
func GetStyle() (script, css string) {
	file, err := ui.Build.ReadFile("build/index.html")
	if err != nil {
		return
	}
	scriptRegexp := regexp.MustCompile(`<script defer="defer" src="(.*)"></script>`)
	scriptData := scriptRegexp.FindStringSubmatch(string(file))
	cssRegexp := regexp.MustCompile(`<link href="(.*)" rel="stylesheet">`)
	cssListData := cssRegexp.FindStringSubmatch(string(file))
	if len(scriptData) == 2 {
		script = scriptData[1]
	}
	if len(cssListData) == 2 {
		css = cssListData[1]
	}
	return
}

// Index question list
func (tc *TemplateController) Index(ctx *gin.Context) {
	req := &schema.QuestionSearch{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	data, count, err := tc.templateRenderController.Index(ctx, req)
	if err != nil {
		tc.Page404(ctx)
		return
	}

	ctx.HTML(http.StatusOK, "question.html", gin.H{
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
		"data":       data,
		"page":       templaterender.Paginator(req.Page, req.PageSize, count),
	})
}

// QuestionInfo question and answers info
func (tc *TemplateController) QuestionInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	answerid := ctx.Param("answerid")
	ctx.HTML(http.StatusOK, "question-detail.html", gin.H{
		"id":         id,
		"answerid":   answerid,
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
	})
}

// TagList tags list
func (tc *TemplateController) TagList(ctx *gin.Context) {
	req := &schema.GetTagWithPageReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	data, err := tc.templateRenderController.TagList(ctx, req)
	if err != nil {
		ctx.HTML(http.StatusOK, "404.html", gin.H{
			"scriptPath": tc.scriptPath,
			"cssPath":    tc.cssPath,
			"err":        err.Error(),
		})
		return
	}
	page := templaterender.Paginator(req.Page, req.PageSize, data.Count)
	ctx.HTML(http.StatusOK, "tags.html", gin.H{
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
		"page":       page,
		"data":       data,
	})
}

// TagInfo taginfo
func (tc *TemplateController) TagInfo(ctx *gin.Context) {
	tag := ctx.Param("tag")

	req := &schema.GetTagInfoReq{}
	req.Name = tag
	taginifo, err := tc.templateRenderController.TagInfo(ctx, req)
	if err != nil {
		ctx.HTML(http.StatusOK, "404.html", gin.H{
			"scriptPath": tc.scriptPath,
			"cssPath":    tc.cssPath,
			"err":        err.Error(),
		})
		return
	}
	ctx.HTML(http.StatusOK, "tag-detail.html", gin.H{
		"tag":        taginifo,
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
	})
}

// UserInfo user info
func (tc *TemplateController) UserInfo(ctx *gin.Context) {
	username := ctx.Param("username")
	req := &schema.GetOtherUserInfoByUsernameReq{}
	req.Username = username
	userinfo, err := tc.templateRenderController.UserInfo(ctx, req)
	if err != nil {
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{
			"scriptPath": tc.scriptPath,
			"cssPath":    tc.cssPath,
			"err":        err.Error(),
		})
		return
	}
	ctx.HTML(http.StatusOK, "homepage.html", gin.H{
		"userinfo":   userinfo,
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
		"bio":        template.HTML(userinfo.Info.BioHTML),
	})
}

func (tc *TemplateController) Page404(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", gin.H{
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
	})
}
