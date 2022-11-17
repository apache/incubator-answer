package controller

import (
	"net/http"
	"path"
	"strings"

	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/middleware"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service"
	"github.com/answerdev/answer/internal/service/action"
	"github.com/answerdev/answer/internal/service/auth"
	"github.com/answerdev/answer/internal/service/export"
	"github.com/answerdev/answer/internal/service/uploader"
	"github.com/gin-gonic/gin"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// UserController user controller
type UserController struct {
	userService     *service.UserService
	authService     *auth.AuthService
	actionService   *action.CaptchaService
	uploaderService *uploader.UploaderService
	emailService    *export.EmailService
}

// NewUserController new controller
func NewUserController(
	authService *auth.AuthService,
	userService *service.UserService,
	actionService *action.CaptchaService,
	emailService *export.EmailService,
	uploaderService *uploader.UploaderService,
) *UserController {
	return &UserController{
		authService:     authService,
		userService:     userService,
		actionService:   actionService,
		uploaderService: uploaderService,
		emailService:    emailService,
	}
}

// GetUserInfoByUserID get user info, if user no login response http code is 200, but user info is null
// @Summary GetUserInfoByUserID
// @Description get user info, if user no login response http code is 200, but user info is null
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} handler.RespBody{data=schema.GetUserToSetShowResp}
// @Router /answer/api/v1/user/info [get]
func (uc *UserController) GetUserInfoByUserID(ctx *gin.Context) {
	userID := middleware.GetLoginUserIDFromContext(ctx)
	token := middleware.ExtractToken(ctx)

	// if user is no login return null in data
	if len(token) == 0 || len(userID) == 0 {
		handler.HandleResponse(ctx, nil, nil)
		return
	}

	resp, err := uc.userService.GetUserInfoByUserID(ctx, token, userID)
	handler.HandleResponse(ctx, err, resp)
}

// GetOtherUserInfoByUsername godoc
// @Summary GetOtherUserInfoByUsername
// @Description GetOtherUserInfoByUsername
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param username query string true "username"
// @Success 200 {object} handler.RespBody{data=schema.GetOtherUserInfoResp}
// @Router /answer/api/v1/personal/user/info [get]
func (uc *UserController) GetOtherUserInfoByUsername(ctx *gin.Context) {
	req := &schema.GetOtherUserInfoByUsernameReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := uc.userService.GetOtherUserInfoByUsername(ctx, req.Username)
	handler.HandleResponse(ctx, err, resp)
}

