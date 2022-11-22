package repo_test

import (
	"context"
	"testing"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/repo/notification"
	"github.com/answerdev/answer/internal/schema"
	"github.com/stretchr/testify/assert"
)

func buildNotificationEntity() *entity.Notification {
	return &entity.Notification{
		UserID:   "1",
		ObjectID: "1",
		Content:  "1",
		Type:     schema.NotificationTypeInbox,
		IsRead:   schema.NotificationNotRead,
		Status:   schema.NotificationStatusNormal,
	}
}

func Test_notificationRepo_ClearIDUnRead(t *testing.T) {
	notificationRepo := notification.NewNotificationRepo(testDataSource)
	ent := buildNotificationEntity()
	err := notificationRepo.AddNotification(context.TODO(), ent)
	assert.NoError(t, err)

	err = notificationRepo.ClearIDUnRead(context.TODO(), ent.UserID, ent.ID)
	assert.NoError(t, err)

	got, exists, err := notificationRepo.GetById(context.TODO(), ent.ID)
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.Equal(t, schema.NotificationRead, got.IsRead)
}

func Test_notificationRepo_ClearUnRead(t *testing.T) {
	notificationRepo := notification.NewNotificationRepo(testDataSource)
	ent := buildNotificationEntity()
	err := notificationRepo.AddNotification(context.TODO(), ent)
	assert.NoError(t, err)

	err = notificationRepo.ClearUnRead(context.TODO(), ent.UserID, ent.Type)
	assert.NoError(t, err)

	got, exists, err := notificationRepo.GetById(context.TODO(), ent.ID)
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.Equal(t, schema.NotificationRead, got.IsRead)
}

func Test_notificationRepo_GetById(t *testing.T) {
	notificationRepo := notification.NewNotificationRepo(testDataSource)
	ent := buildNotificationEntity()
	err := notificationRepo.AddNotification(context.TODO(), ent)
	assert.NoError(t, err)

	got, exists, err := notificationRepo.GetById(context.TODO(), ent.ID)
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.Equal(t, got.ID, ent.ID)
}

func Test_notificationRepo_GetByUserIdObjectIdTypeId(t *testing.T) {
	notificationRepo := notification.NewNotificationRepo(testDataSource)
	ent := buildNotificationEntity()
	err := notificationRepo.AddNotification(context.TODO(), ent)
	assert.NoError(t, err)

	got, exists, err := notificationRepo.GetByUserIdObjectIdTypeId(context.TODO(), ent.UserID, ent.ObjectID, ent.Type)
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.Equal(t, got.ObjectID, ent.ObjectID)
}

func Test_notificationRepo_GetNotificationPage(t *testing.T) {
	notificationRepo := notification.NewNotificationRepo(testDataSource)
	ent := buildNotificationEntity()
	err := notificationRepo.AddNotification(context.TODO(), ent)
	assert.NoError(t, err)

	notificationPage, total, err := notificationRepo.GetNotificationPage(context.TODO(), &schema.NotificationSearch{UserID: ent.UserID})
	assert.NoError(t, err)
	assert.True(t, total > 0)
	assert.Equal(t, notificationPage[0].UserID, ent.UserID)
}

func Test_notificationRepo_UpdateNotificationContent(t *testing.T) {
	notificationRepo := notification.NewNotificationRepo(testDataSource)
	ent := buildNotificationEntity()
	err := notificationRepo.AddNotification(context.TODO(), ent)
	assert.NoError(t, err)

	ent.Content = "test"
	err = notificationRepo.UpdateNotificationContent(context.TODO(), ent)
	assert.NoError(t, err)

	got, exists, err := notificationRepo.GetById(context.TODO(), ent.ID)
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.Equal(t, got.Content, ent.Content)
}
