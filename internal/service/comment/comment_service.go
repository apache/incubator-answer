package comment

import (
	"context"
	"time"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_common"
	"github.com/answerdev/answer/internal/service/activity_queue"
	"github.com/answerdev/answer/internal/service/comment_common"
	"github.com/answerdev/answer/internal/service/export"
	"github.com/answerdev/answer/internal/service/notice_queue"
	"github.com/answerdev/answer/internal/service/object_info"
	"github.com/answerdev/answer/internal/service/permission"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/answerdev/answer/pkg/encryption"
	"github.com/answerdev/answer/pkg/htmltext"
	"github.com/answerdev/answer/pkg/uid"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/i18n"
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

// CommentService user service
type CommentService struct {
	commentRepo       CommentRepo
	commentCommonRepo comment_common.CommentCommonRepo
	userCommon        *usercommon.UserCommon
	voteCommon        activity_common.VoteRepo
	objectInfoService *object_info.ObjService
	emailService      *export.EmailService
	userRepo          usercommon.UserRepo
}

// NewCommentService new comment service
func NewCommentService(
	commentRepo CommentRepo,
	commentCommonRepo comment_common.CommentCommonRepo,
	userCommon *usercommon.UserCommon,
	objectInfoService *object_info.ObjService,
	voteCommon activity_common.VoteRepo,
	emailService *export.EmailService,
	userRepo usercommon.UserRepo,
) *CommentService {
	return &CommentService{
		commentRepo:       commentRepo,
		commentCommonRepo: commentCommonRepo,
		userCommon:        userCommon,
		voteCommon:        voteCommon,
		objectInfoService: objectInfoService,
		emailService:      emailService,
		userRepo:          userRepo,
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
	objInfo.ObjectID = uid.DeShortID(objInfo.ObjectID)
	objInfo.QuestionID = uid.DeShortID(objInfo.QuestionID)
	objInfo.AnswerID = uid.DeShortID(objInfo.AnswerID)
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

	resp = &schema.GetCommentResp{}
	resp.SetFromComment(comment)
	resp.MemberActions = permission.GetCommentPermission(ctx, req.UserID, resp.UserID,
		time.Now(), req.CanEdit, req.CanDelete)

	commentResp, err := cs.addCommentNotification(ctx, req, resp, comment, objInfo)
	if err != nil {
		return commentResp, err
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

	activityMsg := &schema.ActivityMsg{
		UserID:           comment.UserID,
		ObjectID:         comment.ID,
		OriginalObjectID: req.ObjectID,
		ActivityTypeKey:  constant.ActQuestionCommented,
	}
	switch objInfo.ObjectType {
	case constant.QuestionObjectType:
		activityMsg.ActivityTypeKey = constant.ActQuestionCommented
	case constant.AnswerObjectType:
		activityMsg.ActivityTypeKey = constant.ActAnswerCommented
	}
	activity_queue.AddActivity(activityMsg)
	return resp, nil
}

func (cs *CommentService) addCommentNotification(
	ctx context.Context, req *schema.AddCommentReq, resp *schema.GetCommentResp,
	comment *entity.Comment, objInfo *schema.SimpleObjectInfo) (*schema.GetCommentResp, error) {
	// The priority of the notification
	// 1. reply to user
	// 2. comment mention to user
	// 3. answer or question was commented
	alreadyNotifiedUserID := make(map[string]bool)

	// get reply user info
	if len(resp.ReplyUserID) > 0 && resp.ReplyUserID != req.UserID {
		replyUser, exist, err := cs.userCommon.GetUserBasicInfoByID(ctx, resp.ReplyUserID)
		if err != nil {
			return nil, err
		}
		if exist {
			resp.ReplyUsername = replyUser.Username
			resp.ReplyUserDisplayName = replyUser.DisplayName
			resp.ReplyUserStatus = replyUser.Status
		}
		cs.notificationCommentReply(ctx, replyUser.ID, comment.ID, req.UserID)
		alreadyNotifiedUserID[replyUser.ID] = true
		return nil, nil
	}

	if len(req.MentionUsernameList) > 0 {
		alreadyNotifiedUserIDs := cs.notificationMention(
			ctx, req.MentionUsernameList, comment.ID, req.UserID, alreadyNotifiedUserID)
		for _, userID := range alreadyNotifiedUserIDs {
			alreadyNotifiedUserID[userID] = true
		}
		return nil, nil
	}

	if objInfo.ObjectType == constant.QuestionObjectType && !alreadyNotifiedUserID[objInfo.ObjectCreatorUserID] {
		cs.notificationQuestionComment(ctx, objInfo.ObjectCreatorUserID,
			objInfo.QuestionID, objInfo.Title, comment.ID, req.UserID, comment.OriginalText)
	} else if objInfo.ObjectType == constant.AnswerObjectType && !alreadyNotifiedUserID[objInfo.ObjectCreatorUserID] {
		cs.notificationAnswerComment(ctx, objInfo.QuestionID, objInfo.Title, objInfo.AnswerID,
			objInfo.ObjectCreatorUserID, comment.ID, req.UserID, comment.OriginalText)
	}
	return nil, nil
}

// RemoveComment delete comment
func (cs *CommentService) RemoveComment(ctx context.Context, req *schema.RemoveCommentReq) (err error) {
	return cs.commentRepo.RemoveComment(ctx, req.CommentID)
}

// UpdateComment update comment
func (cs *CommentService) UpdateComment(ctx context.Context, req *schema.UpdateCommentReq) (
	resp *schema.GetCommentResp, err error) {
	resp = &schema.GetCommentResp{}

	old, exist, err := cs.commentCommonRepo.GetComment(ctx, req.CommentID)
	if err != nil {
		return
	}
	if !exist {
		return resp, errors.BadRequest(reason.CommentNotFound)
	}

	// user can edit the comment that was posted by himself before deadline.
	if !req.IsAdmin && (time.Now().After(old.CreatedAt.Add(constant.CommentEditDeadline))) {
		return resp, errors.BadRequest(reason.CommentCannotEditAfterDeadline)
	}

	comment := &entity.Comment{}
	_ = copier.Copy(comment, req)
	comment.ID = req.CommentID
	resp.SetFromComment(comment)
	resp.MemberActions = permission.GetCommentPermission(ctx, req.UserID, resp.UserID,
		time.Now(), req.CanEdit, req.CanDelete)
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
	return resp, cs.commentRepo.UpdateComment(ctx, comment)
}

// GetComment get comment one
func (cs *CommentService) GetComment(ctx context.Context, req *schema.GetCommentReq) (resp *schema.GetCommentResp, err error) {
	comment, exist, err := cs.commentCommonRepo.GetComment(ctx, req.ID)
	if err != nil {
		return
	}
	if !exist {
		return nil, errors.BadRequest(reason.CommentNotFound)
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

	resp.MemberActions = permission.GetCommentPermission(ctx, req.UserID, resp.UserID,
		comment.CreatedAt, req.CanEdit, req.CanDelete)
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
		commentResp, err := cs.convertCommentEntity2Resp(ctx, req, comment)
		if err != nil {
			return nil, err
		}
		resp = append(resp, commentResp)
	}

	// if user request the specific comment, add it if not exist.
	if len(req.CommentID) > 0 {
		commentExist := false
		for _, t := range resp {
			if t.CommentID == req.CommentID {
				commentExist = true
				break
			}
		}
		if !commentExist {
			comment, exist, err := cs.commentCommonRepo.GetComment(ctx, req.CommentID)
			if err != nil {
				return nil, err
			}
			if exist && comment.ObjectID == req.ObjectID {
				commentResp, err := cs.convertCommentEntity2Resp(ctx, req, comment)
				if err != nil {
					return nil, err
				}
				resp = append(resp, commentResp)
			}
		}
	}
	return pager.NewPageModel(total, resp), nil
}

func (cs *CommentService) convertCommentEntity2Resp(ctx context.Context, req *schema.GetCommentWithPageReq,
	comment *entity.Comment) (commentResp *schema.GetCommentResp, err error) {
	commentResp = &schema.GetCommentResp{
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

	commentResp.MemberActions = permission.GetCommentPermission(ctx,
		req.UserID, commentResp.UserID, comment.CreatedAt, req.CanEdit, req.CanDelete)
	return commentResp, nil
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
				commentResp.UrlTitle = htmltext.UrlTitle(objInfo.Title)
				commentResp.QuestionID = objInfo.QuestionID
				commentResp.AnswerID = objInfo.AnswerID
				if objInfo.QuestionStatus == entity.QuestionStatusDeleted {
					commentResp.Title = "Deleted question"
				}
			}
		}
		resp = append(resp, commentResp)
	}
	return pager.NewPageModel(total, resp), nil
}

func (cs *CommentService) notificationQuestionComment(ctx context.Context, questionUserID,
	questionID, questionTitle, commentID, commentUserID, commentSummary string) {
	if questionUserID == commentUserID {
		return
	}
	msg := &schema.NotificationMsg{
		ReceiverUserID: questionUserID,
		TriggerUserID:  commentUserID,
		Type:           schema.NotificationTypeInbox,
		ObjectID:       commentID,
	}
	msg.ObjectType = constant.CommentObjectType
	msg.NotificationAction = constant.NotificationCommentQuestion
	notice_queue.AddNotification(msg)

	receiverUserInfo, exist, err := cs.userRepo.GetByUserID(ctx, questionUserID)
	if err != nil {
		log.Error(err)
		return
	}
	if !exist {
		log.Warnf("user %s not found", questionUserID)
		return
	}
	if receiverUserInfo.NoticeStatus == schema.NoticeStatusOff || len(receiverUserInfo.EMail) == 0 {
		return
	}

	rawData := &schema.NewCommentTemplateRawData{
		QuestionTitle:   questionTitle,
		QuestionID:      questionID,
		CommentID:       commentID,
		CommentSummary:  commentSummary,
		UnsubscribeCode: encryption.MD5(receiverUserInfo.Pass),
	}
	commentUser, _, _ := cs.userCommon.GetUserBasicInfoByID(ctx, commentUserID)
	if commentUser != nil {
		rawData.CommentUserDisplayName = commentUser.DisplayName
	}
	codeContent := &schema.EmailCodeContent{
		SourceType: schema.UnsubscribeSourceType,
		Email:      receiverUserInfo.EMail,
		UserID:     receiverUserInfo.ID,
	}

	// If receiver has set language, use it to send email.
	if len(receiverUserInfo.Language) > 0 {
		ctx = context.WithValue(ctx, constant.AcceptLanguageFlag, i18n.Language(receiverUserInfo.Language))
	}
	title, body, err := cs.emailService.NewCommentTemplate(ctx, rawData)
	if err != nil {
		log.Error(err)
		return
	}

	go cs.emailService.SendAndSaveCodeWithTime(
		ctx, receiverUserInfo.EMail, title, body, rawData.UnsubscribeCode, codeContent.ToJSONString(), 7*24*time.Hour)
}

func (cs *CommentService) notificationAnswerComment(ctx context.Context,
	questionID, questionTitle, answerID, answerUserID, commentID, commentUserID, commentSummary string) {
	if answerUserID == commentUserID {
		return
	}
	msg := &schema.NotificationMsg{
		ReceiverUserID: answerUserID,
		TriggerUserID:  commentUserID,
		Type:           schema.NotificationTypeInbox,
		ObjectID:       commentID,
	}
	msg.ObjectType = constant.CommentObjectType
	msg.NotificationAction = constant.NotificationCommentAnswer
	notice_queue.AddNotification(msg)

	receiverUserInfo, exist, err := cs.userRepo.GetByUserID(ctx, answerUserID)
	if err != nil {
		log.Error(err)
		return
	}
	if !exist {
		log.Warnf("user %s not found", answerUserID)
		return
	}
	if receiverUserInfo.NoticeStatus == schema.NoticeStatusOff || len(receiverUserInfo.EMail) == 0 {
		return
	}

	rawData := &schema.NewCommentTemplateRawData{
		QuestionTitle:   questionTitle,
		QuestionID:      questionID,
		AnswerID:        answerID,
		CommentID:       commentID,
		CommentSummary:  commentSummary,
		UnsubscribeCode: encryption.MD5(receiverUserInfo.Pass),
	}
	commentUser, _, _ := cs.userCommon.GetUserBasicInfoByID(ctx, commentUserID)
	if commentUser != nil {
		rawData.CommentUserDisplayName = commentUser.DisplayName
	}
	codeContent := &schema.EmailCodeContent{
		SourceType: schema.UnsubscribeSourceType,
		Email:      receiverUserInfo.EMail,
		UserID:     receiverUserInfo.ID,
	}

	// If receiver has set language, use it to send email.
	if len(receiverUserInfo.Language) > 0 {
		ctx = context.WithValue(ctx, constant.AcceptLanguageFlag, i18n.Language(receiverUserInfo.Language))
	}
	title, body, err := cs.emailService.NewCommentTemplate(ctx, rawData)
	if err != nil {
		log.Error(err)
		return
	}

	go cs.emailService.SendAndSaveCodeWithTime(
		ctx, receiverUserInfo.EMail, title, body, rawData.UnsubscribeCode, codeContent.ToJSONString(), 7*24*time.Hour)
}

func (cs *CommentService) notificationCommentReply(ctx context.Context, replyUserID, commentID, commentUserID string) {
	msg := &schema.NotificationMsg{
		ReceiverUserID: replyUserID,
		TriggerUserID:  commentUserID,
		Type:           schema.NotificationTypeInbox,
		ObjectID:       commentID,
	}
	msg.ObjectType = constant.CommentObjectType
	msg.NotificationAction = constant.NotificationReplyToYou
	notice_queue.AddNotification(msg)
}

func (cs *CommentService) notificationMention(
	ctx context.Context, mentionUsernameList []string, commentID, commentUserID string,
	alreadyNotifiedUserID map[string]bool) (alreadyNotifiedUserIDs []string) {
	for _, username := range mentionUsernameList {
		userInfo, exist, err := cs.userCommon.GetUserBasicInfoByUserName(ctx, username)
		if err != nil {
			log.Error(err)
			continue
		}
		if exist && !alreadyNotifiedUserID[userInfo.ID] {
			msg := &schema.NotificationMsg{
				ReceiverUserID: userInfo.ID,
				TriggerUserID:  commentUserID,
				Type:           schema.NotificationTypeInbox,
				ObjectID:       commentID,
			}
			msg.ObjectType = constant.CommentObjectType
			msg.NotificationAction = constant.NotificationMentionYou
			notice_queue.AddNotification(msg)
			alreadyNotifiedUserIDs = append(alreadyNotifiedUserIDs, userInfo.ID)
		}
	}
	return alreadyNotifiedUserIDs
}
