package controller

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
)

// RevisionController revision controller
type RevisionController struct {
	revisionListService *service.RevisionService
}

// NewRevisionController new controller
func NewRevisionController(revisionListService *service.RevisionService) *RevisionController {
	return &RevisionController{revisionListService: revisionListService}
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
