package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/answerdev/answer/internal/base/handler"
	templaterender "github.com/answerdev/answer/internal/controller/template_render"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	"github.com/answerdev/answer/ui"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/log"
)

type TemplateController struct {
	scriptPath               string
	cssPath                  string
	templateRenderController *templaterender.TemplateRenderController
	siteInfoService          *siteinfo_common.SiteInfoCommonService
}

// NewTemplateController new controller
func NewTemplateController(
	templateRenderController *templaterender.TemplateRenderController,
	siteInfoService *siteinfo_common.SiteInfoCommonService,
) *TemplateController {
	script, css := GetStyle()
	return &TemplateController{
		scriptPath:               script,
		cssPath:                  css,
		templateRenderController: templateRenderController,
		siteInfoService:          siteInfoService,
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
func (tc *TemplateController) SiteInfo(ctx *gin.Context) *schema.TemplateSiteInfoResp {
	var err error
	resp := &schema.TemplateSiteInfoResp{}
	resp.General, err = tc.siteInfoService.GetSiteGeneral(ctx)
	if err != nil {
		log.Error(err)
	}
	resp.Interface, err = tc.siteInfoService.GetSiteInterface(ctx)
	if err != nil {
		log.Error(err)
	}

	resp.Branding, err = tc.siteInfoService.GetSiteBranding(ctx)
	if err != nil {
		log.Error(err)
	}
	resp.Year = fmt.Sprintf("%d", time.Now().Year())
	return resp
}

// Index question list
func (tc *TemplateController) Index(ctx *gin.Context) {
	req := &schema.QuestionSearch{}
	if handler.BindAndCheck(ctx, req) {
		tc.Page404(ctx)
		return
	}

	var page = req.Page

	data, count, err := tc.templateRenderController.Index(ctx, req)
	if err != nil {
		tc.Page404(ctx)
		return
	}
	siteInfo := tc.SiteInfo(ctx)
	siteInfo.Canonical = fmt.Sprintf("%s", siteInfo.General.SiteUrl)
	ctx.HTML(http.StatusOK, "question.html", gin.H{
		"siteinfo":   siteInfo,
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
		"data":       data,
		"page":       templaterender.Paginator(page, req.PageSize, count),
	})
}

func (tc *TemplateController) QuestionList(ctx *gin.Context) {
	req := &schema.QuestionSearch{}
	if handler.BindAndCheck(ctx, req) {
		tc.Page404(ctx)
		return
	}
	var page = req.Page
	data, count, err := tc.templateRenderController.Index(ctx, req)
	if err != nil {
		tc.Page404(ctx)
		return
	}
	siteInfo := tc.SiteInfo(ctx)
	siteInfo.Canonical = fmt.Sprintf("%s/questions", siteInfo.General.SiteUrl)
	ctx.HTML(http.StatusOK, "question.html", gin.H{
		"siteinfo":   siteInfo,
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
		"data":       data,
		"page":       templaterender.Paginator(page, req.PageSize, count),
	})
}

// QuestionInfo question and answers info
func (tc *TemplateController) QuestionInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	answerid := ctx.Param("answerid")
	siteInfo := tc.SiteInfo(ctx)
	encodeTitle := url.QueryEscape("title")
	siteInfo.Canonical = fmt.Sprintf("%s/questions/%s/%s", siteInfo.General.SiteUrl, id, encodeTitle)

	detail, err := tc.templateRenderController.QuestionDetail(ctx, id)
	if err != nil {
		tc.Page404(ctx)
		return
	}

	// answers
	answerReq := &schema.AnswerList{
		QuestionID:  id,
		Order:       "",
		Page:        1,
		PageSize:    999,
		LoginUserID: "",
	}
	answers, _, err := tc.templateRenderController.AnswerList(ctx, answerReq)
	if err != nil {
		tc.Page404(ctx)
		return
	}

	// comments
	objectIDs := []string{id}
	for _, answer := range answers {
		objectIDs = append(objectIDs, answer.ID)
	}
	comments, err := tc.templateRenderController.CommentList(ctx, objectIDs)
	if err != nil {
		tc.Page404(ctx)
		return
	}

	ctx.HTML(http.StatusOK, "question-detail.html", gin.H{
		"id":         id,
		"answerid":   answerid,
		"detail":     detail,
		"answers":    answers,
		"comments":   comments,
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
		"siteinfo":   siteInfo,
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
			"siteinfo":   tc.SiteInfo(ctx),
		})
		return
	}
	page := templaterender.Paginator(req.Page, req.PageSize, data.Count)
	siteInfo := tc.SiteInfo(ctx)
	siteInfo.Canonical = fmt.Sprintf("%s/tags", siteInfo.General.SiteUrl)
	ctx.HTML(http.StatusOK, "tags.html", gin.H{
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
		"page":       page,
		"data":       data,
		"siteinfo":   siteInfo,
	})
}

