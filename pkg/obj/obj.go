package obj

import (
	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/pkg/converter"
	"github.com/segmentfault/pacman/errors"
)

// GetObjectTypeStrByObjectID get object key by object id
func GetObjectTypeStrByObjectID(objectID string) (objectTypeStr string, err error) {
	if err := checkObjectID(objectID); err != nil {
		return "", err
	}
	objectTypeNumber := converter.StringToInt(objectID[1:4])
	objectTypeStr, ok := constant.ObjectTypeNumberMapping[objectTypeNumber]
	if ok {
		return objectTypeStr, nil
	}
	return "", errors.BadRequest(reason.ObjectNotFound)
}

// GetObjectTypeNumberByObjectID get object type by object id
func GetObjectTypeNumberByObjectID(objectID string) (objectTypeNumber int, err error) {
	if err := checkObjectID(objectID); err != nil {
		return 0, err
	}
	return converter.StringToInt(objectID[1:4]), nil
}

func checkObjectID(objectID string) (err error) {
	if len(objectID) < 5 {
		return errors.BadRequest(reason.ObjectNotFound)
	}
	return nil
}
