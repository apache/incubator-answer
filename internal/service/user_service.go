package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/Chain-Zhang/pinyin"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/activity"
	"github.com/answerdev/answer/internal/service/auth"
	"github.com/answerdev/answer/internal/service/export"
	"github.com/answerdev/answer/internal/service/service_config"
	usercommon "github.com/answerdev/answer/internal/service/user_common"
	"github.com/answerdev/answer/pkg/checker"
	"github.com/google/uuid"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
	"golang.org/x/crypto/bcrypt"
)

// UserRepo user repository

// UserService user service
type UserService struct {
	userRepo      usercommon.UserRepo
	userActivity  activity.UserActiveActivityRepo
	serviceConfig *service_config.ServiceConfig
	emailService  *export.EmailService
	authService   *auth.AuthService
}

func NewUserService(userRepo usercommon.UserRepo,
	userActivity activity.UserActiveActivityRepo,
	emailService *export.EmailService,
	authService *auth.AuthService,
	serviceConfig *service_config.ServiceConfig) *UserService {
	return &UserService{
		userRepo:      userRepo,
		userActivity:  userActivity,
		emailService:  emailService,
		serviceConfig: serviceConfig,
		authService:   authService,
	}
}

// GetUserInfoByUserID get user info by user id
func (us *UserService) GetUserInfoByUserID(ctx context.Context, token, userID string) (resp *schema.GetUserResp, err error) {
	userInfo, exist, err := us.userRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.BadRequest(reason.UserNotFound)
	}
	resp = &schema.GetUserResp{}
	resp.GetFromUserEntity(userInfo)
	resp.AccessToken = token
	return resp, nil
}

// GetUserStatus get user info by user id
func (us *UserService) GetUserStatus(ctx context.Context, userID, token string) (resp *schema.GetUserStatusResp, err error) {
	resp = &schema.GetUserStatusResp{}
	if len(userID) == 0 {
		return resp, nil
	}
	userInfo, exist, err := us.userRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.BadRequest(reason.UserNotFound)
	}

	userCacheInfo := &entity.UserCacheInfo{
		UserID:      userID,
		UserStatus:  userInfo.Status,
		EmailStatus: userInfo.MailStatus,
	}
	err = us.authService.UpdateUserCacheInfo(ctx, token, userCacheInfo)
	if err != nil {
		return nil, err
	}
	resp = &schema.GetUserStatusResp{
		Status: schema.UserStatusShow[userInfo.Status],
	}
	return resp, nil
}

func (us *UserService) GetOtherUserInfoByUsername(ctx context.Context, username string) (
	resp *schema.GetOtherUserInfoResp, err error) {
	userInfo, exist, err := us.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	resp = &schema.GetOtherUserInfoResp{}
	if !exist {
		return resp, nil
	}
	resp.Has = true
	resp.Info = &schema.GetOtherUserInfoByUsernameResp{}
	resp.Info.GetFromUserEntity(userInfo)
	return resp, nil
}

