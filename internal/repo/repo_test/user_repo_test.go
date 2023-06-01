package repo_test

import (
	"context"
	"testing"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/repo/user"
	"github.com/stretchr/testify/assert"
)

func Test_userRepo_AddUser(t *testing.T) {
	userRepo := user.NewUserRepo(testDataSource, config_common.NewConfigRepo(testDataSource))
	userInfo := &entity.User{
		Username:    "answer",
		Pass:        "answer",
		EMail:       "answer@example.com",
		MailStatus:  entity.EmailStatusAvailable,
		Status:      entity.UserStatusAvailable,
		DisplayName: "answer",
		IsAdmin:     false,
	}
	err := userRepo.AddUser(context.TODO(), userInfo)
	assert.NoError(t, err)
}

func Test_userRepo_BatchGetByID(t *testing.T) {
	userRepo := user.NewUserRepo(testDataSource, config_common.NewConfigRepo(testDataSource))
	got, err := userRepo.BatchGetByID(context.TODO(), []string{"1"})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(got))
	assert.Equal(t, "admin", got[0].Username)
}

func Test_userRepo_GetByEmail(t *testing.T) {
	userRepo := user.NewUserRepo(testDataSource, config_common.NewConfigRepo(testDataSource))
	got, exist, err := userRepo.GetByEmail(context.TODO(), "admin@admin.com")
	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, "admin", got.Username)
}

func Test_userRepo_GetByUserID(t *testing.T) {
	userRepo := user.NewUserRepo(testDataSource, config_common.NewConfigRepo(testDataSource))
	got, exist, err := userRepo.GetByUserID(context.TODO(), "1")
	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, "admin", got.Username)
}

func Test_userRepo_GetByUsername(t *testing.T) {
	userRepo := user.NewUserRepo(testDataSource, config_common.NewConfigRepo(testDataSource))
	got, exist, err := userRepo.GetByUsername(context.TODO(), "admin")
	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, "admin", got.Username)
}

func Test_userRepo_IncreaseAnswerCount(t *testing.T) {
	userRepo := user.NewUserRepo(testDataSource, config_common.NewConfigRepo(testDataSource))
	err := userRepo.IncreaseAnswerCount(context.TODO(), "1", 1)
	assert.NoError(t, err)

	got, exist, err := userRepo.GetByUserID(context.TODO(), "1")
	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, 1, got.AnswerCount)
}

func Test_userRepo_IncreaseQuestionCount(t *testing.T) {
	userRepo := user.NewUserRepo(testDataSource, config_common.NewConfigRepo(testDataSource))
	err := userRepo.IncreaseQuestionCount(context.TODO(), "1", 1)
	assert.NoError(t, err)

	got, exist, err := userRepo.GetByUserID(context.TODO(), "1")
	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, 1, got.AnswerCount)
}

func Test_userRepo_UpdateEmail(t *testing.T) {
	userRepo := user.NewUserRepo(testDataSource, config_common.NewConfigRepo(testDataSource))
	err := userRepo.UpdateEmail(context.TODO(), "1", "admin@admin.com")
	assert.NoError(t, err)
}

func Test_userRepo_UpdateEmailStatus(t *testing.T) {
	userRepo := user.NewUserRepo(testDataSource, config_common.NewConfigRepo(testDataSource))
	err := userRepo.UpdateEmailStatus(context.TODO(), "1", entity.EmailStatusToBeVerified)
	assert.NoError(t, err)
}

func Test_userRepo_UpdateInfo(t *testing.T) {
	userRepo := user.NewUserRepo(testDataSource, config_common.NewConfigRepo(testDataSource))
	err := userRepo.UpdateInfo(context.TODO(), &entity.User{ID: "1", Bio: "test"})
	assert.NoError(t, err)

	got, exist, err := userRepo.GetByUserID(context.TODO(), "1")
	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, "test", got.Bio)
}

func Test_userRepo_UpdateLastLoginDate(t *testing.T) {
	userRepo := user.NewUserRepo(testDataSource, config_common.NewConfigRepo(testDataSource))
	err := userRepo.UpdateLastLoginDate(context.TODO(), "1")
	assert.NoError(t, err)
}

func Test_userRepo_UpdateNoticeStatus(t *testing.T) {
	userRepo := user.NewUserRepo(testDataSource, config_common.NewConfigRepo(testDataSource))
	err := userRepo.UpdateNoticeStatus(context.TODO(), "1", 1)
	assert.NoError(t, err)
}

func Test_userRepo_UpdatePass(t *testing.T) {
	userRepo := user.NewUserRepo(testDataSource, config_common.NewConfigRepo(testDataSource))
	err := userRepo.UpdatePass(context.TODO(), "1", "admin")
	assert.NoError(t, err)
}