// UserEmailLogin godoc
// @Summary UserEmailLogin
// @Description UserEmailLogin
// @Tags User
// @Accept json
// @Produce json
// @Param data body schema.UserEmailLogin true "UserEmailLogin"
// @Success 200 {object} handler.RespBody{data=schema.GetUserResp}
// @Router /answer/api/v1/user/login/email [post]
func (uc *UserController) UserEmailLogin(ctx *gin.Context) {
	req := &schema.UserEmailLogin{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	captchaPass := uc.actionService.ActionRecordVerifyCaptcha(ctx, schema.ActionRecordTypeLogin, ctx.ClientIP(), req.CaptchaID, req.CaptchaCode)
	if !captchaPass {
		resp := schema.UserVerifyEmailErrorResponse{
			Key:   "captcha_code",
			Value: "error.object.verification_failed",
		}
		resp.Value = translator.GlobalTrans.Tr(handler.GetLang(ctx), resp.Value)
		handler.HandleResponse(ctx, errors.BadRequest(reason.CaptchaVerificationFailed), resp)
		return
	}

	resp, err := uc.userService.EmailLogin(ctx, req)
	if err != nil {
		_, _ = uc.actionService.ActionRecordAdd(ctx, schema.ActionRecordTypeLogin, ctx.ClientIP())
		resp := schema.UserVerifyEmailErrorResponse{
			Key:   "e_mail",
			Value: reason.EmailOrPasswordWrong,
		}
		resp.Value = translator.GlobalTrans.Tr(handler.GetLang(ctx), resp.Value)
		handler.HandleResponse(ctx, errors.BadRequest(reason.EmailOrPasswordWrong), resp)
		return
	}
	uc.actionService.ActionRecordDel(ctx, schema.ActionRecordTypeLogin, ctx.ClientIP())
	handler.HandleResponse(ctx, nil, resp)
}

// RetrievePassWord godoc
// @Summary RetrievePassWord
// @Description RetrievePassWord
// @Tags User
// @Accept  json
// @Produce  json
// @Param data body schema.UserRetrievePassWordRequest  true "UserRetrievePassWordRequest"
// @Success 200 {string} string ""
// @Router /answer/api/v1/user/password/reset [post]
func (uc *UserController) RetrievePassWord(ctx *gin.Context) {
	req := &schema.UserRetrievePassWordRequest{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	captchaPass := uc.actionService.ActionRecordVerifyCaptcha(ctx, schema.ActionRecordTypeFindPass, ctx.ClientIP(), req.CaptchaID, req.CaptchaCode)
	if !captchaPass {
		resp := schema.UserVerifyEmailErrorResponse{
			Key:   "captcha_code",
			Value: "error.object.verification_failed",
		}
		resp.Value = translator.GlobalTrans.Tr(handler.GetLang(ctx), resp.Value)
		handler.HandleResponse(ctx, errors.BadRequest(reason.CaptchaVerificationFailed), resp)
		return
	}
	_, _ = uc.actionService.ActionRecordAdd(ctx, schema.ActionRecordTypeFindPass, ctx.ClientIP())
	code, err := uc.userService.RetrievePassWord(ctx, req)
	handler.HandleResponse(ctx, err, code)
}

// UseRePassWord godoc
// @Summary UseRePassWord
// @Description UseRePassWord
// @Tags User
// @Accept  json
// @Produce  json
// @Param data body schema.UserRePassWordRequest  true "UserRePassWordRequest"
// @Success 200 {string} string ""
// @Router /answer/api/v1/user/password/replacement [post]
func (uc *UserController) UseRePassWord(ctx *gin.Context) {
	req := &schema.UserRePassWordRequest{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.Content = uc.emailService.VerifyUrlExpired(ctx, req.Code)
	if len(req.Content) == 0 {
		handler.HandleResponse(ctx, errors.Forbidden(reason.EmailVerifyURLExpired),
			&schema.ForbiddenResp{Type: schema.ForbiddenReasonTypeURLExpired})
		return
	}

	resp, err := uc.userService.UseRePassword(ctx, req)
	uc.actionService.ActionRecordDel(ctx, schema.ActionRecordTypeFindPass, ctx.ClientIP())
	handler.HandleResponse(ctx, err, resp)
}

// UserLogout user logout
// @Summary user logout
// @Description user logout
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/user/logout [get]
func (uc *UserController) UserLogout(ctx *gin.Context) {
	accessToken := middleware.ExtractToken(ctx)
	_ = uc.authService.RemoveUserCacheInfo(ctx, accessToken)
	handler.HandleResponse(ctx, nil, nil)
}

// UserRegisterByEmail godoc
// @Summary UserRegisterByEmail
// @Description UserRegisterByEmail
// @Tags User
// @Accept json
// @Produce json
// @Param data body schema.UserRegisterReq true "UserRegisterReq"
// @Success 200 {object} handler.RespBody{data=schema.GetUserResp}
// @Router /answer/api/v1/user/register/email [post]
func (uc *UserController) UserRegisterByEmail(ctx *gin.Context) {
	req := &schema.UserRegisterReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.IP = ctx.ClientIP()

	resp, err := uc.userService.UserRegisterByEmail(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}

// UserVerifyEmail godoc
// @Summary UserVerifyEmail
// @Description UserVerifyEmail
// @Tags User
// @Accept json
// @Produce json
// @Param code query string true "code" default()
// @Success 200 {object} handler.RespBody{data=schema.GetUserResp}
// @Router /answer/api/v1/user/email/verification [post]
func (uc *UserController) UserVerifyEmail(ctx *gin.Context) {
	req := &schema.UserVerifyEmailReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.Content = uc.emailService.VerifyUrlExpired(ctx, req.Code)
	if len(req.Content) == 0 {
		handler.HandleResponse(ctx, errors.Forbidden(reason.EmailVerifyURLExpired),
			&schema.ForbiddenResp{Type: schema.ForbiddenReasonTypeURLExpired})
		return
	}

	resp, err := uc.userService.UserVerifyEmail(ctx, req)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}

	uc.actionService.ActionRecordDel(ctx, schema.ActionRecordTypeEmail, ctx.ClientIP())
	handler.HandleResponse(ctx, err, resp)
}

// UserVerifyEmailSend godoc
// @Summary UserVerifyEmailSend
// @Description UserVerifyEmailSend
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param captcha_id query string false "captcha_id"  default()
// @Param captcha_code query string false "captcha_code"  default()
// @Success 200 {string} string ""
// @Router /answer/api/v1/user/email/verification/send [post]
func (uc *UserController) UserVerifyEmailSend(ctx *gin.Context) {
	req := &schema.UserVerifyEmailSendReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	userInfo := middleware.GetUserInfoFromContext(ctx)
	if userInfo == nil {
		handler.HandleResponse(ctx, errors.Unauthorized(reason.UnauthorizedError), nil)
		return
	}

	captchaPass := uc.actionService.ActionRecordVerifyCaptcha(ctx, schema.ActionRecordTypeEmail, ctx.ClientIP(),
		req.CaptchaID, req.CaptchaCode)
	if !captchaPass {
		resp := schema.UserVerifyEmailErrorResponse{
			Key:   "captcha_code",
			Value: "error.object.verification_failed",
		}
		resp.Value = translator.GlobalTrans.Tr(handler.GetLang(ctx), resp.Value)
		handler.HandleResponse(ctx, errors.BadRequest(reason.CaptchaVerificationFailed), resp)

		return
	}
	uc.actionService.ActionRecordAdd(ctx, schema.ActionRecordTypeEmail, ctx.ClientIP())
	err := uc.userService.UserVerifyEmailSend(ctx, userInfo.UserID)
	handler.HandleResponse(ctx, err, nil)
}

// UserModifyPassWord godoc
// @Summary UserModifyPassWord
// @Description UserModifyPassWord
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.UserModifyPassWordRequest  true "UserModifyPassWordRequest"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/user/password [put]
func (uc *UserController) UserModifyPassWord(ctx *gin.Context) {
	req := &schema.UserModifyPassWordRequest{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)

	oldPassVerification, err := uc.userService.UserModifyPassWordVerification(ctx, req)
	if err != nil {
		handler.HandleResponse(ctx, err, nil)
		return
	}
	if !oldPassVerification {
		resp := schema.UserVerifyEmailErrorResponse{
			Key:   "old_pass",
			Value: "error.object.old_password_verification_failed",
		}
		resp.Value = translator.GlobalTrans.Tr(handler.GetLang(ctx), resp.Value)
		handler.HandleResponse(ctx, errors.BadRequest(reason.CaptchaVerificationFailed), resp)
		return
	}
	if req.OldPass == req.Pass {

		resp := schema.UserVerifyEmailErrorResponse{
			Key:   "pass",
			Value: "error.object.new_password_same_as_previous_setting",
		}
		resp.Value = translator.GlobalTrans.Tr(handler.GetLang(ctx), resp.Value)
		handler.HandleResponse(ctx, errors.BadRequest(reason.CaptchaVerificationFailed), resp)
		return
	}
	err = uc.userService.UserModifyPassword(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// UserUpdateInfo update user info
// @Summary UserUpdateInfo update user info
// @Description UserUpdateInfo update user info
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "access-token"
// @Param data body schema.UpdateInfoRequest true "UpdateInfoRequest"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/user/info [put]
func (uc *UserController) UserUpdateInfo(ctx *gin.Context) {
	req := &schema.UpdateInfoRequest{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	err := uc.userService.UpdateInfo(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// UserUpdateInterface update user interface config
// @Summary UserUpdateInterface update user interface config
// @Description UserUpdateInterface update user interface config
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "access-token"
// @Param data body schema.UpdateUserInterfaceRequest true "UpdateInfoRequest"
// @Success 200 {object} handler.RespBody
// @Router /answer/api/v1/user/interface [put]
func (uc *UserController) UserUpdateInterface(ctx *gin.Context) {
	req := &schema.UpdateUserInterfaceRequest{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.UserId = middleware.GetLoginUserIDFromContext(ctx)
	err := uc.userService.UserUpdateInterface(ctx, req)
	handler.HandleResponse(ctx, err, nil)
}

// UploadUserAvatar godoc
// @Summary UserUpdateInfo
// @Description UserUpdateInfo
// @Tags User
// @Accept multipart/form-data
// @Security ApiKeyAuth
// @Param file formData file true "file"
// @Success 200 {object} handler.RespBody{data=string}
// @Router /answer/api/v1/user/avatar/upload [post]
func (uc *UserController) UploadUserAvatar(ctx *gin.Context) {
	// max size
	var filesMax int64 = 5 << 20
	var valuesMax int64 = 5
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, filesMax+valuesMax)
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Error(err.Error())
		handler.HandleResponse(ctx, errors.BadRequest(reason.RequestFormatError), nil)
		return
	}
	fileExt := strings.ToLower(path.Ext(header.Filename))
	if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".jpeg" {
		log.Errorf("upload file format is not supported: %s", fileExt)
		handler.HandleResponse(ctx, errors.BadRequest(reason.RequestFormatError), nil)
		return
	}

	url, err := uc.uploaderService.UploadAvatarFile(ctx, header, fileExt)
	handler.HandleResponse(ctx, err, url)
}

// UploadUserPostFile godoc
// @Summary upload user post file
// @Description upload user post file
// @Tags User
// @Accept multipart/form-data
// @Security ApiKeyAuth
// @Param file formData file true "file"
// @Success 200 {object} handler.RespBody{data=string}
// @Router /answer/api/v1/user/post/file [post]
func (uc *UserController) UploadUserPostFile(ctx *gin.Context) {
	// max size
	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, 10*1024*1024)
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Error(err.Error())
		handler.HandleResponse(ctx, errors.BadRequest(reason.RequestFormatError), nil)
		return
	}
	fileExt := strings.ToLower(path.Ext(header.Filename))
	if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".jpeg" {
		log.Errorf("upload file format is not supported: %s", fileExt)
		handler.HandleResponse(ctx, errors.BadRequest(reason.RequestFormatError), nil)
		return
	}

	url, err := uc.uploaderService.UploadPostFile(ctx, header, fileExt)
	handler.HandleResponse(ctx, err, url)
}

// ActionRecord godoc
// @Summary ActionRecord
// @Description ActionRecord
// @Tags User
// @Param action query string true "action" Enums(login, e_mail, find_pass)
// @Security ApiKeyAuth
// @Success 200 {object} handler.RespBody{data=schema.ActionRecordResp}
// @Router /answer/api/v1/user/action/record [get]
func (uc *UserController) ActionRecord(ctx *gin.Context) {
	req := &schema.ActionRecordReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.IP = ctx.ClientIP()

	resp, err := uc.actionService.ActionRecord(ctx, req)
	handler.HandleResponse(ctx, err, resp)
}

// UserNoticeSet godoc
// @Summary UserNoticeSet
// @Description UserNoticeSet
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.UserNoticeSetRequest true "UserNoticeSetRequest"
// @Success 200 {object} handler.RespBody{data=schema.UserNoticeSetResp}
// @Router /answer/api/v1/user/notice/set [post]
func (uc *UserController) UserNoticeSet(ctx *gin.Context) {
	req := &schema.UserNoticeSetRequest{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	resp, err := uc.userService.UserNoticeSet(ctx, req.UserID, req.NoticeSwitch)
	handler.HandleResponse(ctx, err, resp)
}

// UserChangeEmailSendCode send email to the user email then change their email
// @Summary send email to the user email then change their email
// @Description send email to the user email then change their email
// @Tags User
// @Accept json
// @Produce json
// @Param data body schema.UserChangeEmailSendCodeReq true "UserChangeEmailSendCodeReq"
// @Success 200 {object} handler.RespBody{}
// @Router /answer/api/v1/user/email/change/code [post]
func (uc *UserController) UserChangeEmailSendCode(ctx *gin.Context) {
	req := &schema.UserChangeEmailSendCodeReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	// If the user is not logged in, the api cannot be used.
	// If the user email is not verified, that also can use this api to modify the email.
	if len(req.UserID) == 0 {
		handler.HandleResponse(ctx, errors.Unauthorized(reason.UnauthorizedError), nil)
		return
	}

	captchaPass := uc.actionService.ActionRecordVerifyCaptcha(ctx, schema.ActionRecordTypeEmail, ctx.ClientIP(), req.CaptchaID, req.CaptchaCode)
	if !captchaPass {
		resp := schema.UserVerifyEmailErrorResponse{
			Key:   "captcha_code",
			Value: "error.object.verification_failed",
		}
		resp.Value = translator.GlobalTrans.Tr(handler.GetLang(ctx), resp.Value)
		handler.HandleResponse(ctx, errors.BadRequest(reason.CaptchaVerificationFailed), resp)
		return
	}
	_, _ = uc.actionService.ActionRecordAdd(ctx, schema.ActionRecordTypeEmail, ctx.ClientIP())
	resp, err := uc.userService.UserChangeEmailSendCode(ctx, req)
	if err != nil {
		if resp != nil {
			resp.Value = translator.GlobalTrans.Tr(handler.GetLang(ctx), resp.Value)
		}
		handler.HandleResponse(ctx, err, resp)
		return
	}
	handler.HandleResponse(ctx, err, nil)
}

// UserChangeEmailVerify user change email verification
// @Summary user change email verification
// @Description user change email verification
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param data body schema.UserChangeEmailVerifyReq true "UserChangeEmailVerifyReq"
// @Success 200 {object} handler.RespBody{}
// @Router /answer/api/v1/user/email [put]
func (uc *UserController) UserChangeEmailVerify(ctx *gin.Context) {
	req := &schema.UserChangeEmailVerifyReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}
	req.Content = uc.emailService.VerifyUrlExpired(ctx, req.Code)
	if len(req.Content) == 0 {
		handler.HandleResponse(ctx, errors.Forbidden(reason.EmailVerifyURLExpired),
			&schema.ForbiddenResp{Type: schema.ForbiddenReasonTypeURLExpired})
		return
	}

	err := uc.userService.UserChangeEmailVerify(ctx, req.Content)
	uc.actionService.ActionRecordDel(ctx, schema.ActionRecordTypeEmail, ctx.ClientIP())
	handler.HandleResponse(ctx, err, nil)
}
