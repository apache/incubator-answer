package common

import (
	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/unique"
	"github.com/answerdev/answer/pkg/obj"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

type CommonRepo struct {
	data         *data.Data
	uniqueIDRepo unique.UniqueIDRepo
}

func NewCommonRepo(data *data.Data, uniqueIDRepo unique.UniqueIDRepo) *CommonRepo {
	return &CommonRepo{
		data:         data,
		uniqueIDRepo: uniqueIDRepo,
	}
}

// GetRootObjectID get root object ID
func (cr *CommonRepo) GetRootObjectID(objectID string) (rootObjectID string, err error) {
	var (
		exist      bool
		objectType string
		answer     = entity.Answer{}
		comment    = entity.Comment{}
	)

	objectType, err = obj.GetObjectTypeStrByObjectID(objectID)
	switch objectType {
	case "answer":
		exist, err = cr.data.DB.ID(objectID).Get(&answer)
		if !exist {
			err = errors.BadRequest(reason.ObjectNotFound)
		} else {
			objectID = answer.QuestionID
		}
	case "comment":
		exist, err = cr.data.DB.ID(objectID).Get(&comment)
		if !exist {
			err = errors.BadRequest(reason.ObjectNotFound)
		} else {
			objectID, err = cr.GetRootObjectID(comment.ObjectID)
		}
	default:
		rootObjectID = objectID
	}
	return
}

// GetObjectIDMap get object ID map from object id
func (cr *CommonRepo) GetObjectIDMap(objectID string) (objectIDMap map[string]string, err error) {
	var (
		exist bool
		ID,
		objectType string
		answer  = entity.Answer{}
		comment = entity.Comment{}
	)

	objectIDMap = map[string]string{}
	// 10070000000000450
	objectType, err = obj.GetObjectTypeStrByObjectID(objectID)
	if err != nil {
		log.Error("get report object type:", objectID, ",err:", err)
		return
	}
	switch objectType {
	case "answer":
		exist, err = cr.data.DB.ID(objectID).Get(&answer)
		if !exist {
			err = errors.BadRequest(reason.ObjectNotFound)
		} else {
			objectIDMap, err = cr.GetObjectIDMap(answer.QuestionID)
			ID = answer.ID
		}
	case "comment":
		exist, err = cr.data.DB.ID(objectID).Get(&comment)
		if !exist {
			err = errors.BadRequest(reason.ObjectNotFound)
		} else {
			objectIDMap, err = cr.GetObjectIDMap(comment.ObjectID)
			ID = comment.ID
		}
	case "question":
		ID = objectID
	}
	objectIDMap[objectType] = ID
	return
}
