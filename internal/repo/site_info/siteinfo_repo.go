package site_info

import (
	"context"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	"github.com/segmentfault/pacman/errors"
	"xorm.io/builder"
)

type siteInfoRepo struct {
	data *data.Data
}

func NewSiteInfo(data *data.Data) siteinfo_common.SiteInfoRepo {
	return &siteInfoRepo{
		data: data,
	}
}

// SaveByType save site setting by type
func (sr *siteInfoRepo) SaveByType(ctx context.Context, siteType string, data *entity.SiteInfo) (err error) {
	var (
		old   = &entity.SiteInfo{}
		exist bool
	)
	exist, _ = sr.data.DB.Where(builder.Eq{"type": siteType}).Get(old)
	if exist {
		_, err = sr.data.DB.ID(old.ID).Update(data)
		if err != nil {
			err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
		}
		return
	}

	_, err = sr.data.DB.Insert(data)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetByType get site info by type
func (sr *siteInfoRepo) GetByType(ctx context.Context, siteType string) (siteInfo *entity.SiteInfo, exist bool, err error) {
	siteInfo = &entity.SiteInfo{}
	exist, err = sr.data.DB.Where(builder.Eq{"type": siteType}).Get(siteInfo)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
