package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/answerdev/answer/internal/base/handler"
	templaterender "github.com/answerdev/answer/internal/controller/template_render"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	"github.com/answerdev/answer/pkg/htmltext"
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
	req := &schema.QuestionSearch{
		Order: "newest",
	}
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
	tc.html(ctx, http.StatusOK, "question.html", siteInfo, gin.H{
		"data": data,
		"page": templaterender.Paginator(page, req.PageSize, count),
	})
}

func (tc *TemplateController) QuestionList(ctx *gin.Context) {
	req := &schema.QuestionSearch{
		Order: "newest",
	}
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

	tc.html(ctx, http.StatusOK, "question.html", siteInfo, gin.H{
		"data": data,
		"page": templaterender.Paginator(page, req.PageSize, count),
	})
}

// QuestionInfo question and answers info
func (tc *TemplateController) QuestionInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	answerid := ctx.Param("answerid")

	detail, err := tc.templateRenderController.QuestionDetail(ctx, id)
	if err != nil {
		tc.Page404(ctx)
		return
	}

	// answers
	answerReq := &schema.AnswerListReq{
		QuestionID: id,
		Order:      "",
		Page:       1,
		PageSize:   999,
		UserID:     "",
	}
	answers, answerCount, err := tc.templateRenderController.AnswerList(ctx, answerReq)
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
	siteInfo := tc.SiteInfo(ctx)
	encodeTitle := htmltext.UrlTitle(detail.Title)
	siteInfo.Canonical = fmt.Sprintf("%s/questions/%s/%s", siteInfo.General.SiteUrl, id, encodeTitle)
	if siteInfo.General.PermaLink == schema.PermaLinkQuestionID {
		siteInfo.Canonical = fmt.Sprintf("%s/questions/%s", siteInfo.General.SiteUrl, id)
	}
	jsonLD := &schema.QAPageJsonLD{}
	jsonLD.Context = "https://schema.org"
	jsonLD.Type = "QAPage"
	jsonLD.MainEntity.Type = "Question"
	jsonLD.MainEntity.Name = detail.Title
	jsonLD.MainEntity.Text = htmltext.ClearText(detail.HTML)
	jsonLD.MainEntity.AnswerCount = int(answerCount)
	jsonLD.MainEntity.UpvoteCount = detail.VoteCount
	jsonLD.MainEntity.DateCreated = time.Unix(detail.CreateTime, 0)
	jsonLD.MainEntity.Author.Type = "Person"
	jsonLD.MainEntity.Author.Name = detail.UserInfo.DisplayName
	answerList := make([]*schema.SuggestedAnswerItem, 0)
	for _, answer := range answers {
		item := &schema.SuggestedAnswerItem{}
		item.Type = "Answer"
		item.Text = htmltext.ClearText(answer.HTML)
		item.DateCreated = time.Unix(answer.CreateTime, 0)
		item.UpvoteCount = answer.VoteCount
		item.URL = fmt.Sprintf("%s/%s", siteInfo.Canonical, answer.ID)
		item.Author.Type = "Person"
		item.Author.Name = answer.UserInfo.DisplayName
		answerList = append(answerList, item)
	}
	jsonLD.MainEntity.SuggestedAnswer = answerList
	jsonLDStr, err := json.Marshal(jsonLD)
	if err == nil {
		siteInfo.JsonLD = `<script data-react-helmet="true" type="application/ld+json">` + string(jsonLDStr) + ` </script>`
	}

	siteInfo.Description = htmltext.FetchExcerpt(detail.HTML, "...", 240)
	tags := make([]string, 0)
	for _, tag := range detail.Tags {
		tags = append(tags, tag.DisplayName)
	}
	siteInfo.Keywords = strings.Replace(strings.Trim(fmt.Sprint(tags), "[]"), " ", ",", -1)

	tc.html(ctx, http.StatusOK, "question-detail.html", siteInfo, gin.H{
		"id":       id,
		"answerid": answerid,
		"detail":   detail,
		"answers":  answers,
		"comments": comments,
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
		tc.Page404(ctx)
		return
	}
	page := templaterender.Paginator(req.Page, req.PageSize, data.Count)

	siteInfo := tc.SiteInfo(ctx)
	siteInfo.Canonical = fmt.Sprintf("%s/tags", siteInfo.General.SiteUrl)
	tc.html(ctx, http.StatusOK, "tags.html", siteInfo, gin.H{
		"page": page,
		"data": data,
	})
}

