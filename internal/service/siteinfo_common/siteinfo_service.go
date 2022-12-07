package siteinfo_common

import (
	"context"
	"encoding/json"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
)

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
	return &SiteInfoCommonService{
		siteInfoRepo: siteInfoRepo,
	}
}

// GetSiteGeneral get site info general
func (s *SiteInfoCommonService) GetSiteGeneral(ctx context.Context) (resp *schema.SiteGeneralResp, err error) {
	resp = &schema.SiteGeneralResp{}
	if err = s.getSiteInfoByType(ctx, constant.SiteTypeGeneral, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSiteInterface get site info interface
func (s *SiteInfoCommonService) GetSiteInterface(ctx context.Context) (resp *schema.SiteInterfaceResp, err error) {
	resp = &schema.SiteInterfaceResp{}
	if err = s.getSiteInfoByType(ctx, constant.SiteTypeInterface, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSiteBranding get site info branding
func (s *SiteInfoCommonService) GetSiteBranding(ctx context.Context) (resp *schema.SiteBrandingResp, err error) {
	resp = &schema.SiteBrandingResp{}
	if err = s.getSiteInfoByType(ctx, constant.SiteTypeBranding, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSiteWrite get site info write
func (s *SiteInfoCommonService) GetSiteWrite(ctx context.Context) (resp *schema.SiteWriteResp, err error) {
	resp = &schema.SiteWriteResp{}
	if err = s.getSiteInfoByType(ctx, constant.SiteTypeWrite, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSiteLegal get site info write
func (s *SiteInfoCommonService) GetSiteLegal(ctx context.Context) (resp *schema.SiteLegalResp, err error) {
	resp = &schema.SiteLegalResp{}
	if err = s.getSiteInfoByType(ctx, constant.SiteTypeLegal, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GetSiteLogin get site login config
func (s *SiteInfoCommonService) GetSiteLogin(ctx context.Context) (resp *schema.SiteLoginResp, err error) {
	resp = &schema.SiteLoginResp{}
	if err = s.getSiteInfoByType(ctx, constant.SiteTypeLogin, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SiteInfoCommonService) getSiteInfoByType(ctx context.Context, siteType string, resp interface{}) (err error) {
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
