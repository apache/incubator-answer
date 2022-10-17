package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/base/handler"
	"github.com/segmentfault/answer/internal/base/middleware"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service"
	"github.com/segmentfault/answer/pkg/converter"
	"github.com/segmentfault/pacman/errors"
)

// SearchController tag controller
type SearchController struct {
	searchService *service.SearchService
}

// NewSearchController new controller
func NewSearchController(searchService *service.SearchService) *SearchController {
	return &SearchController{searchService: searchService}
}

// Search godoc
// @Summary search object
// @Description search object
// @Tags Search
// @Produce json
// @Security ApiKeyAuth
// @Param q query string true "query string"
// @Param order query string true "order" Enums(newest,active,score,relevance)
// @Success 200 {object} handler.RespBody{data=schema.SearchListResp}
// @Router /answer/api/v1/search [get]
func (sc *SearchController) Search(ctx *gin.Context) {
	var (
		q,
		order,
		page,
		size string
		ok  bool
		dto schema.SearchDTO
	)
	q, ok = ctx.GetQuery("q")
	if len(q) == 0 || !ok {
		handler.HandleResponse(ctx, errors.BadRequest(reason.RequestFormatError), q)
		return
	}
	page, ok = ctx.GetQuery("page")
	if !ok {
		page = "1"
	}
	size, ok = ctx.GetQuery("size")
	if !ok {
		size = "30"
	}
	order, ok = ctx.GetQuery("order")
	if !ok || (order != "newest" && order != "active" && order != "score" && order != "relevance") {
		order = "newest"
	}

	dto = schema.SearchDTO{
		Query:  q,
		Page:   converter.StringToInt(page),
		Size:   converter.StringToInt(size),
		UserID: middleware.GetLoginUserIDFromContext(ctx),
		Order:  order,
	}

	resp, total, extra, err := sc.searchService.Search(ctx, &dto)

	handler.HandleResponse(ctx, err, schema.SearchListResp{
		Total:      total,
		SearchResp: resp,
		Extra:      extra,
	})
}
