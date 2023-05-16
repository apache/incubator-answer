package reason

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/answerdev/answer/internal/base/handler"
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

func (rr *reasonRepo) ListReasons(ctx context.Context, objectType, action string) (resp []*schema.ReasonItem, err error) {
	lang := handler.GetLangByCtx(ctx)
	reasonAction := fmt.Sprintf("%s.%s.reasons", objectType, action)
	resp = make([]*schema.ReasonItem, 0)

	reasonKeys, err := rr.configRepo.GetArrayString(reasonAction)
	if err != nil {
		return nil, err
	}
	for _, reasonKey := range reasonKeys {
		cfgValue, err := rr.configRepo.GetString(reasonKey)
		if err != nil {
			log.Error(err)
			continue
		}

		reason := &schema.ReasonItem{}
		err = json.Unmarshal([]byte(cfgValue), reason)
		if err != nil {
			log.Error(err)
			continue
		}
		reason.Translate(reasonKey+".", lang)

		reason.ReasonType, err = rr.configRepo.GetConfigType(reasonKey)
		if err != nil {
			log.Error(err)
			continue
		}

		resp = append(resp, reason)
	}
	return resp, nil
}