// EmailLogin email login
func (us *UserService) EmailLogin(ctx context.Context, req *schema.UserEmailLogin) (resp *schema.GetUserResp, err error) {
	userInfo, exist, err := us.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if !exist || userInfo.Status == entity.UserStatusDeleted {
		return nil, errors.BadRequest(reason.EmailOrPasswordWrong)
	}
	if !us.verifyPassword(ctx, req.Pass, userInfo.Pass) {
		return nil, errors.BadRequest(reason.EmailOrPasswordWrong)
	}

	err = us.userRepo.UpdateLastLoginDate(ctx, userInfo.ID)
	if err != nil {
		log.Error("UpdateLastLoginDate", err.Error())
	}

	resp = &schema.GetUserResp{}
	resp.GetFromUserEntity(userInfo)
	userCacheInfo := &entity.UserCacheInfo{
		UserID:      userInfo.ID,
		EmailStatus: userInfo.MailStatus,
		UserStatus:  userInfo.Status,
	}
	resp.AccessToken, err = us.authService.SetUserCacheInfo(ctx, userCacheInfo)
	if err != nil {
		return nil, err
	}
	resp.IsAdmin = userInfo.IsAdmin
	if resp.IsAdmin {
		err = us.authService.SetCmsUserCacheInfo(ctx, resp.AccessToken, userCacheInfo)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

// RetrievePassWord .
func (us *UserService) RetrievePassWord(ctx context.Context, req *schema.UserRetrievePassWordRequest) (string, error) {
	userInfo, has, err := us.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}
	if !has {
		return "", errors.BadRequest(reason.UserNotFound)
	}

	// send email
	data := &schema.EmailCodeContent{
		Email:  req.Email,
		UserID: userInfo.ID,
	}
	code := uuid.NewString()
	verifyEmailUrl := fmt.Sprintf("%s/users/password-reset?code=%s", us.serviceConfig.WebHost, code)
	title, body, err := us.emailService.PassResetTemplate(ctx, verifyEmailUrl)
	if err != nil {
		return "", err
	}
	go us.emailService.Send(ctx, req.Email, title, body, code, data.ToJSONString())
	return code, nil
}

// UseRePassWord
func (us *UserService) UseRePassWord(ctx context.Context, req *schema.UserRePassWordRequest) (resp *schema.GetUserResp, err error) {
	data := &schema.EmailCodeContent{}
	err = data.FromJSONString(req.Content)
	if err != nil {
		return nil, errors.BadRequest(reason.EmailVerifyUrlExpired)
	}

	userInfo, exist, err := us.userRepo.GetByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.BadRequest(reason.UserNotFound)
	}
	enpass, err := us.encryptPassword(ctx, req.Pass)
	if err != nil {
		return nil, err
	}
	userInfo.Pass = enpass
	err = us.userRepo.UpdatePass(ctx, userInfo)
	if err != nil {
		return nil, err
	}
	resp = &schema.GetUserResp{}
	return resp, nil
}

func (us *UserService) UserModifyPassWordVerification(ctx context.Context, request *schema.UserModifyPassWordRequest) (bool, error) {

	userInfo, has, err := us.userRepo.GetByUserID(ctx, request.UserId)
	if err != nil {
		return false, err
	}
	if !has {
		return false, fmt.Errorf("user does not exist")
	}
	isPass := us.verifyPassword(ctx, request.OldPass, userInfo.Pass)
	if !isPass {
		return false, nil
	}

	return true, nil
}

// UserModifyPassWord
func (us *UserService) UserModifyPassWord(ctx context.Context, request *schema.UserModifyPassWordRequest) error {
	enpass, err := us.encryptPassword(ctx, request.Pass)
	if err != nil {
		return err
	}
	userInfo, has, err := us.userRepo.GetByUserID(ctx, request.UserId)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("user does not exist")
	}
	isPass := us.verifyPassword(ctx, request.OldPass, userInfo.Pass)
	if !isPass {
		return fmt.Errorf("the old password verification failed")
	}
	userInfo.Pass = enpass
	err = us.userRepo.UpdatePass(ctx, userInfo)
	if err != nil {
		return err
	}
	return nil
}

// UpdateInfo update user info
func (us *UserService) UpdateInfo(ctx context.Context, req *schema.UpdateInfoRequest) (err error) {
	if len(req.Username) > 0 {
		userInfo, exist, err := us.userRepo.GetByUsername(ctx, req.Username)
		if err != nil {
			return err
		}
		if exist && userInfo.ID != req.UserId {
			return errors.BadRequest(reason.UsernameDuplicate)
		}
	}

	userInfo := entity.User{}
	userInfo.ID = req.UserId
	userInfo.Avatar = req.Avatar
	userInfo.DisplayName = req.DisplayName
	userInfo.Bio = req.Bio
	userInfo.BioHtml = req.BioHtml
	userInfo.Location = req.Location
	userInfo.Website = req.Website
	userInfo.Username = req.Username
	if err := us.userRepo.UpdateInfo(ctx, &userInfo); err != nil {
		return err
	}
	return nil
}

func (us *UserService) UserEmailHas(ctx context.Context, email string) (bool, error) {
	_, has, err := us.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	return has, nil
}

