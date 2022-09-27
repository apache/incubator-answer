package export

import (
	"context"
	"time"

	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/service/export"
	"github.com/segmentfault/pacman/errors"
)

// emailRepo email repository
type emailRepo struct {
	data *data.Data
}

// NewEmailRepo new repository
func NewEmailRepo(data *data.Data) export.EmailRepo {
	return &emailRepo{
		data: data,
	}
}

func (e *emailRepo) SetCode(ctx context.Context, code, content string) error {
	err := e.data.Cache.SetString(ctx, code, content, 10*time.Minute)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (e *emailRepo) VerifyCode(ctx context.Context, code string) (content string, err error) {
	content, err = e.data.Cache.GetString(ctx, code)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