// TagInfo taginfo
func (tc *TemplateController) TagInfo(ctx *gin.Context) {
	tag := ctx.Param("tag")
	req := &schema.GetTamplateTagInfoReq{}
	if handler.BindAndCheck(ctx, req) {
		ctx.HTML(http.StatusOK, "404.html", gin.H{
			"scriptPath": tc.scriptPath,
			"cssPath":    tc.cssPath,
			"err":        "",
			"siteinfo":   tc.SiteInfo(ctx),
		})
		return
	}
	nowPage := req.Page
	req.Name = tag
	taginifo, questionList, questionCount, err := tc.templateRenderController.TagInfo(ctx, req)
	if err != nil {
		ctx.HTML(http.StatusOK, "404.html", gin.H{
			"scriptPath": tc.scriptPath,
			"cssPath":    tc.cssPath,
			"err":        err.Error(),
			"siteinfo":   tc.SiteInfo(ctx),
		})
		return
	}
	page := templaterender.Paginator(nowPage, req.PageSize, questionCount)
	siteInfo := tc.SiteInfo(ctx)
	siteInfo.Canonical = fmt.Sprintf("%s/tags/%s", siteInfo.General.SiteUrl, tag)
	ctx.HTML(http.StatusOK, "tag-detail.html", gin.H{
		"tag":           taginifo,
		"questionList":  questionList,
		"questionCount": questionCount,
		"scriptPath":    tc.scriptPath,
		"cssPath":       tc.cssPath,
		"siteinfo":      siteInfo,
		"page":          page,
	})
}

// UserInfo user info
func (tc *TemplateController) UserInfo(ctx *gin.Context) {
	username := ctx.Param("username")
	req := &schema.GetOtherUserInfoByUsernameReq{}
	req.Username = username
	userinfo, err := tc.templateRenderController.UserInfo(ctx, req)
	if !userinfo.Has {
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{
			"siteinfo":   tc.SiteInfo(ctx),
			"scriptPath": tc.scriptPath,
			"cssPath":    tc.cssPath,
			"err":        "",
		})
		return
	}
	if err != nil {
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{
			"siteinfo":   tc.SiteInfo(ctx),
			"scriptPath": tc.scriptPath,
			"cssPath":    tc.cssPath,
			"err":        err.Error(),
		})
		return
	}

	siteInfo := tc.SiteInfo(ctx)
	siteInfo.Canonical = fmt.Sprintf("%s/users/%s", siteInfo.General.SiteUrl, username)

	ctx.HTML(http.StatusOK, "homepage.html", gin.H{
		"siteinfo":   siteInfo,
		"userinfo":   userinfo,
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
		"bio":        template.HTML(userinfo.Info.BioHTML),
	})
}

func (tc *TemplateController) Page404(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404.html", gin.H{
		"siteinfo":   tc.SiteInfo(ctx),
		"scriptPath": tc.scriptPath,
		"cssPath":    tc.cssPath,
	})
}
