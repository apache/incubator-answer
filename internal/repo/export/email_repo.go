package export

import (
	"context"
	"time"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/service/export"
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

// SetCode The email code is used to verify that the link in the message is out of date
func (e *emailRepo) SetCode(ctx context.Context, code, content string, duration time.Duration) error {
	err := e.data.Cache.SetString(ctx, code, content, duration)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// VerifyCode verify the code if out of date
func (e *emailRepo) VerifyCode(ctx context.Context, code string) (content string, err error) {
	content, err = e.data.Cache.GetString(ctx, code)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
