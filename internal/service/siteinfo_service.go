package service

import (
	"context"
	"encoding/json"

	"github.com/segmentfault/answer/internal/base/reason"
	"github.com/segmentfault/answer/internal/entity"
	"github.com/segmentfault/answer/internal/schema"
	"github.com/segmentfault/answer/internal/service/siteinfo_common"
	"github.com/segmentfault/pacman/errors"
)

type SiteInfoService struct {
	siteInfoRepo siteinfo_common.SiteInfoRepo
}

func NewSiteInfoService(siteInfoRepo siteinfo_common.SiteInfoRepo) *SiteInfoService {
	return &SiteInfoService{
		siteInfoRepo: siteInfoRepo,
	}
}

func (s *SiteInfoService) GetSiteGeneral(ctx context.Context) (resp schema.SiteGeneralResp, err error) {
	var (
		siteType = "general"
		siteInfo *entity.SiteInfo
		exist    bool
	)
	resp = schema.SiteGeneralResp{}

	siteInfo, exist, err = s.siteInfoRepo.GetByType(ctx, siteType)
	if !exist {
		return
	}

	_ = json.Unmarshal([]byte(siteInfo.Content), &resp)
	return
}

func (s *SiteInfoService) GetSiteInterface(ctx context.Context) (resp schema.SiteInterfaceResp, err error) {
	var (
		siteType = "interface"
		siteInfo *entity.SiteInfo
		exist    bool
	)
	resp = schema.SiteInterfaceResp{}

	siteInfo, exist, err = s.siteInfoRepo.GetByType(ctx, siteType)
	if !exist {
		return
	}

	_ = json.Unmarshal([]byte(siteInfo.Content), &resp)
	return
}

func (s *SiteInfoService) SaveSiteGeneral(ctx context.Context, req schema.SiteGeneralReq) (err error) {
	var (
		siteType = "general"
		content  []byte
	)
	content, err = json.Marshal(req)

	data := entity.SiteInfo{
		Type:    siteType,
		Content: string(content),
	}

	err = s.siteInfoRepo.SaveByType(ctx, siteType, &data)
	return
}

func (s *SiteInfoService) SaveSiteInterface(ctx context.Context, req schema.SiteInterfaceReq) (err error) {
	var (
		siteType = "interface"
		themeExist,
		langExist bool
		content []byte
	)

	// check theme
	for _, theme := range schema.GetThemeOptions {
		if theme.Value == req.Theme {
			themeExist = true
			break
		}
	}
	if !themeExist {
		err = errors.BadRequest(reason.ThemeNotFound)
		return
	}

	// check language
	for _, lang := range schema.GetLangOptions {
		if lang.Value == req.Language {
			langExist = true
			break
		}
	}
	if !langExist {
		err = errors.BadRequest(reason.LangNotFound)
		return
	}

	content, err = json.Marshal(req)

	data := entity.SiteInfo{
		Type:    siteType,
		Content: string(content),
	}

	err = s.siteInfoRepo.SaveByType(ctx, siteType, &data)
	return
}