// TagInfo taginfo
func (tc *TemplateController) TagInfo(ctx *gin.Context) {
	tag := ctx.Param("tag")
	req := &schema.GetTamplateTagInfoReq{}
	if handler.BindAndCheck(ctx, req) {
		tc.Page404(ctx)
		return
	}
	nowPage := req.Page
	req.Name = tag
	taginifo, questionList, questionCount, err := tc.templateRenderController.TagInfo(ctx, req)
	if err != nil {
		tc.Page404(ctx)
		return
	}
	page := templaterender.Paginator(nowPage, req.PageSize, questionCount)

	siteInfo := tc.SiteInfo(ctx)
	siteInfo.Canonical = fmt.Sprintf("%s/tags/%s", siteInfo.General.SiteUrl, tag)

	siteInfo.Description = htmltext.FetchExcerpt(taginifo.ParsedText, "...", 240)
	siteInfo.Keywords = taginifo.DisplayName

	tc.html(ctx, http.StatusOK, "tag-detail.html", siteInfo, gin.H{
		"tag":           taginifo,
		"questionList":  questionList,
		"questionCount": questionCount,
		"page":          page,
	})
}

// UserInfo user info
func (tc *TemplateController) UserInfo(ctx *gin.Context) {
	// urlPath := ctx.Request.URL.Path
	// filePath := ""
	// switch urlPath {
	// case "/users/login":
	// 	filePath = "build/index.html"
	// case "/users/register":
	// 	filePath = "build/index.html"
	// default:
	username := ctx.Param("username")
	req := &schema.GetOtherUserInfoByUsernameReq{}
	req.Username = username
	userinfo, err := tc.templateRenderController.UserInfo(ctx, req)
	if !userinfo.Has {
		tc.Page404(ctx)
		return
	}
	if err != nil {
		tc.Page404(ctx)
		return
	}

	siteInfo := tc.SiteInfo(ctx)
	siteInfo.Canonical = fmt.Sprintf("%s/users/%s", siteInfo.General.SiteUrl, username)
	tc.html(ctx, http.StatusOK, "homepage.html", siteInfo, gin.H{
		"userinfo": userinfo,
		"bio":      template.HTML(userinfo.Info.BioHTML),
	})
	// }

	// file, err := ui.Build.ReadFile(filePath)
	// if err != nil {
	// 	log.Error(err)
	// 	ctx.Status(http.StatusNotFound)
	// 	return
	// }
	// ctx.Header("content-type", "text/html;charset=utf-8")
	// ctx.String(http.StatusOK, string(file))

}

func (tc *TemplateController) Page404(ctx *gin.Context) {
	tc.html(ctx, http.StatusNotFound, "404.html", tc.SiteInfo(ctx), gin.H{})
}

func (tc *TemplateController) html(ctx *gin.Context, code int, tpl string, siteInfo *schema.TemplateSiteInfoResp, data gin.H) {
	data["siteinfo"] = siteInfo
	data["scriptPath"] = "" //tc.scriptPath
	data["cssPath"] = tc.cssPath
	data["keywords"] = siteInfo.Keywords
	if siteInfo.Description == "" {
		siteInfo.Description = siteInfo.General.Description
	}
	data["description"] = siteInfo.Description
	data["language"] = handler.GetLang(ctx)
	data["timezone"] = siteInfo.Interface.TimeZone

	ctx.HTML(code, tpl, data)
}
