package repo_test

import (
	"context"
	"testing"

	"github.com/answerdev/answer/internal/repo/export"
	"github.com/stretchr/testify/assert"
)

func Test_emailRepo_VerifyCode(t *testing.T) {
	emailRepo := export.NewEmailRepo(testDataSource)
	code, content := "1111", "test"
	err := emailRepo.SetCode(context.TODO(), code, content)
	assert.NoError(t, err)

	verifyContent, err := emailRepo.VerifyCode(context.TODO(), code)
	assert.NoError(t, err)
	assert.Equal(t, content, verifyContent)
}
