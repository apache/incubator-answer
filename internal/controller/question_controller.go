package controller

import (
	"context"

	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service"
	"github.com/answerdev/answer/internal/service/rank"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
)

// QuestionController question controller
type QuestionController struct {
	questionService *service.QuestionService
	rankService     *rank.RankService
}

// NewQuestionController new controller
func NewQuestionController(questionService *service.QuestionService, rankService *rank.RankService) *QuestionController {
	return &QuestionController{questionService: questionService, rankService: rankService}
}

// RemoveQuestion delete question
// @Summary delete question
// @Description delete question
// @Tags api-question
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.RemoveQuestionReq true "question"
// @Success 200 {object} handler.RespBody
// @Router  /answer/api/v1/question [delete]
func (qc *QuestionController) RemoveQuestion(ctx *gin.Context) {
	req := &schema.RemoveQuestionReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	req.IsAdmin = middleware.GetIsAdminFromContext(ctx)
	if can, err := qc.rankService.CheckRankPermission(ctx, req.UserID, rank.QuestionDeleteRank, req.ID); err != nil || !can {
		handler.HandleResponse(ctx, errors.Forbidden(reason.RankFailToMeetTheCondition), errors.Forbidden(reason.RankFailToMeetTheCondition))
		return
	}

	err := qc.questionService.RemoveQuestion(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// CloseQuestion Close question
// @Summary Close question
// @Description Close question
// @Tags api-question
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.CloseQuestionReq true "question"
// @Success 200 {object} handler.RespBody
// @Router  /answer/api/v1/question/status [put]
func (qc *QuestionController) CloseQuestion(ctx *gin.Context) {
	req := &schema.CloseQuestionReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	req.IsAdmin = middleware.GetIsAdminFromContext(ctx)
	err := qc.questionService.CloseQuestion(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// GetQuestion godoc
// @Summary GetQuestion Question
// @Description GetQuestion Question
// @Tags api-question
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Param id query string true "Question TagID"  default(1)
// @Success 200 {string} string ""
// @Router /answer/api/v1/question/info [get]
func (qc *QuestionController) GetQuestion(ctx *gin.Context) {
	id := ctx.Query("id")
	userID := middleware.GetLoginUserIDFromContext(ctx)
	isAdmin := middleware.GetIsAdminFromContext(ctx)
	info, err := qc.questionService.GetQuestion(ctx, id, userID, true, isAdmin)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, info)
}

// SimilarQuestion godoc
// @Summary Search Similar Question
// @Description Search Similar Question
// @Tags api-question
// @Accept  json
// @Produce  json
// @Param question_id query string true "question_id"  default()
// @Success 200 {string} string ""
// @Router /answer/api/v1/question/similar/tag [get]
func (qc *QuestionController) SimilarQuestion(ctx *gin.Context) {
	questionID := ctx.Query("question_id")
	userID := middleware.GetLoginUserIDFromContext(ctx)
	list, count, err := qc.questionService.SimilarQuestion(ctx, questionID, userID)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, gin.H{
		"list":  list,
		"count": count,
	})
}

// Index godoc
// @Summary SearchQuestionList
// @Description SearchQuestionList <br>  "order"  Enums(newest, active,frequent,score,unanswered)
// @Tags api-question
// @Accept  json
// @Produce  json
// @Param data body schema.QuestionSearch  true "QuestionSearch"
// @Success 200 {string} string ""
// @Router /answer/api/v1/question/page [get]
func (qc *QuestionController) Index(ctx *gin.Context) {
	req := &schema.QuestionSearch{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	userID := middleware.GetLoginUserIDFromContext(ctx)
	list, count, err := qc.questionService.SearchList(ctx, req, userID)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	handler.HandleResponse(ctx, nil, gin.H{
		"list":  list,
		"count": count,
	})
}

// SearchList godoc
// @Summary SearchQuestionList
// @Description SearchQuestionList
// @Tags api-question
// @Accept  json
// @Produce  json
// @Param data body schema.QuestionSearch  true "QuestionSearch"
// @Router  /answer/api/v1/question/search [post]
// @Success 200 {string} string ""
func (qc *QuestionController) SearchList(c *gin.Context) {
	Request := new(schema.QuestionSearch)
	err := c.BindJSON(Request)
	if err != nil {
		handler.HandleResponse(c, err, nil)
		return
	}
	ctx := context.Background()
	userID := middleware.GetLoginUserIDFromContext(c)
	list, count, err := qc.questionService.SearchList(ctx, Request, userID)
	if err != nil {
		handler.HandleResponse(c, err, nil)
		return
	}
	handler.HandleResponse(c, nil, gin.H{
		"list":  list,
		"count": count,
	})
}

// AddQuestion add question
// @Summary add question
// @Description add question
// @Tags api-question
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.QuestionAdd true "question"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/question [post]
func (qc *QuestionController) AddQuestion(ctx *gin.Context) {
	req := &schema.QuestionAdd{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)

	if can, err := qc.rankService.CheckRankPermission(ctx, req.UserID, rank.QuestionAddRank, ""); err != nil || !can {
		handler.HandleResponse(ctx, err, errors.Forbidden(reason.RankFailToMeetTheCondition))
		return
	}

	resp, err := qc.questionService.AddQuestion(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}

// UpdateQuestion update question
// @Summary update question
// @Description update question
// @Tags api-question
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.QuestionUpdate true "question"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/question [put]
func (qc *QuestionController) UpdateQuestion(ctx *gin.Context) {
	req := &schema.QuestionUpdate{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	req.IsAdmin = middleware.GetIsAdminFromContext(ctx)
	if can, err := qc.rankService.CheckRankPermission(ctx, req.UserID, rank.QuestionEditRank, req.ID); err != nil || !can {
		handler.HandleResponse(ctx, err, errors.Forbidden(reason.RankFailToMeetTheCondition))
		return
	}

	resp, err := qc.questionService.UpdateQuestion(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}

// CheckCanUpdateQuestion check can update question
// @Summary check can update question
// @Description check can update question
// @Tags Revision
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id query string true "id"  default(string)
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/revisions/edit/check [get]
func (qc *QuestionController) CheckCanUpdateQuestion(ctx *gin.Context) {
	req := &schema.CheckCanQuestionUpdate{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	req.IsAdmin = middleware.GetIsAdminFromContext(ctx)
	if can, err := qc.rankService.CheckRankPermission(ctx, req.UserID, rank.QuestionEditRank, req.ID); err != nil || !can {
		handler.HandleResponse(ctx, err, errors.Forbidden(reason.RankFailToMeetTheCondition))
		return
	}

	resp, err := qc.questionService.CheckCanUpdate(ctx, req)
	handler.HandleResponse(ctx, err, gin.H{
		"unreviewed": resp,
	})
}

// CloseMsgList close question msg list
// @Summary close question msg list
// @Description close question msg list
// @Tags api-question
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/question/closemsglist [get]
func (qc *QuestionController) CloseMsgList(ctx *gin.Context) {
	resp, err := qc.questionService.CloseMsgList(ctx, handler.GetLang(ctx))
	handler.HandleResponse(ctx, err, resp)
}

// SearchByTitleLike add question title like
// @Summary add question title like
// @Description add question title like
// @Tags api-question
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param title query string true "title"  default(string)
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/question/similar [get]
func (qc *QuestionController) SearchByTitleLike(ctx *gin.Context) {
	title := ctx.Query("title")
	userID := middleware.GetLoginUserIDFromContext(ctx)
	resp, err := qc.questionService.SearchByTitleLike(ctx, title, userID)
	handler.HandleResponse(ctx, err, resp)
}

// UserTop godoc
// @Summary UserTop
// @Description UserTop
// @Tags api-question
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param username query string true "username"  default(string)
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/personal/qa/top [get]
func (qc *QuestionController) UserTop(ctx *gin.Context) {
	userName := ctx.Query("username")
	userID := middleware.GetLoginUserIDFromContext(ctx)
	questionList, answerList, err := qc.questionService.SearchUserTopList(ctx, userName, userID)
	handler.HandleResponse(ctx, err, gin.H{
		"question": questionList,
		"answer":   answerList,
	})
}

// UserList godoc
// @Summary UserList
// @Description UserList
// @Tags api-question
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param username query string true "username"  default(string)
// @Param order query string true "order"  Enums(newest,score)
// @Param page query string true "page"  default(0)
// @Param pagesize query string true "pagesize"  default(20)
// @Success 200 {object} handler.RespBody
// @Router /personal/question/page [get]
func (qc *QuestionController) UserList(ctx *gin.Context) {
	userName := ctx.Query("username")
	order := ctx.Query("order")
	pageStr := ctx.Query("page")
	pageSizeStr := ctx.Query("pagesize")
	page := converter.StringToInt(pageStr)
	pageSize := converter.StringToInt(pageSizeStr)
	userID := middleware.GetLoginUserIDFromContext(ctx)
	questionList, count, err := qc.questionService.SearchUserList(ctx, userName, order, page, pageSize, userID)
	handler.HandleResponse(ctx, err, gin.H{
		"list":  questionList,
		"count": count,
	})
}

// UserAnswerList godoc
// @Summary UserAnswerList
// @Description UserAnswerList
// @Tags api-answer
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param username query string true "username"  default(string)
// @Param order query string true "order"  Enums(newest,score)
// @Param page query string true "page"  default(0)
// @Param pagesize query string true "pagesize"  default(20)
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/personal/answer/page [get]
func (qc *QuestionController) UserAnswerList(ctx *gin.Context) {
	userName := ctx.Query("username")
	order := ctx.Query("order")
	pageStr := ctx.Query("page")
	pageSizeStr := ctx.Query("pagesize")
	page := converter.StringToInt(pageStr)
	pageSize := converter.StringToInt(pageSizeStr)
	userID := middleware.GetLoginUserIDFromContext(ctx)
	questionList, count, err := qc.questionService.SearchUserAnswerList(ctx, userName, order, page, pageSize, userID)
	handler.HandleResponse(ctx, err, gin.H{
		"list":  questionList,
		"count": count,
	})
}

// UserCollectionList godoc
// @Summary UserCollectionList
// @Description UserCollectionList
// @Tags Collection
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query string true "page"  default(0)
// @Param pagesize query string true "pagesize"  default(20)
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/personal/collection/page [get]
func (qc *QuestionController) UserCollectionList(ctx *gin.Context) {
	pageStr := ctx.Query("page")
	pageSizeStr := ctx.Query("pagesize")
	page := converter.StringToInt(pageStr)
	pageSize := converter.StringToInt(pageSizeStr)
	userID := middleware.GetLoginUserIDFromContext(ctx)
	questionList, count, err := qc.questionService.SearchUserCollectionList(ctx, page, pageSize, userID)
	handler.HandleResponse(ctx, err, gin.H{
		"list":  questionList,
		"count": count,
	})
}

// CmsSearchList godoc
// @Summary CmsSearchList
// @Description Status:[available,closed,deleted]
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "page size"
// @Param page_size query int false "page size"
// @Param status query string false "user status" Enums(available, closed, deleted)
// @Param query query string false "question id or title"
// @Success 200 {object} handler.RespBody
// @Router /answer/admin/api/question/page [get]
func (qc *QuestionController) CmsSearchList(ctx *gin.Context) {
	req := &schema.CmsQuestionSearch{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	userID := middleware.GetLoginUserIDFromContext(ctx)
	questionList, count, err := qc.questionService.CmsSearchList(ctx, req, userID)
	handler.HandleResponse(ctx, err, gin.H{
		"list":  questionList,
		"count": count,
	})
}

// CmsSearchAnswerList godoc
// @Summary CmsSearchList
// @Description Status:[available,deleted]
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "page size"
// @Param page_size query int false "page size"
// @Param status query string false "user status" Enums(available,deleted)
// @Param query query string false "answer id or question title"
// @Param question_id query string false "question id"
// @Success 200 {object} handler.RespBody
// @Router /answer/admin/api/answer/page [get]
func (qc *QuestionController) CmsSearchAnswerList(ctx *gin.Context) {
	req := &entity.CmsAnswerSearch{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	userID := middleware.GetLoginUserIDFromContext(ctx)
	questionList, count, err := qc.questionService.CmsSearchAnswerList(ctx, req, userID)
	handler.HandleResponse(ctx, err, gin.H{
		"list":  questionList,
		"count": count,
	})
}

// AdminSetQuestionStatus godoc
// @Summary AdminSetQuestionStatus
// @Description Status:[available,closed,deleted]
// @Tags admin
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.AdminSetQuestionStatusRequest true "AdminSetQuestionStatusRequest"
// @Router /answer/admin/api/question/status [put]
// @Success 200 {object} handler.RespBody
func (qc *QuestionController) AdminSetQuestionStatus(ctx *gin.Context) {
	req := &schema.AdminSetQuestionStatusRequest{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	err := qc.questionService.AdminSetQuestionStatus(ctx, req.QuestionID, req.StatusStr)
	handler.HandleResponse(ctx, err, gin.H{})
}