// UserRegisterByEmail user register
func (us *UserService) UserRegisterByEmail(ctx context.Context, registerUserInfo *schema.UserRegisterReq) (
	resp *schema.GetUserResp, err error) {
	_, has, err := us.userRepo.GetByEmail(ctx, registerUserInfo.Email)
	if err != nil {
		return nil, err
	}
	if has {
		return nil, errors.BadRequest(reason.EmailDuplicate)
	}

	userInfo := &entity.User{}
	userInfo.EMail = registerUserInfo.Email
	userInfo.DisplayName = registerUserInfo.Name
	userInfo.Pass, err = us.encryptPassword(ctx, registerUserInfo.Pass)
	if err != nil {
		return nil, err
	}
	userInfo.Username, err = us.makeUsername(ctx, registerUserInfo.Name)
	if err != nil {
		return nil, err
	}
	userInfo.IPInfo = registerUserInfo.IP
	userInfo.MailStatus = entity.EmailStatusToBeVerified
	userInfo.Status = entity.UserStatusAvailable
	err = us.userRepo.AddUser(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	// send email
	data := &schema.EmailCodeContent{
		Email:  registerUserInfo.Email,
		UserID: userInfo.ID,
	}
	code := uuid.NewString()
	verifyEmailUrl := fmt.Sprintf("%s/users/account-activation?code=%s", us.serviceConfig.WebHost, code)
	title, body, err := us.emailService.RegisterTemplate(ctx, verifyEmailUrl)
	if err != nil {
		return nil, err
	}
	go us.emailService.Send(ctx, userInfo.EMail, title, body, code, data.ToJSONString())

	// return user info and token
	resp = &schema.GetUserResp{}
	resp.GetFromUserEntity(userInfo)
	userCacheInfo := &entity.UserCacheInfo{
		UserID:      userInfo.ID,
		EmailStatus: userInfo.MailStatus,
		UserStatus:  userInfo.Status,
	}
	resp.AccessToken, err = us.authService.SetUserCacheInfo(ctx, userCacheInfo)
	if err != nil {
		return nil, err
	}
	resp.IsAdmin = userInfo.IsAdmin
	if resp.IsAdmin {
		err = us.authService.SetCmsUserCacheInfo(ctx, resp.AccessToken, &entity.UserCacheInfo{UserID: userInfo.ID})
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

func (us *UserService) UserVerifyEmailSend(ctx context.Context, userID string) error {
	userInfo, has, err := us.userRepo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if !has {
		return errors.BadRequest(reason.UserNotFound)
	}

	data := &schema.EmailCodeContent{
		Email:  userInfo.EMail,
		UserID: userInfo.ID,
	}
	code := uuid.NewString()
	verifyEmailUrl := fmt.Sprintf("%s/users/account-activation?code=%s", us.serviceConfig.WebHost, code)
	title, body, err := us.emailService.RegisterTemplate(ctx, verifyEmailUrl)
	if err != nil {
		return err
	}
	go us.emailService.Send(ctx, userInfo.EMail, title, body, code, data.ToJSONString())
	return nil
}

func (us *UserService) UserNoticeSet(ctx context.Context, userId string, noticeSwitch bool) (
	resp *schema.UserNoticeSetResp, err error) {
	userInfo, has, err := us.userRepo.GetByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.BadRequest(reason.UserNotFound)
	}
	if noticeSwitch {
		userInfo.NoticeStatus = schema.Notice_Status_On
	} else {
		userInfo.NoticeStatus = schema.Notice_Status_Off
	}
	err = us.userRepo.UpdateNoticeStatus(ctx, userInfo.ID, userInfo.NoticeStatus)
	return &schema.UserNoticeSetResp{NoticeSwitch: noticeSwitch}, err
}

func (us *UserService) UserVerifyEmail(ctx context.Context, req *schema.UserVerifyEmailReq) (resp *schema.GetUserResp, err error) {
	data := &schema.EmailCodeContent{}
	err = data.FromJSONString(req.Content)
	if err != nil {
		return nil, errors.BadRequest(reason.EmailVerifyUrlExpired)
	}

	userInfo, has, err := us.userRepo.GetByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.BadRequest(reason.UserNotFound)
	}
	userInfo.MailStatus = entity.EmailStatusAvailable
	err = us.userRepo.UpdateEmailStatus(ctx, userInfo.ID, userInfo.MailStatus)
	if err != nil {
		return nil, err
	}
	if err = us.userActivity.UserActive(ctx, userInfo.ID); err != nil {
		log.Error(err)
	}

	resp = &schema.GetUserResp{}
	resp.GetFromUserEntity(userInfo)
	userCacheInfo := &entity.UserCacheInfo{
		UserID:      userInfo.ID,
		EmailStatus: userInfo.MailStatus,
		UserStatus:  userInfo.Status,
	}
	resp.AccessToken, err = us.authService.SetUserCacheInfo(ctx, userCacheInfo)
	if err != nil {
		return nil, err
	}
	resp.IsAdmin = userInfo.IsAdmin
	if resp.IsAdmin {
		err = us.authService.SetCmsUserCacheInfo(ctx, resp.AccessToken, &entity.UserCacheInfo{UserID: userInfo.ID})
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}

// makeUsername
// Generate a unique Username based on the displayName
func (us *UserService) makeUsername(ctx context.Context, displayName string) (username string, err error) {
	// Chinese processing
	if has := checker.IsChinese(displayName); has {
		str, err := pinyin.New(displayName).Split("").Mode(pinyin.WithoutTone).Convert()
		if err != nil {
			return "", err
		} else {
			displayName = str
		}
	}

	username = strings.ReplaceAll(displayName, " ", "_")
	username = strings.ToLower(username)
	suffix := ""

	re := regexp.MustCompile(`^[a-z0-9._-]{4,30}$`)
	match := re.MatchString(username)
	if !match {
		return "", errors.BadRequest(reason.UsernameInvalid)
	}

	for {
		_, has, err := us.userRepo.GetByUsername(ctx, username+suffix)
		if err != nil {
			return "", err
		}
		if !has {
			break
		}
		bytes := make([]byte, 2)
		_, _ = rand.Read(bytes)
		suffix = hex.EncodeToString(bytes)
	}
	return username + suffix, nil
}

// verifyPassword
// Compare whether the password is correct
func (us *UserService) verifyPassword(ctx context.Context, LoginPass, UserPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(UserPass), []byte(LoginPass))
	if err != nil {
		return false
	}
	return true
}

// encryptPassword
// The password does irreversible encryption.
func (us *UserService) encryptPassword(ctx context.Context, Pass string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(Pass), bcrypt.DefaultCost)
	//This encrypted string can be saved to the database and can be used as password matching verification
	return string(hashPwd), err
}

