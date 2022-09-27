package notification

import (
	"context"

	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/base/pager"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/service/notification"
	"github.com/segmentfault/pacman/errors"
)

// notificationReadRepo notificationRead repository
type notificationReadRepo struct {
	data *data.Data
}

// NewNotificationReadRepo new repository
func NewNotificationReadRepo(data *data.Data) notification.NotificationReadRepo {
	return &notificationReadRepo{
		data: data,
	}
}

// AddNotificationRead add notification read record
func (nr *notificationReadRepo) AddNotificationRead(ctx context.Context, notificationRead *entity.NotificationRead) (err error) {
	_, err = nr.data.DB.Insert(notificationRead)
	return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
}

// RemoveNotificationRead delete notification read record
func (nr *notificationReadRepo) RemoveNotificationRead(ctx context.Context, id int) (err error) {
	_, err = nr.data.DB.ID(id).Delete(&entity.NotificationRead{})
	return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
}

// UpdateNotificationRead update notification read record
func (nr *notificationReadRepo) UpdateNotificationRead(ctx context.Context, notificationRead *entity.NotificationRead) (err error) {
	_, err = nr.data.DB.ID(notificationRead.ID).Update(notificationRead)
	return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
}

// GetNotificationRead get notification read record one
func (nr *notificationReadRepo) GetNotificationRead(ctx context.Context, id int) (
	notificationRead *entity.NotificationRead, exist bool, err error) {
	notificationRead = &entity.NotificationRead{}
	exist, err = nr.data.DB.ID(id).Get(notificationRead)
	if err != nil {
		return nil, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetNotificationReadList get notification read record list all
func (nr *notificationReadRepo) GetNotificationReadList(ctx context.Context, notificationRead *entity.NotificationRead) (notificationReadList []*entity.NotificationRead, err error) {
	notificationReadList = make([]*entity.NotificationRead, 0)
	err = nr.data.DB.Find(notificationReadList, notificationRead)
	err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	return
}

// GetNotificationReadPage get notification read record page
func (nr *notificationReadRepo) GetNotificationReadPage(ctx context.Context, page, pageSize int, notificationRead *entity.NotificationRead) (notificationReadList []*entity.NotificationRead, total int64, err error) {
	notificationReadList = make([]*entity.NotificationRead, 0)
	total, err = pager.Help(page, pageSize, notificationReadList, notificationRead, nr.data.DB.NewSession())
	err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	return
}
