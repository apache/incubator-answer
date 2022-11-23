package controller

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service"
	"github.com/answerdev/answer/internal/service/rank"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
)

// RevisionController revision controller
type RevisionController struct {
	revisionListService *service.RevisionService
	rankService         *rank.RankService
}

// NewRevisionController new controller
func NewRevisionController(
	revisionListService *service.RevisionService,
	rankService *rank.RankService,
) *RevisionController {
	return &RevisionController{
		revisionListService: revisionListService,
		rankService:         rankService,
	}
}

// GetRevisionList godoc
// @Summary get revision list
// @Description get revision list
// @Tags Revision
// @Produce json
// @Param object_id query string true "object id"
// @Success 200 {object} handler.RespBody{data=[]schema.GetRevisionResp}
// @Router /answer/api/v1/revisions [get]
func (rc *RevisionController) GetRevisionList(ctx *gin.Context) {
	objectID := ctx.Query("object_id")
	if objectID == "0" || objectID == "" {
		handler.HandleResponse(ctx, errors.BadRequest(reason.RequestFormatError), nil)
		return
	}

	req := &schema.GetRevisionListReq{
		ObjectID: objectID,
	}

	resp, err := rc.revisionListService.GetRevisionList(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}

// GetUnreviewedRevisionList godoc
// @Summary get unreviewed revision list
// @Description get unreviewed revision list
// @Tags Revision
// @Produce json
// @Param page query string true "page id"
// @Success 200 {object} handler.RespBody{data=[]schema.GetRevisionResp}
// @Router /answer/api/v1/revisions/unreviewed [get]
func (rc *RevisionController) GetUnreviewedRevisionList(ctx *gin.Context) {
	pageStr := ctx.Query("page")
	page := converter.StringToInt(pageStr)
	req := &schema.RevisionSearch{
		Page: page,
	}
	userinfo := middleware.GetUserInfoFromContext(ctx)
	if !userinfo.IsAdmin {
		userID := middleware.GetLoginUserIDFromContext(ctx)
		if can, err := rc.rankService.CheckRankPermission(ctx, userID, rank.QuestionEditRank); err != nil || !can {
			handler.HandleResponse(ctx, err, errors.Forbidden(reason.RankFailToMeetTheCondition))
			return
		}
	}

	resp, count, err := rc.revisionListService.GetUnreviewedRevisionList(ctx, req)
	handler.HandleResponse(ctx, err, gin.H{
		"list":  resp,
		"count": count,
	})
}
