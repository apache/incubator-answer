package reason

import (
	"context"

	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/reason_common"
)

type ReasonService struct {
	reasonRepo reason_common.ReasonRepo
}

func NewReasonService(reasonRepo reason_common.ReasonRepo) *ReasonService {
	return &ReasonService{
		reasonRepo: reasonRepo,
	}
}

func (rs ReasonService) GetReasons(ctx context.Context, req schema.ReasonReq) (resp []*schema.ReasonItem, err error) {
	return rs.reasonRepo.ListReasons(ctx, req.ObjectType, req.Action)
}
