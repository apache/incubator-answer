package revision

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/entity"
	repo2 "github.com/segmentfault/answer/internal/repo"
	"github.com/segmentfault/answer/internal/repo/unique"
	"github.com/segmentfault/pacman/log"
	"github.com/stretchr/testify/assert"
)

var (
	dataSource *data.Data
	log        log.log
)

func Init() {
	s, _ := os.LookupEnv("TESTDATA-DB-CONNECTION")
	fmt.Println(s)
	cache, _, _ := data.NewCache(log.Getlog(), &data.CacheConf{})
	dataSource, _, _ = data.NewData(log.Getlog(), data.NewDB(true, &data.Database{
		Connection: s,
	}), cache)
	log = log.Getlog()
}

func TestRevisionRepo_AddRevision(t *testing.T) {
	Init()
	ctx := context.Background()
	uniqueIDRepo := unique.NewUniqueIDRepo(log, dataSource)
	questionRepo := repo2.NewQuestionRepo(log, dataSource, uniqueIDRepo)
	question, _, _ := questionRepo.GetQuestion(ctx, "10010000000000048")
	repo := NewRevisionRepo(log, dataSource, uniqueIDRepo)
	revision := &entity.Revision{
		UserID:     question.UserID,
		ObjectType: 0,
		ObjectID:   question.ID,
		Title:      question.Title,
		Content:    question.OriginalText,
		Status:     1,
	}
	err := repo.AddRevision(ctx, revision, true)
	assert.NoError(t, err)
}
