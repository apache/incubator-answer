package activity_common

import (
	"context"

	"github.com/answerdev/answer/internal/entity"
	"xorm.io/xorm"
)

type ActivityRepo interface {
	GetActivityTypeByObjID(ctx context.Context, objectId string, action string) (activityType, rank int, hasRank int, err error)
	GetActivityTypeByObjKey(ctx context.Context, objectKey, action string) (activityType int, err error)
	GetActivity(ctx context.Context, session *xorm.Session, objectID, userID string, activityType int) (
		existsActivity *entity.Activity, exist bool, err error)
	GetUserIDObjectIDActivitySum(ctx context.Context, userID, objectID string) (int, error)
}
