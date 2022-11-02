package repo_test

import (
	"context"
	"sync"
	"testing"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/repo/tag"
	"github.com/stretchr/testify/assert"
)

var (
	tagRelOnce     sync.Once
	testTagRelList = []*entity.TagRel{
		{
			ObjectID: "1",
			TagID:    "1",
			Status:   entity.TagRelStatusAvailable,
		},
		{
			ObjectID: "2",
			TagID:    "2",
			Status:   entity.TagRelStatusAvailable,
		},
	}
)

func addTagRelList() {
	tagRelRepo := tag.NewTagRelRepo(testDataSource)
	err := tagRelRepo.AddTagRelList(context.TODO(), testTagRelList)
	if err != nil {
		panic(err)
	}
}

func Test_tagListRepo_BatchGetObjectTagRelList(t *testing.T) {
	tagRelOnce.Do(addTagRelList)
	tagRelRepo := tag.NewTagRelRepo(testDataSource)
	relList, err :=
		tagRelRepo.BatchGetObjectTagRelList(context.TODO(), []string{testTagRelList[0].ObjectID, testTagRelList[1].ObjectID})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(relList))
}

func Test_tagListRepo_CountTagRelByTagID(t *testing.T) {
	tagRelOnce.Do(addTagRelList)
	tagRelRepo := tag.NewTagRelRepo(testDataSource)
	count, err := tagRelRepo.CountTagRelByTagID(context.TODO(), "1")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}

func Test_tagListRepo_GetObjectTagRelList(t *testing.T) {
	tagRelOnce.Do(addTagRelList)
	tagRelRepo := tag.NewTagRelRepo(testDataSource)

	relList, err :=
		tagRelRepo.GetObjectTagRelList(context.TODO(), testTagRelList[0].ObjectID)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(relList))
}

func Test_tagListRepo_GetObjectTagRelWithoutStatus(t *testing.T) {
	tagRelOnce.Do(addTagRelList)
	tagRelRepo := tag.NewTagRelRepo(testDataSource)

	relList, err :=
		tagRelRepo.BatchGetObjectTagRelList(context.TODO(), []string{testTagRelList[0].ObjectID, testTagRelList[1].ObjectID})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(relList))

	ids := []int64{relList[0].ID, relList[1].ID}
	err = tagRelRepo.RemoveTagRelListByIDs(context.TODO(), ids)
	assert.NoError(t, err)

	count, err := tagRelRepo.CountTagRelByTagID(context.TODO(), "1")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), count)

	_, exist, err := tagRelRepo.GetObjectTagRelWithoutStatus(context.TODO(), relList[0].ObjectID, relList[0].TagID)
	assert.NoError(t, err)
	assert.True(t, exist)

	err = tagRelRepo.EnableTagRelByIDs(context.TODO(), ids)
	assert.NoError(t, err)

	count, err = tagRelRepo.CountTagRelByTagID(context.TODO(), "1")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}
