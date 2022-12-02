package controller

import (
	"fmt"

	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service"
	"github.com/answerdev/answer/internal/service/dashboard"
	"github.com/answerdev/answer/internal/service/permission"
	"github.com/answerdev/answer/internal/service/rank"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
)

// AnswerController answer controller
type AnswerController struct {
	answerService    *service.AnswerService
	rankService      *rank.RankService
	dashboardService *dashboard.DashboardService
}

// NewAnswerController new controller
func NewAnswerController(answerService *service.AnswerService,
	rankService *rank.RankService,
	dashboardService *dashboard.DashboardService,
) *AnswerController {
	return &AnswerController{
		answerService:    answerService,
		rankService:      rankService,
		dashboardService: dashboardService,
	}
}

// RemoveAnswer delete answer
// @Summary delete answer
// @Description delete answer
// @Tags api-answer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.RemoveAnswerReq true "answer"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/answer [delete]
func (ac *AnswerController) RemoveAnswer(ctx *gin.Context) {
	req := &schema.RemoveAnswerReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	req.IsAdmin = middleware.GetIsAdminFromContext(ctx)
	can, err := ac.rankService.CheckOperationPermission(ctx, req.UserID, permission.AnswerDelete, req.ID)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	if !can {
		handler.HandleResponse(ctx, errors.Forbidden(reason.RankFailToMeetTheCondition), nil)
		return
	}

	err = ac.answerService.RemoveAnswer(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// Get godoc
// @Summary Get Answer
// @Description Get Answer
// @Tags api-answer
// @Accept  json
// @Produce  json
// @Param id query string true "Answer TagID"  default(1)
// @Router  /answer/api/v1/answer/info [get]
// @Success 200 {string} string ""
func (ac *AnswerController) Get(ctx *gin.Context) {
	id := ctx.Query("id")
	userID := middleware.GetLoginUserIDFromContext(ctx)

	info, questionInfo, has, err := ac.answerService.Get(ctx, id, userID)
	if err != nil {
		handler.HandleResponse(ctx, err, gin.H{})
		return
	}
	if !has {
		handler.HandleResponse(ctx, fmt.Errorf(""), gin.H{})
		return
	}
	handler.HandleResponse(ctx, err, gin.H{
		"info":     info,
		"question": questionInfo,
	})
}

// Add godoc
// @Summary Insert Answer
// @Description Insert Answer
// @Tags api-answer
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data body schema.AnswerAddReq  true "AnswerAddReq"
// @Success 200 {string} string ""
// @Router /answer/api/v1/answer [post]
func (ac *AnswerController) Add(ctx *gin.Context) {
	req := &schema.AnswerAddReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)

	can, err := ac.rankService.CheckOperationPermission(ctx, req.UserID, permission.AnswerAdd, "")
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	if !can {
		handler.HandleResponse(ctx, errors.Forbidden(reason.RankFailToMeetTheCondition), nil)
		return
	}

	answerID, err := ac.answerService.Insert(ctx, req)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	info, questionInfo, has, err := ac.answerService.Get(ctx, answerID, req.UserID)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	if !has {
		// todo !has
		handler.HandleResponse(ctx, nil, nil)
		return
	}
	handler.HandleResponse(ctx, nil, gin.H{
		"info":     info,
		"question": questionInfo,
	})
}

// Update godoc
// @Summary Update Answer
// @Description Update Answer
// @Tags api-answer
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data body schema.AnswerUpdateReq  true "AnswerUpdateReq"
// @Success 200 {string} string ""
// @Router /answer/api/v1/answer [put]
func (ac *AnswerController) Update(ctx *gin.Context) {
	req := &schema.AnswerUpdateReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)

	canList, err := ac.rankService.CheckOperationPermissions(ctx, req.UserID, []string{
		permission.AnswerEdit,
		permission.AnswerEditWithoutReview,
	}, "")
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	req.CanEdit = canList[0]
	req.NoNeedReview = canList[1]
	if !req.CanEdit {
		handler.HandleResponse(ctx, errors.Forbidden(reason.RankFailToMeetTheCondition), nil)
		return
	}

	_, err = ac.answerService.Update(ctx, req)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	_, _, _, err = ac.answerService.Get(ctx, req.ID, req.UserID)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, &schema.AnswerUpdateResp{WaitForReview: !req.NoNeedReview})
}

// AnswerList godoc
// @Summary AnswerList
// @Description AnswerList <br> <b>order</b> (default or updated)
// @Tags api-answer
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body schema.AnswerListReq  true "AnswerListReq"
// @Success 200 {string} string ""
// @Router /answer/api/v1/answer/list [get]
func (ac *AnswerController) AnswerList(ctx *gin.Context) {
	req := &schema.AnswerListReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)

	canList, err := ac.rankService.CheckOperationPermissions(ctx, req.UserID, []string{
		permission.AnswerEdit,
		permission.AnswerDelete,
	}, "")
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	req.CanEdit = canList[0]
	req.CanDelete = canList[1]

	list, count, err := ac.answerService.SearchList(ctx, req)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, gin.H{
		"list":  list,
		"count": count,
	})
}

// Adopted godoc
// @Summary Adopted
// @Description Adopted
// @Tags api-answer
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param data body schema.AnswerAdoptedReq  true "AnswerAdoptedReq"
// @Success 200 {string} string ""
// @Router /answer/api/v1/answer/acceptance [post]
func (ac *AnswerController) Adopted(ctx *gin.Context) {
	req := &schema.AnswerAdoptedReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	can, err := ac.rankService.CheckOperationPermission(ctx, req.UserID, permission.AnswerAccept, req.QuestionID)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	if !can {
		handler.HandleResponse(ctx, errors.Forbidden(reason.RankFailToMeetTheCondition), nil)
		return
	}

	err = ac.answerService.UpdateAdopted(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// AdminSetAnswerStatus godoc
// @Summary AdminSetAnswerStatus
// @Description Status:[available,deleted]
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body entity.AdminSetAnswerStatusRequest true "AdminSetAnswerStatusRequest"
// @Router /answer/admin/api/answer/status [put]
// @Success 200 {object} handler.RespBody
func (ac *AnswerController) AdminSetAnswerStatus(ctx *gin.Context) {
	req := &schema.AdminSetAnswerStatusRequest{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.UserID = middleware.GetLoginUserIDFromContext(ctx)

	err := ac.answerService.AdminSetAnswerStatus(ctx, req)
	handler.HandleResponse(ctx, err, gin.H{})
}
