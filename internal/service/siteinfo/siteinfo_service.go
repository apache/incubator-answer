package siteinfo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/handler"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/base/translator"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service/config"
	"github.com/answerdev/answer/internal/service/export"
	questioncommon "github.com/answerdev/answer/internal/service/question_common"
	"github.com/answerdev/answer/internal/service/siteinfo_common"
	tagcommon "github.com/answerdev/answer/internal/service/tag_common"
	"github.com/answerdev/answer/plugin"
	"github.com/jinzhu/copier"
	"github.com/segmentfault/pacman/errors"
	"github.com/segmentfault/pacman/log"
)

type SiteInfoService struct {
	siteInfoRepo          siteinfo_common.SiteInfoRepo
	siteInfoCommonService siteinfo_common.SiteInfoCommonService
	emailService          *export.EmailService
	tagCommonService      *tagcommon.TagCommonService
	configService         *config.ConfigService
	questioncommon        *questioncommon.QuestionCommon
}

func NewSiteInfoService(
	siteInfoRepo siteinfo_common.SiteInfoRepo,
	siteInfoCommonService siteinfo_common.SiteInfoCommonService,
	emailService *export.EmailService,
	tagCommonService *tagcommon.TagCommonService,
	configService *config.ConfigService,
	questioncommon *questioncommon.QuestionCommon,

) *SiteInfoService {
	plugin.RegisterGetSiteURLFunc(func() string {
		generalSiteInfo, err := siteInfoCommonService.GetSiteGeneral(context.Background())
		if err != nil {
			log.Error(err)
			return ""
		}
		return generalSiteInfo.SiteUrl
	})

	return &SiteInfoService{
		siteInfoRepo:          siteInfoRepo,
		siteInfoCommonService: siteInfoCommonService,
		emailService:          emailService,
		tagCommonService:      tagCommonService,
		configService:         configService,
		questioncommon:        questioncommon,
	}
}

// GetSiteGeneral get site info general
func (s *SiteInfoService) GetSiteGeneral(ctx context.Context) (resp *schema.SiteGeneralResp, err error) {
	return s.siteInfoCommonService.GetSiteGeneral(ctx)
}

// GetSiteInterface get site info interface
func (s *SiteInfoService) GetSiteInterface(ctx context.Context) (resp *schema.SiteInterfaceResp, err error) {
	return s.siteInfoCommonService.GetSiteInterface(ctx)
}

// GetSiteBranding get site info branding
func (s *SiteInfoService) GetSiteBranding(ctx context.Context) (resp *schema.SiteBrandingResp, err error) {
	return s.siteInfoCommonService.GetSiteBranding(ctx)
}

// GetSiteUsers get site info about users
func (s *SiteInfoService) GetSiteUsers(ctx context.Context) (resp *schema.SiteUsersResp, err error) {
	return s.siteInfoCommonService.GetSiteUsers(ctx)
}

// GetSiteWrite get site info write
func (s *SiteInfoService) GetSiteWrite(ctx context.Context) (resp *schema.SiteWriteResp, err error) {
	resp = &schema.SiteWriteResp{}
	siteInfo, exist, err := s.siteInfoRepo.GetByType(ctx, constant.SiteTypeWrite)
	if err != nil {
		log.Error(err)
		return resp, nil
	}
	if exist {
		_ = json.Unmarshal([]byte(siteInfo.Content), resp)
	}

	resp.RecommendTags, err = s.tagCommonService.GetSiteWriteRecommendTag(ctx)
	if err != nil {
		log.Error(err)
	}
	resp.ReservedTags, err = s.tagCommonService.GetSiteWriteReservedTag(ctx)
	if err != nil {
		log.Error(err)
	}
	return resp, nil
}

// GetSiteLegal get site legal info
func (s *SiteInfoService) GetSiteLegal(ctx context.Context) (resp *schema.SiteLegalResp, err error) {
	return s.siteInfoCommonService.GetSiteLegal(ctx)
}

// GetSiteLogin get site login info
func (s *SiteInfoService) GetSiteLogin(ctx context.Context) (resp *schema.SiteLoginResp, err error) {
	return s.siteInfoCommonService.GetSiteLogin(ctx)
}

// GetSiteCustomCssHTML get site custom css html config
func (s *SiteInfoService) GetSiteCustomCssHTML(ctx context.Context) (resp *schema.SiteCustomCssHTMLResp, err error) {
	return s.siteInfoCommonService.GetSiteCustomCssHTML(ctx)
}

