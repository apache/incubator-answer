package repo_test

import (
	"context"
	"testing"

	"github.com/answerdev/answer/internal/repo/captcha"
	"github.com/stretchr/testify/assert"
)

var (
	ip         = "127.0.0.1"
	actionType = "actionType"
	amount     = 1
)

func Test_captchaRepo_DelActionType(t *testing.T) {
	captchaRepo := captcha.NewCaptchaRepo(testDataSource)
	err := captchaRepo.SetActionType(context.TODO(), ip, actionType, "", amount)
	assert.NoError(t, err)

	gotAmount, err := captchaRepo.GetActionType(context.TODO(), ip, actionType)
	assert.NoError(t, err)
	assert.Equal(t, amount, gotAmount)

	err = captchaRepo.DelActionType(context.TODO(), ip, actionType)
	assert.NoError(t, err)
}

func Test_captchaRepo_SetCaptcha(t *testing.T) {
	captchaRepo := captcha.NewCaptchaRepo(testDataSource)
	key, capt := "key", "1234"
	err := captchaRepo.SetCaptcha(context.TODO(), key, capt)
	assert.NoError(t, err)

	gotCaptcha, err := captchaRepo.GetCaptcha(context.TODO(), key)
	assert.NoError(t, err)
	assert.Equal(t, capt, gotCaptcha)
}
