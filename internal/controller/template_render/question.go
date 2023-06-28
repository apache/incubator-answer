package templaterender

import (
	"html/template"
	"net/http"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/schema"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/log"
)

func (t *TemplateRenderController) Index(ctx *gin.Context, req *schema.QuestionPageReq) ([]*schema.QuestionPageResp, int64, error) {
	return t.questionService.GetQuestionPage(ctx, req)
}

func (t *TemplateRenderController) QuestionDetail(ctx *gin.Context, id string) (resp *schema.QuestionInfo, err error) {
	return t.questionService.GetQuestion(ctx, id, "", schema.QuestionPermission{})
}

func (t *TemplateRenderController) Sitemap(ctx *gin.Context) {
	general, err := t.siteInfoService.GetSiteGeneral(ctx)
	if err != nil {
		log.Error("get site general failed:", err)
		return
	}
	siteInfo, err := t.siteInfoService.GetSiteSeo(ctx)
	if err != nil {
		log.Error("get site GetSiteSeo failed:", err)
		return
	}

	questions, err := t.questionRepo.SitemapQuestions(ctx, 0, constant.SitemapMaxSize)
	if err != nil {
		log.Errorf("get sitemap questions failed: %s", err)
		return
	}

	ctx.Header("Content-Type", "application/xml")
	if len(questions) < constant.SitemapMaxSize {
		ctx.HTML(
			http.StatusOK, "sitemap.xml", gin.H{
				"xmlHeader": template.HTML(`<?xml version="1.0" encoding="UTF-8"?>`),
				"list":      questions,
				"general":   general,
				"hastitle": siteInfo.PermaLink == constant.PermaLinkQuestionIDAndTitle ||
					siteInfo.PermaLink == constant.PermaLinkQuestionIDAndTitleByShortID,
			},
		)
		return
	}

	questionNum, err := t.questionRepo.GetQuestionCount(ctx)
	if err != nil {
		log.Error("GetQuestionCount error", err)
		return
	}
	var pageList []int64
	for page := int64(1); page*constant.SitemapMaxSize < questionNum; page++ {
		pageList = append(pageList, page)
	}
	ctx.HTML(
		http.StatusOK, "sitemap-list.xml", gin.H{
			"xmlHeader": template.HTML(`<?xml version="1.0" encoding="UTF-8"?>`),
			"page":      pageList,
			"general":   general,
		},
	)
}

func (t *TemplateRenderController) SitemapPage(ctx *gin.Context, page int) error {
	general, err := t.siteInfoService.GetSiteGeneral(ctx)
	if err != nil {
		log.Error("get site general failed:", err)
		return err
	}
	siteInfo, err := t.siteInfoService.GetSiteSeo(ctx)
	if err != nil {
		log.Error("get site GetSiteSeo failed:", err)
		return err
	}

	questions, err := t.questionRepo.SitemapQuestions(ctx, page, constant.SitemapMaxSize)
	if err != nil {
		log.Errorf("get sitemap questions failed: %s", err)
		return err
	}
	ctx.Header("Content-Type", "application/xml")
	ctx.HTML(
		http.StatusOK, "sitemap.xml", gin.H{
			"xmlHeader": template.HTML(`<?xml version="1.0" encoding="UTF-8"?>`),
			"list":      questions,
			"general":   general,
			"hastitle": siteInfo.PermaLink == constant.PermaLinkQuestionIDAndTitle ||
				siteInfo.PermaLink == constant.PermaLinkQuestionIDAndTitleByShortID,
		},
	)
	return nil
}
