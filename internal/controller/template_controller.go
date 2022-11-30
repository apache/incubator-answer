package controller

import (
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
	ctx.HTML(http.StatusOK, "question.html", gin.H{
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
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
	ctx.HTML(http.StatusOK, "tags.html", gin.H{
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
	})
}

// TagInfo taginfo
func (tc *TemplateController) TagInfo(ctx *gin.Context) {
	tag := ctx.Param("tag")
	ctx.HTML(http.StatusOK, "tag-detail.html", gin.H{
		"tag":        tag,
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
	})
}

// UserInfo user info
func (tc *TemplateController) UserInfo(ctx *gin.Context) {
	req := &schema.GetOtherUserInfoByUsernameReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	tc.templateRenderController.UserInfo(ctx, req)

	username := ctx.Param("username")
	ctx.HTML(http.StatusOK, "homepage.html", gin.H{
		"username":   username,
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
	})
}

func (tc *TemplateController) Page404(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "404.html", gin.H{
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
	})
}
