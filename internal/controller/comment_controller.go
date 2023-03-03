package controller

import (
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/comment"
	"github.com/answerdev/answer/internal/service/permission"
	"github.com/answerdev/answer/internal/service/rank"
	"github.com/answerdev/answer/pkg/uid"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
)

// CommentController comment controller
type CommentController struct {
	commentService *comment.CommentService
	rankService    *rank.RankService
}

// NewCommentController new controller
func NewCommentController(
	commentService *comment.CommentService,
	rankService *rank.RankService) *CommentController {
	return &CommentController{commentService: commentService, rankService: rankService}
}

// AddComment add comment
// @Summary add comment
// @Description add comment
// @Tags Comment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.AddCommentReq true "comment"
// @Success 200 {object} handler.RespBody{data=schema.GetCommentResp}
// @Router /answer/api/v1/comment [post]
func (cc *CommentController) AddComment(ctx *gin.Context) {
	req := &schema.AddCommentReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.ObjectID = uid.DeShortID(req.ObjectID)
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	canList, err := cc.rankService.CheckOperationPermissions(ctx, req.UserID, []string{
		permission.CommentAdd,
		permission.CommentEdit,
		permission.CommentDelete,
	})
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	req.CanAdd = canList[0]
	req.CanEdit = canList[1]
	req.CanDelete = canList[2]
	if !req.CanAdd {
		handler.HandleResponse(ctx, errors.Forbidden(reason.RankFailToMeetTheCondition), nil)
		return
	}

	resp, err := cc.commentService.AddComment(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}

// RemoveComment remove comment
// @Summary remove comment
// @Description remove comment
// @Tags Comment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.RemoveCommentReq true "comment"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/comment [delete]
func (cc *CommentController) RemoveComment(ctx *gin.Context) {
	req := &schema.RemoveCommentReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	can, err := cc.rankService.CheckOperationPermission(ctx, req.UserID, permission.CommentDelete, req.CommentID)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	if !can {
		handler.HandleResponse(ctx, errors.Forbidden(reason.RankFailToMeetTheCondition), nil)
		return
	}

	err = cc.commentService.RemoveComment(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// UpdateComment update comment
// @Summary update comment
// @Description update comment
// @Tags Comment
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.UpdateCommentReq true "comment"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/comment [put]
func (cc *CommentController) UpdateComment(ctx *gin.Context) {
	req := &schema.UpdateCommentReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	req.IsAdmin = middleware.GetIsAdminFromContext(ctx)
	can, err := cc.rankService.CheckOperationPermission(ctx, req.UserID, permission.CommentEdit, req.CommentID)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	if !can {
		handler.HandleResponse(ctx, errors.Forbidden(reason.RankFailToMeetTheCondition), nil)
		return
	}

	err = cc.commentService.UpdateComment(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// GetCommentWithPage get comment page
// @Summary get comment page
// @Description get comment page
// @Tags Comment
// @Produce json
// @Param page query int false "page"
// @Param page_size query int false "page size"
// @Param object_id query string true "object id"
// @Param query_cond query string false "query condition" Enums(vote)
// @Success 200 {object} handler.RespBody{data=pager.PageModel{list=[]schema.GetCommentResp}}
// @Router /answer/api/v1/comment/page [get]
func (cc *CommentController) GetCommentWithPage(ctx *gin.Context) {
	req := &schema.GetCommentWithPageReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.ObjectID = uid.DeShortID(req.ObjectID)
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	canList, err := cc.rankService.CheckOperationPermissions(ctx, req.UserID, []string{
		permission.CommentEdit,
		permission.CommentDelete,
	})
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	req.CanEdit = canList[0]
	req.CanDelete = canList[1]

	resp, err := cc.commentService.GetCommentWithPage(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}

// GetCommentPersonalWithPage user personal comment list
// @Summary user personal comment list
// @Description user personal comment list
// @Tags Comment
// @Produce json
// @Param page query int false "page"
// @Param page_size query int false "page size"
// @Param username query string false "username"
// @Success 200 {object} handler.RespBody{data=pager.PageModel{list=[]schema.GetCommentPersonalWithPageResp}}
// @Router /answer/api/v1/personal/comment/page [get]
func (cc *CommentController) GetCommentPersonalWithPage(ctx *gin.Context) {
	req := &schema.GetCommentPersonalWithPageReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.UserID = middleware.GetLoginUserIDFromContext(ctx)

	resp, err := cc.commentService.GetCommentPersonalWithPage(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}

// GetComment godoc
// @Summary get comment by id
// @Description get comment by id
// @Tags Comment
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} handler.RespBody{data=pager.PageModel{list=[]schema.GetCommentResp}}
// @Router /answer/api/v1/comment [get]
func (cc *CommentController) GetComment(ctx *gin.Context) {
	req := &schema.GetCommentReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	canList, err := cc.rankService.CheckOperationPermissions(ctx, req.UserID, []string{
		permission.CommentEdit,
		permission.CommentDelete,
	})
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	req.CanEdit = canList[0]
	req.CanDelete = canList[1]

	resp, err := cc.commentService.GetComment(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}
