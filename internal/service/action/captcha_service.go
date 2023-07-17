package action

import (
	"context"
	"image/color"
	"strings"

	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/mojocn/base64Captcha"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

// CaptchaRepo captcha repository
type CaptchaRepo interface {
	SetCaptcha(ctx context.Context, key, captcha string) (err error)
	GetCaptcha(ctx context.Context, key string) (captcha string, err error)
	DelCaptcha(ctx context.Context, key string) (err error)
	SetActionType(ctx context.Context, unit, actionType, config string, amount int) (err error)
	GetActionType(ctx context.Context, unit, actionType string) (actioninfo *entity.ActionRecordInfo, err error)
	DelActionType(ctx context.Context, unit, actionType string) (err error)
}

// CaptchaService kit service
type CaptchaService struct {
	captchaRepo CaptchaRepo
}

// NewCaptchaService captcha service
func NewCaptchaService(captchaRepo CaptchaRepo) *CaptchaService {
	return &CaptchaService{
		captchaRepo: captchaRepo,
	}
}

// ActionRecord action record
func (cs *CaptchaService) ActionRecord(ctx context.Context, req *schema.ActionRecordReq) (resp *schema.ActionRecordResp, err error) {
	resp = &schema.ActionRecordResp{}
	verificationResult := cs.ValidationStrategy(ctx, req.IP, req.Action)
	if !verificationResult {
		resp.CaptchaID, resp.CaptchaImg, err = cs.GenerateCaptcha(ctx)
		resp.Verify = true
	}
	return
}

func (cs *CaptchaService) UserRegisterCaptcha(ctx context.Context) (resp *schema.ActionRecordResp, err error) {
	resp = &schema.ActionRecordResp{}
	resp.CaptchaID, resp.CaptchaImg, err = cs.GenerateCaptcha(ctx)
	resp.Verify = true
	return
}

func (cs *CaptchaService) UserRegisterVerifyCaptcha(
	ctx context.Context, id string, VerifyValue string,
) bool {
	if id == "" || VerifyValue == "" {
		return false
	}
	pass, err := cs.VerifyCaptcha(ctx, id, VerifyValue)
	if err != nil {
		return false
	}
	return pass
}

// ActionRecordVerifyCaptcha
// Verify that you need to enter a CAPTCHA, and that the CAPTCHA is correct
func (cs *CaptchaService) ActionRecordVerifyCaptcha(
	ctx context.Context, actionType string, unit string, id string, VerifyValue string,
) bool {
	verificationResult := cs.ValidationStrategy(ctx, unit, actionType)
	if !verificationResult {
		if id == "" || VerifyValue == "" {
			return false
		}
		pass, err := cs.VerifyCaptcha(ctx, id, VerifyValue)
		if err != nil {
			return false
		}
		return pass
	}
	return true
}

func (cs *CaptchaService) ActionRecordAdd(ctx context.Context, actionType string, unit string) (int, error) {
	var err error
	info, cahceErr := cs.captchaRepo.GetActionType(ctx, unit, actionType)
	if cahceErr != nil {
		log.Error(err)
	}
	info.Num++
	err = cs.captchaRepo.SetActionType(ctx, unit, actionType, "", info.Num)
	if err != nil {
		return 0, err
	}
	return info.Num, nil
}

func (cs *CaptchaService) ActionRecordDel(ctx context.Context, actionType string, unit string) {
	err := cs.captchaRepo.DelActionType(ctx, unit, actionType)
	if err != nil {
		log.Error(err)
	}
}

// GenerateCaptcha generate captcha
func (cs *CaptchaService) GenerateCaptcha(ctx context.Context) (key, captchaBase64 string, err error) {
	driverString := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor:         &color.RGBA{R: 211, G: 211, B: 211, A: 0},
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	driver := driverString.ConvertFonts()

	id, content, answer := driver.GenerateIdQuestionAnswer()
	item, err := driver.DrawCaptcha(content)
	if err != nil {
		return "", "", errors.InternalServer(reason.UnknownError).WithError(err).WithStack()
	}
	err = cs.captchaRepo.SetCaptcha(ctx, id, answer)
	if err != nil {
		return "", "", err
	}

	captchaBase64 = item.EncodeB64string()
	return id, captchaBase64, nil
}

// VerifyCaptcha generate captcha
func (cs *CaptchaService) VerifyCaptcha(ctx context.Context, key, captcha string) (isCorrect bool, err error) {
	realCaptcha, err := cs.captchaRepo.GetCaptcha(ctx, key)
	if err != nil {
		log.Error("VerifyCaptcha GetCaptcha Error", err.Error())
		return false, nil
	}
	err = cs.captchaRepo.DelCaptcha(ctx, key)
	if err != nil {
		log.Error("VerifyCaptcha DelCaptcha Error", err.Error())
		return false, nil
	}
	return strings.TrimSpace(captcha) == realCaptcha, nil
}
