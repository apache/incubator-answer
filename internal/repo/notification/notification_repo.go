package notification

import (
	"context"
	"time"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	notficationcommon "github.com/answerdev/answer/internal/service/notification_common"
	"github.com/segmentfault/pacman/errors"
)

// notificationRepo notification repository
type notificationRepo struct {
	data *data.Data
}

// NewNotificationRepo new repository
func NewNotificationRepo(data *data.Data) notficationcommon.NotificationRepo {
	return &notificationRepo{
		data: data,
	}
}

// AddNotification add notification
func (nr *notificationRepo) AddNotification(ctx context.Context, notification *entity.Notification) (err error) {
	_, err = nr.data.DB.Insert(notification)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return
	}
	return
}

func (nr *notificationRepo) UpdateNotificationContent(ctx context.Context, notification *entity.Notification) (err error) {
	now := time.Now()
	notification.UpdatedAt = now
	_, err = nr.data.DB.Where("id =?", notification.ID).Cols("content", "updated_at").Update(notification)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return
	}
	return
}

func (nr *notificationRepo) ClearUnRead(ctx context.Context, userID string, notificationType int) (err error) {
	info := &entity.Notification{}
	info.IsRead = schema.NotificationRead
	_, err = nr.data.DB.Where("user_id =?", userID).And("type =?", notificationType).Cols("is_read").Update(info)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return
	}
	return
}

func (nr *notificationRepo) ClearIDUnRead(ctx context.Context, userID string, id string) (err error) {
	info := &entity.Notification{}
	info.IsRead = schema.NotificationRead
	_, err = nr.data.DB.Where("user_id =?", userID).And("id =?", id).Cols("is_read").Update(info)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return
	}
	return
}

func (nr *notificationRepo) GetById(ctx context.Context, id string) (*entity.Notification, bool, error) {
	info := &entity.Notification{}
	exist, err := nr.data.DB.Where("id = ? ", id).Get(info)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return info, false, err
	}
	return info, exist, nil
}

func (nr *notificationRepo) GetByUserIdObjectIdTypeId(ctx context.Context, userID, objectID string, notificationType int) (*entity.Notification, bool, error) {
	info := &entity.Notification{}
	exist, err := nr.data.DB.Where("user_id = ? ", userID).And("object_id = ?", objectID).And("type = ?", notificationType).Get(info)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return info, false, err
	}
	return info, exist, nil
}

func (nr *notificationRepo) SearchList(ctx context.Context, search *schema.NotificationSearch) ([]*entity.Notification, int64, error) {
	var count int64
	var err error

	rows := make([]*entity.Notification, 0)
	if search.UserID == "" {
		return rows, 0, nil
	}
	if search.Page > 0 {
		search.Page = search.Page - 1
	} else {
		search.Page = 0
	}
	if search.PageSize == 0 {
		search.PageSize = constant.Default_PageSize
	}
	offset := search.Page * search.PageSize
	session := nr.data.DB.Where("")
	session = session.And("user_id = ?", search.UserID)
	session = session.And("type = ?", search.Type)
	session = session.OrderBy("updated_at desc")
	session = session.Limit(search.PageSize, offset)
	count, err = session.FindAndCount(&rows)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		return rows, count, err
	}
	return rows, count, nil
}
