package activity

import (
	"context"
	"strings"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/repo/config"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity_common"
	"github.com/answerdev/answer/internal/service/comment_common"
	"github.com/answerdev/answer/internal/service/object_info"
	"github.com/answerdev/answer/internal/service/tag_common"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/segmentfault/pacman/log"
)

// ActivityRepo activity repository
type ActivityRepo interface {
	GetObjectAllActivity(ctx context.Context, objectID string, showVote bool) (activityList []*entity.Activity, err error)
}

// ActivityService activity service
type ActivityService struct {
	activityRepo          ActivityRepo
	userCommon            *usercommon.UserCommon
	activityCommonService *activity_common.ActivityCommon
	tagCommonService      *tag_common.TagCommonService
	objectInfoService     *object_info.ObjService
	commentCommonService  *comment_common.CommentCommonService
}

// NewActivityService new activity service
func NewActivityService(
	activityRepo ActivityRepo,
	userCommon *usercommon.UserCommon,
	activityCommonService *activity_common.ActivityCommon,
	tagCommonService *tag_common.TagCommonService,
	objectInfoService *object_info.ObjService,
	commentCommonService *comment_common.CommentCommonService,
) *ActivityService {
	return &ActivityService{
		objectInfoService:     objectInfoService,
		activityRepo:          activityRepo,
		userCommon:            userCommon,
		activityCommonService: activityCommonService,
		tagCommonService:      tagCommonService,
		commentCommonService:  commentCommonService,
	}
}

// GetObjectTimeline get object timeline
func (as *ActivityService) GetObjectTimeline(ctx context.Context, req *schema.GetObjectTimelineReq) (
	resp *schema.GetObjectTimelineResp, err error) {
	resp = &schema.GetObjectTimelineResp{
		ObjectInfo: &schema.ActObjectInfo{},
		Timeline:   make([]*schema.ActObjectTimeline, 0),
	}

	objInfo, err := as.objectInfoService.GetInfo(ctx, req.ObjectId)
	if err != nil {
		return nil, err
	}
	resp.ObjectInfo.Title = objInfo.Title
	resp.ObjectInfo.ObjectType = objInfo.ObjectType
	resp.ObjectInfo.QuestionID = objInfo.QuestionID
	resp.ObjectInfo.AnswerID = objInfo.AnswerID

	activityList, err := as.activityRepo.GetObjectAllActivity(ctx, req.ObjectId, req.ShowVote)
	if err != nil {
		return nil, err
	}
	for _, act := range activityList {
		item := &schema.ActObjectTimeline{
			ActivityID: act.ID,
			RevisionID: converter.IntToString(act.RevisionID),
			CreatedAt:  act.CreatedAt.Unix(),
			Cancelled:  act.Cancelled == entity.ActivityCancelled,
			ObjectID:   act.ObjectID,
		}
		if item.Cancelled {
			item.CancelledAt = act.CancelledAt.Unix()
		}

		// database save activity type is number, change to activity type string is like "question.asked".
		// so we need to cut the front part of '.'
		item.ObjectType, item.ActivityType, _ = strings.Cut(config.ID2KeyMapping[act.ActivityType], ".")

		isHidden, formattedActivityType := formatActivity(item.ActivityType)
		if isHidden {
			continue
		}
		item.ActivityType = formattedActivityType

		// get user info
		userBasicInfo, exist, err := as.userCommon.GetUserBasicInfoByID(ctx, act.UserID)
		if err != nil {
			return nil, err
		}
		if exist {
			item.Username = userBasicInfo.Username
			item.UserDisplayName = userBasicInfo.DisplayName
		}

		if item.ObjectType == constant.CommentObjectType {
			comment, err := as.commentCommonService.GetComment(ctx, item.ObjectID)
			if err != nil {
				log.Error(err)
			} else {
				item.Comment = comment.ParsedText
			}
		}

		resp.Timeline = append(resp.Timeline, item)
	}
	return
}

func formatActivity(activityType string) (isHidden bool, formattedActivityType string) {
	if activityType == "voted_up" || activityType == "voted_down" || activityType == "accepted" {
		return true, ""
	}
	if activityType == "vote_up" {
		return false, "upvote"
	}
	if activityType == "vote_down" {
		return false, "downvote"
	}
	return false, activityType
}
