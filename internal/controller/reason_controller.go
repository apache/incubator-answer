package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/base/handler"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/reason"
)

// ReasonController answer controller
type ReasonController struct {
	reasonService *reason.ReasonService
}

// NewReasonController new controller
func NewReasonController(answerService *reason.ReasonService) *ReasonController {
	return &ReasonController{reasonService: answerService}
}

// Reasons godoc
// @Summary get reasons by object type and action
// @Description get reasons by object type and action
// @Tags reason
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param object_type query string true "object_type" Enums(question, answer, comment, user)
// @Param action query string true "action" Enums(status, close, flag, review)
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/reasons [get]
// @Router /answer/admin/api/reasons [get]
func (rc *ReasonController) Reasons(ctx *gin.Context) {
	req := &schema.ReasonReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	reasons, err := rc.reasonService.GetReasons(ctx, *req)
	if err != nil {
		err = nil
	}
	handler.HandleResponse(ctx, err, reasons)
}
