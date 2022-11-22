package controller

import (
	"fmt"

	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service"
	"github.com/answerdev/answer/internal/service/dashboard"
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
	if can, err := ac.rankService.CheckRankPermission(ctx, req.UserID, rank.AnswerDeleteRank); err != nil || !can {
		handler.HandleResponse(ctx, err, errors.Forbidden(reason.RankFailToMeetTheCondition))
		return
	}
	userinfo := middleware.GetUserInfoFromContext(ctx)
	req.IsAdmin = userinfo.IsAdmin
	err := ac.answerService.RemoveAnswer(ctx, req)
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

	if can, err := ac.rankService.CheckRankPermission(ctx, req.UserID, rank.AnswerAddRank); err != nil || !can {
		handler.HandleResponse(ctx, err, errors.Forbidden(reason.RankFailToMeetTheCondition))
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
	userinfo := middleware.GetUserInfoFromContext(ctx)
	req.IsAdmin = userinfo.IsAdmin

	if can, err := ac.rankService.CheckRankPermission(ctx, req.UserID, rank.AnswerEditRank); err != nil || !can {
		handler.HandleResponse(ctx, err, errors.Forbidden(reason.RankFailToMeetTheCondition))
		return
	}

	_, err := ac.answerService.Update(ctx, req)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	info, questionInfo, has, err := ac.answerService.Get(ctx, req.ID, req.UserID)
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

// AnswerList godoc
// @Summary AnswerList
// @Description AnswerList <br> <b>order</b> (default or updated)
// @Tags api-answer
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param data body schema.AnswerList  true "AnswerList"
// @Success 200 {string} string ""
// @Router /answer/api/v1/answer/list [get]
func (ac *AnswerController) AnswerList(ctx *gin.Context) {
	req := &schema.AnswerList{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.LoginUserID = middleware.GetLoginUserIDFromContext(ctx)
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
	if can, err := ac.rankService.CheckRankPermission(ctx, req.UserID, rank.AnswerAcceptRank); err != nil || !can {
		handler.HandleResponse(ctx, err, errors.Forbidden(reason.RankFailToMeetTheCondition))
		return
	}

	err := ac.answerService.UpdateAdopted(ctx, req)
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
	req := &entity.AdminSetAnswerStatusRequest{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	err := ac.answerService.AdminSetAnswerStatus(ctx, req.AnswerID, req.StatusStr)
	handler.HandleResponse(ctx, err, gin.H{})
}
