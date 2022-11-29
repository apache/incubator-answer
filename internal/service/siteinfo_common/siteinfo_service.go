package siteinfo_common

import (
	"context"
	"encoding/json"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/schema"
)

type SiteInfoCommonService struct {
	siteInfoRepo SiteInfoRepo
}

func NewSiteInfoCommonService(siteInfoRepo SiteInfoRepo) *SiteInfoCommonService {
	return &SiteInfoCommonService{
		siteInfoRepo: siteInfoRepo,
	}
}

// GetSiteGeneral get site info general
func (s *SiteInfoCommonService) GetSiteGeneral(ctx context.Context) (resp *schema.SiteGeneralResp, err error) {
	resp = &schema.SiteGeneralResp{}
	siteInfo, exist, err := s.siteInfoRepo.GetByType(ctx, constant.SiteTypeGeneral)
	if err != nil {
		return resp, err
	}
	if !exist {
		return resp, nil
	}
	_ = json.Unmarshal([]byte(siteInfo.Content), resp)
	return resp, nil
}

// GetSiteInterface get site info interface
func (s *SiteInfoCommonService) GetSiteInterface(ctx context.Context) (resp *schema.SiteInterfaceResp, err error) {
	resp = &schema.SiteInterfaceResp{}
	siteInfo, exist, err := s.siteInfoRepo.GetByType(ctx, constant.SiteTypeInterface)
	if err != nil {
		return resp, err
	}
	if !exist {
		return resp, nil
	}
	_ = json.Unmarshal([]byte(siteInfo.Content), resp)
	return resp, nil
}

// GetSiteBranding get site info branding
func (s *SiteInfoCommonService) GetSiteBranding(ctx context.Context) (resp *schema.SiteBrandingResp, err error) {
	resp = &schema.SiteBrandingResp{}
	siteInfo, exist, err := s.siteInfoRepo.GetByType(ctx, constant.SiteTypeBranding)
	if err != nil {
		return resp, err
	}
	if !exist {
		return resp, nil
	}
	_ = json.Unmarshal([]byte(siteInfo.Content), resp)
	return resp, nil
}

// GetSiteWrite get site info write
func (s *SiteInfoCommonService) GetSiteWrite(ctx context.Context) (resp *schema.SiteWriteResp, err error) {
	resp = &schema.SiteWriteResp{}
	siteInfo, exist, err := s.siteInfoRepo.GetByType(ctx, constant.SiteTypeWrite)
	if err != nil {
		return resp, err
	}
	if !exist {
		return resp, nil
	}
	_ = json.Unmarshal([]byte(siteInfo.Content), resp)
	return resp, nil
}

// GetSiteLegal get site info write
func (s *SiteInfoCommonService) GetSiteLegal(ctx context.Context) (resp *schema.SiteLegalResp, err error) {
	resp = &schema.SiteLegalResp{}
	siteInfo, exist, err := s.siteInfoRepo.GetByType(ctx, constant.SiteTypeLegal)
	if err != nil {
		return nil, err
	}
	if !exist {
		return resp, nil
	}
	_ = json.Unmarshal([]byte(siteInfo.Content), resp)
	return resp, nil
}