// UserChangeEmailSendCode user change email verification
func (us *UserService) UserChangeEmailSendCode(ctx context.Context, req *schema.UserChangeEmailSendCodeReq) error {
	_, exist, err := us.userRepo.GetByUserID(ctx, req.UserID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.BadRequest(reason.UserNotFound)
	}

	_, exist, err = us.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if exist {
		return errors.BadRequest(reason.EmailDuplicate)
	}

	data := &schema.EmailCodeContent{
		Email:  req.Email,
		UserID: req.UserID,
	}
	code := uuid.NewString()
	verifyEmailUrl := fmt.Sprintf("%s/users/confirm-new-email?code=%s", us.serviceConfig.WebHost, code)
	title, body, err := us.emailService.ChangeEmailTemplate(ctx, verifyEmailUrl)
	if err != nil {
		return err
	}
	log.Infof("send email confirmation %s", verifyEmailUrl)

	go us.emailService.Send(context.Background(), req.Email, title, body, code, data.ToJSONString())
	return nil
}

// UserChangeEmailVerify user change email verify code
func (us *UserService) UserChangeEmailVerify(ctx context.Context, content string) (err error) {
	data := &schema.EmailCodeContent{}
	err = data.FromJSONString(content)
	if err != nil {
		return errors.BadRequest(reason.EmailVerifyUrlExpired)
	}

	_, exist, err := us.userRepo.GetByEmail(ctx, data.Email)
	if err != nil {
		return err
	}
	if exist {
		return errors.BadRequest(reason.EmailDuplicate)
	}

	_, exist, err = us.userRepo.GetByUserID(ctx, data.UserID)
	if err != nil {
		return err
	}
	if !exist {
		return errors.BadRequest(reason.UserNotFound)
	}
	err = us.userRepo.UpdateEmail(ctx, data.UserID, data.Email)
	if err != nil {
		return err
	}
	us.userRepo.UpdateEmailStatus(ctx, data.UserID, entity.EmailStatusAvailable)
	if err != nil {
		return err
	}
	return nil
}
