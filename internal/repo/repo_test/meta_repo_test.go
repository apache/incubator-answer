package repo_test

import (
	"context"
	"testing"

	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/repo/meta"
	"github.com/stretchr/testify/assert"
)

func buildMetaEntity() *entity.Meta {
	return &entity.Meta{
		ObjectID: "1",
		Key:      "1",
		Value:    "1",
	}
}

func Test_metaRepo_GetMetaByObjectIdAndKey(t *testing.T) {
	metaRepo := meta.NewMetaRepo(testDataSource)
	metaEnt := buildMetaEntity()

	err := metaRepo.AddMeta(context.TODO(), metaEnt)
	assert.NoError(t, err)

	gotMeta, exist, err := metaRepo.GetMetaByObjectIdAndKey(context.TODO(), metaEnt.ObjectID, metaEnt.Key)
	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, metaEnt.ID, gotMeta.ID)

	err = metaRepo.RemoveMeta(context.TODO(), metaEnt.ID)
	assert.NoError(t, err)
}

func Test_metaRepo_GetMetaList(t *testing.T) {
	metaRepo := meta.NewMetaRepo(testDataSource)
	metaEnt := buildMetaEntity()

	err := metaRepo.AddMeta(context.TODO(), metaEnt)
	assert.NoError(t, err)

	gotMetaList, err := metaRepo.GetMetaList(context.TODO(), metaEnt)
	assert.NoError(t, err)
	assert.Equal(t, len(gotMetaList), 1)
	assert.Equal(t, gotMetaList[0].ID, metaEnt.ID)

	err = metaRepo.RemoveMeta(context.TODO(), metaEnt.ID)
	assert.NoError(t, err)
}

func Test_metaRepo_GetMetaPage(t *testing.T) {
	metaRepo := meta.NewMetaRepo(testDataSource)
	metaEnt := buildMetaEntity()

	err := metaRepo.AddMeta(context.TODO(), metaEnt)
	assert.NoError(t, err)

	gotMetaList, err := metaRepo.GetMetaList(context.TODO(), metaEnt)
	assert.NoError(t, err)
	assert.Equal(t, len(gotMetaList), 1)
	assert.Equal(t, gotMetaList[0].ID, metaEnt.ID)

	err = metaRepo.RemoveMeta(context.TODO(), metaEnt.ID)
	assert.NoError(t, err)
}

func Test_metaRepo_UpdateMeta(t *testing.T) {
	metaRepo := meta.NewMetaRepo(testDataSource)
	metaEnt := buildMetaEntity()

	err := metaRepo.AddMeta(context.TODO(), metaEnt)
	assert.NoError(t, err)

	metaEnt.Value = "testing"
	err = metaRepo.UpdateMeta(context.TODO(), metaEnt)
	assert.NoError(t, err)

	gotMeta, exist, err := metaRepo.GetMetaByObjectIdAndKey(context.TODO(), metaEnt.ObjectID, metaEnt.Key)
	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, gotMeta.Value, metaEnt.Value)

	err = metaRepo.RemoveMeta(context.TODO(), metaEnt.ID)
	assert.NoError(t, err)
}