// GetSiteTheme get site theme config
func (s *SiteInfoService) GetSiteTheme(ctx context.Context) (resp *schema.SiteThemeResp, err error) {
	return s.siteInfoCommonService.GetSiteTheme(ctx)
}

func (s *SiteInfoService) SaveSiteGeneral(ctx context.Context, req schema.SiteGeneralReq) (err error) {
	req.FormatSiteUrl()
	content, _ := json.Marshal(req)
	data := &entity.SiteInfo{
		Type:    constant.SiteTypeGeneral,
		Content: string(content),
		Status:  1,
	}
	return s.siteInfoRepo.SaveByType(ctx, constant.SiteTypeGeneral, data)
}

func (s *SiteInfoService) SaveSiteInterface(ctx context.Context, req schema.SiteInterfaceReq) (err error) {
	// check language
	if !translator.CheckLanguageIsValid(req.Language) {
		err = errors.BadRequest(reason.LangNotFound)
		return
	}

	content, _ := json.Marshal(req)
	data := entity.SiteInfo{
		Type:    constant.SiteTypeInterface,
		Content: string(content),
	}
	return s.siteInfoRepo.SaveByType(ctx, constant.SiteTypeInterface, &data)
}

// SaveSiteBranding save site branding information
func (s *SiteInfoService) SaveSiteBranding(ctx context.Context, req *schema.SiteBrandingReq) (err error) {
	content, _ := json.Marshal(req)
	data := &entity.SiteInfo{
		Type:    constant.SiteTypeBranding,
		Content: string(content),
		Status:  1,
	}
	return s.siteInfoRepo.SaveByType(ctx, constant.SiteTypeBranding, data)
}

// SaveSiteWrite save site configuration about write
func (s *SiteInfoService) SaveSiteWrite(ctx context.Context, req *schema.SiteWriteReq) (resp interface{}, err error) {
	errData, err := s.tagCommonService.SetSiteWriteTag(ctx, req.RecommendTags, req.ReservedTags, req.UserID)
	if err != nil {
		return errData, err
	}

	content, _ := json.Marshal(req)
	data := &entity.SiteInfo{
		Type:    constant.SiteTypeWrite,
		Content: string(content),
		Status:  1,
	}
	return nil, s.siteInfoRepo.SaveByType(ctx, constant.SiteTypeWrite, data)
}

// SaveSiteLegal save site legal configuration
func (s *SiteInfoService) SaveSiteLegal(ctx context.Context, req *schema.SiteLegalReq) (err error) {
	content, _ := json.Marshal(req)
	data := &entity.SiteInfo{
		Type:    constant.SiteTypeLegal,
		Content: string(content),
		Status:  1,
	}
	return s.siteInfoRepo.SaveByType(ctx, constant.SiteTypeLegal, data)
}

// SaveSiteLogin save site legal configuration
func (s *SiteInfoService) SaveSiteLogin(ctx context.Context, req *schema.SiteLoginReq) (err error) {
	content, _ := json.Marshal(req)
	data := &entity.SiteInfo{
		Type:    constant.SiteTypeLogin,
		Content: string(content),
		Status:  1,
	}
	return s.siteInfoRepo.SaveByType(ctx, constant.SiteTypeLogin, data)
}

// SaveSiteCustomCssHTML save site custom html configuration
func (s *SiteInfoService) SaveSiteCustomCssHTML(ctx context.Context, req *schema.SiteCustomCssHTMLReq) (err error) {
	content, _ := json.Marshal(req)
	data := &entity.SiteInfo{
		Type:    constant.SiteTypeCustomCssHTML,
		Content: string(content),
		Status:  1,
	}
	return s.siteInfoRepo.SaveByType(ctx, constant.SiteTypeCustomCssHTML, data)
}

// SaveSiteTheme save site custom html configuration
func (s *SiteInfoService) SaveSiteTheme(ctx context.Context, req *schema.SiteThemeReq) (err error) {
	content, _ := json.Marshal(req)
	data := &entity.SiteInfo{
		Type:    constant.SiteTypeTheme,
		Content: string(content),
		Status:  1,
	}
	return s.siteInfoRepo.SaveByType(ctx, constant.SiteTypeTheme, data)
}

// SaveSiteUsers save site users
func (s *SiteInfoService) SaveSiteUsers(ctx context.Context, req *schema.SiteUsersReq) (err error) {
	content, _ := json.Marshal(req)
	data := &entity.SiteInfo{
		Type:    constant.SiteTypeUsers,
		Content: string(content),
		Status:  1,
	}
	return s.siteInfoRepo.SaveByType(ctx, constant.SiteTypeUsers, data)
}

