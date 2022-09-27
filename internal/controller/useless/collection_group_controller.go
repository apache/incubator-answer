package useless

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/base/handler"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service"
	"github.com/segmentfault/pacman/log"
)

// CollectionGroupController collectionGroup controller
type CollectionGroupController struct {
	log                    log.log
	collectionGroupService *service.CollectionGroupService
}

// NewCollectionGroupController new controller
func NewCollectionGroupController(collectionGroupService *service.CollectionGroupService) *CollectionGroupController {
	return &CollectionGroupController{collectionGroupService: collectionGroupService}
}

// AddCollectionGroup add collection group
// @Summary add collection group
// @Description add collection group
// @Tags CollectionGroup
// @Accept json
// @Produce json
// @Param data body schema.AddCollectionGroupReq true "collection group"
// @Success 200 {object} handler.RespBody
// Router /collection-group [post]
func (cc *CollectionGroupController) AddCollectionGroup(ctx *gin.Context) {
	req := &schema.AddCollectionGroupReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := cc.collectionGroupService.AddCollectionGroup(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// UpdateCollectionGroup update collection group
// @Summary update collection group
// @Description update collection group
// @Tags CollectionGroup
// @Accept json
// @Produce json
// @Param data body schema.UpdateCollectionGroupReq true "collection group"
// @Success 200 {object} handler.RespBody
// Router /collection-group [put]
func (cc *CollectionGroupController) UpdateCollectionGroup(ctx *gin.Context) {
	req := &schema.UpdateCollectionGroupReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := cc.collectionGroupService.UpdateCollectionGroup(ctx, req, []string{})
	handler.HandleResponse(ctx, err, nil)
}
