package siteinfo_common

import (
	"context"
	"encoding/json"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/pkg/gravatar"
	"github.com/answerdev/answer/pkg/uid"
	"github.com/segmentfault/pacman/log"
)

//go:generate mockgen -source=./siteinfo_service.go -destination=../mock/siteinfo_repo_mock.go -package=mock
type SiteInfoRepo interface {
	SaveByType(ctx context.Context, siteType string, data *entity.SiteInfo) (err error)
	GetByType(ctx context.Context, siteType string) (siteInfo *entity.SiteInfo, exist bool, err error)
}

// SiteInfoCommonService site info common service
type SiteInfoCommonService struct {
	siteInfoRepo SiteInfoRepo
}

// NewSiteInfoCommonService new site info common service
func NewSiteInfoCommonService(siteInfoRepo SiteInfoRepo) *SiteInfoCommonService {
	siteInfo := &SiteInfoCommonService{
		siteInfoRepo: siteInfoRepo,
	}
	seoinfo, err := siteInfo.GetSiteSeo(context.Background())
	if err != nil {
		log.Error("seoinfo error", err)
	}
	if seoinfo.PermaLink == schema.PermaLinkQuestionIDAndTitleByShortID || seoinfo.PermaLink == schema.PermaLinkQuestionIDByShortID {
		uid.ShortIDSwitch = true
	} else {
		uid.ShortIDSwitch = false
	}

	return siteInfo
}

// GetSiteGeneral get site info general
func (s *SiteInfoCommonService) GetSiteGeneral(ctx context.Context) (resp *schema.SiteGeneralResp, err error) {
	resp = &schema.SiteGeneralResp{}
	if err = s.GetSiteInfoByType(ctx, constant.SiteTypeGeneral, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSiteInterface get site info interface
func (s *SiteInfoCommonService) GetSiteInterface(ctx context.Context) (resp *schema.SiteInterfaceResp, err error) {
	resp = &schema.SiteInterfaceResp{}
	if err = s.GetSiteInfoByType(ctx, constant.SiteTypeInterface, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSiteBranding get site info branding
func (s *SiteInfoCommonService) GetSiteBranding(ctx context.Context) (resp *schema.SiteBrandingResp, err error) {
	resp = &schema.SiteBrandingResp{}
	if err = s.GetSiteInfoByType(ctx, constant.SiteTypeBranding, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSiteUsers get site info about users
func (s *SiteInfoCommonService) GetSiteUsers(ctx context.Context) (resp *schema.SiteUsersResp, err error) {
	resp = &schema.SiteUsersResp{}
	if err = s.GetSiteInfoByType(ctx, constant.SiteTypeUsers, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// FormatAvatar format avatar
func (s *SiteInfoCommonService) FormatAvatar(ctx context.Context, originalAvatarData, email string) *schema.AvatarInfo {
	gravatarBaseURL, defaultAvatar := s.getAvatarDefaultConfig(ctx)
	return s.selectedAvatar(originalAvatarData, defaultAvatar, gravatarBaseURL, email)
}

// FormatListAvatar format avatar
func (s *SiteInfoCommonService) FormatListAvatar(ctx context.Context, userList []*entity.User) (
	avatarMapping map[string]*schema.AvatarInfo) {
	gravatarBaseURL, defaultAvatar := s.getAvatarDefaultConfig(ctx)
	avatarMapping = make(map[string]*schema.AvatarInfo)
	for _, user := range userList {
		avatarMapping[user.ID] = s.selectedAvatar(user.Avatar, defaultAvatar, gravatarBaseURL, user.EMail)
	}
	return avatarMapping
}

func (s *SiteInfoCommonService) getAvatarDefaultConfig(ctx context.Context) (string, string) {
	gravatarBaseURL, defaultAvatar := constant.DefaultGravatarBaseURL, constant.DefaultAvatar
	usersConfig, err := s.GetSiteUsers(ctx)
	if err != nil {
		log.Error(err)
	} else {
		gravatarBaseURL = usersConfig.GravatarBaseURL
		defaultAvatar = usersConfig.DefaultAvatar
	}
	return gravatarBaseURL, defaultAvatar
}

func (s *SiteInfoCommonService) selectedAvatar(
	originalAvatarData string, defaultAvatar string, gravatarBaseURL string, email string) *schema.AvatarInfo {
	avatarInfo := &schema.AvatarInfo{}
	_ = json.Unmarshal([]byte(originalAvatarData), avatarInfo)

	if len(avatarInfo.Type) == 0 && defaultAvatar == constant.AvatarTypeGravatar {
		avatarInfo.Type = constant.AvatarTypeGravatar
		avatarInfo.Gravatar = gravatar.GetAvatarURL(gravatarBaseURL, email)
	} else if avatarInfo.Type == constant.AvatarTypeGravatar {
		avatarInfo.Gravatar = gravatar.GetAvatarURL(gravatarBaseURL, email)
	}
	return avatarInfo
}

// GetSiteWrite get site info write
func (s *SiteInfoCommonService) GetSiteWrite(ctx context.Context) (resp *schema.SiteWriteResp, err error) {
	resp = &schema.SiteWriteResp{}
	if err = s.GetSiteInfoByType(ctx, constant.SiteTypeWrite, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSiteLegal get site info write
func (s *SiteInfoCommonService) GetSiteLegal(ctx context.Context) (resp *schema.SiteLegalResp, err error) {
	resp = &schema.SiteLegalResp{}
	if err = s.GetSiteInfoByType(ctx, constant.SiteTypeLegal, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSiteLogin get site login config
func (s *SiteInfoCommonService) GetSiteLogin(ctx context.Context) (resp *schema.SiteLoginResp, err error) {
	resp = &schema.SiteLoginResp{}
	if err = s.GetSiteInfoByType(ctx, constant.SiteTypeLogin, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSiteCustomCssHTML get site custom css html config
func (s *SiteInfoCommonService) GetSiteCustomCssHTML(ctx context.Context) (resp *schema.SiteCustomCssHTMLResp, err error) {
	resp = &schema.SiteCustomCssHTMLResp{}
	if err = s.GetSiteInfoByType(ctx, constant.SiteTypeCustomCssHTML, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSiteTheme get site theme
func (s *SiteInfoCommonService) GetSiteTheme(ctx context.Context) (resp *schema.SiteThemeResp, err error) {
	resp = &schema.SiteThemeResp{
		ThemeOptions: schema.GetThemeOptions,
	}
	if err = s.GetSiteInfoByType(ctx, constant.SiteTypeTheme, resp); err != nil {
		return nil, err
	}
	resp.TrTheme(ctx)
	return resp, nil
}

// GetSiteSeo get site seo
func (s *SiteInfoCommonService) GetSiteSeo(ctx context.Context) (resp *schema.SiteSeoReq, err error) {
	resp = &schema.SiteSeoReq{}
	if err = s.GetSiteInfoByType(ctx, constant.SiteTypeSeo, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SiteInfoCommonService) GetSiteInfoByType(ctx context.Context, siteType string, resp interface{}) (err error) {
	siteInfo, exist, err := s.siteInfoRepo.GetByType(ctx, siteType)
	if err != nil {
		return err
	}
	if !exist {
		return nil
	}
	_ = json.Unmarshal([]byte(siteInfo.Content), resp)
	return nil
}
