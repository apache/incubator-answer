package useless

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/segmentfault/answer/internal/base/handler"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/notification"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// NotificationReadController notificationRead controller
type NotificationReadController struct {
	log                     log.log
	notificationReadService *notification.NotificationReadService
}

// NewNotificationReadController new controller
func NewNotificationReadController(notificationReadService *notification.NotificationReadService) *NotificationReadController {
	return &NotificationReadController{notificationReadService: notificationReadService}
}

// AddNotificationRead add notification read record
// @Summary add notification read record
// @Description add notification read record
// @Tags NotificationRead
// @Accept json
// @Produce json
// @Param data body schema.AddNotificationReadReq true "notification read record"
// @Success 200 {object} handler.RespBody
// Router /notification-read [post]
func (nc *NotificationReadController) AddNotificationRead(ctx *gin.Context) {
	req := &schema.AddNotificationReadReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := nc.notificationReadService.AddNotificationRead(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// RemoveNotificationRead delete notification read record
// @Summary delete notification read record
// @Description delete notification read record
// @Tags NotificationRead
// @Accept json
// @Produce json
// @Param data body schema.RemoveNotificationReadReq true "notification read record"
// @Success 200 {object} handler.RespBody
// Router /notification-read [delete]
func (nc *NotificationReadController) RemoveNotificationRead(ctx *gin.Context) {
	req := &schema.RemoveNotificationReadReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := nc.notificationReadService.RemoveNotificationRead(ctx, req.ID)
	handler.HandleResponse(ctx, err, nil)
}

// UpdateNotificationRead update notification read record
// @Summary update notification read record
// @Description update notification read record
// @Tags NotificationRead
// @Accept json
// @Produce json
// @Param data body schema.UpdateNotificationReadReq true "notification read record"
// @Success 200 {object} handler.RespBody
// Router /notification-read [put]
func (nc *NotificationReadController) UpdateNotificationRead(ctx *gin.Context) {
	req := &schema.UpdateNotificationReadReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := nc.notificationReadService.UpdateNotificationRead(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// GetNotificationRead get notification read record one
// @Summary get notification read record one
// @Description get notification read record one
// @Tags NotificationRead
// @Accept json
// @Produce json
// @Param id path int true "notification read recordid"
// @Success 200 {object} handler.RespBody{data=schema.GetNotificationReadResp}
// Router /notification-read/{id} [get]
func (nc *NotificationReadController) GetNotificationRead(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if id == 0 {
		handler.HandleResponse(ctx, errors.BadRequest(reason.RequestFormatError), nil)
		return
	}

	resp, err := nc.notificationReadService.GetNotificationRead(ctx, id)
	handler.HandleResponse(ctx, err, resp)
}

// GetNotificationReadList get notification read record list
// @Summary get notification read record list
// @Description get notification read record list
// @Tags NotificationRead
// @Produce json
// @Param user_id query string false "user id"
// @Param message_id query string false "message id"
// @Param is_read query string false "read status(unread: 1; read 2)"
// @Success 200 {object} handler.RespBody{data=[]schema.GetNotificationReadResp}
// Router /notification-reads [get]
func (nc *NotificationReadController) GetNotificationReadList(ctx *gin.Context) {
	req := &schema.GetNotificationReadListReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := nc.notificationReadService.GetNotificationReadList(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}

// GetNotificationReadWithPage get notification read record page
// @Summary get notification read record page
// @Description get notification read record page
// @Tags NotificationRead
// @Produce json
// @Param page query int false "page size"
// @Param page_size query int false "page size"
// @Param user_id query string false "user id"
// @Param message_id query string false "message id"
// @Param is_read query string false "read status(unread: 1; read 2)"
// @Success 200 {object} handler.RespBody{data=pager.PageModel{list=[]schema.GetNotificationReadResp}}
// Router /notification-reads/page [get]
func (nc *NotificationReadController) GetNotificationReadWithPage(ctx *gin.Context) {
	req := &schema.GetNotificationReadWithPageReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := nc.notificationReadService.GetNotificationReadWithPage(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}
