package reason

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/config"
	"github.com/answerdev/answer/internal/service/reason_common"
	"github.com/segmentfault/pacman/log"
)

type reasonRepo struct {
	configRepo config.ConfigRepo
}

func NewReasonRepo(configRepo config.ConfigRepo) reason_common.ReasonRepo {
	return &reasonRepo{
		configRepo: configRepo,
	}
}

func (rr *reasonRepo) ListReasons(ctx context.Context, objectType, action string) (resp []schema.ReasonItem, err error) {
	var (
		reasonAction = fmt.Sprintf("%s.%s.reasons", objectType, action)
		reasonKeys   []string
		cfgValue     string
	)
	resp = []schema.ReasonItem{}

	reasonKeys, err = rr.configRepo.GetArrayString(reasonAction)
	if err != nil {
		return
	}
	for _, reasonKey := range reasonKeys {
		var (
			reasonType int
			reason     = schema.ReasonItem{}
		)

		cfgValue, err = rr.configRepo.GetString(reasonKey)
		if err != nil {
			log.Error(err)
			continue
		}

		err = json.Unmarshal([]byte(cfgValue), &reason)
		if err != nil {
			log.Error(err)
			continue
		}
		reasonType, err = rr.configRepo.GetConfigType(reasonKey)
		if err != nil {
			log.Error(err)
			continue
		}

		reason.ReasonType = reasonType
		resp = append(resp, reason)
	}
	return
}
