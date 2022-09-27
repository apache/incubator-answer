package usercommon

import (
	"context"

	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
)

type UserRepo interface {
	AddUser(ctx context.Context, user *entity.User) (err error)
	IncreaseAnswerCount(ctx context.Context, userID string, amount int) (err error)
	IncreaseQuestionCount(ctx context.Context, userID string, amount int) (err error)
	UpdateLastLoginDate(ctx context.Context, userID string) (err error)
	UpdateEmailStatus(ctx context.Context, userID string, emailStatus int) error
	UpdateNoticeStatus(ctx context.Context, userID string, noticeStatus int) error
	UpdateEmail(ctx context.Context, userID, email string) error
	UpdatePass(ctx context.Context, Data *entity.User) error
	UpdateInfo(ctx context.Context, userInfo *entity.User) (err error)
	GetByUserID(ctx context.Context, userID string) (userInfo *entity.User, exist bool, err error)
	BatchGetByID(ctx context.Context, ids []string) ([]*entity.User, error)
	GetByUsername(ctx context.Context, username string) (userInfo *entity.User, exist bool, err error)
	GetByEmail(ctx context.Context, email string) (userInfo *entity.User, exist bool, err error)
}

// UserCommon user service
type UserCommon struct {
	userRepo UserRepo
}

func NewUserCommon(userRepo UserRepo) *UserCommon {
	return &UserCommon{
		userRepo: userRepo,
	}
}

func (us *UserCommon) GetUserBasicInfoByID(ctx context.Context, ID string) (*schema.UserBasicInfo, bool, error) {
	dbInfo, has, err := us.userRepo.GetByUserID(ctx, ID)
	if err != nil {
		return nil, has, err
	}
	info := us.UserBasicInfoFormat(ctx, dbInfo)
	return info, has, nil
}

func (us *UserCommon) GetUserBasicInfoByUserName(ctx context.Context, username string) (*schema.UserBasicInfo, bool, error) {
	userInfo, exist, err := us.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, exist, err
	}
	info := us.UserBasicInfoFormat(ctx, userInfo)
	return info, exist, nil
}

func (us *UserCommon) UpdateAnswerCount(ctx context.Context, userID string, num int) error {
	return us.userRepo.IncreaseAnswerCount(ctx, userID, num)
}

func (us *UserCommon) UpdateQuestionCount(ctx context.Context, userID string, num int) error {
	return us.userRepo.IncreaseQuestionCount(ctx, userID, num)
}

func (us *UserCommon) BatchUserBasicInfoByID(ctx context.Context, IDs []string) (map[string]*schema.UserBasicInfo, error) {
	userMap := make(map[string]*schema.UserBasicInfo)
	dbInfo, err := us.userRepo.BatchGetByID(ctx, IDs)
	if err != nil {
		return userMap, err
	}
	for _, item := range dbInfo {
		info := us.UserBasicInfoFormat(ctx, item)
		userMap[item.ID] = info
	}
	return userMap, nil
}

// UserBasicInfoFormat
func (us *UserCommon) UserBasicInfoFormat(ctx context.Context, dbinfo *entity.User) *schema.UserBasicInfo {
	info := new(schema.UserBasicInfo)
	info.UserId = dbinfo.ID
	info.UserName = dbinfo.Username
	info.Rank = dbinfo.Rank
	info.DisplayName = dbinfo.DisplayName
	info.Avatar = dbinfo.Avatar
	info.Website = dbinfo.Website
	info.Location = dbinfo.Location
	info.IpInfo = dbinfo.IPInfo
	info.Status = dbinfo.Status
	return info
}
