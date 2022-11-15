package siteinfo_common

import (
	"context"
	"encoding/json"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/base/reason"
	"github.com/answerdev/answer/internal/schema"
	"github.com/segmentfault/pacman/errors"
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
	siteInfo, exist, err := s.siteInfoRepo.GetByType(ctx, constant.SiteTypeGeneral)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.BadRequest(reason.SiteInfoNotFound)
	}

	resp = &schema.SiteGeneralResp{}
	_ = json.Unmarshal([]byte(siteInfo.Content), resp)
	return resp, nil
}

// GetSiteInterface get site info interface
func (s *SiteInfoCommonService) GetSiteInterface(ctx context.Context) (resp *schema.SiteInterfaceResp, err error) {
	siteInfo, exist, err := s.siteInfoRepo.GetByType(ctx, constant.SiteTypeInterface)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.BadRequest(reason.SiteInfoNotFound)
	}
	resp = &schema.SiteInterfaceResp{}
	_ = json.Unmarshal([]byte(siteInfo.Content), resp)
	return resp, nil
}

// GetSiteBranding get site info branding
func (s *SiteInfoCommonService) GetSiteBranding(ctx context.Context) (resp *schema.SiteBrandingResp, err error) {
	siteInfo, exist, err := s.siteInfoRepo.GetByType(ctx, constant.SiteTypeBranding)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.BadRequest(reason.SiteInfoNotFound)
	}
	resp = &schema.SiteBrandingResp{}
	_ = json.Unmarshal([]byte(siteInfo.Content), resp)
	return resp, nil
}

// GetSiteWrite get site info write
func (s *SiteInfoCommonService) GetSiteWrite(ctx context.Context) (resp *schema.SiteWriteResp, err error) {
	siteInfo, exist, err := s.siteInfoRepo.GetByType(ctx, constant.SiteTypeWrite)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.BadRequest(reason.SiteInfoNotFound)
	}
	resp = &schema.SiteWriteResp{}
	_ = json.Unmarshal([]byte(siteInfo.Content), resp)
	return resp, nil
}
