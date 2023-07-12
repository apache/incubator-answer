package repo_test

import (
	"context"
	"testing"

	"github.com/answerdev/answer/internal/repo/reason"
	"github.com/stretchr/testify/assert"
)

func Test_reasonRepo_ListReasons(t *testing.T) {
	configRepo := config_common.NewConfigRepo(testDataSource)
	reasonRepo := reason.NewReasonRepo(configRepo)
	reasonItems, err := reasonRepo.ListReasons(context.TODO(), "question", "close")
	assert.NoError(t, err)
	assert.Equal(t, 4, len(reasonItems))
}
