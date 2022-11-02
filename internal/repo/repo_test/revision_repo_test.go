package repo_test

import (
	"context"
	"encoding/json"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/repo/question"
	"github.com/answerdev/answer/internal/repo/revision"
	"github.com/answerdev/answer/internal/repo/unique"
	"github.com/stretchr/testify/assert"
	"testing"
)

var q = &entity.Question{
	ID:               "",
	UserID:           "1",
	Title:            "test",
	OriginalText:     "test",
	ParsedText:       "test",
	Status:           entity.QuestionStatusAvailable,
	ViewCount:        0,
	UniqueViewCount:  0,
	VoteCount:        0,
	AnswerCount:      0,
	CollectionCount:  0,
	FollowCount:      0,
	AcceptedAnswerID: "",
	LastAnswerID:     "",
	RevisionID:       "0",
}

func getRev(objID, title, content string) *entity.Revision {
	return &entity.Revision{
		ID:       "",
		UserID:   "1",
		ObjectID: objID,
		Title:    title,
		Content:  content,
		Log:      "add rev",
	}
}

func Test_revisionRepo_AddRevision(t *testing.T) {
	var (
		uniqueIDRepo = unique.NewUniqueIDRepo(testDataSource)
		revisionRepo = revision.NewRevisionRepo(testDataSource, uniqueIDRepo)
		questionRepo = question.NewQuestionRepo(testDataSource, uniqueIDRepo)
	)

	// create question
	err := questionRepo.AddQuestion(context.TODO(), q)
	assert.NoError(t, err)
	assert.NotEqual(t, "", q.ID)

	content, err := json.Marshal(q)
	// auto update false
	rev := getRev(q.ID, q.Title, string(content))
	err = revisionRepo.AddRevision(context.TODO(), rev, false)
	assert.NoError(t, err)
	qr, _, _ := questionRepo.GetQuestion(context.TODO(), q.ID)
	assert.NotEqual(t, rev.ID, qr.RevisionID)

	// auto update false
	rev = getRev(q.ID, q.Title, string(content))
	err = revisionRepo.AddRevision(context.TODO(), rev, true)
	assert.NoError(t, err)
	qr, _, _ = questionRepo.GetQuestion(context.TODO(), q.ID)
	assert.Equal(t, rev.ID, qr.RevisionID)

	// recovery
	t.Cleanup(func() {
		err = questionRepo.RemoveQuestion(context.TODO(), q.ID)
		assert.NoError(t, err)
	})
}

func Test_revisionRepo_GetLastRevisionByObjectID(t *testing.T) {
	var (
		uniqueIDRepo = unique.NewUniqueIDRepo(testDataSource)
		revisionRepo = revision.NewRevisionRepo(testDataSource, uniqueIDRepo)
	)

	Test_revisionRepo_AddRevision(t)
	rev, exists, err := revisionRepo.GetLastRevisionByObjectID(context.TODO(), q.ID)
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.NotNil(t, rev)
}

func Test_revisionRepo_GetRevisionList(t *testing.T) {
	var (
		uniqueIDRepo = unique.NewUniqueIDRepo(testDataSource)
		revisionRepo = revision.NewRevisionRepo(testDataSource, uniqueIDRepo)
	)
	Test_revisionRepo_AddRevision(t)
	revs, err := revisionRepo.GetRevisionList(context.TODO(), &entity.Revision{ObjectID: q.ID})
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(revs), 1)
}
