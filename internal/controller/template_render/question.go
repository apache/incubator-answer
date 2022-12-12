package templaterender

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/answerdev/answer/internal/schema"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/log"
)

func (t *TemplateRenderController) Index(ctx *gin.Context, req *schema.QuestionSearch) ([]*schema.QuestionInfo, int64, error) {
	return t.questionService.SearchList(ctx, req, req.UserID)
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

	sitemapInfo := &schema.SiteMapList{}
	infoStr, err := t.data.Cache.GetString(ctx, schema.SitemapCachekey)
	if err != nil {
		log.Errorf("get Cache failed: %s", err)
		return
	}
	if err = json.Unmarshal([]byte(infoStr), sitemapInfo); err != nil {
		log.Errorf("get sitemap info failed: %s", err)
		return
	}

	if len(sitemapInfo.QuestionIDs) > 0 {
		//question url list
		ctx.Header("Content-Type", "application/xml")
		ctx.HTML(
			http.StatusOK, "sitemap.xml", gin.H{
				"xmlHeader": template.HTML(`<?xml version="1.0" encoding="UTF-8"?>`),
				"list":      sitemapInfo.QuestionIDs,
				"general":   general,
			},
		)
	} else {
		//question list page
		ctx.Header("Content-Type", "application/xml")
		ctx.HTML(
			http.StatusOK, "sitemap-list.xml", gin.H{
				"xmlHeader": template.HTML(`<?xml version="1.0" encoding="UTF-8"?>`),
				"page":      sitemapInfo.MaxPageNum,
				"general":   general,
			},
		)
		return
	}
}

func (t *TemplateRenderController) SitemapPage(ctx *gin.Context, page int) error {
	sitemapInfo := &schema.SiteMapPageList{}
	general, err := t.siteInfoService.GetSiteGeneral(ctx)
	if err != nil {
		log.Error("get site general failed:", err)
		return err
	}

	cachekey := fmt.Sprintf(schema.SitemapPageCachekey, page)
	infoStr, err := t.data.Cache.GetString(ctx, cachekey)
	if err != nil {
		log.Errorf("get Cache failed: %s", err)
		return err
	}
	if err = json.Unmarshal([]byte(infoStr), sitemapInfo); err != nil {
		log.Errorf("get sitemap info failed: %s", err)
		return err
	}
	ctx.Header("Content-Type", "application/xml")
	ctx.HTML(
		http.StatusOK, "sitemap.xml", gin.H{
			"xmlHeader": template.HTML(`<?xml version="1.0" encoding="UTF-8"?>`),
			"list":      sitemapInfo.PageData,
			"general":   general,
		},
	)
	return nil
}
