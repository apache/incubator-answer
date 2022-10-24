package siteinfo_common

import (
	"context"
	"github.com/answerdev/answer/internal/entity"
)

type SiteInfoRepo interface {
	SaveByType(ctx context.Context, siteType string, data *entity.SiteInfo) (err error)
	GetByType(ctx context.Context, siteType string) (siteInfo *entity.SiteInfo, exist bool, err error)
}
