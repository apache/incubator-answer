package user

import (
	"context"
	"encoding/json"
	"time"

	"github.com/answerdev/answer/internal/base/data"
	"github.com/answerdev/answer/internal/base/pager"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/service/auth"
	"github.com/answerdev/answer/internal/service/user_backyard"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// userBackyardRepo user repository
type userBackyardRepo struct {
	data     *data.Data
	authRepo auth.AuthRepo
}

// NewUserBackyardRepo new repository
func NewUserBackyardRepo(data *data.Data, authRepo auth.AuthRepo) user_backyard.UserBackyardRepo {
	return &userBackyardRepo{
		data:     data,
		authRepo: authRepo,
	}
}

// UpdateUserStatus update user status
func (ur *userBackyardRepo) UpdateUserStatus(ctx context.Context, userID string, userStatus, mailStatus int,
	email string) (err error) {
	cond := &entity.User{Status: userStatus, MailStatus: mailStatus, EMail: email}
	switch userStatus {
	case entity.UserStatusSuspended:
		cond.SuspendedAt = time.Now()
	case entity.UserStatusDeleted:
		cond.DeletedAt = time.Now()
	}
	_, err = ur.data.DB.ID(userID).Update(cond)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}

	userCacheInfo := &entity.UserCacheInfo{
		UserID:      userID,
		EmailStatus: mailStatus,
		UserStatus:  userStatus,
	}
	t, _ := json.Marshal(userCacheInfo)
	log.Infof("user change status: %s", string(t))
	err = ur.authRepo.SetUserStatus(ctx, userID, userCacheInfo)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetUserInfo get user info
func (ur *userBackyardRepo) GetUserInfo(ctx context.Context, userID string) (user *entity.User, exist bool, err error) {
	user = &entity.User{}
	exist, err = ur.data.DB.ID(userID).Get(user)
	if err != nil {
		return nil, false, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetUserPage get user page
func (ur *userBackyardRepo) GetUserPage(ctx context.Context, page, pageSize int, user *entity.User) (users []*entity.User, total int64, err error) {
	users = make([]*entity.User, 0)
	session := ur.data.DB.NewSession()
	if user.Status == entity.UserStatusDeleted {
		session.Desc("deleted_at")
	} else if user.Status == entity.UserStatusSuspended {
		session.Desc("suspended_at")
	} else {
		session.Desc("created_at")
	}
	total, err = pager.Help(page, pageSize, &users, user, session)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
