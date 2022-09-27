package user

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentfault/answer/internal/base/data"
	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/service/config"
	usercommon "github.com/segmentfault/answer/internal/service/user_common"
	"github.com/segmentfault/pacman/errors"
)

// userRepo user repository
type userRepo struct {
	data       *data.Data
	configRepo config.ConfigRepo
}

// NewUserRepo new repository
func NewUserRepo(data *data.Data, configRepo config.ConfigRepo) usercommon.UserRepo {
	return &userRepo{
		data:       data,
		configRepo: configRepo,
	}
}

// AddUser add user
func (ur *userRepo) AddUser(ctx context.Context, user *entity.User) (err error) {
	_, err = ur.data.DB.Insert(user)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// IncreaseAnswerCount increase answer count
func (ur *userRepo) IncreaseAnswerCount(ctx context.Context, userID string, amount int) (err error) {
	user := &entity.User{}
	_, err = ur.data.DB.Where("id = ?", userID).Incr("answer_count", amount).Update(user)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// IncreaseQuestionCount increase question count
func (ur *userRepo) IncreaseQuestionCount(ctx context.Context, userID string, amount int) (err error) {
	user := &entity.User{}
	_, err = ur.data.DB.Where("id = ?", userID).Incr("question_count", amount).Update(user)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// UpdateLastLoginDate update last login date
func (ur *userRepo) UpdateLastLoginDate(ctx context.Context, userID string) (err error) {
	user := &entity.User{LastLoginDate: time.Now()}
	_, err = ur.data.DB.Where("id = ?", userID).Cols("last_login_date").Update(user)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

// UpdateEmailStatus update email status
func (ur *userRepo) UpdateEmailStatus(ctx context.Context, userID string, emailStatus int) error {
	cond := &entity.User{MailStatus: emailStatus}
	_, err := ur.data.DB.Where("id = ?", userID).Cols("mail_status").Update(cond)
	if err != nil {
		return err
	}
	return nil
}

// UpdateNoticeStatus update notice status
func (ur *userRepo) UpdateNoticeStatus(ctx context.Context, userID string, noticeStatus int) error {
	cond := &entity.User{NoticeStatus: noticeStatus}
	_, err := ur.data.DB.Where("id = ?", userID).Cols("notice_status").Update(cond)
	if err != nil {
		return errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return nil
}

func (ur *userRepo) UpdatePass(ctx context.Context, Data *entity.User) error {
	if Data.ID == "" {
		return fmt.Errorf("input error")
	}
	_, err := ur.data.DB.Where("id = ?", Data.ID).Cols("pass").Update(Data)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) UpdateEmail(ctx context.Context, userID, email string) (err error) {
	_, err = ur.data.DB.Where("id = ?", userID).Update(&entity.User{EMail: email})
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// UpdateInfo update user info
func (ur *userRepo) UpdateInfo(ctx context.Context, userInfo *entity.User) (err error) {
	_, err = ur.data.DB.Where("id = ?", userInfo.ID).
		Cols("display_name", "avatar", "bio", "bio_html", "website", "location").Update(userInfo)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetByUserID get user info by user id
func (ur *userRepo) GetByUserID(ctx context.Context, userID string) (userInfo *entity.User, exist bool, err error) {
	userInfo = &entity.User{}
	exist, err = ur.data.DB.Where("id = ?", userID).Get(userInfo)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

func (ur *userRepo) BatchGetByID(ctx context.Context, ids []string) ([]*entity.User, error) {
	list := make([]*entity.User, 0)
	err := ur.data.DB.In("id", ids).Find(&list)
	if err != nil {
		return nil, errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return list, nil
}

// GetByUsername get user by username
func (ur *userRepo) GetByUsername(ctx context.Context, username string) (userInfo *entity.User, exist bool, err error) {
	userInfo = &entity.User{}
	exist, err = ur.data.DB.Where("username = ?", username).Get(userInfo)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}

// GetByEmail get user by email
func (ur *userRepo) GetByEmail(ctx context.Context, email string) (userInfo *entity.User, exist bool, err error) {
	userInfo = &entity.User{}
	exist, err = ur.data.DB.Where("e_mail = ?", email).Get(userInfo)
	if err != nil {
		err = errors.InternalServer(reason.DatabaseError).WithError(err).WithStack()
	}
	return
}