// GetSMTPConfig get smtp config
func (s *SiteInfoService) GetSMTPConfig(ctx context.Context) (resp *schema.GetSMTPConfigResp, err error) {
	emailConfig, err := s.emailService.GetEmailConfig(ctx)
	if err != nil {
		return nil, err
	}
	resp = &schema.GetSMTPConfigResp{}
	_ = copier.Copy(resp, emailConfig)
	return resp, nil
}

// UpdateSMTPConfig get smtp config
func (s *SiteInfoService) UpdateSMTPConfig(ctx context.Context, req *schema.UpdateSMTPConfigReq) (err error) {
	ec := &export.EmailConfig{}
	_ = copier.Copy(ec, req)

	err = s.emailService.SetEmailConfig(ctx, ec)
	if err != nil {
		return err
	}
	if len(req.TestEmailRecipient) > 0 {
		title, body, err := s.emailService.TestTemplate(ctx)
		if err != nil {
			return err
		}
		go s.emailService.SendAndSaveCode(ctx, req.TestEmailRecipient, title, body, "", "")
	}
	return nil
}

func (s *SiteInfoService) GetSeo(ctx context.Context) (resp *schema.SiteSeoReq, err error) {
	resp = &schema.SiteSeoReq{}
	if err = s.siteInfoCommonService.GetSiteInfoByType(ctx, constant.SiteTypeSeo, resp); err != nil {
		return resp, err
	}
	loginConfig, err := s.GetSiteLogin(ctx)
	if err != nil {
		log.Error(err)
		return resp, nil
	}
	// If the site is set to privacy mode, prohibit crawling any page.
	if loginConfig.LoginRequired {
		resp.Robots = "User-agent: *\nDisallow: /"
		return resp, nil
	}
	return resp, nil
}

func (s *SiteInfoService) SaveSeo(ctx context.Context, req schema.SiteSeoReq) (err error) {
	content, _ := json.Marshal(req)
	data := entity.SiteInfo{
		Type:    constant.SiteTypeSeo,
		Content: string(content),
	}
	return s.siteInfoRepo.SaveByType(ctx, constant.SiteTypeSeo, &data)
}

func (s *SiteInfoService) GetPrivilegesConfig(ctx context.Context) (resp *schema.GetPrivilegesConfigResp, err error) {
	privilege := &schema.UpdatePrivilegesConfigReq{}
	if err = s.siteInfoCommonService.GetSiteInfoByType(ctx, constant.SiteTypePrivileges, privilege); err != nil {
		return nil, err
	}
	resp = &schema.GetPrivilegesConfigResp{
		Options:       s.translatePrivilegeOptions(ctx),
		SelectedLevel: schema.PrivilegeLevel3,
	}
	if privilege != nil && privilege.Level > 0 {
		resp.SelectedLevel = privilege.Level
	}
	return resp, nil
}

func (s *SiteInfoService) translatePrivilegeOptions(ctx context.Context) (options []*schema.PrivilegeOption) {
	la := handler.GetLangByCtx(ctx)
	for _, option := range schema.DefaultPrivilegeOptions {
		op := &schema.PrivilegeOption{
			Level:     option.Level,
			LevelDesc: translator.Tr(la, option.LevelDesc),
		}
		for _, privilege := range option.Privileges {
			op.Privileges = append(op.Privileges, &constant.Privilege{
				Key:   privilege.Key,
				Label: translator.Tr(la, privilege.Label),
				Value: privilege.Value,
			})
		}
		options = append(options, op)
	}
	return
}

func (s *SiteInfoService) UpdatePrivilegesConfig(ctx context.Context, req *schema.UpdatePrivilegesConfigReq) (err error) {
	var chooseOption *schema.PrivilegeOption
	for _, option := range schema.DefaultPrivilegeOptions {
		if option.Level == req.Level {
			chooseOption = option
			break
		}
	}
	if chooseOption == nil {
		return nil
	}

	// update site info that user choose which privilege level
	content, _ := json.Marshal(req)
	data := &entity.SiteInfo{
		Type:    constant.SiteTypePrivileges,
		Content: string(content),
		Status:  1,
	}
	err = s.siteInfoRepo.SaveByType(ctx, constant.SiteTypePrivileges, data)
	if err != nil {
		return err
	}

	// update privilege in config
	for _, privilege := range chooseOption.Privileges {
		err = s.configService.UpdateConfig(ctx, privilege.Key, fmt.Sprintf("%d", privilege.Value))
		if err != nil {
			return err
		}
	}
	return
}
