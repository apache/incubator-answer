package repo_test

import (
	"context"
	"testing"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/repo/auth"
	"github.com/stretchr/testify/assert"
)

var (
	token  = "token"
	userID = "1"
)

func Test_authRepo_SetUserCacheInfo(t *testing.T) {
	authRepo := auth.NewAuthRepo(testDataSource)

	err := authRepo.SetUserCacheInfo(context.TODO(), token, &entity.UserCacheInfo{UserID: userID})
	assert.NoError(t, err)

	cacheInfo, err := authRepo.GetUserCacheInfo(context.TODO(), token)
	assert.NoError(t, err)
	assert.Equal(t, userID, cacheInfo.UserID)
}

func Test_authRepo_RemoveUserCacheInfo(t *testing.T) {
	authRepo := auth.NewAuthRepo(testDataSource)

	err := authRepo.SetUserCacheInfo(context.TODO(), token, &entity.UserCacheInfo{UserID: userID})
	assert.NoError(t, err)

	err = authRepo.RemoveUserCacheInfo(context.TODO(), token)
	assert.NoError(t, err)

	_, err = authRepo.GetUserCacheInfo(context.TODO(), token)
	assert.Error(t, err)
}

func Test_authRepo_SetUserStatus(t *testing.T) {
	authRepo := auth.NewAuthRepo(testDataSource)

	err := authRepo.SetUserStatus(context.TODO(), userID, &entity.UserCacheInfo{UserID: userID})
	assert.NoError(t, err)

	cacheInfo, err := authRepo.GetUserStatus(context.TODO(), userID)
	assert.NoError(t, err)
	assert.Equal(t, userID, cacheInfo.UserID)
}
func Test_authRepo_RemoveUserStatus(t *testing.T) {
	authRepo := auth.NewAuthRepo(testDataSource)

	err := authRepo.SetUserStatus(context.TODO(), userID, &entity.UserCacheInfo{UserID: userID})
	assert.NoError(t, err)

	err = authRepo.RemoveUserStatus(context.TODO(), userID)
	assert.NoError(t, err)

	_, err = authRepo.GetUserStatus(context.TODO(), userID)
	assert.Error(t, err)
}

func Test_authRepo_SetBackyardUserCacheInfo(t *testing.T) {
	authRepo := auth.NewAuthRepo(testDataSource)

	err := authRepo.SetBackyardUserCacheInfo(context.TODO(), token, &entity.UserCacheInfo{UserID: userID})
	assert.NoError(t, err)

	cacheInfo, err := authRepo.GetBackyardUserCacheInfo(context.TODO(), token)
	assert.NoError(t, err)
	assert.Equal(t, userID, cacheInfo.UserID)
}

func Test_authRepo_RemoveBackyardUserCacheInfo(t *testing.T) {
	authRepo := auth.NewAuthRepo(testDataSource)

	err := authRepo.SetBackyardUserCacheInfo(context.TODO(), token, &entity.UserCacheInfo{UserID: userID})
	assert.NoError(t, err)

	err = authRepo.RemoveBackyardUserCacheInfo(context.TODO(), token)
	assert.NoError(t, err)

	_, err = authRepo.GetBackyardUserCacheInfo(context.TODO(), token)
	assert.Error(t, err)
}
