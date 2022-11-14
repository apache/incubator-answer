package repo_test

import (
	"context"
	"testing"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/repo/auth"
	"github.com/answerdev/answer/internal/repo/user"
	"github.com/stretchr/testify/assert"
)

func Test_userBackyardRepo_GetUserInfo(t *testing.T) {
	userBackyardRepo := user.NewUserBackyardRepo(testDataSource, auth.NewAuthRepo(testDataSource))
	got, exist, err := userBackyardRepo.GetUserInfo(context.TODO(), "1")
	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, "1", got.ID)
}

func Test_userBackyardRepo_GetUserPage(t *testing.T) {
	userBackyardRepo := user.NewUserBackyardRepo(testDataSource, auth.NewAuthRepo(testDataSource))
	got, total, err := userBackyardRepo.GetUserPage(context.TODO(), 1, 1, &entity.User{Username: "admin"}, "")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.Equal(t, "1", got[0].ID)
}

func Test_userBackyardRepo_UpdateUserStatus(t *testing.T) {
	userBackyardRepo := user.NewUserBackyardRepo(testDataSource, auth.NewAuthRepo(testDataSource))
	got, exist, err := userBackyardRepo.GetUserInfo(context.TODO(), "1")
	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, entity.UserStatusAvailable, got.Status)

	err = userBackyardRepo.UpdateUserStatus(context.TODO(), "1", entity.UserStatusSuspended, entity.EmailStatusAvailable,
		"admin@admin.com")
	assert.NoError(t, err)

	got, exist, err = userBackyardRepo.GetUserInfo(context.TODO(), "1")
	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, entity.UserStatusSuspended, got.Status)

	err = userBackyardRepo.UpdateUserStatus(context.TODO(), "1", entity.UserStatusAvailable, entity.EmailStatusAvailable,
		"admin@admin.com")
	assert.NoError(t, err)

	got, exist, err = userBackyardRepo.GetUserInfo(context.TODO(), "1")
	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, entity.UserStatusAvailable, got.Status)
}
