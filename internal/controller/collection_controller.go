package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/answer/internal/base/handler"
	"github.com/segmentfault/answer/internal/base/middleware"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service"
	"github.com/segmentfault/answer/pkg/converter"
	"github.com/segmentfault/pacman/errors"
)

// CollectionController collection controller
type CollectionController struct {
	collectionService *service.CollectionService
}

// NewCollectionController new controller
func NewCollectionController(collectionService *service.CollectionService) *CollectionController {
	return &CollectionController{collectionService: collectionService}
}

// CollectionSwitch add collection
// @Summary add collection
// @Description add collection
// @Tags Collection
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.CollectionSwitchReq true "collection"
// @Success 200 {object} handler.RespBody{data=schema.CollectionSwitchResp}
// @Router /answer/api/v1/collection/switch [post]
func (cc *CollectionController) CollectionSwitch(ctx *gin.Context) {
	req := &schema.CollectionSwitchReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	dto := &schema.CollectionSwitchDTO{}
	_ = copier.Copy(dto, req)

	dto.UserID = middleware.GetLoginUserIDFromContext(ctx)

	if converter.StringToInt64(req.ObjectID) < 1 {
		return
	}
	if converter.StringToInt64(dto.UserID) < 1 {
		handler.HandleResponse(ctx, errors.Unauthorized(reason.UnauthorizedError), nil)
		return
	}

	resp, err := cc.collectionService.CollectionSwitch(ctx, dto)
	handler.HandleResponse(ctx, err, resp)
}
