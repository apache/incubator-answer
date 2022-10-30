package controller

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/notification"
	"github.com/gin-gonic/gin"
)

// NotificationController notification controller
type NotificationController struct {
	notificationService *notification.NotificationService
}

// NewNotificationController new controller
func NewNotificationController(notificationService *notification.NotificationService) *NotificationController {
	return &NotificationController{notificationService: notificationService}
}

// GetRedDot
// @Summary     GetRedDot
// @Description GetRedDot
// @Tags        Notification
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @Success     200 {object} handler.RespBody
// @Router      /answer/api/v1/notification/status [get]
func (nc *NotificationController) GetRedDot(ctx *gin.Context) {
	userID := middleware.GetLoginUserIDFromContext(ctx)
	RedDot, err := nc.notificationService.GetRedDot(ctx, userID)
	handler.HandleResponse(ctx, err, RedDot)
}

// ClearRedDot
// @Summary     DelRedDot
// @Description DelRedDot
// @Tags        Notification
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @Param       data body     schema.NotificationClearRequest true "NotificationClearRequest"
// @Success     200  {object} handler.RespBody
// @Router      /answer/api/v1/notification/status [put]
func (nc *NotificationController) ClearRedDot(ctx *gin.Context) {
	req := &schema.NotificationClearRequest{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	userID := middleware.GetLoginUserIDFromContext(ctx)
	RedDot, err := nc.notificationService.ClearRedDot(ctx, userID, req.TypeStr)
	handler.HandleResponse(ctx, err, RedDot)
}

// ClearUnRead
// @Summary     ClearUnRead
// @Description ClearUnRead
// @Tags        Notification
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @Param       data body     schema.NotificationClearRequest true "NotificationClearRequest"
// @Success     200  {object} handler.RespBody
// @Router      /answer/api/v1/notification/read/state/all [put]
func (nc *NotificationController) ClearUnRead(ctx *gin.Context) {
	req := &schema.NotificationClearRequest{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	userID := middleware.GetLoginUserIDFromContext(ctx)
	err := nc.notificationService.ClearUnRead(ctx, userID, req.TypeStr)
	handler.HandleResponse(ctx, err, gin.H{})
}

// ClearIDUnRead
// @Summary     ClearUnRead
// @Description ClearUnRead
// @Tags        Notification
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @Param       data body     schema.NotificationClearIDRequest true "NotificationClearIDRequest"
// @Success     200  {object} handler.RespBody
// @Router      /answer/api/v1/notification/read/state [put]
func (nc *NotificationController) ClearIDUnRead(ctx *gin.Context) {
	req := &schema.NotificationClearIDRequest{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	userID := middleware.GetLoginUserIDFromContext(ctx)
	err := nc.notificationService.ClearIDUnRead(ctx, userID, req.ID)
	handler.HandleResponse(ctx, err, gin.H{})
}

// GetList get notification list
// @Summary     get notification list
// @Description get notification list
// @Tags        Notification
// @Accept      json
// @Produce     json
// @Security    ApiKeyAuth
// @Param       page      query    int    false "page size"
// @Param       page_size query    int    false "page size"
// @Param       type      query    string true  "type" Enums(inbox,achievement)
// @Success     200       {object} handler.RespBody
// @Router      /answer/api/v1/notification/page [get]
func (nc *NotificationController) GetList(ctx *gin.Context) {
	req := &schema.NotificationSearch{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	resp, err := nc.notificationService.GetList(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}
