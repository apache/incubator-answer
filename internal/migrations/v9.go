package migrations

import (
	"encoding/json"
	"fmt"

	"github.com/answerdev/answer/internal/base/constant"
	"github.com/answerdev/answer/internal/entity"
	"github.com/answerdev/answer/internal/schema"
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
		_, err = x.ID(loginSiteInfo.ID).Cols("content").Update(loginSiteInfo)
		if err != nil {
			return fmt.Errorf("update site info failed: %w", err)
		}
	}
	return nil
}
