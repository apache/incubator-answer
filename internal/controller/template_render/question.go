package templaterender

import (
	"html/template"
	"net/http"

	"github.com/answerdev/answer/internal/schema"
	"github.com/gin-gonic/gin"
)

func (t *TemplateRenderController) Index(ctx *gin.Context, req *schema.QuestionSearch) ([]*schema.QuestionInfo, int64, error) {
	return t.questionService.SearchList(ctx, req, req.UserID)
}

func (t *TemplateRenderController) QuestionDetail(ctx *gin.Context, id string) (resp *schema.QuestionInfo, err error) {
	return t.questionService.GetQuestion(ctx, id, "", schema.QuestionPermission{})
}

func (t *TemplateRenderController) Sitemap(ctx *gin.Context) {
	if 1 == 1 {
		//question list page
		ctx.Header("Content-Type", "application/xml")
		ctx.HTML(
			http.StatusOK, "sitemap-list.xml", gin.H{
				"xmlHeader": template.HTML(`<?xml version="1.0" encoding="UTF-8"?>`),
				"list":      "string",
			},
		)
		return
	}
	//question url list
	ctx.Header("Content-Type", "application/xml")
	ctx.HTML(
		http.StatusOK, "sitemap.xml", gin.H{
			"xmlHeader": template.HTML(`<?xml version="1.0" encoding="UTF-8"?>`),
			"list":      "string",
		},
	)
}

func (t *TemplateRenderController) SitemapPage(ctx *gin.Context, page int) error {
	ctx.Header("Content-Type", "application/xml")
	ctx.HTML(
		http.StatusOK, "sitemap.xml", gin.H{
			"xmlHeader": template.HTML(`<?xml version="1.0" encoding="UTF-8"?>`),
			"list":      "string",
		},
	)
	return nil
}
