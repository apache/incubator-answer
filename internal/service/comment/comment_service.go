package comment

import (
	"context"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_common"
	"github.com/answerdev/answer/internal/service/activity_queue"
	"github.com/answerdev/answer/internal/service/comment_common"
	"github.com/answerdev/answer/internal/service/notice_queue"
	"github.com/answerdev/answer/internal/service/object_info"
	"github.com/answerdev/answer/internal/service/permission"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// CommentRepo comment repository
type CommentRepo interface {
	AddComment(ctx context.Context, comment *entity.Comment) (err error)
	RemoveComment(ctx context.Context, commentID string) (err error)
	UpdateComment(ctx context.Context, comment *entity.Comment) (err error)
	GetComment(ctx context.Context, commentID string) (comment *entity.Comment, exist bool, err error)
	GetCommentPage(ctx context.Context, commentQuery *CommentQuery) (
		comments []*entity.Comment, total int64, err error)
}

// CommentService user service
type CommentService struct {
	commentRepo       CommentRepo
	commentCommonRepo comment_common.CommentCommonRepo
	userCommon        *usercommon.UserCommon
	voteCommon        activity_common.VoteRepo
	objectInfoService *object_info.ObjService
}

type CommentQuery struct {
	pager.PageCond
	// object id
	ObjectID string
	// query condition
	QueryCond string
	// user id
	UserID string
}

func (c *CommentQuery) GetOrderBy() string {
	if c.QueryCond == "vote" {
		return "vote_count DESC,created_at ASC"
	}
	if c.QueryCond == "created_at" {
		return "created_at DESC"
	}
	return "created_at ASC"
}

// NewCommentService new comment service
func NewCommentService(
	commentRepo CommentRepo,
	commentCommonRepo comment_common.CommentCommonRepo,
	userCommon *usercommon.UserCommon,
	objectInfoService *object_info.ObjService,
	voteCommon activity_common.VoteRepo) *CommentService {
	return &CommentService{
		commentRepo:       commentRepo,
		commentCommonRepo: commentCommonRepo,
		userCommon:        userCommon,
		voteCommon:        voteCommon,
		objectInfoService: objectInfoService,
	}
}

// AddComment add comment
func (cs *CommentService) AddComment(ctx context.Context, req *schema.AddCommentReq) (
	resp *schema.GetCommentResp, err error) {
	comment := &entity.Comment{}
	_ = copier.Copy(comment, req)
	comment.Status = entity.CommentStatusAvailable

	// add question id
	objInfo, err := cs.objectInfoService.GetInfo(ctx, req.ObjectID)
	if err != nil {
		return nil, err
	}
	if objInfo.ObjectType == constant.QuestionObjectType || objInfo.ObjectType == constant.AnswerObjectType {
		comment.QuestionID = objInfo.QuestionID
	}

	if len(req.ReplyCommentID) > 0 {
		replyComment, exist, err := cs.commentCommonRepo.GetComment(ctx, req.ReplyCommentID)
		if err != nil {
			return nil, err
		}
		if !exist {
			return nil, errors.BadRequest(reason.CommentNotFound)
		}
		comment.SetReplyUserID(replyComment.UserID)
		comment.SetReplyCommentID(replyComment.ID)
	} else {
		comment.SetReplyUserID("")
		comment.SetReplyCommentID("")
	}

	err = cs.commentRepo.AddComment(ctx, comment)
	if err != nil {
		return nil, err
	}

	if objInfo.ObjectType == constant.QuestionObjectType {
		cs.notificationQuestionComment(ctx, objInfo.ObjectCreatorUserID, comment.ID, req.UserID)
	} else if objInfo.ObjectType == constant.AnswerObjectType {
		cs.notificationAnswerComment(ctx, objInfo.ObjectCreatorUserID, comment.ID, req.UserID)
	}
	if len(req.MentionUsernameList) > 0 {
		cs.notificationMention(ctx, req.MentionUsernameList, comment.ID, req.UserID)
	}

	resp = &schema.GetCommentResp{}
	resp.SetFromComment(comment)
	resp.MemberActions = permission.GetCommentPermission(ctx, req.UserID, resp.UserID, req.CanEdit, req.CanDelete)

	// get reply user info
	if len(resp.ReplyUserID) > 0 {
		replyUser, exist, err := cs.userCommon.GetUserBasicInfoByID(ctx, resp.ReplyUserID)
		if err != nil {
			return nil, err
		}
		if exist {
			resp.ReplyUsername = replyUser.Username
			resp.ReplyUserDisplayName = replyUser.DisplayName
			resp.ReplyUserStatus = replyUser.Status
		}
		cs.notificationCommentReply(ctx, replyUser.ID, objInfo.QuestionID, req.UserID)
	}

	// get user info
	userInfo, exist, err := cs.userCommon.GetUserBasicInfoByID(ctx, resp.UserID)
	if err != nil {
		return nil, err
	}
	if exist {
		resp.Username = userInfo.Username
		resp.UserDisplayName = userInfo.DisplayName
		resp.UserAvatar = userInfo.Avatar
		resp.UserStatus = userInfo.Status
	}

	activity_queue.AddActivity(&schema.ActivityMsg{
		UserID:           comment.UserID,
		ObjectID:         comment.ID,
		OriginalObjectID: req.ObjectID,
		ActivityTypeKey:  constant.ActQuestionCommented,
	})
	return resp, nil
}

// RemoveComment delete comment
func (cs *CommentService) RemoveComment(ctx context.Context, req *schema.RemoveCommentReq) (err error) {
	if err := cs.checkCommentWhetherOwner(ctx, req.UserID, req.CommentID); err != nil {
		return err
	}
	return cs.commentRepo.RemoveComment(ctx, req.CommentID)
}

// UpdateComment update comment
func (cs *CommentService) UpdateComment(ctx context.Context, req *schema.UpdateCommentReq) (err error) {
	if err := cs.checkCommentWhetherOwner(ctx, req.UserID, req.CommentID); err != nil {
		return err
	}
	comment := &entity.Comment{}
	_ = copier.Copy(comment, req)
	comment.ID = req.CommentID
	return cs.commentRepo.UpdateComment(ctx, comment)
}

// GetComment get comment one
func (cs *CommentService) GetComment(ctx context.Context, req *schema.GetCommentReq) (resp *schema.GetCommentResp, err error) {
	comment, exist, err := cs.commentCommonRepo.GetComment(ctx, req.ID)
	if err != nil {
		return
	}
	if !exist {
		return nil, errors.BadRequest(reason.UnknownError)
	}

	resp = &schema.GetCommentResp{
		CommentID:      comment.ID,
		CreatedAt:      comment.CreatedAt.Unix(),
		UserID:         comment.UserID,
		ReplyUserID:    comment.GetReplyUserID(),
		ReplyCommentID: comment.GetReplyCommentID(),
		ObjectID:       comment.ObjectID,
		VoteCount:      comment.VoteCount,
		OriginalText:   comment.OriginalText,
		ParsedText:     comment.ParsedText,
	}

	// get comment user info
	if len(resp.UserID) > 0 {
		commentUser, exist, err := cs.userCommon.GetUserBasicInfoByID(ctx, resp.UserID)
		if err != nil {
			return nil, err
		}
		if exist {
			resp.Username = commentUser.Username
			resp.UserDisplayName = commentUser.DisplayName
			resp.UserAvatar = commentUser.Avatar
			resp.UserStatus = commentUser.Status
		}
	}

	// get reply user info
	if len(resp.ReplyUserID) > 0 {
		replyUser, exist, err := cs.userCommon.GetUserBasicInfoByID(ctx, resp.ReplyUserID)
		if err != nil {
			return nil, err
		}
		if exist {
			resp.ReplyUsername = replyUser.Username
			resp.ReplyUserDisplayName = replyUser.DisplayName
			resp.ReplyUserStatus = replyUser.Status
		}
	}

	// check if current user vote this comment
	resp.IsVote = cs.checkIsVote(ctx, req.UserID, resp.CommentID)

	resp.MemberActions = permission.GetCommentPermission(ctx, req.UserID, resp.UserID, req.CanEdit, req.CanDelete)
	return resp, nil
}

// GetCommentWithPage get comment list page
func (cs *CommentService) GetCommentWithPage(ctx context.Context, req *schema.GetCommentWithPageReq) (
	pageModel *pager.PageModel, err error) {
	dto := &CommentQuery{
		PageCond:  pager.PageCond{Page: req.Page, PageSize: req.PageSize},
		ObjectID:  req.ObjectID,
		QueryCond: req.QueryCond,
	}
	commentList, total, err := cs.commentRepo.GetCommentPage(ctx, dto)
	if err != nil {
		return nil, err
	}
	resp := make([]*schema.GetCommentResp, 0)
	for _, comment := range commentList {
		commentResp := &schema.GetCommentResp{
			CommentID:      comment.ID,
			CreatedAt:      comment.CreatedAt.Unix(),
			UserID:         comment.UserID,
			ReplyUserID:    comment.GetReplyUserID(),
			ReplyCommentID: comment.GetReplyCommentID(),
			ObjectID:       comment.ObjectID,
			VoteCount:      comment.VoteCount,
			OriginalText:   comment.OriginalText,
			ParsedText:     comment.ParsedText,
		}

		// get comment user info
		if len(commentResp.UserID) > 0 {
			commentUser, exist, err := cs.userCommon.GetUserBasicInfoByID(ctx, commentResp.UserID)
			if err != nil {
				return nil, err
			}
			if exist {
				commentResp.Username = commentUser.Username
				commentResp.UserDisplayName = commentUser.DisplayName
				commentResp.UserAvatar = commentUser.Avatar
				commentResp.UserStatus = commentUser.Status
			}
		}

		// get reply user info
		if len(commentResp.ReplyUserID) > 0 {
			replyUser, exist, err := cs.userCommon.GetUserBasicInfoByID(ctx, commentResp.ReplyUserID)
			if err != nil {
				return nil, err
			}
			if exist {
				commentResp.ReplyUsername = replyUser.Username
				commentResp.ReplyUserDisplayName = replyUser.DisplayName
				commentResp.ReplyUserStatus = replyUser.Status
			}
		}

		// check if current user vote this comment
		commentResp.IsVote = cs.checkIsVote(ctx, req.UserID, commentResp.CommentID)

		commentResp.MemberActions = permission.GetCommentPermission(ctx, req.UserID, commentResp.UserID, req.CanEdit, req.CanDelete)
		resp = append(resp, commentResp)
	}
	return pager.NewPageModel(total, resp), nil
}

func (cs *CommentService) checkCommentWhetherOwner(ctx context.Context, userID, commentID string) error {
	// check comment if user self
	comment, exist, err := cs.commentCommonRepo.GetComment(ctx, commentID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.BadRequest(reason.CommentNotFound)
	}
	if comment.UserID != userID {
		return errors.BadRequest(reason.CommentEditWithoutPermission)
	}
	return nil
}

func (cs *CommentService) checkIsVote(ctx context.Context, userID, commentID string) (isVote bool) {
	status := cs.voteCommon.GetVoteStatus(ctx, commentID, userID)
	return len(status) > 0
}

// GetCommentPersonalWithPage get personal comment list page
func (cs *CommentService) GetCommentPersonalWithPage(ctx context.Context, req *schema.GetCommentPersonalWithPageReq) (
	pageModel *pager.PageModel, err error) {
	if len(req.Username) > 0 {
		userInfo, exist, err := cs.userCommon.GetUserBasicInfoByUserName(ctx, req.Username)
		if err != nil {
			return nil, err
		}
		if !exist {
			return nil, errors.BadRequest(reason.UserNotFound)
		}
		req.UserID = userInfo.ID
	}
	if len(req.UserID) == 0 {
		return nil, errors.BadRequest(reason.UserNotFound)
	}

	dto := &CommentQuery{
		PageCond:  pager.PageCond{Page: req.Page, PageSize: req.PageSize},
		UserID:    req.UserID,
		QueryCond: "created_at",
	}
	commentList, total, err := cs.commentRepo.GetCommentPage(ctx, dto)
	if err != nil {
		return nil, err
	}
	resp := make([]*schema.GetCommentPersonalWithPageResp, 0)
	for _, comment := range commentList {
		commentResp := &schema.GetCommentPersonalWithPageResp{
			CommentID: comment.ID,
			CreatedAt: comment.CreatedAt.Unix(),
			ObjectID:  comment.ObjectID,
			Content:   comment.ParsedText, // todo trim
		}
		if len(comment.ObjectID) > 0 {
			objInfo, err := cs.objectInfoService.GetInfo(ctx, comment.ObjectID)
			if err != nil {
				log.Error(err)
			} else {
				commentResp.ObjectType = objInfo.ObjectType
				commentResp.Title = objInfo.Title
				commentResp.QuestionID = objInfo.QuestionID
				commentResp.AnswerID = objInfo.AnswerID
			}
		}
		resp = append(resp, commentResp)
	}
	return pager.NewPageModel(total, resp), nil
}

func (cs *CommentService) notificationQuestionComment(ctx context.Context, questionUserID, commentID, commentUserID string) {
	msg := &schema.NotificationMsg{
		ReceiverUserID: questionUserID,
		TriggerUserID:  commentUserID,
		Type:           schema.NotificationTypeInbox,
		ObjectID:       commentID,
	}
	msg.ObjectType = constant.CommentObjectType
	msg.NotificationAction = constant.CommentQuestion
	notice_queue.AddNotification(msg)
}

func (cs *CommentService) notificationAnswerComment(ctx context.Context, answerUserID, commentID, commentUserID string) {
	msg := &schema.NotificationMsg{
		ReceiverUserID: answerUserID,
		TriggerUserID:  commentUserID,
		Type:           schema.NotificationTypeInbox,
		ObjectID:       commentID,
	}
	msg.ObjectType = constant.CommentObjectType
	msg.NotificationAction = constant.CommentAnswer
	notice_queue.AddNotification(msg)
}

func (cs *CommentService) notificationCommentReply(ctx context.Context, replyUserID, commentID, commentUserID string) {
	msg := &schema.NotificationMsg{
		ReceiverUserID: replyUserID,
		TriggerUserID:  commentUserID,
		Type:           schema.NotificationTypeInbox,
		ObjectID:       commentID,
	}
	msg.ObjectType = constant.CommentObjectType
	msg.NotificationAction = constant.ReplyToYou
	notice_queue.AddNotification(msg)
}

func (cs *CommentService) notificationMention(ctx context.Context, mentionUsernameList []string, commentID, commentUserID string) {
	for _, username := range mentionUsernameList {
		userInfo, exist, err := cs.userCommon.GetUserBasicInfoByUserName(ctx, username)
		if err != nil {
			log.Error(err)
			continue
		}
		if exist {
			msg := &schema.NotificationMsg{
				ReceiverUserID: userInfo.ID,
				TriggerUserID:  commentUserID,
				Type:           schema.NotificationTypeInbox,
				ObjectID:       commentID,
			}
			msg.ObjectType = constant.CommentObjectType
			msg.NotificationAction = constant.MentionYou
			notice_queue.AddNotification(msg)
		}
	}
}
