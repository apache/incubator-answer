package migrations

import (
	"encoding/json"
	"fmt"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
	"github.com/tidwall/gjson"
	"xorm.io/xorm"
)

func addLoginLimitations(x *xorm.Engine) error {
	loginSiteInfo := &entity.SiteInfo{
		Type: constant.SiteTypeLogin,
	}
	exist, err := x.Get(loginSiteInfo)
	if err != nil {
		return fmt.Errorf("get config failed: %w", err)
	}
	if exist {
		content := &schema.SiteLoginReq{}
		_ = json.Unmarshal([]byte(loginSiteInfo.Content), content)
		content.AllowEmailRegistrations = true
		content.AllowEmailDomains = make([]string, 0)
		_, err = x.ID(loginSiteInfo.ID).Cols("content").Update(loginSiteInfo)
		if err != nil {
			return fmt.Errorf("update site info failed: %w", err)
		}
	}

	interfaceSiteInfo := &entity.SiteInfo{
		Type: constant.SiteTypeInterface,
	}
	exist, err = x.Get(interfaceSiteInfo)
	if err != nil {
		return fmt.Errorf("get config failed: %w", err)
	}
	siteUsers := &schema.SiteUsersReq{
		AllowUpdateDisplayName: true,
		AllowUpdateUsername:    true,
		AllowUpdateAvatar:      true,
		AllowUpdateBio:         true,
		AllowUpdateWebsite:     true,
		AllowUpdateLocation:    true,
	}
	if exist {
		siteUsers.DefaultAvatar = gjson.Get(interfaceSiteInfo.Content, "default_avatar").String()
	}
	data, _ := json.Marshal(siteUsers)

	exist, err = x.Get(&entity.SiteInfo{Type: constant.SiteTypeUsers})
	if err != nil {
		return fmt.Errorf("get config failed: %w", err)
	}
	if !exist {
		usersSiteInfo := &entity.SiteInfo{
			Type:    constant.SiteTypeUsers,
			Content: string(data),
			Status:  1,
		}
		_, err = x.InsertOne(usersSiteInfo)
		if err != nil {
			return fmt.Errorf("insert site info failed: %w", err)
		}
	}
	return nil
}
